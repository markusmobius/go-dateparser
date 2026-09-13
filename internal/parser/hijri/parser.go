package hijri

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
	"github.com/markusmobius/go-dateparser/internal/parser/calendars"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var timeConventions = map[string]string{
	"am": "صباحا",
	"pm": "مساء",
}

func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	// Normalize the string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)

	// Translate the foreign texts
	for latin, arabic := range timeConventions {
		str = strings.ReplaceAll(str, arabic, latin)
	}

	// Sanitize the string
	str = strutil.SanitizeDate(str)
	str = strutil.StripBraces(str)

	// Create parser
	parser := &absolute.Parser{
		Config:              cfg,
		FnGetDateTimeParams: getDateTimeParams,
		FnGetDatePartValue:  getDatePartValue,
		FnCreateDateTime:    createDateTime,
	}

	if err := parser.Init(str); err != nil {
		return date.Date{}, err
	}

	// Parse the string
	return parser.Parse(timezone.OffsetData{})
}

func getDateTimeParams(p *absolute.Parser) (map[string]int, error) {
	// Get current time in Hijri
	current, err := calendars.HijriFromGregorian(p.Now)
	if err != nil {
		return nil, err
	}

	// Get component values
	day, dayExist := p.ComponentValues["day"]
	if !dayExist || day == 0 {
		day = current[2]
	}

	month, monthExist := p.ComponentValues["month"]
	if !monthExist || month == 0 {
		month = current[1]
	}

	year, yearExist := p.ComponentValues["year"]
	if !yearExist || year == 0 {
		year = current[0]
	}

	return map[string]int{
		"year":  year,
		"month": month,
		"day":   day,
	}, nil
}

func getDatePartValue(p *absolute.Parser, component, token, directive string) (int, bool) {
	tokenLength := len(token)
	tokenIsDigit := isDigit(token)

	if component == "year" && tokenLength == 4 && tokenIsDigit {
		year, _ := strconv.Atoi(token)
		if year >= 0 {
			return year, true
		}
	}

	if component == "year" && tokenLength == 2 && tokenIsDigit {
		year, _ := strconv.Atoi(token)
		year = handleTwoDigitYear(year)
		if year >= 0 {
			return year, true
		}
	}

	if component == "month" && tokenLength <= 2 && tokenIsDigit {
		month, _ := strconv.Atoi(token)
		if month >= 1 && month <= 12 {
			return month, true
		}
	}

	if component == "day" && tokenLength <= 2 && tokenIsDigit {
		day, _ := strconv.Atoi(token)
		if day >= 1 && day <= 30 {
			return day, true
		}
	}

	return 0, false
}

func createDateTime(p *absolute.Parser, pms map[string]int, loc *time.Location) (time.Time, error) {
	Y, M, D := pms["year"], pms["month"], pms["day"]
	H, m, s, ns := pms["hour"], pms["minute"], pms["second"], pms["nanosecond"]

	lastDayOfMonth, err := calendars.HijriMonthLength(Y, M)
	if err != nil {
		return time.Time{}, err
	}
	_, explicitDay := p.ComponentTokens["day"]
	_, explicitWeekday := p.ComponentTokens["weekday"]
	if !explicitDay && !explicitWeekday && (D < 1 || D > lastDayOfMonth) {
		D = lastDayOfMonth
	}

	// Convert Hijri to Gregorian
	converted, err := calendars.HijriToGregorian(Y, M, D)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(converted[0], time.Month(converted[1]), converted[2], H, m, s, ns, loc), nil
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func handleTwoDigitYear(year int) int {
	if year >= 90 {
		return year + 1300
	} else {
		return year + 1400
	}
}
