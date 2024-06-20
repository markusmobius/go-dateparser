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
	Time
	Day
	Month
	Year
	Hour
	Minute
	Second
)

func (p Period) String() string {
	switch p {
	case Time:
		return "Time"
	case Second:
		return "Second"
	case Minute:
		return "Minute"
	case Hour:
		return "Hour"
	case Day:
		return "Day"
	case Month:
		return "Month"
	case Year:
		return "Year"
	default:
		return ""
	}
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
