package date

import (
	"time"
)

// Period describes the precision of a parsed date, not a probability of correctness.
// For example, Month means the input identifies a month; finer components may
// be filled from the parser's configuration or defaults.
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

// Date contains a parsed time, its precision, and the translation locale when available.
type Date struct {
	Locale string
	Period Period
	Time   time.Time
}

// IsZero reports whether the date is empty or not.
func (d Date) IsZero() bool {
	return d.Period == None || d.Time.IsZero()
}
