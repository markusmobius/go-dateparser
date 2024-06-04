package jalali

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
		{"سه شنبه سوم شهریور ۱۳۹۴", tt(2015, 8, 25, 0, 0)},
		{"پنجشنبه چهارم تیر ۱۳۹۴", tt(2015, 6, 25, 0, 0)},
		{"شنبه یکم فروردین ۱۳۹۴", tt(2015, 3, 21, 0, 0)},
		{"یکشنبه شانزدهم آذر ۱۳۹۳", tt(2014, 12, 7, 0, 0)},
		{"دوشنبه دوازدهم آبان ۱۳۹۳", tt(2014, 11, 3, 0, 0)},
		{"یکشنبه سیزدهم مهر ۱۳۹۳", tt(2014, 10, 5, 0, 0)},
		{"دوشنبه دوم تیر ۱۳۹۳", tt(2014, 6, 23, 0, 0)},
		{"یکشنبه سوم فروردین ۱۳۹۳", tt(2014, 3, 23, 0, 0)},
		{"پنجشنبه دهم بهمن ۱۳۹۲", tt(2014, 1, 30, 0, 0)},
		{"شنبه چهاردهم دی ۱۳۹۲", tt(2014, 1, 4, 0, 0)},
		{"جمعه بیستم خرداد ۱۳۹۰", tt(2011, 6, 10, 0, 0)},
		{"شنبه نهم مرداد ۱۳۸۹", tt(2010, 7, 31, 0, 0)},
		{"پنجشنبه بیست و سوم اردیبهشت ۱۳۸۹", tt(2010, 5, 13, 0, 0)},
		{"جمعه سی ام اسفند ۱۳۸۷", tt(2009, 3, 20, 0, 0)},
		// dates with time component
		{"پنجشنبه ۲۶ شهريور ۱۳۹۴ ساعت ۹:۳۲", tt(2015, 9, 17, 9, 32)},
		{"دوشنبه ۲۳ شهريور ۱۳۹۴ ساعت ۱۹:۱۱", tt(2015, 9, 14, 19, 11)},
		{"جمعه سی ام اسفند ۱۳۸۷ساعت 19:47", tt(2009, 3, 20, 19, 47)},
		{"شنبه چهاردهم دی ۱۳۹۲ساعت 12:1", tt(2014, 1, 4, 12, 1)},
		{"پنجشنبه 26 شهریور 1394 ساعت ساعت 11 و 01 دقیقه و 47 ثانیه", tt(2015, 9, 17, 11, 1, 47)},
		{"پنجشنبه 26 شهریور 1394 ساعت ساعت 10 و 58 دقیقه و 04 ثانیه", tt(2015, 9, 17, 10, 58, 4)},
		{"سه شنبه 17 شهریور 1394 ساعت ساعت 18 و 21 دقیقه و 44 ثانیه", tt(2015, 9, 8, 18, 21, 44)},
		{"پنجشنبه 11 تیر 1394", tt(2015, 7, 2, 0, 0)},
		{"پنجشنبه 26 شهریور 1394 ساعت ساعت 11 و 01 دقیقه", tt(2015, 9, 17, 11, 1)},
		// Handle two digit year
		{"پنجشنبه 26 شهریور 94 ساعت ساعت 11 و 01 دقیقه", tt(2015, 9, 17, 11, 1)},
		{"چهارشنبه 22 شهریور 02 ساعت ساعت 11 و 01 دقیقه", tt(2023, 9, 13, 11, 1)},
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
