package parser

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParseWithFormats(t *testing.T) {
	// Helper function
	dateIsParsed := func(dt DateData) { assert.NotZero(t, dt.Time) }
	dateIsNotParsed := func(dt DateData) { assert.Zero(t, dt.Time) }
	dateAsExpected := func(dt DateData, expected string, expectedFormat ...string) {
		dateIsParsed(dt)

		format := "2006-01-02"
		if len(expectedFormat) > 0 {
			format = expectedFormat[0]
		}

		str := dt.Time.Format(format)
		assert.Equal(t, expected, str)
	}

	// Prepare configuration
	cfg := setting.DefaultConfig
	cfg.CurrentTime = time.Date(2015, 2, 4, 0, 0, 0, 0, time.Local)

	// No matching format, shouldn't be parsed
	dt := ParseWithFormats(&cfg, "yesterday", "2006-01-02")
	dateIsNotParsed(dt)

	// Matching format, should be parsed
	dt = ParseWithFormats(&cfg, "25-03-14", "02-01-06")
	dateAsExpected(dt, "2014-03-25")

	// Should use current year for dates without year
	dt = ParseWithFormats(&cfg, "09.16", "01.02")
	dateAsExpected(dt, "2015-09-16")

	// Should use day from config for dates without day
	str, format := "August 2014", "January 2006"

	cfg.PreferredDayOfMonth = "first"
	dt = ParseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-01")

	cfg.PreferredDayOfMonth = "last"
	dt = ParseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-31")

	cfg.PreferredDayOfMonth = "current"
	dt = ParseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-04")

}
