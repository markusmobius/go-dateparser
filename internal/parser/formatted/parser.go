package formatted

import (
	"strings"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/dateutil"
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
		formatHasYearDay := strings.Contains(format, "002") || strings.Contains(format, "__2")
		formatHasDay := formatHasYearDay || strings.Contains(checkerText, "4")
		formatHasMonth := formatHasYearDay || strings.Contains(checkerText, "3") || strings.Contains(checkerText, "Mar")

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
		} else if strings.Contains(format, "06") && !strings.Contains(format, "2006") {
			now := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
				currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), t.Location())
			year := t.Year()
			if cfg.PreferredDateSource == setting.Past && now.Before(t) {
				year -= 100
			} else if cfg.PreferredDateSource == setting.Future && !now.Before(t) {
				year += 100
			}
			if t.Month() == time.February && t.Day() == 29 && !dateutil.IsLeapYear(year) {
				year = dateutil.GetLeapYear(year, cfg.PreferredDateSource == setting.Future)
			}
			t = time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
		}

		return date.Date{Time: t, Period: period}
	}

	return date.Date{}
}
