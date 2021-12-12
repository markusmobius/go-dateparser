// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var is_Locale = LocaleData{
	Name:                  "is",
	DateOrder:             "DMY",
	January:               []string{"jan", "janúar"},
	February:              []string{"feb", "febrúar"},
	March:                 []string{"mar", "mars"},
	April:                 []string{"apr", "apríl"},
	May:                   []string{"maí"},
	June:                  []string{"jún", "júní"},
	July:                  []string{"júl", "júlí"},
	August:                []string{"ágú", "ágúst"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "október"},
	November:              []string{"nóv", "nóvember"},
	December:              []string{"des", "desember"},
	Monday:                []string{"mán", "mánudagur"},
	Tuesday:               []string{"þri", "þriðjudagur"},
	Wednesday:             []string{"mið", "miðvikudagur"},
	Thursday:              []string{"fim", "fimmtudagur"},
	Friday:                []string{"fös", "föstudagur"},
	Saturday:              []string{"lau", "laugardagur"},
	Sunday:                []string{"sun", "sunnudagur"},
	AM:                    []string{"fh"},
	PM:                    []string{"eh"},
	Year:                  []string{"ár"},
	Month:                 []string{"mán", "mánuður"},
	Week:                  []string{"v", "vika"},
	Day:                   []string{"d", "dagur"},
	Hour:                  []string{"klst", "klukkustund"},
	Minute:                []string{"mín", "mínúta"},
	Second:                []string{"sek", "sekúnda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`í dag`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`í þessum mán`, `í þessum mánuði`},
		`0 second ago`: {`núna`},
		`0 week ago`:   {`í þessari viku`},
		`0 year ago`:   {`á þessu ári`},
		`1 day ago`:    {`í gær`},
		`1 month ago`:  {`í síðasta mán`, `í síðasta mánuði`},
		`1 week ago`:   {`í síðustu viku`},
		`1 year ago`:   {`á síðasta ári`},
		`in 1 day`:     {`á morgun`},
		`in 1 month`:   {`í næsta mán`, `í næsta mánuði`},
		`in 1 week`:    {`í næstu viku`},
		`in 1 year`:    {`á næsta ári`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) degi`),
			regexp.MustCompile(`(?i)fyrir (\d+) dögum`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) klst`),
			regexp.MustCompile(`(?i)fyrir (\d+) klukkustund`),
			regexp.MustCompile(`(?i)fyrir (\d+) klukkustundum`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) mín`),
			regexp.MustCompile(`(?i)fyrir (\d+) mínútu`),
			regexp.MustCompile(`(?i)fyrir (\d+) mínútum`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) mán`),
			regexp.MustCompile(`(?i)fyrir (\d+) mánuði`),
			regexp.MustCompile(`(?i)fyrir (\d+) mánuðum`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) sek`),
			regexp.MustCompile(`(?i)fyrir (\d+) sekúndu`),
			regexp.MustCompile(`(?i)fyrir (\d+) sekúndum`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) viku`),
			regexp.MustCompile(`(?i)fyrir (\d+) vikum`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)fyrir (\d+) ári`),
			regexp.MustCompile(`(?i)fyrir (\d+) árum`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)eftir (\d+) dag`),
			regexp.MustCompile(`(?i)eftir (\d+) daga`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)eftir (\d+) klst`),
			regexp.MustCompile(`(?i)eftir (\d+) klukkustund`),
			regexp.MustCompile(`(?i)eftir (\d+) klukkustundir`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)eftir (\d+) mín`),
			regexp.MustCompile(`(?i)eftir (\d+) mínútu`),
			regexp.MustCompile(`(?i)eftir (\d+) mínútur`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)eftir (\d+) mán`),
			regexp.MustCompile(`(?i)eftir (\d+) mánuð`),
			regexp.MustCompile(`(?i)eftir (\d+) mánuði`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)eftir (\d+) sek`),
			regexp.MustCompile(`(?i)eftir (\d+) sekúndu`),
			regexp.MustCompile(`(?i)eftir (\d+) sekúndur`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)eftir (\d+) viku`),
			regexp.MustCompile(`(?i)eftir (\d+) vikur`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)eftir (\d+) ár`),
		},
	},
}