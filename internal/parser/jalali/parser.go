package jalali

import (
	"strconv"
	"time"
	"unicode"

	"github.com/jalaali/go-jalaali"
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

func Parse(cfg *setting.Configuration, str string) (date.Date, error) {
	// Normalize the string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)

	// Translate the foreign texts
	str = replaceMonths(str)
	str = replaceWeekdays(str)
	str = replaceDays(str)
	str = replaceTime(str)

	// Sanitize the string
	str = removeWeekdayTranslations(str)
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
	// Get current time in Jalali
	jd := jalaali.From(p.Now)

	// Get component values
	day, dayExist := p.ComponentValues["day"]
	if !dayExist || day == 0 {
		day = jd.Day()
	}

	month, monthExist := p.ComponentValues["month"]
	if !monthExist || month == 0 {
		month = int(jd.Month())
	}

	year, yearExist := p.ComponentValues["year"]
	if !yearExist || year == 0 {
		year = jd.Year()
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

	if component == "month" {
		monthValue, monthValueExist := monthNameValues[token]
		if directive == "January" && monthValueExist {
			return monthValue, true
		}

		if tokenLength <= 2 && tokenIsDigit {
			month, _ := strconv.Atoi(token)
			if month >= 1 && month <= 12 {
				return month, true
			}
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
	isLeapYear, err := jalaali.IsLeapYear(Y)
	if err != nil {
		return time.Time{}, err
	}

	if D == 30 && M == 12 && !isLeapYear {
		Y = getCorrectLeapYear(p.Config, Y)
	}

	// Fix max day
	lastDayOfMonth, err := jalaali.MonthLength(Y, M)
	if err == nil && D > lastDayOfMonth {
		D = lastDayOfMonth
	}

	// Convert Jalali to Gregorian
	gY, gM, gD, err := jalaali.ToGregorian(Y, jalaali.Month(M), D)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(gY, gM, gD, H, m, s, ns, loc), nil
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func getLeapYear(year int, toFuture bool) int {
	step := 1
	if !toFuture {
		step = -1
	}

	originalYear := year

	for {
		year += step
		isLeap, err := jalaali.IsLeapYear(year)
		if isLeap {
			return year
		} else if err != nil {
			return originalYear
		}
	}
}

func getCorrectLeapYear(cfg *setting.Configuration, currentYear int) int {
	var dateSource setting.PreferredDateSource
	if cfg != nil {
		dateSource = cfg.PreferredDateSource
	}

	switch dateSource {
	case setting.Future:
		return getLeapYear(currentYear, true)
	case setting.Past:
		return getLeapYear(currentYear, false)
	default:
		nextLeapYear := getLeapYear(currentYear, true)
		prevLeapYear := getLeapYear(currentYear, false)
		nextLeapYearIsCloser := nextLeapYear-currentYear < currentYear-prevLeapYear
		if nextLeapYearIsCloser {
			return nextLeapYear
		} else {
			return prevLeapYear
		}
	}
}

func handleTwoDigitYear(year int) int {
	if year > 60 {
		return year + 1300
	} else {
		return year + 1400
	}
}
