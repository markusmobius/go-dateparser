# Go-DateParser [![Go Reference][go-ref-icon]][go-ref]

This package provides functionality to easily parse localized dates in almost any string formats commonly found on web pages. This is Go port from the [Python library][original] with the same name.

To use it, install the package inside your project:

```
go get -u -v github.com/markusmobius/go-dateparser
```

## Table of Contents

- [ 1. Features ▲](#-1-features-)
- [ 2. Status ▲](#-2-status-)
- [ 3. Common Use Cases ▲](#-3-common-use-cases-)
- [ 4. Usage ▲](#-4-usage-)
- [ 5. Supported Languages and Locales ▲](#-5-supported-languages-and-locales-)
- [ 6. Supported Patterns ▲](#-6-supported-patterns-)
  - [6.1. Timestamp](#61-timestamp)
  - [6.2. Relative Date](#62-relative-date)
  - [6.3. Absolute Date](#63-absolute-date)
  - [6.4. Custom Specified Formats](#64-custom-specified-formats)
- [ 7. Supported Calendars ▲](#-7-supported-calendars-)
- [ 8. Language Based Date Order ▲](#-8-language-based-date-order-)
- [ 9. Timezone And UTC Offset ▲](#-9-timezone-and-utc-offset-)
- [ 10. Incomplete Date Handling ▲](#-10-incomplete-date-handling-)
- [ 11. Custom Language Detector ▲](#-11-custom-language-detector-)
- [ 12. Handling False Positives ▲](#-12-handling-false-positives-)
- [ 13. Performance ▲](#-13-performance-)
  - [Compiling with cgo under Linux](#compiling-with-cgo-under-linux)
  - [Compiling with cgo under Windows](#compiling-with-cgo-under-windows)
- [ 14. Contribution Guide ▲](#-14-contribution-guide-)
- [ 15. License ▲](#-15-license-)

## <a name="features"></a> 1. Features [▲](#table-of-contents)

- Generic parsing of dates in over 200 language locales plus numerous formats in a language agnostic fashion.
- Generic parsing of relative dates like: `"1 min ago"`, `"2 weeks ago"`, `"3 months, 1 week and 1 day ago"`, `"in 2 days"`, `"tomorrow"`.
- Generic parsing of dates with time zones abbreviations or UTC offsets like: `"August 14, 2015 EST"`, `"July 4, 2013 PST"`, `"21 July 2013 10:15 pm +0500"`.
- Date lookup in longer texts.
- Support for non-Gregorian calendar systems. See [Supported Calendars](#supported-calendars).
- Extensive test coverage.

## <a name="status"></a> 2. Status [▲](#table-of-contents)

This package is up to date with the original dateparser until commit [748e48a][original-commit] (several commits after [v1.2.0][original-tag]). There are several behavior and implementation differences between this port and the original:

- In Python, timezone data is not included in date and time objects. Meanwhile in Go timezone data is required.
- Regex in Go is pretty slow, so in this port whenever possible we use basic strings or runes operations instead of regex to improve the performance.

## <a name="common-use-cases"></a> 3. Common Use Cases [▲](#table-of-contents)

**dateparser** can be used for many different purposes, but it stands out when it comes to:

Consuming data from different sources:

- **Scraping**: extract dates from different places with several different formats and languages
- **IoT**: consuming data coming from different sources with different date formats
- **Tooling**: consuming dates from different logs / sources
- **Format transformations**: when transforming dates coming from different files (PDF, CSV, etc.) to other formats (database, etc).

Offering natural interaction with users:

- **Tooling and CLI**: allow users to write `3 days ago` to retrieve information.
- **Search engine**: allow people to search by date in an easiest / natural format.
- **Bots**: allow users to interact with a bot easily

## <a name="usage"></a> 4. Usage [▲](#table-of-contents)

The most straightforward way is to use the [`dateparser.Parse`][dps-parse] function that wraps around most of the functionality:

```go
dt, err := dateparser.Parse(nil, "6 July 2020")
// locale: en, time: 2020-07-06 00:00:00, confidence: Day
```

You can also extract dates from longer strings of text by using [`dateparser.Search`][dps-search]. The extracted dates are returned as list of [`SearchResult`][dps-search-result] containing data of the extracted date and its original string:

> Support for searching dates is really limited and needs a lot of improvement.

```go
_, dates, _ := dps.Search(nil, "The client arrived to the office for the first time "+
	"in March 3rd, 2004 and got serviced, after a couple of months, on May 6th 2004, "+
	"the customer returned indicating a defect on the part")

// Output, formatted:
// "in March 3rd, 2004 and" => time: 2004-03-03 00:00:00, confidence: Day
// "on May 6th 2004" => time: 2004-05-06 00:00:00, confidence: Day
```

If you need more control over what is being parsed, check the documentation for [`Configuration`][dps-config] and [`Parser`][dps-parser].

## <a name="supported-languages-and-locales"></a> 5. Supported Languages and Locales [▲](#table-of-contents)

You can check the supported locales by checking out [this code][dps-locale-list]. In total there are almost 300 locales across 205 languages.

## <a name="supported-patterns"></a> 6. Supported Patterns [▲](#table-of-contents)

There are several patterns that can be parsed by **dateparser**:

- Timestamp
- Relative date
- Absolute date
- Custom specified formats

### 6.1. Timestamp

Timestamp is the date time string in Unix time format:

```go
dt, _ := dps.Parse(nil, "1570308760263")
// time: 2019-10-05 20:52:40 +0000, confidence: Day
```

### 6.2. Relative Date

Relative date is pattern where the date is specified in relative value like "1 week ago" or "5 years ago". **Dateparser** will try to parse a date from the given string, attempting to detect the language for each time:

```go
// Current time is 2015-06-01 00:00:00
cfg := &dps.Configuration{
	CurrentTime: time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC),
}

var dt date.Date
dt, _ = dps.Parse(cfg, "1 hour ago")
// time: 2015-05-31 23:00:00 UTC, confidence: Day

dt, _ = dps.Parse(cfg, "Il ya 2 heures") // French (2 hours ago)
// time: 2015-05-31 22:00:00 UTC, confidence: Day

dt, _ = dps.Parse(cfg, "1 anno 2 mesi") // Italian (1 year 2 months)
// time: 2014-04-01 00:00:00 UTC, confidence: Month

dt, _ = dps.Parse(cfg, "yaklaşık 23 saat önce") // Turkish (23 hours ago)
// time: 2015-05-31 01:00:00 UTC, confidence: Day

dt, _ = dps.Parse(cfg, "Hace una semana") // Spanish (a week ago)
// time: 2015-05-25 00:00:00 UTC, confidence: Day

dt, _ = dps.Parse(cfg, "2小时前") // Chinese (2 hours ago)
// time: 2015-05-31 22:00:00 UTC, confidence: Day
```

### 6.3. Absolute Date

Absolute date is pattern where the date (and time) is explicitly mentioned, e.g. "12 August 2021" or "23 January, 15:05". Like the relative time, **dateparser** will try to detect the language for each time:

```go
dt, _ = dps.Parse(nil, "12/12/12")
// time: 2012-12-12 00:00:00 UTC, confidence: Day

dt, _ = dps.Parse(nil, "Fri, 12 Dec 2014 10:55:50")
// time: 2014-12-12 10:55:50 UTC, confidence: Day

dt, _ = dps.Parse(nil, "Martes 21 de Octubre de 2014") // Spanish (Tuesday 21 October 2014)
// time: 2014-10-21 00:00:00 UTC, confidence: Day

dt, _ = dps.Parse(nil, "Le 11 Décembre 2014 à 09:00") // French (11 December 2014 at 09:00)
// time: 2014-12-11 09:00:00 UTC, confidence: Day

dt, _ = dps.Parse(nil, "13 января 2015 г. в 13:34") // Russian (13 January 2015 at 13:34)
// time: 2015-01-13 13:34:00 UTC, confidence: Day

dt, _ = dps.Parse(nil, "1 เดือนตุลาคม 2005, 1:00 AM") // Thai (1 October 2005, 1:00 AM)
// time: 2005-10-01 01:00:00 UTC, confidence: Day
```

### 6.4. Custom Specified Formats

If you know the possible formats of the dates, you can specify it yourself using the Go time's layout:

```go
dt, _ = dps.Parse(nil, "22 Décembre 2010", "02 January 2006")
// time: 2010-12-22 00:00:00 UTC, confidence: Day
```

## <a name="supported-calendars"></a> 7. Supported Calendars [▲](#table-of-contents)

Apart from the Georgian calendar, **dateparser** also supports the Persian Jalali calendar and the Hijri calendar which can be used by calling [`dateparser.ParseJalali`][dps-jalali] and [`dateparser.ParseHijri`][dps-hijri].

Persian Jalali calendar is also often called Solar Hijri calendar and commonly used in Iran and Afghanistan.

Hijri calendar is a Lunar calendar which commonly used in Islamic countries. Some only use it for religious purposes (e.g. calculating when Ramadhan or Hajj started) while some also use it for both administrative purposes.

There are several variations of Hijri calendar, e.g. [Tabular][tabular-hijri] and [Umm al-Qura][umm-al-qura]. In **dateparser** we use Umm al-Qura calendar since it's the one that apparently mostly used by Islamic world.

```go
dt, _ = dps.ParseJalali(nil, "جمعه سی ام اسفند ۱۳۸۷")
// time: 2009-03-20 00:00:00, confidence: Day

dt, _ = dps.ParseHijri(nil, "17-01-1437 هـ 08:30 مساءً")
// time: 2015-10-30 20:30:00, confidence: Day
```

> Note: Hijri and Jalali parser on support the absolute date pattern, so relative and timestamp pattern won't work.

## <a name="language-based-date-order"></a> 8. Language Based Date Order [▲](#table-of-contents)

By default **dateparser** will use date order from the detected language:

```go
dt, _ = dps.Parse(nil, "02-03-2016") // assumes english language, uses MDY date order
// time: 2016-02-03 00:00:00 UTC, confidence: Day

dt, _ = dps.Parse(nil, "le 02-03-2016") // detects french, uses DMY date order
// time: 2016-03-02 00:00:00 UTC, confidence: Day
```

However, ordering is not locale based. So, since most English speakers come from America which uses _MDY_ order, it will be used as default order. That's why you can't expect to use _DMY_ order for UK/Australia English by default.

To solve this issue, you can explicitly specify the date order in `Configuration`:

```go
cfg = &dps.Configuration{DateOrder: dps.DMY}
dt, _ = dps.Parse(cfg, "02-03-2016")
// time: 2016-03-02 00:00:00 UTC, confidence: Day
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

dt, _ = dps.Parse(cfg, "02-03-2016") // english language, now use DMY date order
// locale: en, time: 2016-03-02 00:00:00 UTC, confidence: Day

dt, _ = dps.Parse(cfg, "miy 02-03-2016") // philippines, keep using MDY date order
// locale: fil, time: 2016-02-03 00:00:00 UTC, confidence: Day
```

## <a name="timezone-and-utc-offset"></a> 9. Timezone And UTC Offset [▲](#table-of-contents)

By default, **dateparser** uses the timezone that is present in the date string. If the timezone is not present, then it will use timezone from the `CurrentTime`:

```go
// Use case: I'm from Jakarta, parsing London's news article.
// So, my current time use GMT+7, however the parser's default timezone is GMT.
london, _ := time.LoadLocation("Europe/London")
jakarta, _ := time.LoadLocation("Asia/Jakarta")

cfg := &dps.Configuration{
	DefaultTimezone: london,
	CurrentTime:     time.Date(2015, 6, 1, 12, 0, 0, 0, jakarta),
}

// Here the timezone is explicitly mentioned in string
dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM EST")
// time: 2012-01-12 22:00:00 EST, confidence: Day

dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM -0500")
// time: 2012-01-12 22:00:00 UTC-05:00, confidence: Day

dt, _ = dps.Parse(cfg, "2 hours ago EST")
// time: 2015-05-31 22:00:00 EST, confidence: Day

dt, _ = dps.Parse(cfg, "2 hours ago -0500")
// time: 2015-05-31 22:00:00 UTC-05:00, confidence: Day

// Now timezone is not specified, so parser will use `DefaultTimezone`
dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM")
// time: 2012-01-12 22:00:00 GMT, confidence: Day

dt, _ = dps.Parse(cfg, "2 hours ago")
// time: 2015-06-01 04:00:00 BST, confidence: Day
// notice the timezone become BST, because we move back to Daylight Saving Time.

// Now we remove the default timezone, so it use timezone from `CurrentTime`
cfg.DefaultTimezone = nil

dt, _ = dps.Parse(cfg, "January 12, 2012 10:00 PM")
// time: 2012-01-12 22:00:00 WIB, confidence: Day

dt, _ = dps.Parse(cfg, "2 hours ago")
// time: 2015-06-01 10:00:00 WIB, confidence: Day
```

## <a name="incomplete-date-handling"></a> 10. Incomplete Date Handling [▲](#table-of-contents)

By default **dateparser** will try to fill incomplete dates with value from current time:

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

You can also ignore parsing incomplete dates altogether by setting `StrictParsing` in config:

```go
cfg = &dps.Configuration{StrictParsing: true}
dt, _ = dps.Parse(cfg, "March")
fmt.Println(dt.IsZero()) // true
```

## <a name="custom-language-detector"></a> 11. Custom Language Detector [▲](#table-of-contents)

**Dateparser** allows to use a language detection behavior by using the `DetectLanguagesFunction` in the `Parser`. Using this, you can use any language detectors that you want. However since language detection often fail or inaccurate for short strings, it's highly recommended to use `DetectLanguagesFunction` while specifying `DefaultLanguages` in `Configuration`.

Here is an example on how to use [`lingua-go`][lingua-go] as detector:

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
	// Prepare detector
	detector := lingua.
		NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()

	// Create custom Parser
	p := dps.Parser{
		DetectLanguagesFunction: func(s string) []string {
			var languages []string

			candidates := detector.ComputeLanguageConfidenceValues(s)
			for _, c := range candidates {
				isoCode := c.Language().IsoCode639_1().String()
				isoCode = strings.ToLower(isoCode)

				if c.Value() >= 0.9 && dps.IsKnownLocale(isoCode) {
					languages = append(languages, isoCode)
				}
			}

			return languages
		},
	}

	// Use the custom Parser to parse
	dt, _ := p.Parse(nil, "Sabtu, 13 Maret 2021")
	printDate(dt) // locale: id, time: 2021-03-13 00:00:00 UTC, confidence: Day

	// It will fail for short strings
	dt, _ = p.Parse(nil, "13 Maret 2021")
	fmt.Println(dt.IsZero()) // true

	// So we need to specify default languages
	cfg := &dps.Configuration{DefaultLanguages: []string{"id"}}
	dt, _ = p.Parse(cfg, "13 Maret 2021")
	printDate(dt) // locale: id, time: 2021-03-13 00:00:00 UTC, confidence: Day
}

func printDate(dt date.Date) {
	fmt.Printf("locale: %s, time: %s, confidence: %s\n",
		dt.Locale, dt.Time.Format("2006-01-02 15:04:05 MST"), dt.Period)
}
```

## <a name="handling-false-positives"></a> 12. Handling False Positives [▲](#table-of-contents)

By default **dateparser** will do its best to return a date, dealing with multiple formats and different locales. For that reason it is important that the input is a valid date, otherwise it could return false positives.

To reduce the possibility of receiving false positives, make sure that:

- The input string is a valid date and it doesn't contain any other words or numbers.
- If you know the languages or locales beforehand, you can specify them in the `Configuration`.

Besides those, you also could exclude any of the [parser method][dps-parser-type] (timestamp, relative-time...) or change the order in which they are executed.

## <a name="performance"></a> 13. Performance [▲](#table-of-contents)

This library heavily uses regular expression for various purposes, from detecting language, extracting relevant content, and translating the date value. Unfortunately, as commonly known, Go's regular expression is pretty slow. This is because:

- The regex engine in other language usually implemented in C, while in Go it's implemented from scratch in Go language. As expected, C implementation is still faster than Go's.
- Since Go is usually used for web service, its regex is designed to finish in time linear to the length of the input, which useful for protecting server from ReDoS attack. However, this comes with performance cost.

If you want to parse a huge amount of data, it would be preferrable to have a better performance. So, this package provides C++ [`re2`][re2] as an alternative regex engine using binding from [go-re2]. To activate it, you can build your app using tag `re2_wasm` or `re2_cgo`, for example:

```
go build -tags re2_cgo .
```

More detailed instructions in how to prepare your system for compiling with cgo are provided below.

When using `re2_wasm` tag, it will make your app uses `re2` that packaged as WebAssembly module so it should be runnable even without cgo. However, if your input is too small, it might be even slower than using Go's standard regex engine.

When using `re2_cgo` tag, it will make your app uses `re2` library that wrapped using cgo. It's a lot faster than Go's standard regex and `re2_wasm`, however to use it cgo must be available and `re2` should be installed in your system.

In my test, `re2_cgo` will always be faster than the standard library, however `re2_wasm` only faster than standard regex when processing more than 10,000 short strings (e.g. "26 gennaio 2014"). So, when possible you might be better using `re2_cgo`. You could run `scripts/speedtest` to test it yourself.

Do note that this alternative regex engine is experimental, so use on your own risk.

### Compiling with cgo under Linux

On Ubuntu install the gcc tool chain and the re2 library as follows:

```bash
sudo apt install build-essential
sudo apt-get install -y libre2-dev
```

### Compiling with cgo under Windows

On Windows start by installing [MSYS2][msys2]. Then open the MINGW64 terminal and install the gcc toolchain and re2 via pacman:

```bash
pacman -S mingw-w64-x86_64-gcc
pacman -S mingw-w64-x86_64-re2
```

If you want to run the resulting exe program outside the MINGW64 terminal you need to add a path to the MinGW-w64 libraries to the PATH environmental variable (adjust as needed for your system):

```cmd
SET PATH=C:\msys64\mingw64\bin;%PATH%
```

## <a name="contribution-guide"></a> 14. Contribution Guide [▲](#table-of-contents)

Like another open source project, feel free to open issues and PRs to improve this project. However, special care is needed in case you want to submit additional date translation.

In DateParser, English is used as the base language. So, every submitted dates will be translated into their English equivalents before they are parsed. The language data that used by DateParser to translate are located in [`internal/data`](internal/data/) directory. For each supported language, there is a Go file containing translation data.

Those language data are generated by scripts in [`scripts/codegen`](scripts/codegen/), which use raw data from the following sources:

- [Unicode CLDR][cldr] data in JSON format, which will be automatically downloaded by the scripts.
- Additional data from community in YAML format which located in [`data-supplementary/date-translation`](data-supplementary/date-translation) directory.

If you wish to extend the data of an existing language, or add data for a new language, you must:

1. **Edit or create the corresponding file within [`data-supplementary/date-translation`](data-supplementary/date-translation) directory.**

   See the existing files as examples, and see `SupplementaryData` struct in `scripts/codegen/03-structs.go` to learn how they defined. Make sure to never modify `internal/data` directly, because changes in that directory will be removed by code generator scripts.

2. **Regenerate data by running:**

   ```
   go run scripts/codegen/*.go
   ```

   Use `-h` flag to see the command guide.

3. **Write tests that cover your changes.**

   You should put your test either in `parser_test.go` or `parser-relative_test.go`, depending on what kind of new translation that you make. The former file is for absolute time while the latter file is for relative time.

   Just follow the test format in those files and you'll be good. If in doubt, feel free to put a new issues.

## <a name="license"></a> 15. License [▲](#table-of-contents)

Just like the original, this package is licensed under [BSD-3 License][bsd3].

[go-ref-icon]: https://pkg.go.dev/badge/github.com/markusmobius/go-dateparser.svg
[go-ref]: https://pkg.go.dev/github.com/markusmobius/go-dateparser
[original]: https://github.com/scrapinghub/dateparser
[original-commit]: https://github.com/scrapinghub/dateparser/tree/748e48a
[original-tag]: https://github.com/scrapinghub/dateparser/releases/tag/v1.2.0
[dps-parse]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Parse
[dps-jalali]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParseJalali
[dps-hijri]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParseHijri
[dps-search]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Search
[dps-search-result]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#SearchResult
[dps-config]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Configuration
[dps-parser]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#Parser
[dps-parser-type]: https://pkg.go.dev/github.com/markusmobius/go-dateparser#ParserType
[dps-locale-list]: https://github.com/markusmobius/go-dateparser/blob/main/internal/data/04-locale-order.go
[tabular-hijri]: https://en.wikipedia.org/wiki/Tabular_Islamic_calendar
[umm-al-qura]: https://webspace.science.uu.nl/~gent0113/islam/ummalqura.htm
[lingua-go]: https://github.com/pemistahl/lingua-go
[bsd3]: https://tldrlegal.com/license/bsd-3-clause-license-(revised)
[re2]: https://github.com/google/re2
[go-re2]: https://github.com/wasilibs/go-re2
[msys2]: https://www.msys2.org/
[cldr]: https://cldr.unicode.org/
