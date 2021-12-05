package data

import "regexp"

var quLocale = LocaleData{
	Name:                  "qu",
	DateOrder:             "DMY",
	January:               []string{"qul", "qulla puquy"},
	February:              []string{"hat", "hatun puquy"},
	March:                 []string{"pau", "pauqar waray"},
	April:                 []string{"ayr", "ayriwa"},
	May:                   []string{"aym", "aymuray"},
	June:                  []string{"int", "inti raymi"},
	July:                  []string{"ant", "anta sitwa"},
	August:                []string{"qha", "qhapaq sitwa"},
	September:             []string{"uma", "uma raymi"},
	October:               []string{"kan", "kantaray"},
	November:              []string{"aya", "ayamarq'a"},
	December:              []string{"kap", "kapaq raymi"},
	Monday:                []string{"lun", "lunes"},
	Tuesday:               []string{"mar", "martes"},
	Wednesday:             []string{"mié", "miércoles"},
	Thursday:              []string{"jue", "jueves"},
	Friday:                []string{"vie", "viernes"},
	Saturday:              []string{"sab", "sábado"},
	Sunday:                []string{"dom", "domingo"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"year"},
	Month:                 []string{"month"},
	Week:                  []string{"week"},
	Day:                   []string{"day"},
	Hour:                  []string{"hour"},
	Minute:                []string{"minute"},
	Second:                []string{"second"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`today`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`yesterday`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`tomorrow`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"qu-BO": quBOLocale,
		"qu-EC": quECLocale,
	},
}

var quBOLocale = LocaleData{
	Name:                  "qu-BO",
	DateOrder:             "",
}

var quECLocale = LocaleData{
	Name:                  "qu-EC",
	DateOrder:             "",
}
