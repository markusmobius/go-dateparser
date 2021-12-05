package data

import "regexp"

var ksLocale = LocaleData{
	Name:                  "ks",
	DateOrder:             "MDY",
	January:               []string{"جنؤری"},
	February:              []string{"فرؤری"},
	March:                 []string{"مارٕچ"},
	April:                 []string{"اپریل"},
	May:                   []string{"میٔ"},
	June:                  []string{"جوٗن"},
	July:                  []string{"جوٗلایی"},
	August:                []string{"اگست"},
	September:             []string{"ستمبر"},
	October:               []string{"اکتوٗبر"},
	November:              []string{"نومبر"},
	December:              []string{"دسمبر"},
	Monday:                []string{"ژٔنٛدرٕروار", "ژٔنٛدٕروار"},
	Tuesday:               []string{"بوٚموار"},
	Wednesday:             []string{"بودوار"},
	Thursday:              []string{"برٛٮ۪سوار"},
	Friday:                []string{"جُمہ"},
	Saturday:              []string{"بٹوار"},
	Sunday:                []string{"آتھوار", "اَتھوار"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ؤری"},
	Month:                 []string{"رٮ۪تھ"},
	Week:                  []string{"ہفتہٕ"},
	Day:                   []string{"دۄہ"},
	Hour:                  []string{"گٲنٛٹہٕ"},
	Minute:                []string{"مِنَٹ"},
	Second:                []string{"سٮ۪کَنڑ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`اَز`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`راتھ`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`پگاہ`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
}
