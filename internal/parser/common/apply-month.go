package common

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
)

func ApplyMonthFromConfig(cfg *setting.Configuration, t time.Time, currentMonth ...time.Month) time.Time {
	if cfg == nil {
		return t
	}

	var newMonth time.Month
	switch cfg.PreferredMonthOfYear {
	case setting.FirstMonth:
		newMonth = 1
	case setting.LastMonth:
		newMonth = 12
	default: // current day
		if len(currentMonth) > 0 {
			newMonth = time.Month(currentMonth[0])
		} else {
			if cfg.CurrentTime.IsZero() {
				newMonth = time.Now().UTC().Month()
			} else {
				newMonth = cfg.CurrentTime.Month()
			}
		}
	}

	return time.Date(t.Year(), newMonth, t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		t.Location())
}
