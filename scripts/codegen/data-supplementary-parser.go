package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/elliotchance/pie/v2"
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
	addTranslationFromStrings := func(data *LocaleData, translation string, entries []string) {
		for _, entry := range entries {
			data.AddTranslation(entry, translation, false)
		}
	}

	addTranslationFromMapStrings := func(data *LocaleData, mapStrings map[string][]string) {
		for translation, entries := range mapStrings {
			fnAdder := data.AddTranslation
			if strings.HasPrefix(translation, "in ") || strings.HasSuffix(translation, " ago") {
				fnAdder = data.AddRelativeType
			}

			for _, entry := range entries {
				fnAdder(entry, translation, false)
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

	addTranslationFromStrings(&data, "monday", yamlData.Monday)
	addTranslationFromStrings(&data, "tuesday", yamlData.Tuesday)
	addTranslationFromStrings(&data, "wednesday", yamlData.Wednesday)
	addTranslationFromStrings(&data, "thursday", yamlData.Thursday)
	addTranslationFromStrings(&data, "friday", yamlData.Friday)
	addTranslationFromStrings(&data, "saturday", yamlData.Saturday)
	addTranslationFromStrings(&data, "sunday", yamlData.Sunday)

	addTranslationFromStrings(&data, "january", yamlData.January)
	addTranslationFromStrings(&data, "february", yamlData.February)
	addTranslationFromStrings(&data, "march", yamlData.March)
	addTranslationFromStrings(&data, "april", yamlData.April)
	addTranslationFromStrings(&data, "may", yamlData.May)
	addTranslationFromStrings(&data, "june", yamlData.June)
	addTranslationFromStrings(&data, "july", yamlData.July)
	addTranslationFromStrings(&data, "august", yamlData.August)
	addTranslationFromStrings(&data, "september", yamlData.September)
	addTranslationFromStrings(&data, "october", yamlData.October)
	addTranslationFromStrings(&data, "november", yamlData.November)
	addTranslationFromStrings(&data, "december", yamlData.December)

	addTranslationFromStrings(&data, "decade", yamlData.Decade)
	addTranslationFromStrings(&data, "year", yamlData.Year)
	addTranslationFromStrings(&data, "month", yamlData.Month)
	addTranslationFromStrings(&data, "week", yamlData.Week)
	addTranslationFromStrings(&data, "day", yamlData.Day)
	addTranslationFromStrings(&data, "hour", yamlData.Hour)
	addTranslationFromStrings(&data, "minute", yamlData.Minute)
	addTranslationFromStrings(&data, "second", yamlData.Second)
	addTranslationFromStrings(&data, "ago", yamlData.Ago)
	addTranslationFromStrings(&data, "in", yamlData.In)

	addTranslationFromStrings(&data, "am", yamlData.AM)
	addTranslationFromStrings(&data, "pm", yamlData.PM)

	addTranslationFromMapStrings(&data, yamlData.RelativeType)
	addTranslationFromMapStrings(&data, yamlData.RelativeTypeRegex)

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
	skipWords := cloneSlice(data.SkipWords)
	pertainWords := cloneSlice(data.PertainWords)
	addTranslationFromStrings(&data, "", skipWords)
	addTranslationFromStrings(&data, "", pertainWords)

	return &data, nil
}
