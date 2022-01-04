package parser

import (
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

func parseFreshnessPattern(cfg *setting.Configuration, str string) DateData {
	// Prepare string
	str = rxBraces.ReplaceAllString(str, "")
	str, tzData := timezone.PopTzOffset(str)

	// Parse time
	t, _ := parseFreshnessTime(str)

	// Find current time
	now := time.Now()
	if cfg != nil && !cfg.CurrentTime.IsZero() {
		now = cfg.CurrentTime
	}

	if !tzData.IsZero() {
		loc := time.FixedZone(tzData.Name, tzData.Offset)
		now = now.In(loc)
	}

	// Get relative date
	dt, period := parseFreshnessDate(cfg, str, now)

	if !dt.IsZero() && !t.IsZero() {
		dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			dt.Location())
		if cfg != nil && cfg.ReturnTimeAsPeriod {
			period = Time
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
		if _, weekExist := relTimes["week"]; weekExist {
			period = Week
		} else if _, monthExist := relTimes["month"]; monthExist {
			period = Month
		} else if _, yearExist := relTimes["year"]; yearExist {
			period = Year
		}
	}

	date := now
	days := relTimes["day"] + relTimes["week"]*7
	if (rxIn.MatchString(str) || cfg.PreferFutureTime) && !rxAgo.MatchString(str) {
		date = date.AddDate(relTimes["year"], relTimes["month"], days)
		date = date.Add(time.Duration(relTimes["hour"]) * time.Hour)
		date = date.Add(time.Duration(relTimes["minute"]) * time.Minute)
		date = date.Add(time.Duration(relTimes["second"]) * time.Second)
	} else {
		date = date.AddDate(-relTimes["year"], -relTimes["month"], -days)
		date = date.Add(-time.Duration(relTimes["hour"]) * time.Hour)
		date = date.Add(-time.Duration(relTimes["minute"]) * time.Minute)
		date = date.Add(-time.Duration(relTimes["second"]) * time.Second)
	}

	return date, period
}

func allWordsAreUnits(s string) bool {
	s = strings.Join(strings.Fields(s), " ")

	var words []string
	var wordCount int
	for _, word := range rxNonWord.Split(s, -1) {
		if word == "" {
			continue
		}

		if !rxFreshnessSkipWord.MatchString(word) {
			wordCount++
			words = append(words, word)
		}
	}

	return wordCount == 0
}

func getRelativeTimes(s string) map[string]int {
	relativeTimes := map[string]int{}

	for _, parts := range rxFreshnessPattern.FindAllStringSubmatch(s, -1) {
		period := parts[2]
		value, _ := strconv.Atoi(parts[1])
		relativeTimes[period] = value
	}

	if decades, exist := relativeTimes["decade"]; exist {
		relativeTimes["year"] += decades * 10
		delete(relativeTimes, "decade")
	}

	return relativeTimes
}
