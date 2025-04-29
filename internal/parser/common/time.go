package common

import (
	"fmt"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
)

// ParseTime parses the specified string using one of the common time formats.
func ParseTime(s string) (time.Time, date.Period, error) {
	timeFormats := []struct {
		Pattern string
		Period  date.Period
	}{
		{"15:4:5", date.Second},
		{"3:4:5 PM", date.Second},
		{"15:4", date.Minute},
		{"3:4 PM", date.Minute},
		{"3 PM", date.Hour},
		{"15:4:5.999999", date.Second},
		{"3:4:5.999999 PM", date.Second},
		{"15:4 PM", date.Minute},
	}

	tmp := strings.TrimSpace(s)
	tmp = strings.ToUpper(tmp)
	for _, format := range timeFormats {
		t, err := time.Parse(format.Pattern, tmp)
		if err != nil {
			continue
		}

		return t, format.Period, nil
	}

	err := fmt.Errorf("%s doesn't seem to be a valid time", s)
	return time.Time{}, date.None, err
}
