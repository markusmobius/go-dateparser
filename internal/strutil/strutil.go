package strutil

import (
	"encoding/json"
	"strings"
	"unicode"

	"github.com/markusmobius/go-dateparser/internal/regexp"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/unicode/rangetable"
)

var (
	nfkcTransformer    = transform.Chain(norm.NFKD, norm.NFKC)
	unicodeTransformer = transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)

	aposthropeTransformer = runes.Map(func(r rune) rune {
		switch r {
		case '\u2019', // right single quotation mark
			'\u02bc', // modifier letter apostrophe
			'\u02bb', // modifier letter turned comma
			'\u055a', // armenian apostrophe
			'\ua78c', // latin small letter saltillo
			'\u2032', // prime
			'\u2035', // reversed prime
			'\u02b9', // modifier letter prime
			'\uff07': // fullwidth apostrophe
			return '\''
		default:
			return r
		}
	})

	braceTransformer = runes.Remove(runes.In(rangetable.New(
		'{', '}',
		'(', ')',
		'<', '>',
		'[', ']',
	)))

	rxSanitizeSkip     = regexp.MustCompile(`\t|\n|\r|\x{00bb}|,\s\x{0432}\b|\x{200e}|\x{b7}|\x{200f}|\x{064e}|\x{064f}`)
	rxSanitizeRussian  = regexp.MustCompile(`(?i)([\PL\pN])\x{0433}\.`)
	rxSanitizeCroatian = regexp.MustCompile(`(?i)(\pN+)\.\s?(\pN+)\.\s?(\pN+)\.( u)?`)
	rxSanitizePeriod   = regexp.MustCompile(`(?i)([^\pN\.\s])\.`)
	rxSanitizeOn       = regexp.MustCompile(`(?i)^.*?on:\s+(.*)`)
)

// SanitizeSpaces replaces non-breaking spaces into a normal one,
// then remove any excess whitespaces.
func SanitizeSpaces(s string) string {
	s = strings.ReplaceAll(s, "\u00A0", "")
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

// NormalizeAposthrope replace apostrope lookalike chars into the ordinary aposthrope.
func NormalizeAposthrope(s string) string {
	normalized, _, _ := transform.String(aposthropeTransformer, s)
	return normalized
}

// NormalizeCharset is used to normalize charset in a string. Used before
// detecting language of a string.
func NormalizeCharset(str string) string {
	normalized, _, err := transform.String(nfkcTransformer, str)
	if err == nil {
		str = normalized
	}

	str = NormalizeAposthrope(str)
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
	s = rxSanitizeRussian.ReplaceAllString(s, "$1 ")        // remove 'г.' (Russian for year) but not in words
	s = rxSanitizeCroatian.ReplaceAllString(s, "$1.$2.$3 ") // extra '.' and 'u' interferes with parsing relative fractional dates
	s = SanitizeSpaces(s)
	s = rxSanitizePeriod.ReplaceAllString(s, "$1")
	s = rxSanitizeOn.ReplaceAllString(s, "$1")
	s = strings.TrimSuffix(s, ":")
	s = NormalizeAposthrope(s)
	s = strings.TrimSpace(s)
	return s
}

// StripBraces, as it name implies, will remove any braces from the string.
func StripBraces(s string) string {
	stripped, _, _ := transform.String(braceTransformer, s)
	return stripped
}

// RemoveRegion removes region code from a locale, returning the language of locale.
func RemoveRegion(locale string) string {
	lastIdx := strings.LastIndex(locale, "-")
	if lastIdx >= 0 {
		suffix := locale[lastIdx+1:]
		if suffix != "" && suffix == strings.ToUpper(suffix) {
			return locale[:lastIdx]
		}
	}

	return locale
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

// IsNumberOnly check if string only consist of number.
func IsNumberOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Jsonify is used to get a JSON rpresentation of specified data.
func Jsonify(data interface{}) string {
	bt, _ := json.Marshal(data)
	return string(bt)
}
