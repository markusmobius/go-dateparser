package language

import (
	"strings"

	"golang.org/x/text/transform"
)

func normalizeUnicode(str string) string {
	normalized, _, err := transform.String(unicodeTransformer, str)
	if err != nil {
		return str
	}

	return normalized
}

func normalizeString(str string) string {
	str = normalizeUnicode(str)
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}
