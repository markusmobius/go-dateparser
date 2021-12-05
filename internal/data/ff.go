package data

import "regexp"

var ffLocale = LocaleData{
	Name:                  "ff",
	DateOrder:             "DMY",
	January:               []string{"sii", "siilo"},
	February:              []string{"col", "colte"},
	March:                 []string{"mbo", "mbooy"},
	April:                 []string{"see", "seeɗto"},
	May:                   []string{"duu", "duujal"},
	June:                  []string{"kor", "korse"},
	July:                  []string{"mor", "morso"},
	August:                []string{"juk", "juko"},
	September:             []string{"siilto", "slt"},
	October:               []string{"yar", "yarkomaa"},
	November:              []string{"jol", "jolal"},
	December:              []string{"bow", "bowte"},
	Monday:                []string{"aaɓ", "aaɓnde"},
	Tuesday:               []string{"maw", "mawbaare"},
	Wednesday:             []string{"nje", "njeslaare"},
	Thursday:              []string{"naa", "naasaande"},
	Friday:                []string{"mawnde", "mwd"},
	Saturday:              []string{"hbi", "hoore-biir"},
	Sunday:                []string{"dew", "dewo"},
	AM:                    []string{"subaka"},
	PM:                    []string{"kikiiɗe"},
	Year:                  []string{"hitaande"},
	Month:                 []string{"lewru"},
	Week:                  []string{"yontere"},
	Day:                   []string{"ñalnde"},
	Hour:                  []string{"waktu"},
	Minute:                []string{"hoƴom"},
	Second:                []string{"majaango"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hannde`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`haŋki`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`jaŋngo`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"ff-CM": ffCMLocale,
		"ff-GN": ffGNLocale,
		"ff-MR": ffMRLocale,
	},
}

var ffCMLocale = LocaleData{
	Name:                  "ff-CM",
	DateOrder:             "",
}

var ffGNLocale = LocaleData{
	Name:                  "ff-GN",
	DateOrder:             "",
}

var ffMRLocale = LocaleData{
	Name:                  "ff-MR",
	DateOrder:             "",
}
