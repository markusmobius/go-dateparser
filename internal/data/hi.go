package data

import "regexp"

var hiLocale = LocaleData{
	Name:                  "hi",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 3,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "के", "को", "बजे", "सन्", "से", "，"},
	January:               []string{"जनवरी", "जन॰"},
	February:              []string{"फ़रवरी", "फ़र॰"},
	March:                 []string{"मार्च"},
	April:                 []string{"अप्रैल"},
	May:                   []string{"मई"},
	June:                  []string{"जून"},
	July:                  []string{"जुलाई", "जुल॰"},
	August:                []string{"अगस्त", "अग॰"},
	September:             []string{"सितंबर", "सितम्बर", "सित॰"},
	October:               []string{"अक्टूबर", "अक्तूबर", "अक्तू॰"},
	November:              []string{"नवंबर", "नवम्बर", "नव॰"},
	December:              []string{"दिसंबर", "दिसम्बर", "दिस॰"},
	Monday:                []string{"सोम", "सोमवार"},
	Tuesday:               []string{"मंगल", "मंगलवार"},
	Wednesday:             []string{"बुध", "बुधवार"},
	Thursday:              []string{"गुरु", "गुरुवार"},
	Friday:                []string{"शुक्र", "शुक्रवार"},
	Saturday:              []string{"शनि", "शनिवार"},
	Sunday:                []string{"रवि", "रविवार"},
	AM:                    []string{"पूर्वाह्न"},
	PM:                    []string{"अपराह्न"},
	Year:                  []string{"वर्ष", "वर्षों", "साल"},
	Month:                 []string{"महीना", "महीने", "मास", "माह"},
	Week:                  []string{"सप्ताह"},
	Day:                   []string{"दिन", "दिवस"},
	Hour:                  []string{"घं", "घंटा", "घंटे"},
	Minute:                []string{"मि", "मिनट"},
	Second:                []string{"से", "सेकंड"},
	In:                    []string{"बाद", "में"},
	Ago:                   []string{"पहले", "पूर्व"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`आज`},
		`0 hour ago`:   {`यह घंटा`},
		`0 minute ago`: {`यह मिनट`},
		`0 month ago`:  {`इस माह`},
		`0 second ago`: {`अब`},
		`0 week ago`:   {`इस सप्ताह`},
		`0 year ago`:   {`इस वर्ष`},
		`1 day ago`:    {`कल`},
		`1 month ago`:  {`पिछला माह`},
		`1 week ago`:   {`पिछला सप्ताह`},
		`1 year ago`:   {`पिछला वर्ष`},
		`2 day ago`:    {`परसों`},
		`in 1 day`:     {`कल`},
		`in 1 month`:   {`अगला माह`},
		`in 1 week`:    {`अगला सप्ताह`},
		`in 1 year`:    {`अगला वर्ष`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) दिन पहले`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) घं पहले`),
			regexp.MustCompile(`(?i)(\d+) घंटे पहले`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) मि पहले`),
			regexp.MustCompile(`(?i)(\d+) मिनट पहले`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) माह पहले`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) से पहले`),
			regexp.MustCompile(`(?i)(\d+) सेकंड पहले`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) सप्ताह पहले`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) वर्ष पहले`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) दिन में`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) घं में`),
			regexp.MustCompile(`(?i)(\d+) घंटे में`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) मि में`),
			regexp.MustCompile(`(?i)(\d+) मिनट में`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) माह में`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) से में`),
			regexp.MustCompile(`(?i)(\d+) सेकंड में`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) सप्ताह में`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) वर्ष में`),
		},
	},
}
