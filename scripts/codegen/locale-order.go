package main

func createLocaleOrder(languageLocalesMap map[string][]string, languageOrder []string) map[string]int {
	// Create language order map
	languageOrderMap := map[string]int{}
	for i, language := range languageOrder {
		languageOrderMap[language] = i
	}

	// Create locale order
	localeOrderMap := map[string]int{}
	for language, locales := range languageLocalesMap {
		localeOrderMap[language] = languageOrderMap[language]
		for _, locale := range locales {
			localeOrderMap[locale] = languageOrderMap[language]
		}
	}

	return localeOrderMap
}
