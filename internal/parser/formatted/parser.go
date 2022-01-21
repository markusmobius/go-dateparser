package formatted

import (
	"regexp"
	"time"

	"github.com/markusmobius/go-dateparser/internal/parser/common"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

var (
	rxDayNumber = regexp.MustCompile(`2.{0,3}`)
)

func Parse(cfg *setting.Configuration, str string, formats ...string) date.Date {
	period := date.Day

	currentTime := time.Now().UTC()
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

	return date.Date{Period: period}
}
