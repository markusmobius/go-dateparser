package nospace

import (
	"fmt"
	"strings"

	"github.com/itchyny/timefmt-go"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/tokenizer"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	if !isEligible(str) {
		return date.Date{}, fmt.Errorf("unable to parse date from %s", str)
	}

	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, ":", "")
	if str == "" {
		return date.Date{}, fmt.Errorf("string is empty")
	}

	var order string
	if cfg != nil && cfg.DateOrder != "" {
		order = resolveDateOrder(cfg.DateOrder)
	} else {
		order = resolveDateOrder("MDY")
		if rxEightDigit.MatchString(str) {
			dt := findBestMatchingDate(str)
			if dt.IsValid() {
				return dt, nil
			}
		}
	}

	formats, exist := dateTimeFormats[order]
	if !exist {
		return date.Date{}, fmt.Errorf("order %s is invalid", order)
	}

	var ambiguousDate date.Date
	for _, token := range tokenizer.Tokenize(str) {
		if token.Type != tokenizer.Digit {
			continue
		}

		for _, format := range formats {
			period := getPeriodFromFormat(format)
			candidates, candidateFormat := createParseCandidates(token.Text, format)

			for _, candidate := range candidates {
				t, _ := timefmt.Parse(candidate, candidateFormat)
				if t.IsZero() {
					continue
				}

				dt := date.Date{Time: t, Period: period}
				if t.Year() < 1000 {
					ambiguousDate = dt
					continue
				}

				if err := checkStrictParsing(cfg, format); err != nil {
					continue
				}

				return dt, nil
			}
		}
	}

	if ambiguousDate.IsValid() {
		return ambiguousDate, nil
	}

	return date.Date{}, fmt.Errorf("unable to parse date from %s", str)
}

func isEligible(str string) bool {
	match := rxCompatible.FindString(str)
	if match == "" || match == ":" {
		return true
	}
	return false
}

func findBestMatchingDate(str string) date.Date {
	for _, format := range preferredFormatsOrdered8Digit {
		t, _ := timefmt.Parse(format, str)
		if !t.IsZero() && t.Year() >= 1000 {
			period := getPeriodFromFormat(format)
			return date.Date{Time: t, Period: period}
		}
	}

	return date.Date{}
}

func getPeriodFromFormat(format string) date.Period {
	for _, dayKey := range []string{"%d", "%H", "%M", "%S"} {
		if strings.Contains(format, dayKey) {
			return date.Day
		}
	}

	if strings.Contains(format, "%m") {
		return date.Month
	}

	return date.Year
}

func checkStrictParsing(cfg *setting.Configuration, format string) error {
	if cfg == nil || !cfg.StrictParsing || len(cfg.RequiredParts) == 0 {
		return nil
	}

	missingParts := map[string]struct{}{}
	for key, part := range formatMapping {
		if !strings.Contains(format, key) {
			missingParts[part] = struct{}{}
		}
	}

	if len(missingParts) > 0 {
		if cfg.StrictParsing {
			return fmt.Errorf("several part of date are missing")
		}

		for _, requiredPart := range cfg.RequiredParts {
			if _, isMissing := missingParts[requiredPart]; isMissing {
				return fmt.Errorf("required part \"%s\" is missing", requiredPart)
			}
		}
	}

	return nil
}

func resolveDateOrder(order string) string {
	order = strings.ToUpper(order)
	return dateOrderMap[order]
}
