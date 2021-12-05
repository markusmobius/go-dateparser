package data

import "regexp"

var swLocale = LocaleData{
	Name:                  "sw",
	DateOrder:             "DMY",
	January:               []string{"jan", "januari"},
	February:              []string{"feb", "februari"},
	March:                 []string{"mac", "machi"},
	April:                 []string{"apr", "aprili"},
	May:                   []string{"mei"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "julai"},
	August:                []string{"ago", "agosti"},
	September:             []string{"sep", "septemba"},
	October:               []string{"okt", "oktoba"},
	November:              []string{"nov", "novemba"},
	December:              []string{"des", "desemba"},
	Monday:                []string{"jumatatu"},
	Tuesday:               []string{"jumanne"},
	Wednesday:             []string{"jumatano"},
	Thursday:              []string{"alhamisi"},
	Friday:                []string{"ijumaa"},
	Saturday:              []string{"jumamosi"},
	Sunday:                []string{"jumapili"},
	AM:                    []string{"am", "asubuhi"},
	PM:                    []string{"mchana", "pm"},
	Year:                  []string{"mwaka"},
	Month:                 []string{"mwezi"},
	Week:                  []string{"wiki"},
	Day:                   []string{"siku"},
	Hour:                  []string{"saa"},
	Minute:                []string{"dak", "dakika"},
	Second:                []string{"sek", "sekunde"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`leo`},
		`0 hour ago`:   {`saa hii`},
		`0 minute ago`: {`dakika hii`},
		`0 month ago`:  {`mwezi huu`},
		`0 second ago`: {`sasa hivi`},
		`0 week ago`:   {`wiki hii`},
		`0 year ago`:   {`mwaka huu`},
		`1 day ago`:    {`jana`},
		`1 month ago`:  {`mwezi uliopita`},
		`1 week ago`:   {`wiki iliyopita`},
		`1 year ago`:   {`mwaka uliopita`},
		`in 1 day`:     {`kesho`},
		`in 1 month`:   {`mwezi ujao`},
		`in 1 week`:    {`wiki ijayo`},
		`in 1 year`:    {`mwaka ujao`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)siku (\d+) iliyopita`),
			regexp.MustCompile(`(?i)siku (\d+) zilizopita`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)saa (\d+) iliyopita`),
			regexp.MustCompile(`(?i)saa (\d+) zilizopita`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)dakika (\d+) iliyopita`),
			regexp.MustCompile(`(?i)dakika (\d+) zilizopita`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)miezi (\d+) iliyopita`),
			regexp.MustCompile(`(?i)mwezi (\d+) uliopita`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)sekunde (\d+) iliyopita`),
			regexp.MustCompile(`(?i)sekunde (\d+) zilizopita`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)wiki (\d+) iliyopita`),
			regexp.MustCompile(`(?i)wiki (\d+) zilizopita`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)miaka (\d+) iliyopita`),
			regexp.MustCompile(`(?i)mwaka (\d+) uliopita`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)baada ya siku (\d+)`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)baada ya saa (\d+)`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)baada ya dakika (\d+)`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)baada ya miezi (\d+)`),
			regexp.MustCompile(`(?i)baada ya mwezi (\d+)`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)baada ya sekunde (\d+)`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)baada ya wiki (\d+)`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)baada ya miaka (\d+)`),
			regexp.MustCompile(`(?i)baada ya mwaka (\d+)`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"sw-CD": swCDLocale,
		"sw-KE": swKELocale,
		"sw-UG": swUGLocale,
	},
}

var swKELocale = LocaleData{
	Name:                  "sw-KE",
	DateOrder:             "",
}

var swUGLocale = LocaleData{
	Name:                  "sw-UG",
	DateOrder:             "",
}

var swCDLocale = LocaleData{
	Name:                  "sw-CD",
	DateOrder:             "",
	Week:                  []string{"juma"},
}
