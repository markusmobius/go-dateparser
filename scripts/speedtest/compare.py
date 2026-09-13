import argparse
import hashlib
import os
from pathlib import Path
import subprocess
import sys
import tempfile


REFERENCE_REPOSITORY = "https://github.com/markusmobius/rust-dateparser.git"
REFERENCE_COMMIT = "37b447f048d2dc9cd5f251191d76959b582c4f04"
FIXTURE_SHA256 = "c07b2bb77d75a959b86364a5c2f1fbd6a24b9c6b3619b70099141ab75cb2f24f"


def shared_reference(root):
    cache = root / ".benchmarks"
    cache.mkdir(parents=True, exist_ok=True)
    destination = cache / f"shared-suite-{REFERENCE_COMMIT}"
    if not destination.exists():
        with tempfile.TemporaryDirectory(prefix="shared-download-", dir=cache) as temporary:
            checkout = Path(temporary) / "suite"
            subprocess.run(
                ["git", "clone", "--quiet", "--filter=blob:none", "--no-checkout",
                 REFERENCE_REPOSITORY, str(checkout)], check=True,
            )
            subprocess.run(
                ["git", "-C", str(checkout), "checkout", "--quiet", "--detach", REFERENCE_COMMIT],
                check=True,
            )
            checkout.rename(destination)
    actual = subprocess.check_output(
        ["git", "-C", str(destination), "rev-parse", "HEAD"], text=True,
    ).strip()
    changes = subprocess.check_output(
        ["git", "-C", str(destination), "status", "--porcelain", "--untracked-files=all"], text=True,
    ).strip()
    if actual != REFERENCE_COMMIT or changes:
        raise RuntimeError(f"Shared benchmark checkout is not the clean pinned revision: {destination}")
    fixture = destination / "testdata/go-core.json"
    if hashlib.sha256(fixture.read_bytes()).hexdigest() != FIXTURE_SHA256:
        raise RuntimeError("Shared benchmark fixture checksum differs from the published suite")
    return destination


def main():
    parser = argparse.ArgumentParser(description="Go release comparison on the shared Go/Rust suite")
    parser.add_argument("--baseline", choices=("v1.4.3",), default="v1.4.3")
    parser.add_argument("--candidate-version", choices=("v1.4.4",), default="v1.4.4",
                        help="pinned published release; the current checkout is not benchmarked")
    parser.add_argument("--runs", type=int, default=6)
    parser.add_argument("--passes", type=int, default=8)
    parser.add_argument("--cpu", type=int, default=2)
    parser.add_argument("--cohort", action="append", choices=("auto", "explicit", "htmldate"))
    parser.add_argument("--output", type=Path, default=Path(".benchmarks/shared-dateparser.json"))
    arguments = parser.parse_args()
    if arguments.runs < 2 or arguments.runs % 2 or arguments.passes < 1:
        parser.error("use an even number of runs >= 2 and positive passes")
    if arguments.cohort and len(set(arguments.cohort)) != len(arguments.cohort):
        parser.error("select each cohort at most once")
    if not hasattr(os, "sched_getaffinity") or arguments.cpu not in os.sched_getaffinity(0):
        parser.error("run under Linux/WSL and select an available CPU")

    root = Path(__file__).resolve().parents[2]
    reference = shared_reference(root)
    destination = arguments.output
    if not destination.is_absolute():
        destination = root / destination
    command = [
        sys.executable, str(reference / "tools/benchmark.py"),
        "--engine", f"go-{arguments.baseline}",
        "--engine", f"go-{arguments.candidate_version}",
        "--runs", str(arguments.runs), "--passes", str(arguments.passes),
        "--cpu", str(arguments.cpu), "--output", str(destination),
    ]
    for cohort in arguments.cohort or []:
        command.extend(["--cohort", cohort])
    print(f"Shared benchmark: rust-dateparser@{REFERENCE_COMMIT}", flush=True)
    subprocess.run(command, cwd=reference, check=True)


if __name__ == "__main__":
    main()