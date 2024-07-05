package relative

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var (
	rxNonWord          = regexp.MustCompile(`\W`)
	rxIn               = regexp.MustCompile(`(?i)\bin\b`)
	rxAgo              = regexp.MustCompile(`(?i)\bago\b`)
	rxInAgo            = regexp.MustCompile(`(?i)\b(?:ago|in)\b`)
	rxRelativePattern  = regexp.MustCompile(`(?i)(\d+[.,]?\d*)\s*(` + relativeUnits + `)\b`)
	rxRelativeSkipWord = regexp.MustCompile(`(?i)^(?:` + relativeUnits + `|ago|in|\d+|:|[ap]m)`)
	relativeUnits      = `decade|year|month|week|day|hour|minute|second`
)

// Parse parses date string like "1 year, 2 months ago" and "3 hours, 50 minutes ago".
func Parse(cfg *setting.Configuration, str string) date.Date {
	// Prepare string
	str = strutil.StripBraces(str)
	str, tzData := timezone.PopTzOffset(str)

	// Parse time
	t, _ := parseTime(str)

	// Find current time
	now := time.Now().UTC()
	if cfg != nil && !cfg.CurrentTime.IsZero() {
		now = cfg.CurrentTime
	}

	// Apply timezone
	if !tzData.IsZero() {
		loc := time.FixedZone(tzData.Name, tzData.Offset)
		now = now.In(loc)
	} else if cfg != nil && cfg.DefaultTimezone != nil {
		now = now.In(cfg.DefaultTimezone)
	}

	// Get relative date
	dt, period := parseDate(cfg, str, now)

	if !dt.IsZero() && !t.IsZero() {
		dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			dt.Location())
		if period > date.Hour {
			period = date.Hour
		}
	}

	if cfg != nil && !cfg.ReturnTimeAsPeriod && period.IsTime() {
		period = date.Day
	}

	// Create date data
	return date.Date{Time: dt, Period: period}
}

func parseTime(s string) (time.Time, error) {
	s = rxRelativePattern.ReplaceAllString(s, "")
	s = rxInAgo.ReplaceAllString(s, "")
	return common.ParseTime(s)
}

func parseDate(cfg *setting.Configuration, str string, now time.Time) (time.Time, date.Period) {
	if !allWordsAreUnits(str) {
		return time.Time{}, 0
	}

	// Retrieve relative durations
	relDurations := getRelativeDurations(str)
	if len(relDurations) == 0 {
		return time.Time{}, 0
	}

	// Extract period from relative durations
	period := date.Day

	switch {
	case keyExist(relDurations, "second"):
		period = date.Second
	case keyExist(relDurations, "minute"):
		period = date.Minute
	case keyExist(relDurations, "hour"):
		period = date.Hour
	case keyExist(relDurations, "day"):
		period = date.Day
	case keyExist(relDurations, "month"):
		period = date.Month
	case keyExist(relDurations, "year"):
		period = date.Year
	}

	// Convert relative durations (which in float64) to usable format
	year := int(relDurations["year"])
	month := int(relDurations["month"])
	day := int(relDurations["day"])
	hour := time.Duration(relDurations["hour"]) * time.Hour
	minute := time.Duration(relDurations["minute"]) * time.Minute
	second := time.Duration(relDurations["second"]) * time.Second

	date := now
	if (rxIn.MatchString(str) || cfg.PreferredDateSource == setting.Future) && !rxAgo.MatchString(str) {
		date = date.AddDate(year, month, day)
		date = date.Add(hour)
		date = date.Add(minute)
		date = date.Add(second)
	} else {
		date = date.AddDate(-year, -month, -day)
		date = date.Add(-hour)
		date = date.Add(-minute)
		date = date.Add(-second)
	}

	return date, period
}

func allWordsAreUnits(s string) bool {
	s = strutil.SanitizeSpaces(s)

	var wordCount int
	for _, word := range rxNonWord.Split(s, -1) {
		if word == "" {
			continue
		}

		if !rxRelativeSkipWord.MatchString(word) {
			wordCount++
		}
	}

	return wordCount == 0
}

