package data

import "regexp"

var dzLocale = LocaleData{
	Name:                  "dz",
	DateOrder:             "YMD",
	January:               []string{"ཟླ་༡", "ཟླ་དངཔ་", "སྤྱི་ཟླ་དངཔ་"},
	February:              []string{"ཟླ་༢", "ཟླ་གཉིས་པ་", "སྤྱི་ཟླ་གཉིས་པ་"},
	March:                 []string{"ཟླ་༣", "ཟླ་གསུམ་པ་", "སྤྱི་ཟླ་གསུམ་པ་"},
	April:                 []string{"ཟླ་༤", "ཟླ་བཞི་པ་", "སྤྱི་ཟླ་བཞི་པ"},
	May:                   []string{"ཟླ་༥", "ཟླ་ལྔ་པ་", "སྤྱི་ཟླ་ལྔ་པ་"},
	June:                  []string{"ཟླ་༦", "ཟླ་དྲུག་པ", "སྤྱི་ཟླ་དྲུག་པ"},
	July:                  []string{"ཟླ་༧", "ཟླ་བདུན་པ་", "སྤྱི་ཟླ་བདུན་པ་"},
	August:                []string{"ཟླ་༨", "ཟླ་བརྒྱད་པ་", "སྤྱི་ཟླ་བརྒྱད་པ་"},
	September:             []string{"ཟླ་༩", "ཟླ་དགུ་པ་", "སྤྱི་ཟླ་དགུ་པ་"},
	October:               []string{"ཟླ་༡༠", "ཟླ་བཅུ་པ་", "སྤྱི་ཟླ་བཅུ་པ་"},
	November:              []string{"ཟླ་༡༡", "ཟླ་བཅུ་གཅིག་པ་", "སྤྱི་ཟླ་བཅུ་གཅིག་པ་"},
	December:              []string{"ཟླ་༡༢", "ཟླ་བཅུ་གཉིས་པ་", "སྤྱི་ཟླ་བཅུ་གཉིས་པ་"},
	Monday:                []string{"གཟའ་མིག་དམར་", "མིར་"},
	Tuesday:               []string{"གཟའ་ལྷག་པ་", "ལྷག་"},
	Wednesday:             []string{"གཟའ་ཕུར་བུ་", "ཕུར་"},
	Thursday:              []string{"གཟའ་པ་སངས་", "སངས་"},
	Friday:                []string{"གཟའ་སྤེན་པ་", "སྤེན་"},
	Saturday:              []string{"གཟའ་ཉི་མ་", "ཉི་"},
	Sunday:                []string{"གཟའ་ཟླ་བ་", "ཟླ་"},
	AM:                    []string{"སྔ་ཆ་"},
	PM:                    []string{"ཕྱི་ཆ་"},
	Year:                  []string{"ལོ"},
	Month:                 []string{"ཟླ་ཝ་"},
	Week:                  []string{"བདུན་ཕྲག"},
	Day:                   []string{"ཚེས་"},
	Hour:                  []string{"ཆུ་ཚོད"},
	Minute:                []string{"སྐར་མ"},
	Second:                []string{"སྐར་ཆཱ་"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ད་རིས་`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`ཁ་ཙ་`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`ནངས་པ་`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)ཉིནམ་ (\d+) ཧེ་མ་`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)ཆུ་ཚོད་ (\d+) ཧེ་མ་`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)སྐར་མ་ (\d+) ཧེ་མ་`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)ཟླཝ་ (\d+) ཧེ་མ་`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)སྐར་ཆ་ (\d+) ཧེ་མ་`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)བངུན་ཕྲག་ (\d+) ཧེ་མ་`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)ལོ་འཁོར་ (\d+) ཧེ་མ་`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)ཉིནམ་ (\d+) ནང་`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)ཆུ་ཚོད་ (\d+) ནང་`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)སྐར་མ་ (\d+) ནང་`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)ཟླཝ་ (\d+) ནང་`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)སྐར་ཆ་ (\d+) ནང་`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)བངུན་ཕྲག་ (\d+) ནང་`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)ལོ་འཁོར་ (\d+) ནང་`),
		},
	},
}
