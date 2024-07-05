package dateparser

import (
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var defaultParser = &Parser{}

// Parse parses string representing date and/or time in recognizable localized formats,
// using the default Parser. Useful for quick use.
func Parse(cfg *Configuration, str string, formats ...string) (date.Date, error) {
	return defaultParser.Parse(cfg, str, formats...)
}

// Search detect the suitable language of the text, then find all substrings of the given
// string which represent date and/or time and parse them using the default parser. Useful
// for quick use.
func Search(cfg *Configuration, text string) (string, []SearchResult, error) {
	return defaultParser.Search(cfg, text)
}

// IsKnownLocale is helper function to check if the specified locale or language can be parsed
// by the Parser. The code must be in ISO639 format.
func IsKnownLocale(code string) bool {
	_, exist := data.GetLocaleData(code)
	return exist
}

// PopTzOffset is helper function to extracts timezone data from str then return the
// str with the timezone code removed.
func PopTzOffset(str string) (cleanStr string, tzName string, tzOffset int) {
	cleanStr, data := timezone.PopTzOffset(str)
	tzName, tzOffset = data.Name, data.Offset
	return
}
