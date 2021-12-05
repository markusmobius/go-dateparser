package data

import "regexp"

var srLatnLocale = LocaleData{
	Name:                  "sr-Latn",
	DateOrder:             "DMY.",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mar", "mart"},
	April:                 []string{"apr", "april"},
	May:                   []string{"maj"},
	June:                  []string{"jun"},
	July:                  []string{"jul"},
	August:                []string{"avg", "avgust"},
	September:             []string{"sep", "septembar"},
	October:               []string{"okt", "oktobar"},
	November:              []string{"nov", "novembar"},
	December:              []string{"dec", "decembar"},
	Monday:                []string{"pon", "ponedeljak"},
	Tuesday:               []string{"uto", "utorak"},
	Wednesday:             []string{"sre", "sreda"},
	Thursday:              []string{"čet", "četvrtak"},
	Friday:                []string{"pet", "petak"},
	Saturday:              []string{"sub", "subota"},
	Sunday:                []string{"ned", "nedelja"},
	AM:                    []string{"pre podne"},
	PM:                    []string{"po podne"},
	Year:                  []string{"g", "god", "godina"},
	Month:                 []string{"m", "mes", "mesec"},
	Week:                  []string{"n", "ned", "nedelja"},
	Day:                   []string{"d", "dan"},
	Hour:                  []string{"sat", "č"},
	Minute:                []string{"min", "minut"},
	Second:                []string{"s", "sek", "sekund"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`danas`},
		`0 hour ago`:   {`ovog sata`},
		`0 minute ago`: {`ovog minuta`},
		`0 month ago`:  {`ovog meseca`},
		`0 second ago`: {`sada`},
		`0 week ago`:   {`ove nedelje`},
		`0 year ago`:   {`ove godine`},
		`1 day ago`:    {`juče`},
		`1 month ago`:  {`prošlog meseca`},
		`1 week ago`:   {`prošle nedelje`},
		`1 year ago`:   {`prošle godine`},
		`in 1 day`:     {`sutra`},
		`in 1 month`:   {`sledećeg meseca`},
		`in 1 week`:    {`sledeće nedelje`},
		`in 1 year`:    {`sledeće godine`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)pre (\d+) d`),
			regexp.MustCompile(`(?i)pre (\d+) dana`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)pre (\d+) sata`),
			regexp.MustCompile(`(?i)pre (\d+) sati`),
			regexp.MustCompile(`(?i)pre (\d+) č`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)pre (\d+) min`),
			regexp.MustCompile(`(?i)pre (\d+) minuta`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)pre (\d+) m`),
			regexp.MustCompile(`(?i)pre (\d+) mes`),
			regexp.MustCompile(`(?i)pre (\d+) meseca`),
			regexp.MustCompile(`(?i)pre (\d+) meseci`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)pre (\d+) s`),
			regexp.MustCompile(`(?i)pre (\d+) sek`),
			regexp.MustCompile(`(?i)pre (\d+) sekunde`),
			regexp.MustCompile(`(?i)pre (\d+) sekundi`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)pre (\d+) n`),
			regexp.MustCompile(`(?i)pre (\d+) ned`),
			regexp.MustCompile(`(?i)pre (\d+) nedelja`),
			regexp.MustCompile(`(?i)pre (\d+) nedelje`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)pre (\d+) g`),
			regexp.MustCompile(`(?i)pre (\d+) god`),
			regexp.MustCompile(`(?i)pre (\d+) godina`),
			regexp.MustCompile(`(?i)pre (\d+) godine`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)za (\d+) d`),
			regexp.MustCompile(`(?i)za (\d+) dan`),
			regexp.MustCompile(`(?i)za (\d+) dana`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)za (\d+) sat`),
			regexp.MustCompile(`(?i)za (\d+) sati`),
			regexp.MustCompile(`(?i)za (\d+) č`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)za (\d+) min`),
			regexp.MustCompile(`(?i)za (\d+) minut`),
			regexp.MustCompile(`(?i)za (\d+) minuta`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)za (\d+) m`),
			regexp.MustCompile(`(?i)za (\d+) mes`),
			regexp.MustCompile(`(?i)za (\d+) mesec`),
			regexp.MustCompile(`(?i)za (\d+) meseci`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)za (\d+) s`),
			regexp.MustCompile(`(?i)za (\d+) sek`),
			regexp.MustCompile(`(?i)za (\d+) sekundi`),
			regexp.MustCompile(`(?i)za (\d+) sekundu`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)za (\d+) n`),
			regexp.MustCompile(`(?i)za (\d+) ned`),
			regexp.MustCompile(`(?i)za (\d+) nedelja`),
			regexp.MustCompile(`(?i)za (\d+) nedelju`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)za (\d+) g`),
			regexp.MustCompile(`(?i)za (\d+) god`),
			regexp.MustCompile(`(?i)za (\d+) godina`),
			regexp.MustCompile(`(?i)za (\d+) godinu`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"sr-Latn-BA": srLatnBALocale,
		"sr-Latn-ME": srLatnMELocale,
		"sr-Latn-XK": srLatnXKLocale,
	},
}

var srLatnXKLocale = LocaleData{
	Name:                  "sr-Latn-XK",
	DateOrder:             "",
	September:             []string{"sept"},
	Tuesday:               []string{"ut"},
	Wednesday:             []string{"sr"},
}

var srLatnBALocale = LocaleData{
	Name:                  "sr-Latn-BA",
	DateOrder:             "",
	September:             []string{"sept"},
	Tuesday:               []string{"ut"},
	Wednesday:             []string{"sr", "srijeda"},
	Sunday:                []string{"nedjelja"},
	AM:                    []string{"prije podne"},
}

var srLatnMELocale = LocaleData{
	Name:                  "sr-Latn-ME",
	DateOrder:             "",
	September:             []string{"sept"},
	Tuesday:               []string{"ut"},
	Wednesday:             []string{"sr", "srijeda"},
	Sunday:                []string{"nedjelja"},
	AM:                    []string{"prije podne"},
}
