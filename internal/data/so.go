package data

import "regexp"

var soLocale = LocaleData{
	Name:                  "so",
	DateOrder:             "DMY",
	January:               []string{"bisha koobaad", "kob"},
	February:              []string{"bisha labaad", "lab"},
	March:                 []string{"bisha saddexaad", "sad"},
	April:                 []string{"afr", "bisha afraad"},
	May:                   []string{"bisha shanaad", "sha"},
	June:                  []string{"bisha lixaad", "lix"},
	July:                  []string{"bisha todobaad", "tod"},
	August:                []string{"bisha sideedaad", "sid"},
	September:             []string{"bisha sagaalaad", "sag"},
	October:               []string{"bisha tobnaad", "tob"},
	November:              []string{"bisha kow iyo tobnaad", "kit"},
	December:              []string{"bisha laba iyo tobnaad", "lit"},
	Monday:                []string{"isn", "isniin"},
	Tuesday:               []string{"tal", "talaado"},
	Wednesday:             []string{"arb", "arbaco"},
	Thursday:              []string{"kha", "khamiis"},
	Friday:                []string{"jim", "jimco"},
	Saturday:              []string{"sab", "sabti"},
	Sunday:                []string{"axad", "axd"},
	AM:                    []string{"sn"},
	PM:                    []string{"gn"},
	Year:                  []string{"year"},
	Month:                 []string{"month"},
	Week:                  []string{"week"},
	Day:                   []string{"day"},
	Hour:                  []string{"hour"},
	Minute:                []string{"minute"},
	Second:                []string{"second"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`maanta`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`shalay`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`berri`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"so-DJ": soDJLocale,
		"so-ET": soETLocale,
		"so-KE": soKELocale,
	},
}

var soDJLocale = LocaleData{
	Name:                  "so-DJ",
	DateOrder:             "",
}

var soETLocale = LocaleData{
	Name:                  "so-ET",
	DateOrder:             "",
}

var soKELocale = LocaleData{
	Name:                  "so-KE",
	DateOrder:             "",
}
