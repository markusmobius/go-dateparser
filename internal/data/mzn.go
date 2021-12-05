package data

import "regexp"

var mznLocale = LocaleData{
	Name:                  "mzn",
	DateOrder:             "YMD",
	January:               []string{"ژانویه"},
	February:              []string{"فوریه"},
	March:                 []string{"مارس"},
	April:                 []string{"آوریل"},
	May:                   []string{"مه"},
	June:                  []string{"ژوئن"},
	July:                  []string{"ژوئیه"},
	August:                []string{"اوت"},
	September:             []string{"سپتامبر"},
	October:               []string{"اکتبر"},
	November:              []string{"نوامبر"},
	December:              []string{"دسامبر"},
	Monday:                []string{"mon"},
	Tuesday:               []string{"tue"},
	Wednesday:             []string{"wed"},
	Thursday:              []string{"thu"},
	Friday:                []string{"fri"},
	Saturday:              []string{"sat"},
	Sunday:                []string{"sun"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"سال"},
	Month:                 []string{"ماه"},
	Week:                  []string{"هفته"},
	Day:                   []string{"روز"},
	Hour:                  []string{"ساعت", "ساعِت"},
	Minute:                []string{"دقیقه", "دَقه"},
	Second:                []string{"ثانیه"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`اَمروز`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`این ماه`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`این هفته`},
		`0 year ago`:   {`امسال`},
		`1 day ago`:    {`دیروز`},
		`1 month ago`:  {`ماه قبل`},
		`1 week ago`:   {`قبلی هفته`},
		`1 year ago`:   {`پارسال`},
		`in 1 day`:     {`فِردا`},
		`in 1 month`:   {`ماه ِبعد`},
		`in 1 week`:    {`بعدی هفته`},
		`in 1 year`:    {`سال دیگه`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) روز پیش`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ساعت پیش`),
			regexp.MustCompile(`(?i)(\d+) ساعِت پیش`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) دَقه پیش`),
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
			regexp.MustCompile(`(?i)(\d+) روز دله`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ساعت دله`),
			regexp.MustCompile(`(?i)(\d+) ساعِت دله`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) دقیقه دله`),
			regexp.MustCompile(`(?i)(\d+) دَقه دله`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ماه دله`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) ثانیه دله`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) هفته دله`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) سال دله`),
		},
	},
}
