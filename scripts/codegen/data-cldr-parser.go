package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseAllCldrData(languageLocalesMap map[string][]string) (map[string]LocaleData, error) {
	result := map[string]LocaleData{}

	for language, locales := range languageLocalesMap {
		locales = append([]string{language}, locales...)
		for _, locale := range locales {
			localeData, err := parseCldrData(locale)
			if os.IsNotExist(err) {
				continue
			} else if err != nil {
				return nil, err
			}

			log.Info().Msgf("parsed cldr for %s", locale)
			result[locale] = *localeData
		}
	}

	return result, nil
}

func parseCldrData(locale string) (*LocaleData, error) {
	// Parse data
	cldrGregorian, err := parseCldrGregorianData(locale)
	if err != nil {
		return nil, err
	}

	cldrDateFields, err := parseCldrDateFieldsData(locale)
	if err != nil {
		return nil, err
	}

	// Fetch main data
	gregorianData := cldrGregorian.Main[locale].Dates.Calendars.Gregorian
	dateFieldsData := cldrDateFields.Main[locale].Dates.Fields

	// Generate locale data
	setLocaleMonthData := func(month int) []string {
		strMonth := strconv.Itoa(month)
		return cleanList(nil, localeMonthFilter,
			gregorianData.Months.Format.Wide[strMonth],
			gregorianData.Months.Format.Abbreviated[strMonth],
			gregorianData.Months.StandAlone.Wide[strMonth],
			gregorianData.Months.StandAlone.Abbreviated[strMonth])
	}

	setLocaleWeekDays := func(weekDays string) []string {
		return cleanList(nil, nil,
			gregorianData.Days.Format.Wide[weekDays],
			gregorianData.Days.Format.Abbreviated[weekDays],
			gregorianData.Days.StandAlone.Wide[weekDays],
			gregorianData.Days.StandAlone.Abbreviated[weekDays])
	}

	setLocaleDayPeriods := func(period string) []string {
		periodList := []string{
			gregorianData.DayPeriods.Format.Wide[period],
			gregorianData.DayPeriods.Format.Abbreviated[period],
			gregorianData.DayPeriods.StandAlone.Wide[period],
			gregorianData.DayPeriods.StandAlone.Abbreviated[period],
		}

		for i, p := range periodList {
			periodList[i] = rxAmPmPattern.ReplaceAllString(p, period)
		}

		return cleanList(nil, nil, periodList...)
	}

	setLocaleDateFields := func(fieldName string, attr string) []string {
		finalList := make([]string, 3)
		keys := []string{
			fieldName,
			fieldName + "-short",
			fieldName + "-narrow",
		}

		for i, key := range keys {
			switch attr {
			case "display":
				finalList[i] = dateFieldsData[key].DisplayName
			case "relative-past":
				finalList[i] = dateFieldsData[key].RelativeTypeMin1
			case "relative-now":
				finalList[i] = dateFieldsData[key].RelativeType0
			case "relative-future":
				finalList[i] = dateFieldsData[key].RelativeType1
			}
		}

		return cleanList(nil, nil, finalList...)
	}

	setLocaleRegexes := func(fieldName string, relativeType string) []string {
		var finalList []string
		keys := []string{
			fieldName,
			fieldName + "-short",
			fieldName + "-narrow",
		}

		if relativeType == "future" {
			for _, key := range keys {
				finalList = append(finalList,
					dateFieldsData[key].RelativeTimeTypeFuture.CountOne,
					dateFieldsData[key].RelativeTimeTypeFuture.CountOther)
			}
		} else {
			for _, key := range keys {
				finalList = append(finalList,
					dateFieldsData[key].RelativeTimeTypePast.CountOne,
					dateFieldsData[key].RelativeTimeTypePast.CountOther)
			}
		}

		return cleanList(nil, regexFilter, finalList...)
	}

	data := LocaleData{
		Name:      locale,
		DateOrder: rxDateOrderPattern.ReplaceAllString(gregorianData.DateFormats.Short, "$1$2$3"),
		January:   setLocaleMonthData(1),
		February:  setLocaleMonthData(2),
		March:     setLocaleMonthData(3),
		April:     setLocaleMonthData(4),
		May:       setLocaleMonthData(5),
		June:      setLocaleMonthData(6),
		July:      setLocaleMonthData(7),
		August:    setLocaleMonthData(8),
		September: setLocaleMonthData(9),
		October:   setLocaleMonthData(10),
		November:  setLocaleMonthData(11),
		December:  setLocaleMonthData(12),
		Monday:    setLocaleWeekDays("mon"),
		Tuesday:   setLocaleWeekDays("tue"),
		Wednesday: setLocaleWeekDays("wed"),
		Thursday:  setLocaleWeekDays("thu"),
		Friday:    setLocaleWeekDays("fri"),
		Saturday:  setLocaleWeekDays("sat"),
		Sunday:    setLocaleWeekDays("sun"),
		AM:        setLocaleDayPeriods("am"),
		PM:        setLocaleDayPeriods("pm"),
		Year:      setLocaleDateFields("year", "display"),
		Month:     setLocaleDateFields("month", "display"),
		Week:      setLocaleDateFields("week", "display"),
		Day:       setLocaleDateFields("day", "display"),
		Hour:      setLocaleDateFields("hour", "display"),
		Minute:    setLocaleDateFields("minute", "display"),
		Second:    setLocaleDateFields("second", "display"),
		RelativeType: cleanEmptyStringListMap(map[string][]string{
			"1 year ago":   setLocaleDateFields("year", "relative-past"),
			"0 year ago":   setLocaleDateFields("year", "relative-now"),
			"in 1 year":    setLocaleDateFields("year", "relative-future"),
			"1 month ago":  setLocaleDateFields("month", "relative-past"),
			"0 month ago":  setLocaleDateFields("month", "relative-now"),
			"in 1 month":   setLocaleDateFields("month", "relative-future"),
			"1 week ago":   setLocaleDateFields("week", "relative-past"),
			"0 week ago":   setLocaleDateFields("week", "relative-now"),
			"in 1 week":    setLocaleDateFields("week", "relative-future"),
			"1 day ago":    setLocaleDateFields("day", "relative-past"),
			"0 day ago":    setLocaleDateFields("day", "relative-now"),
			"in 1 day":     setLocaleDateFields("day", "relative-future"),
			"0 hour ago":   setLocaleDateFields("hour", "relative-now"),
			"0 minute ago": setLocaleDateFields("minute", "relative-now"),
			"0 second ago": setLocaleDateFields("second", "relative-now"),
		}),
		RelativeTypeRegex: cleanEmptyStringListMap(map[string][]string{
			"in \\1 year":    setLocaleRegexes("year", "future"),
			"\\1 year ago":   setLocaleRegexes("year", "past"),
			"in \\1 month":   setLocaleRegexes("month", "future"),
			"\\1 month ago":  setLocaleRegexes("month", "past"),
			"in \\1 week":    setLocaleRegexes("week", "future"),
			"\\1 week ago":   setLocaleRegexes("week", "past"),
			"in \\1 day":     setLocaleRegexes("day", "future"),
			"\\1 day ago":    setLocaleRegexes("day", "past"),
			"in \\1 hour":    setLocaleRegexes("hour", "future"),
			"\\1 hour ago":   setLocaleRegexes("hour", "past"),
			"in \\1 minute":  setLocaleRegexes("minute", "future"),
			"\\1 minute ago": setLocaleRegexes("minute", "past"),
			"in \\1 second":  setLocaleRegexes("second", "future"),
			"\\1 second ago": setLocaleRegexes("second", "past"),
		}),
	}

	return &data, nil
}

