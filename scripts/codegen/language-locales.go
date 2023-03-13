package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/zyedidia/generic/mapset"
)

func createLanguageLocales() (map[string][]string, error) {
	cldrDatesFullDir := filepath.Join(RAW_DIR, "cldr-dates-full", "main")

	// Fetch all CLDR languages and their locales
	dirItems, err := ioutil.ReadDir(cldrDatesFullDir)
	if err != nil {
		return nil, err
	}

	languageNames := mapset.New[string]()
	languageLocales := map[string][]string{}
	for _, item := range dirItems {
		if !item.IsDir() {
			continue
		}

		// Check if dir name contains locale. If false, then this locale is the
		// "language" which will become parents for several sub-locales.
		name := item.Name()
		locale := rxLocale.FindString(name)

		if locale == "" {
			languageNames.Put(name)
			if _, exist := languageLocales[name]; !exist {
				languageLocales[name] = []string{}
			}
		} else {
			language := strings.TrimSuffix(name, locale)
			languageLocales[language] = append(languageLocales[language], name)
		}
	}

	// Remove invalid or excluded language
	for language := range languageLocales {
		isValid := languageNames.Has(language)
		_, isExcluded := excludedLanguages[language]
		if !isValid || isExcluded || language == "root" {
			delete(languageLocales, language)
		}
	}

	// Add supplementary languages to the map
	supplementaryLanguages, err := getSupplementaryLanguages()
	if err != nil {
		return nil, err
	}

	for _, language := range supplementaryLanguages {
		if _, exist := languageLocales[language]; !exist {
			languageLocales[language] = []string{}
		}
	}

	return languageLocales, nil
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
