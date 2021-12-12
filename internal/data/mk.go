// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mk_Locale = LocaleData{
	Name:                  "mk",
	DateOrder:             "DMY",
	January:               []string{"јан", "јануари"},
	February:              []string{"фев", "февруари"},
	March:                 []string{"мар", "март"},
	April:                 []string{"апр", "април"},
	May:                   []string{"мај"},
	June:                  []string{"јун", "јуни"},
	July:                  []string{"јул", "јули"},
	August:                []string{"авг", "август"},
	September:             []string{"септ", "септември"},
	October:               []string{"окт", "октомври"},
	November:              []string{"ноем", "ноември"},
	December:              []string{"дек", "декември"},
	Monday:                []string{"пон", "понеделник"},
	Tuesday:               []string{"вт", "вто", "вторник"},
	Wednesday:             []string{"сре", "среда"},
	Thursday:              []string{"чет", "четврток"},
	Friday:                []string{"пет", "петок"},
	Saturday:              []string{"саб", "сабота"},
	Sunday:                []string{"нед", "недела"},
	AM:                    []string{"претпл", "претпладне"},
	PM:                    []string{"попл", "попладне"},
	Year:                  []string{"год", "година"},
	Month:                 []string{"мес", "месец"},
	Week:                  []string{"недела", "сед"},
	Day:                   []string{"ден"},
	Hour:                  []string{"час"},
	Minute:                []string{"мин", "минута"},
	Second:                []string{"сек", "секунда"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`денес`},
		`0 hour ago`:   {`часов`},
		`0 minute ago`: {`оваа минута`},
		`0 month ago`:  {`овој месец`},
		`0 second ago`: {`сега`},
		`0 week ago`:   {`оваа седмица`},
		`0 year ago`:   {`оваа година`},
		`1 day ago`:    {`вчера`},
		`1 month ago`:  {`минатиот месец`},
		`1 week ago`:   {`минатата седмица`},
		`1 year ago`:   {`минатата година`},
		`in 1 day`:     {`утре`},
		`in 1 month`:   {`следниот месец`},
		`in 1 week`:    {`следната седмица`},
		`in 1 year`:    {`следната година`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)пред (\d+) ден`),
			regexp.MustCompile(`(?i)пред (\d+) дена`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)пред (\d+) час`),
			regexp.MustCompile(`(?i)пред (\d+) часа`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)пред (\d+) минута`),
			regexp.MustCompile(`(?i)пред (\d+) минути`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)пред (\d+) месец`),
			regexp.MustCompile(`(?i)пред (\d+) месеци`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)пред (\d+) секунда`),
			regexp.MustCompile(`(?i)пред (\d+) секунди`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)пред (\d+) седмица`),
			regexp.MustCompile(`(?i)пред (\d+) седмици`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)пред (\d+) година`),
			regexp.MustCompile(`(?i)пред (\d+) години`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)за (\d+) ден`),
			regexp.MustCompile(`(?i)за (\d+) дена`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)за (\d+) час`),
			regexp.MustCompile(`(?i)за (\d+) часа`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)за (\d+) минута`),
			regexp.MustCompile(`(?i)за (\d+) минути`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)за (\d+) месец`),
			regexp.MustCompile(`(?i)за (\d+) месеци`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)за (\d+) секунда`),
			regexp.MustCompile(`(?i)за (\d+) секунди`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)за (\d+) седмица`),
			regexp.MustCompile(`(?i)за (\d+) седмици`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)за (\d+) година`),
			regexp.MustCompile(`(?i)за (\d+) години`),
		},
	},
}