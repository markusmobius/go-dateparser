package parser

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/setting"
)

func parseFormattedDate(cfg *setting.Configuration, str string, formats ...string) DateData {
	period := Day

	currentTime := time.Now()
	if !cfg.CurrentTime.IsZero() {
		currentTime = cfg.CurrentTime
	}

	for _, format := range formats {
		// Parse time
		t, err := time.Parse(format, str)
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
			period = Month
			t = applyDayFromConfig(t, cfg)
		}

		// Check if format has year
		if t.Year() == 0 {
			t = time.Date(currentTime.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
				t.Location())
		}

		return DateData{Time: t, Period: period}
	}

	return DateData{Period: period}
}
