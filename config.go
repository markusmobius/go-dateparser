package dateparser

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

// DateOrder returns the component order for a language or locale.
// The result must contain M, D, and Y exactly once, representing month, day, and year.
type DateOrder func(locale string) string

var (
	YMD = func(_ string) string { return "YMD" }
	YDM = func(_ string) string { return "YDM" }
	MYD = func(_ string) string { return "MYD" }
	MDY = func(_ string) string { return "MDY" }
	DYM = func(_ string) string { return "DYM" }
	DMY = func(_ string) string { return "DMY" }

	DefaultDateOrder = func(locale string) string {
		if ld, exist := data.GetLocaleData(locale); exist {
			return ld.DateOrder
		} else {
			return "YMD"
		}
	}
)

// PreferredDateSource selects how incomplete dates are interpreted relative to CurrentTime.
type PreferredDateSource uint8

const (
	// CurrentPeriod fills missing components from the current calendar period.
	CurrentPeriod PreferredDateSource = iota
	// Past prefers an interpretation before CurrentTime.
	Past
	// Future prefers an interpretation after CurrentTime.
	Future
)

// PreferredDayOfMonth selects the day for a date with no explicit day,
// such as "2021-12" or "February 2000".
type PreferredDayOfMonth uint8

const (
	// Current uses the day from CurrentTime, clamped to the target month's last day.
	Current PreferredDayOfMonth = iota
	// First uses the first day of the target month.
	First
	// Last uses the last day of the target month.
	Last
)

// PreferredMonthOfYear selects the month for a date with no explicit month,
// such as "2021".
type PreferredMonthOfYear uint8

const (
	// CurrentMonth uses the month from CurrentTime.
	CurrentMonth PreferredMonthOfYear = iota
	// FirstMonth uses January.
	FirstMonth
	// LastMonth uses December.
	LastMonth
)

// Configuration controls date parsing and search behavior.
// Calls clone the configuration; callers must not mutate it or its slices during a call.
type Configuration struct {
	// Locales restricts translation to locale codes such as "fr-PF", "qu-EC", or "af-NA".
	// When specified, Languages and Region are ignored.
	Locales []string

	// Languages restricts translation to language codes such as "en", "es", or "zh-Hant".
	// When Locales is empty, Languages and Region select the translation locales.
	Languages []string

	// Region is a region code such as "IN", "001", or "NE" used with Languages
	// when Locales is empty.
	Region string

	// TryPreviousLocales tries previously successful translation locales first.
	TryPreviousLocales bool

	// UseGivenOrder tries locales in the supplied order.
	UseGivenOrder bool

	// DefaultLanguages supplies fallback language codes, such as "en" or "fr", when
	// parsing with detected languages fails. It is useful with DetectLanguagesFunction.
	DefaultLanguages []string

	// DateOrder specifies the order of year, month, and day in ambiguous dates.
	// If nil, each locale's date order is used.
	DateOrder DateOrder

	// CurrentTime is the base datetime to use for interpreting partial or relative date
	// strings. Defaults to the current date and time in UTC.
	CurrentTime time.Time

	// DefaultTimezone is used when the input contains no timezone.
	// If nil, the location of CurrentTime is used. Unix timestamps retain time.Local.
	DefaultTimezone *time.Location

	// PreferredDayOfMonth selects a missing day. Defaults to Current.
	PreferredDayOfMonth PreferredDayOfMonth

	// PreferredMonthOfYear selects a missing month. Defaults to CurrentMonth.
	PreferredMonthOfYear PreferredMonthOfYear

	// PreferredDateSource selects the interpretation of incomplete dates.
	// Defaults to CurrentPeriod.
	PreferredDateSource PreferredDateSource

	// StrictParsing requires an explicit day, month, and year. Defaults to false.
	StrictParsing bool

	// IgnoreSurroundingText retries failed parses without unknown leading or trailing tokens.
	IgnoreSurroundingText bool

	// RequiredParts lists required date components: "day", "month", and/or "year".
	// Defaults to nil.
	RequiredParts []string

	// SkipTokens lists tokens to discard during translation and language detection.
	// An empty list defaults to []string{"t"}, skipping the T in ISO date-time strings.
	SkipTokens []string

	// ReturnTimeAsPeriod permits Second, Minute, or Hour precision in Date.Period
	// when the input specifies a time. Otherwise time precision is reported as Day.
	// Defaults to false.
	ReturnTimeAsPeriod bool

	// SearchStrategy selects "split" (default) or longest-token "ngram" search.
	SearchStrategy string

	// ReturnTimeSpan appends the first matching English time span's start and end
	// to search results. Boundaries retain the reference time of day.
	ReturnTimeSpan bool

	// DefaultStartOfWeek is "monday" (default) or "sunday" for time-span searches.
	DefaultStartOfWeek string

	// DefaultDaysInMonth is the length of an unnumbered month span; zero defaults to 30.
	DefaultDaysInMonth int

	// PreserveEndOfMonth is retained for source compatibility.
	// Deprecated: month-end clamping now always follows Python dateutil.
	PreserveEndOfMonth bool
}

