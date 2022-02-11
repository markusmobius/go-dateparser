package dateparser

import (
	"fmt"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/jalali"
)

// ParseJalali parses a Jalali date (which also called Persian or Solar
// Hijri date) which commonly used in Iran and Afghanistan.
func ParseJalali(cfg *Configuration, str string) (date.Date, error) {
	// Prepare config
	if cfg == nil {
		cfg = &Configuration{}
	}

	// Iran and Afghanistan mostly uses DMY format,
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
	return jalali.Parse(iCfg, str)
}
