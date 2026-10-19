package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocaleData_BuildsLocaleOnFirstRequest(t *testing.T) {
	assert.Empty(t, zu_Locale.Name)

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
