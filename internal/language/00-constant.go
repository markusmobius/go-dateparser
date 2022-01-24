package language

import (
	"errors"
	"regexp"

	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/pemistahl/lingua-go"
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

	alwaysKeptTokens = strutil.NewDict("+", ":", ".", " ", "-", "/")
	freshnessWords   = strutil.NewDict("day", "week", "month", "year", "hour", "minute", "second")

	externalLanguageDetector = lingua.NewLanguageDetectorBuilder().FromAllSpokenLanguages().Build()
)