// Clone copies the configuration and its slices.
func (c Configuration) Clone() *Configuration {
	return &Configuration{
		Locales:               slices.Clone(c.Locales),
		Languages:             slices.Clone(c.Languages),
		Region:                c.Region,
		TryPreviousLocales:    c.TryPreviousLocales,
		UseGivenOrder:         c.UseGivenOrder,
		DefaultLanguages:      slices.Clone(c.DefaultLanguages),
		DateOrder:             c.DateOrder,
		CurrentTime:           c.CurrentTime,
		DefaultTimezone:       c.DefaultTimezone,
		PreferredDayOfMonth:   c.PreferredDayOfMonth,
		PreferredMonthOfYear:  c.PreferredMonthOfYear,
		PreferredDateSource:   c.PreferredDateSource,
		StrictParsing:         c.StrictParsing,
		IgnoreSurroundingText: c.IgnoreSurroundingText,
		RequiredParts:         slices.Clone(c.RequiredParts),
		SkipTokens:            slices.Clone(c.SkipTokens),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
		SearchStrategy:        c.SearchStrategy,
		ReturnTimeSpan:        c.ReturnTimeSpan,
		DefaultStartOfWeek:    c.DefaultStartOfWeek,
		DefaultDaysInMonth:    c.DefaultDaysInMonth,
		PreserveEndOfMonth:    c.PreserveEndOfMonth,
	}
}

// validate validates the configuration and return error if it's not valid.
func (c Configuration) validate() error {
	if c.SearchStrategy != "" && c.SearchStrategy != "split" && c.SearchStrategy != "ngram" {
		return fmt.Errorf("invalid search strategy: %s", c.SearchStrategy)
	}
	if c.DefaultStartOfWeek != "" && c.DefaultStartOfWeek != "monday" && c.DefaultStartOfWeek != "sunday" {
		return fmt.Errorf("invalid default start of week: %s", c.DefaultStartOfWeek)
	}

	// Validate preferred day of month
	if dom := c.PreferredDayOfMonth; dom > Last {
		return fmt.Errorf("invalid preferred day of month: %d", dom)
	}

	// Validate preferred month of year
	if moy := c.PreferredMonthOfYear; moy > LastMonth {
		return fmt.Errorf("invalid preferred month of year: %d", moy)
	}

	// Validate preferred date source
	if ds := c.PreferredDateSource; ds > Future {
		return fmt.Errorf("invalid preferred date source: %d", ds)
	}

	// Validate required parts
	for _, part := range c.RequiredParts {
		switch strings.ToLower(part) {
		case "day", "month", "year":
		default:
			return fmt.Errorf("invalid component in required parts: %s", part)
		}
	}

	return nil
}

// initiate normalize config value and apply the defaults.
func (c *Configuration) initiate() *Configuration {
	c = c.Clone()

	if c.CurrentTime.IsZero() {
		c.CurrentTime = time.Now().UTC()
	}

	if len(c.SkipTokens) == 0 {
		c.SkipTokens = []string{"t"}
	}

	if c.DefaultStartOfWeek == "" {
		c.DefaultStartOfWeek = "monday"
	}
	if c.DefaultDaysInMonth == 0 {
		c.DefaultDaysInMonth = 30
	}

	return c
}

func (c Configuration) toInternalConfig() *setting.Configuration {
	return &setting.Configuration{
		DateOrderIsExplicit:   c.DateOrder != nil,
		CurrentTime:           c.CurrentTime,
		DefaultTimezone:       c.DefaultTimezone,
		PreferredDayOfMonth:   setting.PreferredDayOfMonth(c.PreferredDayOfMonth),
		PreferredMonthOfYear:  setting.PreferredMonthOfYear(c.PreferredMonthOfYear),
		PreferredDateSource:   setting.PreferredDateSource(c.PreferredDateSource),
		StrictParsing:         c.StrictParsing,
		IgnoreSurroundingText: c.IgnoreSurroundingText,
		RequiredParts:         slices.Clone(c.RequiredParts),
		SkipTokens:            slices.Clone(c.SkipTokens),
		DefaultLanguages:      slices.Clone(c.DefaultLanguages),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
		PreserveEndOfMonth:    c.PreserveEndOfMonth,
	}
}

func configFromInternal(c *setting.Configuration) *Configuration {
	return &Configuration{
		CurrentTime:           c.CurrentTime,
		DefaultTimezone:       c.DefaultTimezone,
		PreferredDayOfMonth:   PreferredDayOfMonth(c.PreferredDayOfMonth),
		PreferredMonthOfYear:  PreferredMonthOfYear(c.PreferredMonthOfYear),
		PreferredDateSource:   PreferredDateSource(c.PreferredDateSource),
		StrictParsing:         c.StrictParsing,
		IgnoreSurroundingText: c.IgnoreSurroundingText,
		RequiredParts:         slices.Clone(c.RequiredParts),
		SkipTokens:            slices.Clone(c.SkipTokens),
		DefaultLanguages:      slices.Clone(c.DefaultLanguages),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
		PreserveEndOfMonth:    c.PreserveEndOfMonth,
	}
}
