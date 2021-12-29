package parser

import "time"

type DatePeriod uint8

const (
	None DatePeriod = iota
	Day
	Month
	Year
	Time
)

type DateData struct {
	Locale string
	Period DatePeriod
	Time   time.Time
}
