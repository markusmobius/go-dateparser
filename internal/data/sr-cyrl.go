package data

import "regexp"

var srCyrlLocale = LocaleData{
	Name:                  "sr-Cyrl",
	DateOrder:             "DMY.",
	January:               []string{"јан", "јануар"},
	February:              []string{"феб", "фебруар"},
	March:                 []string{"мар", "март"},
	April:                 []string{"апр", "април"},
	May:                   []string{"мај"},
	June:                  []string{"јун"},
	July:                  []string{"јул"},
	August:                []string{"авг", "август"},
	September:             []string{"сеп", "септембар"},
	October:               []string{"окт", "октобар"},
	November:              []string{"нов", "новембар"},
	December:              []string{"дец", "децембар"},
	Monday:                []string{"пон", "понедељак"},
	Tuesday:               []string{"уто", "уторак"},
	Wednesday:             []string{"сре", "среда"},
	Thursday:              []string{"чет", "четвртак"},
	Friday:                []string{"пет", "петак"},
	Saturday:              []string{"суб", "субота"},
	Sunday:                []string{"нед", "недеља"},
	AM:                    []string{"пре подне"},
	PM:                    []string{"по подне"},
	Year:                  []string{"г", "год", "година"},
	Month:                 []string{"м", "мес", "месец"},
	Week:                  []string{"н", "нед", "недеља"},
	Day:                   []string{"д", "дан"},
	Hour:                  []string{"сат", "ч"},
	Minute:                []string{"мин", "минут"},
	Second:                []string{"с", "сек", "секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`данас`},
		`0 hour ago`:   {`овог сата`},
		`0 minute ago`: {`овог минута`},
		`0 month ago`:  {`овог месеца`},
		`0 second ago`: {`сада`},
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
			regexp.MustCompile(`(?i)пре (\d+) д`),
			regexp.MustCompile(`(?i)пре (\d+) дана`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)пре (\d+) сата`),
			regexp.MustCompile(`(?i)пре (\d+) сати`),
			regexp.MustCompile(`(?i)пре (\d+) ч`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)пре (\d+) мин`),
			regexp.MustCompile(`(?i)пре (\d+) минута`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)пре (\d+) м`),
			regexp.MustCompile(`(?i)пре (\d+) мес`),
			regexp.MustCompile(`(?i)пре (\d+) месеца`),
			regexp.MustCompile(`(?i)пре (\d+) месеци`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)пре (\d+) с`),
			regexp.MustCompile(`(?i)пре (\d+) сек`),
			regexp.MustCompile(`(?i)пре (\d+) секунде`),
			regexp.MustCompile(`(?i)пре (\d+) секунди`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)пре (\d+) н`),
			regexp.MustCompile(`(?i)пре (\d+) нед`),
			regexp.MustCompile(`(?i)пре (\d+) недеља`),
			regexp.MustCompile(`(?i)пре (\d+) недеље`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)пре (\d+) г`),
			regexp.MustCompile(`(?i)пре (\d+) год`),
			regexp.MustCompile(`(?i)пре (\d+) година`),
			regexp.MustCompile(`(?i)пре (\d+) године`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)за (\d+) д`),
			regexp.MustCompile(`(?i)за (\d+) дан`),
			regexp.MustCompile(`(?i)за (\d+) дана`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)за (\d+) сат`),
			regexp.MustCompile(`(?i)за (\d+) сати`),
			regexp.MustCompile(`(?i)за (\d+) ч`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)за (\d+) мин`),
			regexp.MustCompile(`(?i)за (\d+) минут`),
			regexp.MustCompile(`(?i)за (\d+) минута`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)за (\d+) м`),
			regexp.MustCompile(`(?i)за (\d+) мес`),
			regexp.MustCompile(`(?i)за (\d+) месец`),
			regexp.MustCompile(`(?i)за (\d+) месеци`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)за (\d+) с`),
			regexp.MustCompile(`(?i)за (\d+) сек`),
			regexp.MustCompile(`(?i)за (\d+) секунди`),
			regexp.MustCompile(`(?i)за (\d+) секунду`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)за (\d+) н`),
			regexp.MustCompile(`(?i)за (\d+) нед`),
			regexp.MustCompile(`(?i)за (\d+) недеља`),
			regexp.MustCompile(`(?i)за (\d+) недељу`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)за (\d+) г`),
			regexp.MustCompile(`(?i)за (\d+) год`),
			regexp.MustCompile(`(?i)за (\d+) година`),
			regexp.MustCompile(`(?i)за (\d+) годину`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"sr-Cyrl-BA": srCyrlBALocale,
		"sr-Cyrl-ME": srCyrlMELocale,
		"sr-Cyrl-XK": srCyrlXKLocale,
	},
}

var srCyrlMELocale = LocaleData{
	Name:                  "sr-Cyrl-ME",
	DateOrder:             "",
	September:             []string{"септ"},
	Tuesday:               []string{"ут"},
	Wednesday:             []string{"ср", "сриједа"},
	Sunday:                []string{"недјеља"},
	AM:                    []string{"прије подне"},
}

var srCyrlXKLocale = LocaleData{
	Name:                  "sr-Cyrl-XK",
	DateOrder:             "",
	September:             []string{"септ"},
	Tuesday:               []string{"ут"},
	Wednesday:             []string{"ср"},
}

var srCyrlBALocale = LocaleData{
	Name:                  "sr-Cyrl-BA",
	DateOrder:             "",
	September:             []string{"септ"},
	Tuesday:               []string{"ут"},
	Wednesday:             []string{"ср", "сриједа"},
	Sunday:                []string{"недјеља"},
	AM:                    []string{"прије подне"},
}
