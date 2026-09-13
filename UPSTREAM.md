# Upstream Compatibility

Source: https://github.com/scrapinghub/dateparser

- Previous documented baseline: `02bd2e5dd4477b4f6db98c5e98149458eb3cc821` (v1.2.1).
- Target: `9ce60b1958f1b285886bcfbb743f6419feacfc92` (v1.4.3).
- Range: all 66 commits accounted for, in topological order.
- Baseline verification: `go test -mod=mod ./...` passed on Windows with Go 1.27.1.

`Ported` means the applicable change and regression coverage are present.
`Equivalent` means existing Go behavior covers the upstream change.
`N/A` records a Python-specific or project-maintenance change with no Go runtime change.

Intentional Go differences remain: Go time layouts instead of strptime directives,
time values always have a location, finer hour/minute/second periods, relative
weeks represented by the Day period, and the
optional PreserveEndOfMonth behavior. Tests use a fixed CurrentTime where the
upstream test depends on today's date. Go uses RE2-compatible expressions and
string operations, not Python regex syntax or Python serialization caches.
The existing UseGivenOrder setting supplies upstream's given-language-order
behavior, and Go retains its supported language set and detection ranking.

## Verification

### v1.4.5 Python Reference

Python dateparser **1.4.3** remains the behavioral authority. The independent
generator in [scripts/python-reference/export.py](scripts/python-reference/export.py)
uses pinned packages from
[scripts/python-reference/requirements.txt](scripts/python-reference/requirements.txt).
It checks calendar dependency versions and records SHA-256 hashes of the
imported Python sources. Those installed wheel sources have not all been
independently compared byte-for-byte with the target Git commit.

[testdata/python-features.json](testdata/python-features.json) contains
7,504 calendar cases and 188 search cases. Of the search cases, 178 compare
exact text, detected language, datetimes, offsets and callback inputs; ten
Python exception inputs check safe handling instead of reproducing Python's
splitter `IndexError` or arithmetic `OverflowError`. Calendar errors compare
success/failure, not exception wording. Go's finer time periods are normalized
when compared with Python. Tests explicitly select `PreserveEndOfMonth` where
needed without changing its optional Go default.

The corrections cover MDY calendar defaults, reference-date conversion using
local date components, missing fields, actual month lengths, explicit-day
rollover and supported ranges; search word evidence and detector behavior;
Chinese/Japanese joining and Cantonese bounds; and relative/time-span overflow.
The existing Hindi/Lao exceptions below remain. This fixture is substantial
regression coverage, not a claim of universal Python parity.

The corrected runtime passed the complete Go 1.27.1 cgo-disabled suite, the
Go 1.26.0 cgo-disabled suite, the Go 1.27.1 race suite in UTC,
America/New_York and Asia/Kolkata, the native Windows Go 1.27.1 cgo-disabled
suite, `go vet`, and `go mod verify`. The actual `re2_wasm` backend also passed
the full suite on Go 1.26.0 with `CGO_ENABLED=1`. Native `re2_cgo` remains
unqualified because its separate RE2 library is not installed. Python fixture
and conversion-data regeneration passed `--check` byte-for-byte. These are
local worktree checks, not hosted CI results for a published tag.

Reproduce the reference data in an isolated Python environment with the pinned
requirements installed:

```sh
python scripts/python-reference/export.py --features-output testdata/python-features.json --check
go test -mod=readonly ./... -run TestPython -count=1
```

Fixture SHA-256:
`b138df4fe7beb7dd40c51f01267f78856031a2de4ddbf5e958c95b33a9d5e418`.
Calendar conversion-data SHA-256:
`ad31d61fb903d5f2f5d22a38df22e6cba6c3dcf7322d6f325a27df417ebfc2e2`.

### Historical v1.4.4 Verification

The Go-only v1.4.4 optimization keeps the Python v1.4.3 baseline and the
compatibility ledger below unchanged. On 2026-09-12, the final runtime passed
the complete Go 1.27.1 race suite in UTC, America/New_York and Asia/Kolkata,
the complete Go 1.26.0 suite with cgo disabled, `go vet`, and `go mod verify`.
The complete native Windows/amd64 suite passed with Go 1.27.1 and cgo disabled;
the optional `re2_wasm` suite passed under Linux with Go 1.26.0 and cgo enabled.
The native `re2_cgo` backend was not requalified because its separate RE2
library is not installed.
Generated locale/matcher files, dependencies and both license files are unchanged.
The shared benchmark additionally compares exact results for 2,951 stateless
parsing cases in both published Go releases. Automatic, explicit and HtmlDate
cohorts contain 226, 2,530 and 195 cases, with 222, 2,388 and 167 accepted.
The inputs, settings and v1.4.4 samples are also used by the Rust comparison.
This is additional regression coverage, not a new independent Python comparison.

