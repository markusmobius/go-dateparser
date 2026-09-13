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
PARSE_COHORTS = ("auto", "explicit", "htmldate")
FEATURE_COHORTS = ("search-auto", "search-split", "search-ngram", "time-span", "jalali", "hijri")


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


def local_reference(root, source):
    source = source.resolve()
    if not (source / "tools/benchmark.py").is_file():
        raise RuntimeError(f"Rust checkout has no shared benchmark runner: {source}")
    fixture = Path("testdata/python-features.json")
    expected = hashlib.sha256((root / fixture).read_bytes()).hexdigest()
    actual = hashlib.sha256((source / fixture).read_bytes()).hexdigest()
    if actual != expected:
        raise RuntimeError("Go and Rust Python feature fixtures must be byte-identical")
    return source


def main():
    parser = argparse.ArgumentParser(description="Single-core Go/Rust comparison on shared fixtures")
    parser.add_argument("--baseline", choices=("v1.4.3",), help="historical Go-only baseline (default: v1.4.3)")
    parser.add_argument("--candidate-version", choices=("v1.4.4", "v1.4.5"),
                        help="published Go release: v1.4.5 for features, v1.4.4 for the historical suite")
    parser.add_argument("--features", action="store_true", help="compare Rust with Go on search, spans, Jalali and Hijri")
    parser.add_argument("--rust-source", type=Path, help="explicit Rust checkout containing the feature benchmark; may be uncommitted")
    parser.add_argument("--worktree", action="store_true", help="benchmark this Go checkout instead of a release; recorded as worktree, never v1.4.5")
    parser.add_argument("--iterations", type=int, default=64, help="feature corpus repetitions per timed pass")
    parser.add_argument("--runs", type=int, default=6)
    parser.add_argument("--passes", type=int, default=8)
    parser.add_argument("--cpu", type=int, default=2)
    parser.add_argument("--cohort", action="append", choices=PARSE_COHORTS + FEATURE_COHORTS)
    parser.add_argument("--output", type=Path)
    arguments = parser.parse_args()
    candidate = arguments.candidate_version or ("v1.4.5" if arguments.features else "v1.4.4")
    cohorts = FEATURE_COHORTS if arguments.features else PARSE_COHORTS
    if arguments.features:
        if not arguments.rust_source or arguments.baseline:
            parser.error("--features requires --rust-source and compares Rust, not a Go --baseline")
        if arguments.worktree and arguments.candidate_version:
            parser.error("--worktree cannot be labelled with --candidate-version")
        if not arguments.worktree and candidate != "v1.4.5":
            parser.error("the Python-qualified feature comparison requires Go v1.4.5")
    elif arguments.rust_source or arguments.worktree or candidate != "v1.4.4":
        parser.error("use --features --rust-source for the Rust/Go v1.4.5 comparison")
    if arguments.iterations < 1 or (arguments.cohort and any(cohort not in cohorts for cohort in arguments.cohort)):
        parser.error("select cohorts belonging to the chosen suite and positive iterations")
    if arguments.runs < 2 or arguments.runs % 2 or arguments.passes < 1:
        parser.error("use an even number of runs >= 2 and positive passes")
    if arguments.cohort and len(set(arguments.cohort)) != len(arguments.cohort):
        parser.error("select each cohort at most once")
    if not hasattr(os, "sched_getaffinity") or arguments.cpu not in os.sched_getaffinity(0):
        parser.error("run under Linux/WSL and select an available CPU")

    root = Path(__file__).resolve().parents[2]
    reference = local_reference(root, arguments.rust_source) if arguments.features else shared_reference(root)
    destination = arguments.output or Path(".benchmarks/shared-features.json" if arguments.features else ".benchmarks/shared-dateparser.json")
    if not destination.is_absolute():
        destination = root / destination
    command = [
        sys.executable, str(reference / "tools/benchmark.py"),
        "--engine", "rust" if arguments.features else f"go-{arguments.baseline or 'v1.4.3'}",
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
    if arguments.features:
        print(f"Feature benchmark: Rust checkout {reference}; Go {'worktree' if arguments.worktree else candidate}", flush=True)
    else:
        print(f"Shared benchmark: rust-dateparser@{REFERENCE_COMMIT}", flush=True)
    subprocess.run(command, cwd=reference, check=True)


if __name__ == "__main__":
    main()