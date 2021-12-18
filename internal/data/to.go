// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var to_Locale = merge(nil, LocaleData{
	Name:      "to",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"mahina kuo'osi": "1 month ago",
		"mahina kaha'u":  "in 1 month",
		"uike kuo'osi":   "1 week ago",
		"ta'u kuo'osi":   "1 year ago",
		"'apongipongi":   "in 1 day",
		"tu'apulelulu":   "thursday",
		"this minute":    "0 minute ago",
		"uike kaha'u":    "in 1 week",
		"ta'u kaha'u":    "in 1 year",
		"hengihengi":     "am",
		"this hour":      "0 hour ago",
		"mahina ni":      "0 month ago",
		"taimi ni":       "0 second ago",
		"'epeleli":       "april",
		"'okatopa":       "october",
		"tokonaki":       "saturday",
		"sepitema":       "september",
		"pulelulu":       "wednesday",
		"'aho ni":        "0 day ago",
		"uike ni":        "0 week ago",
		"ta'u ni":        "0 year ago",
		"'aneafi":        "1 day ago",
		"'aokosi":        "august",
		"fepueli":        "february",
		"falaite":        "friday",
		"sanuali":        "january",
		"tisema":         "december",
		"siulai":         "july",
		"ma'asi":         "march",
		"miniti":         "minute",
		"monite":         "monday",
		"mahina":         "month",
		"novema":         "november",
		"efiafi":         "pm",
		"sekoni":         "second",
		"sapate":         "sunday",
		"tusite":         "tuesday",
		"'epe":           "april",
		"'aok":           "august",
		"'aho":           "day",
		"houa":           "hour",
		"sune":           "june",
		"ma'a":           "march",
		"'oka":           "october",
		"tu'a":           "thursday",
		"uike":           "week",
		"ta'u":           "year",
		"tis":            "december",
		"fep":            "february",
		"fal":            "friday",
		"gmt":            "gmt",
		"san":            "january",
		"siu":            "july",
		"sun":            "june",
		"mon":            "monday",
		"nov":            "november",
		"tok":            "saturday",
		"sep":            "september",
		"sap":            "sunday",
		"tus":            "tuesday",
		"utc":            "utc",
		"pul":            "wednesday",
		"am":             "am",
		"hh":             "am",
		"me":             "may",
		"ea":             "pm",
		"pm":             "pm",
		"'":              "",
		",":              "",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"|":              "",
		" ":              " ",
		"+":              "+",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		"z":              "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)miniti 'e (\d+) kuo'osi`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)mahina 'e (\d+) kuo'osi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)sekoni 'e (\d+) kuo'osi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)'aho 'e (\d+) kuo'osi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)houa 'e (\d+) kuo'osi`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)uike 'e (\d+) kuo'osi`), "$1 week ago"},
		{regexp.MustCompile(`(?i)ta'u 'e (\d+) kuo'osi`), "$1 year ago"},
		{regexp.MustCompile(`(?i)'i he miniti 'e (\d+)`), "in $1 minute"},
		{regexp.MustCompile(`(?i)'i he mahina 'e (\d+)`), "in $1 month"},
		{regexp.MustCompile(`(?i)'i he sekoni 'e (\d+)`), "in $1 second"},
		{regexp.MustCompile(`(?i)'i he 'aho 'e (\d+)`), "in $1 day"},
		{regexp.MustCompile(`(?i)'i he houa 'e (\d+)`), "in $1 hour"},
		{regexp.MustCompile(`(?i)'i he uike 'e (\d+)`), "in $1 week"},
		{regexp.MustCompile(`(?i)'i he ta'u 'e (\d+)`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(mahina 'e \d+ kuo'osi|miniti 'e \d+ kuo'osi|sekoni 'e \d+ kuo'osi|'aho 'e \d+ kuo'osi|'i he mahina 'e \d+|'i he miniti 'e \d+|'i he sekoni 'e \d+|houa 'e \d+ kuo'osi|ta'u 'e \d+ kuo'osi|uike 'e \d+ kuo'osi|'i he 'aho 'e \d+|'i he houa 'e \d+|'i he ta'u 'e \d+|'i he uike 'e \d+)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(mahina 'e \d+ kuo'osi|miniti 'e \d+ kuo'osi|sekoni 'e \d+ kuo'osi|'aho 'e \d+ kuo'osi|'i he mahina 'e \d+|'i he miniti 'e \d+|'i he sekoni 'e \d+|houa 'e \d+ kuo'osi|ta'u 'e \d+ kuo'osi|uike 'e \d+ kuo'osi|'i he 'aho 'e \d+|'i he houa 'e \d+|'i he ta'u 'e \d+|'i he uike 'e \d+)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(mahina kuo'osi|mahina kaha'u|'apongipongi|ta'u kuo'osi|tu'apulelulu|uike kuo'osi|ta'u kaha'u|this minute|uike kaha'u|hengihengi|mahina ni|this hour|'epeleli|'okatopa|pulelulu|sepitema|taimi ni|tokonaki|'aho ni|'aneafi|'aokosi|falaite|fepueli|sanuali|ta'u ni|uike ni|efiafi|ma'asi|mahina|miniti|monite|novema|sapate|sekoni|siulai|tisema|tusite|'aho|'aok|'epe|'oka|houa|ma'a|sune|ta'u|tu'a|uike|fal|fep|gmt|mon|nov|pul|san|sap|sep|siu|sun|tis|tok|tus|utc|\+|\.|\[|\]|\||am|ea|hh|me|pm| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
