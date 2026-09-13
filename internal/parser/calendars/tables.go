package calendars

import (
	_ "embed"
	"encoding/json"
	"errors"
	"sort"
	"time"
)

//go:embed data.json
var encodedData []byte

type calendarData struct {
	FirstYear    int   `json:"first_year"`
	YearStarts   []int `json:"year_starts"`
	MonthStarts  []int `json:"month_starts"`
	GregorianMin int   `json:"gregorian_min"`
	GregorianMax int   `json:"gregorian_max"`
}

var data = func() struct {
	Jalali calendarData `json:"jalali"`
	Hijri  calendarData `json:"hijri"`
} {
	var result struct {
		Jalali calendarData `json:"jalali"`
		Hijri  calendarData `json:"hijri"`
	}
	if err := json.Unmarshal(encodedData, &result); err != nil {
		panic(err)
	}
	return result
}()

var (
	errJalaliRange = errors.New("date is outside the supported Jalali range")
	errHijriRange  = errors.New("date is outside Umm al-Qura scope")
)

func gregorianDay(value time.Time) int {
	return int(time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, time.UTC).Unix() / 86400)
}

func gregorianParts(day int) ([3]int, error) {
	value := time.Unix(int64(day)*86400, 0).UTC()
	if value.Year() < 1 || value.Year() > 9999 {
		return [3]int{}, errJalaliRange
	}
	return [3]int{value.Year(), int(value.Month()), value.Day()}, nil
}

func JalaliFromGregorian(value time.Time) ([3]int, error) {
	day := gregorianDay(value)
	if day < data.Jalali.GregorianMin || day > data.Jalali.GregorianMax {
		return [3]int{}, errJalaliRange
	}
	index := sort.Search(len(data.Jalali.YearStarts), func(index int) bool {
		return data.Jalali.YearStarts[index] > day
	}) - 1
	remaining := day - data.Jalali.YearStarts[index]
	month, monthDay := remaining/31+1, remaining%31+1
	if remaining >= 186 {
		month, monthDay = (remaining-186)/30+7, (remaining-186)%30+1
	}
	return [3]int{data.Jalali.FirstYear + index, month, monthDay}, nil
}

func JalaliToGregorian(year, month, day int) ([3]int, error) {
	if year < data.Jalali.FirstYear || year >= data.Jalali.FirstYear+len(data.Jalali.YearStarts) || month < 1 || month > 12 || day < 1 {
		return [3]int{}, errJalaliRange
	}
	start := data.Jalali.YearStarts[year-data.Jalali.FirstYear]
	monthDays := (month - 1) * 31
	if month > 7 {
		monthDays = (month-1)*30 + 6
	}
	return gregorianParts(start + monthDays + day - 1)
}

func JalaliMonthLength(year, month int) (int, error) {
	if year < data.Jalali.FirstYear || year >= data.Jalali.FirstYear+len(data.Jalali.YearStarts) || month < 1 || month > 12 {
		return 0, errJalaliRange
	}
	if month <= 6 {
		return 31, nil
	}
	if month < 12 {
		return 30, nil
	}
	index := year - data.Jalali.FirstYear
	if index+1 >= len(data.Jalali.YearStarts) {
		return 0, errJalaliRange
	}
	return data.Jalali.YearStarts[index+1] - data.Jalali.YearStarts[index] - 336, nil
}

func HijriFromGregorian(value time.Time) ([3]int, error) {
	day := gregorianDay(value)
	if day < data.Hijri.GregorianMin || day > data.Hijri.GregorianMax {
		return [3]int{}, errHijriRange
	}
	index := sort.Search(len(data.Hijri.MonthStarts), func(index int) bool {
		return data.Hijri.MonthStarts[index] > day
	}) - 1
	return [3]int{data.Hijri.FirstYear + index/12, index%12 + 1, day - data.Hijri.MonthStarts[index] + 1}, nil
}

func hijriIndex(year, month int) (int, error) {
	if year < data.Hijri.FirstYear || year >= data.Hijri.FirstYear+(len(data.Hijri.MonthStarts)-1)/12 || month < 1 || month > 12 {
		return 0, errHijriRange
	}
	return (year-data.Hijri.FirstYear)*12 + month - 1, nil
}

func HijriToGregorian(year, month, day int) ([3]int, error) {
	index, err := hijriIndex(year, month)
	if err != nil || day < 1 {
		return [3]int{}, errHijriRange
	}
	return gregorianParts(data.Hijri.MonthStarts[index] + day - 1)
}

func HijriMonthLength(year, month int) (int, error) {
	index, err := hijriIndex(year, month)
	if err != nil {
		return 0, err
	}
	return data.Hijri.MonthStarts[index+1] - data.Hijri.MonthStarts[index], nil
}
