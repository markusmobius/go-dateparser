package parser

import (
	"fmt"
	"time"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/language"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/formatted"
	"github.com/markusmobius/go-dateparser/internal/parser/nospace"
	"github.com/markusmobius/go-dateparser/internal/parser/relative"
	"github.com/markusmobius/go-dateparser/internal/parser/timestamp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

type ParseMethod func(
	cfg *setting.Configuration,
	translation string,
	translationWithFormat string,
	formats ...string,
) date.Date

func Parse(cfg *setting.Configuration, d *language.LocaleDetector, str string, formats ...string) (date.Date, error) {
	// Try to parse with specified formats
	var dt date.Date
	if dt = formatted.Parse(cfg, str, formats...); !dt.IsZero() {
		return dt, nil
	}

	// Sanitize string
	originalStr := str
	str = strutil.SanitizeDate(str)

	print("ORIGINAL STRING:", originalStr)
	print("SANITIZED:", str)

	// Try to parse as timestamp
	dt = timestamp.Parse(cfg, str)
	if !dt.IsZero() {
		return dt, nil
	}

	// Find the suitable locales for this string
	var locales []*data.LocaleData
	if d != nil {
		var err error
		locales, err = d.Detect(str)
		if err != nil {
			return date.Date{}, err
		}
	}
	print("SUITABLE LOCALES:")

	for i := range locales {
		print("\t", locales[i].Name)
	}

	// Process each locale
	parseMethods := []ParseMethod{
		tryRelativeParser,
		tryFormattedParser,
		tryAbsoluteParser,
		tryNospaceParser,
	}

	for _, locale := range locales {
		print("PARSE USING LOCALE:", locale.Name)

		// Clone config
		clonedCfg := cfg.Clone()
		if clonedCfg.DateOrder == "" || clonedCfg.PreferLocaleDateOrder {
			print("\tYOOOO:", clonedCfg.DateOrder, locale.Name, locale.DateOrder)
			clonedCfg.DateOrder = locale.DateOrder
		}

		// Translate string
		translation := language.Translate(&clonedCfg, locale, str, false)
		translationWithFormat := language.Translate(&clonedCfg, locale, str, true)

		print("\tTRANSLATION:", translation)

		for i := range parseMethods {
			dt = parseMethods[i](&clonedCfg, translation, translationWithFormat)
			if !dt.IsZero() {
				dt.Locale = locale.Name
				return dt, nil
			}
		}
	}

	return date.Date{}, fmt.Errorf("failed to parse \"%s\": unknown format", originalStr)
}

func tryRelativeParser(cfg *setting.Configuration, t string, _ string, _ ...string) date.Date {
	return relative.Parse(cfg, t)
}

func tryFormattedParser(cfg *setting.Configuration, _ string, twf string, formats ...string) date.Date {
	return formatted.Parse(cfg, twf, formats...)
}

func tryAbsoluteParser(cfg *setting.Configuration, t string, _ string, _ ...string) date.Date {
	if t, tz := stripBracesAndTimezones(t); t != "" {
		dt, _ := absolute.Parse(cfg, t)
		return applyTimezone(dt, tz)
	}
	return date.Date{}
}

func tryNospaceParser(cfg *setting.Configuration, t string, _ string, _ ...string) date.Date {
	if t, tz := stripBracesAndTimezones(t); t != "" {
		dt, _ := nospace.Parse(cfg, t)
		return applyTimezone(dt, tz)
	}
	return date.Date{}
}

func stripBracesAndTimezones(s string) (string, timezone.TimezoneOffsetData) {
	s = strutil.StripBraces(s)
	return timezone.PopTzOffset(s)
}

func applyTimezone(dt date.Date, tz timezone.TimezoneOffsetData) date.Date {
	if dt.IsZero() || tz.IsZero() {
		return dt
	}

	dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(),
		dt.Time.Hour(), dt.Time.Minute(), dt.Time.Second(), dt.Time.Nanosecond(),
		time.FixedZone(tz.Name, tz.Offset))

	return dt
}

func print(args ...interface{}) {
	// fmt.Println(args...)
}
