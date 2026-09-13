package dateparser

import (
	"strconv"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateutil/v2/relativedelta"
)

const timeSpanPrefix = `(?i)\b(?:for\s+the\s+|during\s+the\s+|in\s+the\s+)?`

var timeSpanPatterns = []struct {
	Pattern *regexp.Regexp
	Unit    string
	Future  bool
}{
	{regexp.MustCompile(timeSpanPrefix + `(?:past|last|previous)\s+month\b`), "month", false},
	{regexp.MustCompile(timeSpanPrefix + `(?:past|last|previous)\s+week\b`), "week", false},
	{regexp.MustCompile(timeSpanPrefix + `(?:past|last|previous)\s+(\d+)\s+days?\b`), "days", false},
	{regexp.MustCompile(timeSpanPrefix + `(?:past|last|previous)\s+(\d+)\s+weeks?\b`), "weeks", false},
	{regexp.MustCompile(timeSpanPrefix + `(?:past|last|previous)\s+(\d+)\s+months?\b`), "months", false},
	{regexp.MustCompile(timeSpanPrefix + `(?:next|coming|following)\s+month\b`), "month", true},
	{regexp.MustCompile(timeSpanPrefix + `(?:next|coming|following)\s+week\b`), "week", true},
	{regexp.MustCompile(timeSpanPrefix + `(?:next|coming|following)\s+(\d+)\s+days?\b`), "days", true},
	{regexp.MustCompile(timeSpanPrefix + `(?:next|coming|following)\s+(\d+)\s+weeks?\b`), "weeks", true},
	{regexp.MustCompile(timeSpanPrefix + `(?:next|coming|following)\s+(\d+)\s+months?\b`), "months", true},
}

func searchTimeSpan(cfg *Configuration, lang, text string) []SearchResult {
	for _, pattern := range timeSpanPatterns {
		match := pattern.Pattern.FindStringSubmatch(text)
		if len(match) == 0 {
			continue
		}

		number := 1
		if len(match) > 1 {
			var err error
			number, err = strconv.Atoi(match[1])
			if err != nil {
				return nil
			}
		}
		if number > 366*10000 || pattern.Unit == "month" && cfg.DefaultDaysInMonth > 366*10000 {
			return nil
		}

		base := cfg.CurrentTime
		if base.IsZero() {
			base = time.Now().UTC()
		}
		start, end := base, base
		direction := -1
		if pattern.Future {
			direction = 1
		}

		boundary := base
		var arithmeticError error
		switch pattern.Unit {
		case "month":
			boundary, arithmeticError = (relativedelta.Delta{Days: float64(direction * cfg.DefaultDaysInMonth)}).Apply(base)
		case "week":
			daysBack := (int(base.Weekday()) + 6) % 7
			if cfg.DefaultStartOfWeek == "sunday" {
				daysBack = int(base.Weekday())
			}
			weekStart := base.AddDate(0, 0, -daysBack)
			start = weekStart.AddDate(0, 0, direction*7)
			end = start.AddDate(0, 0, 6)
		case "days":
			boundary = base.AddDate(0, 0, direction*number)
		case "weeks":
			if number > int(^uint(0)>>1)/7 {
				return nil
			}
			boundary = base.AddDate(0, 0, direction*number*7)
		case "months":
			boundary, arithmeticError = (relativedelta.Delta{Months: float64(direction * number)}).Apply(base)
		}
		if arithmeticError != nil {
			return nil
		}
		if pattern.Unit != "week" {
			if pattern.Future {
				end = boundary
			} else {
				start = boundary
			}
		}
		if start.Year() < 1 || start.Year() > 9999 || end.Year() < 1 || end.Year() > 9999 {
			return nil
		}

		return []SearchResult{
			{Date: date.Date{Time: start, Period: date.Day, Locale: lang}, Text: match[0] + " (start)"},
			{Date: date.Date{Time: end, Period: date.Day, Locale: lang}, Text: match[0] + " (end)"},
		}
	}
	return nil
}
