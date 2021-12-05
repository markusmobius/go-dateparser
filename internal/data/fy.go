package data

import "regexp"

var fyLocale = LocaleData{
	Name:                  "fy",
	DateOrder:             "DMY",
	January:               []string{"jan", "jannewaris"},
	February:              []string{"feb", "febrewaris"},
	March:                 []string{"maart", "mrt"},
	April:                 []string{"apr", "april"},
	May:                   []string{"maaie", "mai"},
	June:                  []string{"jun", "juny"},
	July:                  []string{"jul", "july"},
	August:                []string{"aug", "augustus"},
	September:             []string{"sep", "septimber"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "novimber"},
	December:              []string{"des", "desimber"},
	Monday:                []string{"mo", "moandei"},
	Tuesday:               []string{"ti", "tiisdei"},
	Wednesday:             []string{"wo", "woansdei"},
	Thursday:              []string{"to", "tongersdei"},
	Friday:                []string{"fr", "freed"},
	Saturday:              []string{"sneon", "so"},
	Sunday:                []string{"si", "snein"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"jier"},
	Month:                 []string{"moanne"},
	Week:                  []string{"wike"},
	Day:                   []string{"dei"},
	Hour:                  []string{"oere"},
	Minute:                []string{"minút"},
	Second:                []string{"sekonde"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`vandaag`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`dizze moanne`},
		`0 second ago`: {`nu`},
		`0 week ago`:   {`dizze wike`},
		`0 year ago`:   {`dit jier`},
		`1 day ago`:    {`gisteren`},
		`1 month ago`:  {`foarige moanne`},
		`1 week ago`:   {`foarige wike`},
		`1 year ago`:   {`foarich jier`},
		`in 1 day`:     {`morgen`},
		`in 1 month`:   {`folgjende moanne`},
		`in 1 week`:    {`folgjende wike`},
		`in 1 year`:    {`folgjend jier`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) dei lyn`),
			regexp.MustCompile(`(?i)(\d+) deien lyn`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) oere lyn`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) minuten lyn`),
			regexp.MustCompile(`(?i)(\d+) minút lyn`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) moanne lyn`),
			regexp.MustCompile(`(?i)(\d+) moannen lyn`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sekonde lyn`),
			regexp.MustCompile(`(?i)(\d+) sekonden lyn`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) wike lyn`),
			regexp.MustCompile(`(?i)(\d+) wiken lyn`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) jier lyn`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)oer (\d+) dei`),
			regexp.MustCompile(`(?i)oer (\d+) deien`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)oer (\d+) oere`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)oer (\d+) minuten`),
			regexp.MustCompile(`(?i)oer (\d+) minút`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)oer (\d+) moanne`),
			regexp.MustCompile(`(?i)oer (\d+) moannen`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)oer (\d+) sekonde`),
			regexp.MustCompile(`(?i)oer (\d+) sekonden`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)oer (\d+) wike`),
			regexp.MustCompile(`(?i)oer (\d+) wiken`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)oer (\d+) jier`),
		},
	},
}
