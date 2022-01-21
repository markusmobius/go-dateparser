package dateutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLeapYear(t *testing.T) {
	type testScenario struct {
		Year   int
		IsLeap bool
	}

	tests := []testScenario{
		// Not leap year
		{1999, false},
		{1700, false},
		// Is leap year
		{1696, true},
		{1996, true},
		{2000, true},
		{2016, true},
		{2096, true},
		{2392, true},
	}

	for _, test := range tests {
		message := fmt.Sprintf("%d, IS LEAP YEAR: %v", test.Year, test.IsLeap)
		assert.Equal(t, test.IsLeap, IsLeapYear(test.Year), message)
	}
}

func TestGetLeapYear_future(t *testing.T) {
	type testScenario struct {
		Year         int
		NextLeapYear int
	}

	tests := []testScenario{
		{2020, 2024},
		{1996, 2000},
		{2096, 2104},
		{1696, 1704},
		{2396, 2400},
		{0, 4},
	}

	for _, test := range tests {
		message := fmt.Sprintf("%d, NEXT LEAP YEAR IS: %d", test.Year, test.NextLeapYear)
		assert.Equal(t, test.NextLeapYear, GetLeapYear(test.Year, true), message)
	}
}

func TestGetLeapYear_past(t *testing.T) {
	type testScenario struct {
		Year         int
		PrevLeapYear int
	}

	tests := []testScenario{
		{2020, 2016},
		{2000, 1996},
		{2104, 2096},
		{1704, 1696},
		{2396, 2392},
		{0, -4},
	}

	for _, test := range tests {
		message := fmt.Sprintf("%d, PREV LEAP YEAR WAS: %d", test.Year, test.PrevLeapYear)
		assert.Equal(t, test.PrevLeapYear, GetLeapYear(test.Year, false), message)
	}
}

func TestGetLastDayOfMonth(t *testing.T) {
	type testScenario struct {
		Year    int
		Month   int
		LastDay int
	}

	tests := []testScenario{
		{2111, 1, 31},
		{1999, 2, 28},
		{1996, 2, 29},
		{2000, 2, 29},
		{1700, 2, 28},
		{2020, 3, 31},
		{1987, 4, 30},
		{1000, 5, 31},
		{1534, 6, 30},
		{1777, 7, 31},
		{1234, 8, 31},
		{1678, 9, 30},
		{1947, 10, 31},
		{2015, 11, 30},
		{2300, 12, 31},
	}

	for _, test := range tests {
		message := fmt.Sprintf("%d-%02d, LAST DAY IS: %d", test.Year, test.Month, test.LastDay)
		assert.Equal(t, test.LastDay, GetLastDayOfMonth(test.Year, test.Month), message)
	}
}
