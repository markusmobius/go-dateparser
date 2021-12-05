package data

import "regexp"

var myLocale = LocaleData{
	Name:                  "my",
	DateOrder:             "DMY",
	January:               []string{"ဇန်", "ဇန်နဝါရီ"},
	February:              []string{"ဖေ", "ဖေဖော်ဝါရီ"},
	March:                 []string{"မတ်"},
	April:                 []string{"ဧ", "ဧပြီ"},
	May:                   []string{"မေ"},
	June:                  []string{"ဇွန်"},
	July:                  []string{"ဇူ", "ဇူလိုင်"},
	August:                []string{"ဩ", "ဩဂုတ်"},
	September:             []string{"စက်", "စက်တင်ဘာ"},
	October:               []string{"အောက်", "အောက်တိုဘာ"},
	November:              []string{"နို", "နိုဝင်ဘာ"},
	December:              []string{"ဒီ", "ဒီဇင်ဘာ"},
	Monday:                []string{"တနင်္လာ"},
	Tuesday:               []string{"အင်္ဂါ"},
	Wednesday:             []string{"ဗုဒ္ဓဟူး"},
	Thursday:              []string{"ကြာသပတေး"},
	Friday:                []string{"သောကြာ"},
	Saturday:              []string{"စနေ"},
	Sunday:                []string{"တနင်္ဂနွေ"},
	AM:                    []string{"နံနက်"},
	PM:                    []string{"ညနေ"},
	Year:                  []string{"နှစ်"},
	Month:                 []string{"လ"},
	Week:                  []string{"ပတ်"},
	Day:                   []string{"ရက်"},
	Hour:                  []string{"နာရီ"},
	Minute:                []string{"မိနစ်"},
	Second:                []string{"စက္ကန့်"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ယနေ့`},
		`0 hour ago`:   {`ဤအချိန်`},
		`0 minute ago`: {`ဤမိနစ်`},
		`0 month ago`:  {`ယခုလ`},
		`0 second ago`: {`ယခု`},
		`0 week ago`:   {`ယခု သီတင်းပတ်`},
		`0 year ago`:   {`ယခုနှစ်`},
		`1 day ago`:    {`မနေ့က`},
		`1 month ago`:  {`ပြီးခဲ့သည့်လ`},
		`1 week ago`:   {`ပြီးခဲ့သည့် သီတင်းပတ်`},
		`1 year ago`:   {`ယမန်နှစ်`},
		`in 1 day`:     {`မနက်ဖြန်`},
		`in 1 month`:   {`လာမည့်လ`},
		`in 1 week`:    {`လာမည့် သီတင်းပတ်`},
		`in 1 year`:    {`လာမည့်နှစ်`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) ရက်`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) နာရီ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) မိနစ်`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) လ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) စက္ကန့်`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) ပတ်`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)ပြီးခဲ့သည့် (\d+) နှစ်`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) ရက်အတွင်း`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) နာရီအတွင်း`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) မိနစ်အတွင်း`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) လအတွင်း`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) စက္ကန့်အတွင်း`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ပတ်အတွင်း`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) နှစ်အတွင်း`),
		},
	},
}
