package common

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
)

func ApplyDayFromConfig(cfg *setting.Configuration, t time.Time, currentDay ...int) time.Time {
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
