package data

import "regexp"

var amLocale = LocaleData{
	Name:                  "am",
	DateOrder:             "DMY",
	January:               []string{"ጃንዩ", "ጃንዩወሪ"},
	February:              []string{"ፌብሩ", "ፌብሩወሪ"},
	March:                 []string{"ማርች"},
	April:                 []string{"ኤፕሪ", "ኤፕሪል"},
	May:                   []string{"ሜይ"},
	June:                  []string{"ጁን"},
	July:                  []string{"ጁላይ"},
	August:                []string{"ኦገስ", "ኦገስት"},
	September:             []string{"ሴፕቴ", "ሴፕቴምበር"},
	October:               []string{"ኦክቶ", "ኦክቶበር"},
	November:              []string{"ኖቬም", "ኖቬምበር"},
	December:              []string{"ዲሴም", "ዲሴምበር"},
	Monday:                []string{"ሰኞ"},
	Tuesday:               []string{"ማክሰ", "ማክሰኞ"},
	Wednesday:             []string{"ረቡዕ"},
	Thursday:              []string{"ሐሙስ"},
	Friday:                []string{"ዓርብ"},
	Saturday:              []string{"ቅዳሜ"},
	Sunday:                []string{"እሑድ"},
	AM:                    []string{"ጥዋት"},
	PM:                    []string{"ከሰዓት"},
	Year:                  []string{"ዓመት"},
	Month:                 []string{"ወር"},
	Week:                  []string{"ሳምንት"},
	Day:                   []string{"ቀን"},
	Hour:                  []string{"ሰዓት"},
	Minute:                []string{"ደቂቃ"},
	Second:                []string{"ሰከንድ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ዛሬ`},
		`0 hour ago`:   {`ይህ ሰዓት`},
		`0 minute ago`: {`ይህ ደቂቃ`},
		`0 month ago`:  {`በዚህ ወር`},
		`0 second ago`: {`አሁን`},
		`0 week ago`:   {`በዚህ ሣምንት`, `በዚህ ሳምንት`},
		`0 year ago`:   {`በዚህ ዓመት`},
		`1 day ago`:    {`ትላንትና`, `ትናንት`},
		`1 month ago`:  {`ያለፈው ወር`},
		`1 week ago`:   {`ባለፈው ሳምንት`, `ያለፈው ሳምንት`},
		`1 year ago`:   {`ያለፈው ዓመት`},
		`in 1 day`:     {`ነገ`},
		`in 1 month`:   {`የሚቀጥለው ወር`},
		`in 1 week`:    {`የሚቀጥለው ሳምንት`},
		`in 1 year`:    {`የሚቀጥለው ዓመት`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)ከ (\d+) ቀን በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ቀናት በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ቀን በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ቀኖች በፊት`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ሰዓት በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ሰዓቶች በፊት`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ደቂቃ በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ደቂቃዎች በፊት`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ወራት በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ወር በፊት`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ሰከንድ በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ሰከንዶች በፊት`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ሳምንታት በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ሳምንት በፊት`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)ከ(\d+) ዓመታት በፊት`),
			regexp.MustCompile(`(?i)ከ(\d+) ዓመት በፊት`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)በ(\d+) ቀናት ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ቀን ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ቀኖች ውስጥ`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)በ(\d+) ሰዓት ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ሰዓቶች ውስጥ`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)በ(\d+) ደቂቃ ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ደቂቃዎች ውስጥ`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)በ(\d+) ወራት ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ወር ውስጥ`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)በ(\d+) ሰከንድ ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ሰከንዶች ውስጥ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)በ(\d+) ሳምንታት ውስጥ`),
			regexp.MustCompile(`(?i)በ(\d+) ሳምንት ውስጥ`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)በ(\d+) ዓመታት ውስጥ`),
		},
	},
}
