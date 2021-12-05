package data

import "regexp"

var gswLocale = LocaleData{
	Name:                  "gsw",
	DateOrder:             "DMY",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mär", "märz"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mai"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "auguscht"},
	September:             []string{"sep", "septämber"},
	October:               []string{"okt", "oktoober"},
	November:              []string{"nov", "novämber"},
	December:              []string{"dez", "dezämber"},
	Monday:                []string{"mä", "määntig"},
	Tuesday:               []string{"zi", "ziischtig"},
	Wednesday:             []string{"mi", "mittwuch"},
	Thursday:              []string{"du", "dunschtig"},
	Friday:                []string{"fr", "friitig"},
	Saturday:              []string{"sa", "samschtig"},
	Sunday:                []string{"su", "sunntig"},
	AM:                    []string{"am vormittag", "vorm", "vormittag"},
	PM:                    []string{"am namittag", "nam", "namittag"},
	Year:                  []string{"jaar"},
	Month:                 []string{"monet"},
	Week:                  []string{"wuche"},
	Day:                   []string{"tag"},
	Hour:                  []string{"schtund"},
	Minute:                []string{"minuute"},
	Second:                []string{"sekunde"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hüt`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`geschter`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`moorn`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"gsw-FR": gswFRLocale,
		"gsw-LI": gswLILocale,
	},
}

var gswFRLocale = LocaleData{
	Name:                  "gsw-FR",
	DateOrder:             "",
}

var gswLILocale = LocaleData{
	Name:                  "gsw-LI",
	DateOrder:             "",
}
