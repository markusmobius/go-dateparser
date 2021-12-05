package data

import "regexp"

var klLocale = LocaleData{
	Name:                  "kl",
	DateOrder:             "YMD",
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"mar", "martsi"},
	April:                 []string{"apr", "aprili"},
	May:                   []string{"maj", "maji"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "augustusi"},
	September:             []string{"sep", "septemberi"},
	October:               []string{"okt", "oktoberi"},
	November:              []string{"nov", "novemberi"},
	December:              []string{"dec", "decemberi"},
	Monday:                []string{"ata", "ataasinngorneq"},
	Tuesday:               []string{"mar", "marlunngorneq"},
	Wednesday:             []string{"pin", "pingasunngorneq"},
	Thursday:              []string{"sis", "sisamanngorneq"},
	Friday:                []string{"tal", "tallimanngorneq"},
	Saturday:              []string{"arf", "arfininngorneq"},
	Sunday:                []string{"sab", "sabaat"},
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
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)for (\d+) ulloq unnuarlu siden`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)for (\d+) nalunaaquttap-akunnera siden`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)for (\d+) minutsi siden`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)for (\d+) qaammat siden`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)for (\d+) sekundi siden`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)for (\d+) sapaatip-akunnera siden`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)for (\d+) ukioq siden`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)om (\d+) ulloq unnuarlu`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)om (\d+) nalunaaquttap-akunnera`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)om (\d+) minutsi`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)om (\d+) qaammat`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)om (\d+) sekundi`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)om (\d+) sapaatip-akunnera`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)om (\d+) ukioq`),
		},
	},
}
