package data

import "regexp"

var deLocale = LocaleData{
	Name:                  "de",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "etwa", "uhr", "um", "und", "|", "，"},
	January:               []string{"jan", "januar", "jänner"},
	February:              []string{"feb", "feber", "februar"},
	March:                 []string{"mrz", "mär", "märz"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mai"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"dez", "dezember"},
	Monday:                []string{"mo", "mon", "montag"},
	Tuesday:               []string{"di", "die", "dienstag"},
	Wednesday:             []string{"mi", "mit", "mittwoch"},
	Thursday:              []string{"do", "don", "donnerstag"},
	Friday:                []string{"fr", "fre", "freitag"},
	Saturday:              []string{"sa", "sam", "samstag"},
	Sunday:                []string{"so", "son", "sonntag"},
	AM:                    []string{"vorm"},
	PM:                    []string{"nachm"},
	Year:                  []string{"j", "jahr", "jahre", "jahren"},
	Month:                 []string{"m", "monat", "monate", "monaten"},
	Week:                  []string{"w", "woche", "wochen"},
	Day:                   []string{"tag", "tage", "tagen"},
	Hour:                  []string{"std", "stunde", "stunden"},
	Minute:                []string{"min", "minute", "minuten"},
	Second:                []string{"sek", "sekunde", "sekunden"},
	In:                    []string{"im", "in"},
	Ago:                   []string{"vor"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`heute`},
		`0 hour ago`:   {`in dieser stunde`},
		`0 minute ago`: {`in dieser minute`},
		`0 month ago`:  {`diesen monat`},
		`0 second ago`: {`jetzt`},
		`0 week ago`:   {`diese woche`},
		`0 year ago`:   {`dieses jahr`},
		`1 day ago`:    {`gestern`},
		`1 month ago`:  {`letzten monat`},
		`1 week ago`:   {`letzte woche`},
		`1 year ago`:   {`letztes jahr`},
		`2 day ago`:    {`vorgestern`},
		`in 1 day`:     {`morgen`},
		`in 1 month`:   {`nächsten monat`},
		`in 1 week`:    {`nächste woche`},
		`in 1 year`:    {`nächstes jahr`},
		`in 2 day`:     {`übermorgen`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)vor (\d+) tag`),
			regexp.MustCompile(`(?i)vor (\d+) tagen`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)vor (\d+)\s*h`),
			regexp.MustCompile(`(?i)vor (\d+) std`),
			regexp.MustCompile(`(?i)vor (\d+) stunde`),
			regexp.MustCompile(`(?i)vor (\d+) stunden`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)vor (\d+)\s*m`),
			regexp.MustCompile(`(?i)vor (\d+) m`),
			regexp.MustCompile(`(?i)vor (\d+) min`),
			regexp.MustCompile(`(?i)vor (\d+) minute`),
			regexp.MustCompile(`(?i)vor (\d+) minuten`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)vor (\d+) monat`),
			regexp.MustCompile(`(?i)vor (\d+) monaten`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)vor (\d+)\s*s`),
			regexp.MustCompile(`(?i)vor (\d+) s`),
			regexp.MustCompile(`(?i)vor (\d+) sek`),
			regexp.MustCompile(`(?i)vor (\d+) sekunde`),
			regexp.MustCompile(`(?i)vor (\d+) sekunden`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)vor (\d+) wo`),
			regexp.MustCompile(`(?i)vor (\d+) woche`),
			regexp.MustCompile(`(?i)vor (\d+) wochen`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)vor (\d+) jahr`),
			regexp.MustCompile(`(?i)vor (\d+) jahren`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)in (\d+) tag`),
			regexp.MustCompile(`(?i)in (\d+) tagen`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)in (\d+) std`),
			regexp.MustCompile(`(?i)in (\d+) stunde`),
			regexp.MustCompile(`(?i)in (\d+) stunden`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)in (\d+) m`),
			regexp.MustCompile(`(?i)in (\d+) min`),
			regexp.MustCompile(`(?i)in (\d+) minute`),
			regexp.MustCompile(`(?i)in (\d+) minuten`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)in (\d+) monat`),
			regexp.MustCompile(`(?i)in (\d+) monaten`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)in (\d+) s`),
			regexp.MustCompile(`(?i)in (\d+) sek`),
			regexp.MustCompile(`(?i)in (\d+) sekunde`),
			regexp.MustCompile(`(?i)in (\d+) sekunden`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)in (\d+) wo`),
			regexp.MustCompile(`(?i)in (\d+) woche`),
			regexp.MustCompile(`(?i)in (\d+) wochen`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)in (\d+) jahr`),
			regexp.MustCompile(`(?i)in (\d+) jahren`),
		},
	},
	Simplifications: map[string]string{
		`acht`:   `8`,
		`drei`:   `3`,
		`ein`:    `1`,
		`einem`:  `1`,
		`einer`:  `1`,
		`elf`:    `11`,
		`fünf`:   `5`,
		`neun`:   `9`,
		`sechs`:  `6`,
		`sieben`: `7`,
		`vier`:   `4`,
		`zehn`:   `10`,
		`zwei`:   `2`,
		`zwölf`:  `12`,
	},
	LocaleSpecific: map[string]LocaleData{
		"de-AT": deATLocale,
		"de-BE": deBELocale,
		"de-CH": deCHLocale,
		"de-IT": deITLocale,
		"de-LI": deLILocale,
		"de-LU": deLULocale,
	},
}

var deLILocale = LocaleData{
	Name:                  "de-LI",
	DateOrder:             "",
}

var deLULocale = LocaleData{
	Name:                  "de-LU",
	DateOrder:             "",
}

var deATLocale = LocaleData{
	Name:                  "de-AT",
	DateOrder:             "",
	January:               []string{"jän", "jänner"},
}

var deBELocale = LocaleData{
	Name:                  "de-BE",
	DateOrder:             "",
}

var deCHLocale = LocaleData{
	Name:                  "de-CH",
	DateOrder:             "",
}

var deITLocale = LocaleData{
	Name:                  "de-IT",
	DateOrder:             "",
	January:               []string{"jän", "jänner"},
}
