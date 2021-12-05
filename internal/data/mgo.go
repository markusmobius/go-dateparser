package data

import "regexp"

var mgoLocale = LocaleData{
	Name:                  "mgo",
	DateOrder:             "YMD",
	January:               []string{"iməg mbegtug", "mbegtug"},
	February:              []string{"imeg àbùbì"},
	March:                 []string{"imeg mbəŋchubi"},
	April:                 []string{"iməg ngwə̀t"},
	May:                   []string{"iməg fog"},
	June:                  []string{"iməg ichiibɔd"},
	July:                  []string{"iməg àdùmbə̀ŋ"},
	August:                []string{"iməg ichika"},
	September:             []string{"iməg kud"},
	October:               []string{"iməg tèsi'e"},
	November:              []string{"iməg zò"},
	December:              []string{"iməg krizmed"},
	Monday:                []string{"aneg 2"},
	Tuesday:               []string{"aneg 3"},
	Wednesday:             []string{"aneg 4"},
	Thursday:              []string{"aneg 5"},
	Friday:                []string{"aneg 6"},
	Saturday:              []string{"aneg 7"},
	Sunday:                []string{"aneg 1"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"fitu'"},
	Month:                 []string{"iməg"},
	Week:                  []string{"nkap"},
	Day:                   []string{"anəg"},
	Hour:                  []string{"hour"},
	Minute:                []string{"minute"},
	Second:                []string{"second"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`tèchɔ̀ŋ`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`ikwiri`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`isu`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
}
