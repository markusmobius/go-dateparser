import argparse
import hashlib
import os
from pathlib import Path
import subprocess
import sys


PARSE_COHORTS = ("auto", "explicit", "htmldate")
FEATURE_COHORTS = ("search-auto", "search-split", "search-ngram", "time-span", "jalali", "hijri")


def local_reference(root, source):
    source = source.resolve()
    if not (source / "tools/benchmark.py").is_file():
        raise RuntimeError(f"Suite checkout has no shared benchmark runner: {source}")
    fixture = Path("testdata/python-features.json")
    expected = hashlib.sha256((root / fixture).read_bytes()).hexdigest()
    actual = hashlib.sha256((source / fixture).read_bytes()).hexdigest()
    if actual != expected:
        raise RuntimeError("Go and shared-suite Python feature fixtures must be byte-identical")
    return source


def main():
    parser = argparse.ArgumentParser(description="Single-core Go v1.4.3 versus v1.4.5 comparison")
    parser.add_argument("--baseline", choices=("v1.4.3",), default="v1.4.3")
    parser.add_argument("--candidate-version", choices=("v1.4.5",), help="published Go release (default: v1.4.5)")
    parser.add_argument("--features", action="store_true", help="measure the six search, time-span and calendar cohorts instead of the three parsing cohorts")
    parser.add_argument("--suite-source", type=Path, required=True, help="checkout containing the shared tools/benchmark.py runner and fixtures")
    parser.add_argument("--worktree", action="store_true", help="benchmark this Go checkout instead of a release; recorded as worktree, never v1.4.5")
    parser.add_argument("--iterations", type=int, default=16, help="feature corpus repetitions per timed pass")
    parser.add_argument("--runs", type=int, default=6)
    parser.add_argument("--passes", type=int, default=8)
    parser.add_argument("--cpu", type=int, default=2)
    parser.add_argument("--cohort", action="append", choices=PARSE_COHORTS + FEATURE_COHORTS)
    parser.add_argument("--output", type=Path)
    arguments = parser.parse_args()
    candidate = arguments.candidate_version or "v1.4.5"
    cohorts = FEATURE_COHORTS if arguments.features else PARSE_COHORTS
    if arguments.worktree and (not arguments.features or arguments.candidate_version):
        parser.error("--worktree requires --features and cannot be labelled with --candidate-version")
    if arguments.iterations < 1 or (arguments.cohort and any(cohort not in cohorts for cohort in arguments.cohort)):
        parser.error("select cohorts belonging to the chosen suite and positive iterations")
    if arguments.runs < 2 or arguments.runs % 2 or arguments.passes < 1:
        parser.error("use an even number of runs >= 2 and positive passes")
    if arguments.cohort and len(set(arguments.cohort)) != len(arguments.cohort):
        parser.error("select each cohort at most once")
    if not hasattr(os, "sched_getaffinity") or arguments.cpu not in os.sched_getaffinity(0):
        parser.error("run under Linux/WSL and select an available CPU")

    root = Path(__file__).resolve().parents[2]
    reference = local_reference(root, arguments.suite_source)
    destination = arguments.output or Path(".benchmarks/shared-go-features.json" if arguments.features else ".benchmarks/shared-go-core.json")
    if not destination.is_absolute():
        destination = root / destination
    command = [
        sys.executable, str(reference / "tools/benchmark.py"),
        "--engine", f"go-{arguments.baseline}",
        "--engine", "go-worktree" if arguments.worktree else f"go-{candidate}",
        "--runs", str(arguments.runs), "--passes", str(arguments.passes),
        "--cpu", str(arguments.cpu), "--output", str(destination),
    ]
    if arguments.features:
        command.extend(["--features", "--iterations", str(arguments.iterations)])
        if arguments.worktree:
            command.extend(["--go-source", str(root)])
    for cohort in arguments.cohort or []:
        command.extend(["--cohort", cohort])
    print(f"Go benchmark: {arguments.baseline} versus {'worktree' if arguments.worktree else candidate}", flush=True)
    subprocess.run(command, cwd=reference, check=True)


if __name__ == "__main__":
    main()