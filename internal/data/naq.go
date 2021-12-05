package data

import "regexp"

var naqLocale = LocaleData{
	Name:                  "naq",
	DateOrder:             "DMY",
	January:               []string{"jan", "ǃkhanni"},
	February:              []string{"feb", "ǃkhanǀgôab"},
	March:                 []string{"mar", "ǀkhuuǁkhâb"},
	April:                 []string{"apr", "ǃhôaǂkhaib"},
	May:                   []string{"may", "ǃkhaitsâb"},
	June:                  []string{"gamaǀaeb", "jun"},
	July:                  []string{"jul", "ǂkhoesaob"},
	August:                []string{"aoǁkhuumûǁkhâb", "aug"},
	September:             []string{"sep", "taraǀkhuumûǁkhâb"},
	October:               []string{"oct", "ǂnûǁnâiseb"},
	November:              []string{"nov", "ǀhooǂgaeb"},
	December:              []string{"dec", "hôasoreǁkhâb"},
	Monday:                []string{"ma", "mantaxtsees"},
	Tuesday:               []string{"de", "denstaxtsees"},
	Wednesday:             []string{"wu", "wunstaxtsees"},
	Thursday:              []string{"do", "dondertaxtsees"},
	Friday:                []string{"fr", "fraitaxtsees"},
	Saturday:              []string{"sat", "satertaxtsees"},
	Sunday:                []string{"son", "sontaxtsees"},
	AM:                    []string{"ǁgoagas"},
	PM:                    []string{"ǃuias"},
	Year:                  []string{"kurib"},
	Month:                 []string{"ǁkhâb"},
	Week:                  []string{"wekheb"},
	Day:                   []string{"tsees"},
	Hour:                  []string{"iiri"},
	Minute:                []string{"haib"},
	Second:                []string{"ǀgâub"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`neetsee`},
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
