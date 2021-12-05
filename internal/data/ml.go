package data

import "regexp"

var mlLocale = LocaleData{
	Name:                  "ml",
	DateOrder:             "DMY",
	January:               []string{"ജനു", "ജനുവരി"},
	February:              []string{"ഫെബ്രു", "ഫെബ്രുവരി"},
	March:                 []string{"മാർ", "മാർച്ച്"},
	April:                 []string{"ഏപ്രി", "ഏപ്രിൽ"},
	May:                   []string{"മേയ്"},
	June:                  []string{"ജൂൺ"},
	July:                  []string{"ജൂലൈ"},
	August:                []string{"ഓഗ", "ഓഗസ്റ്റ്"},
	September:             []string{"സെപ്റ്റം", "സെപ്റ്റംബർ"},
	October:               []string{"ഒക്ടോ", "ഒക്‌ടോബർ"},
	November:              []string{"നവം", "നവംബർ"},
	December:              []string{"ഡിസം", "ഡിസംബർ"},
	Monday:                []string{"തിങ്കളാഴ്‌ച", "തിങ്കൾ"},
	Tuesday:               []string{"ചൊവ്വ", "ചൊവ്വാഴ്ച", "ചൊവ്വാഴ്‌ച"},
	Wednesday:             []string{"ബുധനാഴ്‌ച", "ബുധൻ"},
	Thursday:              []string{"വ്യാഴം", "വ്യാഴാഴ്‌ച"},
	Friday:                []string{"വെള്ളി", "വെള്ളിയാഴ്‌ച"},
	Saturday:              []string{"ശനി", "ശനിയാഴ്‌ച"},
	Sunday:                []string{"ഞായറാഴ്‌ച", "ഞായർ"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"വ", "വർഷം"},
	Month:                 []string{"മാ", "മാസം"},
	Week:                  []string{"ആ", "ആഴ്ച"},
	Day:                   []string{"ദിവസം"},
	Hour:                  []string{"മ", "മണിക്കൂർ"},
	Minute:                []string{"മി", "മിനിറ്റ്"},
	Second:                []string{"സെ", "സെക്കൻഡ്"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ഇന്ന്`},
		`0 hour ago`:   {`ഈ മണിക്കൂറിൽ`},
		`0 minute ago`: {`ഈ മിനിറ്റിൽ`},
		`0 month ago`:  {`ഈ മാസം`},
		`0 second ago`: {`ഇപ്പോൾ`},
		`0 week ago`:   {`ഈ ആഴ്ച`},
		`0 year ago`:   {`ഈ വർ‌ഷം`},
		`1 day ago`:    {`ഇന്നലെ`},
		`1 month ago`:  {`കഴിഞ്ഞ മാസം`},
		`1 week ago`:   {`കഴിഞ്ഞ ആഴ്‌ച`},
		`1 year ago`:   {`കഴിഞ്ഞ വർഷം`},
		`in 1 day`:     {`നാളെ`},
		`in 1 month`:   {`അടുത്ത മാസം`},
		`in 1 week`:    {`അടുത്ത ആഴ്ച`},
		`in 1 year`:    {`അടുത്തവർഷം`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ദിവസം മുമ്പ്`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) മണിക്കൂർ മുമ്പ്`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) മിനിറ്റ് മുമ്പ്`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) മാസം മുമ്പ്`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) സെക്കൻഡ് മുമ്പ്`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ആഴ്ച മുമ്പ്`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) വർഷം മുമ്പ്`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) ദിവസത്തിൽ`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) മണിക്കൂറിൽ`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) മിനിറ്റിൽ`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) മാസത്തിൽ`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) സെക്കൻഡിൽ`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ആഴ്ചയിൽ`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) വർഷത്തിൽ`),
		},
	},
}
