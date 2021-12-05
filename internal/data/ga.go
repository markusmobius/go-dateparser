package data

import "regexp"

var gaLocale = LocaleData{
	Name:                  "ga",
	DateOrder:             "DMY",
	January:               []string{"ean", "eanáir"},
	February:              []string{"feabh", "feabhra"},
	March:                 []string{"márta"},
	April:                 []string{"aib", "aibreán"},
	May:                   []string{"beal", "bealtaine"},
	June:                  []string{"meith", "meitheamh"},
	July:                  []string{"iúil"},
	August:                []string{"lún", "lúnasa"},
	September:             []string{"meán fómhair", "mfómh"},
	October:               []string{"deireadh fómhair", "dfómh"},
	November:              []string{"samh", "samhain"},
	December:              []string{"noll", "nollaig"},
	Monday:                []string{"dé luain", "luan"},
	Tuesday:               []string{"dé máirt", "máirt"},
	Wednesday:             []string{"céad", "dé céadaoin"},
	Thursday:              []string{"déar", "déardaoin"},
	Friday:                []string{"aoine", "dé haoine"},
	Saturday:              []string{"dé sathairn", "sath"},
	Sunday:                []string{"domh", "dé domhnaigh"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"bl", "bliain"},
	Month:                 []string{"mí"},
	Week:                  []string{"scht", "seachtain"},
	Day:                   []string{"lá"},
	Hour:                  []string{"u", "uair"},
	Minute:                []string{"n", "nóim", "nóiméad"},
	Second:                []string{"s", "soic", "soicind"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`inniu`},
		`0 hour ago`:   {`an uair seo`},
		`0 minute ago`: {`an nóiméad seo`},
		`0 month ago`:  {`an mhí seo`},
		`0 second ago`: {`anois`},
		`0 week ago`:   {`an tscht seo`, `an tseachtain seo`},
		`0 year ago`:   {`an bhl seo`, `an bhliain seo`},
		`1 day ago`:    {`inné`},
		`1 month ago`:  {`an mhí seo caite`},
		`1 week ago`:   {`an tscht seo caite`, `an tseachtain seo caite`},
		`1 year ago`:   {`anuraidh`},
		`in 1 day`:     {`amárach`},
		`in 1 month`:   {`an mhí seo chugainn`},
		`in 1 week`:    {`an tscht seo chugainn`, `an tseachtain seo chugainn`},
		`in 1 year`:    {`an bhl seo chugainn`, `an bhliain seo chugainn`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) lá ó shin`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) uair an chloig ó shin`),
			regexp.MustCompile(`(?i)(\d+) uair ó shin`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) nóim ó shin`),
			regexp.MustCompile(`(?i)(\d+) nóiméad ó shin`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mhí ó shin`),
			regexp.MustCompile(`(?i)(\d+) mí ó shin`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) soic ó shin`),
			regexp.MustCompile(`(?i)(\d+) soicind ó shin`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) scht ó shin`),
			regexp.MustCompile(`(?i)(\d+) seachtain ó shin`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) bhl ó shin`),
			regexp.MustCompile(`(?i)(\d+) bhliain ó shin`),
			regexp.MustCompile(`(?i)(\d+) bl ó shin`),
			regexp.MustCompile(`(?i)(\d+) bliain ó shin`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)i gceann (\d+) lá`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)i gceann (\d+) uair`),
			regexp.MustCompile(`(?i)i gceann (\d+) uair an chloig`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)i gceann (\d+) nóim`),
			regexp.MustCompile(`(?i)i gceann (\d+) nóiméad`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)i gceann (\d+) mhí`),
			regexp.MustCompile(`(?i)i gceann (\d+) mí`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)i gceann (\d+) soic`),
			regexp.MustCompile(`(?i)i gceann (\d+) soicind`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)i gceann (\d+) scht`),
			regexp.MustCompile(`(?i)i gceann (\d+) seachtain`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)i gceann (\d+) bhliain`),
			regexp.MustCompile(`(?i)i gceann (\d+) bl`),
			regexp.MustCompile(`(?i)i gceann (\d+) bliain`),
		},
	},
}
