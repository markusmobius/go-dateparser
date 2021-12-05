package data

import "regexp"

var ptLocale = LocaleData{
	Name:                  "pt",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "cerca", "de", "e", "|", "às", "，"},
	PertainWords:          []string{"de"},
	January:               []string{"jan", "janeiro"},
	February:              []string{"fev", "fevereiro"},
	March:                 []string{"mar", "março"},
	April:                 []string{"abr", "abril"},
	May:                   []string{"mai", "maio"},
	June:                  []string{"jun", "junho"},
	July:                  []string{"jul", "julho"},
	August:                []string{"ago", "agosto"},
	September:             []string{"septembro", "set", "setembro"},
	October:               []string{"out", "outubro"},
	November:              []string{"nov", "novembro"},
	December:              []string{"dez", "dezembro"},
	Monday:                []string{"seg", "segunda", "segunda-feira"},
	Tuesday:               []string{"ter", "terça", "terça-feira"},
	Wednesday:             []string{"qua", "quarta", "quarta-feira"},
	Thursday:              []string{"qui", "quinta", "quinta-feira"},
	Friday:                []string{"sex", "sexta", "sexta-feira"},
	Saturday:              []string{"sab", "sáb", "sábado"},
	Sunday:                []string{"dom", "domingo"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ano", "anos"},
	Month:                 []string{"meses", "mês"},
	Week:                  []string{"sem", "semana", "semanas"},
	Day:                   []string{"dia", "dias"},
	Hour:                  []string{"h", "hora", "horas"},
	Minute:                []string{"m", "min", "minuto", "minutos"},
	Second:                []string{"s", "seg", "segundo", "segundos"},
	In:                    []string{"em"},
	Ago:                   []string{"atrás", "há"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hoje`},
		`0 hour ago`:   {`esta hora`},
		`0 minute ago`: {`este minuto`},
		`0 month ago`:  {`este mês`},
		`0 second ago`: {`agora`},
		`0 week ago`:   {`esta semana`},
		`0 year ago`:   {`este ano`},
		`1 day ago`:    {`ontem`},
		`1 month ago`:  {`mês passado`},
		`1 week ago`:   {`semana passada`},
		`1 year ago`:   {`ano passado`},
		`2 day ago`:    {`anteontem`},
		`in 1 day`:     {`amanhã`},
		`in 1 month`:   {`próximo mês`},
		`in 1 week`:    {`próxima semana`},
		`in 1 year`:    {`próximo ano`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)há (\d+) dia`),
			regexp.MustCompile(`(?i)há (\d+) dias`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)há (\d+) h`),
			regexp.MustCompile(`(?i)há (\d+) hora`),
			regexp.MustCompile(`(?i)há (\d+) horas`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)há (\d+) min`),
			regexp.MustCompile(`(?i)há (\d+) mins`),
			regexp.MustCompile(`(?i)há (\d+) minuto`),
			regexp.MustCompile(`(?i)há (\d+) minutos`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)há (\d+) meses`),
			regexp.MustCompile(`(?i)há (\d+) mês`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) seg`),
			regexp.MustCompile(`(?i)há (\d+) segundo`),
			regexp.MustCompile(`(?i)há (\d+) segundos`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)há (\d+) sem`),
			regexp.MustCompile(`(?i)há (\d+) semana`),
			regexp.MustCompile(`(?i)há (\d+) semanas`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)há (\d+) ano`),
			regexp.MustCompile(`(?i)há (\d+) anos`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)em (\d+) dia`),
			regexp.MustCompile(`(?i)em (\d+) dias`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)em (\d+) h`),
			regexp.MustCompile(`(?i)em (\d+) hora`),
			regexp.MustCompile(`(?i)em (\d+) horas`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)em (\d+) min`),
			regexp.MustCompile(`(?i)em (\d+) mins`),
			regexp.MustCompile(`(?i)em (\d+) minuto`),
			regexp.MustCompile(`(?i)em (\d+) minutos`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)em (\d+) meses`),
			regexp.MustCompile(`(?i)em (\d+) mês`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)em (\d+) seg`),
			regexp.MustCompile(`(?i)em (\d+) segs`),
			regexp.MustCompile(`(?i)em (\d+) segundo`),
			regexp.MustCompile(`(?i)em (\d+) segundos`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)em (\d+) sem`),
			regexp.MustCompile(`(?i)em (\d+) semana`),
			regexp.MustCompile(`(?i)em (\d+) semanas`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)em (\d+) ano`),
			regexp.MustCompile(`(?i)em (\d+) anos`),
		},
	},
	Simplifications: map[string]string{
		`alguns segundos`: `44 segundos`,
		`um`:              `1`,
		`uma`:             `1`,
	},
	LocaleSpecific: map[string]LocaleData{
		"pt-AO": ptAOLocale,
		"pt-CH": ptCHLocale,
		"pt-CV": ptCVLocale,
		"pt-GQ": ptGQLocale,
		"pt-GW": ptGWLocale,
		"pt-LU": ptLULocale,
		"pt-MO": ptMOLocale,
		"pt-MZ": ptMZLocale,
		"pt-PT": ptPTLocale,
		"pt-ST": ptSTLocale,
		"pt-TL": ptTLLocale,
	},
}

var ptPTLocale = LocaleData{
	Name:                  "pt-PT",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptAOLocale = LocaleData{
	Name:                  "pt-AO",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptMOLocale = LocaleData{
	Name:                  "pt-MO",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptMZLocale = LocaleData{
	Name:                  "pt-MZ",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptLULocale = LocaleData{
	Name:                  "pt-LU",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptSTLocale = LocaleData{
	Name:                  "pt-ST",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptTLLocale = LocaleData{
	Name:                  "pt-TL",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptCHLocale = LocaleData{
	Name:                  "pt-CH",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptCVLocale = LocaleData{
	Name:                  "pt-CV",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptGWLocale = LocaleData{
	Name:                  "pt-GW",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}

var ptGQLocale = LocaleData{
	Name:                  "pt-GQ",
	DateOrder:             "",
	Monday:                []string{"segunda"},
	Tuesday:               []string{"terça"},
	Wednesday:             []string{"quarta"},
	Thursday:              []string{"quinta"},
	Friday:                []string{"sexta"},
	AM:                    []string{"da manhã", "manhã"},
	PM:                    []string{"da tarde", "tarde"},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 second ago`: {
			regexp.MustCompile(`(?i)há (\d+) s`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)dentro de (\d+) dia`),
			regexp.MustCompile(`(?i)dentro de (\d+) dias`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) meses`),
			regexp.MustCompile(`(?i)dentro de (\d+) mês`),
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
			regexp.MustCompile(`(?i)dentro de (\d+) ano`),
			regexp.MustCompile(`(?i)dentro de (\d+) anos`),
		},
	},
}
