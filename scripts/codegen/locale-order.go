package main

func createLocaleOrder(languageLocales map[string][]string, languageOrder []string) map[string]int {
	// Create language order map
	languageOrders := map[string]int{}
	for i, language := range languageOrder {
		languageOrders[language] = i
	}

	// Create locale order
	localeOrders := map[string]int{}
	for language, locales := range languageLocales {
		localeOrders[language] = languageOrders[language]
		for _, locale := range locales {
			localeOrders[locale] = languageOrders[language]
		}
	}

	return localeOrders
}
