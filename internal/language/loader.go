package language

import (
	"fmt"
	"sort"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
)

func GetLocales(languages []string, locales []string, region string, useGivenOrder bool, allowConflictingLocales bool) ([]*data.LocaleData, error) {
	var validLocales []string
	localeTracker := map[string]struct{}{}
	languageTracker := map[string]struct{}{}

	// If locales specified, process it first
	if len(locales) > 0 {
		var invalidLocales []string

		for _, locale := range locales {
			// If locale already registered, skip
			if _, exist := localeTracker[locale]; exist {
				continue
			}

			// Check if the locale valid
			if _, exist := data.LocaleDataMap[locale]; !exist {
				invalidLocales = append(invalidLocales, locale)
				continue
			}

			// Make sure language is only used once
			if !allowConflictingLocales {
				language := rxRegionRemover.ReplaceAllString(locale, "")
				if _, exist := languageTracker[language]; exist {
					return nil, ErrConflictingLocales
				}

				languageTracker[language] = struct{}{}
			}

			// Save the locale
			localeTracker[locale] = struct{}{}
			validLocales = append(validLocales, locale)
		}

		// If there are invalid locales, stop
		if len(invalidLocales) > 0 {
			return nil, fmt.Errorf("%w: %s", ErrUnknownLocales, strings.Join(invalidLocales, ", "))
		}
	} else {
		// If languages is not specified, generate it from language order
		if len(languages) == 0 {
			languages = make([]string, len(data.LanguageOrder))
			for lang, order := range data.LanguageOrder {
				languages[order] = lang
			}
		}

		// Tidy up region value
		region = strings.ToUpper(region)
		region = strings.TrimSpace(region)

		// Generate locales from languages and regions
		var unknownLanguages []string

		for _, language := range languages {
			// If language already checked, skip
			if _, exist := languageTracker[language]; exist {
				continue
			} else {
				languageTracker[language] = struct{}{}
			}

			// Check if language is known
			if _, exist := data.LanguageOrder[language]; !exist {
				unknownLanguages = append(unknownLanguages, language)
				continue
			}

			// Generate and save locale
			locale := language
			if region != "" {
				locale = language + "-" + region
			}

			// If locale already registered, skip
			if _, exist := localeTracker[locale]; exist {
				continue
			}

			// Make sure generated locale is valid
			if _, exist := data.LocaleDataMap[locale]; !exist {
				continue
			}

			// Save the generated locale
			localeTracker[locale] = struct{}{}
			validLocales = append(validLocales, locale)
		}

		// If there are unknown languages, stop
		if len(unknownLanguages) > 0 {
			return nil, fmt.Errorf("%w: %s", ErrUnknownLanguages, strings.Join(unknownLanguages, ", "))
		}
	}

	// If required, sort the locales
	if !useGivenOrder {
		sort.Slice(validLocales, func(a, b int) bool {
			localeA := validLocales[a]
			localeB := validLocales[b]
			languageA := rxRegionRemover.ReplaceAllString(localeA, "")
			languageB := rxRegionRemover.ReplaceAllString(localeB, "")

			if languageA != languageB {
				orderA := data.LanguageOrder[languageA]
				orderB := data.LanguageOrder[languageB]
				return orderA < orderB
			} else {
				return localeA < localeB
			}
		})
	}

	// Fetch the locale data
	listLocaleData := make([]*data.LocaleData, len(validLocales))
	for i, locale := range validLocales {
		listLocaleData[i] = data.LocaleDataMap[locale]
	}

	if len(listLocaleData) == 0 {
		return nil, ErrNotFound
	}

	return listLocaleData, nil
}

func GetLocale(shortname string) (*data.LocaleData, error) {
	locales, err := GetLocales(nil, []string{shortname}, "", false, false)
	if err != nil {
		return nil, err
	}

	return locales[0], nil
}