func getRelativeDurations(s string) map[string]float64 {
	// Extract durations using regex
	floatDurations := map[string]float64{}
	for _, parts := range rxRelativePattern.FindAllStringSubmatch(s, -1) {
		period := parts[2]
		strValue := strings.Replace(parts[1], ",", ".", -1)
		value, _ := strconv.ParseFloat(strValue, 64)
		floatDurations[period] = value
	}

	// Convert decade to year
	if decades, exist := floatDurations["decade"]; exist {
		floatDurations["year"] += decades * 10
		delete(floatDurations, "decade")
	}

	// Convert week to days
	if week, exist := floatDurations["week"]; exist {
		floatDurations["day"] += week * 7
		delete(floatDurations, "week")
	}

	// Convert fractional values to lower unit
	for _, unit := range strings.Split(relativeUnits, "|") {
		// Make sure duration exist
		value, exist := floatDurations[unit]
		if !exist {
			continue
		}

		// If value doesn't have fractional unit, don't change anything
		isNegative := value < 0
		value, fraction, hasFraction := splitFraction(value)
		if !hasFraction {
			continue
		}

		switch unit {
		case "year":
			year := value
			month, fraction, _ := splitFraction(fraction * 12)
			day, fraction, _ := splitFraction(fraction * 30)
			hour, fraction, _ := splitFraction(fraction * 24)
			minute, fraction, _ := splitFraction(fraction * 60)
			second, _, _ := splitFraction(fraction * 60)

			if isNegative {
				year, month, day = -year, -month, -day
				hour, minute, second = -hour, -minute, -second
			}

			floatDurations["year"] = year
			addMapValue(floatDurations, "month", month)
			addMapValue(floatDurations, "day", day)
			addMapValue(floatDurations, "hour", hour)
			addMapValue(floatDurations, "minute", minute)
			addMapValue(floatDurations, "second", second)

		case "month":
			month := value
			day, fraction, _ := splitFraction(fraction * 30)
			hour, fraction, _ := splitFraction(fraction * 24)
			minute, fraction, _ := splitFraction(fraction * 60)
			second, _, _ := splitFraction(fraction * 60)

			if isNegative {
				month, day = -month, -day
				hour, minute, second = -hour, -minute, -second
			}

			floatDurations["month"] = month
			addMapValue(floatDurations, "day", day)
			addMapValue(floatDurations, "hour", hour)
			addMapValue(floatDurations, "minute", minute)
			addMapValue(floatDurations, "second", second)

		case "day":
			day := value
			hour, fraction, _ := splitFraction(fraction * 24)
			minute, fraction, _ := splitFraction(fraction * 60)
			second, _, _ := splitFraction(fraction * 60)

			if isNegative {
				day = -day
				hour, minute, second = -hour, -minute, -second
			}

			floatDurations["day"] = day
			addMapValue(floatDurations, "hour", hour)
			addMapValue(floatDurations, "minute", minute)
			addMapValue(floatDurations, "second", second)

		case "hour":
			hour := value
			minute, fraction, _ := splitFraction(fraction * 60)
			second, _, _ := splitFraction(fraction * 60)

			if isNegative {
				hour, minute, second = -hour, -minute, -second
			}

			floatDurations["hour"] = hour
			addMapValue(floatDurations, "minute", minute)
			addMapValue(floatDurations, "second", second)

		case "minute":
			minute := value
			second, _, _ := splitFraction(fraction * 60)

			if isNegative {
				minute, second = -minute, -second
			}

			floatDurations["minute"] = minute
			addMapValue(floatDurations, "second", second)

		case "second":
			second := math.Round(value + fraction)

			if isNegative {
				second = -second
			}

			floatDurations["second"] = second
		}
	}

	return floatDurations
}

func splitFraction(fl float64) (intPart, fractionPart float64, hasFraction bool) {
	if fl == 0 {
		return 0, 0, false
	}

	value := math.Abs(fl)
	floorValue := math.Floor(value)
	return floorValue, value - floorValue, value != floorValue
}

func addMapValue(m map[string]float64, key string, value float64) {
	if value != 0 {
		m[key] += value
	}
}

func keyExist(m map[string]float64, key string) bool {
	_, exist := m[key]
	return exist
}
