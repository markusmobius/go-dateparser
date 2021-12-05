package data

import "regexp"

var huLocale = LocaleData{
	Name:                  "hu",
	DateOrder:             "YMD.",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "-a", "-ai", "-akor", "-e", "-ei", "-ekor", "-es", "-i", "-ig", "-je", "-jei", "-ji", "-kor", "-tól", "-től", "-áig", "-án", "-ától", "-éig", "-én", "-étől", "-ös", "/", ";", "@", "[", "]", "|", "，"},
	January:               []string{"i", "jan", "január"},
	February:              []string{"feb", "febr", "február", "ii"},
	March:                 []string{"iii", "már", "márc", "március"},
	April:                 []string{"iv", "ápr", "április"},
	May:                   []string{"máj", "május", "v"},
	June:                  []string{"jún", "június", "vi"},
	July:                  []string{"júl", "július", "vii"},
	August:                []string{"aug", "augusztus", "viii"},
	September:             []string{"ix", "szept", "szeptember"},
	October:               []string{"okt", "október", "x"},
	November:              []string{"nov", "november", "xi"},
	December:              []string{"dec", "december", "xii"},
	Monday:                []string{"h", "hétfő"},
	Tuesday:               []string{"k", "kedd"},
	Wednesday:             []string{"sze", "szerda"},
	Thursday:              []string{"cs", "csütörtök"},
	Friday:                []string{"p", "péntek"},
	Saturday:              []string{"szo", "szombat"},
	Sunday:                []string{"v", "vas", "vasárnap"},
	AM:                    []string{"de"},
	PM:                    []string{"du"},
	Year:                  []string{"év", "éve", "évek", "évvel"},
	Month:                 []string{"hó", "hónap", "hónapja", "hónapok", "hónappal"},
	Week:                  []string{"hete", "hetek", "hét", "héttel"},
	Day:                   []string{"nap", "napja", "napok", "nappal"},
	Hour:                  []string{"ó", "óra", "óráig", "órája", "órák", "órától", "órával"},
	Minute:                []string{"p", "perc", "perccel", "perce", "percek", "percig", "perctől"},
	Second:                []string{"mp", "másodperc", "másodperccel", "másodperce", "másodpercek", "másodpercig", "másodperctől"},
	In:                    []string{"múlva"},
	Ago:                   []string{"ezelőtt"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ma`},
		`0 hour ago`:   {`ebben az órában`},
		`0 minute ago`: {`ebben a percben`},
		`0 month ago`:  {`ez a hónap`},
		`0 second ago`: {`most`},
		`0 week ago`:   {`ez a hét`},
		`0 year ago`:   {`ez az év`},
		`1 day ago`:    {`tegnap`},
		`1 month ago`:  {`előző hónap`},
		`1 week ago`:   {`előző hét`},
		`1 year ago`:   {`előző év`},
		`2 day ago`:    {`tegnapelőtt`},
		`in 1 day`:     {`holnap`},
		`in 1 month`:   {`következő hónap`},
		`in 1 week`:    {`következő hét`},
		`in 1 year`:    {`következő év`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) napja`),
			regexp.MustCompile(`(?i)(\d+) nappal ezelőtt`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) órával ezelőtt`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) perccel ezelőtt`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) hónappal ezelőtt`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) másodperccel ezelőtt`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) héttel ezelőtt`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) évvel ezelőtt`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) nap múlva`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) óra múlva`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) perc múlva`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) hónap múlva`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) másodperc múlva`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) hét múlva`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) év múlva`),
		},
	},
	Simplifications: map[string]string{
		`egy`: `1`,
	},
}
