import argparse
import hashlib
import json
import math
import os
from pathlib import Path
import platform
import statistics
import subprocess
import tempfile
import time
import zipfile


def measure(binary, cohort, passes, environment, root, include_results=False):
    command = [str(binary), "-json", "-cohort", cohort, "-passes", str(passes)]
    if include_results:
        command.append("-results")
    started = time.perf_counter()
    ready = result = latency = None
    output = []
    with subprocess.Popen(
        command, cwd=root, env=environment, stdout=subprocess.PIPE,
        stderr=subprocess.STDOUT, text=True, encoding="utf-8",
    ) as process:
        for line in process.stdout:
            output.append(line)
            if line.startswith("BENCHMARK_READY "):
                latency = (time.perf_counter() - started) * 1000
                ready = json.loads(line.partition(" ")[2])
            elif line.startswith("BENCHMARK_RESULT "):
                result = json.loads(line.partition(" ")[2])
        status = process.wait()
    if status != 0 or ready is None or result is None:
        raise RuntimeError(f"Benchmark failed ({status}):\n{''.join(output[-20:])}")
    if ready != result["metadata"] or ready["cohort"] != cohort:
        raise RuntimeError("Benchmark metadata changed")
    if len(result["pass_ms"]) != passes or any(
        not math.isfinite(value) or value <= 0 for value in result["pass_ms"]
    ):
        raise RuntimeError("Invalid benchmark samples")
    return result | {"launch_to_ready_ms": latency}


def summarize(records):
    summary = {}
    for cohort in sorted({record["cohort"] for record in records}):
        summary[cohort] = {}
        for engine in ("baseline", "candidate"):
            selected = [record for record in records
                        if record["cohort"] == cohort and record["engine"] == engine]
            medians = [statistics.median(record["pass_ms"]) for record in selected]
            summary[cohort][engine] = {
                "cases": selected[0]["metadata"]["cases"],
                "parsed": selected[0]["metadata"]["parsed"],
                "warm_pass_median_ms": statistics.median(medians),
                "warm_pass_median_range_ms": [min(medians), max(medians)],
                "launch_to_ready_median_ms": statistics.median(
                    record["launch_to_ready_ms"] for record in selected
                ),
                "retained_heap_median_bytes": statistics.median(
                    record["retained_heap_bytes"] for record in selected
                ),
            }
        summary[cohort]["baseline_over_candidate_time"] = (
            summary[cohort]["baseline"]["warm_pass_median_ms"]
            / summary[cohort]["candidate"]["warm_pass_median_ms"]
        )
    return summary


