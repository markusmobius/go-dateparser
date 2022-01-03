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
	rxSanitizeRussian    = regexp.MustCompile(`(?i)([\W\d])\x{0433}\.`)
	rxSanitizePeriod     = regexp.MustCompile(`(?i)(\D+)\.`)
	rxSanitizeOn         = regexp.MustCompile(`(?i)^.*?on:\s+(.*)`)
	rxSanitizeAposthrope = regexp.MustCompile(strings.Join(apostropheLookAlikeChars, "|"))
)

func NormalizeUnicode(str string) string {
	normalized, _, err := transform.String(unicodeTransformer, str)
	if err != nil {
		return str
	}
	return normalized
}

func NormalizeString(str string) string {
	str = NormalizeUnicode(str)
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}

func SanitizeSpaces(s string) string {
	s = rxNbsp.ReplaceAllString(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func SanitizeDate(s string) string {
	s = rxSanitizeSkip.ReplaceAllString(s, " ")
	s = rxSanitizeRussian.ReplaceAllString(s, "$1 ")
	s = SanitizeSpaces(s)
	s = rxSanitizePeriod.ReplaceAllString(s, "$1")
	s = rxSanitizeOn.ReplaceAllString(s, "$1")
	s = strings.TrimSuffix(s, ":")
	s = rxSanitizeAposthrope.ReplaceAllString(s, "'")
	s = strings.TrimSpace(s)
	return s
}

func Jsonify(data interface{}) string {
	bt, _ := json.Marshal(data)
	return string(bt)
}
