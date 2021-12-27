package parser

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
)

func applyDayFromConfig(t time.Time, cfg *setting.Configuration, currentDay ...int) time.Time {
	if cfg == nil {
		return t
	}

	var newDay int
	switch cfg.PreferredDayOfMonth {
	case "first":
		newDay = 1
	case "last":
		newDay = t.AddDate(0, 1, -t.Day()).Day()
	case "current":
		if len(currentDay) > 0 {
			newDay = currentDay[0]
		} else {
			if cfg.CurrentTime.IsZero() {
				newDay = time.Now().Day()
			} else {
				newDay = cfg.CurrentTime.Day()
			}
		}
	}

	return time.Date(t.Year(), t.Month(), newDay,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		t.Location())
}

func applyTimezoneFromConfig(t time.Time, cfg *setting.Configuration) time.Time {
	if cfg == nil {
		return t
	}

	if cfg.Timezone != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			cloneTimezone(cfg.Timezone))
	}

	if cfg.ToTimezone != nil {
		tz := cloneTimezone(cfg.ToTimezone)
		t = t.In(tz)
	}

	if !cfg.ReturnAsTimezoneAware {
		t = time.Date(t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
			time.UTC)
	}

	return t
}

func cloneTimezone(tz *time.Location) *time.Location {
	return time.FixedZone(time.Now().In(tz).Zone())
}