def main():
    parser = argparse.ArgumentParser(description="Single-core Go version comparison")
    parser.add_argument("--baseline", default="v1.4.3")
    parser.add_argument("--candidate-version", help="use a published version instead of this checkout")
    parser.add_argument("--runs", type=int, default=6)
    parser.add_argument("--passes", type=int, default=8)
    parser.add_argument("--cpu", type=int, default=2)
    parser.add_argument("--cohort", action="append", choices=("auto", "explicit", "htmldate"))
    parser.add_argument("--output", type=Path, default=Path(".benchmarks/go-versions.json"))
    arguments = parser.parse_args()
    if arguments.runs < 2 or arguments.runs % 2 or arguments.passes < 1:
        parser.error("use an even number of runs >= 2 and positive passes")
    if not hasattr(os, "sched_getaffinity") or arguments.cpu not in os.sched_getaffinity(0):
        parser.error("run under Linux/WSL and select an available CPU")

    root = Path(__file__).resolve().parents[2]
    runner = Path(__file__).with_name("main.go").resolve()
    build_directory = root / ".benchmarks/runners"
    build_directory.mkdir(parents=True, exist_ok=True)
    environment = os.environ | {
        "GOTOOLCHAIN": "go1.27.1", "GOMAXPROCS": "1", "CGO_ENABLED": "0",
        "GOAMD64": "v1", "GOFLAGS": "", "GOEXPERIMENT": "", "GODEBUG": "",
        "GOGC": "100", "GOMEMLIMIT": "off", "TZ": "UTC", "LC_ALL": "C.UTF-8",
    }

    def module(version):
        record = json.loads(subprocess.check_output(
            ["go", "mod", "download", "-json", f"github.com/markusmobius/go-dateparser@{version}"],
            cwd=root, env=environment, text=True,
        ))
        if "Error" in record or not record.get("Sum"):
            raise RuntimeError(f"Could not verify module {version}: {record}")
        return Path(record["Dir"]), record

    baseline_root, baseline = module(arguments.baseline)
    candidate_root, candidate = (
        module(arguments.candidate_version) if arguments.candidate_version else
        (root, {
            "checkout": str(root),
            "commit": subprocess.check_output(
                ["git", "rev-parse", "HEAD"], cwd=root, text=True,
            ).strip(),
            "tracked_diff_sha256": hashlib.sha256(subprocess.check_output(
                ["git", "diff", "HEAD"], cwd=root,
            )).hexdigest(),
        })
    )
    binaries = {}
    for engine, directory, source in (
        ("baseline", baseline_root, baseline), ("candidate", candidate_root, candidate)
    ):
        with tempfile.TemporaryDirectory(dir=build_directory) as temporary:
            if "Zip" in source:
                with zipfile.ZipFile(source["Zip"]) as archive:
                    archive.extractall(temporary)
                directory = Path(temporary) / f"{source['Path']}@{source['Version']}"
            overlay = build_directory / f"{engine}-overlay.json"
            overlay.write_text(json.dumps({"Replace": {
                str(directory / "scripts/speedtest/main.go"): str(runner),
            }}), encoding="utf-8")
            binary = build_directory / engine
            subprocess.run(
                ["go", "build", "-mod=readonly", "-trimpath", "-buildvcs=false",
                 "-overlay", str(overlay), "-o", str(binary), "./scripts/speedtest"],
                cwd=directory, env=environment, check=True,
            )
            binaries[engine] = binary

    os.sched_setaffinity(0, {arguments.cpu})
    records = []
    reference_results = {}
    for cohort in arguments.cohort or ("auto", "explicit", "htmldate"):
        expected = measure(binaries["baseline"], cohort, 1, environment, root, True)
        actual = measure(binaries["candidate"], cohort, 1, environment, root, True)
        if actual["results"] != expected["results"]:
            differences = [
                {"index": index, "baseline": before, "candidate": after}
                for index, (before, after) in enumerate(zip(expected["results"], actual["results"]))
                if before != after
            ]
            raise RuntimeError(f"Exact result mismatch: {differences[:3]}")
        identity = {key: value for key, value in expected["metadata"].items()
                    if key != "first_pass_ms"}
        reference_results[cohort] = expected["results"]
        print(f"{cohort}: exact results match; preflights discarded", flush=True)
        for round_index in range(arguments.runs):
            order = ("baseline", "candidate") if round_index % 2 == 0 else ("candidate", "baseline")
            for engine in order:
                record = measure(binaries[engine], cohort, arguments.passes, environment, root)
                observed = {key: value for key, value in record["metadata"].items()
                            if key != "first_pass_ms"}
                if observed != identity:
                    raise RuntimeError(f"Workload or results changed: {engine}/{cohort}")
                record.update(engine=engine, cohort=cohort, round=round_index + 1)
                records.append(record)
                print(
                    f"{cohort} {engine} {round_index + 1}/{arguments.runs}: "
                    f"warm {statistics.median(record['pass_ms']):.3f} ms, "
                    f"heap {record['retained_heap_bytes'] / 2**20:.2f} MiB", flush=True,
                )

    report = {
        "date": time.strftime("%Y-%m-%d"), "platform": platform.platform(),
        "cpu_model": next(line.partition(":")[2].strip()
                          for line in Path("/proc/cpuinfo").read_text().splitlines()
                          if line.startswith("model name")),
        "cpu": arguments.cpu,
        "go_version": subprocess.check_output(["go", "version"], env=environment, text=True).strip(),
        "baseline": baseline, "candidate": candidate,
        "runs_per_engine": arguments.runs, "passes_per_run": arguments.passes,
        "runner_sha256": hashlib.sha256(runner.read_bytes()).hexdigest(),
        "binary_sha256": {engine: hashlib.sha256(binary.read_bytes()).hexdigest()
                          for engine, binary in binaries.items()},
        "summary": summarize(records), "runs": records, "reference_results": reference_results,
    }
    destination = arguments.output
    if not destination.is_absolute():
        destination = root / destination
    destination.parent.mkdir(parents=True, exist_ok=True)
    destination.write_text(json.dumps(report, indent=2) + "\n", encoding="utf-8")
    print(json.dumps(report["summary"], indent=2))
    print(f"Report: {destination}")


if __name__ == "__main__":
    main()