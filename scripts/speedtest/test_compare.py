import contextlib
import importlib.util
import io
import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest
from unittest import mock

import compare


class CompareTests(unittest.TestCase):
    def setUp(self):
        temporary = tempfile.TemporaryDirectory()
        self.addCleanup(temporary.cleanup)
        self.go_root = Path(temporary.name) / "go"
        self.rust_root = Path(temporary.name) / "rust"
        for root in (self.go_root, self.rust_root):
            (root / "testdata").mkdir(parents=True)
            (root / "testdata/python-features.json").write_text("{}\n", encoding="utf-8")
        (self.rust_root / "tools").mkdir()
        (self.rust_root / "tools/benchmark.py").touch()

    def invoke(self, *arguments):
        with (
            mock.patch.object(sys, "argv", ["compare.py", "--suite-source", str(self.rust_root), *arguments]),
            mock.patch.object(compare.os, "sched_getaffinity", return_value={2}, create=True),
            mock.patch.object(compare, "local_reference", return_value=self.rust_root),
            mock.patch.object(compare.subprocess, "run") as run,
            contextlib.redirect_stdout(io.StringIO()),
            contextlib.redirect_stderr(io.StringIO()),
        ):
            compare.main()
            run.assert_called_once()
            return run.call_args.args[0]

    def test_matching_local_fixtures(self):
        self.assertEqual(compare.local_reference(self.go_root, self.rust_root), self.rust_root.resolve())

    def test_mismatched_local_fixtures_are_rejected(self):
        (self.rust_root / "testdata/python-features.json").write_text('{"changed":true}\n', encoding="utf-8")
        with self.assertRaisesRegex(RuntimeError, "byte-identical"):
            compare.local_reference(self.go_root, self.rust_root)

    def test_missing_local_runner_is_rejected(self):
        (self.rust_root / "tools/benchmark.py").unlink()
        with self.assertRaisesRegex(RuntimeError, "no shared benchmark runner"):
            compare.local_reference(self.go_root, self.rust_root)

    def test_core_comparison_selects_two_published_go_versions(self):
        command = self.invoke()
        self.assertIn("go-v1.4.3", command)
        self.assertIn("go-v1.4.5", command)
        self.assertNotIn("rust", command)
        self.assertNotIn("go-v1.4.4", command)
        self.assertNotIn("--features", command)
        self.assertNotIn("--iterations", command)

    def test_feature_comparison_selects_only_published_go_versions(self):
        command = self.invoke("--features", "--cohort", "search-ngram")
        self.assertIn("go-v1.4.3", command)
        self.assertIn("go-v1.4.5", command)
        self.assertNotIn("rust", command)
        self.assertNotIn("go-v1.4.4", command)
        self.assertIn("search-ngram", command)
        self.assertIn("--features", command)
        self.assertEqual(command[command.index("--iterations") + 1], "16")
        self.assertNotIn("--go-source", command)

    def test_worktree_mode_is_not_labelled_as_release(self):
        command = self.invoke("--features", "--worktree")
        self.assertIn("go-worktree", command)
        self.assertIn("--go-source", command)
        self.assertNotIn("go-v1.4.5", command)

    def test_invalid_combinations_fail_before_launching(self):
        source = ["--features"]
        for arguments in (
            ["--worktree"],
            [*source, "--candidate-version", "v1.4.4"],
            [*source, "--worktree", "--candidate-version", "v1.4.5"],
            [*source, "--baseline", "v1.4.4"],
            [*source, "--cohort", "auto"],
            [*source, "--cohort", "jalali", "--cohort", "jalali"],
            [*source, "--iterations", "0"],
            [*source, "--runs", "3"],
        ):
            with self.subTest(arguments=arguments), self.assertRaises(SystemExit) as raised:
                self.invoke(*arguments)
            self.assertEqual(raised.exception.code, 2)


class SharedRunnerTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        source = os.environ.get("RUST_DATEPARSER_SOURCE")
        if not source:
            raise unittest.SkipTest("set RUST_DATEPARSER_SOURCE to test the shared Rust benchmark driver")
        path = Path(source) / "tools/benchmark.py"
        spec = importlib.util.spec_from_file_location("shared_dateparser_benchmark", path)
        cls.runner = importlib.util.module_from_spec(spec)
        spec.loader.exec_module(cls.runner)

    def module(self, **changes):
        module = {
            "Path": self.runner.GO_MODULE, "Version": "v1.4.5",
            "Sum": "h1:" + "A" * 43 + "=", "Origin": {"Hash": "a" * 40},
        }
        module.update(changes)
        return json.dumps(module)

    def tag(self, annotated=True):
        commit = "b" * 40 if annotated else "a" * 40
        output = f"{commit}\trefs/tags/v1.4.5\n"
        if annotated:
            output += f"{'a' * 40}\trefs/tags/v1.4.5^{{}}\n"
        return subprocess.CompletedProcess(["git"], 0, output, "")

    def test_unpublished_tag_does_not_query_module_proxy(self):
        with (
            mock.patch.object(self.runner.subprocess, "run", return_value=subprocess.CompletedProcess(["git"], 2, "", "")),
            mock.patch.object(self.runner.subprocess, "check_output") as download,
            self.assertRaisesRegex(RuntimeError, "not published"),
        ):
            self.runner.published_reference("v1.4.5", Path("temporary.mod"), Path.cwd(), {})
        download.assert_not_called()

    def test_annotated_and_lightweight_tags_match_the_module(self):
        for annotated in (False, True):
            with (
                self.subTest(annotated=annotated),
                mock.patch.object(self.runner.subprocess, "run", return_value=self.tag(annotated)),
                mock.patch.object(self.runner.subprocess, "check_output", return_value=self.module()),
            ):
                reference = self.runner.published_reference("v1.4.5", Path("temporary.mod"), Path.cwd(), {})
            self.assertEqual(reference["version"], "v1.4.5")
            self.assertEqual(reference["commit"], "a" * 40)
            self.assertEqual(reference["module_sum"], "h1:" + "A" * 43 + "=")

    def test_wrong_module_identity_is_rejected(self):
        for changes in ({"Origin": {"Hash": "c" * 40}}, {"Sum": ""}, {"Version": "v1.4.4"}, {"Path": "other/module"}):
            with (
                self.subTest(changes=changes),
                mock.patch.object(self.runner.subprocess, "run", return_value=self.tag()),
                mock.patch.object(self.runner.subprocess, "check_output", return_value=self.module(**changes)),
                self.assertRaises(RuntimeError),
            ):
                self.runner.published_reference("v1.4.5", Path("temporary.mod"), Path.cwd(), {})

    def test_old_release_does_not_need_new_tag_resolution(self):
        with (
            mock.patch.object(self.runner.subprocess, "run") as resolve,
            mock.patch.object(self.runner.subprocess, "check_output", return_value=self.module(Version="v1.4.4")),
        ):
            reference = self.runner.published_reference("v1.4.4", Path("temporary.mod"), Path.cwd(), {})
        resolve.assert_not_called()
        self.assertEqual(reference["version"], "v1.4.4")

    def test_summary_compares_rust_with_145(self):
        records = [
            {
                "cohort": "search-auto", "engine": engine, "pass_ms": [duration],
                "launch_to_ready_ms": 1,
                "metadata": {"cases": 3, "parsed": 3, "iterations": 1},
            }
            for engine, duration in (("rust", 1), ("go-v1.4.3", 6), ("go-v1.4.5", 3))
        ]
        summary = self.runner.summarize(records)
        self.assertEqual(summary["search-auto"]["go_over_rust_time"], 3)
        self.assertEqual(summary["search-auto"]["go_old_over_new_time"], 2)

    def test_historical_outcome_variability_is_reported(self):
        records = [
            {
                "cohort": "jalali", "engine": "go-v1.4.3", "pass_ms": [1],
                "launch_to_ready_ms": 1,
                "metadata": {
                    "cases": 1311, "parsed": parsed, "matched_dates": parsed,
                    "iterations": 1, "python_matches": matched,
                    "python_mismatches": 1311 - matched, "panics": 0,
                    "outcome_sha256": str(parsed),
                },
            }
            for parsed, matched in ((952, 797), (940, 809))
        ]
        summary = self.runner.summarize(records)["jalali"]["go-v1.4.3"]
        self.assertEqual(summary["parsed_range"], [940, 952])
        self.assertEqual(summary["python_mismatches_range"], [502, 514])
        self.assertEqual(summary["outcome_variants"], 2)


if __name__ == "__main__":
    unittest.main()