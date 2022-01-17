package date

import "time"

type Period uint8

const (
	None Period = iota
	Time
	Day
	Week
	Month
	Year
)

type Date struct {
	Locale string
	Period Period
	Time   time.Time
}

func (d Date) IsZero() bool {
	return d.Period == None || d.Time.IsZero()
}
