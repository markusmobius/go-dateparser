package data

import "regexp"

var haLocale = LocaleData{
	Name:                  "ha",
	DateOrder:             "DMY",
	January:               []string{"jan", "janairu"},
	February:              []string{"fab", "faburairu"},
	March:                 []string{"mar", "maris"},
	April:                 []string{"afi", "afirilu"},
	May:                   []string{"may", "mayu"},
	June:                  []string{"yun", "yuni"},
	July:                  []string{"yul", "yuli"},
	August:                []string{"agu", "agusta"},
	September:             []string{"sat", "satumba"},
	October:               []string{"okt", "oktoba"},
	November:              []string{"nuw", "nuwamba"},
	December:              []string{"dis", "disamba"},
	Monday:                []string{"lit", "litinin"},
	Tuesday:               []string{"tal", "talata"},
	Wednesday:             []string{"lar", "laraba"},
	Thursday:              []string{"alh", "alhamis"},
	Friday:                []string{"jum", "jumma'a"},
	Saturday:              []string{"asa", "asabar"},
	Sunday:                []string{"lah", "lahadi"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"shekara"},
	Month:                 []string{"wata"},
	Week:                  []string{"mako"},
	Day:                   []string{"kwana"},
	Hour:                  []string{"awa"},
	Minute:                []string{"minti"},
	Second:                []string{"daƙiƙa"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`yau`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`jiya`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`gobe`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"ha-GH": haGHLocale,
		"ha-NE": haNELocale,
	},
}

var haNELocale = LocaleData{
	Name:                  "ha-NE",
	DateOrder:             "",
}

var haGHLocale = LocaleData{
	Name:                  "ha-GH",
	DateOrder:             "",
}
