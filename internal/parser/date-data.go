package parser

import "time"

type DatePeriod uint8

const (
	None DatePeriod = iota
	Day
	Week
	Month
	Year
	Time
)

type DateData struct {
	Locale string
	Period DatePeriod
	Time   time.Time
}

func (dd DateData) IsValid() bool {
	return dd.Period != None && !dd.Time.IsZero()
}
