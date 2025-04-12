package main

import (
	"os"
	"path/filepath"

	"github.com/elliotchance/pie/v2"
	"golang.org/x/exp/slices"
)

func parseAllSupplementaryData(languageLocalesMap map[string][]string) (map[string]LocaleData, error) {
	// Parse base data
	baseData, err := parseSupplementaryFile(SUPPLEMENTARY_BASE_PATH)
	if err != nil {
		return nil, err
	}

	// Parse all language
	result := map[string]LocaleData{}

	for language := range languageLocalesMap {
		baseClone := baseData.Clone()
		baseClone.Name = language

		fPath := filepath.Join(SUPPLEMENTARY_DIR, language+".yaml")
		localeData, err := parseSupplementaryFile(fPath)
		if os.IsNotExist(err) {
			result[language] = baseClone
			continue
		} else if err != nil {
			return nil, err
		}

		log.Info().Msgf("parsed supplementary %s", language)
		result[language] = baseClone.Merge(*localeData)
	}

	return result, nil
}

func parseSupplementaryFile(fPath string) (*LocaleData, error) {
	// Parse YAML
	var yamlData SupplementaryData
	err := parseYamlFile(&yamlData, fPath)
	if err != nil {
		return nil, err
	}

	// Prepare helper function
	addTranslations := func(data *LocaleData, translation string, entries []string) {
		for _, entry := range entries {
			data.AddTranslation(entry, translation, false)
		}
	}

	addRelativeTypes := func(data *LocaleData, mapStrings map[string][]string) {
		for translation, entries := range mapStrings {
			for _, entry := range entries {
				data.AddRelativeType(entry, translation, false)
			}
		}
	}

	// Generate locale data
	data := LocaleData{
		SkipWords:             cleanList(false, yamlData.SkipWords...),
		PertainWords:          cleanList(false, yamlData.PertainWords...),
		NoWordSpacing:         yamlData.NoWordSpacing,
		SentenceSplitterGroup: yamlData.SentenceSplitterGroup,
	}

	addTranslations(&data, "monday", yamlData.Monday)
	addTranslations(&data, "tuesday", yamlData.Tuesday)
	addTranslations(&data, "wednesday", yamlData.Wednesday)
	addTranslations(&data, "thursday", yamlData.Thursday)
	addTranslations(&data, "friday", yamlData.Friday)
	addTranslations(&data, "saturday", yamlData.Saturday)
	addTranslations(&data, "sunday", yamlData.Sunday)

	addTranslations(&data, "january", yamlData.January)
	addTranslations(&data, "february", yamlData.February)
	addTranslations(&data, "march", yamlData.March)
	addTranslations(&data, "april", yamlData.April)
	addTranslations(&data, "may", yamlData.May)
	addTranslations(&data, "june", yamlData.June)
	addTranslations(&data, "july", yamlData.July)
	addTranslations(&data, "august", yamlData.August)
	addTranslations(&data, "september", yamlData.September)
	addTranslations(&data, "october", yamlData.October)
	addTranslations(&data, "november", yamlData.November)
	addTranslations(&data, "december", yamlData.December)

	addTranslations(&data, "decade", yamlData.Decade)
	addTranslations(&data, "year", yamlData.Year)
	addTranslations(&data, "month", yamlData.Month)
	addTranslations(&data, "week", yamlData.Week)
	addTranslations(&data, "day", yamlData.Day)
	addTranslations(&data, "hour", yamlData.Hour)
	addTranslations(&data, "minute", yamlData.Minute)
	addTranslations(&data, "second", yamlData.Second)
	addTranslations(&data, "ago", yamlData.Ago)
	addTranslations(&data, "in", yamlData.In)

	addTranslations(&data, "am", yamlData.AM)
	addTranslations(&data, "pm", yamlData.PM)

	addRelativeTypes(&data, yamlData.RelativeType)
	addRelativeTypes(&data, yamlData.RelativeTypeRegex)

	for _, simplification := range yamlData.Simplifications {
		for pattern, replacement := range simplification {
			data.AddSimplification(pattern, replacement)
		}
	}

	// If skip words for some reason redefined again, then it's not skipped
	var checkedSkipWords []string
	for _, w := range data.SkipWords {
		_, transExist := data.Translations[w]
		isPertained := pie.Contains(data.PertainWords, w)
		if !isPertained && !transExist {
			checkedSkipWords = append(checkedSkipWords, w)
		}
	}
	data.SkipWords = checkedSkipWords

	// Save the skipped and pertained words to dictionary
	skipWords := slices.Clone(data.SkipWords)
	pertainWords := slices.Clone(data.PertainWords)
	addTranslations(&data, "", skipWords)
	addTranslations(&data, "", pertainWords)

	return &data, nil
}
