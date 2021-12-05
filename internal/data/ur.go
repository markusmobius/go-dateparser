package data

import "regexp"

var urLocale = LocaleData{
	Name:                  "ur",
	DateOrder:             "DMY",
	January:               []string{"جنوری"},
	February:              []string{"فروری"},
	March:                 []string{"مارچ"},
	April:                 []string{"اپریل"},
	May:                   []string{"مئی"},
	June:                  []string{"جون"},
	July:                  []string{"جولائی"},
	August:                []string{"اگست"},
	September:             []string{"ستمبر"},
	October:               []string{"اکتوبر"},
	November:              []string{"نومبر"},
	December:              []string{"دسمبر"},
	Monday:                []string{"سوموار"},
	Tuesday:               []string{"منگل"},
	Wednesday:             []string{"بدھ"},
	Thursday:              []string{"جمعرات"},
	Friday:                []string{"جمعہ"},
	Saturday:              []string{"ہفتہ"},
	Sunday:                []string{"اتوار"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"سال"},
	Month:                 []string{"ماہ", "مہینہ"},
	Week:                  []string{"ہفتہ"},
	Day:                   []string{"دن"},
	Hour:                  []string{"گھنٹہ"},
	Minute:                []string{"منٹ"},
	Second:                []string{"سیکنڈ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`آج`},
		`0 hour ago`:   {`اس گھنٹے`},
		`0 minute ago`: {`اس منٹ`},
		`0 month ago`:  {`اس مہینہ`},
		`0 second ago`: {`اب`},
		`0 week ago`:   {`اس ہفتہ`},
		`0 year ago`:   {`اس سال`},
		`1 day ago`:    {`گزشتہ کل`},
		`1 month ago`:  {`پچھلے مہینہ`},
		`1 week ago`:   {`پچھلے ہفتہ`},
		`1 year ago`:   {`گزشتہ سال`},
		`in 1 day`:     {`آئندہ کل`},
		`in 1 month`:   {`اگلے مہینہ`},
		`in 1 week`:    {`اگلے ہفتہ`},
		`in 1 year`:    {`اگلے سال`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) دن پہلے`),
			regexp.MustCompile(`(?i)(\d+) دنوں پہلے`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) گھنٹہ پہلے`),
			regexp.MustCompile(`(?i)(\d+) گھنٹے پہلے`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) منٹ پہلے`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ماہ قبل`),
			regexp.MustCompile(`(?i)(\d+) ماہ پہلے`),
			regexp.MustCompile(`(?i)(\d+) مہینہ پہلے`),
			regexp.MustCompile(`(?i)(\d+) مہینے پہلے`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) سیکنڈ پہلے`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ہفتہ پہلے`),
			regexp.MustCompile(`(?i)(\d+) ہفتے پہلے`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) سال پہلے`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) دن میں`),
			regexp.MustCompile(`(?i)(\d+) دنوں میں`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) گھنٹوں میں`),
			regexp.MustCompile(`(?i)(\d+) گھنٹہ میں`),
			regexp.MustCompile(`(?i)(\d+) گھنٹے میں`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) منٹ میں`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ماہ میں`),
			regexp.MustCompile(`(?i)(\d+) مہینہ میں`),
			regexp.MustCompile(`(?i)(\d+) مہینے میں`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) سیکنڈ میں`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ہفتہ میں`),
			regexp.MustCompile(`(?i)(\d+) ہفتے میں`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) سال میں`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ur-IN": urINLocale,
	},
}

var urINLocale = LocaleData{
	Name:                  "ur-IN",
	DateOrder:             "",
	Monday:                []string{"پیر"},
	RelativeType: map[string][]string{
		`0 month ago`: {`اس ماہ`},
		`1 month ago`: {`گزشتہ ماہ`},
		`1 week ago`:  {`گزشتہ ہفتہ`},
		`in 1 month`:  {`اگلے ماہ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) دن قبل`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) گھنٹہ قبل`),
			regexp.MustCompile(`(?i)(\d+) گھنٹے قبل`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) منٹ قبل`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) سیکنڈ قبل`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ہفتہ قبل`),
			regexp.MustCompile(`(?i)(\d+) ہفتے قبل`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) سالوں پہلے`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ہفتوں میں`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) سالوں میں`),
		},
	},
}
