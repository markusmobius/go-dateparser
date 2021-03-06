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
	DateOrder string

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

func (c Configuration) Clone() *Configuration {
	return &Configuration{
		DateOrder:           c.DateOrder,
		CurrentTime:         c.CurrentTime,
		PreferredDayOfMonth: c.PreferredDayOfMonth,
		PreferredDateSource: c.PreferredDateSource,
		StrictParsing:       c.StrictParsing,
		RequiredParts:       append([]string{}, c.RequiredParts...),
		SkipTokens:          append([]string{}, c.SkipTokens...),
		DefaultLanguages:    append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:  c.ReturnTimeAsPeriod,
	}
}
