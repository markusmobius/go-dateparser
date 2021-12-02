package main

import (
	"os"
	"path/filepath"
)

func parseAllSupplementaryData(languageLocalesMap map[string][]string) (map[string]LocaleData, error) {
	// Parse base data
	var baseData LocaleData
	err := parseYamlFile(&baseData, SUPPLEMENTARY_BASE_PATH)
	if err != nil {
		return nil, err
	}

	// Parse all language
	result := map[string]LocaleData{}

	for language := range languageLocalesMap {
		localeData, err := parseSupplementaryData(language)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil, err
		}

		log.Info().Msgf("parsed supplementary %s", language)
		result[language] = mergeLocaleData(*localeData, baseData)
	}

	return result, nil
}

func parseSupplementaryData(language string) (*LocaleData, error) {
	// Parse YAML
	var data LocaleData
	fPath := filepath.Join(SUPPLEMENTARY_DIR, language+".yaml")
	err := parseYamlFile(&data, fPath)
	if err != nil {
		return nil, err
	}

	// Convert list of simplication to map
	data.Simplifications = map[string]string{}
	for _, simplification := range data.SimplificationList {
		for key, value := range simplification {
			if value != "" {
				data.Simplifications[key] = value
			}
		}
	}

	data.SimplificationList = nil
	if len(data.Simplifications) == 0 {
		data.Simplifications = nil
	}

	return &data, nil
}
