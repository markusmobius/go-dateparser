package hijri

import (
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/hablullah/go-hijri"
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
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
	str, tz := timezone.PopTzOffset(str)

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
	dt, err := parser.Parse(tz)
	if err != nil {
		return date.Date{}, err
	}

	// Apply the popped timezone
	if !dt.IsZero() && !tz.IsZero() {
		dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(),
			dt.Time.Hour(), dt.Time.Minute(), dt.Time.Second(), dt.Time.Nanosecond(),
			time.FixedZone(tz.Name, tz.Offset))
	}

	return dt, nil
}

func getDateTimeParams(p *absolute.Parser) map[string]int {
	// Get current time in Hijri
	hd, _ := hijri.CreateUmmAlQuraDate(p.Now)

	// Get component values
	day, dayExist := p.ComponentValues["day"]
	if !dayExist || day == 0 {
		day = int(hd.Day)
	}

	month, monthExist := p.ComponentValues["month"]
	if !monthExist || month == 0 {
		month = int(hd.Month)
	}

	year, yearExist := p.ComponentValues["year"]
	if !yearExist || year == 0 {
		year = int(hd.Year)
	}

	return map[string]int{
		"year":  year,
		"month": month,
		"day":   day,
	}
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

	// Fix leap year
	if D == 30 && M == 12 && !isHijriLeapYear(Y) {
		Y = getCorrectHijriLeapYear(p.Config, Y)
	}

	// Fix max day
	lastDayOfMonth := getLastDayOfHijriMonth(Y, M)
	if D > lastDayOfMonth {
		D = lastDayOfMonth
	}

	// Convert Hijri to Gregorian
	gd := hijri.UmmAlQuraDate{Year: int64(Y), Month: int64(M), Day: int64(D)}.ToGregorian()
	return time.Date(gd.Year(), gd.Month(), gd.Day(), H, m, s, ns, loc), nil
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func isHijriLeapYear(year int) bool {
	hd := hijri.UmmAlQuraDate{Year: int64(year), Month: 12, Day: 30}
	gd := hd.ToGregorian()
	rhd, _ := hijri.CreateUmmAlQuraDate(gd)
	return rhd.Day == 30
}

func getHijriLeapYear(year int, toFuture bool) int {
	step := 1
	if !toFuture {
		step = -1
	}

	for {
		year += step
		if isHijriLeapYear(year) {
			return year
		}
	}
}

func getCorrectHijriLeapYear(cfg *setting.Configuration, currentYear int) int {
	var dateSource setting.PreferredDateSource
	if cfg != nil {
		dateSource = cfg.PreferredDateSource
	}

	switch dateSource {
	case setting.Future:
		return getHijriLeapYear(currentYear, true)
	case setting.Past:
		return getHijriLeapYear(currentYear, false)
	default:
		nextLeapYear := getHijriLeapYear(currentYear, true)
		prevLeapYear := getHijriLeapYear(currentYear, false)
		nextLeapYearIsCloser := nextLeapYear-currentYear < currentYear-prevLeapYear
		if nextLeapYearIsCloser {
			return nextLeapYear
		} else {
			return prevLeapYear
		}
	}
}

func getLastDayOfHijriMonth(year, month int) int {
	switch {
	case month == 12 && isHijriLeapYear(year):
		return 30
	case month%2 == 0:
		return 29
	default:
		return 30
	}
}

func handleTwoDigitYear(year int) int {
	if year >= 90 {
		return year + 1300
	} else {
		return year + 1400
	}
}
