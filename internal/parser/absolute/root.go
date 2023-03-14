package absolute

import (
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

func getDateTimeParams(p *Parser) map[string]int {
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
	}
}

func createDateTime(p *Parser, pms map[string]int, loc *time.Location) (time.Time, error) {
	Y, M, D := pms["year"], pms["month"], pms["day"]
	H, m, s, ns := pms["hour"], pms["minute"], pms["second"], pms["nanosecond"]

	// Fix leap year
	if D == 29 && M == 2 && !dateutil.IsLeapYear(Y) {
		Y = p.getCorrectLeapYear(Y)
	}

	// Fix max day
	lastDayOfMonth := dateutil.GetLastDayOfMonth(Y, M)
	if D > lastDayOfMonth {
		D = lastDayOfMonth
	}

	return time.Date(Y, time.Month(M), D, H, m, s, ns, loc), nil
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
