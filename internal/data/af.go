// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var af_Locale = merge(nil, LocaleData{
	Name:      "af",
	DateOrder: "YMD",
	Charset:   []rune(`bcdefghijklnorstuvwyzô`),
	Translations: map[string][]string{
		"donderdag": {"thursday"},
		"februarie": {"february"},
		"september": {"september"},
		"augustus":  {"august"},
		"desember":  {"december"},
		"januarie":  {"january"},
		"november":  {"november"},
		"saterdag":  {"saturday"},
		"woensdag":  {"wednesday"},
		"dinsdag":   {"tuesday"},
		"maandag":   {"monday"},
		"oktober":   {"october"},
		"sekonde":   {"second"},
		"minuut":    {"minute"},
		"sondag":    {"sunday"},
		"vrydag":    {"friday"},
		"april":     {"april"},
		"julie":     {"july"},
		"junie":     {"june"},
		"maand":     {"month"},
		"maart":     {"march"},
		"jaar":      {"year"},
		"week":      {"week"},
		"apr":       {"april"},
		"aug":       {"august"},
		"dag":       {"day"},
		"des":       {"december"},
		"feb":       {"february"},
		"gmt":       {"gmt"},
		"jan":       {"january"},
		"jul":       {"july"},
		"jun":       {"june"},
		"mei":       {"may"},
		"min":       {"minute"},
		"mrt":       {"march"},
		"nov":       {"november"},
		"okt":       {"october"},
		"sek":       {"second"},
		"sep":       {"september"},
		"utc":       {"utc"},
		"uur":       {"hour"},
		"am":        {"am"},
		"di":        {"tuesday"},
		"do":        {"thursday"},
		"ma":        {"monday"},
		"md":        {"month"},
		"nm":        {"pm"},
		"pm":        {"pm"},
		"sa":        {"saturday"},
		"so":        {"sunday"},
		"vm":        {"am"},
		"vr":        {"friday"},
		"wk":        {"week"},
		"wo":        {"wednesday"},
		" ":         {" "},
		"'":         {""},
		"+":         {"+"},
		",":         {""},
		"-":         {"-"},
		".":         {"."},
		"/":         {"/"},
		":":         {":"},
		";":         {""},
		"@":         {""},
		"[":         {""},
		"]":         {""},
		"d":         {"day"},
		"j":         {"year"},
		"m":         {"minute"},
		"s":         {"second"},
		"u":         {"hour"},
		"z":         {"z"},
		"|":         {""},
	},
	RelativeType: map[string]string{
		"hierdie minuut": "0 minute ago",
		"volgende maand": "in 1 month",
		"verlede maand":  "1 month ago",
		"volgende jaar":  "in 1 year",
		"volgende week":  "in 1 week",
		"hierdie jaar":   "0 year ago",
		"vandeesmaand":   "0 month ago",
		"verlede jaar":   "1 year ago",
		"verlede week":   "1 week ago",
		"hierdie uur":    "0 hour ago",
		"vandeesweek":    "0 week ago",
		"gister":         "1 day ago",
		"vandag":         "0 day ago",
		"more":           "in 1 day",
		"nou":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekondes gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekonde gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) maande gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minute gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuut gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) maand gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) sekondes`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) jaar gelede`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) week gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) weke gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) sekonde`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dae gelede`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dag gelede`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min gelede`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sek gelede`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) uur gelede`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) minuut`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) md gelede`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) w gelede`), "$1 week ago"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) jaar`), "in $1 year"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) week`), "in $1 week"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) weke`), "in $1 week"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) dae`), "in $1 day"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) uur`), "in $1 hour"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) md`), "in $1 month"},
		{regexp.MustCompile(`(?i)oor (\d+[.,]?\d*) w`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* sekondes gelede|\d+[.,]?\d* sekonde gelede|\d+[.,]?\d* maande gelede|\d+[.,]?\d* minute gelede|\d+[.,]?\d* minuut gelede|\d+[.,]?\d* maand gelede|oor \d+[.,]?\d* sekondes|\d+[.,]?\d* jaar gelede|\d+[.,]?\d* week gelede|\d+[.,]?\d* weke gelede|oor \d+[.,]?\d* sekonde|\d+[.,]?\d* dae gelede|\d+[.,]?\d* dag gelede|\d+[.,]?\d* min gelede|\d+[.,]?\d* sek gelede|\d+[.,]?\d* uur gelede|oor \d+[.,]?\d* minuut|\d+[.,]?\d* md gelede|\d+[.,]?\d* w gelede|oor \d+[.,]?\d* jaar|oor \d+[.,]?\d* week|oor \d+[.,]?\d* weke|oor \d+[.,]?\d* dae|oor \d+[.,]?\d* dag|oor \d+[.,]?\d* min|oor \d+[.,]?\d* sek|oor \d+[.,]?\d* uur|oor \d+[.,]?\d* md|oor \d+[.,]?\d* w)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* sekondes gelede|\d+[.,]?\d* sekonde gelede|\d+[.,]?\d* maande gelede|\d+[.,]?\d* minute gelede|\d+[.,]?\d* minuut gelede|\d+[.,]?\d* maand gelede|oor \d+[.,]?\d* sekondes|\d+[.,]?\d* jaar gelede|\d+[.,]?\d* week gelede|\d+[.,]?\d* weke gelede|oor \d+[.,]?\d* sekonde|\d+[.,]?\d* dae gelede|\d+[.,]?\d* dag gelede|\d+[.,]?\d* min gelede|\d+[.,]?\d* sek gelede|\d+[.,]?\d* uur gelede|oor \d+[.,]?\d* minuut|\d+[.,]?\d* md gelede|\d+[.,]?\d* w gelede|oor \d+[.,]?\d* jaar|oor \d+[.,]?\d* week|oor \d+[.,]?\d* weke|oor \d+[.,]?\d* dae|oor \d+[.,]?\d* dag|oor \d+[.,]?\d* min|oor \d+[.,]?\d* sek|oor \d+[.,]?\d* uur|oor \d+[.,]?\d* md|oor \d+[.,]?\d* w)$`),
	KnownWords:      []string{"hierdie minuut", "volgende maand", "verlede maand", "volgende jaar", "volgende week", "hierdie jaar", "vandeesmaand", "verlede jaar", "verlede week", "hierdie uur", "vandeesweek", "donderdag", "februarie", "september", "augustus", "desember", "januarie", "november", "saterdag", "woensdag", "dinsdag", "maandag", "oktober", "sekonde", "gister", "minuut", "sondag", "vandag", "vrydag", "april", "julie", "junie", "maand", "maart", "jaar", "more", "week", "apr", "aug", "dag", "des", "feb", "gmt", "jan", "jul", "jun", "mei", "min", "mrt", "nou", "nov", "okt", "sek", "sep", "utc", "uur", "am", "di", "do", "ma", "md", "nm", "pm", "sa", "so", "vm", "vr", "wk", "wo", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "j", "m", "s", "u", "z", "|"},
})

var af_NA_Locale = merge(&af_Locale, LocaleData{
	Name:      "af-NA",
	DateOrder: "YMD",
})
