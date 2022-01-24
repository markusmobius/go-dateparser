package language

import (
	"fmt"
	"strings"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	// Helper function
	sp := func(items ...string) []string { return items }

	// Prepare scenarios
	type testScenario struct {
		Languages    []string
		Locales      []string
		Region       string
		Expected     []string
		NoGivenOrder bool
	}

	tests := []testScenario{
		// Basic test
		{
			Languages: sp("es", "fr", "en"),
			Expected:  sp("es", "fr", "en"),
		}, {
			Locales:  sp("fr-BF", "ar-SO", "asa"),
			Expected: sp("fr-BF", "ar-SO", "asa"),
		}, {
			Locales:  sp("nl-CW", "so-KE", "sr-Latn-XK", "zgh"),
			Expected: sp("nl-CW", "so-KE", "sr-Latn-XK", "zgh"),
		}, {
			Languages: sp("pt", "zh-Hant", "zh-Hans"),
			Region:    "MO",
			Expected:  sp("pt-MO", "zh-Hant-MO", "zh-Hans-MO"),
		}, {
			Languages: sp("pt", "he"),
			Locales:   sp("ru-UA", "ckb-IR", "sq-XK"),
			Region:    "MO",
			Expected:  sp("ru-UA", "ckb-IR", "sq-XK"),
		}, {
			Languages: sp("da", "ja"),
			Locales:   sp("shi-Latn", "teo-KE", "ewo", "vun"),
			Region:    "AO",
			Expected:  sp("shi-Latn", "teo-KE", "ewo", "vun"),
		},
		// Load by given order
		{
			Locales:  sp("es-MX", "ar-EG", "fr-DJ"),
			Expected: sp("es-MX", "ar-EG", "fr-DJ"),
		}, {
			Locales:  sp("pt-MO", "ru-KZ", "es-CU"),
			Expected: sp("pt-MO", "ru-KZ", "es-CU"),
		}, {
			Locales:  sp("zh-Hans-HK", "en-VG", "my"),
			Expected: sp("zh-Hans-HK", "en-VG", "my"),
		}, {
			Locales:  sp("tl", "nl-SX", "de-BE"),
			Expected: sp("tl", "nl-SX", "de-BE"),
		}, {
			Locales:  sp("sw-KE", "ru-UA", "he"),
			Expected: sp("sw-KE", "ru-UA", "he"),
		}, {
			Locales:  sp("de-IT", "ta-MY", "pa-Arab"),
			Expected: sp("de-IT", "ta-MY", "pa-Arab"),
		}, {
			Locales:  sp("bn-IN", "pt-ST", "ko-KP", "ta"),
			Expected: sp("bn-IN", "pt-ST", "ko-KP", "ta"),
		}, {
			Locales:  sp("fr-NE", "ar-SY"),
			Expected: sp("fr-NE", "ar-SY"),
		},
		// Load NOT by given order
		{
			Locales:      sp("os-RU", "ln-CF", "ee-TG"),
			Expected:     sp("ee-TG", "ln-CF", "os-RU"),
			NoGivenOrder: true,
		}, {
			Locales:      sp("fo-DK", "khq"),
			Expected:     sp("khq", "fo-DK"),
			NoGivenOrder: true,
		}, {
			Locales:      sp("en-CC", "fr-BE", "ar-KW"),
			Expected:     sp("en-CC", "fr-BE", "ar-KW"),
			NoGivenOrder: true,
		}, {
			Locales:      sp("en-FJ", "pt-CV", "fr-RW"),
			Expected:     sp("en-FJ", "fr-RW", "pt-CV"),
			NoGivenOrder: true,
		}, {
			Locales:      sp("pt-AO", "hi", "zh-Hans-SG", "vi"),
			Expected:     sp("pt-AO", "vi", "zh-Hans-SG", "hi"),
			NoGivenOrder: true,
		}, {
			Locales:      sp("gsw-FR", "es-BZ", "ca-IT", "qu-EC"),
			Expected:     sp("es-BZ", "qu-EC", "ca-IT", "gsw-FR"),
			NoGivenOrder: true,
		},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare log message
		message := strutil.Jsonify(&test)

		// Fetch the locales
		loadedLocales, _ := GetLocales(test.Locales, test.Languages, test.Region, !test.NoGivenOrder, false)
		localeNames := make([]string, len(loadedLocales))
		for i, l := range loadedLocales {
			localeNames[i] = l.Name
		}

		passed := assert.Equal(t, test.Expected, localeNames, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestGetLocales_error(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Languages              []string
		Locales                []string
		AllowConflictingLocale bool
		Error                  error
		ErrorEntries           []string
	}

	sp := func(items ...string) []string { return items }

	unknownLanguages := func(languages, errorEntries []string) testScenario {
		return testScenario{Languages: languages, Error: ErrUnknownLanguages, ErrorEntries: errorEntries}
	}

	unknownLocales := func(locales, errorEntries []string) testScenario {
		return testScenario{Locales: locales, Error: ErrUnknownLocales, ErrorEntries: errorEntries}
	}

	conflictingLocales := func(locales ...string) testScenario {
		return testScenario{Locales: locales, Error: ErrConflictingLocales}
	}

	noConflictingLocales := func(locales ...string) testScenario {
		return testScenario{Locales: locales, AllowConflictingLocale: true}
	}

	tests := []testScenario{
		unknownLanguages(sp("es", "ar-001", "xx"), sp("ar-001", "xx")),
		unknownLanguages(sp("sr-Latn", "sq", "ii-Latn"), sp("ii-Latn")),
		unknownLanguages(sp("vol", "bc"), sp("vol", "bc")),

		unknownLocales(sp("es-AB", "ar-001", "fr-DJ"), sp("es-AB", "ar-001")),
		unknownLocales(sp("ru-MD", "my-MY", "zz"), sp("my-MY", "zz")),
		unknownLocales(sp("nl-SX", "be-BE", "ca-FR"), sp("be-BE")),
		unknownLocales(sp("ca-TA", "ca", "fr"), sp("ca-TA")),

		conflictingLocales("en-TK", "en-TO", "zh"),
		conflictingLocales("es-PY", "es-IC", "ja", "es-DO"),

		noConflictingLocales("en-TK", "en-TO", "zh"),
		noConflictingLocales("es-PY", "es-IC", "ja", "es-DO"),
		noConflictingLocales("af-NA", "da", "af"),
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare log message
		message := strutil.Jsonify(&test)

		// Fetch the locales
		_, err := GetLocales(test.Locales, test.Languages, "", true, test.AllowConflictingLocale)
		passed := assert.ErrorIs(t, err, test.Error, message)
		if passed && len(test.ErrorEntries) > 0 {
			str := strings.Join(test.ErrorEntries, ", ")
			str = fmt.Sprintf("%v: %s", test.Error, str)
			assert.Equal(t, str, err.Error())
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
