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

	// Prepare helper function
	addMonthTranslations := func(data *LocaleData, monthNumber int) {
		strMonthNumber := strconv.Itoa(monthNumber)
		localNames := filterList(localeMonthFilter,
			gregorianData.Months.Format.Wide[strMonthNumber],
			gregorianData.Months.Format.Abbreviated[strMonthNumber],
			gregorianData.Months.StandAlone.Wide[strMonthNumber],
			gregorianData.Months.StandAlone.Abbreviated[strMonthNumber])

		enMonthName := enMonthNames[monthNumber]
		for _, localName := range localNames {
			data.AddTranslation(localName, enMonthName, true)
		}
	}

	addWeekDayTranslations := func(data *LocaleData, weekDayNumber int) {
		enDayName := enDayNames[weekDayNumber]
		enShortDayName := enDayName[:3]
		localNames := []string{
			gregorianData.Days.Format.Wide[enShortDayName],
			gregorianData.Days.Format.Abbreviated[enShortDayName],
			gregorianData.Days.StandAlone.Wide[enShortDayName],
			gregorianData.Days.StandAlone.Abbreviated[enShortDayName]}

		for _, localName := range localNames {
			data.AddTranslation(localName, enDayName, true)
		}
	}

	addDayPeriodTranslations := func(data *LocaleData, period string) {
		localNames := []string{
			gregorianData.DayPeriods.Format.Wide[period],
			gregorianData.DayPeriods.Format.Abbreviated[period],
			gregorianData.DayPeriods.StandAlone.Wide[period],
			gregorianData.DayPeriods.StandAlone.Abbreviated[period],
		}

		for _, localName := range localNames {
			localName = rxAmPmPattern.ReplaceAllString(localName, period)
			data.AddTranslation(localName, period, true)
		}
	}

	addDateFieldTranslations := func(data *LocaleData, fieldName string) {
		// Create dateField keys
		cleanFieldName := rxSanitizeDateField.ReplaceAllString(fieldName, "")
		cleanFieldName = strings.TrimSpace(cleanFieldName)

		keys := []string{
			cleanFieldName,
			cleanFieldName + "-short",
			cleanFieldName + "-narrow",
		}

		// Fetch patterns
		var localPatterns []string
		var filter func(str string) bool

		if strings.Contains(fieldName, "$") { // Regex field
			filter = regexFilter

			if strings.Contains(fieldName, "in") { // Is future time
				for _, key := range keys {
					localPatterns = append(localPatterns,
						dateFieldsData[key].RelativeTimeTypeFuture.CountOne,
						dateFieldsData[key].RelativeTimeTypeFuture.CountOther)
				}
			} else if strings.Contains(fieldName, "ago") { // Is past time
				for _, key := range keys {
					localPatterns = append(localPatterns,
						dateFieldsData[key].RelativeTimeTypePast.CountOne,
						dateFieldsData[key].RelativeTimeTypePast.CountOther)
				}
			}
		} else { // Handle normal fields
			localPatterns = make([]string, len(keys))

			for i, key := range keys {
				switch {
				case strings.Contains(fieldName, "0"): // Is current time
					localPatterns[i] = dateFieldsData[key].RelativeType0
				case strings.Contains(fieldName, "ago"): // Is past time
					localPatterns[i] = dateFieldsData[key].RelativeTypeMin1
				case strings.Contains(fieldName, "in"): // Is future time
					localPatterns[i] = dateFieldsData[key].RelativeType1
				default: // Is display
					localPatterns[i] = dateFieldsData[key].DisplayName
				}
			}
		}

		// Save the translations
		localPatterns = filterList(filter, localPatterns...)

		for _, localPattern := range localPatterns {
			data.AddTranslation(localPattern, fieldName, true)
		}
	}

	// Generate locale data
	data := LocaleData{
		Name:      locale,
		DateOrder: rxDateOrderPattern.ReplaceAllString(gregorianData.DateFormats.Short, "$1$2$3"),
	}

	for month := 1; month <= 12; month++ {
		addMonthTranslations(&data, month)
	}

	for day := 1; day <= 7; day++ {
		addWeekDayTranslations(&data, day)
	}

	for _, period := range enDayPeriods {
		addDayPeriodTranslations(&data, period)
	}

	for _, dateField := range enDateFields {
		addDateFieldTranslations(&data, dateField)
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
