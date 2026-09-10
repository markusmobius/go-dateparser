// Package dateparser parses localized date strings and extracts dates from text.
// Generated re2go matchers are enabled by default with or without cgo; users do
// not need the generator or a build tag. Optional RE2 backends affect only the
// remaining regex operations.
package dateparser

import (
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var defaultParser = &Parser{}

// Parse parses a localized date or time using the shared default Parser.
// A nil configuration uses the defaults. Optional formats are Go time layouts.
func Parse(cfg *Configuration, str string, formats ...string) (date.Date, error) {
	return defaultParser.Parse(cfg, str, formats...)
}

// Search detects a language and extracts dates using the shared default Parser.
// It returns the selected language code, matches, and any error.
// A nil configuration uses the default split search strategy.
func Search(cfg *Configuration, text string) (string, []SearchResult, error) {
	return defaultParser.Search(cfg, text)
}

// IsKnownLocale reports whether a language or locale code is supported,
// for example "en", "fr-PF", or "zh-Hant".
func IsKnownLocale(code string) bool {
	_, exist := data.GetLocaleData(code)
	return exist
}

// PopTzOffset extracts a timezone from str, returning the remaining string,
// the timezone name, and its offset east of UTC in seconds.
func PopTzOffset(str string) (cleanStr string, tzName string, tzOffset int) {
	cleanStr, data := timezone.PopTzOffset(str)
	tzName, tzOffset = data.Name, data.Offset
	return
}
