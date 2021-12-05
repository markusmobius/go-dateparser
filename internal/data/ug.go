package data

import "regexp"

var ugLocale = LocaleData{
	Name:                  "ug",
	DateOrder:             "YMD",
	January:               []string{"يانۋار"},
	February:              []string{"فېۋرال"},
	March:                 []string{"مارت"},
	April:                 []string{"ئاپرېل"},
	May:                   []string{"ماي"},
	June:                  []string{"ئىيۇن"},
	July:                  []string{"ئىيۇل"},
	August:                []string{"ئاۋغۇست"},
	September:             []string{"سېنتەبىر"},
	October:               []string{"ئۆكتەبىر"},
	November:              []string{"نويابىر"},
	December:              []string{"دېكابىر"},
	Monday:                []string{"دۈ", "دۈشەنبە"},
	Tuesday:               []string{"سە", "سەيشەنبە"},
	Wednesday:             []string{"چا", "چارشەنبە"},
	Thursday:              []string{"پە", "پەيشەنبە"},
	Friday:                []string{"جۈ", "جۈمە"},
	Saturday:              []string{"شە", "شەنبە"},
	Sunday:                []string{"يە", "يەكشەنبە"},
	AM:                    []string{"چب", "چۈشتىن بۇرۇن"},
	PM:                    []string{"چك", "چۈشتىن كېيىن"},
	Year:                  []string{"يىل"},
	Month:                 []string{"ئاي"},
	Week:                  []string{"ھەپتە"},
	Day:                   []string{"كۈن"},
	Hour:                  []string{"سائەت"},
	Minute:                []string{"مىنۇت"},
	Second:                []string{"سېكۇنت"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`بۈگۈن`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`بۇ ئاي`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`بۇ ھەپتە`},
		`0 year ago`:   {`بۇ يىل`},
		`1 day ago`:    {`تۈنۈگۈن`},
		`1 month ago`:  {`ئۆتكەن ئاي`},
		`1 week ago`:   {`ئۆتكەن ھەپتە`},
		`1 year ago`:   {`ئۆتكەن يىل`},
		`in 1 day`:     {`ئەتە`},
		`in 1 month`:   {`كېلەر ئاي`},
		`in 1 week`:    {`كېلەر ھەپتە`},
		`in 1 year`:    {`كېلەر يىل`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) كۈن ئىلگىرى`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) سائەت ئىلگىرى`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) مىنۇت ئىلگىرى`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ئاي ئىلگىرى`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) سېكۇنت ئىلگىرى`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ھەپتە ئىلگىرى`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) يىل ئىلگىرى`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) كۈندىن كېيىن`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) سائەتتىن كېيىن`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) مىنۇتتىن كېيىن`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ئايدىن كېيىن`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) سېكۇنتتىن كېيىن`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ھەپتىدىن كېيىن`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) يىلدىن كېيىن`),
		},
	},
}
