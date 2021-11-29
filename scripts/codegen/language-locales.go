package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func createLanguageLocales() (map[string][]string, error) {
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")

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
	dirItems, err := ioutil.ReadDir(SUPPLEMENTARY_DIR)
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
