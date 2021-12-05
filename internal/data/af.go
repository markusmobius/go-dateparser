package data

import "regexp"

var afLocale = LocaleData{
	Name:                  "af",
	DateOrder:             "YMD",
	January:               []string{"jan", "januarie"},
	February:              []string{"feb", "februarie"},
	March:                 []string{"maart", "mrt"},
	April:                 []string{"apr", "april"},
	May:                   []string{"mei"},
	June:                  []string{"jun", "junie"},
	July:                  []string{"jul", "julie"},
	August:                []string{"aug", "augustus"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"des", "desember"},
	Monday:                []string{"ma", "maandag"},
	Tuesday:               []string{"di", "dinsdag"},
	Wednesday:             []string{"wo", "woensdag"},
	Thursday:              []string{"do", "donderdag"},
	Friday:                []string{"vr", "vrydag"},
	Saturday:              []string{"sa", "saterdag"},
	Sunday:                []string{"so", "sondag"},
	AM:                    []string{"vm"},
	PM:                    []string{"nm"},
	Year:                  []string{"j", "jaar"},
	Month:                 []string{"maand", "md"},
	Week:                  []string{"week", "wk"},
	Day:                   []string{"d", "dag"},
	Hour:                  []string{"u", "uur"},
	Minute:                []string{"m", "min", "minuut"},
	Second:                []string{"s", "sek", "sekonde"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`vandag`},
		`0 hour ago`:   {`hierdie uur`},
		`0 minute ago`: {`hierdie minuut`},
		`0 month ago`:  {`vandeesmaand`},
		`0 second ago`: {`nou`},
		`0 week ago`:   {`vandeesweek`},
		`0 year ago`:   {`hierdie jaar`},
		`1 day ago`:    {`gister`},
		`1 month ago`:  {`verlede maand`},
		`1 week ago`:   {`verlede week`},
		`1 year ago`:   {`verlede jaar`},
		`in 1 day`:     {`môre`},
		`in 1 month`:   {`volgende maand`},
		`in 1 week`:    {`volgende week`},
		`in 1 year`:    {`volgende jaar`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) dae gelede`),
			regexp.MustCompile(`(?i)(\d+) dag gelede`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) uur gelede`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min gelede`),
			regexp.MustCompile(`(?i)(\d+) minute gelede`),
			regexp.MustCompile(`(?i)(\d+) minuut gelede`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) maand gelede`),
			regexp.MustCompile(`(?i)(\d+) maande gelede`),
			regexp.MustCompile(`(?i)(\d+) md gelede`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sek gelede`),
			regexp.MustCompile(`(?i)(\d+) sekonde gelede`),
			regexp.MustCompile(`(?i)(\d+) sekondes gelede`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) w gelede`),
			regexp.MustCompile(`(?i)(\d+) week gelede`),
			regexp.MustCompile(`(?i)(\d+) weke gelede`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) jaar gelede`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)oor (\d+) dae`),
			regexp.MustCompile(`(?i)oor (\d+) dag`),
			regexp.MustCompile(`(?i)oor (\d+) minuut`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)oor (\d+) uur`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)oor (\d+) min`),
			regexp.MustCompile(`(?i)oor (\d+) minuut`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)oor (\d+) md`),
			regexp.MustCompile(`(?i)oor (\d+) minuut`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)oor (\d+) sek`),
			regexp.MustCompile(`(?i)oor (\d+) sekonde`),
			regexp.MustCompile(`(?i)oor (\d+) sekondes`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)oor (\d+) w`),
			regexp.MustCompile(`(?i)oor (\d+) week`),
			regexp.MustCompile(`(?i)oor (\d+) weke`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)oor (\d+) jaar`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"af-NA": afNALocale,
	},
}

var afNALocale = LocaleData{
	Name:                  "af-NA",
	DateOrder:             "",
}
