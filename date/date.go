package date

import (
	"time"
)

// Period is representation of confidence level for the parsed date.
// For example, if the parsed date has period level `Month`, then
// the parser is only confident up to the month, while the day and
// time is taken from the current time.
type Period uint8

const (
	None Period = iota
	Second
	Minute
	Hour
	Day
	Month
	Year
)

var periodString = map[Period]string{
	Second: "Second",
	Minute: "Minute",
	Hour:   "Hour",
	Day:    "Day",
	Month:  "Month",
	Year:   "Year",
}

func (p Period) String() string {
	return periodString[p]
}

func (p Period) IsTime() bool {
	return p == Second || p == Minute || p == Hour
}

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
