package strutil

import (
	"encoding/json"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	nfkcTransformer    = transform.Chain(norm.NFKD, norm.NFKC)
	unicodeTransformer = transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)

	apostropheLookAlikeChars = []string{
		`\x{2019}`, // right single quotation mark
		`\x{02bc}`, // modifier letter apostrophe
		`\x{02bb}`, // modifier letter turned comma
		`\x{055a}`, // armenian apostrophe
		`\x{a78c}`, // latin small letter saltillo
		`\x{2032}`, // prime
		`\x{2035}`, // reversed prime
		`\x{02b9}`, // modifier letter prime
		`\x{ff07}`, // fullwidth apostrophe
	}

	rxNbsp               = regexp.MustCompile(`\x{a0}`)
	rxSanitizeSkip       = regexp.MustCompile(`\t|\n|\r|\x{00bb}|,\s\x{0432}\b|\x{200e}|\x{b7}|\x{200f}|\x{064e}|\x{064f}`)
	rxSanitizeRussian    = regexp.MustCompile(`(?i)([\PL\pN])\x{0433}\.`)
	rxSanitizePeriod     = regexp.MustCompile(`(?i)([^\pN\.])\.`)
	rxSanitizeOn         = regexp.MustCompile(`(?i)^.*?on:\s+(.*)`)
	rxSanitizeAposthrope = regexp.MustCompile(strings.Join(apostropheLookAlikeChars, "|"))
	rxBraces             = regexp.MustCompile(`[{}()<>\[\]]`)
)

// SanitizeSpaces replaces non-breaking spaces into a normal one,
// then remove any excess whitespaces.
func SanitizeSpaces(s string) string {
	s = rxNbsp.ReplaceAllString(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

// NormalizeUnicode removes Nonspacing Mark (Mn) characters then
// use NFKC as the normal forms.
func NormalizeUnicode(str string) string {
	normalized, _, err := transform.String(unicodeTransformer, str)
	if err != nil {
		return str
	}
	return normalized
}

// NormalizeCharset is used to normalize charset in a string. Used before
// detecting language of a string.
func NormalizeCharset(str string) string {
	normalized, _, err := transform.String(nfkcTransformer, str)
	if err == nil {
		str = normalized
	}

	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = SanitizeSpaces(str)
	return str
}

// NormalizeString is used to normalize the string before translated
// into another language. This function will normalizes the unicode,
// convert the string to lowercase then remove any excess whitespaces.
func NormalizeString(str string) string {
	str = NormalizeUnicode(str)
	str = strings.ToLower(str)
	str = SanitizeSpaces(str)
	return str
}

// SanitizeDate is used to sanitize the specified date string before
// it parsed by the parser.
func SanitizeDate(s string) string {
	s = rxSanitizeSkip.ReplaceAllString(s, " ")
	s = rxSanitizeRussian.ReplaceAllString(s, "$1 ") // remove 'г.' (Russian for year) but not in words
	s = SanitizeSpaces(s)
	s = rxSanitizePeriod.ReplaceAllString(s, "$1")
	s = rxSanitizeOn.ReplaceAllString(s, "$1")
	s = strings.TrimSuffix(s, ":")
	s = rxSanitizeAposthrope.ReplaceAllString(s, "'")
	s = strings.TrimSpace(s)
	return s
}

// StripBraces, as it name implies, will remove any braces from the string.
func StripBraces(s string) string {
	return rxBraces.ReplaceAllString(s, "")
}

// TrimChars will trim all character in beginning and end of the string
// that matched with the specified characters.
func TrimChars(s string, chars string) string {
	charMap := map[rune]struct{}{}
	for _, c := range chars {
		charMap[c] = struct{}{}
	}

	return strings.TrimFunc(s, func(r rune) bool {
		_, exist := charMap[r]
		return exist
	})
}

// Jsonify is used to get a JSON rpresentation of specified data.
func Jsonify(data interface{}) string {
	bt, _ := json.Marshal(data)
	return string(bt)
}
