package data

import "regexp"

var nusLocale = LocaleData{
	Name:                  "nus",
	DateOrder:             "DMY",
	January:               []string{"tiop", "tiop thar pɛt"},
	February:              []string{"pɛt"},
	March:                 []string{"duɔ̱ɔ̱", "duɔ̱ɔ̱ŋ"},
	April:                 []string{"guak"},
	May:                   []string{"duä", "duät"},
	June:                  []string{"kor", "kornyoot"},
	July:                  []string{"pay", "pay yie̱tni"},
	August:                []string{"thoo", "tho̱o̱r"},
	September:             []string{"tɛɛ", "tɛɛr"},
	October:               []string{"laa", "laath"},
	November:              []string{"kur"},
	December:              []string{"tid", "tio̱p in di̱i̱t"},
	Monday:                []string{"jiec", "jiec la̱t"},
	Tuesday:               []string{"rɛw", "rɛw lätni"},
	Wednesday:             []string{"diɔ̱k", "diɔ̱k lätni"},
	Thursday:              []string{"ŋuaan", "ŋuaan lätni"},
	Friday:                []string{"dhieec", "dhieec lätni"},
	Saturday:              []string{"bäkɛl", "bäkɛl lätni"},
	Sunday:                []string{"cäŋ", "cäŋ kuɔth"},
	AM:                    []string{"rw"},
	PM:                    []string{"tŋ"},
	Year:                  []string{"ruɔ̱n"},
	Month:                 []string{"pay"},
	Week:                  []string{"jiɔk"},
	Day:                   []string{"cäŋ"},
	Hour:                  []string{"thaak"},
	Minute:                []string{"minit"},
	Second:                []string{"thɛkɛni"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`walɛ`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`pan`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`ruun`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
}
