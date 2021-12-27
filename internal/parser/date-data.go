package parser

import "time"

type DatePeriod uint8

const (
	Day DatePeriod = iota
	Month
	Year
)

type DateData struct {
	Locale string
	Period DatePeriod
	Time   time.Time
}
