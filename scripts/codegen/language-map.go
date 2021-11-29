package main

import "strings"

func createLanguageMap(languageOrder []string) map[string][]string {
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
