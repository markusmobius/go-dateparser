package language

import (
	"errors"
	"regexp"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	ErrNotFound           = errors.New("locale not found")
	ErrUnknownLocales     = errors.New("unknown locale(s)")
	ErrUnknownLanguages   = errors.New("unknown language(s)")
	ErrConflictingLocales = errors.New("locales should not have same language and different region")
)

const (
	splitSeparator = `|#=#=#|`
)

var (
	rxKeepToken1       = regexp.MustCompile(`^.*[^\W_].*$`)
	rxKeepToken2       = regexp.MustCompile(`^.*[^\P{L}\d_].*$`)
	rxRegionRemover    = regexp.MustCompile(`-([A-Z]+)$`)
	rxNumeral          = regexp.MustCompile(`(\d+)`)
	unicodeTransformer = transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)

	alwaysKeptTokens = map[string]struct{}{
		"+": {},
		":": {},
		".": {},
		" ": {},
		"-": {},
		"/": {},
	}

	freshnessWords = map[string]struct{}{
		"day":    {},
		"week":   {},
		"month":  {},
		"year":   {},
		"hour":   {},
		"minute": {},
		"second": {},
	}
)
