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
	"github.com/markusmobius/go-dateutil/v2/relativedelta"
)

var (
	rxNonWord          = regexp.MustCompile(`\W`)
	rxIn               = regexp.MustCompile(`(?i)\bin\b`)
	rxAgo              = regexp.MustCompile(`(?i)\bago\b`)
	rxInAgo            = regexp.MustCompile(`(?i)\b(?:ago|in)\b`)
	rxRelativePattern  = regexp.MustCompile(`(?i)([+-]?\s*\d+[.,]?\d*)\s*(` + relativeUnits + `)\b`)
	rxRelativeSkipWord = regexp.MustCompile(`(?i)^(?:` + relativeUnits + `|ago|in|\d+|:|[ap]m)`)
	relativeUnits      = `decade|year|month|week|day|hour|minute|second`
)

// Parse parses date string like "1 year, 2 months ago" and "3 hours, 50 minutes ago".
func Parse(cfg *setting.Configuration, str string) date.Date {
	// Prepare string
	str = strutil.StripBraces(str)
	str, tzData := timezone.PopTzOffset(str)

	// Parse time
	t, tPeriod, _ := parseTime(str)

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
		period = min(period, tPeriod)
	}

	if cfg != nil && !cfg.ReturnTimeAsPeriod && period.IsTime() {
		period = date.Day
	}

	// Create date data
	return date.Date{Time: dt, Period: period}
}

func parseTime(s string) (time.Time, date.Period, error) {
	s = rxRelativePattern.ReplaceAllString(s, "")
	s = rxInAgo.ReplaceAllString(s, "")
	return common.ParseTime(s)
}

func parseDate(cfg *setting.Configuration, str string, now time.Time) (time.Time, date.Period) {
	if !allWordsAreUnits(str) {
		return time.Time{}, 0
	}

	// Retrieve relative durations
	goingForward := rxIn.MatchString(str) || (cfg.PreferredDateSource == setting.Future && !rxAgo.MatchString(str))
	relDurations := getRelativeDurations(str, goingForward)
	if len(relDurations) == 0 {
		return time.Time{}, 0
	}
	for unit, maximum := range map[string]float64{
		"year": 10000, "month": 12 * 10000, "day": 366 * 10000,
		"hour": 24 * 366 * 10000, "minute": 60 * 24 * 366 * 10000,
		"second": 60 * 60 * 24 * 366 * 10000,
	} {
		value := relDurations[unit]
		if math.IsInf(value, 0) || math.IsNaN(value) || math.Abs(value) > maximum {
			return time.Time{}, 0
		}
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
	if cfg.ReturnTimeAsPeriod {
		fractionalHours := math.Mod(relDurations["day"], 1) * 24
		if math.Trunc(fractionalHours) != 0 {
			period = min(period, date.Hour)
		}
		fractionalMinutes := math.Mod(relDurations["hour"]+fractionalHours, 1) * 60
		if math.Trunc(fractionalMinutes) != 0 {
			period = min(period, date.Minute)
		}
		fractionalSeconds := math.Mod(relDurations["minute"]+fractionalMinutes, 1) * 60
		if math.Trunc(fractionalSeconds) != 0 {
			period = min(period, date.Second)
		}
	}

	date, err := (relativedelta.Delta{
		Years: relDurations["year"], Months: relDurations["month"], Days: relDurations["day"],
		Hours: relDurations["hour"], Minutes: relDurations["minute"], Seconds: relDurations["second"],
	}).Apply(now)
	if err != nil {
		return time.Time{}, 0
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

func getRelativeDurations(s string, goingForward bool) map[string]float64 {
	// Extract durations using regex
	floatDurations := map[string]float64{}
	for _, parts := range rxRelativePattern.FindAllStringSubmatch(s, -1) {
		period := parts[2]
		strValue := strings.ReplaceAll(strings.Join(strings.Fields(parts[1]), ""), ",", ".")
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil || math.IsInf(value, 0) || math.IsNaN(value) {
			return nil
		}
		if !goingForward && strValue[0] != '+' && strValue[0] != '-' {
			value = -value
		}
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

	return floatDurations
}

func keyExist(m map[string]float64, key string) bool {
	_, exist := m[key]
	return exist
}
