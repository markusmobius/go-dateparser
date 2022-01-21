package relative

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var (
	rxNonWord          = regexp.MustCompile(`\W`)
	rxIn               = regexp.MustCompile(`(?i)\bin\b`)
	rxAgo              = regexp.MustCompile(`(?i)\bago\b`)
	rxInAgo            = regexp.MustCompile(`(?i)\b(?:ago|in)\b`)
	rxRelativePattern  = regexp.MustCompile(`(?i)(\d+)\s*(` + relativeUnits + `)\b`)
	rxRelativeSkipWord = regexp.MustCompile(`(?i)^(?:` + relativeUnits + `|ago|in|\d+|:|[ap]m)`)
	relativeUnits      = `decade|year|month|week|day|hour|minute|second`
)

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

	if !tzData.IsZero() {
		loc := time.FixedZone(tzData.Name, tzData.Offset)
		now = now.In(loc)
	}

	// Get relative date
	dt, period := parseDate(cfg, str, now)

	if !dt.IsZero() && !t.IsZero() {
		dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			dt.Location())
		if cfg != nil && cfg.ReturnTimeAsPeriod {
			period = date.Time
		}
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

	relTimes := getRelativeTimes(str)
	if len(relTimes) == 0 {
		return time.Time{}, 0
	}

	period := date.Day
	if _, dayExist := relTimes["day"]; !dayExist {
		if _, weekExist := relTimes["week"]; weekExist {
			period = date.Week
		} else if _, monthExist := relTimes["month"]; monthExist {
			period = date.Month
		} else if _, yearExist := relTimes["year"]; yearExist {
			period = date.Year
		}
	}

	date := now
	days := relTimes["day"] + relTimes["week"]*7
	if (rxIn.MatchString(str) || cfg.PreferredDateSource == setting.Future) && !rxAgo.MatchString(str) {
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

func getRelativeTimes(s string) map[string]int {
	relativeTimes := map[string]int{}

	for _, parts := range rxRelativePattern.FindAllStringSubmatch(s, -1) {
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
