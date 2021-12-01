package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func parseAllCldrData(languageLocalesMap map[string][]string) (map[string]LocaleData, error) {
	result := map[string]LocaleData{}

	for language, locales := range languageLocalesMap {
		log.Info().Msgf("parsing language %s", language)
		localeData, err := parseCldrData(language, locales...)
		if err != nil {
			return nil, err
		}

		result[language] = *localeData
	}

	return result, nil
}

func parseCldrData(locale string, sublocales ...string) (*LocaleData, error) {
	// Parse data
	cldrGregorian, err := parseCldrGregorianData(locale)
	if err != nil {
		return nil, err
	}

	cldrDateFields, err := parseCldrDateFieldsData(locale)
	if err != nil {
		return nil, err
	}

	if cldrGregorian == nil || cldrDateFields == nil {
		return &LocaleData{}, nil
	}

	// Fetch main data
	gregorianData := cldrGregorian.Main[locale].Dates.Calendars.Gregorian
	dateFieldsData := cldrDateFields.Main[locale].Dates.Fields

	// Generate locale data
	setLocaleMonthData := func(month int) []string {
		strMonth := strconv.Itoa(month)
		return clearList(nil, localeMonthFilter,
			gregorianData.Months.Format.Wide[strMonth],
			gregorianData.Months.Format.Abbreviated[strMonth],
			gregorianData.Months.StandAlone.Wide[strMonth],
			gregorianData.Months.StandAlone.Abbreviated[strMonth])
	}

	setLocaleWeekDays := func(weekDays string) []string {
		return clearList(nil, nil,
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

		return clearList(nil, nil, periodList...)
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

		return clearList(nil, nil, finalList...)
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

		return clearList(nil, regexFilter, finalList...)
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
		RelativeType: clearEmptyStringListMap(map[string][]string{
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
		RelativeTypeRegex: clearEmptyStringListMap(map[string][]string{
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

	// Process sublocales
	if len(sublocales) > 0 {
		data.LocaleSpecific = map[string]LocaleData{}

		for _, sublocale := range sublocales {
			subData, err := parseCldrData(sublocale)
			if err != nil {
				return nil, err
			}

			cleanedData := removeLocaleDataDuplicate(data, *subData)
			data.LocaleSpecific[sublocale] = cleanedData
		}
	}

	return &data, nil
}

func parseCldrGregorianData(locale string) (*CldrGregorianData, error) {
	// Prepare file path
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")
	fPath := filepath.Join(cldrDatesFullDir, locale, "ca-gregorian.json")

	// Open file
	f, err := os.Open(fPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn().Msgf("gregorian data for locale %s not exist, skipped", locale)
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	// Parse JSON
	var data CldrGregorianData
	err = json.NewDecoder(f).Decode(&data)
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
	// Prepare file path
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")
	fPath := filepath.Join(cldrDatesFullDir, locale, "dateFields.json")

	// Open file
	f, err := os.Open(fPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn().Msgf("date fields data for locale %s not exist, skipped", locale)
			return nil, nil
		}
		return nil, err
	}
	defer f.Close()

	// Parse JSON
	var data CldrDateFieldsData
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func cleanString(str string) string {
	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}

func clearList(initialItems []string, fnFilter func(string) bool, items ...string) []string {
	// Prepare tracker
	tracker := map[string]struct{}{}
	if initialItems != nil && len(initialItems) > 0 {
		for _, item := range initialItems {
			tracker[item] = struct{}{}
		}
	}

	// Create final list
	var finalList []string
	for _, item := range items {
		item = cleanString(item)

		// Make sure item not empty
		if item == "" {
			continue
		}

		// Make sure item tracked yet
		if _, exist := tracker[item]; exist {
			continue
		}

		// Make sure item passed the filter
		if fnFilter != nil && !fnFilter(item) {
			continue
		}

		// Save the item
		tracker[item] = struct{}{}
		finalList = append(finalList, item)
	}

	// Sort final list
	sort.Strings(finalList)

	if len(finalList) == 0 {
		return nil
	} else {
		return finalList
	}
}

func clearStringListMap(parentMap, childMap map[string][]string) map[string][]string {
	for key, list := range childMap {
		parentList := parentMap[key]
		cleanedList := clearList(parentList, nil, list...)

		if len(cleanedList) == 0 {
			delete(childMap, key)
		} else {
			childMap[key] = cleanedList
		}
	}

	if len(childMap) == 0 {
		return nil
	} else {
		return childMap
	}
}

func removeLocaleDataDuplicate(parent LocaleData, child LocaleData) LocaleData {
	var cleanedDateOrder string
	if child.DateOrder != parent.DateOrder {
		cleanedDateOrder = child.DateOrder
	}

	return LocaleData{
		Name:      child.Name,
		DateOrder: cleanedDateOrder,
		January:   clearList(parent.January, nil, child.January...),
		February:  clearList(parent.February, nil, child.February...),
		March:     clearList(parent.March, nil, child.March...),
		April:     clearList(parent.April, nil, child.April...),
		May:       clearList(parent.May, nil, child.May...),
		June:      clearList(parent.June, nil, child.June...),
		July:      clearList(parent.July, nil, child.July...),
		August:    clearList(parent.August, nil, child.August...),
		September: clearList(parent.September, nil, child.September...),
		October:   clearList(parent.October, nil, child.October...),
		November:  clearList(parent.November, nil, child.November...),
		December:  clearList(parent.December, nil, child.December...),
		Monday:    clearList(parent.Monday, nil, child.Monday...),
		Tuesday:   clearList(parent.Tuesday, nil, child.Tuesday...),
		Wednesday: clearList(parent.Wednesday, nil, child.Wednesday...),
		Thursday:  clearList(parent.Thursday, nil, child.Thursday...),
		Friday:    clearList(parent.Friday, nil, child.Friday...),
		Saturday:  clearList(parent.Saturday, nil, child.Saturday...),
		Sunday:    clearList(parent.Sunday, nil, child.Sunday...),
		AM:        clearList(parent.AM, nil, child.AM...),
		PM:        clearList(parent.PM, nil, child.PM...),
		Year:      clearList(parent.Year, nil, child.Year...),
		Month:     clearList(parent.Month, nil, child.Month...),
		Week:      clearList(parent.Week, nil, child.Week...),
		Day:       clearList(parent.Day, nil, child.Day...),
		Hour:      clearList(parent.Hour, nil, child.Hour...),
		Minute:    clearList(parent.Minute, nil, child.Minute...),
		Second:    clearList(parent.Second, nil, child.Second...),

		RelativeType:      clearStringListMap(parent.RelativeType, child.RelativeType),
		RelativeTypeRegex: clearStringListMap(parent.RelativeTypeRegex, child.RelativeTypeRegex),
	}
}

func clearEmptyStringListMap(m map[string][]string) map[string][]string {
	for key, list := range m {
		if len(list) == 0 {
			delete(m, key)
		}
	}

	if len(m) == 0 {
		return nil
	}

	return m
}

func localeMonthFilter(str string) bool {
	return !rxDefaultMonthPattern.MatchString(str)
}

func regexFilter(str string) bool {
	return strings.Contains(str, "{0}") &&
		!rxNotRelativePattern.MatchString(str) &&
		!rxParenthesisPattern.MatchString(str)
}
