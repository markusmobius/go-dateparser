package parser

import (
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

func parseFreshnessDateTime(cfg *setting.Configuration, str string) DateData {
	// Prepare string
	str = rxBraces.ReplaceAllString(str, "")
	str, tzData := timezone.PopTzOffset(str)

	// Parse time
	t, _ := parseFreshnessTime(str)

	// Find current time
	var now time.Time
	if cfg != nil && !cfg.RelativeTimeBase.IsZero() {
		now = cfg.RelativeTimeBase
	} else {
		now = time.Now()
	}

	if !tzData.IsZero() {
		loc := time.FixedZone(tzData.Name, tzData.Offset)
		now = now.In(loc)
	}

	if cfg != nil && cfg.Timezone != nil {
		now = now.In(cfg.Timezone)
	}

	// Get relative date
	dt, period := parseFreshnessDate(cfg, str, now)
	if !dt.IsZero() {
		if !t.IsZero() {
			dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				dt.Location())

			if cfg != nil && cfg.ReturnTimeAsPeriod {
				period = Time
			}
		}

		if cfg != nil && cfg.ToTimezone != nil {
			dt = dt.In(cfg.ToTimezone)
		}

		if cfg != nil && !cfg.ReturnAsTimezoneAware {
			dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
				dt.Hour(), dt.Minute(), dt.Second(), dt.Nanosecond(),
				time.Local)
		}
	}

	// Create date data
	return DateData{Time: dt, Period: period}
}

func parseFreshnessTime(s string) (time.Time, error) {
	s = rxFreshnessPattern.ReplaceAllString(s, "")
	s = rxInAgo.ReplaceAllString(s, "")
	return parseTime(s)
}

func parseFreshnessDate(cfg *setting.Configuration, str string, now time.Time) (time.Time, DatePeriod) {
	if !allWordsAreUnits(str) {
		return time.Time{}, 0
	}

	relTimes := getRelativeTimes(str)
	if len(relTimes) == 0 {
		return time.Time{}, 0
	}

	period := Day
	if _, dayExist := relTimes["day"]; !dayExist {
		if _, monthExist := relTimes["month"]; monthExist {
			period = Month
		} else if _, yearExist := relTimes["year"]; yearExist {
			period = Year
		}
	}

	date := now
	if rxFreshnessFuture.MatchString(str) && !rxFreshnessPast.MatchString(str) {
		date = date.AddDate(relTimes["year"], relTimes["month"], relTimes["day"])
		date = date.Add(time.Duration(relTimes["hour"]) * time.Hour)
		date = date.Add(time.Duration(relTimes["minute"]) * time.Minute)
		date = date.Add(time.Duration(relTimes["second"]) * time.Second)
	} else {
		date = date.AddDate(-relTimes["year"], -relTimes["month"], -relTimes["day"])
		date = date.Add(-time.Duration(relTimes["hour"]) * time.Hour)
		date = date.Add(-time.Duration(relTimes["minute"]) * time.Minute)
		date = date.Add(-time.Duration(relTimes["second"]) * time.Second)
	}

	return date, period
}

func allWordsAreUnits(s string) bool {
	s = strings.Join(strings.Fields(s), " ")

	var wordCount int
	for _, word := range rxNonWord.Split(s, -1) {
		if word != "" && !rxFreshnessSkipWord.MatchString(word) {
			wordCount++
		}
	}

	return wordCount == 0
}

func getRelativeTimes(s string) map[string]int {
	relativeTimes := map[string]int{}

	for _, parts := range rxFreshnessPattern.FindAllStringSubmatch(s, -1) {
		period := parts[1]
		value, _ := strconv.Atoi(parts[0])
		relativeTimes[period] = value
	}

	if decades, exist := relativeTimes["decade"]; exist {
		relativeTimes["year"] += decades * 10
		delete(relativeTimes, "decade")
	}

	if weeks, exist := relativeTimes["week"]; exist {
		relativeTimes["day"] += weeks * 7
		delete(relativeTimes, "week")
	}

	return relativeTimes
}
