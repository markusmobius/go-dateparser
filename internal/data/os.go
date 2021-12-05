package data

import "regexp"

var osLocale = LocaleData{
	Name:                  "os",
	DateOrder:             "DMY",
	January:               []string{"янв", "январы", "январь"},
	February:              []string{"фев", "февр", "февралы", "февраль"},
	March:                 []string{"мар", "март", "мартъи", "мартъийы"},
	April:                 []string{"апр", "апрелы", "апрель"},
	May:                   []string{"май", "майы"},
	June:                  []string{"июны", "июнь"},
	July:                  []string{"июлы", "июль"},
	August:                []string{"авг", "август", "августы"},
	September:             []string{"сен", "сент", "сентябры", "сентябрь"},
	October:               []string{"окт", "октябры", "октябрь"},
	November:              []string{"ноя", "нояб", "ноябры", "ноябрь"},
	December:              []string{"дек", "декабры", "декабрь"},
	Monday:                []string{"крс", "къуырисӕр"},
	Tuesday:               []string{"дцг", "дыццӕг"},
	Wednesday:             []string{"ӕрт", "ӕртыццӕг"},
	Thursday:              []string{"цпр", "цыппӕрӕм"},
	Friday:                []string{"майрӕмбон", "мрб"},
	Saturday:              []string{"сабат", "сбт"},
	Sunday:                []string{"хуыцаубон", "хцб"},
	AM:                    []string{"am", "ӕмбисбоны размӕ"},
	PM:                    []string{"pm", "ӕмбисбоны фӕстӕ"},
	Year:                  []string{"аз"},
	Month:                 []string{"мӕй"},
	Week:                  []string{"къуыри"},
	Day:                   []string{"бон"},
	Hour:                  []string{"сахат"},
	Minute:                []string{"минут"},
	Second:                []string{"секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`абон`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`знон`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`сом`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) бон раздӕр`),
			regexp.MustCompile(`(?i)(\d+) боны размӕ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) сахаты размӕ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) боны фӕстӕ`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) сахаты фӕстӕ`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"os-RU": osRULocale,
	},
}

var osRULocale = LocaleData{
	Name:                  "os-RU",
	DateOrder:             "",
}
