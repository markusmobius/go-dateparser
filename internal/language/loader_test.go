package language

import (
	"fmt"
	"strings"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestGetLocales(t *testing.T) {
	// Helper functions
	assertLocaleNames := func(result []*data.LocaleData, expectedNames ...string) {
		names := make([]string, len(result))
		for i, data := range result {
			names[i] = data.Name
		}
		assert.Equal(t, expectedNames, names)
	}

	assertError := func(err error, expected error, expectedEntries ...string) {
		assert.ErrorIs(t, err, expected)
		if len(expectedEntries) > 0 {
			str := strings.Join(expectedEntries, ", ")
			str = fmt.Sprintf("%v: %s", expected, str)
			assert.Equal(t, str, err.Error())
		}
	}

	// Test loading data using various parameters
	languages := []string{"es", "fr", "en"}
	result, _ := GetLocales(languages, nil, "", true, false)
	assertLocaleNames(result, languages...)

	locales := []string{"fr-BF", "ar-SO", "asa"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"nl-CW", "so-KE", "sr-Latn-XK", "zgh"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	region := "MO"
	languages = []string{"pt", "zh-Hant", "zh-Hans"}
	result, _ = GetLocales(languages, nil, region, true, false)
	assertLocaleNames(result, "pt-MO", "zh-Hant-MO", "zh-Hans-MO")

	region = "MO"
	languages = []string{"pt", "he"}
	locales = []string{"ru-UA", "ckb-IR", "sq-XK"}
	result, _ = GetLocales(languages, locales, region, true, false)
	assertLocaleNames(result, locales...)

	region = "AO"
	languages = []string{"da", "ja"}
	locales = []string{"shi-Latn", "teo-KE", "ewo", "vun"}
	result, _ = GetLocales(languages, locales, region, true, false)
	assertLocaleNames(result, locales...)

	// Test loading data using locales with given order
	locales = []string{"es-MX", "ar-EG", "fr-DJ"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"pt-MO", "ru-KZ", "es-CU"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"zh-Hans-HK", "en-VG", "my"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"tl", "nl-SX", "de-BE"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"sw-KE", "ru-UA", "he"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"de-IT", "ta-MY", "pa-Arab"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"bn-IN", "pt-ST", "ko-KP", "ta"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	locales = []string{"fr-NE", "ar-SY"}
	result, _ = GetLocales(nil, locales, "", true, false)
	assertLocaleNames(result, locales...)

	// Test loading data using locales but not with given order
	locales = []string{"os-RU", "ln-CF", "ee-TG"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "ee-TG", "ln-CF", "os-RU")

	locales = []string{"fo-DK", "khq"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "khq", "fo-DK")

	locales = []string{"en-CC", "fr-BE", "ar-KW"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "en-CC", "fr-BE", "ar-KW")

	locales = []string{"en-FJ", "pt-CV", "fr-RW"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "en-FJ", "fr-RW", "pt-CV")

	locales = []string{"pt-AO", "hi", "zh-Hans-SG", "vi"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "pt-AO", "vi", "zh-Hans-SG", "hi")

	locales = []string{"gsw-FR", "es-BZ", "ca-IT", "qu-EC"}
	result, _ = GetLocales(nil, locales, "", false, false)
	assertLocaleNames(result, "es-BZ", "qu-EC", "ca-IT", "gsw-FR")

	// Test error raised for unknown languages
	languages = []string{"es", "ar-001", "xx"}
	_, err := GetLocales(languages, nil, "", true, false)
	assertError(err, ErrUnknownLanguages, "ar-001", "xx")

	languages = []string{"sr-Latn", "sq", "ii-Latn"}
	_, err = GetLocales(languages, nil, "", true, false)
	assertError(err, ErrUnknownLanguages, "ii-Latn")

	languages = []string{"vol", "bc"}
	_, err = GetLocales(languages, nil, "", true, false)
	assertError(err, ErrUnknownLanguages, "vol", "bc")

	// Test error raised for unknown locales
	locales = []string{"es-AB", "ar-001", "fr-DJ"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrUnknownLocales, "es-AB", "ar-001")

	locales = []string{"ru-MD", "my-MY", "zz"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrUnknownLocales, "my-MY", "zz")

	locales = []string{"nl-SX", "be-BE", "ca-FR"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrUnknownLocales, "be-BE")

	locales = []string{"ca-TA", "ca", "fr"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrUnknownLocales, "ca-TA")

	// Test error raised for conflicting locales
	locales = []string{"en-TK", "en-TO", "zh"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrConflictingLocales)

	locales = []string{"es-PY", "es-IC", "ja", "es-DO"}
	_, err = GetLocales(nil, locales, "", true, false)
	assertError(err, ErrConflictingLocales)

	// Test no error raised when conflicting locales allowed
	locales = []string{"en-TK", "en-TO", "zh"}
	_, err = GetLocales(nil, locales, "", true, true)
	assert.Nil(t, err)

	locales = []string{"es-PY", "es-IC", "ja", "es-DO"}
	_, err = GetLocales(nil, locales, "", true, true)
	assert.Nil(t, err)

	locales = []string{"af-NA", "da", "af"}
	_, err = GetLocales(nil, locales, "", true, true)
	assert.Nil(t, err)
}
