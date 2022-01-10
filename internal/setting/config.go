package setting

import "time"

type Configuration struct {
	// Date order
	DateOrder             string
	PreferLocaleDateOrder bool

	// Incomplete dates
	CurrentTime         time.Time
	PreferredDayOfMonth string
	PreferFutureTime    bool
	StrictParsing       bool
	RequiredParts       []string

	// Language detection
	SkipTokens       []string
	DefaultLanguages []string

	// Others
	ReturnTimeAsPeriod bool
}

var DefaultConfig = Configuration{
	// Date order
	DateOrder:             "MDY",
	PreferLocaleDateOrder: true,

	// Incomplete dates
	CurrentTime:         time.Now(),
	PreferredDayOfMonth: "current",

	// Language detection
	SkipTokens: []string{"t"},
}
