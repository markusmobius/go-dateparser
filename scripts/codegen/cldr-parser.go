package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func parseCldrData(languageLocalesMap map[string][]string) error {
	for language, locales := range languageLocalesMap {
		log.Info().Msgf("parsing data for language %s", language)

		_, err := parseCldrGregorianData(language)
		if err != nil {
			return err
		}

		_, err = parseCldrDateFieldsData(language)
		if err != nil {
			return err
		}

		for _, locale := range locales {
			_, err = parseCldrGregorianData(locale)
			if err != nil {
				return err
			}

			_, err = parseCldrDateFieldsData(locale)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
