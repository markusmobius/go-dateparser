package language

import (
	"errors"
	"regexp"

	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/pemistahl/lingua-go"
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
	rxKeepToken1    = regexp.MustCompile(`^.*[^\W_].*$`)
	rxKeepToken2    = regexp.MustCompile(`^.*[^\P{L}\d_].*$`)
	rxRegionRemover = regexp.MustCompile(`-([A-Z]+)$`)
	rxNumeral       = regexp.MustCompile(`(\d+)`)
	rxNumberOnly    = regexp.MustCompile(`^\d+$`)

	rxSentenceSplitters = map[int]*regexp.Regexp{
		// The most common splitter, used in European, Tagalog, Hebrew, Georgian, Indonesian, Vietnamese
		1: regexp.MustCompile(`([^\s.]*)[\.!?;…\r\n]+(?:\s|$)*`),
		// Splitter for Spanish
		2: regexp.MustCompile(`([^\s.]*)(?:[¡¿]+|[\.!?;…\r\n]+(?:\s|$))+`),
		// Splitter for Hindi and Bangla
		3: regexp.MustCompile(`([^\s.]*)[|!?;\r\n]+(?:\s|$)+`),
		// Splitter for Japanese and Chinese
		4: regexp.MustCompile(`([^\s.]*)[。…‥\.!?？！;\r\n]+(?:\s|$)+`),
		// Splitter for Thai
		5: regexp.MustCompile(`([^\s.]*)[\r\n]+`),
		// Splitter for Arabic and Farsi
		6: regexp.MustCompile(`([^\s.]*)[\r\n؟!\.…]+(?:\s|$)+`),
	}

	alwaysKeptTokens   = strutil.NewDict("+", ":", ".", " ", "-", "/")
	freshnessWords     = strutil.NewDict("day", "week", "month", "year", "hour", "minute", "second")
	langWithDigitAbbrs = strutil.NewDict("fi", "cs", "hu", "de", "da")
	commonCharset      = rangetable.New([]rune(`0123456789()\-/.,:' `)...)

	externalLanguageDetector = lingua.NewLanguageDetectorBuilder().FromAllSpokenLanguages().Build()
)
