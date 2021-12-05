package data

import "regexp"

var paGuruLocale = LocaleData{
	Name:                  "pa-Guru",
	DateOrder:             "DMY",
	January:               []string{"ਜਨ", "ਜਨਵਰੀ"},
	February:              []string{"ਫ਼ਰ", "ਫ਼ਰਵਰੀ"},
	March:                 []string{"ਮਾਰਚ"},
	April:                 []string{"ਅਪ੍ਰੈ", "ਅਪ੍ਰੈਲ"},
	May:                   []string{"ਮਈ"},
	June:                  []string{"ਜੂਨ"},
	July:                  []string{"ਜੁਲਾ", "ਜੁਲਾਈ"},
	August:                []string{"ਅਗ", "ਅਗਸਤ"},
	September:             []string{"ਸਤੰ", "ਸਤੰਬਰ"},
	October:               []string{"ਅਕਤੂ", "ਅਕਤੂਬਰ"},
	November:              []string{"ਨਵੰ", "ਨਵੰਬਰ"},
	December:              []string{"ਦਸੰ", "ਦਸੰਬਰ"},
	Monday:                []string{"ਸੋਮ", "ਸੋਮਵਾਰ"},
	Tuesday:               []string{"ਮੰਗਲ", "ਮੰਗਲਵਾਰ"},
	Wednesday:             []string{"ਬੁੱਧ", "ਬੁੱਧਵਾਰ"},
	Thursday:              []string{"ਵੀਰ", "ਵੀਰਵਾਰ"},
	Friday:                []string{"ਸ਼ੁੱਕਰ", "ਸ਼ੁੱਕਰਵਾਰ"},
	Saturday:              []string{"ਸ਼ਨਿੱਚਰ", "ਸ਼ਨਿੱਚਰਵਾਰ"},
	Sunday:                []string{"ਐਤ", "ਐਤਵਾਰ"},
	AM:                    []string{"ਪੂਦੁ"},
	PM:                    []string{"ਬਾਦੁ"},
	Year:                  []string{"ਸਾਲ"},
	Month:                 []string{"ਮਹੀਨਾ"},
	Week:                  []string{"ਹਫ਼ਤਾ"},
	Day:                   []string{"ਦਿਨ"},
	Hour:                  []string{"ਘੰ", "ਘੰਟਾ"},
	Minute:                []string{"ਮਿੰਟ"},
	Second:                []string{"ਸਕਿੰਟ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ਅੱਜ`},
		`0 hour ago`:   {`ਇਸ ਘੰਟੇ`},
		`0 minute ago`: {`ਇਸ ਮਿੰਟ`},
		`0 month ago`:  {`ਇਹ ਮਹੀਨਾ`},
		`0 second ago`: {`ਹੁਣ`},
		`0 week ago`:   {`ਇਹ ਹਫ਼ਤਾ`},
		`0 year ago`:   {`ਇਹ ਸਾਲ`},
		`1 day ago`:    {`ਬੀਤਿਆ ਕੱਲ੍ਹ`},
		`1 month ago`:  {`ਪਿਛਲਾ ਮਹੀਨਾ`},
		`1 week ago`:   {`ਪਿਛਲਾ ਹਫ਼ਤਾ`},
		`1 year ago`:   {`ਪਿਛਲਾ ਸਾਲ`},
		`in 1 day`:     {`ਭਲਕੇ`},
		`in 1 month`:   {`ਅਗਲਾ ਮਹੀਨਾ`},
		`in 1 week`:    {`ਅਗਲਾ ਹਫ਼ਤਾ`},
		`in 1 year`:    {`ਅਗਲਾ ਸਾਲ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ਦਿਨ ਪਹਿਲਾਂ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ਘੰਟਾ ਪਹਿਲਾਂ`),
			regexp.MustCompile(`(?i)(\d+) ਘੰਟੇ ਪਹਿਲਾਂ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) ਮਿੰਟ ਪਹਿਲਾਂ`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ਮਹੀਨਾ ਪਹਿਲਾਂ`),
			regexp.MustCompile(`(?i)(\d+) ਮਹੀਨੇ ਪਹਿਲਾਂ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) ਸਕਿੰਟ ਪਹਿਲਾਂ`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ਹਫ਼ਤਾ ਪਹਿਲਾਂ`),
			regexp.MustCompile(`(?i)(\d+) ਹਫ਼ਤੇ ਪਹਿਲਾਂ`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ਸਾਲ ਪਹਿਲਾਂ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) ਦਿਨ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਦਿਨਾਂ ਵਿੱਚ`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ਘੰਟਿਆਂ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਘੰਟੇ ਵਿੱਚ`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) ਮਿੰਟ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਮਿੰਟਾਂ ਵਿੱਚ`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ਮਹੀਨਿਆਂ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਮਹੀਨੇ ਵਿੱਚ`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) ਸਕਿੰਟ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਸਕਿੰਟਾਂ ਵਿੱਚ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ਹਫ਼ਤਿਆਂ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਹਫ਼ਤੇ ਵਿੱਚ`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) ਸਾਲ ਵਿੱਚ`),
			regexp.MustCompile(`(?i)(\d+) ਸਾਲਾਂ ਵਿੱਚ`),
		},
	},
}
