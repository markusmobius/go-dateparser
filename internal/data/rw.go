package data

import "regexp"

var rwLocale = LocaleData{
	Name:                  "rw",
	DateOrder:             "YMD",
	January:               []string{"mut", "mutarama"},
	February:              []string{"gas", "gashyantare"},
	March:                 []string{"wer", "werurwe"},
	April:                 []string{"mat", "mata"},
	May:                   []string{"gic", "gicuransi"},
	June:                  []string{"kam", "kamena"},
	July:                  []string{"nya", "nyakanga"},
	August:                []string{"kan", "kanama"},
	September:             []string{"nze", "nzeli"},
	October:               []string{"ukw", "ukwakira"},
	November:              []string{"ugu", "ugushyingo"},
	December:              []string{"uku", "ukuboza"},
	Monday:                []string{"kuwa mbere", "mbe"},
	Tuesday:               []string{"kab", "kuwa kabiri"},
	Wednesday:             []string{"gtu", "kuwa gatatu"},
	Thursday:              []string{"kan", "kuwa kane"},
	Friday:                []string{"gnu", "kuwa gatanu"},
	Saturday:              []string{"gnd", "kuwa gatandatu"},
	Sunday:                []string{"cyu", "ku cyumweru"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"year"},
	Month:                 []string{"month"},
	Week:                  []string{"week"},
	Day:                   []string{"day"},
	Hour:                  []string{"hour"},
	Minute:                []string{"minute"},
	Second:                []string{"second"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`today`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`yesterday`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`tomorrow`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
}
