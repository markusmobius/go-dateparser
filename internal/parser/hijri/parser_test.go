package hijri

import (
	"fmt"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text         string
		ExpectedTime time.Time
	}

	tests := []testScenario{
		{"14-09-1432 هـ, 09:40 صباحاً", tt(2011, 8, 14, 9, 40)},
		{"20-02-1430 هـ, 07:21 صباحاً", tt(2009, 2, 15, 7, 21)},
		{"11-08-1434 هـ, 09:38 صباحاً", tt(2013, 6, 20, 9, 38)},
		{" 17-01-1437 هـ 08:30 مساءً", tt(2015, 10, 30, 20, 30)},
		{"29-02-1433 هـ, 06:22 صباحاً", tt(2012, 1, 23, 6, 22)},
		{"04-03-1433 هـ, 10:08 مساءً", tt(2012, 1, 27, 22, 8)},
		{"30-02-1433", tt(2012, 1, 24)},
		// Handle two digit year
		{"30-02-33", tt(2012, 1, 24)},
		{"10-03-90", tt(1970, 5, 16)},
	}

	// Prepare config
	cfg := &setting.Configuration{
		DateOrder: "DMY",
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => \"%s\"", test.Text,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Parse text
		dt, _ := Parse(cfg, test.Text)
		if passed := assertResult(t, dt, test.ExpectedTime, date.Day, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParsePythonCalendarRules(t *testing.T) {
	cfg := &setting.Configuration{
		DateOrder:   "DMY",
		CurrentTime: tt(2024, 3, 20, 12, 0),
	}
	for _, test := range []struct {
		input string
		want  time.Time
	}{
		{"1/1/1343", tt(1924, 8, 1)},
		{"1/1/1355", tt(1936, 3, 24)},
		{"30/12/1356", tt(1938, 3, 2)},
		{"01/02", tt(2023, 8, 17)},
		{"23:30", tt(2024, 3, 20, 23, 30)},
	} {
		t.Run(test.input, func(t *testing.T) {
			parsed, err := Parse(cfg, test.input)
			if assert.NoError(t, err) {
				assert.True(t, parsed.Time.Equal(test.want), "got %s, want %s", parsed.Time, test.want)
			}
		})
	}
	for _, input := range []string{"1/1/1342", "1/1/1501", "30/12/9999"} {
		_, err := Parse(cfg, input)
		assert.Error(t, err, input)
	}
	for _, reference := range []time.Time{tt(1924, 7, 31), tt(2077, 11, 17)} {
		cfg.CurrentTime = reference
		_, err := Parse(cfg, "1/1/1444")
		assert.Error(t, err, reference)
	}
}

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

func assertResult(t *testing.T, dt date.Date, expectedTime time.Time, expectedPeriod date.Period, messages ...string) bool {
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
