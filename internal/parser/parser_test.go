package parser

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func Test_parseTimestamp(t *testing.T) {
	cfg := &setting.DefaultConfig
	zero := time.Time{}

	// Test milliseconds timestamp
	assert.Equal(t,
		time.Unix(1570308760, 263*1_000_000),
		parseTimestamp(cfg, "1570308760263"),
	)

	// Test microseconds timestamp
	assert.Equal(t,
		time.Unix(1570308760, 263_111*1_000),
		parseTimestamp(cfg, "1570308760263111"),
	)

	// Test wrong timestamp
	assert.Equal(t, zero, parseTimestamp(cfg, "15703087602631"))
	assert.Equal(t, zero, parseTimestamp(cfg, "157030876026xx"))
	assert.Equal(t, zero, parseTimestamp(cfg, "1570308760263x"))
	assert.Equal(t, zero, parseTimestamp(cfg, "157030876026311"))
	assert.Equal(t, zero, parseTimestamp(cfg, "15703087602631x"))
	assert.Equal(t, zero, parseTimestamp(cfg, "15703087602631xx"))
	assert.Equal(t, zero, parseTimestamp(cfg, "15703087602631111"))
	assert.Equal(t, zero, parseTimestamp(cfg, "1570308760263111x"))
	assert.Equal(t, zero, parseTimestamp(cfg, "1570308760263111xx"))
	assert.Equal(t, zero, parseTimestamp(cfg, "1570308760263111222"))
}

func Test_parseWithFormats(t *testing.T) {
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
	dt := parseWithFormats(&cfg, "yesterday", "2006-01-02")
	dateIsNotParsed(dt)

	// Matching format, should be parsed
	dt = parseWithFormats(&cfg, "25-03-14", "02-01-06")
	dateAsExpected(dt, "2014-03-25")

	// Should use current year for dates without year
	dt = parseWithFormats(&cfg, "09.16", "01.02")
	dateAsExpected(dt, "2015-09-16")

	// Should use day from config for dates without day
	str, format := "August 2014", "January 2006"

	cfg.PreferredDayOfMonth = "first"
	dt = parseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-01")

	cfg.PreferredDayOfMonth = "last"
	dt = parseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-31")

	cfg.PreferredDayOfMonth = "current"
	dt = parseWithFormats(&cfg, str, format)
	dateAsExpected(dt, "2014-08-04")

}
