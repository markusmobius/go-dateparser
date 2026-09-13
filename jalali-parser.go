package dateparser

import (
	"fmt"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/jalali"
)

// ParseJalali converts an absolute Jalali (Persian or Solar Hijri) date to a
// Gregorian date. It defaults to MDY order and does not parse relative dates.
func ParseJalali(cfg *Configuration, str string) (date.Date, error) {
	// Prepare config
	if cfg == nil {
		cfg = &Configuration{}
	}

	dateOrder := "MDY"
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
	return jalali.Parse(iCfg, str)
}
