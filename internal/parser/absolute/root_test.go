package absolute

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
	"github.com/stretchr/testify/assert"
)

func TestParse_error(t *testing.T) {
	// Prepare scenarios
	cfgStrict := &setting.Configuration{StrictParsing: true}
	cfgAllRequired := &setting.Configuration{RequiredParts: []string{"day", "month", "year"}}
	cfgDayRequired := &setting.Configuration{RequiredParts: []string{"day"}}
	cfgMonthRequired := &setting.Configuration{RequiredParts: []string{"month"}}
	cfgYearRequired := &setting.Configuration{RequiredParts: []string{"year"}}

	type testScenario struct {
		Text   string
		Config *setting.Configuration
	}

	tests := []testScenario{
		// Empty string
		{"", nil},

		// Not date
		{"invalid date string", nil},
		{"Aug 7, 2014Aug 7, 2014", nil},
		{"24h ago", nil},
		{"2015-03-17t16:37:51+00:002015-03-17t15:24:37+00:00", nil},
		{"8 enero 2013 martes 7:03 AM EST 8 enero 2013 martes 7:03 AM EST", nil},
		{"12/09/18567", nil},

		// Wrong day
		{"29 February 2015", nil},
		{"32 January 2015", nil},
		{"31 April 2015", nil},
		{"31 June 2015", nil},
		{"31 September 2015", nil},

		// Incomplete dates with strict parsing enabled
		{"april 2010", cfgStrict},
		{"11 March", cfgStrict},
		{"March", cfgStrict},
		{"31 2010", cfgStrict},
		{"31/2010", cfgStrict},

		// Partially complete dates, all date components required
		{"april 2010", cfgAllRequired},
		{"11 March", cfgAllRequired},
		{"March", cfgAllRequired},
		{"31 2010", cfgAllRequired},
		{"31/2010", cfgAllRequired},

		// Day is required
		{"april 2010", cfgDayRequired},
		{"March", cfgDayRequired},
		{"2010", cfgDayRequired},

		// Month required
		{"31 2010", cfgMonthRequired},
		{"31/2010", cfgMonthRequired},

		// Year required
		{"11 March", cfgYearRequired},
		{"March", cfgYearRequired},
	}

	// Start test
	var nFailed int
	var tz timezone.OffsetData
	for _, test := range tests {
		_, err := Parse(test.Config, test.Text, tz)
		passed := assert.Error(t, err, test.Text)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
