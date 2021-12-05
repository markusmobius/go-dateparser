package data

import "regexp"

var siLocale = LocaleData{
	Name:                  "si",
	DateOrder:             "YMD",
	January:               []string{"ජන", "ජනවාරි"},
	February:              []string{"පෙබ", "පෙබරවාරි"},
	March:                 []string{"මාර්", "මාර්තු"},
	April:                 []string{"අප්‍රේල්"},
	May:                   []string{"මැයි"},
	June:                  []string{"ජූනි"},
	July:                  []string{"ජූලි"},
	August:                []string{"අගෝ", "අගෝස්තු"},
	September:             []string{"සැප්", "සැප්තැම්බර්"},
	October:               []string{"ඔක්", "ඔක්තෝබර්"},
	November:              []string{"නොවැ", "නොවැම්බර්"},
	December:              []string{"දෙසැ", "දෙසැම්බර්"},
	Monday:                []string{"සඳුදා"},
	Tuesday:               []string{"අඟහ", "අඟහරුවාදා"},
	Wednesday:             []string{"බදාදා"},
	Thursday:              []string{"බ්‍රහස්", "බ්‍රහස්පතින්දා"},
	Friday:                []string{"සිකු", "සිකුරාදා"},
	Saturday:              []string{"සෙන", "සෙනසුරාදා"},
	Sunday:                []string{"ඉරිදා"},
	AM:                    []string{"පෙව"},
	PM:                    []string{"පව"},
	Year:                  []string{"වර්", "වර්ෂය"},
	Month:                 []string{"මාස", "මාසය"},
	Week:                  []string{"සති", "සතිය"},
	Day:                   []string{"දිනය"},
	Hour:                  []string{"පැ", "පැය"},
	Minute:                []string{"මි", "මිනි", "මිනිත්තුව"},
	Second:                []string{"ත", "තත්", "තත්පරය"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`අද`},
		`0 hour ago`:   {`මෙම පැය`},
		`0 minute ago`: {`මෙම මිනිත්තුව`},
		`0 month ago`:  {`මෙම මාස`, `මෙම මාසය`},
		`0 second ago`: {`දැන්`},
		`0 week ago`:   {`මෙම සති`, `මෙම සතිය`},
		`0 year ago`:   {`මෙම වසර`},
		`1 day ago`:    {`ඊයේ`},
		`1 month ago`:  {`පසුගිය මාස`, `පසුගිය මාසය`},
		`1 week ago`:   {`පසුගිය සති`, `පසුගිය සතිය`},
		`1 year ago`:   {`පසුගිය වසර`},
		`in 1 day`:     {`හෙට`},
		`in 1 month`:   {`ඊළඟ මාස`, `ඊළඟ මාසය`},
		`in 1 week`:    {`ඊළඟ සති`, `ඊළඟ සතිය`},
		`in 1 year`:    {`ඊළඟ වසර`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)දින (\d+)කට පෙර`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)පැය (\d+)කට පෙර`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)මිනිත්තු (\d+)කට පෙර`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)මාස (\d+)කට පෙර`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)තත්පර (\d+)කට පෙර`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)සති (\d+)කට පෙර`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)වසර (\d+)කට පෙර`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)දින (\d+)න්`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)පැය (\d+)කින්`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)මිනිත්තු (\d+)කින්`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)මාස (\d+)කින්`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)තත්පර (\d+)කින්`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)සති (\d+)කින්`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)වසර (\d+)කින්`),
		},
	},
}
