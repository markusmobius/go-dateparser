package data

import "regexp"

var yoLocale = LocaleData{
	Name:                  "yo",
	DateOrder:             "DMY",
	January:               []string{"oṣù ṣẹ́rẹ́", "ṣẹ́rẹ́"},
	February:              []string{"oṣù èrèlè", "èrèlè"},
	March:                 []string{"oṣù ẹrẹ̀nà", "ẹrẹ̀nà"},
	April:                 []string{"oṣù ìgbé", "ìgbé"},
	May:                   []string{"oṣù ẹ̀bibi", "ẹ̀bibi"},
	June:                  []string{"oṣù òkúdu", "òkúdu"},
	July:                  []string{"agẹmọ", "oṣù agẹmọ"},
	August:                []string{"oṣù ògún", "ògún"},
	September:             []string{"owewe", "oṣù owewe"},
	October:               []string{"oṣù ọ̀wàrà", "ọ̀wàrà"},
	November:              []string{"bélú", "oṣù bélú"},
	December:              []string{"oṣù ọ̀pẹ̀", "ọ̀pẹ̀"},
	Monday:                []string{"ajé", "ọjọ́ ajé"},
	Tuesday:               []string{"ìsẹ́gun", "ọjọ́ ìsẹ́gun"},
	Wednesday:             []string{"ọjọ́rú"},
	Thursday:              []string{"ọjọ́bọ"},
	Friday:                []string{"ẹtì", "ọjọ́ ẹtì"},
	Saturday:              []string{"àbámẹ́ta", "ọjọ́ àbámẹ́ta"},
	Sunday:                []string{"àìkú", "ọjọ́ àìkú"},
	AM:                    []string{"àárọ̀"},
	PM:                    []string{"ọ̀sán"},
	Year:                  []string{"ọdún"},
	Month:                 []string{"osù"},
	Week:                  []string{"ọ̀sè"},
	Day:                   []string{"ọjọ́"},
	Hour:                  []string{"wákàtí"},
	Minute:                []string{"ìsẹ́jú"},
	Second:                []string{"ìsẹ́jú ààyá"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`òní`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`àná`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`ọ̀la`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	LocaleSpecific: map[string]LocaleData{
		"yo-BJ": yoBJLocale,
	},
}

var yoBJLocale = LocaleData{
	Name:                  "yo-BJ",
	DateOrder:             "",
	January:               []string{"oshù shɛ́rɛ́", "shɛ́rɛ́"},
	February:              []string{"oshù èrèlè"},
	March:                 []string{"oshù ɛrɛ̀nà", "ɛrɛ̀nà"},
	April:                 []string{"oshù ìgbé"},
	May:                   []string{"oshù ɛ̀bibi", "ɛ̀bibi"},
	June:                  []string{"oshù òkúdu"},
	July:                  []string{"agɛmɔ", "oshù agɛmɔ"},
	August:                []string{"oshù ògún"},
	September:             []string{"oshù owewe"},
	October:               []string{"oshù ɔ̀wàrà", "ɔ̀wàrà"},
	November:              []string{"oshù bélú"},
	December:              []string{"oshù ɔ̀pɛ̀", "ɔ̀pɛ̀"},
	Monday:                []string{"ɔjɔ́ ajé"},
	Tuesday:               []string{"ìsɛ́gun", "ɔjɔ́ ìsɛ́gun"},
	Wednesday:             []string{"ɔjɔ́rú"},
	Thursday:              []string{"ɔjɔ́bɔ"},
	Friday:                []string{"ɔjɔ́ ɛtì", "ɛtì"},
	Saturday:              []string{"àbámɛ́ta", "ɔjɔ́ àbámɛ́ta"},
	Sunday:                []string{"ɔjɔ́ àìkú"},
	AM:                    []string{"àárɔ̀"},
	PM:                    []string{"ɔ̀sán"},
	Year:                  []string{"ɔdún"},
	Week:                  []string{"ɔ̀sè"},
	Day:                   []string{"ɔjɔ́"},
	Minute:                []string{"ìsɛ́jú"},
	Second:                []string{"ìsɛ́jú ààyá"},
	RelativeType: map[string][]string{
		`in 1 day`: {`ɔ̀la`},
	},
}
