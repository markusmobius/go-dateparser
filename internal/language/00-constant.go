package language

import (
	"errors"
	"regexp"
)

var (
	rxRegionRemover = regexp.MustCompile(`-([A-Z]+)$`)

	ErrNotFound           = errors.New("locale not found")
	ErrUnknownLocales     = errors.New("unknown locale(s)")
	ErrUnknownLanguages   = errors.New("unknown language(s)")
	ErrConflictingLocales = errors.New("locales should not have same language and different region")
)
