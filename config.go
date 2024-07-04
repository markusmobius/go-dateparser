package dateparser

import (
	"fmt"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

// DateOrder is function that returns date order string for specified language and/or locale.
// The returned date order MUST only uses characters M, D or Y which represents month, day and year.
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

// PreferredDateSource is the variable to set date source to fill incomplete dates value.
type PreferredDateSource uint8

const (
	// CurrentPeriod means parser will use current period to fill the incomplete dates value.
	// So, for current year 2021 and date string "10 December", parser will return 2021-12-10.
	CurrentPeriod PreferredDateSource = iota
	// Past means parser will use past period to fill the incomplete dates value. So, for current
	// year 2021 and date string "10 December", parser will return 2020-12-10.
	Past
	// Future means parser will use future period to fill the incomplete dates value. So, for
	// current year 2021 and date string "10 December", parser will return 2022-12-10.
	Future
)

// PreferredDayOfMonth is the variable to set day value for date that has month and year, but
// missing the day value. For example, date like "2021-12" or "February 2000".
type PreferredDayOfMonth uint8

const (
	// Current means parser will use current day to fill the day value. So, if today is
	// 2022-01-20 and date string is "February 2020", parser will return 2020-02-20.
	Current PreferredDayOfMonth = iota
	// First means parser will use first day of the month to fill the day value. So, if
	// date string is "February 2020", parser will return 2020-02-01.
	First
	// Last means parser will use last day of the month to fill the day value. So, if
	// date string is "February 2020", parser will return 2020-02-29.
	Last
)

// PreferredMonthOfYear is the variable to set month value for date that has year, but
// missing the month value. For example, date like "2021".
type PreferredMonthOfYear uint8

const (
	// CurrentMonth means parser will use current month to fill the month value. So,
	// if today is 2022-05-20 and date string is "2020", parser will return 2020-05-20.
	CurrentMonth PreferredMonthOfYear = iota
	// FirstMonth means parser will use January to fill the month value. So, if date
	// string is "2020", parser will return 2020-01-20.
	FirstMonth
	// LastMonth means parser will use December to fill the month value. So, if date
	// string is "2020", parser will return 2020-12-20.
	LastMonth
)

// Configuration is object to control and configure parsing behavior of date parser.
type Configuration struct {
	// Locales is a list of locale codes, e.g. ['fr-PF', 'qu-EC', 'af-NA'].
	// The parser uses only these locales to translate date string. When specified,
	// Languages and Region will be ignored.
	Locales []string

	// Languages is a list of language codes, e.g. ['en', 'es', 'zh-Hant']. If
	// locales are not given, languages and region are used to construct locales
	// for translation.
	Languages []string

	// Region is a region code, e.g. 'IN', '001', 'NE'. If locales are not given,
	// languages and region are used to construct locales for translation.
	Region string

	// If true, locales previously used to translate date are tried first.
	TryPreviousLocales bool

	// If true, locales are tried for translation of date string in the order in
	// which they are given.
	UseGivenOrder bool

	// Default languages is a list of language codes in ISO 639 (e.g. "en", "fr") that will be
	// used as default languages for parsing when language detection fails. When using this
	// setting, these languages will be tried after trying with the detected languages with no
	// success. It is especially useful when using the `DetectLanguagesFunction`.
	DefaultLanguages []string

	// DateOrder is function that specifies the order in which date components year, month
	// and day are expected while parsing ambiguous dates. If empty, parser will use each
	// language specific date order.
	DateOrder DateOrder

	// CurrentTime is the base datetime to use for interpreting partial or relative date
	// strings. Defaults to the current date and time in UTC.
	CurrentTime time.Time

	// DefaultTimezone is the default timezone to use when string doesn't contains any timezone.
	DefaultTimezone *time.Location

	// PreferredDayOfMonth specify the day for string with missing day. Defaults to `Current`.
	PreferredDayOfMonth PreferredDayOfMonth

	// PreferredMonthOfYear specify the month for string with missing month. Defaults to `CurrentMonth`.
	PreferredMonthOfYear PreferredMonthOfYear

	// PreferredDateSource specify the date source to fill incomplete date values. Defaults
	// to `CurrentPeriod`.
	PreferredDateSource PreferredDateSource

	// StrictParsing when set to true will make the parser returns a date only if the date
	// is complete, i.e. has day, month and year value. Defaults to false.
	StrictParsing bool

	// RequiredParts is list of date components that required by the parser. Defaults to
	// `nil` and can accept "day", "month" and "year".
	RequiredParts []string

	// SkipTokens is a list of tokens to discard while detecting language. Defaults to
	// []string{"t"} which skips T in iso format datetime string e.g. 2015-05-02T10:20:19+0000.
	SkipTokens []string

	// ReturnTimeAsPeriod returns `Time` as period in date object, if time component is present
	// in date string. Defaults to false.
	ReturnTimeAsPeriod bool
}

// Clone clones the config to a new, separate one.
func (c Configuration) Clone() *Configuration {
	return &Configuration{
		Locales:              append([]string{}, c.Locales...),
		Languages:            append([]string{}, c.Languages...),
		Region:               c.Region,
		TryPreviousLocales:   c.TryPreviousLocales,
		UseGivenOrder:        c.UseGivenOrder,
		DefaultLanguages:     append([]string{}, c.DefaultLanguages...),
		DateOrder:            c.DateOrder,
		CurrentTime:          c.CurrentTime,
		DefaultTimezone:      c.DefaultTimezone,
		PreferredDayOfMonth:  c.PreferredDayOfMonth,
		PreferredMonthOfYear: c.PreferredMonthOfYear,
		PreferredDateSource:  c.PreferredDateSource,
		StrictParsing:        c.StrictParsing,
		RequiredParts:        append([]string{}, c.RequiredParts...),
		SkipTokens:           append([]string{}, c.SkipTokens...),
		ReturnTimeAsPeriod:   c.ReturnTimeAsPeriod,
	}
}

// validate validates the configuration and return error if it's not valid.
func (c Configuration) validate() error {
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

	return c
}

func (c Configuration) toInternalConfig() *setting.Configuration {
	return &setting.Configuration{
		CurrentTime:          c.CurrentTime,
		DefaultTimezone:      c.DefaultTimezone,
		PreferredDayOfMonth:  setting.PreferredDayOfMonth(c.PreferredDayOfMonth),
		PreferredMonthOfYear: setting.PreferredMonthOfYear(c.PreferredMonthOfYear),
		PreferredDateSource:  setting.PreferredDateSource(c.PreferredDateSource),
		StrictParsing:        c.StrictParsing,
		RequiredParts:        append([]string{}, c.RequiredParts...),
		SkipTokens:           append([]string{}, c.SkipTokens...),
		DefaultLanguages:     append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:   c.ReturnTimeAsPeriod,
	}
}

func configFromInternal(c *setting.Configuration) *Configuration {
	return &Configuration{
		CurrentTime:          c.CurrentTime,
		DefaultTimezone:      c.DefaultTimezone,
		PreferredDayOfMonth:  PreferredDayOfMonth(c.PreferredDayOfMonth),
		PreferredMonthOfYear: PreferredMonthOfYear(c.PreferredMonthOfYear),
		PreferredDateSource:  PreferredDateSource(c.PreferredDateSource),
		StrictParsing:        c.StrictParsing,
		RequiredParts:        append([]string{}, c.RequiredParts...),
		SkipTokens:           append([]string{}, c.SkipTokens...),
		DefaultLanguages:     append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:   c.ReturnTimeAsPeriod,
	}
}
