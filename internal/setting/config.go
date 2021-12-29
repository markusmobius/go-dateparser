package setting

import "time"

type Configuration struct {
	// Timezone related
	Timezone              *time.Location
	ToTimezone            *time.Location
	ReturnAsTimezoneAware bool

	// Incomplete dates
	CurrentTime         time.Time
	RelativeTimeBase    time.Time
	PreferredDayOfMonth string

	// Language detection
	SkipTokens       []string
	DefaultLanguages []string

	// Others
	ReturnTimeAsPeriod bool
}

var DefaultConfig = Configuration{
	// Timezone related
	Timezone:              time.Local,
	ToTimezone:            nil,
	ReturnAsTimezoneAware: false,

	// Incomplete dates
	CurrentTime:         time.Now(),
	PreferredDayOfMonth: "current",

	// Language detection
	SkipTokens: []string{"t"},
}
