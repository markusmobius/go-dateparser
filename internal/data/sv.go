package data

import "regexp"

var svLocale = LocaleData{
	Name:                  "sv",
	DateOrder:             "YMD",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "den", "på", "|", "，"},
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"mars"},
	April:                 []string{"apr", "april"},
	May:                   []string{"maj"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "augusti"},
	September:             []string{"sep", "sept", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"dec", "december"},
	Monday:                []string{"mån", "månd", "måndag"},
	Tuesday:               []string{"tis", "tisd", "tisdag"},
	Wednesday:             []string{"ons", "onsd", "onsdag"},
	Thursday:              []string{"tors", "torsd", "torsdag"},
	Friday:                []string{"fre", "fred", "fredag"},
	Saturday:              []string{"lör", "lörd", "lördag"},
	Sunday:                []string{"sön", "sönd", "söndag"},
	AM:                    []string{"fm", "förmiddag"},
	PM:                    []string{"eftermiddag", "em"},
	Year:                  []string{"år", "året"},
	Month:                 []string{"m", "mån", "månad", "månaden", "månader"},
	Week:                  []string{"v", "vecka", "veckor"},
	Day:                   []string{"dag", "dagar"},
	Hour:                  []string{"h", "t", "tim", "timmar", "timme"},
	Minute:                []string{"m", "min", "minut", "minuter"},
	Second:                []string{"s", "sek", "sekund", "sekunder"},
	In:                    []string{"från nu", "om"},
	Ago:                   []string{"sedan"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`i dag`, `idag`},
		`0 hour ago`:   {`denna timme`},
		`0 minute ago`: {`denna minut`},
		`0 month ago`:  {`denna mån`, `denna månad`},
		`0 second ago`: {`nu`},
		`0 week ago`:   {`denna v`, `denna vecka`},
		`0 year ago`:   {`i år`},
		`1 day ago`:    {`i går`, `igår`},
		`1 month ago`:  {`förra mån`, `förra månaden`},
		`1 week ago`:   {`förra v`, `förra veckan`},
		`1 year ago`:   {`förra året`, `i fjol`},
		`2 day ago`:    {`förrgår`},
		`in 1 day`:     {`i morgon`, `imorgon`},
		`in 1 month`:   {`nästa mån`, `nästa månad`},
		`in 1 week`:    {`nästa v`, `nästa vecka`},
		`in 1 year`:    {`nästa år`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)för (\d+) d sedan`),
			regexp.MustCompile(`(?i)för (\d+) dag sedan`),
			regexp.MustCompile(`(?i)för (\d+) dagar sedan`),
			regexp.MustCompile(`(?i)−(\d+) d`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)för (\d+) tim sedan`),
			regexp.MustCompile(`(?i)för (\d+) timmar sedan`),
			regexp.MustCompile(`(?i)för (\d+) timme sedan`),
			regexp.MustCompile(`(?i)−(\d+) h`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)för (\d+) min sedan`),
			regexp.MustCompile(`(?i)för (\d+) minut sedan`),
			regexp.MustCompile(`(?i)för (\d+) minuter sedan`),
			regexp.MustCompile(`(?i)−(\d+) min`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)för (\d+) mån sedan`),
			regexp.MustCompile(`(?i)för (\d+) månad sedan`),
			regexp.MustCompile(`(?i)för (\d+) månader sedan`),
			regexp.MustCompile(`(?i)−(\d+) mån`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)för (\d+) sek sedan`),
			regexp.MustCompile(`(?i)för (\d+) sekund sedan`),
			regexp.MustCompile(`(?i)för (\d+) sekunder sedan`),
			regexp.MustCompile(`(?i)−(\d+) s`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)för (\d+) v sedan`),
			regexp.MustCompile(`(?i)för (\d+) vecka sedan`),
			regexp.MustCompile(`(?i)för (\d+) veckor sedan`),
			regexp.MustCompile(`(?i)−(\d+) v`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)för (\d+) år sedan`),
			regexp.MustCompile(`(?i)−(\d+) år`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)om (\d+) d`),
			regexp.MustCompile(`(?i)om (\d+) dag`),
			regexp.MustCompile(`(?i)om (\d+) dagar`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)om (\d+) tim`),
			regexp.MustCompile(`(?i)om (\d+) timmar`),
			regexp.MustCompile(`(?i)om (\d+) timme`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)om (\d+) min`),
			regexp.MustCompile(`(?i)om (\d+) minut`),
			regexp.MustCompile(`(?i)om (\d+) minuter`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)om (\d+) mån`),
			regexp.MustCompile(`(?i)om (\d+) månad`),
			regexp.MustCompile(`(?i)om (\d+) månader`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)om (\d+) sek`),
			regexp.MustCompile(`(?i)om (\d+) sekund`),
			regexp.MustCompile(`(?i)om (\d+) sekunder`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)om (\d+) v`),
			regexp.MustCompile(`(?i)om (\d+) vecka`),
			regexp.MustCompile(`(?i)om (\d+) veckor`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)om (\d+) år`),
		},
	},
	Simplifications: map[string]string{
		`en`:  `1`,
		`två`: `2`,
	},
	LocaleSpecific: map[string]LocaleData{
		"sv-AX": svAXLocale,
		"sv-FI": svFILocale,
	},
}

var svAXLocale = LocaleData{
	Name:                  "sv-AX",
	DateOrder:             "",
}

var svFILocale = LocaleData{
	Name:                  "sv-FI",
	DateOrder:             "DMY",
}
