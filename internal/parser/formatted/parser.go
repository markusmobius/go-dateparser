package formatted

import (
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

// Parse the specified string using one of the specified formats.
func Parse(cfg *setting.Configuration, str string, formats ...string) date.Date {
	// Create initial period
	period := date.Day

	// Check if string contain timezone
	var timeLoc *time.Location
	if _, tzData := timezone.PopTzOffset(str); tzData.IsZero() {
		if cfg.DefaultTimezone != nil {
			timeLoc = cfg.DefaultTimezone
		} else {
			timeLoc = time.UTC
		}
	} else {
		timeLoc = time.FixedZone(tzData.Name, tzData.Offset)
	}

	// Fetch current time
	currentTime := time.Now().UTC()
	if !cfg.CurrentTime.IsZero() {
		currentTime = cfg.CurrentTime
	}

	// Try each format
	checker := time.Date(12, 3, 4, 5, 6, 7, 8, time.UTC)

	for _, format := range formats {
		// Parse time
		t, err := time.ParseInLocation(format, str, timeLoc)
		if err != nil {
			continue
		}

		// Check if format has day or month
		checkerText := checker.Format(format)
		formatHasDay := strings.Contains(checkerText, "4")
		formatHasMonth := strings.Contains(checkerText, "3") || strings.Contains(checkerText, "Mar")

		if !formatHasMonth && !formatHasDay {
			period = date.Year
			t = common.ApplyMonthFromConfig(cfg, t)
			t = common.ApplyDayFromConfig(cfg, t)
		} else if !formatHasMonth {
			period = date.Year
			t = common.ApplyMonthFromConfig(cfg, t)
		} else if !formatHasDay {
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
