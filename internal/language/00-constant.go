package language

import (
	"errors"
	"unicode"

	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"golang.org/x/text/unicode/rangetable"
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
	rxSpaces     = regexp.MustCompile(`\s{2,}`)
	rxKeepToken1 = regexp.MustCompile(`^.*[^\W_].*$`)
	rxKeepToken2 = regexp.MustCompile(`^.*[^\P{L}\d_].*$`)
	rxNumeral    = regexp.MustCompile(`(\d+)`)
	rxTokenDigit = regexp.MustCompile(`[\d\.:\-/]+`)

	disallowedKnownWordRunes = rangetable.Merge(
		unicode.Letter,
		unicode.Mark,
		unicode.Number,
	)

	rxSentenceSplitters = map[int]*regexp.Regexp{
		// The most common splitter, used in European, Tagalog, Hebrew, Georgian, Indonesian, Vietnamese
		1: regexp.MustCompile(`([^\s.]*)[\.!?;…\r\n]+(?:\s|$)*`),
		// Splitter for Spanish
		2: regexp.MustCompile(`([^\s.]*)[\.!?;…\r\n]+(\s*[¡¿]*|$)|[¡¿]+`),
		// Splitter for Hindi and Bangla
		3: regexp.MustCompile(`([^\s.]*)[|!?;\r\n]+(?:\s|$)+`),
		// Splitter for Japanese and Chinese
		4: regexp.MustCompile(`([^\s.]*)[。…‥\.!?？！;\r\n]+(?:\s|$)+`),
		// Splitter for Thai
		5: regexp.MustCompile(`([^\s.]*)[\r\n]+`),
		// Splitter for Arabic and Farsi
		6: regexp.MustCompile(`([^\s.]*)[\r\n؟!\.…]+(?:\s|$)+`),
	}

	dashes             = strutil.NewDict("-", "——", "—", "～")
	alwaysKeptTokens   = strutil.NewDict("+", ":", ".", " ", "-", "/")
	freshnessWords     = strutil.NewDict("day", "week", "month", "year", "hour", "minute", "second")
	langWithDigitAbbrs = strutil.NewDict("fi", "cs", "hu", "de", "da")
	commonCharset      = rangetable.New([]rune(`0123456789()\-/.,:' `)...)
)
