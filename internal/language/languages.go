package language

import (
	"fmt"
	"sort"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func GetLanguages(locales []string, languages []string, useGivenOrder bool) ([]string, error) {
	// If locales specified, languages is ignored and replaced
	// with languages from locales
	if len(locales) > 0 {
		languages = []string{}
		tracker := strutil.NewDict()

		for _, locale := range locales {
			language := strutil.RemoveRegion(locale)
			if !tracker.Contain(language) {
				tracker.Add(language)
				languages = append(languages, language)
			}
		}
	}

	// If list of languages defined, validate it
	// If not, take from the known language order
	if len(languages) > 0 {
		for _, lang := range languages {
			if _, exist := data.LanguageMap[lang]; !exist {
				return nil, fmt.Errorf("unknown language: %s", lang)
			}
		}
	} else if !useGivenOrder {
		languages = make([]string, len(data.LanguageOrder))
		for lang, order := range data.LanguageOrder {
			languages[order] = lang
		}
	}

	// If required, sort the languages
	if !useGivenOrder {
		sort.Slice(languages, func(a, b int) bool {
			orderA := data.LanguageOrder[languages[a]]
			orderB := data.LanguageOrder[languages[b]]
			if orderA != orderB {
				return orderA < orderB
			}

			return languages[a] < languages[b]
		})
	}

	return languages, nil
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
