package absolute

import (
	"fmt"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/dateutil"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

func Parse(cfg *setting.Configuration, str string, tz timezone.OffsetData) (date.Date, error) {
	parser := &Parser{
		Config:              cfg,
		FnGetDateTimeParams: getDateTimeParams,
		FnGetDatePartValue:  getDatePartValue,
		FnCreateDateTime:    createDateTime,
	}

	if err := parser.Init(str); err != nil {
		return date.Date{}, err
	}

	return parser.Parse(tz)
}

func getDateTimeParams(p *Parser) (map[string]int, error) {
	// Get component values
	day, dayExist := p.ComponentValues["day"]
	if !dayExist || day == 0 {
		day = p.Now.Day()
	}

	month, monthExist := p.ComponentValues["month"]
	if !monthExist || month == 0 {
		month = int(p.Now.Month())
	}

	year, yearExist := p.ComponentValues["year"]
	if !yearExist || year == 0 {
		year = p.Now.Year()
	}

	return map[string]int{
		"year":  year,
		"month": month,
		"day":   day,
	}, nil
}

func createDateTime(p *Parser, pms map[string]int, loc *time.Location) (time.Time, error) {
	year, month, day := pms["year"], pms["month"], pms["day"]
	_, explicitYear := p.ComponentValues["year"]
	_, explicitDay := p.ComponentValues["day"]
	if year < 1 || year > 9999 || month < 1 || month > 12 || day < 1 {
		return time.Time{}, fmt.Errorf("invalid calendar date: %d-%d-%d", year, month, day)
	}
	if day == 29 && month == 2 && !dateutil.IsLeapYear(year) && !explicitYear {
		year = p.getCorrectLeapYear(year)
	}
	lastDay := dateutil.GetLastDayOfMonth(year, month)
	if day > lastDay {
		if explicitDay {
			return time.Time{}, fmt.Errorf("invalid calendar day: %d-%d-%d", year, month, day)
		}
		day = lastDay
	}
	return time.Date(year, time.Month(month), day, pms["hour"], pms["minute"], pms["second"], pms["nanosecond"], loc), nil
}

func getDatePartValue(p *Parser, component, token, directive string) (int, bool) {
	token = strings.ToLower(token)
	if t, _ := time.Parse(directive, token); !t.IsZero() {
		switch component {
		case "day":
			return t.Day(), true
		case "weekday":
			return int(t.Weekday()), true
		case "month":
			return int(t.Month()), true
		case "year":
			return t.Year(), true
		}
	}

	return 0, false
}
