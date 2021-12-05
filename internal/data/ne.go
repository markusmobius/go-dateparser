package data

import "regexp"

var neLocale = LocaleData{
	Name:                  "ne",
	DateOrder:             "YMD",
	January:               []string{"जनवरी"},
	February:              []string{"फेब्रुअरी"},
	March:                 []string{"मार्च"},
	April:                 []string{"अप्रिल"},
	May:                   []string{"मई", "मे"},
	June:                  []string{"जुन"},
	July:                  []string{"जुलाई"},
	August:                []string{"अगस्ट"},
	September:             []string{"सेप्टेम्बर"},
	October:               []string{"अक्टोबर"},
	November:              []string{"नोभेम्बर"},
	December:              []string{"डिसेम्बर"},
	Monday:                []string{"सोम", "सोमबार"},
	Tuesday:               []string{"मङ्गल", "मङ्गलबार"},
	Wednesday:             []string{"बुध", "बुधबार"},
	Thursday:              []string{"बिहि", "बिहिबार"},
	Friday:                []string{"शुक्र", "शुक्रबार"},
	Saturday:              []string{"शनि", "शनिबार"},
	Sunday:                []string{"आइत", "आइतबार"},
	AM:                    []string{"पूर्वाह्न"},
	PM:                    []string{"अपराह्न"},
	Year:                  []string{"बर्ष", "वर्ष"},
	Month:                 []string{"महिना"},
	Week:                  []string{"हप्ता"},
	Day:                   []string{"बार"},
	Hour:                  []string{"घण्टा"},
	Minute:                []string{"मिनेट"},
	Second:                []string{"सेकेन्ड"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`आज`},
		`0 hour ago`:   {`यो घडीमा`},
		`0 minute ago`: {`यही मिनेटमा`},
		`0 month ago`:  {`यो महिना`},
		`0 second ago`: {`अब`},
		`0 week ago`:   {`यो हप्ता`},
		`0 year ago`:   {`यो वर्ष`},
		`1 day ago`:    {`हिजो`},
		`1 month ago`:  {`गत महिना`},
		`1 week ago`:   {`गत हप्ता`},
		`1 year ago`:   {`गत वर्ष`},
		`in 1 day`:     {`भोलि`},
		`in 1 month`:   {`अर्को महिना`},
		`in 1 week`:    {`आउने हप्ता`},
		`in 1 year`:    {`अर्को वर्ष`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) दिन पहिले`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) घण्टा पहिले`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) मिनेट पहिले`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) महिना पहिले`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) सेकेण्ड पहिले`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) हप्ता पहिले`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) वर्ष अघि`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) दिनमा`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) घण्टामा`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) मिनेटमा`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) महिनामा`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) सेकेण्डमा`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) हप्तामा`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) वर्षमा`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ne-IN": neINLocale,
	},
}

var neINLocale = LocaleData{
	Name:                  "ne-IN",
	DateOrder:             "",
}
