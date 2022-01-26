package dateparser

import (
	"fmt"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
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

// Configuration is object to control and configure parsing behavior of date parser.
type Configuration struct {
	// DateOrder specifies the order in which date components year, month and day are
	// expected while parsing ambiguous dates. It defaults to MDY which translates to
	// month first, day second and year last order. Characters M, D or Y can be shuffled
	// to meet required order. For example, DMY specifies day first, month second and
	// year last order.
	DateOrder string
	// PreferConfigDateOrder defaults to false. Most languages have a default date order
	// specified for them. For example, for French it is DMY. If this option set to true,
	// the language date order will be ignored in favor of the date order from config.
	PreferConfigDateOrder bool
	// CurrentTime is the base datetime to use for interpreting partial or relative date
	// strings. Defaults to the current date and time in UTC.
	CurrentTime time.Time
	// PreferredDayOfMonth specify the day for date with missing day. Defaults to `Current`.
	PreferredDayOfMonth PreferredDayOfMonth
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
	// Default languages is a list of language codes in ISO 639 (e.g. "en", "fr") that will be
	// used as default languages for parsing when language detection fails. When using this
	// setting, these languages will be tried after trying with the detected languages with no
	// success. It is especially useful when using the `DetectLanguagesFunction`.
	DefaultLanguages []string
	// ReturnTimeAsPeriod returns `Time` as period in date object, if time component is present
	// in date string. Defaults to false.
	ReturnTimeAsPeriod bool
}

func (c Configuration) Clone() Configuration {
	return Configuration{
		DateOrder:             c.DateOrder,
		PreferConfigDateOrder: c.PreferConfigDateOrder,
		CurrentTime:           c.CurrentTime,
		PreferredDayOfMonth:   c.PreferredDayOfMonth,
		PreferredDateSource:   c.PreferredDateSource,
		StrictParsing:         c.StrictParsing,
		RequiredParts:         append([]string{}, c.RequiredParts...),
		SkipTokens:            append([]string{}, c.SkipTokens...),
		DefaultLanguages:      append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
	}
}

// validate validates the configuration and return error if it's not valid.
func (c Configuration) validate() error {
	// Validate date order
	c.DateOrder = strings.ToUpper(c.DateOrder)
	for _, char := range c.DateOrder {
		switch char {
		case 'D', 'M', 'Y':
		default:
			return fmt.Errorf("invalid component in date order: %q", char)
		}
	}

	// Validate preferred day of month
	if dom := c.PreferredDayOfMonth; dom < 0 || dom > Last {
		return fmt.Errorf("invalid preferred day of month: %d", dom)
	}

	// Validate preferred date source
	if ds := c.PreferredDateSource; ds < 0 || ds > Future {
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
func (c *Configuration) initiate() {
	c.DateOrder = strings.ToUpper(c.DateOrder)
	if c.DateOrder == "" && c.PreferConfigDateOrder {
		c.DateOrder = "MDY"
	}

	if c.CurrentTime.IsZero() {
		c.CurrentTime = time.Now().UTC()
	}

	if len(c.SkipTokens) == 0 {
		c.SkipTokens = []string{"t"}
	}
}

func (c Configuration) toInternalConfig() setting.Configuration {
	return setting.Configuration{
		DateOrder:             c.DateOrder,
		PreferConfigDateOrder: c.PreferConfigDateOrder,
		CurrentTime:           c.CurrentTime,
		PreferredDayOfMonth:   setting.PreferredDayOfMonth(c.PreferredDayOfMonth),
		PreferredDateSource:   setting.PreferredDateSource(c.PreferredDateSource),
		StrictParsing:         c.StrictParsing,
		RequiredParts:         append([]string{}, c.RequiredParts...),
		SkipTokens:            append([]string{}, c.SkipTokens...),
		DefaultLanguages:      append([]string{}, c.DefaultLanguages...),
		ReturnTimeAsPeriod:    c.ReturnTimeAsPeriod,
	}
}
