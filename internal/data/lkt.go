// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lkt_Locale = merge(nil, LocaleData{
	Name:      "lkt",
	DateOrder: "YMD",
	Charset:   []rune("+,-./;@[]ceghiklnorstuwyz|ŋ"),
	Translations: map[string]string{
		"wipazukha-waste wi": "june",
		"istawichayazaŋ wi":  "march",
		"chaŋwape-kasna wi":  "october",
		"thiyoheyuŋka wi":    "february",
		"owaphe oh'aŋkho":    "minute",
		"thahekapsuŋ wi":     "december",
		"chaŋphasapa wi":     "july",
		"chaŋwapetho wi":     "may",
		"wiothehika wi":      "january",
		"owaŋgyuzazapi":      "saturday",
		"chaŋwapegi wi":      "september",
		"aŋpetuzaptaŋ":       "friday",
		"aŋpetuwakhaŋ":       "sunday",
		"phezitho wi":        "april",
		"wasuthuŋ wi":        "august",
		"aŋpetuwaŋzi":        "monday",
		"waniyetu wi":        "november",
		"aŋpetunuŋpa":        "tuesday",
		"aŋpetuyamni":        "wednesday",
		"aŋpetutopa":         "thursday",
		"aŋpetu":             "day",
		"owaphe":             "hour",
		"omakha":             "year",
		"okpi":               "second",
		"gmt":                "gmt",
		"utc":                "utc",
		"oko":                "week",
		"am":                 "am",
		"wi":                 "month",
		"pm":                 "pm",
		"'":                  "",
		",":                  "",
		";":                  "",
		"@":                  "",
		"[":                  "",
		"]":                  "",
		"|":                  "",
		" ":                  " ",
		"+":                  "+",
		"-":                  "-",
		".":                  ".",
		"/":                  "/",
		":":                  ":",
		"z":                  "z",
	},
	RelativeType: map[string]string{
		"thokata omakha kiŋhaŋ": "in 1 year",
		"thokata oko kiŋhaŋ":    "in 1 week",
		"omakha k'uŋ hehaŋ":     "1 year ago",
		"thokata wi kiŋhaŋ":     "in 1 month",
		"hiŋhaŋni kiŋhaŋ":       "in 1 day",
		"oko k'uŋ hehaŋ":        "1 week ago",
		"le aŋpetu kiŋ":         "0 day ago",
		"le omakha kiŋ":         "0 year ago",
		"wi k'uŋ hehaŋ":         "1 month ago",
		"this minute":           "0 minute ago",
		"le oko kiŋ":            "0 week ago",
		"this hour":             "0 hour ago",
		"le wi kiŋ":             "0 month ago",
		"htalehaŋ":              "1 day ago",
		"now":                   "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)hekta oh'aŋkho (\d+) k'uŋ hehaŋ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)hekta wiyawapi (\d+) k'uŋ hehaŋ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)letaŋhaŋ oh'aŋkho (\d+) kiŋhaŋ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)letaŋhaŋ wiyawapi (\d+) kiŋhaŋ`), "in $1 month"},
		{regexp.MustCompile(`(?i)hekta owaphe (\d+) k'uŋ hehaŋ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)hekta omakha (\d+) k'uŋ hehaŋ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)letaŋhaŋ owaphe (\d+) kiŋhaŋ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)letaŋhaŋ omakha (\d+) kiŋhaŋ`), "in $1 year"},
		{regexp.MustCompile(`(?i)hekta (\d+)-chaŋ k'uŋ hehaŋ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)hekta okpi (\d+) k'uŋ hehaŋ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)hekta oko (\d+) k'uŋ hehaŋ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)letaŋhaŋ (\d+)-chaŋ kiŋhaŋ`), "in $1 day"},
		{regexp.MustCompile(`(?i)letaŋhaŋ okpi (\d+) kiŋhaŋ`), "in $1 second"},
		{regexp.MustCompile(`(?i)letaŋhaŋ oko (\d+) kiŋhaŋ`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(hekta oh'aŋkho \d+ k'uŋ hehaŋ|hekta wiyawapi \d+ k'uŋ hehaŋ|letaŋhaŋ oh'aŋkho \d+ kiŋhaŋ|letaŋhaŋ wiyawapi \d+ kiŋhaŋ|hekta omakha \d+ k'uŋ hehaŋ|hekta owaphe \d+ k'uŋ hehaŋ|letaŋhaŋ omakha \d+ kiŋhaŋ|letaŋhaŋ owaphe \d+ kiŋhaŋ|hekta \d+-chaŋ k'uŋ hehaŋ|hekta okpi \d+ k'uŋ hehaŋ|hekta oko \d+ k'uŋ hehaŋ|letaŋhaŋ \d+-chaŋ kiŋhaŋ|letaŋhaŋ okpi \d+ kiŋhaŋ|letaŋhaŋ oko \d+ kiŋhaŋ)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(hekta oh'aŋkho \d+ k'uŋ hehaŋ|hekta wiyawapi \d+ k'uŋ hehaŋ|letaŋhaŋ oh'aŋkho \d+ kiŋhaŋ|letaŋhaŋ wiyawapi \d+ kiŋhaŋ|hekta omakha \d+ k'uŋ hehaŋ|hekta owaphe \d+ k'uŋ hehaŋ|letaŋhaŋ omakha \d+ kiŋhaŋ|letaŋhaŋ owaphe \d+ kiŋhaŋ|hekta \d+-chaŋ k'uŋ hehaŋ|hekta okpi \d+ k'uŋ hehaŋ|hekta oko \d+ k'uŋ hehaŋ|letaŋhaŋ \d+-chaŋ kiŋhaŋ|letaŋhaŋ okpi \d+ kiŋhaŋ|letaŋhaŋ oko \d+ kiŋhaŋ)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(thokata omakha kiŋhaŋ|thokata oko kiŋhaŋ|wipazukha-waste wi|chaŋwape-kasna wi|istawichayazaŋ wi|omakha k'uŋ hehaŋ|thokata wi kiŋhaŋ|hiŋhaŋni kiŋhaŋ|owaphe oh'aŋkho|thiyoheyuŋka wi|chaŋphasapa wi|chaŋwapetho wi|oko k'uŋ hehaŋ|thahekapsuŋ wi|chaŋwapegi wi|le aŋpetu kiŋ|le omakha kiŋ|owaŋgyuzazapi|wi k'uŋ hehaŋ|wiothehika wi|aŋpetuwakhaŋ|aŋpetuzaptaŋ|aŋpetunuŋpa|aŋpetuwaŋzi|aŋpetuyamni|phezitho wi|this minute|waniyetu wi|wasuthuŋ wi|aŋpetutopa|le oko kiŋ|le wi kiŋ|this hour|htalehaŋ|aŋpetu|omakha|owaphe|okpi|gmt|now|oko|utc|\+|\.|\[|\]|\||am|pm|wi| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
