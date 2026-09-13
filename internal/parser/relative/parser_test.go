package relative

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func TestFractionalPeriodPrecision(t *testing.T) {
	actual := Parse(&setting.Configuration{CurrentTime: time.Date(2026, 9, 12, 12, 30, 45, 0, time.UTC), ReturnTimeAsPeriod: true}, "1.1.1 day ago")
	if actual.Period != date.Minute {
		t.Fatalf("fractional day period = %v, want minute", actual.Period)
	}
}

func TestPythonRelativeArithmetic(t *testing.T) {
	base := time.Date(2023, 1, 31, 12, 0, 0, 0, time.UTC)
	for _, testCase := range []struct{ input, expected string }{
		{"in 1 month", "2023-02-28T12:00:00Z"},
		{"1 month ago", "2022-12-31T12:00:00Z"},
		{"in 1.5 day", "2023-02-02T00:00:00Z"},
		{"0.5 second ago", "2023-01-31T11:59:59.5Z"},
		{"in 1.5 year", ""},
		{"in 1.5 month", ""},
	} {
		for _, preserve := range []bool{false, true} {
			actual := Parse(&setting.Configuration{CurrentTime: base, PreserveEndOfMonth: preserve}, testCase.input).Time
			value := ""
			if !actual.IsZero() {
				value = actual.Format(time.RFC3339Nano)
			}
			if value != testCase.expected {
				t.Errorf("%q preserve=%t: %s, Python expects %s", testCase.input, preserve, value, testCase.expected)
			}
		}
	}
}
