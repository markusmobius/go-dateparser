package main

import (
	"os"
	"path/filepath"
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
	// 1. Parse data
	cldrGregorian, err := parseCldrGregorianData(locale)
	if err != nil {
		return nil, err
	}

	cldrDateFields, err := parseCldrDateFieldsData(locale)
	if err != nil {
		return nil, err
	}

	// 2. Extract main data
	gregorianData := cldrGregorian.Dates.Calendars.Gregorian
	dateFieldsData := cldrDateFields.Main[locale].Dates.Fields

	// 3. Prepare initial locale data
	dateOrder := strings.ToUpper(gregorianData.DateFormats.Short.Value)

	data := LocaleData{
		Name:      locale,
		DateOrder: rxDateOrderPattern.ReplaceAllString(dateOrder, "$1$2$3"),
	}

	// 4. Save all local month names to locale data
	allMonthNames := [][]string{
		gregorianData.Months.Format.Wide.List(),
		gregorianData.Months.Format.Abbreviated.List(),
		gregorianData.Months.StandAlone.Wide.List(),
		gregorianData.Months.StandAlone.Abbreviated.List(),
	}

	for _, monthNames := range allMonthNames {
		for idx, monthName := range monthNames {
			// Skip if the month name is a placeholder pattern, e.g
			// "M01" for January,
			// "M02" for February,
			// and so on.
			if !rxDefaultMonthPattern.MatchString(monthName) {
				data.AddTranslation(monthName, enMonthNames[idx], true)
			}
		}
	}

	// 5. Save all local weekday names to locale data
	allWeekdayNames := [][]string{
		gregorianData.Days.Format.Wide.List(),
		gregorianData.Days.Format.Abbreviated.List(),
		gregorianData.Days.StandAlone.Wide.List(),
		gregorianData.Days.StandAlone.Abbreviated.List(),
	}

	for _, weekdayNames := range allWeekdayNames {
		for idx, weekdayName := range weekdayNames {
			data.AddTranslation(weekdayName, enDayNames[idx], true)
		}
	}

	// 6. Save all day period names to locale data
	amPeriodNames := []string{
		gregorianData.DayPeriods.Format.Wide.Am,
		gregorianData.DayPeriods.Format.Abbreviated.Am,
		gregorianData.DayPeriods.StandAlone.Wide.Am,
		gregorianData.DayPeriods.StandAlone.Abbreviated.Am,
	}

	pmPeriodNames := []string{
		gregorianData.DayPeriods.Format.Wide.Pm,
		gregorianData.DayPeriods.Format.Abbreviated.Pm,
		gregorianData.DayPeriods.StandAlone.Wide.Pm,
		gregorianData.DayPeriods.StandAlone.Abbreviated.Pm,
	}

	for _, amName := range amPeriodNames {
		amName = rxAmPmPattern.ReplaceAllString(amName, "am")
		data.AddTranslation(amName, "am", true)
	}

	for _, pmName := range pmPeriodNames {
		pmName = rxAmPmPattern.ReplaceAllString(pmName, "pm")
		data.AddTranslation(pmName, "pm", true)
	}

	// 7. Add data from CLDR Date Fields
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

			if strings.HasPrefix(fieldName, "in ") { // Is future time
				for _, key := range keys {
					localPatterns = append(localPatterns,
						dateFieldsData[key].RelativeTimeTypeFuture.CountOne,
						dateFieldsData[key].RelativeTimeTypeFuture.CountOther)
				}
			} else if strings.HasSuffix(fieldName, " ago") { // Is past time
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
				case strings.HasPrefix(fieldName, "0 "): // Is current time
					localPatterns[i] = dateFieldsData[key].RelativeType0
				case strings.HasSuffix(fieldName, " ago"): // Is past time
					localPatterns[i] = dateFieldsData[key].RelativeTypeMin1
				case strings.HasPrefix(fieldName, "in "): // Is future time
					localPatterns[i] = dateFieldsData[key].RelativeType1
				default: // Is display
					localPatterns[i] = dateFieldsData[key].DisplayName
				}
			}
		}

		// Save the translations
		localPatterns = filterList(filter, localPatterns...)

		fnAdder := data.AddTranslation
		if strings.HasPrefix(fieldName, "in ") || strings.HasSuffix(fieldName, " ago") {
			fnAdder = data.AddRelativeType
		}

		for _, localPattern := range localPatterns {
			fnAdder(localPattern, fieldName, true)
		}
	}

	for _, dateField := range enDateFields {
		addDateFieldTranslations(&data, dateField)
	}

	return &data, nil
}

func parseCldrGregorianData(locale string) (*CldrGregorian, error) {
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")
	fPath := filepath.Join(cldrDatesFullDir, locale, "ca-gregorian.json")

	var container CldrGregorianContainer
	err := parseJsonFile(&container, fPath)
	if err != nil {
		return nil, err
	}

	return &container.Data, nil
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

func regexFilter(str string) bool {
	return strings.Contains(str, "{0}") &&
		!rxNotRelativePattern.MatchString(str) &&
		!rxParenthesisPattern.MatchString(str)
}
