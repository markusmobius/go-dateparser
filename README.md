# Go-DateParser [![Go Reference][go-ref-icon]][go-ref]

This package parses localized dates in string formats commonly found on web pages. It is a Go port of the [Python library][original] with the same name.

Requires **Go 1.26.0 or newer**. The module recommends Go 1.27.1 for development; normal builds do not require cgo or re2go.

To use it, install the package inside your project:

```sh
go get github.com/markusmobius/go-dateparser@v1.4.3
```

## Table of Contents

- [1. Features](#features)
- [2. Status](#status)
- [3. Common Use Cases](#common-use-cases)
- [4. Usage](#usage)
	- [Search Strategies](#search-strategies)
	- [Time Spans](#time-spans)
	- [Concurrent Use](#concurrent-use)
- [5. Supported Languages and Locales](#supported-languages-and-locales)
- [6. Supported Patterns](#supported-patterns)
  - [6.1. Timestamp](#61-timestamp)
  - [6.2. Relative Date](#62-relative-date)
  - [6.3. Absolute Date](#63-absolute-date)
	- [6.4. Custom Formats](#64-custom-formats)
- [7. Supported Calendars](#supported-calendars)
- [8. Language-Based Date Order](#language-based-date-order)
- [9. Timezones and UTC Offsets](#timezone-and-utc-offset)
- [10. Incomplete Date Handling](#incomplete-date-handling)
- [11. Custom Language Detector](#custom-language-detector)
- [12. Handling False Positives](#handling-false-positives)
	- [Ignoring Surrounding Text](#ignoring-surrounding-text)
- [13. Performance](#performance)
	- [Current Measurements](#current-measurements)
	- [Build Modes](#build-modes)
  - [Compiling with cgo under Linux](#compiling-with-cgo-under-linux)
  - [Compiling with cgo under Windows](#compiling-with-cgo-under-windows)
- [14. Contribution Guide](#contribution-guide)
- [15. License](#license)

## <a name="features"></a> 1. Features [▲](#table-of-contents)

- Parsing of dates in more than 200 languages and their regional variants, with automatic language detection.
- Parsing of relative dates such as `"1 min ago"`, `"2 weeks ago"`, `"3 months, 1 week and 1 day ago"`, `"in 2 days"`, and `"tomorrow"`.
- Parsing of dates with timezone abbreviations or UTC offsets, such as `"August 14, 2015 EST"`, `"July 4, 2013 PST"`, and `"21 July 2013 10:15 pm +0500"`.
- Date lookup in longer texts.
- Support for non-Gregorian calendar systems. See [Supported Calendars](#supported-calendars).
- Extensive test coverage.

## <a name="status"></a> 2. Status [▲](#table-of-contents)

This README describes Go-DateParser v1.4.3. Upgrade to this version to receive the changes listed in [CHANGELOG.md](CHANGELOG.md), including the default generated matchers and additional fast paths.

This package tracks the applicable changes through Python dateparser [v1.4.3][original-tag], commit [9ce60b1][original-commit]. The previous baseline, `02bd2e5`, was the v1.2.1 release commit. [UPSTREAM.md](UPSTREAM.md) accounts for all 66 intervening commits, including Python-specific changes that do not apply to Go, verification results, and known compatibility exceptions.

The existing Go APIs and intentional differences are preserved:

- Custom formats use Go time layouts, not Python strptime directives.
- Python can return naive datetimes; Go `time.Time` always has a location. `CurrentTime` defaults to UTC.
- Go provides finer hour/minute/second periods and represents relative weeks with the `Day` period.
- `PreserveEndOfMonth` remains an optional Go setting for relative month/year arithmetic.
- `UseGivenOrder` provides the equivalent of Python's `USE_GIVEN_LANGUAGE_ORDER` setting.
- Translation uses RE2-compatible expressions and Go string operations, with generated locale data from CLDR 44.1.0.

## <a name="common-use-cases"></a> 3. Common Use Cases [▲](#table-of-contents)

**dateparser** supports date parsing and extraction in several workflows:

Consuming data from different sources:

- **Scraping**: extract dates from different places with several different formats and languages
- **IoT**: consuming data coming from different sources with different date formats
- **Tooling**: consuming dates from logs and other sources
- **Format transformations**: converting dates from documents or CSV files into database values or other formats

Offering natural interaction with users:

- **Tooling and CLI**: allow users to write `3 days ago` to retrieve information.
- **Search engines**: allow people to search using natural-language dates.
- **Bots**: allow users to interact with a bot easily

## <a name="usage"></a> 4. Usage [▲](#table-of-contents)

Use [`dateparser.Parse`][dps-parse] for a single date:

```go
package main

import (
	"fmt"

	dps "github.com/markusmobius/go-dateparser"
)

func main() {
	dt, err := dps.Parse(nil, "6 July 2020")
	if err != nil {
		panic(err)
	}
	fmt.Printf("locale: %s, time: %s, period: %s\n",
		dt.Locale, dt.Time.Format("2006-01-02 15:04:05 MST"), dt.Period)
	// Output: locale: en, time: 2020-07-06 00:00:00 UTC, period: Day
}
```

The remaining snippets use the same `dps` import alias; some also use `time` and the package's `date` subpackage. Errors are omitted in shorter examples but should be handled in application code.

`Date.Period` describes precision, not a probability of correctness. By default, times are reported with the `Day` period; set `ReturnTimeAsPeriod` to receive `Hour`, `Minute`, or `Second` where applicable.

You can also extract dates from longer text using [`dateparser.Search`][dps-search]. It returns a language code and a list of [`SearchResult`][dps-search-result] values containing each parsed date and its matching text:

Search is heuristic: it can miss dates or include surrounding words. Restrict the languages when known, and choose the search strategy that fits your input.

```go
_, dates, _ := dps.Search(nil, "The client arrived to the office for the first time "+
	"in March 3rd, 2004 and got serviced, after a couple of months, on May 6th 2004, "+
	"the customer returned indicating a defect on the part")

// Output, formatted:
// "in March 3rd, 2004 and" => time: 2004-03-03 00:00:00 UTC, period: Day
// "on May 6th 2004" => time: 2004-05-06 00:00:00 UTC, period: Day
```

If you need more control over what is being parsed, check the documentation for [`Configuration`][dps-config] and [`Parser`][dps-parser].

### Search Strategies

`SearchStrategy` defaults to `"split"`, which keeps the existing translation-and-splitting behavior. The opt-in `"ngram"` strategy tries the longest sequences of up to seven tokens first. It returns non-overlapping matches with their original text, at the cost of more parse attempts:

```go
cfg := &dps.Configuration{
	Languages:      []string{"en"},
	SearchStrategy: "ngram",
}
_, dates, err := dps.Search(cfg, "The first satellite was launched on 4 October 1957.")
```

This returns `"4 October 1957"` with the date 1957-10-04. With multiple explicit languages or locales, the default split strategy tries the detected language first, then the supplied alternatives until one produces results; it does not combine results across languages. N-gram search tries those locales for each candidate and reports the outer search's detected language on its results.

### Time Spans

Set `ReturnTimeSpan` to append start/end results for English span expressions such as `"past month"`, `"next week"`, or `"last 3 days"`. This works with either search strategy:

```go
cfg := &dps.Configuration{
	Languages:          []string{"en"},
	CurrentTime:        time.Date(2025, 2, 15, 12, 0, 0, 0, time.UTC),
	ReturnTimeSpan:     true,
	DefaultStartOfWeek: "monday",
	DefaultDaysInMonth: 30,
}
_, dates, err := dps.Search(cfg, "Messages received for the past week")
```

The span results are `"for the past week (start)"` at 2025-02-03 12:00 and `"for the past week (end)"` at 2025-02-09 12:00. Both have the `Day` period. Like upstream, boundaries retain the reference time of day; they are not midnight/end-of-day bounds.

`DefaultStartOfWeek` accepts `"monday"` (default) or `"sunday"`. An unnumbered month is a rolling `DefaultDaysInMonth` days, with zero defaulting to 30. Numbered months use calendar arithmetic with the destination day clamped to the month's last day. The first matching span pattern is returned, not every span in the text.

### Concurrent Use

A `Parser` can be reused by concurrent `Parse` and `Search` calls. Configure its exported fields before starting those calls, and do not mutate supplied configurations or their slices while calls are in flight. Custom callbacks must support concurrent invocation. Per-call date-order and relative-reference changes do not modify the caller's configuration.

## <a name="supported-languages-and-locales"></a> 5. Supported Languages and Locales [▲](#table-of-contents)

The supported languages and regional variants are listed in the [generated locale list][dps-locale-list]. The CLDR refresh retains the existing supported language set and detection ranking, including locale codes removed from newer CLDR data.

## <a name="supported-patterns"></a> 6. Supported Patterns [▲](#table-of-contents)

There are several patterns that can be parsed by **dateparser**:

- Timestamp
- Relative date
- Absolute date
- Custom Go layouts

### 6.1. Timestamp

A timestamp represents time since the Unix epoch. The parser accepts 10-digit seconds, 13-digit milliseconds, and 16-digit microseconds, with an optional minus sign. Timestamp results retain Go's `time.Local` location; use `UTC()` or `In()` to choose the display location:

```go
dt, _ := dps.Parse(nil, "1570308760263")
fmt.Println(dt.Time.UTC().Format(time.RFC3339Nano))
// Output: 2019-10-05T20:52:40.263Z
```

### 6.2. Relative Date

A relative date describes an offset such as "1 week ago" or "5 years ago". The parser detects the language and resolves the offset against `CurrentTime`:

```go
// Current time is 2015-06-01 00:00:00
cfg := &dps.Configuration{
	CurrentTime: time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC),
}

var dt date.Date
dt, _ = dps.Parse(cfg, "1 hour ago")
// time: 2015-05-31 23:00:00 UTC, period: Day

dt, _ = dps.Parse(cfg, "Il ya 2 heures") // French (2 hours ago)
// time: 2015-05-31 22:00:00 UTC, period: Day

dt, _ = dps.Parse(cfg, "1 anno 2 mesi") // Italian (1 year 2 months)
// time: 2014-04-01 00:00:00 UTC, period: Month

dt, _ = dps.Parse(cfg, "yaklaşık 23 saat önce") // Turkish (23 hours ago)
// time: 2015-05-31 01:00:00 UTC, period: Day

dt, _ = dps.Parse(cfg, "Hace una semana") // Spanish (a week ago)
// time: 2015-05-25 00:00:00 UTC, period: Day

dt, _ = dps.Parse(cfg, "2小时前") // Chinese (2 hours ago)
// time: 2015-05-31 22:00:00 UTC, period: Day
```

### 6.3. Absolute Date

An absolute date names calendar components explicitly, for example "12 August 2021" or "23 January, 15:05". The parser detects the language for each input:

```go
dt, _ = dps.Parse(nil, "12/12/12")
// time: 2012-12-12 00:00:00 UTC, period: Day

dt, _ = dps.Parse(nil, "Fri, 12 Dec 2014 10:55:50")
// time: 2014-12-12 10:55:50 UTC, period: Day

dt, _ = dps.Parse(nil, "Martes 21 de Octubre de 2014") // Spanish (Tuesday 21 October 2014)
// time: 2014-10-21 00:00:00 UTC, period: Day

dt, _ = dps.Parse(nil, "Le 11 Décembre 2014 à 09:00") // French (11 December 2014 at 09:00)
// time: 2014-12-11 09:00:00 UTC, period: Day

dt, _ = dps.Parse(nil, "13 января 2015 г. в 13:34") // Russian (13 January 2015 at 13:34)
// time: 2015-01-13 13:34:00 UTC, period: Day

dt, _ = dps.Parse(nil, "1 เดือนตุลาคม 2005, 1:00 AM") // Thai (1 October 2005, 1:00 AM)
// time: 2005-10-01 01:00:00 UTC, period: Day
```

### 6.4. Custom Formats

If you know the possible formats, supply them as Go time layouts:

```go
dt, _ = dps.Parse(nil, "22 Décembre 2010", "02 January 2006")
// time: 2010-12-22 00:00:00 UTC, period: Day
```

Ordinal-day layouts such as `"2006-002"` and `"2006 __2"` are supported and reject invalid days for the parsed year. For layouts with a two-digit year (`"06"`), `PreferredDateSource` can select a past or future century relative to `CurrentTime`. Missing-year layouts also use `CurrentTime`.

## <a name="supported-calendars"></a> 7. Supported Calendars [▲](#table-of-contents)

In addition to the Gregorian calendar, **dateparser** supports the Persian Jalali and Hijri calendars through [`dateparser.ParseJalali`][dps-jalali] and [`dateparser.ParseHijri`][dps-hijri]. Both return Gregorian `time.Time` values.

The Persian Jalali calendar, also called the Solar Hijri calendar, is used in Iran and Afghanistan.

The Hijri calendar is a lunar calendar with several variants, including [Tabular][tabular-hijri] and [Umm al-Qura][umm-al-qura]. This package uses Umm al-Qura.

```go
dt, _ = dps.ParseJalali(nil, "جمعه سی ام اسفند ۱۳۸۷")
// time: 2009-03-20 00:00:00 UTC, period: Day

dt, _ = dps.ParseHijri(nil, "17-01-1437 هـ 08:30 مساءً")
// time: 2015-10-30 20:30:00 UTC, period: Day
```

The Hijri and Jalali parsers support absolute dates only, not relative dates or Unix timestamps.

## <a name="language-based-date-order"></a> 8. Language-Based Date Order [▲](#table-of-contents)

By default, **dateparser** uses the detected locale's date order:

```go
dt, _ = dps.Parse(nil, "02-03-2016") // assumes English, uses MDY date order
// time: 2016-02-03 00:00:00 UTC, period: Day

dt, _ = dps.Parse(nil, "le 02-03-2016") // detects French, uses DMY date order
// time: 2016-03-02 00:00:00 UTC, period: Day
```

Regional date orders are used where locale data defines them. The explicit `en-US` locale uses _MDY_; it can also be selected with `Languages: []string{"en"}` and `Region: "US"`. Do not assume that every regional variant has a distinct order: use an explicit `DateOrder` when the input's convention is known.

To override the locale's convention, specify the date order in `Configuration`:

```go
cfg = &dps.Configuration{DateOrder: dps.DMY}
dt, _ = dps.Parse(cfg, "02-03-2016")
// time: 2016-03-02 00:00:00 UTC, period: Day
```

You can also specify the date order only for a specific locale:

```go
cfg = &dps.Configuration{
	DateOrder: func(locale string) string {
		if locale == "en" {
			return "DMY"
		}
		return dps.DefaultDateOrder(locale)
	},
}

dt, _ = dps.Parse(cfg, "02-03-2016") // English now uses DMY date order
// locale: en, time: 2016-03-02 00:00:00 UTC, period: Day

dt, _ = dps.Parse(cfg, "miy 02-03-2016") // Filipino keeps its MDY date order
// locale: fil, time: 2016-02-03 00:00:00 UTC, period: Day
```

## <a name="timezone-and-utc-offset"></a> 9. Timezones and UTC Offsets [▲](#table-of-contents)

For non-timestamp inputs, an explicit timezone takes precedence. Otherwise, the parser uses `DefaultTimezone`, or the location of `CurrentTime` if no default timezone is set. An unset `CurrentTime` defaults to the current time in UTC:

```go
// Parse a London news article with a reference time in Jakarta.
london, _ := time.LoadLocation("Europe/London")
jakarta, _ := time.LoadLocation("Asia/Jakarta")

cfg := &dps.Configuration{
	DefaultTimezone: london,
	CurrentTime:     time.Date(2015, 6, 1, 12, 0, 0, 0, jakarta),
}

// The timezone is explicitly specified in the input.
dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM EST")
// time: 2012-01-12 22:00:00 EST, period: Day

dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM -0500")
// time: 2012-01-12 22:00:00 UTC-05:00, period: Day

dt, _ = dps.Parse(cfg, "2 hours ago EST")
// time: 2015-05-31 22:00:00 EST, period: Day

dt, _ = dps.Parse(cfg, "2 hours ago -0500")
// time: 2015-05-31 22:00:00 UTC-05:00, period: Day

// Without an explicit timezone, the parser uses DefaultTimezone.
dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM")
// time: 2012-01-12 22:00:00 GMT, period: Day

dt, _ = dps.Parse(cfg, "2 hours ago")
// time: 2015-06-01 04:00:00 BST, period: Day
// London observes British Summer Time in June.

// Remove the default timezone to use the location of CurrentTime.
cfg.DefaultTimezone = nil

dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM")
// time: 2012-01-12 22:00:00 WIB, period: Day

dt, _ = dps.Parse(cfg, "2 hours ago")
// time: 2015-06-01 10:00:00 WIB, period: Day
```

## <a name="incomplete-date-handling"></a> 10. Incomplete Date Handling [▲](#table-of-contents)

By default, **dateparser** fills missing calendar components from `CurrentTime`, clamping invalid days to the target month's last day:

```go
// Current time is 2015-07-31 12:00:00 UTC
cfg := &dps.Configuration{
	CurrentTime: time.Date(2015, 7, 31, 12, 0, 0, 0, time.UTC),
}

dt, _ = dps.Parse(cfg, "December 2015")
// time: 2015-12-31 00:00:00 UTC (day from current time)

dt, _ = dps.Parse(cfg, "February 2020")
// time: 2020-02-29 00:00:00 UTC (day from current time, corrected for leap year)

dt, _ = dps.Parse(cfg, "December")
// time: 2015-12-31 00:00:00 UTC (year and day from current time)

dt, _ = dps.Parse(cfg, "2015")
// time: 2015-07-31 00:00:00 UTC (day and month from current time)

dt, _ = dps.Parse(cfg, "Sunday")
// time: 2015-07-26 00:00:00 UTC (the closest Sunday from current time)
```

You can change the behavior by using `PreferredMonthOfYear`, `PreferredDayOfMonth` and `PreferredDateSource` in `Configuration`:

```go
// Current time is 2015-07-10 12:00:00 UTC
cfg = &dps.Configuration{
	CurrentTime: time.Date(2015, 7, 10, 12, 0, 0, 0, time.UTC),
}

cfg.PreferredDayOfMonth = dps.Current
dt, _ = dps.Parse(cfg, "December 2015")
// time: 2015-12-10 00:00:00 UTC

cfg.PreferredDayOfMonth = dps.First
dt, _ = dps.Parse(cfg, "December 2015")
// time: 2015-12-01 00:00:00 UTC

cfg.PreferredDayOfMonth = dps.Last
dt, _ = dps.Parse(cfg, "December 2015")
// time: 2015-12-31 00:00:00 UTC
```

```go
// Current time is 2015-10-10 12:00:00 UTC
cfg = &dps.Configuration{
	CurrentTime: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC),
}

cfg.PreferredDayOfMonth = dps.Current
cfg.PreferredMonthOfYear = dps.CurrentMonth
dt, _ = dps.Parse(cfg, "2020")
// time: 2020-10-10 00:00:00 UTC

cfg.PreferredDayOfMonth = dps.Last
cfg.PreferredMonthOfYear = dps.FirstMonth
dt, _ = dps.Parse(cfg, "2020")
// time: 2020-01-31 00:00:00 UTC

cfg.PreferredDayOfMonth = dps.First
cfg.PreferredMonthOfYear = dps.LastMonth
dt, _ = dps.Parse(cfg, "2015")
// time: 2015-12-01 00:00:00 UTC
```

```go
// Current time is 2015-10-10 12:00:00 UTC
cfg = &dps.Configuration{
	CurrentTime: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC),
}

cfg.PreferredDateSource = dps.CurrentPeriod
dt, _ = dps.Parse(cfg, "March")
// time: 2015-03-10 00:00:00 UTC

cfg.PreferredDateSource = dps.Future
dt, _ = dps.Parse(cfg, "March")
// time: 2016-03-10 00:00:00 UTC

cfg.PreferredDateSource = dps.Past
dt, _ = dps.Parse(cfg, "August")
// time: 2015-08-10 00:00:00 UTC
```

Set `StrictParsing` to reject inputs that omit a day, month, or year:

```go
cfg = &dps.Configuration{StrictParsing: true}
dt, _ = dps.Parse(cfg, "March")
fmt.Println(dt.IsZero()) // true
```

For a less restrictive requirement, use `RequiredParts`, for example `[]string{"year", "month"}`. Without an explicit `DateOrder`, ambiguous numbers can be retried as a year when a year is required and a day is not. Explicit caller-provided date orders remain authoritative.

## <a name="custom-language-detector"></a> 11. Custom Language Detector [▲](#table-of-contents)

Set `Parser.DetectLanguagesFunction` to supply candidate language codes from a custom detector. Short date strings can be difficult to classify, so use `Configuration.DefaultLanguages` as a fallback.

This example uses [`lingua-go`][lingua-go] to choose between English and Indonesian, with Indonesian as a fallback when the detector supplies no usable candidate:

```go
package main

import (
	"fmt"
	"strings"

	dps "github.com/markusmobius/go-dateparser"
	"github.com/markusmobius/go-dateparser/date"
	"github.com/pemistahl/lingua-go"
)

func main() {
	detector := lingua.
		NewLanguageDetectorBuilder().
		FromLanguages(lingua.English, lingua.Indonesian).
		Build()

	parser := dps.Parser{
		DetectLanguagesFunction: func(text string) []string {
			var languages []string

			candidates := detector.ComputeLanguageConfidenceValues(text)
			for _, candidate := range candidates {
				isoCode := candidate.Language().IsoCode639_1().String()
				isoCode = strings.ToLower(isoCode)

				if candidate.Value() >= 0.9 && dps.IsKnownLocale(isoCode) {
					languages = append(languages, isoCode)
				}
			}

			return languages
		},
	}

	cfg := &dps.Configuration{DefaultLanguages: []string{"id"}}
	for _, text := range []string{"Sabtu, 13 Maret 2021", "13 Maret 2021"} {
		dt, err := parser.Parse(cfg, text)
		if err != nil {
			panic(err)
		}
		printDate(dt)
	}
	// Both inputs return locale: id, time: 2021-03-13 00:00:00 UTC, period: Day.
}

func printDate(dt date.Date) {
	fmt.Printf("locale: %s, time: %s, period: %s\n",
		dt.Locale, dt.Time.Format("2006-01-02 15:04:05 MST"), dt.Period)
}
```

## <a name="handling-false-positives"></a> 12. Handling False Positives [▲](#table-of-contents)

**dateparser** tries multiple formats and locales. Inputs that are not dates can therefore produce false positives.

To reduce the possibility of receiving false positives, make sure that:

- The input string is a valid date and it doesn't contain any other words or numbers.
- If you know the languages or locales beforehand, you can specify them in the `Configuration`.

You can also restrict the [parser types][dps-parser-type] through `Parser.ParserTypes` or change their order.

### Ignoring Surrounding Text

`IgnoreSurroundingText` is an opt-in fallback for a single date wrapped in unknown text:

```go
cfg := &dps.Configuration{
	Languages:             []string{"en"},
	IgnoreSurroundingText: true,
}
dt, err := dps.Parse(cfg, "Published on 16 April 2019")
```

The ordinary parse is tried first. On failure, unknown leading and trailing tokens are removed and the remaining date is retried. This can increase false positives: numeric prefixes, recognized date words, and unknown words inside the date are not discarded. `StrictParsing` and `RequiredParts` still apply to the remaining date. Use `Search` when extracting dates from longer text.

A recognized trailing timezone can be retained, but a leading timezone or a timezone followed by more unknown text can be discarded during this fallback. Prefer clean date strings when timezone interpretation must be unambiguous. The fixed abbreviations `BST` and `HDT` now mean UTC+01:00 and UTC-09:00; ambiguous abbreviations such as `CST` retain their existing mapping.

## <a name="performance"></a> 13. Performance [▲](#table-of-contents)

Exact whole-token locale matching uses pure Go functions generated by [re2go], enabled by default with both `CGO_ENABLED=0` and `CGO_ENABLED=1`. The generated source is checked in, so users do not need to install re2go, enable a build tag, or change configuration to receive these improvements. Malformed UTF-8 retains the existing regex behavior.

Generated matching covers the boolean whole-token matchers, including Unicode input. Additional ASCII fast paths skip impossible timezone matches and classify split tokens without regexes; these fast paths retain the original regex behavior for non-ASCII input. Actual timezone extraction, capture-based splitting, and replacement are unchanged.

### Current Measurements

The following corpus timings were measured on 2026-09-10 after the dependency refresh, using Go 1.27.1 on Windows/AMD Ryzen AI 7 PRO 350, `CGO_ENABLED=1`, and no build tags. Each configuration used the same inputs and fixed reference time, with randomized per-input execution order, one discarded warmup, and six measured rounds:

| Workload | Retained Regexes | re2go Only | Current Defaults | Total Time Reduction |
| --- | --- | --- | --- | --- |
| Parse, 762 inputs | 7.684 s | 6.677 s | 5.278 s | 31.3% |
| Search, 62 texts | 3.157 s | 3.251 s | 1.995 s | 36.8% |

Values are median time per complete corpus, not per-date latency. The baseline uses the current parser and dependencies with these optimizations disabled, not an older release or Python dateparser. All outputs matched across configurations. The additional fast paths reduced time by 20.9% for parse and 38.7% for search relative to re2go alone in this run.

Machine load varied substantially. In particular, the isolated re2go search result was noisy and 3.0% slower by these medians, so it does not establish a search speedup by itself. The complete current defaults were faster in every measured round. These figures are workload-specific, not guarantees.

The generated source is 1.89 MB. With the current dependencies, enabling generated matchers increased the root test executable from 14.75 MB to 16.81 MB, approximately 14%. See [UPSTREAM.md](UPSTREAM.md#current-performance) for the full breakdown, measurement scope, and exact sizes.

### Build Modes

| `CGO_ENABLED` | Build Tags | Generated Matchers | Remaining Regex Operations |
| --- | --- | --- | --- |
| `0` | None, `re2_wasm`, or `re2_cgo` | Enabled | Go standard engine |
| `1` | None | Enabled | Go standard engine |
| `1` | `re2_wasm` | Enabled | RE2 through WebAssembly |
| `1` | `re2_cgo` | Enabled | Native RE2 through cgo |

The remaining regex operations use Go's standard engine by default, including when cgo is enabled without a build tag. C++ [RE2][re2] is available through [go-re2] as an optional alternative, selected with the `re2_wasm` or `re2_cgo` build tag, for example:

```sh
go build -tags re2_cgo .
```

Instructions for the native RE2 dependencies are provided below.

The `re2_wasm` tag uses RE2 packaged as WebAssembly without a separately installed native RE2 library. This package's current backend-selection tags require `CGO_ENABLED=1` for either optional backend; with cgo disabled, the remaining regex operations still use Go's standard engine.

The `re2_cgo` tag uses native RE2 through cgo and requires its library and compiler toolchain to be installed.

Measure these alternatives on your workload with the existing parse/search benchmarks or [scripts/speedtest](scripts/speedtest). Neither tag changes the generated exact matchers, and a faster regex engine does not necessarily make the complete workload faster.

The optional regex backends are experimental. Validate them against your inputs before deployment.

### Compiling with cgo under Linux

On Ubuntu, install the GCC toolchain, pkg-config, and the RE2 library:

```bash
sudo apt-get install -y build-essential pkg-config libre2-dev
```

### Compiling with cgo under Windows

On Windows start by installing [MSYS2][msys2]. Then open the MINGW64 terminal and install the gcc toolchain and re2 via pacman:

```bash
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-re2
pacman -S mingw-w64-x86_64-pkg-config
```

To run the resulting executable outside the MINGW64 terminal, add the MinGW-w64 library directory to the `PATH` environment variable (adjust the path for your installation):

```cmd
SET PATH=C:\msys64\mingw64\bin;%PATH%
```

## <a name="contribution-guide"></a> 14. Contribution Guide [▲](#table-of-contents)

Issues and pull requests are welcome. Translation changes require regenerating the locale data and its exact matchers.

English is the base language for parsing. Locale-specific vocabulary is translated into English equivalents using the generated files in [internal/data](internal/data).

The generator in [scripts/codegen](scripts/codegen) uses the following sources:

- [Unicode CLDR][cldr] 44.1.0 date, core, and unit data in JSON format, automatically downloaded from the pinned tag.
- Community-maintained YAML data in [data-supplementary/date-translation](data-supplementary/date-translation).

Regeneration also compiles exact locale matchers using **re2go 4.4**, the Go backend of re2c. Install it on `PATH`, or pass its executable with `--re2go`, for example `--re2go C:/msys64/usr/bin/re2go.exe` on Windows. This is a generation-time dependency only. The converter in [scripts/codegen/exact-matcher.go](scripts/codegen/exact-matcher.go) preserves Go's Unicode case folding and rejects unsupported regex operations before replacing existing locale data.

If you wish to extend the data of an existing language, or add data for a new language, you must:

1. **Edit or create the corresponding file in [data-supplementary/date-translation](data-supplementary/date-translation).**

	See the existing files as examples, and see `SupplementaryData` in [scripts/codegen/03-structs.go](scripts/codegen/03-structs.go) for the schema. Regional overrides use `locale-specific` entries with optional `date-order` and vocabulary. Never modify [internal/data](internal/data) directly, because regeneration replaces it.

2. **Regenerate data by running:**

	```sh
	go run ./scripts/codegen --keep-language-order
	```

	`--keep-language-order` preserves the existing detection ranking instead of adopting a fresh W3Techs ranking. After the first download, add `--skip-raw` to regenerate from the cached CLDR data. Use `-h` to see all options.

3. **Write tests that cover your changes.**

	Use [parser_test.go](parser_test.go) for absolute-date translations and [parser-relative_test.go](parser-relative_test.go) for relative-date translations. Add search regressions to [search_test.go](search_test.go) when applicable.

	Follow the existing test format. Open an issue when a translation's intended behavior is unclear.

## <a name="license"></a> 15. License [▲](#table-of-contents)

Like the original, this package is licensed under the [BSD 3-Clause License][bsd3].

[go-ref-icon]: https://pkg.go.dev/badge/github.com/markusmobius/go-dateparser.svg
[go-ref]: https://pkg.go.dev/github.com/markusmobius/go-dateparser
[original]: https://github.com/scrapinghub/dateparser
[original-commit]: https://github.com/scrapinghub/dateparser/tree/9ce60b1958f1b285886bcfbb743f6419feacfc92
[original-tag]: https://github.com/scrapinghub/dateparser/releases/tag/v1.4.3
[dps-parse]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Parse
[dps-jalali]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParseJalali
[dps-hijri]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParseHijri
[dps-search]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Search
[dps-search-result]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#SearchResult
[dps-config]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Configuration
[dps-parser]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Parser
[dps-parser-type]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParserType
[dps-locale-list]: https://github.com/markusmobius/go-dateparser/blob/v1.4.3/internal/data/04-locale-order.go
[tabular-hijri]: https://en.wikipedia.org/wiki/Tabular_Islamic_calendar
[umm-al-qura]: https://webspace.science.uu.nl/~gent0113/islam/ummalqura.htm
[lingua-go]: https://github.com/pemistahl/lingua-go
[bsd3]: https://tldrlegal.com/license/bsd-3-clause-license-(revised)
[re2]: https://github.com/google/re2
[re2go]: https://re2c.org/manual/manual_go.html
[go-re2]: https://github.com/wasilibs/go-re2
[msys2]: https://www.msys2.org/
[cldr]: https://cldr.unicode.org/
