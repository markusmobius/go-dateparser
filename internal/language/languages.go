package language

import (
	"fmt"
	"sort"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func GetLanguages(locales, languages []string, useGivenOrder bool) ([]string, error) {
	var result []string
	languageTracker := strutil.NewDict()

	// The search order is:
	// - Locales
	// - Languages
	// - External detector
	if len(locales) > 0 {
		var invalidLocales []string

		for _, locale := range locales {
			// Check if the locale valid
			if _, exist := data.GetLocaleData(locale); !exist {
				invalidLocales = append(invalidLocales, locale)
				continue
			}

			// Save locale's language to tracker
			language := strutil.RemoveRegion(locale)
			if !languageTracker.Contain(language) {
				languageTracker.Add(language)
				result = append(result, language)
			}
		}

		// If there are invalid locales, stop
		if len(invalidLocales) > 0 {
			return nil, fmt.Errorf("unknown locale(s): %s", strings.Join(invalidLocales, ", "))
		}
	} else {
		// Fetch languages from params
		if len(languages) > 0 {
			result = append([]string{}, languages...)
		}

		// Check for unknown languages
		var unknownLanguages []string
		for _, lang := range languages {
			if _, exist := data.LanguageOrder[lang]; !exist {
				unknownLanguages = append(unknownLanguages, lang)
			}
		}

		if len(unknownLanguages) > 0 {
			return nil, fmt.Errorf("unknown language(s): %s", strings.Join(unknownLanguages, ", "))
		}
	}

	// If required, sort the languages
	if !useGivenOrder {
		sort.Slice(result, func(a, b int) bool {
			orderA := data.LanguageOrder[result[a]]
			orderB := data.LanguageOrder[result[b]]
			if orderA != orderB {
				return orderA < orderB
			}

			return result[a] < result[b]
		})
	}

	// At this point, if there are still no languages just generate it from language order
	if len(result) == 0 {
		result = make([]string, len(data.LanguageOrder))
		for lang, order := range data.LanguageOrder {
			result[order] = lang
		}
	}

	return result, nil
}

func GetUniqueCharsets(languages []string) map[string][]rune {
	// Prepare result
	uniqueCharsets := map[string][]rune{}

	// Process each charset
	for _, language := range languages {
		// Fetch locale data
		ld, exist := data.GetLocaleData(language)
		if !exist {
			continue
		}

		// Initiate map to contain the unique chars for this language
		mapUniqueChar := map[rune]struct{}{}
		for _, char := range ld.Charset {
			mapUniqueChar[char] = struct{}{}
		}

		// Look for charset in the other languages
		for _, otherLanguage := range languages {
			if otherLanguage == language {
				continue
			}

			ld, exist := data.GetLocaleData(otherLanguage)
			if !exist {
				continue
			}

			// If the charset of this language exist in map, remove it
			for _, char := range ld.Charset {
				delete(mapUniqueChar, char)
			}
		}

		// Convert map to list
		var uniqueCharset []rune
		for char := range mapUniqueChar {
			uniqueCharset = append(uniqueCharset, char)
		}

		uniqueCharsets[language] = uniqueCharset
	}

	return uniqueCharsets
}
