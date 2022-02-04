// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kl_Locale = merge(nil, LocaleData{
	Name:      "kl",
	DateOrder: "YMD",
	Charset:   []rune(`bcdefghijklnorstuvwxyz`),
	Translations: map[string]string{
		"tallimanngorneq": "friday",
		"pingasunngorneq": "wednesday",
		"ataasinngorneq":  "monday",
		"arfininngorneq":  "saturday",
		"sisamanngorneq":  "thursday",
		"marlunngorneq":   "tuesday",
		"septemberi":      "september",
		"augustusi":       "august",
		"decemberi":       "december",
		"novemberi":       "november",
		"februari":        "february",
		"oktoberi":        "october",
		"januari":         "january",
		"aprili":          "april",
		"martsi":          "march",
		"minute":          "minute",
		"second":          "second",
		"sabaat":          "sunday",
		"month":           "month",
		"hour":            "hour",
		"juli":            "july",
		"juni":            "june",
		"maji":            "may",
		"week":            "week",
		"year":            "year",
		"apr":             "april",
		"aug":             "august",
		"day":             "day",
		"dec":             "december",
		"feb":             "february",
		"tal":             "friday",
		"gmt":             "gmt",
		"jan":             "january",
		"jul":             "july",
		"jun":             "june",
		"mar":             "march",
		"maj":             "may",
		"ata":             "monday",
		"nov":             "november",
		"okt":             "october",
		"arf":             "saturday",
		"sep":             "september",
		"sab":             "sunday",
		"sis":             "thursday",
		"utc":             "utc",
		"pin":             "wednesday",
		"am":              "am",
		"pm":              "pm",
		"'":               "",
		",":               "",
		";":               "",
		"@":               "",
		"[":               "",
		"]":               "",
		"|":               "",
		" ":               " ",
		"+":               "+",
		"-":               "-",
		".":               ".",
		"/":               "/",
		":":               ":",
		"z":               "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"yesterday":   "1 day ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"tomorrow":    "in 1 day",
		"today":       "0 day ago",
		"now":         "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)for (\d+) nalunaaquttap-akunnera siden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) sapaatip-akunnera siden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)om (\d+) nalunaaquttap-akunnera`), "in $1 hour"},
		{regexp.MustCompile(`(?i)for (\d+) ulloq unnuarlu siden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)om (\d+) sapaatip-akunnera`), "in $1 week"},
		{regexp.MustCompile(`(?i)for (\d+) minutsi siden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) qaammat siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) sekundi siden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)om (\d+) ulloq unnuarlu`), "in $1 day"},
		{regexp.MustCompile(`(?i)for (\d+) ukioq siden`), "$1 year ago"},
		{regexp.MustCompile(`(?i)om (\d+) minutsi`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) qaammat`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) sekundi`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) ukioq`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+ nalunaaquttap-akunnera siden|for \d+ sapaatip-akunnera siden|om \d+ nalunaaquttap-akunnera|for \d+ ulloq unnuarlu siden|om \d+ sapaatip-akunnera|for \d+ minutsi siden|for \d+ qaammat siden|for \d+ sekundi siden|om \d+ ulloq unnuarlu|for \d+ ukioq siden|om \d+ minutsi|om \d+ qaammat|om \d+ sekundi|om \d+ ukioq)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(for \d+ nalunaaquttap-akunnera siden|for \d+ sapaatip-akunnera siden|om \d+ nalunaaquttap-akunnera|for \d+ ulloq unnuarlu siden|om \d+ sapaatip-akunnera|for \d+ minutsi siden|for \d+ qaammat siden|for \d+ sekundi siden|om \d+ ulloq unnuarlu|for \d+ ukioq siden|om \d+ minutsi|om \d+ qaammat|om \d+ sekundi|om \d+ ukioq)$`),
	KnownWords:      []string{"pingasunngorneq", "tallimanngorneq", "arfininngorneq", "ataasinngorneq", "sisamanngorneq", "marlunngorneq", "this minute", "last month", "next month", "septemberi", "this month", "augustusi", "decemberi", "last week", "last year", "next week", "next year", "novemberi", "this hour", "this week", "this year", "yesterday", "februari", "oktoberi", "tomorrow", "januari", "aprili", "martsi", "minute", "sabaat", "second", "month", "today", "hour", "juli", "juni", "maji", "week", "year", "apr", "arf", "ata", "aug", "day", "dec", "feb", "gmt", "jan", "jul", "jun", "maj", "mar", "nov", "now", "okt", "pin", "sab", "sep", "sis", "tal", "utc", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
