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

func Test_parseTime(t *testing.T) {
	// Helper function
	isEqual := func(s string, hour, min, sec, msec int) {
		expected := time.Date(0, 1, 1, hour, min, sec, msec*1_000, time.UTC)
		result, _ := parseTime(s)
		assert.Equal(t, expected, result, s)
	}

	isZero := func(s string) {
		result, _ := parseTime(s)
		assert.True(t, result.IsZero(), s)
	}

	// Time should be parsed
	isEqual("11:30:14", 11, 30, 14, 0)
	isEqual("11:30", 11, 30, 0, 0)
	isEqual("11:30 PM", 23, 30, 0, 0)
	isEqual("13:30 PM", 13, 30, 0, 0)
	isEqual("16:14 AM", 16, 14, 0, 0)
	isEqual("23:30 AM", 23, 30, 0, 0)
	isEqual("1:30 AM", 1, 30, 0, 0)
	isEqual("1:30:15.330 AM", 1, 30, 15, 330000)
	isEqual("1:30:15.330 PM", 13, 30, 15, 330000)
	isEqual("1:30:15.3301 PM", 13, 30, 15, 330100)
	isEqual("11:20:05 AM", 11, 20, 5, 0)
	isEqual("14:30:15.330100", 14, 30, 15, 330100)

	// Time is invalid
	isZero("11")
	isZero("22:12:12 PM")
	isZero("22:12:10:16")
	isZero("10:14.123 PM")
	isZero("2:13:88")
	isZero("23:01:56.34 PM")
	isZero("2.45 PM")
}
