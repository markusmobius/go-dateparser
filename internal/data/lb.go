package data

import "regexp"

var lbLocale = LocaleData{
	Name:                  "lb",
	DateOrder:             "DMY",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mäe", "mäerz"},
	April:                 []string{"abr", "abrëll"},
	May:                   []string{"mee"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"dez", "dezember"},
	Monday:                []string{"méi", "méindeg"},
	Tuesday:               []string{"dën", "dënschdeg"},
	Wednesday:             []string{"mët", "mëttwoch"},
	Thursday:              []string{"don", "donneschdeg"},
	Friday:                []string{"fre", "freideg"},
	Saturday:              []string{"sam", "samschdeg"},
	Sunday:                []string{"son", "sonndeg"},
	AM:                    []string{"moies"},
	PM:                    []string{"nomëttes"},
	Year:                  []string{"j", "joer"},
	Month:                 []string{"m", "mount"},
	Week:                  []string{"w", "woch"},
	Day:                   []string{"d", "dag"},
	Hour:                  []string{"st", "stonn"},
	Minute:                []string{"min", "minutt"},
	Second:                []string{"sek", "sekonn"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`haut`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`dëse mount`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`dës woch`},
		`0 year ago`:   {`dëst joer`},
		`1 day ago`:    {`gëschter`},
		`1 month ago`:  {`leschte mount`},
		`1 week ago`:   {`lescht woch`},
		`1 year ago`:   {`lescht joer`},
		`in 1 day`:     {`muer`},
		`in 1 month`:   {`nächste mount`},
		`in 1 week`:    {`nächst woch`},
		`in 1 year`:    {`nächst joer`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)virun (\d+) d`),
			regexp.MustCompile(`(?i)virun (\d+) dag`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)virun (\d+) st`),
			regexp.MustCompile(`(?i)virun (\d+) stonn`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)virun (\d+) min`),
			regexp.MustCompile(`(?i)virun (\d+) minutt`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)virun (\d+) m`),
			regexp.MustCompile(`(?i)virun (\d+) mount`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)virun (\d+) sek`),
			regexp.MustCompile(`(?i)virun (\d+) sekonn`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)virun (\d+) w`),
			regexp.MustCompile(`(?i)virun (\d+) woch`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)virun (\d+) j`),
			regexp.MustCompile(`(?i)virun (\d+) joer`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)an (\d+) d`),
			regexp.MustCompile(`(?i)an (\d+) dag`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)an (\d+) st`),
			regexp.MustCompile(`(?i)an (\d+) stonn`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)an (\d+) min`),
			regexp.MustCompile(`(?i)an (\d+) minutt`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)an (\d+) m`),
			regexp.MustCompile(`(?i)an (\d+) mount`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)an (\d+) sek`),
			regexp.MustCompile(`(?i)an (\d+) sekonn`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)an (\d+) w`),
			regexp.MustCompile(`(?i)an (\d+) woch`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)an (\d+) j`),
			regexp.MustCompile(`(?i)an (\d+) joer`),
		},
	},
}
