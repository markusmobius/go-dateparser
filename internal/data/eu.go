package data

import "regexp"

var euLocale = LocaleData{
	Name:                  "eu",
	DateOrder:             "YMD",
	January:               []string{"urt", "urtarrila"},
	February:              []string{"ots", "otsaila"},
	March:                 []string{"mar", "martxoa"},
	April:                 []string{"api", "apirila"},
	May:                   []string{"mai", "maiatza"},
	June:                  []string{"eka", "ekaina"},
	July:                  []string{"uzt", "uztaila"},
	August:                []string{"abu", "abuztua"},
	September:             []string{"ira", "iraila"},
	October:               []string{"urr", "urria"},
	November:              []string{"aza", "azaroa"},
	December:              []string{"abe", "abendua"},
	Monday:                []string{"al", "astelehena"},
	Tuesday:               []string{"ar", "asteartea"},
	Wednesday:             []string{"asteazkena", "az"},
	Thursday:              []string{"og", "osteguna"},
	Friday:                []string{"or", "ostirala"},
	Saturday:              []string{"larunbata", "lr"},
	Sunday:                []string{"ig", "igandea"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"urtea"},
	Month:                 []string{"hil", "hilabetea"},
	Week:                  []string{"ast", "astea"},
	Day:                   []string{"eg", "eguna"},
	Hour:                  []string{"h", "ordua"},
	Minute:                []string{"min", "minutua"},
	Second:                []string{"s", "segundoa"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`gaur`},
		`0 hour ago`:   {`ordu honetan`},
		`0 minute ago`: {`minutu honetan`},
		`0 month ago`:  {`hilabete hau`},
		`0 second ago`: {`orain`},
		`0 week ago`:   {`aste hau`},
		`0 year ago`:   {`aurten`},
		`1 day ago`:    {`atzo`},
		`1 month ago`:  {`aurreko hilabetea`},
		`1 week ago`:   {`aurreko astea`},
		`1 year ago`:   {`aurreko urtea`},
		`in 1 day`:     {`bihar`},
		`in 1 month`:   {`hurrengo hilabetea`},
		`in 1 week`:    {`hurrengo astea`},
		`in 1 year`:    {`hurrengo urtea`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)duela (\d+) egun`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)duela (\d+) ordu`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)duela (\d+) minutu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)duela (\d+) hilabete`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)duela (\d+) segundo`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)duela (\d+) aste`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)duela (\d+) urte`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) egun barru`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ordu barru`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) minutu barru`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) hilabete barru`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) segundo barru`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) aste barru`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) urte barru`),
		},
	},
}
