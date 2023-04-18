package jalali

import (
	stdregex "regexp"
	"strconv"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/regexp"
)

var (
	rxHourPattern   = stdregex.MustCompile(`ساعت\s+\d{2}`)
	rxMinutePattern = stdregex.MustCompile(`\d{2}\s+دقیقه`)
	rxSecondPattern = stdregex.MustCompile(`\d{2}\s+ثانیه`)

	rxColonPattern   = regexp.MustCompile(`\s+و\s+`)
	rxOrdinalPattern = regexp.MustCompile(`ام|م|ین`)
	rxNonDigit       = regexp.MustCompile(`\D`)
)

var monthNames = []struct {
	Latin    string
	Persians []string
}{
	{"Farvardin", []string{"فروردین"}},
	{"Ordibehesht", []string{"اردیبهشت"}},
	{"Khordad", []string{"خرداد"}},
	{"Tir", []string{"تیر"}},
	{"Mordad", []string{"امرداد", "مرداد"}},
	{"Shahrivar", []string{"شهریور", "شهريور"}},
	{"Mehr", []string{"مهر"}},
	{"Aban", []string{"ابان"}},
	{"Azar", []string{"اذر"}},
	{"Dey", []string{"دی"}},
	{"Bahman", []string{"بهمن", "بهن"}},
	{"Esfand", []string{"اسفند"}},
}

var monthNameValues = map[string]int{
	"Farvardin":   1,
	"Ordibehesht": 2,
	"Khordad":     3,
	"Tir":         4,
	"Mordad":      5,
	"Shahrivar":   6,
	"Mehr":        7,
	"Aban":        8,
	"Azar":        9,
	"Dey":         10,
	"Bahman":      11,
	"Esfand":      12,
}

var weekdayNames = []struct {
	Translation string
	Persians    []string
}{
	{"IekShanbe", []string{"یکشنبه"}},
	{"DoShanbe", []string{"دوشنبه"}},
	{"SeShanbe", []string{"سهشنبه", "سه شنبه"}},
	{"ChaharShanbe", []string{"چهارشنبه", "چهار شنبه"}},
	{"PanjShanbe", []string{"پنجشنبه", "پنج شنبه"}},
	{"Jome", []string{"جمعه"}},
	{"Shanbe", []string{"روز شنبه", "شنبه"}},
}

var weekdayNameValues = map[string]int{
	"Shanbe":       0,
	"IekShanbe":    1,
	"DoShanbe":     2,
	"SeShanbe":     3,
	"ChaharShanbe": 4,
	"PanjShanbe":   5,
	"Jome":         6,
}

var dayNumbers = []struct {
	Number   int
	Persians []string
}{
	{31, []string{"سی و یک"}},
	{13, []string{"سیزده"}},
	{29, []string{"بیست و نه"}},
	{28, []string{"بیست و هشت"}},
	{27, []string{"بیست و هفت"}},
	{26, []string{"بیست و شش"}},
	{25, []string{"بیست و پنج"}},
	{24, []string{"بیست و چهار"}},
	{23, []string{"بیست و سه", "بیست و سو"}},
	{22, []string{"بیست و دو", "بیست ثانیه"}},
	{21, []string{"بیست و یک"}},
	{20, []string{"بیست"}},
	{19, []string{"نوزده"}},
	{18, []string{"هجده"}},
	{17, []string{"هفده"}},
	{16, []string{"شانزده"}},
	{15, []string{"پانزده"}},
	{14, []string{"چهارده"}},
	{30, []string{"سی"}},
	{12, []string{"دوازده"}},
	{11, []string{"یازده"}},
	{10, []string{"ده"}},
	{9, []string{"نه"}},
	{8, []string{"هشت"}},
	{7, []string{"هفت"}},
	{6, []string{"شش"}},
	{5, []string{"پنج"}},
	{4, []string{"چهار"}},
	{3, []string{"سه", "سو"}},
	{2, []string{"دو"}},
	{1, []string{"یک", "اول"}},
	{0, []string{"صفر"}},
}

func replaceMonths(s string) string {
	for _, mn := range monthNames {
		for _, persian := range mn.Persians {
			s = strings.ReplaceAll(s, persian, " "+mn.Latin+" ")
		}
	}
	return s
}

func replaceWeekdays(s string) string {
	for _, wn := range weekdayNames {
		for _, persian := range wn.Persians {
			s = strings.ReplaceAll(s, persian, " "+wn.Translation+" ")
		}
	}
	return s
}

func replaceTime(s string) string {
	fnNumberOnly := func(s string) string {
		return rxNonDigit.ReplaceAllString(s, " ")
	}

	s = rxHourPattern.ReplaceAllStringFunc(s, fnNumberOnly)
	s = rxMinutePattern.ReplaceAllStringFunc(s, fnNumberOnly)
	s = rxSecondPattern.ReplaceAllStringFunc(s, fnNumberOnly)
	s = rxColonPattern.ReplaceAllString(s, ":")
	s = strings.ReplaceAll(s, "ساعت", "")
	return s
}

func replaceDays(s string) string {
	s = rxOrdinalPattern.ReplaceAllString(s, "")
	for _, dn := range dayNumbers {
		number := strconv.Itoa(dn.Number)
		for _, persian := range dn.Persians {
			s = strings.ReplaceAll(s, persian, number)
		}
	}
	return s
}

func removeWeekdayTranslations(s string) string {
	for _, wn := range weekdayNames {
		s = strings.ReplaceAll(s, wn.Translation, " ")
	}
	return s
}
