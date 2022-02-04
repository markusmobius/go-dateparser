// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var af_Locale = merge(nil, LocaleData{
	Name:      "af",
	DateOrder: "YMD",
	Charset:   []rune(`bcdefghijklnorstuvwyzô`),
	Translations: map[string]string{
		"februarie": "february",
		"september": "september",
		"donderdag": "thursday",
		"augustus":  "august",
		"desember":  "december",
		"januarie":  "january",
		"november":  "november",
		"saterdag":  "saturday",
		"woensdag":  "wednesday",
		"maandag":   "monday",
		"oktober":   "october",
		"sekonde":   "second",
		"dinsdag":   "tuesday",
		"vrydag":    "friday",
		"minuut":    "minute",
		"sondag":    "sunday",
		"april":     "april",
		"julie":     "july",
		"junie":     "june",
		"maart":     "march",
		"maand":     "month",
		"week":      "week",
		"jaar":      "year",
		"apr":       "april",
		"aug":       "august",
		"dag":       "day",
		"des":       "december",
		"feb":       "february",
		"gmt":       "gmt",
		"uur":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mrt":       "march",
		"mei":       "may",
		"min":       "minute",
		"nov":       "november",
		"okt":       "october",
		"sek":       "second",
		"sep":       "september",
		"utc":       "utc",
		"am":        "am",
		"vm":        "am",
		"vr":        "friday",
		"ma":        "monday",
		"md":        "month",
		"nm":        "pm",
		"pm":        "pm",
		"sa":        "saturday",
		"so":        "sunday",
		"do":        "thursday",
		"di":        "tuesday",
		"wo":        "wednesday",
		"wk":        "week",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"d":         "day",
		"u":         "hour",
		"m":         "minute",
		"s":         "second",
		"j":         "year",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"hierdie minuut": "0 minute ago",
		"volgende maand": "in 1 month",
		"verlede maand":  "1 month ago",
		"volgende week":  "in 1 week",
		"volgende jaar":  "in 1 year",
		"vandeesmaand":   "0 month ago",
		"hierdie jaar":   "0 year ago",
		"verlede week":   "1 week ago",
		"verlede jaar":   "1 year ago",
		"hierdie uur":    "0 hour ago",
		"vandeesweek":    "0 week ago",
		"vandag":         "0 day ago",
		"gister":         "1 day ago",
		"more":           "in 1 day",
		"nou":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) sekondes gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) sekonde gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minute gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) minuut gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) maande gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) maand gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)oor (\d+) sekondes`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) week gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) weke gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) jaar gelede`), "$1 year ago"},
		{regexp.MustCompile(`(?i)oor (\d+) sekonde`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) dae gelede`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) dag gelede`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) uur gelede`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) min gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sek gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)oor (\d+) minuut`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) md gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) w gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)oor (\d+) week`), "in $1 week"},
		{regexp.MustCompile(`(?i)oor (\d+) weke`), "in $1 week"},
		{regexp.MustCompile(`(?i)oor (\d+) jaar`), "in $1 year"},
		{regexp.MustCompile(`(?i)oor (\d+) dae`), "in $1 day"},
		{regexp.MustCompile(`(?i)oor (\d+) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)oor (\d+) uur`), "in $1 hour"},
		{regexp.MustCompile(`(?i)oor (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)oor (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)oor (\d+) md`), "in $1 month"},
		{regexp.MustCompile(`(?i)oor (\d+) w`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ sekondes gelede|\d+ sekonde gelede|\d+ maande gelede|\d+ minute gelede|\d+ minuut gelede|\d+ maand gelede|oor \d+ sekondes|\d+ jaar gelede|\d+ week gelede|\d+ weke gelede|oor \d+ sekonde|\d+ dae gelede|\d+ dag gelede|\d+ min gelede|\d+ sek gelede|\d+ uur gelede|oor \d+ minuut|\d+ md gelede|\d+ w gelede|oor \d+ jaar|oor \d+ week|oor \d+ weke|oor \d+ dae|oor \d+ dag|oor \d+ min|oor \d+ sek|oor \d+ uur|oor \d+ md|oor \d+ w)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ sekondes gelede|\d+ sekonde gelede|\d+ maande gelede|\d+ minute gelede|\d+ minuut gelede|\d+ maand gelede|oor \d+ sekondes|\d+ jaar gelede|\d+ week gelede|\d+ weke gelede|oor \d+ sekonde|\d+ dae gelede|\d+ dag gelede|\d+ min gelede|\d+ sek gelede|\d+ uur gelede|oor \d+ minuut|\d+ md gelede|\d+ w gelede|oor \d+ jaar|oor \d+ week|oor \d+ weke|oor \d+ dae|oor \d+ dag|oor \d+ min|oor \d+ sek|oor \d+ uur|oor \d+ md|oor \d+ w)$`),
	KnownWords:      []string{"hierdie minuut", "volgende maand", "verlede maand", "volgende jaar", "volgende week", "hierdie jaar", "vandeesmaand", "verlede jaar", "verlede week", "hierdie uur", "vandeesweek", "donderdag", "februarie", "september", "augustus", "desember", "januarie", "november", "saterdag", "woensdag", "dinsdag", "maandag", "oktober", "sekonde", "gister", "minuut", "sondag", "vandag", "vrydag", "april", "julie", "junie", "maand", "maart", "jaar", "more", "week", "apr", "aug", "dag", "des", "feb", "gmt", "jan", "jul", "jun", "mei", "min", "mrt", "nou", "nov", "okt", "sek", "sep", "utc", "uur", "am", "di", "do", "ma", "md", "nm", "pm", "sa", "so", "vm", "vr", "wk", "wo", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "j", "m", "s", "u", "z", "|"},
})

var af_NA_Locale = merge(&af_Locale, LocaleData{
	Name:      "af-NA",
	DateOrder: "YMD",
})
