package formatted

import (
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var (
	rxDayNumber = regexp.MustCompile(`2.{0,3}`)
)

// Parse the specified string using one of the specified formats.
func Parse(cfg *setting.Configuration, str string, formats ...string) date.Date {
	// Create initial period
	period := date.Day

	// Check if string contain timezone
	var timeLoc *time.Location
	if _, tzData := timezone.PopTzOffset(str); tzData.IsZero() {
		timeLoc = time.UTC
	} else {
		timeLoc = time.FixedZone(tzData.Name, tzData.Offset)
	}

	// Fetch current time
	currentTime := time.Now().UTC()
	if !cfg.CurrentTime.IsZero() {
		currentTime = cfg.CurrentTime
	}

	// Try each format
	for _, format := range formats {
		// Parse time
		t, err := time.ParseInLocation(format, str, timeLoc)
		if err != nil {
			continue
		}

		// Check if format has day
		var formatHasDay bool
		for _, match := range rxDayNumber.FindAllString(format, -1) {
			if match != "2006" {
				formatHasDay = true
				break
			}
		}

		if !formatHasDay {
			period = date.Month
			t = common.ApplyDayFromConfig(cfg, t)
		}

		// Check if format has year
		if t.Year() == 0 {
			t = time.Date(currentTime.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				t.Location())
		}

		return date.Date{Time: t, Period: period}
	}

	return date.Date{}
}
