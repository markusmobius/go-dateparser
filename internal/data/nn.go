package data

import "regexp"

var nnLocale = LocaleData{
	Name:                  "nn",
	DateOrder:             "DMY",
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
	Monday:                []string{"må", "mån", "måndag"},
	Tuesday:               []string{"ty", "tys", "tysdag"},
	Wednesday:             []string{"on", "ons", "onsdag"},
	Thursday:              []string{"to", "tor", "torsdag"},
	Friday:                []string{"fr", "fre", "fredag"},
	Saturday:              []string{"la", "lau", "laurdag"},
	Sunday:                []string{"sø", "søn", "søndag"},
	AM:                    []string{"fm", "formiddag"},
	PM:                    []string{"em", "ettermiddag"},
	Year:                  []string{"år"},
	Month:                 []string{"månad"},
	Week:                  []string{"veke"},
	Day:                   []string{"dag"},
	Hour:                  []string{"time"},
	Minute:                []string{"minutt"},
	Second:                []string{"sekund"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`i dag`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`i går`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`i morgon`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)for (\d+) døgn siden`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)for (\d+) time siden`),
			regexp.MustCompile(`(?i)for (\d+) timer siden`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)for (\d+) minutt siden`),
			regexp.MustCompile(`(?i)for (\d+) minutter siden`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)for (\d+) måned siden`),
			regexp.MustCompile(`(?i)for (\d+) måneder siden`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)for (\d+) sekund siden`),
			regexp.MustCompile(`(?i)for (\d+) sekunder siden`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)for (\d+) uke siden`),
			regexp.MustCompile(`(?i)for (\d+) uker siden`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)for (\d+) år siden`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)om (\d+) døgn`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)om (\d+) time`),
			regexp.MustCompile(`(?i)om (\d+) timer`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)om (\d+) minutt`),
			regexp.MustCompile(`(?i)om (\d+) minutter`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)om (\d+) måned`),
			regexp.MustCompile(`(?i)om (\d+) måneder`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)om (\d+) sekund`),
			regexp.MustCompile(`(?i)om (\d+) sekunder`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)om (\d+) uke`),
			regexp.MustCompile(`(?i)om (\d+) uker`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)om (\d+) år`),
		},
	},
}
