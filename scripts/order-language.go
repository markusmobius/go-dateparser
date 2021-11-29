//go:build ignore

package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type CldrTerritoryData struct {
	Supplemental struct {
		TerritoryInfo map[string]struct {
			Gdp                string `json:"_gdp"`
			LiteracyPercent    string `json:"_literacyPercent"`
			Population         string `json:"_population"`
			LanguagePopulation map[string]struct {
				PopulationPercent string `json:"_populationPercent"`
				OfficialStatus    string `json:"_officialStatus"`
			} `json:"languagePopulation"`
		} `json:"territoryInfo"`
	} `json:"supplemental"`
}

var (
	rxLocale        = regexp.MustCompile(`-[A-Z0-9]+$`)
	rxLocaleCleaner = regexp.MustCompile(`-\w+`)

	// Languages with insufficient translation data are excluded
	excludedLanguages = map[string]struct{}{
		"cu":       {},
		"kkj":      {},
		"nds":      {},
		"prg":      {},
		"tk":       {},
		"vai":      {},
		"vai-Latn": {},
		"vai-Vaii": {},
		"vo":       {},
	}

	// List of most common locales from https://w3techs.com/technologies/overview/content_language
	// Last updated on 30.10.2020
	mostCommonLocales = map[string]int{
		"en": 1,
		"ru": 2,
		"es": 3,
		"tr": 4,
		"fa": 5,
		"fr": 6,
		"de": 7,
		"ja": 8,
		"pt": 9,
		"vi": 10,
		"zh": 11,
		"ar": 12,
		"it": 13,
		"pl": 14,
		"id": 15,
		"el": 16,
		"nl": 17,
		"ko": 18,
		"th": 19,
		"he": 20,
		"uk": 21,
		"cs": 22,
		"sv": 23,
		"ro": 24,
		"hu": 25,
		"da": 26,
		"sr": 27,
		"sk": 28,
		"fi": 29,
		"bg": 30,
		"hr": 31,
		"lt": 32,
		"hi": 33,
		"nb": 34,
		"sl": 35,
		"nn": 36,
		"et": 37,
		"lv": 38,
	}
)

func getLanguageLocaleMap(rawDir string) (map[string][]string, error) {
	cldrDatesFullDir := filepath.Join(rawDir, "cldr-dates-full", "main")

	// Fetch all CLDR languages and their locales
	dirItems, err := ioutil.ReadDir(cldrDatesFullDir)
	if err != nil {
		return nil, err
	}

	languageNames := map[string]struct{}{}
	languageLocaleMap := map[string][]string{}
	for _, item := range dirItems {
		if !item.IsDir() {
			continue
		}

		// Check if dir name contains locale
		name := item.Name()
		locale := rxLocale.FindString(name)

		if locale == "" {
			languageNames[name] = struct{}{}
			if _, exist := languageLocaleMap[name]; !exist {
				languageLocaleMap[name] = []string{}
			}
		} else {
			language := strings.TrimSuffix(name, locale)
			languageLocaleMap[language] = append(languageLocaleMap[language], name)
		}
	}

	// Remove invalid or excluded language
	for language := range languageLocaleMap {
		_, isValid := languageNames[language]
		_, isExcluded := excludedLanguages[language]
		if !isValid || isExcluded || language == "root" {
			delete(languageLocaleMap, language)
		}
	}

	// Add supplementary languages to the map
	supplementaryLanguages, err := getSupplementaryLanguages()
	if err != nil {
		return nil, err
	}

	for _, language := range supplementaryLanguages {
		if _, exist := languageLocaleMap[language]; !exist {
			languageLocaleMap[language] = []string{}
		}
	}

	return languageLocaleMap, nil
}

func getSupplementaryLanguages() ([]string, error) {
	supplementaryDir := "./data/supplementary/date_translation_data"
	dirItems, err := ioutil.ReadDir(supplementaryDir)
	if err != nil {
		return nil, err
	}

	var languages []string
	for _, item := range dirItems {
		if item.IsDir() {
			continue
		}

		name := item.Name()
		ext := filepath.Ext(name)
		if ext != ".yaml" {
			continue
		}

		name = strings.TrimSuffix(name, ext)
		languages = append(languages, name)
	}

	return languages, nil
}

func getLanguageOrder(rawDir string, languageLocaleMap map[string][]string) ([]string, error) {
	// Open territory info file
	fPath := filepath.Join(rawDir, "cldr-core/supplemental/territoryInfo.json")
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
	for language := range languageLocaleMap {
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

func getLanguageMap(languageOrder []string) map[string][]string {
	languageMap := map[string][]string{}

	for _, language := range languageOrder {
		key := language
		if strings.Contains(key, "-") {
			key = rxLocaleCleaner.ReplaceAllString(key, "")
		}
		languageMap[key] = append(languageMap[key], language)
	}

	return languageMap
}

func main() {
	// Find raw dir path in arguments
	rawDir := "./raw_data"
	if len(os.Args) > 1 {
		rawDir = os.Args[1]
	}

	// Generate language locale map
	languageLocales, err := getLanguageLocaleMap(rawDir)
	if err != nil {
		panic(err)
	}

	// Order languages by its popularity
	languageOrder, err := getLanguageOrder(rawDir, languageLocales)
	if err != nil {
		panic(err)
	}

	// Get language map
	getLanguageMap(languageOrder)
}
