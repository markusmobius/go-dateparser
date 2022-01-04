package setting

import "time"

type Configuration struct {
	// Incomplete dates
	CurrentTime         time.Time
	PreferredDayOfMonth string
	PreferFutureTime    bool

	// Language detection
	SkipTokens       []string
	DefaultLanguages []string

	// Others
	ReturnTimeAsPeriod bool
}

var DefaultConfig = Configuration{
	// Incomplete dates
	CurrentTime:         time.Now(),
	PreferredDayOfMonth: "current",

	// Language detection
	SkipTokens: []string{"t"},
}
