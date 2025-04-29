package common

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	// Helper function
	isEqual := func(s string, hour, min, sec, msec int, expectedPeriod date.Period) {
		expectedTime := time.Date(0, 1, 1, hour, min, sec, msec*1_000, time.UTC)
		result, period, _ := ParseTime(s)
		assert.Equal(t, expectedTime, result, s)
		assert.Equal(t, expectedPeriod, period, s)
	}

	isZero := func(s string) {
		result, period, _ := ParseTime(s)
		assert.True(t, result.IsZero(), s)
		assert.Equal(t, date.None, period)
	}

	// Time should be parsed
	isEqual("11:30:14", 11, 30, 14, 0, date.Second)
	isEqual("11:30", 11, 30, 0, 0, date.Minute)
	isEqual("11:30 PM", 23, 30, 0, 0, date.Minute)
	isEqual("13:30 PM", 13, 30, 0, 0, date.Minute)
	isEqual("16:14 AM", 16, 14, 0, 0, date.Minute)
	isEqual("23:30 AM", 23, 30, 0, 0, date.Minute)
	isEqual("1:30 AM", 1, 30, 0, 0, date.Minute)
	isEqual("1:30:15.330 AM", 1, 30, 15, 330000, date.Second)
	isEqual("1:30:15.330 PM", 13, 30, 15, 330000, date.Second)
	isEqual("1:30:15.3301 PM", 13, 30, 15, 330100, date.Second)
	isEqual("11:20:05 AM", 11, 20, 5, 0, date.Second)
	isEqual("14:30:15.330100", 14, 30, 15, 330100, date.Second)
	isEqual("11 AM", 11, 0, 0, 0, date.Hour)
	isEqual("11 PM", 23, 0, 0, 0, date.Hour)

	// Time is invalid
	isZero("11")
	isZero("22:12:12 PM")
	isZero("22:12:10:16")
	isZero("10:14.123 PM")
	isZero("2:13:88")
	isZero("23:01:56.34 PM")
	isZero("2.45 PM")
}
