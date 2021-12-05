package data

import "regexp"

var uzLocale = LocaleData{
	Name:                  "uz",
	DateOrder:             "DMY",
	January:               []string{"yan", "yanvar"},
	February:              []string{"fev", "fevral"},
	March:                 []string{"mar", "mart"},
	April:                 []string{"apr", "aprel"},
	May:                   []string{"may"},
	June:                  []string{"iyn", "iyun"},
	July:                  []string{"iyl", "iyul"},
	August:                []string{"avg", "avgust"},
	September:             []string{"sen", "sentabr"},
	October:               []string{"okt", "oktabr"},
	November:              []string{"noy", "noyabr"},
	December:              []string{"dek", "dekabr"},
	Monday:                []string{"dush", "dushanba"},
	Tuesday:               []string{"sesh", "seshanba"},
	Wednesday:             []string{"chor", "chorshanba"},
	Thursday:              []string{"pay", "payshanba"},
	Friday:                []string{"jum", "juma"},
	Saturday:              []string{"shan", "shanba"},
	Sunday:                []string{"yak", "yakshanba"},
	AM:                    []string{"to"},
	PM:                    []string{"tk"},
	Year:                  []string{"yil"},
	Month:                 []string{"oy"},
	Week:                  []string{"hafta"},
	Day:                   []string{"kun"},
	Hour:                  []string{"soat"},
	Minute:                []string{"daq", "daqiqa"},
	Second:                []string{"son", "soniya"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`bugun`},
		`0 hour ago`:   {`shu soatda`},
		`0 minute ago`: {`shu daqiqada`},
		`0 month ago`:  {`shu oy`},
		`0 second ago`: {`hozir`},
		`0 week ago`:   {`shu hafta`},
		`0 year ago`:   {`bu yil`, `shu yil`},
		`1 day ago`:    {`kecha`},
		`1 month ago`:  {`o‘tgan oy`},
		`1 week ago`:   {`o‘tgan hafta`},
		`1 year ago`:   {`o'tgan yil`, `o‘tgan yil`},
		`in 1 day`:     {`ertaga`},
		`in 1 month`:   {`keyingi oy`},
		`in 1 week`:    {`keyingi hafta`},
		`in 1 year`:    {`keyingi yil`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) kun oldin`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) soat oldin`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) daqiqa oldin`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) oy oldin`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) soniya oldin`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) hafta oldin`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) yil oldin`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) kundan keyin`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) soatdan keyin`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) daqiqadan keyin`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) oydan keyin`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) soniyadan keyin`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) haftadan keyin`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) yildan keyin`),
		},
	},
}
