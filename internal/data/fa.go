package data

import "regexp"

var faLocale = LocaleData{
	Name:                  "fa",
	DateOrder:             "YMD",
	SentenceSplitterGroup: 6,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "，"},
	January:               []string{"جنوری", "ژانویه", "ژانویهٔ"},
	February:              []string{"فبروری", "فوریه", "فوریهٔ"},
	March:                 []string{"مارس", "مارچ"},
	April:                 []string{"آوریل", "اپریل"},
	May:                   []string{"مه", "مهٔ", "می"},
	June:                  []string{"جون", "ژوئن"},
	July:                  []string{"جولای", "ژوئیه", "ژوئیهٔ"},
	August:                []string{"آگست", "اوت"},
	September:             []string{"سپتامبر", "سپتمبر"},
	October:               []string{"اکتبر", "اکتوبر"},
	November:              []string{"نوامبر", "نومبر"},
	December:              []string{"دسامبر", "دسمبر"},
	Monday:                []string{"دوشنبه"},
	Tuesday:               []string{"سهشنبه", "سه‌شنبه"},
	Wednesday:             []string{"چهار شنبه", "چهارشنبه"},
	Thursday:              []string{"پنج شنبه", "پنجشنبه"},
	Friday:                []string{"جمعه"},
	Saturday:              []string{"د", "دو شنبه", "روز شنبه", "شنبه"},
	Sunday:                []string{"یکشنبه"},
	AM:                    []string{"قبل‌ازظهر", "قظ"},
	PM:                    []string{"بظ", "بعدازظهر"},
	Year:                  []string{"سال"},
	Month:                 []string{"ماه"},
	Week:                  []string{"هفته"},
	Day:                   []string{"روز"},
	Hour:                  []string{"ساعت"},
	Minute:                []string{"دقیقه"},
	Second:                []string{"ثانیه", "دوم"},
	In:                    []string{"در"},
	Ago:                   []string{"پیش"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`امروز`},
		`0 hour ago`:   {`همین ساعت`},
		`0 minute ago`: {`همین دقیقه`},
		`0 month ago`:  {`این ماه`},
		`0 second ago`: {`اکنون`},
		`0 week ago`:   {`این هفته`},
		`0 year ago`:   {`امسال`},
		`1 day ago`:    {`دیروز`},
		`1 month ago`:  {`ماه پیش`, `ماه گذشته`},
		`1 week ago`:   {`هفتهٔ گذشته`},
		`1 year ago`:   {`سال گذشته`},
		`in 1 day`:     {`فردا`},
		`in 1 month`:   {`ماه آینده`},
		`in 1 week`:    {`هفتهٔ آینده`},
		`in 1 year`:    {`سال آینده`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) روز پیش`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ساعت پیش`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) دقیقه پیش`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ماه پیش`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) ثانیه پیش`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) هفته پیش`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) سال پیش`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) روز بعد`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ساعت بعد`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) دقیقه بعد`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ماه بعد`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) ثانیه بعد`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) هفته بعد`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) سال بعد`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"fa-AF": faAFLocale,
	},
}

var faAFLocale = LocaleData{
	Name:                  "fa-AF",
	DateOrder:             "",
	January:               []string{"جنو", "جنوری"},
	February:              []string{"فبروری"},
	March:                 []string{"مارچ"},
	April:                 []string{"اپریل"},
	May:                   []string{"می"},
	June:                  []string{"جون"},
	July:                  []string{"جول", "جولای"},
	August:                []string{"اگست"},
	September:             []string{"سپتمبر"},
	October:               []string{"اکتوبر"},
	November:              []string{"نومبر"},
	December:              []string{"دسم", "دسمبر"},
}
