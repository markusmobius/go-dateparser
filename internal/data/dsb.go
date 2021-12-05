package data

import "regexp"

var dsbLocale = LocaleData{
	Name:                  "dsb",
	DateOrder:             "DMY",
	January:               []string{"jan", "januar", "januara"},
	February:              []string{"feb", "februar", "februara"},
	March:                 []string{"měr", "měrc", "měrca"},
	April:                 []string{"apr", "apryl", "apryla"},
	May:                   []string{"maj", "maja"},
	June:                  []string{"jun", "junij", "junija"},
	July:                  []string{"jul", "julij", "julija"},
	August:                []string{"awg", "awgust", "awgusta"},
	September:             []string{"sep", "september", "septembra"},
	October:               []string{"okt", "oktober", "oktobra"},
	November:              []string{"now", "nowember", "nowembra"},
	December:              []string{"dec", "december", "decembra"},
	Monday:                []string{"pón", "pónjeźele"},
	Tuesday:               []string{"wał", "wałtora"},
	Wednesday:             []string{"srj", "srjoda"},
	Thursday:              []string{"stw", "stwórtk"},
	Friday:                []string{"pět", "pětk"},
	Saturday:              []string{"sob", "sobota"},
	Sunday:                []string{"nje", "njeźela"},
	AM:                    []string{"dopołdnja"},
	PM:                    []string{"wótpołdnja"},
	Year:                  []string{"l", "lěto"},
	Month:                 []string{"mjas", "mjasec"},
	Week:                  []string{"tyź", "tyźeń"},
	Day:                   []string{"ź", "źeń"},
	Hour:                  []string{"g", "góź", "góźina"},
	Minute:                []string{"m", "min", "minuta"},
	Second:                []string{"s", "sek", "sekunda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`źinsa`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`ten mjasec`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`ten tyźeń`},
		`0 year ago`:   {`lětosa`},
		`1 day ago`:    {`cora`},
		`1 month ago`:  {`slědny mjasec`},
		`1 week ago`:   {`slědny tyźeń`},
		`1 year ago`:   {`łoni`},
		`in 1 day`:     {`witśe`},
		`in 1 month`:   {`pśiducy mjasec`},
		`in 1 week`:    {`pśiducy tyźeń`},
		`in 1 year`:    {`znowa`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) d`),
			regexp.MustCompile(`(?i)pśed (\d+) dnj`),
			regexp.MustCompile(`(?i)pśed (\d+) dnjami`),
			regexp.MustCompile(`(?i)pśed (\d+) dnjom`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) g`),
			regexp.MustCompile(`(?i)pśed (\d+) góź`),
			regexp.MustCompile(`(?i)pśed (\d+) góźinami`),
			regexp.MustCompile(`(?i)pśed (\d+) góźinu`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) m`),
			regexp.MustCompile(`(?i)pśed (\d+) min`),
			regexp.MustCompile(`(?i)pśed (\d+) minutami`),
			regexp.MustCompile(`(?i)pśed (\d+) minutu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) mjas`),
			regexp.MustCompile(`(?i)pśed (\d+) mjasecami`),
			regexp.MustCompile(`(?i)pśed (\d+) mjasecom`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) s`),
			regexp.MustCompile(`(?i)pśed (\d+) sek`),
			regexp.MustCompile(`(?i)pśed (\d+) sekundami`),
			regexp.MustCompile(`(?i)pśed (\d+) sekundu`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) tyź`),
			regexp.MustCompile(`(?i)pśed (\d+) tyźenjami`),
			regexp.MustCompile(`(?i)pśed (\d+) tyźenjom`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)pśed (\d+) l`),
			regexp.MustCompile(`(?i)pśed (\d+) lětami`),
			regexp.MustCompile(`(?i)pśed (\d+) lětom`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)za (\d+) dnj`),
			regexp.MustCompile(`(?i)za (\d+) dnjow`),
			regexp.MustCompile(`(?i)za (\d+) ź`),
			regexp.MustCompile(`(?i)za (\d+) źeń`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)za (\d+) g`),
			regexp.MustCompile(`(?i)za (\d+) góź`),
			regexp.MustCompile(`(?i)za (\d+) góźin`),
			regexp.MustCompile(`(?i)za (\d+) góźinu`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)za (\d+) m`),
			regexp.MustCompile(`(?i)za (\d+) min`),
			regexp.MustCompile(`(?i)za (\d+) minutow`),
			regexp.MustCompile(`(?i)za (\d+) minutu`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)za (\d+) mjas`),
			regexp.MustCompile(`(?i)za (\d+) mjasec`),
			regexp.MustCompile(`(?i)za (\d+) mjasecow`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)za (\d+) s`),
			regexp.MustCompile(`(?i)za (\d+) sek`),
			regexp.MustCompile(`(?i)za (\d+) sekundow`),
			regexp.MustCompile(`(?i)za (\d+) sekundu`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)za (\d+) tyź`),
			regexp.MustCompile(`(?i)za (\d+) tyźenjow`),
			regexp.MustCompile(`(?i)za (\d+) tyźeń`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)za (\d+) l`),
			regexp.MustCompile(`(?i)za (\d+) lět`),
			regexp.MustCompile(`(?i)za (\d+) lěto`),
		},
	},
}
