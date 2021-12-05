package data

import "regexp"

var esLocale = LocaleData{
	Name:                  "es",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 2,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "a las", "cerca", "de", "del", "en", "y", "|", "，"},
	PertainWords:          []string{"de", "del"},
	January:               []string{"ene", "enero"},
	February:              []string{"feb", "febrero"},
	March:                 []string{"mar", "marzo"},
	April:                 []string{"abr", "abril"},
	May:                   []string{"may", "mayo"},
	June:                  []string{"jun", "junio"},
	July:                  []string{"jul", "julio"},
	August:                []string{"ago", "agosto"},
	September:             []string{"sep", "sept", "septiembre", "set", "setiembre"},
	October:               []string{"oct", "octubre"},
	November:              []string{"nov", "noviembre"},
	December:              []string{"dic", "diciembre"},
	Monday:                []string{"lu", "lun", "lunes"},
	Tuesday:               []string{"mar", "martes"},
	Wednesday:             []string{"mi", "mié", "miércoles"},
	Thursday:              []string{"ju", "jue", "jueves"},
	Friday:                []string{"vi", "vie", "viernes"},
	Saturday:              []string{"sa", "sáb", "sábado"},
	Sunday:                []string{"do", "dom", "domingo"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"a", "año", "años"},
	Month:                 []string{"m", "mes", "meses"},
	Week:                  []string{"sem", "semana", "semanas"},
	Day:                   []string{"d", "día", "días"},
	Hour:                  []string{"h", "hora", "horas"},
	Minute:                []string{"min", "minuto", "minutos"},
	Second:                []string{"s", "segundo", "segundos"},
	In:                    []string{"en"},
	Ago:                   []string{"hace"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hoy`},
		`0 hour ago`:   {`esta hora`},
		`0 minute ago`: {`este minuto`},
		`0 month ago`:  {`este mes`},
		`0 second ago`: {`ahora`},
		`0 week ago`:   {`esta semana`},
		`0 year ago`:   {`este año`},
		`1 day ago`:    {`ayer`},
		`1 month ago`:  {`el mes pasado`},
		`1 week ago`:   {`la semana pasada`},
		`1 year ago`:   {`el año pasado`},
		`2 day ago`:    {`anteayer`},
		`in 1 day`:     {`mañana`},
		`in 1 month`:   {`el próximo mes`},
		`in 1 week`:    {`la próxima semana`},
		`in 1 year`:    {`el próximo año`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)hace (\d+) día`),
			regexp.MustCompile(`(?i)hace (\d+) días`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)hace (\d+) h`),
			regexp.MustCompile(`(?i)hace (\d+) hora`),
			regexp.MustCompile(`(?i)hace (\d+) horas`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)hace (\d+) min`),
			regexp.MustCompile(`(?i)hace (\d+) minuto`),
			regexp.MustCompile(`(?i)hace (\d+) minutos`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)hace (\d+) m`),
			regexp.MustCompile(`(?i)hace (\d+) mes`),
			regexp.MustCompile(`(?i)hace (\d+) meses`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)hace (\d+) s`),
			regexp.MustCompile(`(?i)hace (\d+) segundo`),
			regexp.MustCompile(`(?i)hace (\d+) segundos`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)hace (\d+) sem`),
			regexp.MustCompile(`(?i)hace (\d+) semana`),
			regexp.MustCompile(`(?i)hace (\d+) semanas`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)hace (\d+) a`),
			regexp.MustCompile(`(?i)hace (\d+) año`),
			regexp.MustCompile(`(?i)hace (\d+) años`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) día`),
			regexp.MustCompile(`(?i)dentro de (\d+) días`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)dentro de (\d+) h`),
			regexp.MustCompile(`(?i)dentro de (\d+) hora`),
			regexp.MustCompile(`(?i)dentro de (\d+) horas`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)dentro de (\d+) min`),
			regexp.MustCompile(`(?i)dentro de (\d+) minuto`),
			regexp.MustCompile(`(?i)dentro de (\d+) minutos`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)dentro de (\d+) m`),
			regexp.MustCompile(`(?i)dentro de (\d+) mes`),
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dentro de (\d+) s`),
			regexp.MustCompile(`(?i)dentro de (\d+) segundo`),
			regexp.MustCompile(`(?i)dentro de (\d+) segundos`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)dentro de (\d+) sem`),
			regexp.MustCompile(`(?i)dentro de (\d+) semana`),
			regexp.MustCompile(`(?i)dentro de (\d+) semanas`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)dentro de (\d+) a`),
			regexp.MustCompile(`(?i)dentro de (\d+) año`),
			regexp.MustCompile(`(?i)dentro de (\d+) años`),
		},
	},
	Simplifications: map[string]string{
		`un`:  `1`,
		`una`: `1`,
	},
	LocaleSpecific: map[string]LocaleData{
		"es-419": es419Locale,
		"es-AR":  esARLocale,
		"es-BO":  esBOLocale,
		"es-BR":  esBRLocale,
		"es-BZ":  esBZLocale,
		"es-CL":  esCLLocale,
		"es-CO":  esCOLocale,
		"es-CR":  esCRLocale,
		"es-CU":  esCULocale,
		"es-DO":  esDOLocale,
		"es-EA":  esEALocale,
		"es-EC":  esECLocale,
		"es-GQ":  esGQLocale,
		"es-GT":  esGTLocale,
		"es-HN":  esHNLocale,
		"es-IC":  esICLocale,
		"es-MX":  esMXLocale,
		"es-NI":  esNILocale,
		"es-PA":  esPALocale,
		"es-PE":  esPELocale,
		"es-PH":  esPHLocale,
		"es-PR":  esPRLocale,
		"es-PY":  esPYLocale,
		"es-SV":  esSVLocale,
		"es-US":  esUSLocale,
		"es-UY":  esUYLocale,
		"es-VE":  esVELocale,
	},
}

var esBZLocale = LocaleData{
	Name:                  "es-BZ",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esVELocale = LocaleData{
	Name:                  "es-VE",
	DateOrder:             "",
}

var esCLLocale = LocaleData{
	Name:                  "es-CL",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esCOLocale = LocaleData{
	Name:                  "es-CO",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esBRLocale = LocaleData{
	Name:                  "es-BR",
	DateOrder:             "",
	September:             []string{"sep"},
}

var es419Locale = LocaleData{
	Name:                  "es-419",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esGQLocale = LocaleData{
	Name:                  "es-GQ",
	DateOrder:             "",
}

var esUSLocale = LocaleData{
	Name:                  "es-US",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esGTLocale = LocaleData{
	Name:                  "es-GT",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esPRLocale = LocaleData{
	Name:                  "es-PR",
	DateOrder:             "MDY",
	September:             []string{"sep"},
}

var esUYLocale = LocaleData{
	Name:                  "es-UY",
	DateOrder:             "",
	September:             []string{"set", "setiembre"},
}

var esECLocale = LocaleData{
	Name:                  "es-EC",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esDOLocale = LocaleData{
	Name:                  "es-DO",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esEALocale = LocaleData{
	Name:                  "es-EA",
	DateOrder:             "",
}

var esHNLocale = LocaleData{
	Name:                  "es-HN",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esPALocale = LocaleData{
	Name:                  "es-PA",
	DateOrder:             "MDY",
	September:             []string{"sep"},
}

var esBOLocale = LocaleData{
	Name:                  "es-BO",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esCULocale = LocaleData{
	Name:                  "es-CU",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esARLocale = LocaleData{
	Name:                  "es-AR",
	DateOrder:             "",
	September:             []string{"sep"},
	Second:                []string{"seg"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)hace (\d+) seg`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dentro de (\d+) seg`),
		},
	},
}

var esMXLocale = LocaleData{
	Name:                  "es-MX",
	DateOrder:             "",
	September:             []string{"sep"},
	RelativeType: map[string][]string{
		`in 1 month`: {`el mes próximo`},
		`in 1 week`:  {`la semana próxima`},
		`in 1 year`:  {`el año próximo`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`in \1 day`: {
			regexp.MustCompile(`(?i)en (\d+) día`),
			regexp.MustCompile(`(?i)en (\d+) días`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)en (\d+) h`),
			regexp.MustCompile(`(?i)en (\d+) n`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)en (\d+) min`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)en (\d+) m`),
			regexp.MustCompile(`(?i)en (\d+) mes`),
			regexp.MustCompile(`(?i)en (\d+) meses`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)en (\d+) s`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)en (\d+) sem`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)en (\d+) a`),
		},
	},
}

var esCRLocale = LocaleData{
	Name:                  "es-CR",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esSVLocale = LocaleData{
	Name:                  "es-SV",
	DateOrder:             "",
	September:             []string{"sep"},
}

var esPHLocale = LocaleData{
	Name:                  "es-PH",
	DateOrder:             "",
}

var esPELocale = LocaleData{
	Name:                  "es-PE",
	DateOrder:             "",
	September:             []string{"set", "setiembre"},
}

var esPYLocale = LocaleData{
	Name:                  "es-PY",
	DateOrder:             "",
	Second:                []string{"seg"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)hace (\d+) seg`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)dentro de (\d+) seg`),
		},
	},
}

var esICLocale = LocaleData{
	Name:                  "es-IC",
	DateOrder:             "",
}

var esNILocale = LocaleData{
	Name:                  "es-NI",
	DateOrder:             "",
	September:             []string{"sep"},
}
