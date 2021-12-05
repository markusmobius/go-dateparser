package data

import "regexp"

var lnLocale = LocaleData{
	Name:                  "ln",
	DateOrder:             "DMY",
	January:               []string{"sánzá ya yambo", "yan"},
	February:              []string{"fbl", "sánzá ya míbalé"},
	March:                 []string{"msi", "sánzá ya mísáto"},
	April:                 []string{"apl", "sánzá ya mínei"},
	May:                   []string{"mai", "sánzá ya mítáno"},
	June:                  []string{"sánzá ya motóbá", "yun"},
	July:                  []string{"sánzá ya nsambo", "yul"},
	August:                []string{"agt", "sánzá ya mwambe"},
	September:             []string{"stb", "sánzá ya libwa"},
	October:               []string{"sánzá ya zómi", "ɔtb"},
	November:              []string{"nvb", "sánzá ya zómi na mɔ̌kɔ́"},
	December:              []string{"dsb", "sánzá ya zómi na míbalé"},
	Monday:                []string{"mokɔlɔ mwa yambo", "ybo"},
	Tuesday:               []string{"mbl", "mokɔlɔ mwa míbalé"},
	Wednesday:             []string{"mokɔlɔ mwa mísáto", "mst"},
	Thursday:              []string{"min", "mokɔlɔ ya mínéi"},
	Friday:                []string{"mokɔlɔ ya mítáno", "mtn"},
	Saturday:              []string{"mps", "mpɔ́sɔ"},
	Sunday:                []string{"eye", "eyenga"},
	AM:                    []string{"ntɔ́ngɔ́"},
	PM:                    []string{"mpókwa"},
	Year:                  []string{"mobú"},
	Month:                 []string{"sánzá"},
	Week:                  []string{"pɔ́sɔ"},
	Day:                   []string{"mokɔlɔ"},
	Hour:                  []string{"ngonga"},
	Minute:                []string{"monúti"},
	Second:                []string{"sɛkɔ́ndɛ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`lɛlɔ́`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`lóbi elékí`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`lóbi ekoyâ`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"ln-AO": lnAOLocale,
		"ln-CF": lnCFLocale,
		"ln-CG": lnCGLocale,
	},
}

var lnCGLocale = LocaleData{
	Name:                  "ln-CG",
	DateOrder:             "",
}

var lnAOLocale = LocaleData{
	Name:                  "ln-AO",
	DateOrder:             "",
}

var lnCFLocale = LocaleData{
	Name:                  "ln-CF",
	DateOrder:             "",
}
