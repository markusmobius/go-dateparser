package main

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func createLanguageOrder(languageLocalesMap map[string][]string) ([]string, error) {
	// Open territory info file
	fPath := filepath.Join(RAW_DIR, "cldr-core/supplemental/territoryInfo.json")
	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decode JSON
	var data CldrTerritoryData
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return nil, err
	}

	// Get population per language
	languagePopulationMap := map[string]int{}
	for _, territoryData := range data.Supplemental.TerritoryInfo {
		population, _ := strconv.ParseFloat(territoryData.Population, 64)
		for language, languageData := range territoryData.LanguagePopulation {
			language = strings.ReplaceAll(language, "_", "-")
			percentage, _ := strconv.ParseFloat(languageData.PopulationPercent, 64)
			languagePopulation := math.Round(population * percentage)
			languagePopulationMap[language] += int(languagePopulation)
		}
	}

	// Separate between languages that has population data or not
	var languageList []string
	var unusedLanguages []string
	for language := range languageLocalesMap {
		population := languagePopulationMap[language]
		if population > 0 {
			languageList = append(languageList, language)
		} else {
			unusedLanguages = append(unusedLanguages, language)
		}
	}

	// Sort languages based on how common is it and how many speaker it has
	sort.Slice(languageList, func(i, j int) bool {
		langI := languageList[i]
		langJ := languageList[j]

		// Check for common language
		iIdx, iIsCommon := mostCommonLocales[langI]
		jIdx, jIsCommon := mostCommonLocales[langJ]

		switch {
		case iIsCommon && jIsCommon:
			return iIdx < jIdx
		case iIsCommon && !jIsCommon:
			return true
		case !iIsCommon && jIsCommon:
			return false
		}

		// Check for population
		iPopulation := languagePopulationMap[langI]
		jPopulation := languagePopulationMap[langJ]
		if iPopulation == jPopulation {
			return langI > langJ
		} else {
			return iPopulation > jPopulation
		}
	})

	// Put back the unused languages
	sort.Strings(unusedLanguages)

	parentLanguageIndexes := map[string]int{}
	for _, language := range unusedLanguages {
		parentLanguage := rxLocaleCleaner.ReplaceAllString(language, "")
		parentIdx, exist := parentLanguageIndexes[parentLanguage]
		if !exist {
			parentIdx = -1
			for i, lang := range languageList {
				if lang == parentLanguage {
					parentIdx = i + 1
					break
				}
			}
			parentLanguageIndexes[parentLanguage] = parentIdx
		}

		if parentIdx >= 0 && parentIdx != len(languageList) {
			languageList = append(languageList, "")
			copy(languageList[parentIdx+1:], languageList[parentIdx:])
			languageList[parentIdx] = language
			parentLanguageIndexes[parentLanguage]++
		} else {
			languageList = append(languageList, language)
		}
	}

	return languageList, nil
}
