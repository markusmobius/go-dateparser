package nospace

import (
	"fmt"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/parser/tokenizer"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

// Parse parses the date string that written without spaces, for example
// 2021-10-11 is written as 20211011.
func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	// Make sure string can be parsed by this parser.
	if !isEligible(str) {
		return date.Date{}, fmt.Errorf("unable to parse date from %s", str)
	}

	// Sanitize string
	str = strutil.SanitizeSpaces(str)
	str = strings.ReplaceAll(str, ":", "")
	if str == "" {
		return date.Date{}, fmt.Errorf("string is empty")
	}

	// Fetch date order
	var order string
	if cfg != nil && cfg.DateOrder != "" {
		order = strings.ToLower(cfg.DateOrder)
	} else {
		order = strings.ToLower("MDY")
		if rxEightDigit.MatchString(str) {
			dt := findBestMatchingDate(str)
			if !dt.IsZero() {
				return dt, nil
			}
		}
	}

	formats, exist := dateTimeFormats[order]
	if !exist {
		return date.Date{}, fmt.Errorf("order %s is invalid", order)
	}

	// Process each token in string
	var ambiguousDate date.Date
	for _, token := range tokenizer.Tokenize(str) {
		// We only want digit token
		if token.Type != tokenizer.Digit {
			continue
		}

		// Try each format
		for _, format := range formats {
			period := getPeriodFromFormat(format)
			candidates, goFormat := createParseCandidates(token.Text, format)

			for _, candidate := range candidates {
				t, _ := time.Parse(goFormat, candidate)
				if t.IsZero() {
					continue
				}

				dt := date.Date{Time: t, Period: period}
				if t.Year() < 1000 {
					ambiguousDate = dt
					continue
				}

				missingParts := strutil.NewDict()
				for key, part := range formatMapping {
					if !strings.Contains(format, key) {
						missingParts.Add(part)
					}
				}

				if err := common.CheckStrictParsing(cfg, missingParts); err != nil {
					continue
				}

				return dt, nil
			}
		}
	}

	if !ambiguousDate.IsZero() {
		return ambiguousDate, nil
	}

	return date.Date{}, fmt.Errorf("unable to parse date from %s", str)
}

func isEligible(str string) bool {
	match := rxCompatible.FindString(str)
	return match == "" || match == ":"
}

func findBestMatchingDate(str string) date.Date {
	for _, format := range preferredFormatsOrdered8Digit {
		t, _ := time.Parse(format, str)
		if !t.IsZero() && t.Year() >= 1000 {
			period := getPeriodFromFormat(format)
			return date.Date{Time: t, Period: period}
		}
	}

	return date.Date{}
}

func getPeriodFromFormat(format string) date.Period {
	for _, dayKey := range []string{"02", "d", "H", "M", "S", "f"} {
		if strings.Contains(format, dayKey) {
			return date.Day
		}
	}

	for _, monthKey := range []string{"01", "m"} {
		if strings.Contains(format, monthKey) {
			return date.Month
		}
	}

	return date.Year
}
