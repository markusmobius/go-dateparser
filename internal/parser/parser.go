package parser

import (
	"fmt"
	"strings"
	"time"
)

func parseTime(s string) (time.Time, error) {
	timeFormats := []string{
		"15:04:05",
		"3:04:05 PM",
		"15:04",
		"3:04 PM",
		"3 PM",
		"15:04:05.999999",
		"3:04:05.999999 PM",
		"15:04 PM",
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
