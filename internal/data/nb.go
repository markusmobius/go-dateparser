package data

import "regexp"

var nbLocale = LocaleData{
	Name:                  "nb",
	DateOrder:             "DMY",
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "，"},
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mar", "mars"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mai"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"des", "desember"},
	Monday:                []string{"man", "mandag"},
	Tuesday:               []string{"tir", "tirsdag"},
	Wednesday:             []string{"ons", "onsdag"},
	Thursday:              []string{"tor", "torsdag"},
	Friday:                []string{"fre", "fredag"},
	Saturday:              []string{"lør", "lørdag"},
	Sunday:                []string{"søn", "søndag"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"år"},
	Month:                 []string{"md", "mnd", "måned"},
	Week:                  []string{"u", "uke", "uker"},
	Day:                   []string{"d", "dag", "dager"},
	Hour:                  []string{"t", "time", "timer"},
	Minute:                []string{"m", "min", "minutt"},
	Second:                []string{"s", "sek", "sekund"},
	In:                    []string{"om"},
	Ago:                   []string{"siden"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`i dag`},
		`0 hour ago`:   {`denne timen`},
		`0 minute ago`: {`dette minuttet`},
		`0 month ago`:  {`denne md`, `denne måneden`},
		`0 second ago`: {`nå`},
		`0 week ago`:   {`denne uken`},
		`0 year ago`:   {`i år`},
		`1 day ago`:    {`i går`},
		`1 month ago`:  {`forrige md`, `forrige måned`},
		`1 week ago`:   {`forrige uke`},
		`1 year ago`:   {`i fjor`},
		`in 1 day`:     {`i morgen`},
		`in 1 month`:   {`neste md`, `neste måned`},
		`in 1 week`:    {`neste uke`},
		`in 1 year`:    {`neste år`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)for (\d+) dager siden`),
			regexp.MustCompile(`(?i)for (\d+) d siden`),
			regexp.MustCompile(`(?i)for (\d+) døgn siden`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)for (\d+) t siden`),
			regexp.MustCompile(`(?i)for (\d+) time siden`),
			regexp.MustCompile(`(?i)for (\d+) timer siden`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)for (\d+) min siden`),
			regexp.MustCompile(`(?i)for (\d+) minutt siden`),
			regexp.MustCompile(`(?i)for (\d+) minutter siden`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)for (\d+) md siden`),
			regexp.MustCompile(`(?i)for (\d+) måned siden`),
			regexp.MustCompile(`(?i)for (\d+) måneder siden`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)for (\d+) sek siden`),
			regexp.MustCompile(`(?i)for (\d+) sekund siden`),
			regexp.MustCompile(`(?i)for (\d+) sekunder siden`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)for (\d+) u siden`),
			regexp.MustCompile(`(?i)for (\d+) uke siden`),
			regexp.MustCompile(`(?i)for (\d+) uker siden`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)for (\d+) år siden`),
			regexp.MustCompile(`(?i)–(\d+) år`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)om (\d+) d`),
			regexp.MustCompile(`(?i)om (\d+) døgn`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)om (\d+) t`),
			regexp.MustCompile(`(?i)om (\d+) time`),
			regexp.MustCompile(`(?i)om (\d+) timer`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)om (\d+) min`),
			regexp.MustCompile(`(?i)om (\d+) minutt`),
			regexp.MustCompile(`(?i)om (\d+) minutter`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)om (\d+) md`),
			regexp.MustCompile(`(?i)om (\d+) måned`),
			regexp.MustCompile(`(?i)om (\d+) måneder`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)om (\d+) sek`),
			regexp.MustCompile(`(?i)om (\d+) sekund`),
			regexp.MustCompile(`(?i)om (\d+) sekunder`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)om (\d+) u`),
			regexp.MustCompile(`(?i)om (\d+) uke`),
			regexp.MustCompile(`(?i)om (\d+) uker`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)om (\d+) år`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"nb-SJ": nbSJLocale,
	},
}

var nbSJLocale = LocaleData{
	Name:                  "nb-SJ",
	DateOrder:             "",
}
