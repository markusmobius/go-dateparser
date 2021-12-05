package data

import "regexp"

var keaLocale = LocaleData{
	Name:                  "kea",
	DateOrder:             "DMY",
	January:               []string{"jan", "janeru"},
	February:              []string{"feb", "febreru"},
	March:                 []string{"mar", "marsu"},
	April:                 []string{"abr", "abril"},
	May:                   []string{"mai", "maiu"},
	June:                  []string{"jun", "junhu"},
	July:                  []string{"jul", "julhu"},
	August:                []string{"ago", "agostu"},
	September:             []string{"set", "setenbru"},
	October:               []string{"otu", "otubru"},
	November:              []string{"nuv", "nuvenbru"},
	December:              []string{"diz", "dizenbru"},
	Monday:                []string{"sig", "sigunda-fera"},
	Tuesday:               []string{"ter", "tersa-fera"},
	Wednesday:             []string{"kua", "kuarta-fera"},
	Thursday:              []string{"kin", "kinta-fera"},
	Friday:                []string{"ses", "sesta-fera"},
	Saturday:              []string{"sab", "sabadu", "sábadu"},
	Sunday:                []string{"dum", "dumingu"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"anu"},
	Month:                 []string{"mes"},
	Week:                  []string{"sim", "simana"},
	Day:                   []string{"dia"},
	Hour:                  []string{"h", "ora"},
	Minute:                []string{"m", "min", "minutu"},
	Second:                []string{"s", "sig", "sigundu"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`oji`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`es mes li`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`es simana li`},
		`0 year ago`:   {`es anu li`},
		`1 day ago`:    {`onti`},
		`1 month ago`:  {`mes pasadu`},
		`1 week ago`:   {`simana pasadu`},
		`1 year ago`:   {`anu pasadu`},
		`in 1 day`:     {`manha`},
		`in 1 month`:   {`prósimu mes`},
		`in 1 week`:    {`prósimu simana`},
		`in 1 year`:    {`prósimu anu`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) dia`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) ora`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) m`),
			regexp.MustCompile(`(?i)a ten (\d+) min`),
			regexp.MustCompile(`(?i)a ten (\d+) minutu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) mes`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) s`),
			regexp.MustCompile(`(?i)a ten (\d+) sig`),
			regexp.MustCompile(`(?i)a ten (\d+) sigundu`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) sim`),
			regexp.MustCompile(`(?i)a ten (\d+) simana`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)a ten (\d+) anu`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)di li (\d+) dia`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)di li (\d+) ora`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)di li (\d+) m`),
			regexp.MustCompile(`(?i)di li (\d+) min`),
			regexp.MustCompile(`(?i)di li (\d+) minutu`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)di li (\d+) mes`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)di li (\d+) s`),
			regexp.MustCompile(`(?i)di li (\d+) sig`),
			regexp.MustCompile(`(?i)di li (\d+) sigundu`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)di li (\d+) sim`),
			regexp.MustCompile(`(?i)di li (\d+) simana`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)di li (\d+) anu`),
		},
	},
}
