package dateparser_test

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/stretchr/testify/assert"
)

func tt(Y, M, D int, times ...int) time.Time {
	nTimes := len(times)
	var H, m, s, ms int

	if nTimes >= 1 {
		H = times[0]
	}
	if nTimes >= 2 {
		m = times[1]
	}
	if nTimes >= 3 {
		s = times[2]
	}
	if nTimes >= 4 {
		ms = times[3]
	}

	return time.Date(Y, time.Month(M), D, H, m, s, ms*1000, time.UTC)
}

func assertParseResult(t *testing.T, dt date.Date, expectedTime time.Time, expectedPeriod date.Period, messages ...string) bool {
	var message string
	if len(messages) > 0 {
		message = messages[0]
	}

	passed := assert.False(t, dt.IsZero(), message)
	if passed {
		diff := expectedTime.Sub(dt.Time)
		passed = assert.Zero(t, diff, message)
	}

	if passed {
		passed = assert.Equal(t, expectedPeriod, dt.Period, message)
	}

	return passed
}
