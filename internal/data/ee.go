package data

import "regexp"

var eeLocale = LocaleData{
	Name:                  "ee",
	DateOrder:             "MDY",
	January:               []string{"dzove", "dzv"},
	February:              []string{"dzd", "dzodze"},
	March:                 []string{"ted", "tedoxe"},
	April:                 []string{"afɔ", "afɔfĩe"},
	May:                   []string{"dam", "dama"},
	June:                  []string{"mas", "masa"},
	July:                  []string{"sia", "siamlɔm"},
	August:                []string{"dea", "deasiamime"},
	September:             []string{"any", "anyɔnyɔ"},
	October:               []string{"kel", "kele"},
	November:              []string{"ade", "adeɛmekpɔxe"},
	December:              []string{"dzm", "dzome"},
	Monday:                []string{"dzo", "dzoɖa"},
	Tuesday:               []string{"bla", "blaɖa"},
	Wednesday:             []string{"kuɖ", "kuɖa"},
	Thursday:              []string{"yaw", "yawoɖa"},
	Friday:                []string{"fiɖ", "fiɖa"},
	Saturday:              []string{"mem", "memleɖa"},
	Sunday:                []string{"kɔs", "kɔsiɖa"},
	AM:                    []string{"ŋdi"},
	PM:                    []string{"ɣetrɔ"},
	Year:                  []string{"ƒe"},
	Month:                 []string{"ɣleti"},
	Week:                  []string{"kɔsiɖa ɖeka"},
	Day:                   []string{"ŋkeke"},
	Hour:                  []string{"gaƒoƒo"},
	Minute:                []string{"aɖabaƒoƒo"},
	Second:                []string{"sekend"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`egbe`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`ɣleti sia`},
		`0 second ago`: {`fifi`},
		`0 week ago`:   {`kɔsiɖa sia`},
		`0 year ago`:   {`ƒe sia`},
		`1 day ago`:    {`etsɔ si va yi`},
		`1 month ago`:  {`ɣleti si va yi`},
		`1 week ago`:   {`kɔsiɖa si va yi`},
		`1 year ago`:   {`ƒe si va yi`},
		`in 1 day`:     {`etsɔ si gbɔna`},
		`in 1 month`:   {`ɣleti si gbɔ na`},
		`in 1 week`:    {`kɔsiɖa si gbɔ na`},
		`in 1 year`:    {`ƒe si gbɔ na`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)ŋkeke (\d+) si va yi`),
			regexp.MustCompile(`(?i)ŋkeke (\d+) si wo va yi`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)gaƒoƒo (\d+) si va yi`),
			regexp.MustCompile(`(?i)gaƒoƒo (\d+) si wo va yi`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)aɖabaƒoƒo (\d+) si va yi`),
			regexp.MustCompile(`(?i)aɖabaƒoƒo (\d+) si wo va yi`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)ɣleti (\d+) si va yi`),
			regexp.MustCompile(`(?i)ɣleti (\d+) si wo va yi`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)sekend (\d+) si va yi`),
			regexp.MustCompile(`(?i)sekend (\d+) si wo va yi`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)kɔsiɖa (\d+) si va yi`),
			regexp.MustCompile(`(?i)kɔsiɖa (\d+) si wo va yi`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)le ƒe (\d+) si va yi me`),
			regexp.MustCompile(`(?i)ƒe (\d+) si va yi`),
			regexp.MustCompile(`(?i)ƒe (\d+) si va yi me`),
			regexp.MustCompile(`(?i)ƒe (\d+) si wo va yi`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)le ŋkeke (\d+) me`),
			regexp.MustCompile(`(?i)le ŋkeke (\d+) wo me`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)le gaƒoƒo (\d+) me`),
			regexp.MustCompile(`(?i)le gaƒoƒo (\d+) wo me`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)le aɖabaƒoƒo (\d+) me`),
			regexp.MustCompile(`(?i)le aɖabaƒoƒo (\d+) wo me`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)le ɣleti (\d+) me`),
			regexp.MustCompile(`(?i)le ɣleti (\d+) wo me`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)le sekend (\d+) me`),
			regexp.MustCompile(`(?i)le sekend (\d+) wo me`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)le kɔsiɖa (\d+) me`),
			regexp.MustCompile(`(?i)le kɔsiɖa (\d+) wo me`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)le ƒe (\d+) me`),
			regexp.MustCompile(`(?i)le ƒe (\d+) si gbɔna me`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ee-TG": eeTGLocale,
	},
}

var eeTGLocale = LocaleData{
	Name:                  "ee-TG",
	DateOrder:             "",
}
