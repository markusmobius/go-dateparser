package data

import "regexp"

var nlLocale = LocaleData{
	Name:                  "nl",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "om", "|", "，"},
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"maart", "mrt"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mei"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "augustus"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"dec", "december"},
	Monday:                []string{"ma", "maandag"},
	Tuesday:               []string{"di", "dinsdag"},
	Wednesday:             []string{"wo", "woensdag"},
	Thursday:              []string{"do", "donderdag"},
	Friday:                []string{"vr", "vrijdag"},
	Saturday:              []string{"za", "zaterdag"},
	Sunday:                []string{"zo", "zondag"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"jaar", "jr"},
	Month:                 []string{"maand", "maanden", "mnd"},
	Week:                  []string{"week", "weken", "wk"},
	Day:                   []string{"dag", "dagen"},
	Hour:                  []string{"uur"},
	Minute:                []string{"min", "minuten", "minuut"},
	Second:                []string{"s", "sec", "seconde", "seconden"},
	In:                    []string{"in"},
	Ago:                   []string{"geleden"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`vandaag`},
		`0 hour ago`:   {`binnen een uur`},
		`0 minute ago`: {`binnen een minuut`},
		`0 month ago`:  {`deze maand`},
		`0 second ago`: {`nu`},
		`0 week ago`:   {`deze week`},
		`0 year ago`:   {`dit jaar`},
		`1 day ago`:    {`gisteren`},
		`1 month ago`:  {`vorige maand`},
		`1 week ago`:   {`vorige week`},
		`1 year ago`:   {`vorig jaar`, `vorige jaar`},
		`2 day ago`:    {`eergisteren`},
		`in 1 day`:     {`morgen`},
		`in 1 month`:   {`volgende maand`},
		`in 1 week`:    {`volgende week`},
		`in 1 year`:    {`volgend jaar`},
		`in 2 day`:     {`overmorgen`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) dag geleden`),
			regexp.MustCompile(`(?i)(\d+) dagen geleden`),
			regexp.MustCompile(`(?i)(\d+) dgn geleden`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) uur geleden`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min geleden`),
			regexp.MustCompile(`(?i)(\d+) minuten geleden`),
			regexp.MustCompile(`(?i)(\d+) minuut geleden`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) maand geleden`),
			regexp.MustCompile(`(?i)(\d+) maanden geleden`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sec geleden`),
			regexp.MustCompile(`(?i)(\d+) seconde geleden`),
			regexp.MustCompile(`(?i)(\d+) seconden geleden`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) week geleden`),
			regexp.MustCompile(`(?i)(\d+) weken geleden`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) jaar geleden`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)over (\d+) dag`),
			regexp.MustCompile(`(?i)over (\d+) dagen`),
			regexp.MustCompile(`(?i)over (\d+) dgn`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)over (\d+) uur`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)over (\d+) min`),
			regexp.MustCompile(`(?i)over (\d+) minuten`),
			regexp.MustCompile(`(?i)over (\d+) minuut`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)over (\d+) maand`),
			regexp.MustCompile(`(?i)over (\d+) maanden`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)over (\d+) sec`),
			regexp.MustCompile(`(?i)over (\d+) seconde`),
			regexp.MustCompile(`(?i)over (\d+) seconden`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)over (\d+) week`),
			regexp.MustCompile(`(?i)over (\d+) weken`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)over (\d+) jaar`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"nl-AW": nlAWLocale,
		"nl-BE": nlBELocale,
		"nl-BQ": nlBQLocale,
		"nl-CW": nlCWLocale,
		"nl-SR": nlSRLocale,
		"nl-SX": nlSXLocale,
	},
}

var nlAWLocale = LocaleData{
	Name:                  "nl-AW",
	DateOrder:             "",
}

var nlBELocale = LocaleData{
	Name:                  "nl-BE",
	DateOrder:             "",
}

var nlBQLocale = LocaleData{
	Name:                  "nl-BQ",
	DateOrder:             "",
}

var nlCWLocale = LocaleData{
	Name:                  "nl-CW",
	DateOrder:             "",
}

var nlSRLocale = LocaleData{
	Name:                  "nl-SR",
	DateOrder:             "",
}

var nlSXLocale = LocaleData{
	Name:                  "nl-SX",
	DateOrder:             "",
}
