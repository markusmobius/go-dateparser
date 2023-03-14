package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/go-shiori/dom"
)

func createLanguageOrder(languageLocales map[string][]string) ([]string, error) {
	// Parse CLDR territory info
	languagePopulationMap, err := parseCldrTerritory()
	if err != nil {
		return nil, err
	}

	// Parse most common languages from W3Techs
	mostCommonLanguages, err := parseW3CommonLanguages()
	if err != nil {
		return nil, err
	}

	// Separate between languages that has population data or not
	var languages []string
	var unusedLanguages []string
	for language := range languageLocales {
		population := languagePopulationMap[language]
		if population > 0 {
			languages = append(languages, language)
		} else {
			unusedLanguages = append(unusedLanguages, language)
		}
	}

	// Sort languages based on how common is it and how many speaker it has
	sort.Slice(languages, func(i, j int) bool {
		langI := languages[i]
		langJ := languages[j]

		// Check for common language
		iIdx, iIsCommon := mostCommonLanguages[langI]
		jIdx, jIsCommon := mostCommonLanguages[langJ]

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

	for _, unusedLanguage := range unusedLanguages {
		parentLanguage := rxLocaleCleaner.ReplaceAllString(unusedLanguage, "")
		parentIdx := pie.FindFirstUsing(languages, func(v string) bool {
			return v == parentLanguage
		})

		if parentIdx >= 0 {
			languages = pie.Insert(languages, parentIdx+1, unusedLanguage)
		} else {
			languages = append(languages, unusedLanguage)
		}
	}

	return languages, nil
}

func parseCldrTerritory() (map[string]int, error) {
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
	languagePopulations := map[string]int{}
	for _, territoryData := range data.Supplemental.TerritoryInfo {
		population, _ := strconv.ParseFloat(territoryData.Population, 64)
		for language, languageData := range territoryData.LanguagePopulation {
			language = strings.ReplaceAll(language, "_", "-")
			percentage, _ := strconv.ParseFloat(languageData.PopulationPercent, 64)
			languagePopulation := math.Round(population * percentage)
			languagePopulations[language] += int(languagePopulation)
		}
	}

	return languagePopulations, nil
}

func parseW3CommonLanguages() (map[string]int, error) {
	// Open HTML file
	fPath := filepath.Join(RAW_DIR, "w3techs", "content_language.html")
	f, err := os.Open(fPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Parse HTML
	doc, err := dom.Parse(f)
	if err != nil {
		return nil, err
	}

	// Fetch the common languages
	commonLanguages := map[string]int{}
	for i, a := range dom.QuerySelectorAll(doc, "table.bars th a") {
		href := dom.GetAttribute(a, "href")
		langCode := path.Base(href)
		langCode = strings.ToLower(langCode)
		langCode = strings.TrimPrefix(langCode, "cl")
		langCode = strings.Trim(langCode, "-")
		commonLanguages[langCode] = i
	}

	// If English is not the most common language, it's error
	if commonLanguages["en"] != 0 {
		return nil, fmt.Errorf("english is not the first language")
	}

	return commonLanguages, nil
}
