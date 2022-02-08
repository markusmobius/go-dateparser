package dateparser

import "github.com/markusmobius/go-dateparser/date"

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
