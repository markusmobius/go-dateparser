### v1.4.6 - 2026-09-13

- Use published Go-Dateutil v2.9.0 for relative parsing and month-based time spans, following Python dateparser 1.4.3 and python-dateutil 2.9.0.post0. No local dependency replacement or Python runtime is required.
- Clamp month/year offsets to the target month's last valid day, preserve fractional seconds at Python microsecond precision, and reject fractional years/months instead of approximating them.
- Retain `PreserveEndOfMonth` as a deprecated source-compatible field with no effect. Month-end clamping is now unconditional, as in Python.
- Reject explicitly invalid Gregorian dates while allowing omitted days to clamp and omitted years to select a valid leap year.
- Add Python-grounded arithmetic and calendar regressions. Preserve Go's finer period labels separately from the arithmetic, without promoting floating-point noise to a finer period.
- Preserve the Go 1.26 minimum, locale and conversion data, caller-owned concurrency, and both root license files.

### v1.4.5 - 2026-09-13

- Roll up v1.4.4's single-thread improvements: lazy locale applicability, shared normalization/digit conversion, and cached default locale ordering. No internal worker goroutines or new build flags.
- Correct Jalali and Hijri parsing against Python dateparser 1.4.3: MDY public defaults, calendar-local missing fields, actual month lengths, omitted-day clamping, explicit-day rollover, two-digit years, and reference-date bounds. Callers relying on the previous DMY default should set `Configuration.DateOrder` explicitly.
- Replace the Go Jalali/Hijri conversion dependencies with immutable native tables generated from pinned Python `convertdate` and `hijridate` packages. Remove the now-unused `go-jalaali`, `go-hijri`, and `go-juliandays` dependencies; Python is not required at runtime.
- Correct search language evidence and detector handling. Explicit languages/locales bypass the callback, returned detector languages retain their order, and translated fragments do not invoke it again. Fix Chinese/Japanese timestamp joining and Cantonese splitter bounds.
- Reject relative dates and time spans outside Python's supported datetime range instead of overflowing; preserve valid large whole-second offsets without overflowing `time.Duration`.
- Add a reproducible independent Python fixture: 7,504 calendar cases, 178 exact search cases, and 10 exception inputs checked for safe handling. Record dependency versions and imported source hashes; retain known compatibility exceptions in `UPSTREAM.md`.
- Add a Go-repository benchmark entry point comparing published Go v1.4.3 with v1.4.5 across nine workloads: three parsing cohorts plus search, n-grams, time spans, Jalali, and Hijri. Both versions use identical inputs; historical output differences and panics are recorded. Explicit pre-release runs are labelled worktree.
- Preserve both root license files, generated locale/matcher data, and the Go 1.26 minimum. Python dateparser remains pinned to v1.4.3.

### v1.4.4 - 2026-09-12

- Check locale applicability lazily in the existing priority order, stopping after a successful parse while preserving eager configuration validation and callback behavior.
- Share locale-independent normalization and digit conversion across candidate locales, retaining the detector's original normalized input.
- Cache both unrestricted default locale orderings and return independent slices to callers. Explicit locale/language/region validation is unchanged.
- Add focused regressions for cached ordering and slice ownership, detector inputs, and invalid configuration after previous-locale success.
- Add a reproducible single-core version comparison. Post-release documentation and tooling now use the same 2,951-case suite as RustDateParser: warmed automatic detection measured 3.85x and HtmlDate strict/past 2.78x versus v1.4.3; explicit-locale ranges overlap. See the README and published raw report for exact-output checks and measurement limits. The v1.4.4 tag and module contents are unchanged.
- Preserve dependencies, generated locale/regex data, public APIs, licensing, and Python dateparser v1.4.3 compatibility. No internal worker goroutines or new build flags are introduced.

### v1.4.3 - 2026-09-10

- Require Go 1.26.0 (recommended toolchain 1.27.1) and refresh runtime, test, and code-generator dependencies.
- Refresh README examples, public API documentation, and performance comparisons after the dependency update; clarify that generated matchers are enabled with or without cgo.
- Catch up with Python dateparser v1.4.3, accounting for all 66 commits after v1.2.1 in [UPSTREAM.md](UPSTREAM.md).
- Refresh generated locale data to CLDR 44.1.0, retain older accepted vocabulary and locale codes, and support supplementary regional overrides including en-US.
- Enable checked-in re2go-generated exact locale matchers by default, with no cgo requirement or build flags. Regenerate them alongside locale data and retain regex fallback behavior for malformed UTF-8.
- Add ASCII fast paths for timezone candidate rejection and split-token classification, preserving timezone extraction and backend-specific Unicode and malformed-UTF-8 behavior.
- Add the optional n-gram search strategy and time-span results, including configurable week starts and rolling month lengths.
- Add opt-in parsing of dates surrounded by unknown edge text, preserving strict and required-parts checks.
- Port signed relative offsets, Russian search and compound-number fixes, whitespace handling, explicit-language ISO dates, year-first/Japanese component fixes, and the upstream language additions.
- Correct required-parts validation and ambiguous-year retries. Support ordinal-day custom layouts and preferred centuries for two-digit years.
- Correct BST to UTC+01:00 and HDT to UTC-09:00 without changing other ambiguous timezone abbreviations.
- Isolate shared normalization and search state for concurrent parsing and searching; add race regressions and a Linux multi-timezone CI matrix.
- Preserve Go time layouts, period types, existing language ordering, and the optional end-of-month behavior. See the compatibility ledger for known differences and verification limitations.

### 2021-12-10

- Implement locale data loader and its unit tests.

### 2021-12-07

- Implement timezone parser and its unit tests.

### 2021-12-05

- Implement script to generate Go code. At this point all Python script is ported.
- Restructure directories layout.

### 2021-12-04

No commit today as well, back to home.

### 2021-12-03

No commit today, going to meet family.

### 2021-12-02

- Implement script to generate locale data from supplementary data.

### 2021-12-01

- Implement script to generate simplified locale data from CLDR data.

### 2021-11-30

- Implement script for parsing CLDR gregorian data.
- Implement script for parsing CLDR date fields data.

### 2021-11-29

- Implement script for generating language maps and ordering languages by their popularity.
- Restructure scripts into one single CLI.

### 2021-11-28

- Implement script for fetching raw data from `CLDR` repository.

### 2021-11-28

- Repository created.
