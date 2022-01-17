package common

import (
	"fmt"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func CheckStrictParsing(cfg *setting.Configuration, missingParts strutil.Dict) error {
	// Read config
	var strictParsing bool
	var requiredParts []string
	if cfg != nil {
		strictParsing = cfg.StrictParsing
		requiredParts = cfg.RequiredParts
	}

	// If strict parsing enabled, make sure there are no missing parts
	if strictParsing && len(missingParts) > 0 {
		return fmt.Errorf("fields missing from date string: %s", missingParts)
	}

	// Check if there are any missing required parts
	for _, requiredPart := range requiredParts {
		if missingParts.Contain(requiredPart) {
			return fmt.Errorf("required part \"%s\" is missing", requiredPart)
		}
	}

	return nil
}
