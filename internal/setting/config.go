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

type PreferredMonthOfYear uint8

const (
	CurrentMonth PreferredMonthOfYear = iota
	FirstMonth
	LastMonth
)

type Configuration struct {
	// Date order
	DateOrder string

	// Incomplete dates
	CurrentTime          time.Time
	DefaultTimezone      *time.Location
	PreferredDayOfMonth  PreferredDayOfMonth
	PreferredMonthOfYear PreferredMonthOfYear
	PreferredDateSource  PreferredDateSource
	StrictParsing        bool
	RequiredParts        []string

	// Language detection
	SkipTokens       []string
	DefaultLanguages []string

	// Others
	ReturnTimeAsPeriod bool
}

func (c Configuration) Clone() *Configuration {
	return &Configuration{
		DateOrder:            c.DateOrder,
		CurrentTime:          c.CurrentTime,
		DefaultTimezone:      c.DefaultTimezone,
		PreferredDayOfMonth:  c.PreferredDayOfMonth,
		PreferredMonthOfYear: c.PreferredMonthOfYear,
		PreferredDateSource:  c.PreferredDateSource,
		StrictParsing:        c.StrictParsing,
		RequiredParts:        append([]string{}, c.RequiredParts...),
		SkipTokens:           append([]string{}, c.SkipTokens...),
		DefaultLanguages:     append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:   c.ReturnTimeAsPeriod,
	}
}
