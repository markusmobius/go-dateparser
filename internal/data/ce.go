// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ceLocale = LocaleData{
	Name:                  "ce",
	DateOrder:             "YMD",
	January:               []string{"янв", "январь"},
	February:              []string{"фев", "февраль"},
	March:                 []string{"мар", "март"},
	April:                 []string{"апр", "апрель"},
	May:                   []string{"май"},
	June:                  []string{"июн", "июнь"},
	July:                  []string{"июл", "июль"},
	August:                []string{"авг", "август"},
	September:             []string{"сен", "сентябрь"},
	October:               []string{"окт", "октябрь"},
	November:              []string{"ноя", "ноябрь"},
	December:              []string{"дек", "декабрь"},
	Monday:                []string{"оршотан де"},
	Tuesday:               []string{"шинарин де"},
	Wednesday:             []string{"кхаарин де"},
	Thursday:              []string{"еарин де"},
	Friday:                []string{"пӏераскан де"},
	Saturday:              []string{"шот де"},
	Sunday:                []string{"кӏиранан де"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ш", "шо"},
	Month:                 []string{"бут", "бутт"},
	Week:                  []string{"кӏир", "кӏира"},
	Day:                   []string{"де"},
	Hour:                  []string{"сахь", "сахьт"},
	Minute:                []string{"мин", "минот"},
	Second:                []string{"сек", "секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`тахана`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`карарчу баттахь`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`карарчу кӏирнахь`},
		`0 year ago`:   {`карарчу шарахь`},
		`1 day ago`:    {`селхана`},
		`1 month ago`:  {`баханчу баттахь`},
		`1 week ago`:   {`даханчу кӏирнахь`},
		`1 year ago`:   {`даханчу шарахь`},
		`in 1 day`:     {`кхана`},
		`in 1 month`:   {`рогӏерчу баттахь`},
		`in 1 week`:    {`рогӏерчу кӏирнахь`},
		`in 1 year`:    {`рогӏерчу шарахь`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) д хьалха`),
			regexp.MustCompile(`(?i)(\d+) де хьалха`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) сахь хьалха`),
			regexp.MustCompile(`(?i)(\d+) сахьт хьалха`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) мин хьалха`),
			regexp.MustCompile(`(?i)(\d+) минот хьалха`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) б хьалха`),
			regexp.MustCompile(`(?i)(\d+) бутт хьалха`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) сек хьалха`),
			regexp.MustCompile(`(?i)(\d+) секунд хьалха`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) кӏир хьалха`),
			regexp.MustCompile(`(?i)(\d+) кӏира хьалха`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ш хьалха`),
			regexp.MustCompile(`(?i)(\d+) шо хьалха`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) д даьлча`),
			regexp.MustCompile(`(?i)(\d+) де даьлча`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) сахь даьлча`),
			regexp.MustCompile(`(?i)(\d+) сахьт даьлча`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) мин яьлча`),
			regexp.MustCompile(`(?i)(\d+) минот яьлча`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) б баьлча`),
			regexp.MustCompile(`(?i)(\d+) бутт баьлча`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) сек яьлча`),
			regexp.MustCompile(`(?i)(\d+) секунд яьлча`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) кӏир даьлча`),
			regexp.MustCompile(`(?i)(\d+) кӏира даьлча`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) ш даьлча`),
			regexp.MustCompile(`(?i)(\d+) шо даьлча`),
		},
	},
}
