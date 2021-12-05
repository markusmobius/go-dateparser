package data

import "regexp"

var waeLocale = LocaleData{
	Name:                  "wae",
	DateOrder:             "YMD",
	January:               []string{"jen", "jenner"},
	February:              []string{"hor", "hornig"},
	March:                 []string{"mär", "märze"},
	April:                 []string{"abr", "abrille"},
	May:                   []string{"mei", "meije"},
	June:                  []string{"brá", "bráčet"},
	July:                  []string{"hei", "heiwet"},
	August:                []string{"öig", "öigšte"},
	September:             []string{"her", "herbštmánet"},
	October:               []string{"wím", "wímánet"},
	November:              []string{"win", "wintermánet"},
	December:              []string{"chr", "chrištmánet"},
	Monday:                []string{"män", "mäntag"},
	Tuesday:               []string{"ziš", "zištag"},
	Wednesday:             []string{"mit", "mittwuč"},
	Thursday:              []string{"fró", "fróntag"},
	Friday:                []string{"fri", "fritag"},
	Saturday:              []string{"sam", "samštag"},
	Sunday:                []string{"sun", "sunntag"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"jár"},
	Month:                 []string{"mánet"},
	Week:                  []string{"wuča"},
	Day:                   []string{"tag"},
	Hour:                  []string{"schtund"},
	Minute:                []string{"mínütta"},
	Second:                []string{"sekunda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hitte`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`gešter`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`móre`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)vor (\d+) tag`),
			regexp.MustCompile(`(?i)vor (\d+) täg`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)vor (\d+) stund`),
			regexp.MustCompile(`(?i)vor (\d+) stunde`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)vor (\d+) minüta`),
			regexp.MustCompile(`(?i)vor (\d+) minüte`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)vor (\d+) mánet`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)vor (\d+) sekund`),
			regexp.MustCompile(`(?i)vor (\d+) sekunde`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)cor (\d+) wučä`),
			regexp.MustCompile(`(?i)vor (\d+) wuča`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)cor (\d+) jár`),
			regexp.MustCompile(`(?i)vor (\d+) jár`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)i (\d+) tag`),
			regexp.MustCompile(`(?i)i (\d+) täg`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)i (\d+) stund`),
			regexp.MustCompile(`(?i)i (\d+) stunde`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)i (\d+) minüta`),
			regexp.MustCompile(`(?i)i (\d+) minüte`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)i (\d+) mánet`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)i (\d+) sekund`),
			regexp.MustCompile(`(?i)i (\d+) sekunde`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)i (\d+) wuča`),
			regexp.MustCompile(`(?i)i (\d+) wučä`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)i (\d+) jár`),
		},
	},
}
