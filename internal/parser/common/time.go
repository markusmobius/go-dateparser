package common

import (
	"fmt"
	"strings"
	"time"
)

// ParseTime parses the specified string using one of the common time formats.
func ParseTime(s string) (time.Time, error) {
	timeFormats := []string{
		"15:4:5",
		"3:4:5 PM",
		"15:4",
		"3:4 PM",
		"3 PM",
		"15:4:5.999999",
		"3:4:5.999999 PM",
		"15:4 PM",
	}

	tmp := strings.TrimSpace(s)
	tmp = strings.ToUpper(tmp)
	for _, format := range timeFormats {
		t, err := time.Parse(format, tmp)
		if err != nil {
			continue
		}

		return t, nil
	}

	err := fmt.Errorf("%s doesn't seem to be a valid time", s)
	return time.Time{}, err
}
