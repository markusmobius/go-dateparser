package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
)

func parseTimestamp(cfg *setting.Configuration, str string) time.Time {
	parts := rxTimestamp.FindStringSubmatch(str)
	if len(parts) > 0 {
		seconds, _ := strconv.ParseInt(parts[1], 10, 64)
		millis, _ := strconv.ParseInt(parts[2], 10, 64)
		micros, _ := strconv.ParseInt(parts[3], 10, 64)
		nanos := micros*1_000 + millis*1_000_000

		t := time.Unix(seconds, nanos)
		return t
	}

	return time.Time{}
}

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
