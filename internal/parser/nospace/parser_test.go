package nospace

import (
	"fmt"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParse_notParsed(t *testing.T) {
	tests := []string{
		// Date with spaces
		"2013 25 12",
		// Empty string
		"",
		":",
		// Only has colons
		"::",
		":::",
		// Has alphabets
		"12AUG2015",
		// Insane date
		"12345678901234567890",
		"987654321234567890123456789",
	}

	nFailed := 0
	for _, test := range tests {
		_, err := Parse(nil, test)
		passed := assert.NotNil(t, err, test)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParse_parsed(t *testing.T) {
	// Prepare struct
	type testScenario struct {
		Text           string
		ExpectedTime   time.Time
		ExpectedPeriod date.Period
		Order          string
	}

	// Helper function
	tt := func(Y, M, D int, times ...int) time.Time {
		var H, m, s int
		switch len(times) {
		case 1:
			H = times[0]
		case 2:
			H = times[0]
			m = times[1]
		case 3:
			H = times[0]
			m = times[1]
			s = times[2]
		}
		return time.Date(Y, time.Month(M), D, H, m, s, 0, time.UTC)
	}

	// Prepare scenarios
	tests := []testScenario{
		// Order is specified
		{"201115", tt(2015, 11, 20), date.Day, "DMY"},
		{"18092017", tt(2017, 9, 18), date.Day, "DMY"},
		{"912958:15:10", tt(9581, 12, 9, 5, 1), date.Day, "DMY"},
		{"20201511", tt(2015, 11, 20), date.Day, "DYM"},
		{"171410", tt(2014, 10, 17), date.Day, "DYM"},
		{"71995121046", tt(1995, 12, 7, 10, 4, 6), date.Day, "DYM"},
		{"112015", tt(2015, 1, 1), date.Day, "MDY"},
		{"12595", tt(1995, 12, 5), date.Day, "MDY"},
		{"459712:15:07.54", tt(4597, 12, 15, 0, 7), date.Day, "MDY"},
		{"11012015", tt(2015, 11, 1), date.Day, "MDY"},
		{"12201511", tt(2015, 12, 11), date.Day, "MYD"},
		{"21813", tt(2018, 2, 13), date.Day, "MYD"},
		{"12937886", tt(2937, 1, 8, 8, 6), date.Day, "MYD"},
		{"20151211", tt(2015, 12, 11), date.Day, "YMD"},
		{"18216", tt(2018, 2, 16), date.Day, "YMD"},
		{"1986411:5", tt(1986, 4, 1, 1, 5), date.Day, "YMD"},
		{"20153011", tt(2015, 11, 30), date.Day, "YDM"},
		{"14271", tt(2014, 1, 27), date.Day, "YDM"},
		{"2010111110:11", tt(2010, 11, 11, 10, 1, 1), date.Day, "YDM"},
		{"10:11:2", tt(2010, 2, 11, 0, 0), date.Day, "YDM"},
		// Order is not specified
		{"10032017", tt(2017, 10, 3), date.Day, ""},
		{"19991215:07:08:04.54", tt(1999, 12, 15, 7, 8, 4), date.Day, ""},
		// 8 digit, order not specified
		{"20110101", tt(2011, 1, 1), date.Day, ""},
		{"01201702", tt(1702, 1, 20), date.Day, ""},
		{"01202020", tt(2020, 1, 20), date.Day, ""},
		{"20202001", tt(2020, 1, 20), date.Day, ""},
		{"20202020", tt(2002, 2, 2, 0, 2), date.Day, ""},
		{"12345678", tt(1234, 5, 6, 7, 8), date.Day, ""},
	}

	// Prepare configuration
	cfg := setting.DefaultConfig

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%s) => \"%s\"",
			test.Text, test.Order,
			test.ExpectedTime.Format("2006-01-02 15:04:05"))

		// Parse text
		cfg.DateOrder = test.Order
		dt, err := Parse(&cfg, test.Text)

		// Check the result
		passed := assert.Nil(t, err, message)
		if passed {
			passed = assert.Equal(t, test.ExpectedTime, dt.Time, message)
		}
		if passed {
			passed = assert.Equal(t, test.ExpectedPeriod, dt.Period, message)
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func Test_getPeriodFromFormat(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Format   string
		Expected date.Period
	}

	tests := []testScenario{
		{"d", date.Day},
		{"H", date.Day},
		{"M", date.Day},
		{"S", date.Day},
		{"m", date.Month},
		{"y", date.Year},
		{"", date.Year},
		{"md", date.Day},
		{"Ym", date.Month},
		{"dmy", date.Day},
		{"YmdHM", date.Day},
		{"YmdHMS.f", date.Day},
		{"HM", date.Day},
		{"MS.f", date.Day},
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		result := getPeriodFromFormat(test.Format)
		passed := assert.Equal(t, test.Expected, result, test.Format)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
