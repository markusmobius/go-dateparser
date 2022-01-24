package timezone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopTzOffset_nothingExtracted(t *testing.T) {
	// Helper function
	emptyTz := func(str string) {
		_, tzData := PopTzOffset(str)
		assert.True(t, tzData.IsZero())
	}

	emptyTz("15 May 2004")
	emptyTz("Wed Aug 05 12:00:00 EDTERR 2015")
}

func TestPopTzOffset_getOffset(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text     string
		Expected float64
	}

	tests := []testScenario{
		{"Sep 03 2014 | 4:32 pm EDT", -4},
		{"17th October, 2034 @ 01:08 am PDT", -7},
		{"17th October, 2034 @ 01:08 am (PDT)", -7},
		{"October 17, 2014 at 7:30 am PST", -8},
		{"20 Oct 2014 13:08 CET", +1},
		{"20 Oct 2014 13:08cet", +1},
		{"Nov 25 2014 | 10:17 pm EST", -5},
		{"Nov 25 2014 | 10:17 pm +0600", +6},
		{"Nov 25 2014 | 10:17 pm -0930", -9.5},
		{"20 Oct 2014 | 05:17 am -1200", -12},
		{"20 Oct 2014 | 05:17 am +0000", 0},
		{"20 Oct 2014 | 05:17 am -0000", 0},
		{"Wed Aug 05 12:00:00 EDT 2015", -4},
		{"April 10, 2016 at 12:00:00 UTC", 0},
		{"April 10, 2016 at 12:00:00 MEZ", 1},
		{"April 10, 2016 at 12:00:00 MESZ", 2},
		{"April 10, 2016 at 12:00:00 GMT+2", 2},
		{"April 10, 2016 at 12:00:00 UTC+2:00", 2},
		{"April 10, 2016 at 12:00:00 GMT+02:00", 2},
		{"April 10, 2016 at 12:00:00 UTC+5:30", 5.5},
		{"April 10, 2016 at 12:00:00 GMT+05:30", 5.5},
		{"April 10, 2016 at 12:00:00 UTC-2", -2},
		{"April 10, 2016 at 12:00:00 GMT-2:00", -2},
		{"April 10, 2016 at 12:00:00 UTC-02:00", -2},
		{"April 10, 2016 at 12:00:00 GMT-9:30", -9.5},
		{"April 10, 2016 at 12:00:00 UTC-09:30", -9.5},
		{"Thu, 24 Nov 2016 16:03:00 UT", 0},
		{"Fri Sep 23 2016 10:34:51 GMT+0800 (CST)", 8},
		{"Fri Sep 23 2016 10:34:51 GMT+12", 12},
		{"Fri Sep 23 2016 10:34:51 UTC+13", 13},
		{"Fri Sep 23 2016 10:34:51 GMT+1245 (CST)", 12.75},
		{"Fri Sep 23 2016 10:34:51 UTC+1245", 12.75},
		{"2019-07-17T12:30:00.000-03:30", -3.5},
		{"2019-07-17T12:30:00.000-02:30", -2.5},
		{"16. srpna 2021 9:59:44 SELČ", 2},
		{"16. srpna 2021 9:59:44 SEČ", 1},
		{"16. srpna 2021 9:59:44 ZEČ", 0},
		{"16. srpna 2021 9:59:44 VEČ", 2},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => %.f", test.Text, test.Expected)

		// Check the offset
		_, tzData := PopTzOffset(test.Text)
		expected := int(test.Expected * 3_600)
		passed := assert.Equal(t, expected, tzData.Offset, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestPopTzOffset_getLeftoverString(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text     string
		Expected string
	}

	tests := []testScenario{
		// Valid timezone must be removed from string
		{"Sep 03 2014 | 4:32 pm EDT", "Sep 03 2014 | 4:32 pm "},
		{"17th October, 2034 @ 01:08 am PDT", "17th October, 2034 @ 01:08 am "},
		{"October 17, 2014 at 7:30 am PST", "October 17, 2014 at 7:30 am "},
		{"20 Oct 2014 13:08 CET", "20 Oct 2014 13:08 "},
		{"20 Oct 2014 13:08cet", "20 Oct 2014 13:08"},
		{"Nov 25 2014 | 10:17 pm EST", "Nov 25 2014 | 10:17 pm "},
		{"17th October, 2034 @ 01:08 am +0700", "17th October, 2034 @ 01:08 am "},
		{"Sep 03 2014 4:32 pm +0630", "Sep 03 2014 4:32 pm "},
		// String won't be changed if there are no valid timezone
		{"15 May 2004", "15 May 2004"},
	}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => \"%s\"", test.Text, test.Expected)

		// Check the offset
		leftover, _ := PopTzOffset(test.Text)
		passed := assert.Equal(t, test.Expected, leftover, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
