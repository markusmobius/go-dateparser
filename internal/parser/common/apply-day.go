package common

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/dateutil"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func ApplyDayFromConfig(cfg *setting.Configuration, t time.Time, currentDay ...int) time.Time {
	if cfg == nil {
		return t
	}

	var newDay int
	switch cfg.PreferredDayOfMonth {
	case setting.First:
		newDay = 1
	case setting.Last:
		newDay = t.AddDate(0, 1, -t.Day()).Day()
	default: // current day
		if len(currentDay) > 0 {
			newDay = currentDay[0]
		} else {
			if cfg.CurrentTime.IsZero() {
				newDay = time.Now().UTC().Day()
			} else {
				newDay = cfg.CurrentTime.Day()
			}
		}
	}

	// Make sure new day is not bigger than max day
	lastDayOfMonth := dateutil.GetLastDayOfMonth(t.Year(), int(t.Month()))
	if newDay > lastDayOfMonth {
		newDay = lastDayOfMonth
	}

	return time.Date(t.Year(), t.Month(), newDay,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		t.Location())
}