Dependencies were refreshed on 2026-09-10, with a new minimum Go version of 1.26.0 and recommended toolchain 1.27.1. The current performance measurements below use the refreshed dependencies.

- Full `go test ./... -count=1` passed on Windows with the declared minimum Go 1.26.0 and cgo disabled, including the default generated exact matchers.
- Full `go test -race ./... -count=1` passed with Go 1.27.1 and GCC 16.2.0. Expanded concurrent shared-parser, independent-parser, relative-base, and detector/previous-locale regressions also passed under the race detector.
- Full `go test -tags re2_wasm ./... -count=1` passed on Go 1.26.0 with cgo enabled. The native `re2_cgo` backend was not tested because its separate RE2 library is not installed.
- `go mod verify` passed. `govulncheck` v1.8.0 found no reachable vulnerabilities or vulnerable imported packages in the default build on Go 1.27.1. It reported module-only advisory [GO-2026-5932](https://pkg.go.dev/vuln/GO-2026-5932) for the unimported, unmaintained `golang.org/x/crypto/openpgp` package; no fixed version is listed.
- Before the dependency refresh, regeneration with `go run ./scripts/codegen --skip-raw --keep-language-order` and re2go 4.4 reproduced all 211 generated locale/matcher source files byte-for-byte. Use `--re2go` to select an executable outside PATH. A missing generator returned a nonzero exit status and left existing output unchanged.
- Before the dependency refresh, a temporary differential driver compared 575 fixed-reference cases with Python dateparser 1.4.3: upstream regression scenarios plus month and relative phrases sampled across the refreshed locale files. It compared parse success, wall-clock values, normalized periods, locales, and search result text/order. 573 matched; the two inherited exceptions below were confirmed against the original Go HEAD. This sample is not exhaustive parity proof.
- The Linux CI matrix uses Go 1.26.x/stable in UTC, America/New_York, and Asia/Kolkata, with `GOTOOLCHAIN=local` to prevent automatic toolchain switching. Both the [branch run](https://github.com/markusmobius/go-dateparser/actions/runs/34727618390) and [tag run](https://github.com/markusmobius/go-dateparser/actions/runs/34727618210) passed for the published v1.4.4 source commit.
- With current dependencies and no build tags, `TestExactCombinedMatchers` and its fallback regression passed separately with `CGO_ENABLED=0` and `CGO_ENABLED=1`, performing 241,049 comparisons in each build. Both selected the Go standard backend for remaining regex operations.
- The README's Go examples were compiled and executed in a temporary Markdown/Go AST harness, including the complete quick-start and lingua-go v1.4.0 examples. Timestamp output is explicitly converted to UTC; the parser itself preserves `time.Local` for Unix timestamps.

## Generated Matchers

Exact whole-token locale matching uses 127 deduplicated re2go-generated functions by default, with cgo enabled or disabled and regardless of optional regex build tags. Locale data stores direct function references and inherits them alongside the original regexes. Locales without a combined expression need no matcher. Malformed UTF-8 falls back to the existing regex engine. There is no benchmark switch, additional runtime dependency, or required build tag. Capture-based regex operations retain their existing implementation and optional backends.

The permanent `TestExactCombinedMatchers` checks all locale bindings and performs 241,049 comparisons against the original regexes, covering generated matching strings, locale vocabulary, Unicode folds, truncated inputs, NULs, and malformed UTF-8. The initial isolated prototype also passed 1,188,746 comparisons. The enabled implementation passes the full minimum-Go, race-enabled, and WebAssembly RE2 suites listed above.

## Additional Fast Paths

Two further optimizations are enabled alongside the generated matchers. A conservative ASCII candidate filter skips timezone scans when no signed offset or possible timezone word is present; non-ASCII input and unusually long tokens retain the existing path. Candidate extraction still uses the original regexes and offset ordering. Split-token classification checks ASCII letters and digits directly while preserving newline rejection, formatting rules, and the original regex behavior for non-ASCII and malformed UTF-8.

Permanent regressions compare timezone candidate rejection against every configured timezone name across boundary combinations and compare token classification against the original regexes on ASCII, Unicode, newline, and malformed-UTF-8 cases. The temporary comparison switches and workload instrumentation remain outside the repository.

## Performance

The current **Go v1.4.3 versus v1.4.5** comparison is displayed directly in
[README.md](README.md#current-measurements) and the v1.4.5 release notes.
It covers the three parsing cohorts and six search, time-span and calendar
cohorts with the same inputs and settings for both published modules.

The runner uses one caller on one CPU, six balanced launches per version and
eight warm passes per launch. Feature passes repeat their corpus 16 times;
the table reports time per complete corpus traversal. Initial output checks
are outside the timers. Current feature results must match the independent
Python reference; historical result differences and recovered panics are
recorded rather than silently dropping inputs or claiming output equivalence.

The [reproduction commands](README.md#reproduce-measurements) select only the
two published Go versions. An explicitly selected worktree is labelled as
such and is never presented as a release. No published tag or module is moved
by these documentation and benchmark-tool updates.

### Historical v1.4.3 Measurements

These measurements replace the earlier pre-refresh comparisons. On 2026-09-10, a four-way comparison used the current dependency graph, Go 1.27.1, Windows/amd64 on an AMD Ryzen AI 7 PRO 350, `CGO_ENABLED=1`, and no optional regex tags. All four configurations used Go's standard engine for remaining regex operations:

1. **Retained regexes:** generated matcher bindings disabled, with the original timezone and token-classification paths.
2. **re2go only:** generated matcher bindings enabled, with the original timezone and token-classification paths.
3. **Plus timezone filter:** re2go and the conservative timezone filter enabled.
4. **Current defaults:** re2go, the timezone filter, and ASCII token classification enabled.

The corpus contains 762 parse inputs and 62 search texts from the existing benchmarks. All outputs, including locales and search text/order, matched across all four configurations at a fixed reference time of 2026-09-10 12:00 UTC. Timing used that same reference time, independent parser instances, randomized configuration order for each input, one discarded warmup, and six measured rounds. Setup and comparison switches were outside the timed regions.

| Workload | Retained Regexes | re2go Only | Plus Timezone Filter | Current Defaults |
| --- | --- | --- | --- | --- |
| Parse, 762 inputs | 7.684 s | 6.677 s | 5.857 s | 5.278 s |
| Search, 62 texts | 3.157 s | 3.251 s | 2.071 s | 1.995 s |

These are median times per complete corpus. Relative to retained regexes, current defaults used **31.3% less parse time** (1.46x throughput) and **36.8% less search time** (1.58x). Relative to re2go alone, the additional fast paths reduced time by **20.9%** and **38.7%**, respectively. These are ratios from the same run, not percentages added from separate experiments.

Machine load varied substantially across rounds. The isolated re2go effect was 13.1% less parse time, but search was noisy and 3.0% slower by these medians. This run does not establish an isolated re2go search benefit. Current defaults were faster than retained regexes in every measured round. These results are not per-date latency, a comparison with the original Go release or Python, or a guarantee for other workloads. The optional RE2 engines were not benchmarked here.

The generated source remains **1,890,891 bytes**. A separate size comparison on the current dependency graph and the same Go/cgo settings built the root test executable with and without generated matcher bindings; both retained the ASCII fast paths. Sizes were **14,748,160 bytes** with retained regexes and **16,805,888 bytes** with generated matching, an increase of **2,057,728 bytes (14.0%)**. Temporary build overlays and measurement tools remain outside the repository; no production switches were added.

## Retained Exceptions

At a reference date of 2025-09-15 12:00 UTC, with the language explicitly supplied:

| Input | Go (before and after this update) | Python 1.4.3 |
| --- | --- | --- |
| Hindi U+0915 U+0932 (`kal`) | 2025-09-14 12:00, Day | 2025-09-16 12:00, day |
| Lao U+0EA1 U+0EB7 U+0EC9 U+0E99 U+0EB5 U+0EC9 (today) | 2025-09-15 12:00, Day | 2025-03-15 00:00, month |

The Hindi token appears under both yesterday and tomorrow in the upstream data,
including at the original baseline. These pre-existing Go results are retained;
the CLDR refresh does not attempt to reproduce every ambiguous dictionary or
normalization collision in Python. Custom-layout parsing can also leave Go's
Locale empty when no translation was needed, unlike Python's locale reporting.

## Calendar Data Provenance

[internal/parser/calendars/data.json](internal/parser/calendars/data.json)
contains integer epoch-day boundaries generated with `convertdate` 2.4.1
and `hijridate` 2.6.0. Jalali year starts are calculated through `convertdate`
using PyMeeus 0.5.12; Umm al-Qura month starts are converted from `hijridate`'s
table. The generator verifies round trips and month/year lengths. Go embeds
only the resulting numeric data and implements lookup and boundary arithmetic
natively; none of these Python packages is linked or bundled at runtime.

PyMeeus's source declares LGPL-3.0-or-later and is an external generation-time
dependency. Its source and license are available from
[PyMeeus](https://github.com/architest/pymeeus). Reproducing the data requires
installing the pinned packages with their own accompanying licenses.

The MIT notices for the calendar-data sources are retained here. The two
existing project license files remain unchanged.

### MIT Calendar Source Notices

`convertdate` 2.4.1: Copyright (c) 2014-2022 Neil Freeman.

`hijridate` 2.6.0: Copyright (c) Mohammed Alshehri.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

## Commit Ledger

| Commit | Change | Status / Evidence |
| --- | --- | --- |
| f857dfc | Russian search preposition | Ported: Search preprocessing; TestParser_SearchRussianPrepositions. |
| e9324ee | Cache timezone offsets at Python import | N/A: Go has compiled timezone tables; no pickle cache or Python regex dependency. |
| 114f5c5 | Python 3.14 error messages | N/A: Go does not match Python datetime exception messages. |
| 180e697 | Weekday search across month boundaries | Equivalent: existing weekday correction; exact upstream case added to TestParser_Parse_OnlyWeekdays. |
| 84511a3 | v1.2.2 release notes and tox | N/A: Python release documentation and test environment. |
| f69e9b2 | v1.2.2 version bump | N/A: Python package version. |
| abf7130 | Russian compound word numbers | Ported: supplementary patterns, Simplify, and all 13 upstream search examples. |
| 0667cb3 | British English relative phrases | Ported: supplementary week patterns and UpstreamTranslations tests. |
| bdc57d3 | Time span search and settings | Ported: time-span.go, Configuration, SearchTimeSpan tests. Reference time of day is preserved as upstream implements it. |
| b7bb1a5 | Current year for custom formats | Equivalent: formatted.Parse uses CurrentTime; formatMissingYear regression. |
| cbe2979 | Python support and CI | N/A: Python versions, packaging, and ruamel.yaml API migration; Go uses its own YAML loader. |
| 6b23348 | Stabilize Russian search test year | Equivalent: ported search test uses a fixed CurrentTime instead of the wall clock. |
| 920d725 | Short English weekdays | Ported: supplementary weekday names and ordinal-suffix filtering; upstream's known-failing Mo case remains excluded. |
| f0f9102 | Finnish klo regression cases | Ported: Finnish skip token and all five upstream parsing cases. |
| d2e236f | Signed relative offsets | Ported: per-component signs before unit aggregation; relative_signedOffsets tests. |
| 0e87891 | v1.3.0 release notes | N/A: Python release documentation; Go changes tracked here. |
| c2ff31c | v1.3.0 version bump | N/A: Python package version. |
| 24dac43 | Python MD5 in FIPS environments | N/A: no MD5-based settings cache in Go. |
| 35e7f61 | Test in multiple timezones | Ported: Go CI configured to run the race suite in UTC, America/New_York, and Asia/Kolkata on Linux. |
| 0f04199 | Bosnian Cyrillic translations | Ported: supplementary bs-Cyrl data and UpstreamTranslations regression. |
| 169c4ce | Browser demo | N/A: updates the Python project's demo link. |
| 6ca550f | English from now expressions | Ported: English supplementary patterns and UpstreamTranslations tests. |
| cd5f226 | Required parts and ambiguous years | Ported: correct RequiredParts validation and year-order retries; requiredParts tests. |
| e0939ec | Python installation documentation | N/A: Python installation and publishing commands. |
| 8cdcf80 | Logo | N/A: Python project artwork. |
| e22e62e | Word-number relative phrases | Ported: English one-through-twelve later expressions and translation tests. |
| 5add61c | Python publishing workflow | N/A: PyPI publishing is independent of Go module distribution. |
| 341d308 | Python deployment | N/A: PyPI workflow permissions. |
| 62710c0 | Security policy | N/A: Python support policy and package classifiers; not a runtime change. |
| 7a3113e | en-US locale | Ported: supplementary locale metadata, generated alias, and USLocale tests for explicit locale and region. |
| d405657 | v1.4.0 release and security fixes | Equivalent: typed boolean locale data and compiled-in timezone tables; Go never evaluates locale strings or loads pickle files. |
| 60b8a75 | v1.4.0 version bump | N/A: Python package version and support policy. |
| 3b4a816 | Read the Docs environment | N/A: Python documentation hosting environment. |
| b72ed09 | Remove fastText and unblock NumPy 2 | N/A: Go has no bundled fastText/NumPy integration; external language detection remains a callback. |
| 373ede9 | Preserve translation whitespace | Ported: alphabetic skip-word removal and maximum adjacent spacing; TestTranslateWhitespace and existing fixtures. |
| 33ec7ef | Czech translations | Ported: supplementary Czech vocabulary and clock expressions with translation regressions. |
| 075712f | Spaces around relative operators | Ported with d2e236f; spaced positive and negative offset regressions. |
| 85af00e | Region docs and reference month | Equivalent: absolute parser already passes p.Now.Month(); fixed May 31 regression and region-selection test. |
| 24bfa9e | Search language fallback | Ported: detected language followed by explicit candidates; SearchLanguageFallback invoice regression. |
| b903e69 | English mon/mons units | Ported: supplementary unit patterns, positive relative test, and negative word-boundary tests. |
| b107ac3 | Python ordered dictionaries | N/A: Python dict now preserves insertion order; existing Go ordered slices and explicit sorting remain unchanged. |
| 5c437a1 | Python packaging metadata | N/A: pyproject migration and Python documentation reorganization. |
| 333e519 | Packaging merge repair | N/A: Python data-generator indentation repair. |
| 081d251 | Norwegian Bokmal translations | Ported: supplementary numbers, units, and skip words with regressions. |
| 0d11d82 | Given language order setting | Equivalent: existing UseGivenOrder already applies to top-level Parse; all eight upstream ordering cases tested. |
| 5c6d97f | Signs in mixed decades and years | Ported with d2e236f; mixed-sign decades/years tested against Python 1.4.3. |
| 98b9c32 | Relative-regex backtracking | Equivalent: Go's linear-time regex engine; 3,200/6,400-digit regressions cover translation and relative parsing. |
| a049fd1 | Korean alternative expressions | Ported: supplementary relative expressions and translation tests. |
| 08c78d3 | v1.4.1 release | N/A: Python release notes and version metadata. |
| eb111c2 | Czech inflections and relatives | Ported: additional Czech inflections and relative vocabulary with regressions. |
| 33e913c | Italian relative expressions | Ported: un'ora/un ora and alle; positive and word-boundary regressions. |
| 273eb23 | Day-of-year custom formats | Ported: 002/__2 layouts preserve parsed month/day; TestParseDayOfYear. |
| 762dfff | Preferred dates for two-digit years | Ported: custom-layout century preferences with reference-date and leap-year regressions. |
| 93b3c63 | CLDR 44.1 and duration/hour data | Ported: pinned date/core/unit generation, supplementary vocabulary deltas and regional overrides; TestSupplementaryLocaleOverrides, retainedCLDRLocales, Catalan hour translation, and reproducible generation. |
| cf6d3e5 | N-gram search strategy | Ported: SearchStrategy and ngram-search.go; TestParser_SearchNgram covers upstream examples, exact substrings, blacklist, relative dates, and spans. |
| 315e396 | ISO dates with explicit languages | Ported: inferred four-digit year keeps month/day order unless caller overrides it; TestParser_Parse_yearFirstComponents and Python comparisons. |
| 06ca032 | Year-first date component handling | Ported: defer two-digit-year interpretation when later components are present; Japanese/month-day regressions and shared calendar parser suites. |
| d886e23 | Ignore surrounding text setting | Ported: opt-in strict-first retry, edge token filtering and exact timezone recognition; ignoreSurroundingText and surroundingTimezones regressions include detector, previous-locale, format, and strictness paths. |
| ff3dcb9 | README RST metadata | N/A: Python README markup and documentation include directives. |
| b1f603e | v1.4.2 release | N/A: Python release notes and version metadata. |
| 764091d | BST and HDT timezone offsets | Ported: corrected runtime table and TestTimezoneUpstreamOffsets/Parse_correctedTimezoneOffsets, including IANA comparisons and unchanged ambiguous offsets. Python's development-only abbreviation-conflict reporter is not shipped; Go uses focused regression checks. |
| de7b986 | CodSpeed benchmarks | N/A: Python-specific benchmark runner; existing Go parse/search benchmarks cover the corresponding operations. |
| cd9bb13 | Shared parser state and caches | Ported: retain locked Parse/per-locale config clones and compiled locale data; use exclusively owned normalization transformers. TestParser_ConcurrentParseAndSearch and the full race suite pass. Python-only cache eviction/loader locks are N/A to compiled Go tables. |
| 6b5f17d | Reject invalid day-of-year formats | Equivalent: time.ParseInLocation rejects invalid ordinal days; all upstream boundary cases tested in TestParseDayOfYear. |
| 27182cd | Isolate search detection state | Ported: detection and locale choices are per call; read charset-cache snapshot under lock. Concurrent mixed-language Search and Parse leave caller configs/reference dates unchanged. |
| 9ce60b1 | v1.4.3 release | N/A: Python release notes and version metadata; target pinned above. |