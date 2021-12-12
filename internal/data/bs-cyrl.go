// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Cyrl_Locale = LocaleData{
	Name:                  "bs-Cyrl",
	DateOrder:             "DMY.",
	January:               []string{"јан", "јануар"},
	February:              []string{"феб", "фебруар"},
	March:                 []string{"мар", "март"},
	April:                 []string{"апр", "април"},
	May:                   []string{"мај"},
	June:                  []string{"јун", "јуни"},
	July:                  []string{"јул", "јули"},
	August:                []string{"авг", "август"},
	September:             []string{"сеп", "септембар"},
	October:               []string{"окт", "октобар"},
	November:              []string{"нов", "новембар"},
	December:              []string{"дец", "децембар"},
	Monday:                []string{"пон", "понедељак"},
	Tuesday:               []string{"уто", "уторак"},
	Wednesday:             []string{"сри", "сриједа"},
	Thursday:              []string{"чет", "четвртак"},
	Friday:                []string{"пет", "петак"},
	Saturday:              []string{"суб", "субота"},
	Sunday:                []string{"нед", "недеља"},
	AM:                    []string{"пре подне"},
	PM:                    []string{"поподне"},
	Year:                  []string{"година"},
	Month:                 []string{"месец"},
	Week:                  []string{"недеља"},
	Day:                   []string{"дан"},
	Hour:                  []string{"час"},
	Minute:                []string{"минут"},
	Second:                []string{"секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`данас`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`овог месеца`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`ове недеље`},
		`0 year ago`:   {`ове године`},
		`1 day ago`:    {`јуче`},
		`1 month ago`:  {`прошлог месеца`},
		`1 week ago`:   {`прошле недеље`},
		`1 year ago`:   {`прошле године`},
		`in 1 day`:     {`сутра`},
		`in 1 month`:   {`следећег месеца`},
		`in 1 week`:    {`следеће недеље`},
		`in 1 year`:    {`следеће године`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)пре (\d+) дан`),
			regexp.MustCompile(`(?i)пре (\d+) дана`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)пре (\d+) сат`),
			regexp.MustCompile(`(?i)пре (\d+) сати`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)пре (\d+) минут`),
			regexp.MustCompile(`(?i)пре (\d+) минута`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)пре (\d+) месец`),
			regexp.MustCompile(`(?i)пре (\d+) месеци`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)пре (\d+) секунд`),
			regexp.MustCompile(`(?i)пре (\d+) секунди`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)пре (\d+) недеља`),
			regexp.MustCompile(`(?i)пре (\d+) недељу`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)пре (\d+) година`),
			regexp.MustCompile(`(?i)пре (\d+) годину`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)за (\d+) дан`),
			regexp.MustCompile(`(?i)за (\d+) дана`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)за (\d+) сат`),
			regexp.MustCompile(`(?i)за (\d+) сати`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)за (\d+) минут`),
			regexp.MustCompile(`(?i)за (\d+) минута`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)за (\d+) месец`),
			regexp.MustCompile(`(?i)за (\d+) месеци`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)за (\d+) секунд`),
			regexp.MustCompile(`(?i)за (\d+) секунди`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)за (\d+) недеља`),
			regexp.MustCompile(`(?i)за (\d+) недељу`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)за (\d+) година`),
			regexp.MustCompile(`(?i)за (\d+) годину`),
		},
	},
}