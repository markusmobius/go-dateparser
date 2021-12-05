package data

import "regexp"

var kkLocale = LocaleData{
	Name:                  "kk",
	DateOrder:             "DMY",
	January:               []string{"қаң", "қаңтар"},
	February:              []string{"ақп", "ақпан"},
	March:                 []string{"нау", "наурыз"},
	April:                 []string{"сәу", "сәуір"},
	May:                   []string{"мам", "мамыр"},
	June:                  []string{"мау", "маусым"},
	July:                  []string{"шіл", "шілде"},
	August:                []string{"там", "тамыз"},
	September:             []string{"қыр", "қыркүйек"},
	October:               []string{"қаз", "қазан"},
	November:              []string{"қар", "қараша"},
	December:              []string{"жел", "желтоқсан"},
	Monday:                []string{"дс", "дүйсенбі"},
	Tuesday:               []string{"сейсенбі", "сс"},
	Wednesday:             []string{"ср", "сәрсенбі"},
	Thursday:              []string{"бейсенбі", "бс"},
	Friday:                []string{"жм", "жұма"},
	Saturday:              []string{"сб", "сенбі"},
	Sunday:                []string{"жексенбі", "жс"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ж", "жыл"},
	Month:                 []string{"ай"},
	Week:                  []string{"ап", "апта"},
	Day:                   []string{"күн"},
	Hour:                  []string{"сағ", "сағат"},
	Minute:                []string{"мин", "минут"},
	Second:                []string{"с", "секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`бүгін`},
		`0 hour ago`:   {`осы сағат`},
		`0 minute ago`: {`осы минут`},
		`0 month ago`:  {`осы ай`},
		`0 second ago`: {`қазір`},
		`0 week ago`:   {`осы апта`},
		`0 year ago`:   {`биылғы жыл`},
		`1 day ago`:    {`кеше`},
		`1 month ago`:  {`өткен ай`},
		`1 week ago`:   {`өткен апта`},
		`1 year ago`:   {`былтырғы жыл`},
		`in 1 day`:     {`ертең`},
		`in 1 month`:   {`келесі ай`},
		`in 1 week`:    {`келесі апта`},
		`in 1 year`:    {`келесі жыл`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) күн бұрын`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) сағ бұрын`),
			regexp.MustCompile(`(?i)(\d+) сағат бұрын`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) мин бұрын`),
			regexp.MustCompile(`(?i)(\d+) минут бұрын`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ай бұрын`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) сек бұрын`),
			regexp.MustCompile(`(?i)(\d+) секунд бұрын`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ап бұрын`),
			regexp.MustCompile(`(?i)(\d+) апта бұрын`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ж бұрын`),
			regexp.MustCompile(`(?i)(\d+) жыл бұрын`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) күннен кейін`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) сағ кейін`),
			regexp.MustCompile(`(?i)(\d+) сағаттан кейін`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) мин кейін`),
			regexp.MustCompile(`(?i)(\d+) минуттан кейін`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) айдан кейін`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) сек кейін`),
			regexp.MustCompile(`(?i)(\d+) секундтан кейін`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ап кейін`),
			regexp.MustCompile(`(?i)(\d+) аптадан кейін`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) ж кейін`),
			regexp.MustCompile(`(?i)(\d+) жылдан кейін`),
		},
	},
}
