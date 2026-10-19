package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Loaders run at most once per process, so record which locales were built
// before any test requested them; this keeps the check valid under -count.
var builtBeforeTests map[string]bool

func TestMain(m *testing.M) {
	builtBeforeTests = map[string]bool{
		"en": en_Locale.Name != "",
		"fr": fr_Locale.Name != "",
		"zu": zu_Locale.Name != "",
	}
	os.Exit(m.Run())
}

func TestGetLocaleData_DoesNotBuildLocalesAtInit(t *testing.T) {
	for locale, built := range builtBeforeTests {
		assert.False(t, built, "locale %q was built before it was requested", locale)
	}
}

func TestGetLocaleData_BuildsLocaleOnFirstRequest(t *testing.T) {
	locale, ok := GetLocaleData("zu")

	assert.True(t, ok)
	assert.Equal(t, "zu", locale.Name)
	assert.NotNil(t, locale.RxCombined)
}

func TestGetLocaleData_BuildsSubLocaleWithItsParent(t *testing.T) {
	locale, ok := GetLocaleData("fr-CA")

	assert.True(t, ok)
	assert.Equal(t, "fr-CA", locale.Name)
	assert.Equal(t, fr_Locale.RxCombined, locale.RxCombined)
	assert.NotEmpty(t, fr_Locale.Name)
}

// Regexes are compiled lazily, so a bad pattern would otherwise only panic
// when its locale is first used.
func TestGetLocaleData_BuildsEveryLocale(t *testing.T) {
	for name := range LocaleOrder {
		locale, ok := GetLocaleData(name)
		if assert.True(t, ok, name) {
			assert.Equal(t, name, locale.Name)
		}
	}
}

func TestGetLocaleData_UnknownLocale(t *testing.T) {
	locale, ok := GetLocaleData("xx-unknown")

	assert.False(t, ok)
	assert.Nil(t, locale)
}
