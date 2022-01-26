package dateparser

import (
	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/hijri"
)

// ParseHijri parses a Hijri date string using Umm al-Qura Calendar.
func ParseHijri(cfg *Configuration, str string) (date.Date, error) {
	iCfg := cfg.toInternalConfig()
	return hijri.Parse(&iCfg, str)
}
