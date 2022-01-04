package parser

import (
	"fmt"
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
	fmt.Println(tzData, tzData.IsZero())

	// Parse time
	t, _ := parseFreshnessTime(str)
	fmt.Println("TIME:", t)

	// Find current time
	now := time.Now()
	if cfg != nil && !cfg.CurrentTime.IsZero() {
		now = cfg.CurrentTime
	}

	fmt.Println("NOW 0:", now)
	if !tzData.IsZero() {
		loc := time.FixedZone(tzData.Name, tzData.Offset)
		now = now.In(loc)
		fmt.Println("NOW 1:", now)
	}

	// Get relative date
	// fmt.Println("NOW:", now)
	dt, period := parseFreshnessDate(cfg, str, now)
	fmt.Println("NOW 2:", now)
	fmt.Println("DATE 1:", dt)

	if !dt.IsZero() && !t.IsZero() {
		dt = time.Date(dt.Year(), dt.Month(), dt.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			dt.Location())
		fmt.Println("DATE 2:", dt)
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
	wordsAreUnit := allWordsAreUnits(str)
	// fmt.Println("ALL WORDS ARE UNIT", wordsAreUnit)
	if !wordsAreUnit {
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

	// fmt.Println("SKIP:", rxFreshnessSkipWord)
	// fmt.Println("SPLIT:", strutil.Jsonify(rxNonWord.Split(s, -1)))

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

	// fmt.Println("WORDS:", strutil.Jsonify(&words))
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