func parseCldrGregorianData(locale string) (*CldrGregorianData, error) {
	// Parse JSON
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")
	fPath := filepath.Join(cldrDatesFullDir, locale, "ca-gregorian.json")

	var data CldrGregorianData
	err := parseJsonFile(&data, fPath)
	if err != nil {
		return nil, err
	}

	// Check for dateFormats.short value
	mainData := data.Main[locale]
	shortItfValue := mainData.Dates.Calendars.Gregorian.DateFormats.ShortItf
	switch v := shortItfValue.(type) {
	case string:
		mainData.Dates.Calendars.Gregorian.DateFormats.Short = strings.ToUpper(v)
	case map[string]interface{}:
		if value, ok := v["_value"].(string); ok {
			mainData.Dates.Calendars.Gregorian.DateFormats.Short = strings.ToUpper(value)
		}
	}
	data.Main[locale] = mainData

	return &data, nil
}

func parseCldrDateFieldsData(locale string) (*CldrDateFieldsData, error) {
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")
	fPath := filepath.Join(cldrDatesFullDir, locale, "dateFields.json")

	var data CldrDateFieldsData
	err := parseJsonFile(&data, fPath)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func localeMonthFilter(str string) bool {
	return !rxDefaultMonthPattern.MatchString(str)
}

func regexFilter(str string) bool {
	return strings.Contains(str, "{0}") &&
		!rxNotRelativePattern.MatchString(str) &&
		!rxParenthesisPattern.MatchString(str)
}
