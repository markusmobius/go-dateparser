package data

import "regexp"

var seLocale = LocaleData{
	Name:                  "se",
	DateOrder:             "YMD",
	January:               []string{"ođđajagemánnu", "ođđj"},
	February:              []string{"guov", "guovvamánnu"},
	March:                 []string{"njuk", "njukčamánnu"},
	April:                 []string{"cuo", "cuoŋománnu"},
	May:                   []string{"mies", "miessemánnu"},
	June:                  []string{"geas", "geassemánnu"},
	July:                  []string{"suoi", "suoidnemánnu"},
	August:                []string{"borg", "borgemánnu"},
	September:             []string{"čakč", "čakčamánnu"},
	October:               []string{"golg", "golggotmánnu"},
	November:              []string{"skáb", "skábmamánnu"},
	December:              []string{"juov", "juovlamánnu"},
	Monday:                []string{"vuos", "vuossárga"},
	Tuesday:               []string{"maŋ", "maŋŋebárga"},
	Wednesday:             []string{"gask", "gaskavahkku"},
	Thursday:              []string{"duor", "duorasdat"},
	Friday:                []string{"bear", "bearjadat"},
	Saturday:              []string{"láv", "lávvardat"},
	Sunday:                []string{"sotn", "sotnabeaivi"},
	AM:                    []string{"ib", "iđitbeaivet", "iđitbeaivi"},
	PM:                    []string{"eahketbeaivet", "eahketbeaivi", "eb"},
	Year:                  []string{"jáhki"},
	Month:                 []string{"mánnu"},
	Week:                  []string{"váhkku"},
	Day:                   []string{"beaivi"},
	Hour:                  []string{"diibmu"},
	Minute:                []string{"minuhtta"},
	Second:                []string{"sekunda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`odne`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`ikte`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`ihttin`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) jándor árat`),
			regexp.MustCompile(`(?i)(\d+) jándora árat`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) diibmu árat`),
			regexp.MustCompile(`(?i)(\d+) diibmur árat`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) minuhta árat`),
			regexp.MustCompile(`(?i)(\d+) minuhtta árat`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mánotbadji árat`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sekunda árat`),
			regexp.MustCompile(`(?i)(\d+) sekundda árat`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) vahkku árat`),
			regexp.MustCompile(`(?i)(\d+) vahku árat`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) jahki árat`),
			regexp.MustCompile(`(?i)(\d+) jahkki árat`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) jándor maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) jándora maŋŋilit`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) diibmu maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) diibmur maŋŋilit`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) minuhta maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) minuhtta maŋŋilit`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) mánotbadji maŋŋilit`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) sekunda maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) sekundda maŋŋilit`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) vahkku maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) vahku maŋŋilit`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) jahki maŋŋilit`),
			regexp.MustCompile(`(?i)(\d+) jahkki maŋŋilit`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"se-FI": seFILocale,
		"se-SE": seSELocale,
	},
}

var seSELocale = LocaleData{
	Name:                  "se-SE",
	DateOrder:             "",
}

var seFILocale = LocaleData{
	Name:                  "se-FI",
	DateOrder:             "",
	Monday:                []string{"vuossárgga"},
	Tuesday:               []string{"maŋŋebárgga"},
	Wednesday:             []string{"gaskavahku"},
	Thursday:              []string{"duorastaga"},
	Friday:                []string{"bearjadaga"},
	Saturday:              []string{"lávvardaga"},
	Year:                  []string{"j", "jahki"},
	Week:                  []string{"vahkku"},
	RelativeType: map[string][]string{
		`0 year ago`: {`dán jagi`},
		`1 year ago`: {`mannan jagi`},
		`in 1 year`:  {`boahtte jagi`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) jagi árat`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) jagi siste`),
		},
	},
}
