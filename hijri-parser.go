package dateparser

import (
	"fmt"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/hijri"
)

// ParseHijri parses a Hijri date string using Umm al-Qura Calendar
// which commonly used in Islamic country.
func ParseHijri(cfg *Configuration, str string) (date.Date, error) {
	// Prepare config
	if cfg == nil {
		cfg = &Configuration{}
	}

	// Hijri usually used in Eastern country which uses DMY,
	// so here it used as default.
	dateOrder := "DMY"
	if cfg.DateOrder != nil {
		do := cfg.DateOrder("")
		if do, valid := validateDateOrder(do); valid {
			dateOrder = do
		}
	}

	cfg = cfg.initiate()
	err := cfg.validate()
	if err != nil {
		return date.Date{}, fmt.Errorf("config error: %w", err)
	}

	// Start parser
	iCfg := cfg.toInternalConfig()
	iCfg.DateOrder = dateOrder
	return hijri.Parse(iCfg, str)
}
