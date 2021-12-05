package data

import "regexp"

var caLocale = LocaleData{
	Name:                  "ca",
	DateOrder:             "DMY",
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "de", "del", "i", "l'", "|", "，"},
	PertainWords:          []string{"de", "del"},
	January:               []string{"de gen", "de gener", "gen", "gener"},
	February:              []string{"de febr", "de febrer", "febr", "febrer"},
	March:                 []string{"de març", "març"},
	April:                 []string{"abr", "abril", "d'abr", "d'abril"},
	May:                   []string{"de maig", "maig"},
	June:                  []string{"de juny", "juny"},
	July:                  []string{"de jul", "de juliol", "jul", "juliol"},
	August:                []string{"ag", "agost", "d'ag", "d'agost"},
	September:             []string{"de set", "de setembre", "set", "setembre"},
	October:               []string{"d'oct", "d'octubre", "oct", "octubre"},
	November:              []string{"de nov", "de novembre", "nov", "novembre"},
	December:              []string{"de des", "de desembre", "des", "desembre"},
	Monday:                []string{"dilluns", "dl"},
	Tuesday:               []string{"dimarts", "dt"},
	Wednesday:             []string{"dc", "dimecres"},
	Thursday:              []string{"dijous", "dj"},
	Friday:                []string{"divendres", "dv"},
	Saturday:              []string{"dissabte", "ds"},
	Sunday:                []string{"dg", "diumenge"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"any"},
	Month:                 []string{"mes"},
	Week:                  []string{"setm", "setmana"},
	Day:                   []string{"dia"},
	Hour:                  []string{"h", "hora"},
	Minute:                []string{"min", "minut"},
	Second:                []string{"s", "segon"},
	In:                    []string{"en"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`avui`, `hui`},
		`0 hour ago`:   {`aquesta hora`},
		`0 minute ago`: {`aquest minut`},
		`0 month ago`:  {`aquest mes`},
		`0 second ago`: {`ara`},
		`0 week ago`:   {`aquesta setm`, `aquesta setmana`},
		`0 year ago`:   {`enguany`},
		`1 day ago`:    {`ahir`},
		`1 month ago`:  {`el mes passat`, `mes passat`},
		`1 week ago`:   {`la setm passada`, `la setmana passada`, `setm passada`},
		`1 year ago`:   {`l'any passat`},
		`2 day ago`:    {`despús-ahir`, `abans-d’ahir`, `dellà-ahir`},
		`in 1 day`:     {`demà`},
		`in 1 month`:   {`el mes que ve`, `mes vinent`},
		`in 1 week`:    {`la propera setmana`, `la pròxima setmana`, `la setm que ve`, `la setmana que ve`, `la setmana vinent`, `setm vinent`},
		`in 1 year`:    {`l'any que ve`},
		`in 2 day`:     {`endemà`, `sendemà`, `despús-demà`, `demà passat`, `passat demà`},
		`in 3 day`:     {`endemà passat`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)fa (\d+) dia`),
			regexp.MustCompile(`(?i)fa (\d+) dies`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)fa (\d+) h`),
			regexp.MustCompile(`(?i)fa (\d+) hora`),
			regexp.MustCompile(`(?i)fa (\d+) hores`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)fa (\d+) min`),
			regexp.MustCompile(`(?i)fa (\d+) minut`),
			regexp.MustCompile(`(?i)fa (\d+) minuts`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)fa (\d+) mes`),
			regexp.MustCompile(`(?i)fa (\d+) mesos`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)fa (\d+) s`),
			regexp.MustCompile(`(?i)fa (\d+) segon`),
			regexp.MustCompile(`(?i)fa (\d+) segons`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)fa (\d+) setm`),
			regexp.MustCompile(`(?i)fa (\d+) setmana`),
			regexp.MustCompile(`(?i)fa (\d+) setmanes`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)fa (\d+) any`),
			regexp.MustCompile(`(?i)fa (\d+) anys`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) dia`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) dies`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) h`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) hora`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) hores`),
			regexp.MustCompile(`(?i)d‘aquí a (\d+) h`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) min`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) minut`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) minuts`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) mes`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) mesos`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) s`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) segon`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) segons`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) setm`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) setmana`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) setmanes`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)d'aquí a (\d+) any`),
			regexp.MustCompile(`(?i)d'aquí a (\d+) anys`),
		},
	},
	Simplifications: map[string]string{
		`un`:  `1`,
		`una`: `1`,
	},
	LocaleSpecific: map[string]LocaleData{
		"ca-AD": caADLocale,
		"ca-FR": caFRLocale,
		"ca-IT": caITLocale,
	},
}

var caFRLocale = LocaleData{
	Name:                  "ca-FR",
	DateOrder:             "",
}

var caITLocale = LocaleData{
	Name:                  "ca-IT",
	DateOrder:             "",
}

var caADLocale = LocaleData{
	Name:                  "ca-AD",
	DateOrder:             "",
}
