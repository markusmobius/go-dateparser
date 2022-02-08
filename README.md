# Go-DateParser

This package provides functionality to easily parse localized dates in almost any string formats commonly found on web pages. This is Go port from the Python library with the same name.

## Key Features

- Support for almost every existing date format: absolute dates, relative dates (`"two weeks ago"` or `"tomorrow"`), timestamps, etc.
- Support for more than 200 language locales.
- Language autodetection
- Customizable behavior through settings.
- Support for non-Gregorian calendar systems.
- Support for dates with timezones abbreviations or UTC offsets (`"August 14, 2015 EST"`, `"21 July 2013 10:15 pm +0500"`, etc)
- Search dates in longer texts.

## Status

This package is up to date with the original dateparser until commit [ff439d1][0] (several commits after [v1.1.0][1]). There are several behavior and implementation differences between this port and the original:

- In Python, timezone data is not included in date and time objects. Meanwhile in Go timezone data is required. So, in this parser if the timezone is not specified in the string, parser will use timezone data from the `CurrentTime`.
- Regex in Go is pretty slow, so in this port, whenever possible we use basic strings or runes operations instead of regex to improve the performance.

## Basic Usage

Install the package inside your project:

```
go get -u -v github.com/markusmobius/go-dateparser
```

Then use the parser in your code. The most straightforward way to parse dates with **dateparser** is to use the `dateparser.Parse()` function, that wraps around most of the functionality of the module.

```go
package main

import (
	"fmt"
	"time"

	dps "github.com/markusmobius/go-dateparser"
)

func main() {
	texts := []string{
		"Fri, 12 Dec 2014 10:55:50",
		"1991-05-17",
		"In two months", // today is 1st Aug 2020
		"1484823450",    // timestamp
		"January 12, 2012 10:00 PM EST",
	}

	cfg := &dps.Configuration{
		CurrentTime: time.Date(2020, 8, 1, 0, 0, 0, 0, time.UTC),
	}

	for _, text := range texts {
		dt, _ := dps.Parse(cfg, text)
		fmt.Printf("locale: %s, time: %s, confidence: %s\n",
			dt.Locale, dt.Time, dt.Period)
	}
}

/*
Output:

locale: en, time: 2014-12-12 10:55:50 +0000 UTC, confidence: Day
locale: en, time: 1991-05-17 00:00:00 +0000 UTC, confidence: Day
locale: en, time: 2020-10-01 00:00:00 +0000 UTC, confidence: Month
locale: en, time: 2017-01-19 17:57:30 +0700 WIB, confidence: Day
locale: en, time: 2012-01-12 22:00:00 -0500 EST, confidence: Day
*/
```

As you can see, **dateparser** works with different date formats, but it can also be used directly with strings in different languages:

```go
package main

import (
	"fmt"
	"time"

	dps "github.com/markusmobius/go-dateparser"
)

func main() {
	texts := []string{
		"Martes 21 de Octubre de 2014", // Spanish (Tuesday 21 October 2014)
		"Le 11 Décembre 2014 à 09:00",  // French (11 December 2014 at 09:00)
		"13 января 2015 г. в 13:34",    // Russian (13 January 2015 at 13:34)
		"1 เดือนตุลาคม 2005, 1:00 AM",  // Thai (1 October 2005, 1:00 AM)
		"yaklaşık 23 saat önce",        // Turkish (23 hours ago), current time: 23:46
		"2小时前",                         // Chinese (2 hours ago), current time: 23:46
	}

	cfg := &dps.Configuration{
		CurrentTime: time.Date(2020, 8, 1, 23, 46, 0, 0, time.UTC),
	}

	for _, text := range texts {
		dt, _ := dps.Parse(cfg, text)
		fmt.Printf("locale: %s, time: %s, confidence: %s\n",
			dt.Locale, dt.Time, dt.Period)
	}
}

/*
Output:

locale: es, time: 2014-10-21 00:00:00 +0000 UTC, confidence: Day
locale: fr, time: 2014-12-11 09:00:00 +0000 UTC, confidence: Day
locale: ru, time: 2015-01-13 13:34:00 +0000 UTC, confidence: Day
locale: th, time: 2005-10-01 01:00:00 +0000 UTC, confidence: Day
locale: tr, time: 2020-08-01 00:46:00 +0000 UTC, confidence: Day
locale: zh, time: 2020-08-01 21:46:00 +0000 UTC, confidence: Day
*/
```

You can control multiple behaviors by using the `Configuration` parameter:

```go
package main

import (
	"fmt"
	"time"

	dps "github.com/markusmobius/go-dateparser"
	"github.com/markusmobius/go-dateparser/date"
)

func main() {
	// Initial configuration
	cfg := &dps.Configuration{
		CurrentTime: time.Date(2020, 9, 23, 0, 0, 0, 0, time.UTC),
	}

	// Custom date order
	cfg.DateOrder = "YMD"
	cfg.PreferConfigDateOrder = true
	dt, _ := dps.Parse(cfg, "2014-10-12")
	printDate(dt)

	cfg.DateOrder = "YDM"
	cfg.PreferConfigDateOrder = true
	dt, _ = dps.Parse(cfg, "2014-10-12")
	printDate(dt)

	// Has preferred date source
	cfg.PreferredDateSource = dps.Future
	dt, _ = dps.Parse(cfg, "1 year")
	printDate(dt)
}

func printDate(dt date.Date) {
	fmt.Printf("locale: %s, time: %s, confidence: %s\n",
		dt.Locale, dt.Time, dt.Period)
}

/*
Output:

locale: en, time: 2014-10-12 00:00:00 +0000 UTC, confidence: Day
locale: en, time: 2014-12-10 00:00:00 +0000 UTC, confidence: Day
locale: en, time: 2021-09-23 00:00:00 +0000 UTC, confidence: Year
*/
```

To see more, check out the advanced [docs](docs).

## False Positives

By default **dateparser** will do its best to return a date, dealing with multiple formats and different locales. For that reason it is important that the input is a valid date, otherwise it could return false positives.

To reduce the possibility of receiving false positives, make sure that:

- The input string is a valid date and it doesn't contain any other words or numbers.
- If you know the language or languages beforehand you add them through the languages or locales properties of the `Parser`. You also could exclude any of the parser method (timestamp, relative-time...) or change the order in which they are executed.

## Common Use Cases

**dateparser** can be used with a really different number of purposes, but it stands out when it comes to:

### Consuming data from different sources:

- **Scraping**: extract dates from different places with several different formats and languages
- **IoT**: consuming data coming from different sources with different date formats
- **Tooling**: consuming dates from different logs / sources
- **Format transformations**: when transforming dates coming from different files (PDF, CSV, etc.) to other formats (database, etc).

Offering natural interaction with users:

- **Tooling and CLI**: allow users to write `3 days ago` to retrieve information.
- **Search engine**: allow people to search by date in an easiest / natural format.
- **Bots**: allow users to interact with a bot easily

## License

Just like the original, this package is licensed under [BSD-3 License][2].

[0]: https://github.com/scrapinghub/dateparser/commit/ff439d1c5f87ef997951e66c31861720726faffb
[1]: https://github.com/scrapinghub/dateparser/tree/v1.1.0
[2]: https://tldrlegal.com/license/bsd-3-clause-license-(revised)
