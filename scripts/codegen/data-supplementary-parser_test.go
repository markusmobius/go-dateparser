package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupplementaryLocaleOverrides(t *testing.T) {
	base, locales := parseSupplementaryData(SupplementaryData{
		Monday: []string{"base-monday"},
		LocaleSpecific: map[string]SupplementaryData{
			"en-US": {DateOrder: "MDY", Monday: []string{"local-monday"}, SkipWords: []string{"skip"}},
		},
	})
	assert.Equal(t, []string{"monday"}, base.Translations["base-monday"])
	assert.NotContains(t, base.Translations, "local-monday")
	assert.Equal(t, "en-US", locales["en-US"].Name)
	assert.Equal(t, "MDY", locales["en-US"].DateOrder)
	assert.Equal(t, []string{"monday"}, locales["en-US"].Translations["local-monday"])
	assert.Equal(t, []string{""}, locales["en-US"].Translations["skip"])
}
