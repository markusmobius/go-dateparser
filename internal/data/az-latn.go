package data

import "regexp"

var azLatnLocale = LocaleData{
	Name:                  "az-Latn",
	DateOrder:             "DMY",
	January:               []string{"yan", "yanvar"},
	February:              []string{"fev", "fevral"},
	March:                 []string{"mar", "mart"},
	April:                 []string{"apr", "aprel"},
	May:                   []string{"may"},
	June:                  []string{"iyn", "iyun"},
	July:                  []string{"iyl", "iyul"},
	August:                []string{"avq", "avqust"},
	September:             []string{"sen", "sentyabr"},
	October:               []string{"okt", "oktyabr"},
	November:              []string{"noy", "noyabr"},
	December:              []string{"dek", "dekabr"},
	Monday:                []string{"bazar ertəsi", "be"},
	Tuesday:               []string{"ça", "çərşənbə axşamı"},
	Wednesday:             []string{"ç", "çərşənbə"},
	Thursday:              []string{"ca", "cümə axşamı"},
	Friday:                []string{"c", "cümə"},
	Saturday:              []string{"ş", "şənbə"},
	Sunday:                []string{"b", "bazar"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"il"},
	Month:                 []string{"ay"},
	Week:                  []string{"həftə"},
	Day:                   []string{"gün"},
	Hour:                  []string{"saat"},
	Minute:                []string{"dəq", "dəqiqə"},
	Second:                []string{"san", "saniyə"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`bu gün`},
		`0 hour ago`:   {`bu saat`},
		`0 minute ago`: {`bu dəqiqə`},
		`0 month ago`:  {`bu ay`},
		`0 second ago`: {`indi`},
		`0 week ago`:   {`bu həftə`},
		`0 year ago`:   {`bu il`},
		`1 day ago`:    {`dünən`},
		`1 month ago`:  {`keçən ay`},
		`1 week ago`:   {`keçən həftə`},
		`1 year ago`:   {`keçən il`},
		`in 1 day`:     {`sabah`},
		`in 1 month`:   {`gələn ay`},
		`in 1 week`:    {`gələn həftə`},
		`in 1 year`:    {`gələn il`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) gün öncə`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) saat öncə`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) dəqiqə öncə`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ay öncə`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) saniyə öncə`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) həftə öncə`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) il öncə`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) gün ərzində`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) saat ərzində`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) dəqiqə ərzində`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ay ərzində`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) saniyə ərzində`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) həftə ərzində`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) il ərzində`),
		},
	},
}
