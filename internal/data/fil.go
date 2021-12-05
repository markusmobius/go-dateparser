// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var filLocale = LocaleData{
	Name:                  "fil",
	DateOrder:             "MDY",
	January:               []string{"ene", "enero"},
	February:              []string{"peb", "pebrero"},
	March:                 []string{"mar", "marso"},
	April:                 []string{"abr", "abril"},
	May:                   []string{"may", "mayo"},
	June:                  []string{"hun", "hunyo"},
	July:                  []string{"hul", "hulyo"},
	August:                []string{"ago", "agosto"},
	September:             []string{"set", "setyembre"},
	October:               []string{"okt", "oktubre"},
	November:              []string{"nob", "nobyembre"},
	December:              []string{"dis", "disyembre"},
	Monday:                []string{"lun", "lunes"},
	Tuesday:               []string{"mar", "martes"},
	Wednesday:             []string{"miy", "miyerkules"},
	Thursday:              []string{"huw", "huwebes"},
	Friday:                []string{"biy", "biyernes"},
	Saturday:              []string{"sab", "sabado"},
	Sunday:                []string{"lin", "linggo"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"taon"},
	Month:                 []string{"buwan"},
	Week:                  []string{"linggo"},
	Day:                   []string{"araw"},
	Hour:                  []string{"oras"},
	Minute:                []string{"min", "minuto"},
	Second:                []string{"seg", "segundo"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ngayong araw`},
		`0 hour ago`:   {`ngayong oras`},
		`0 minute ago`: {`sa minutong ito`},
		`0 month ago`:  {`ngayong buwan`},
		`0 second ago`: {`ngayon`},
		`0 week ago`:   {`ngayong linggo`, `sa linggong ito`},
		`0 year ago`:   {`ngayong taon`},
		`1 day ago`:    {`kahapon`},
		`1 month ago`:  {`nakaraang buwan`},
		`1 week ago`:   {`nakalipas na linggo`, `nakaraang linggo`},
		`1 year ago`:   {`nakaraang taon`},
		`in 1 day`:     {`bukas`},
		`in 1 month`:   {`susunod na buwan`},
		`in 1 week`:    {`susunod na linggo`},
		`in 1 year`:    {`susunod na taon`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) araw ang nakalipas`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) oras ang nakalipas`),
			regexp.MustCompile(`(?i)(\d+) oras nakalipas`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min ang nakalipas`),
			regexp.MustCompile(`(?i)(\d+) minuto ang nakalipas`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) buwan ang nakalipas`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) seg ang nakalipas`),
			regexp.MustCompile(`(?i)(\d+) seg nakalipas`),
			regexp.MustCompile(`(?i)(\d+) segundo ang nakalipas`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) linggo ang nakalipas`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) taon ang nakalipas`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)sa (\d+) araw`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)sa (\d+) oras`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)sa (\d+) min`),
			regexp.MustCompile(`(?i)sa (\d+) minuto`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)sa (\d+) buwan`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)sa (\d+) seg`),
			regexp.MustCompile(`(?i)sa (\d+) segundo`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)sa (\d+) linggo`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)sa (\d+) taon`),
		},
	},
}
