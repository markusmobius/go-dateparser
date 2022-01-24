package godateparser

import (
	"time"

	"github.com/markusmobius/go-dateparser/internal/parser/date"
)

// Period is representation of confidence level for the parsed date.
// For example, if the parsed date has period level `Month`, then
// the parser is only confident up to the month, while the day and
// time is taken from the current time.
type Period uint8

const (
	None Period = iota
	Time
	Day
	Month
	Year
)

// Date is object that represents the parsed date with useful information.
type Date struct {
	Locale string
	Period Period
	Time   time.Time
}

// IsZero reports whether the date is empty or not.
func (d Date) IsZero() bool {
	return d.Period == None || d.Time.IsZero()
}

func dateFromInternal(d date.Date) Date {
	return Date{
		Locale: d.Locale,
		Period: Period(d.Period),
		Time:   d.Time,
	}
}
