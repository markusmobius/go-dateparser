package formatted

import (
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParseDayOfYear(t *testing.T) {
	cfg := &setting.Configuration{
		CurrentTime:          time.Date(2025, 9, 15, 12, 0, 0, 0, time.UTC),
		PreferredDayOfMonth:  setting.Last,
		PreferredMonthOfYear: setting.FirstMonth,
	}
	tests := []struct{ Text, Format, Expected string }{
		{"2023-100", "2006-002", "2023-04-10"},
		{"2024-060", "2006-002", "2024-02-29"},
		{"2023 060", "2006 002", "2023-03-01"},
		{"2024  60", "2006 __2", "2024-02-29"},
		{"1999001", "2006002", "1999-01-01"},
		{"1999050", "2006002", "1999-02-19"},
		{"1999032", "2006002", "1999-02-01"},
		{"1999059", "2006002", "1999-02-28"},
		{"1999060", "2006002", "1999-03-01"},
		{"1999365", "2006002", "1999-12-31"},
		{"2000060", "2006002", "2000-02-29"},
		{"2000366", "2006002", "2000-12-31"},
		{"1999366", "2006002", ""},
		{"2023-366", "2006-002", ""},
		{"1900366", "2006002", ""},
		{"2100 366", "2006 002", ""},
		{"1999000", "2006002", ""},
		{"1999367", "2006002", ""},
		{"1999777", "2006002", ""},
	}
	for _, test := range tests {
		t.Run(test.Text, func(t *testing.T) {
			parsed := Parse(cfg, test.Text, test.Format)
			if test.Expected == "" {
				assert.True(t, parsed.IsZero())
			} else {
				assert.Equal(t, test.Expected, parsed.Time.Format("2006-01-02"))
				assert.Equal(t, date.Day, parsed.Period)
			}
		})
	}
}

func TestParseCenturyPreference(t *testing.T) {
	tests := []struct {
		Text, Format string
		Preference   setting.PreferredDateSource
		BaseYear     int
		Expected     string
	}{
		{"1/15/64", "1/2/06", setting.Past, 2026, "1964-01-15"},
		{"1/15/64", "1/2/06", setting.Future, 2026, "2064-01-15"},
		{"1/15/24", "1/2/06", setting.Future, 2026, "2124-01-15"},
		{"1/15/35", "1/2/06", setting.Past, 2026, "1935-01-15"},
		{"1/15/2050", "1/2/2006", setting.Past, 2026, "2050-01-15"},
		{"1/15/64", "1/2/06", setting.Past, 1980, "1964-01-15"},
		{"1/15/64", "1/2/06", setting.Future, 2020, "2064-01-15"},
		{"2/29/00", "1/2/06", setting.Future, 2026, "2104-02-29"},
		{"2/29/00", "1/2/06", setting.Past, 1980, "1896-02-29"},
		{"1/1/26", "1/2/06", setting.Future, 2026, "2126-01-01"},
	}
	for _, test := range tests {
		t.Run(test.Text+"/"+test.Expected, func(t *testing.T) {
			cfg := &setting.Configuration{
				CurrentTime:         time.Date(test.BaseYear, 1, 1, 0, 0, 0, 0, time.FixedZone("base", 2*3600)),
				PreferredDateSource: test.Preference,
			}
			parsed := Parse(cfg, test.Text, test.Format)
			assert.Equal(t, test.Expected, parsed.Time.Format("2006-01-02"))
			assert.Equal(t, date.Day, parsed.Period)
		})
	}
}

func TestParse(t *testing.T) {
	// Helper function
	dateIsParsed := func(dt date.Date) { assert.NotZero(t, dt.Time) }
	dateIsNotParsed := func(dt date.Date) { assert.Zero(t, dt.Time) }
	dateAsExpected := func(dt date.Date, expected string, expectedFormat ...string) {
		dateIsParsed(dt)

		format := "2006-01-02"
		if len(expectedFormat) > 0 {
			format = expectedFormat[0]
		}

		str := dt.Time.Format(format)
		assert.Equal(t, expected, str)
	}

	// Prepare configuration
	cfg := setting.Configuration{
		CurrentTime: time.Date(2015, 2, 4, 0, 0, 0, 0, time.UTC),
	}

	// No matching format, shouldn't be parsed
	dt := Parse(&cfg, "yesterday", "2006-01-02")
	dateIsNotParsed(dt)

	// Matching format, should be parsed
	dt = Parse(&cfg, "25-03-14", "02-01-06")
	dateAsExpected(dt, "2014-03-25")

	// Should use current year for dates without year
	dt = Parse(&cfg, "09.16", "01.02")
	dateAsExpected(dt, "2015-09-16")

	// Should use day from config for dates without day
	str, format := "August 2014", "January 2006"
	cfg = setting.Configuration{
		CurrentTime: time.Date(2014, 8, 12, 0, 0, 0, 0, time.UTC),
	}

	cfg.PreferredDayOfMonth = setting.First
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-01")

	cfg.PreferredDayOfMonth = setting.Last
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-31")

	cfg.PreferredDayOfMonth = setting.Current
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-12")

	// Should use month from config for dates without month
	str, format = "2014", "2006"
	cfg = setting.Configuration{
		CurrentTime: time.Date(2014, 7, 1, 0, 0, 0, 0, time.UTC),
	}

	cfg.PreferredMonthOfYear = setting.FirstMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-01-01")

	cfg.PreferredMonthOfYear = setting.LastMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-12-01")

	cfg.PreferredMonthOfYear = setting.CurrentMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-07-01")
}
