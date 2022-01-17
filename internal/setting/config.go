package setting

import "time"

type PreferredDateSource uint8

const (
	CurrentPeriod PreferredDateSource = iota
	Past
	Future
)

type PreferredDayOfMonth uint8

const (
	Current PreferredDayOfMonth = iota
	First
	Last
)

type Configuration struct {
	// Date order
	DateOrder             string
	PreferLocaleDateOrder bool

	// Incomplete dates
	CurrentTime         time.Time
	PreferredDayOfMonth PreferredDayOfMonth
	PreferredDateSource PreferredDateSource
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
	PreferredDayOfMonth: Current,

	// Language detection
	SkipTokens: []string{"t"},
}
