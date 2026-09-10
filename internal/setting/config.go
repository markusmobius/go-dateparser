package setting

import (
	"slices"
	"time"
)

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

type PreferredMonthOfYear uint8

const (
	CurrentMonth PreferredMonthOfYear = iota
	FirstMonth
	LastMonth
)

type Configuration struct {
	// Date order
	DateOrder           string
	DateOrderIsExplicit bool

	// Incomplete dates
	CurrentTime           time.Time
	DefaultTimezone       *time.Location
	PreferredDayOfMonth   PreferredDayOfMonth
	PreferredMonthOfYear  PreferredMonthOfYear
	PreferredDateSource   PreferredDateSource
	StrictParsing         bool
	IgnoreSurroundingText bool
	RequiredParts         []string

	// Language detection
	SkipTokens       []string
	DefaultLanguages []string

	// Others
	ReturnTimeAsPeriod bool
	PreserveEndOfMonth bool
}

func (c Configuration) Clone() *Configuration {
	return &Configuration{
		DateOrder:             c.DateOrder,
		DateOrderIsExplicit:   c.DateOrderIsExplicit,
		CurrentTime:           c.CurrentTime,
		DefaultTimezone:       c.DefaultTimezone,
		PreferredDayOfMonth:   c.PreferredDayOfMonth,
		PreferredMonthOfYear:  c.PreferredMonthOfYear,
		PreferredDateSource:   c.PreferredDateSource,
		StrictParsing:         c.StrictParsing,
		IgnoreSurroundingText: c.IgnoreSurroundingText,
		RequiredParts:         slices.Clone(c.RequiredParts),
		SkipTokens:            slices.Clone(c.SkipTokens),
		DefaultLanguages:      slices.Clone(c.DefaultLanguages),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
		PreserveEndOfMonth:    c.PreserveEndOfMonth,
	}
}
