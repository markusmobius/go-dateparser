package data

import "regexp"

var hrLocale = LocaleData{
	Name:                  "hr",
	DateOrder:             "DMY.",
	January:               []string{"sij", "siječanj", "siječnja"},
	February:              []string{"velj", "veljača", "veljače"},
	March:                 []string{"ožu", "ožujak", "ožujka"},
	April:                 []string{"tra", "travanj", "travnja"},
	May:                   []string{"svi", "svibanj", "svibnja"},
	June:                  []string{"lip", "lipanj", "lipnja"},
	July:                  []string{"srp", "srpanj", "srpnja"},
	August:                []string{"kol", "kolovoz", "kolovoza"},
	September:             []string{"ruj", "rujan", "rujna"},
	October:               []string{"lis", "listopad", "listopada"},
	November:              []string{"stu", "studeni", "studenoga"},
	December:              []string{"pro", "prosinac", "prosinca"},
	Monday:                []string{"pon", "ponedjeljak"},
	Tuesday:               []string{"uto", "utorak"},
	Wednesday:             []string{"sri", "srijeda"},
	Thursday:              []string{"čet", "četvrtak"},
	Friday:                []string{"pet", "petak"},
	Saturday:              []string{"sub", "subota"},
	Sunday:                []string{"ned", "nedjelja"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"g", "godina"},
	Month:                 []string{"m", "mj", "mjesec"},
	Week:                  []string{"tj", "tjedan"},
	Day:                   []string{"d", "dan"},
	Hour:                  []string{"h", "sat"},
	Minute:                []string{"min", "minuta"},
	Second:                []string{"s", "sekunda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`danas`},
		`0 hour ago`:   {`ovaj sat`},
		`0 minute ago`: {`ova minuta`},
		`0 month ago`:  {`ovaj mj`, `ovaj mjesec`},
		`0 second ago`: {`sad`},
		`0 week ago`:   {`ovaj tj`, `ovaj tjedan`},
		`0 year ago`:   {`ove g`, `ove god`, `ove godine`},
		`1 day ago`:    {`jučer`},
		`1 month ago`:  {`prošli mj`, `prošli mjesec`},
		`1 week ago`:   {`prošli tj`, `prošli tjedan`},
		`1 year ago`:   {`prošle g`, `prošle god`, `prošle godine`},
		`in 1 day`:     {`sutra`},
		`in 1 month`:   {`sljedeći mj`, `sljedeći mjesec`},
		`in 1 week`:    {`sljedeći tj`, `sljedeći tjedan`},
		`in 1 year`:    {`sljedeće g`, `sljedeće god`, `sljedeće godine`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)prije (\d+) d`),
			regexp.MustCompile(`(?i)prije (\d+) dan`),
			regexp.MustCompile(`(?i)prije (\d+) dana`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)prije (\d+) h`),
			regexp.MustCompile(`(?i)prije (\d+) sat`),
			regexp.MustCompile(`(?i)prije (\d+) sati`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)prije (\d+) min`),
			regexp.MustCompile(`(?i)prije (\d+) minuta`),
			regexp.MustCompile(`(?i)prije (\d+) minutu`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)prije (\d+) mj`),
			regexp.MustCompile(`(?i)prije (\d+) mjesec`),
			regexp.MustCompile(`(?i)prije (\d+) mjeseci`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)prije (\d+) s`),
			regexp.MustCompile(`(?i)prije (\d+) sekundi`),
			regexp.MustCompile(`(?i)prije (\d+) sekundu`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)prije (\d+) tj`),
			regexp.MustCompile(`(?i)prije (\d+) tjedan`),
			regexp.MustCompile(`(?i)prije (\d+) tjedana`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)prije (\d+) g`),
			regexp.MustCompile(`(?i)prije (\d+) godina`),
			regexp.MustCompile(`(?i)prije (\d+) godinu`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)za (\d+) d`),
			regexp.MustCompile(`(?i)za (\d+) dan`),
			regexp.MustCompile(`(?i)za (\d+) dana`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)za (\d+) h`),
			regexp.MustCompile(`(?i)za (\d+) sat`),
			regexp.MustCompile(`(?i)za (\d+) sati`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)za (\d+) min`),
			regexp.MustCompile(`(?i)za (\d+) minuta`),
			regexp.MustCompile(`(?i)za (\d+) minutu`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)za (\d+) mj`),
			regexp.MustCompile(`(?i)za (\d+) mjesec`),
			regexp.MustCompile(`(?i)za (\d+) mjeseci`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)za (\d+) s`),
			regexp.MustCompile(`(?i)za (\d+) sekundi`),
			regexp.MustCompile(`(?i)za (\d+) sekundu`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)za (\d+) tj`),
			regexp.MustCompile(`(?i)za (\d+) tjedan`),
			regexp.MustCompile(`(?i)za (\d+) tjedana`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)za (\d+) g`),
			regexp.MustCompile(`(?i)za (\d+) godina`),
			regexp.MustCompile(`(?i)za (\d+) godinu`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"hr-BA": hrBALocale,
	},
}

var hrBALocale = LocaleData{
	Name:                  "hr-BA",
	DateOrder:             "",
}
