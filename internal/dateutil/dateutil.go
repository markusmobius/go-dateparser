package dateutil

import "time"

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	} else {
		return false
	}
}

func GetLeapYear(year int, toFuture bool) int {
	step := 1
	if !toFuture {
		step = -1
	}

	for {
		year += step
		if IsLeapYear(year) {
			return year
		}
	}
}

func GetLastDayOfMonth(year, month int) int {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 1, -1).Day()
}
