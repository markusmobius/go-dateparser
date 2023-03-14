// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kl_Locale = merge(nil, LocaleData{
	Name:      "kl",
	DateOrder: "YMD",
	Charset:   []rune(`bcdefghijklnorstuvwxyz`),
	Translations: map[string][]string{
		"pingasunngorneq": {"wednesday"},
		"tallimanngorneq": {"friday"},
		"arfininngorneq":  {"saturday"},
		"ataasinngorneq":  {"monday"},
		"sisamanngorneq":  {"thursday"},
		"marlunngorneq":   {"tuesday"},
		"septemberi":      {"september"},
		"augustusi":       {"august"},
		"decemberi":       {"december"},
		"novemberi":       {"november"},
		"februari":        {"february"},
		"oktoberi":        {"october"},
		"januari":         {"january"},
		"aprili":          {"april"},
		"martsi":          {"march"},
		"minute":          {"minute"},
		"sabaat":          {"sunday"},
		"second":          {"second"},
		"month":           {"month"},
		"hour":            {"hour"},
		"juli":            {"july"},
		"juni":            {"june"},
		"maji":            {"may"},
		"week":            {"week"},
		"year":            {"year"},
		"apr":             {"april"},
		"arf":             {"saturday"},
		"ata":             {"monday"},
		"aug":             {"august"},
		"day":             {"day"},
		"dec":             {"december"},
		"feb":             {"february"},
		"gmt":             {"gmt"},
		"jan":             {"january"},
		"jul":             {"july"},
		"jun":             {"june"},
		"maj":             {"may"},
		"mar":             {"march", "tuesday"},
		"nov":             {"november"},
		"okt":             {"october"},
		"pin":             {"wednesday"},
		"sab":             {"sunday"},
		"sep":             {"september"},
		"sis":             {"thursday"},
		"tal":             {"friday"},
		"utc":             {"utc"},
		"am":              {"am"},
		"pm":              {"pm"},
		" ":               {" "},
		"'":               {""},
		"+":               {"+"},
		",":               {""},
		"-":               {"-"},
		".":               {"."},
		"/":               {"/"},
		":":               {":"},
		";":               {""},
		"@":               {""},
		"[":               {""},
		"]":               {""},
		"z":               {"z"},
		"|":               {""},
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this month":  "0 month ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"yesterday":   "1 day ago",
		"tomorrow":    "in 1 day",
		"today":       "0 day ago",
		"now":         "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) nalunaaquttap-akunnera siden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sapaatip-akunnera siden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) nalunaaquttap-akunnera`), "in $1 hour"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) ulloq unnuarlu siden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sapaatip-akunnera`), "in $1 week"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minutsi siden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) qaammat siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekundi siden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) ulloq unnuarlu`), "in $1 day"},
		{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) ukioq siden`), "$1 year ago"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minutsi`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) qaammat`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekundi`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) ukioq`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+[.,]?\d* nalunaaquttap-akunnera siden|for \d+[.,]?\d* sapaatip-akunnera siden|om \d+[.,]?\d* nalunaaquttap-akunnera|for \d+[.,]?\d* ulloq unnuarlu siden|om \d+[.,]?\d* sapaatip-akunnera|for \d+[.,]?\d* minutsi siden|for \d+[.,]?\d* qaammat siden|for \d+[.,]?\d* sekundi siden|om \d+[.,]?\d* ulloq unnuarlu|for \d+[.,]?\d* ukioq siden|om \d+[.,]?\d* minutsi|om \d+[.,]?\d* qaammat|om \d+[.,]?\d* sekundi|om \d+[.,]?\d* ukioq)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(for \d+[.,]?\d* nalunaaquttap-akunnera siden|for \d+[.,]?\d* sapaatip-akunnera siden|om \d+[.,]?\d* nalunaaquttap-akunnera|for \d+[.,]?\d* ulloq unnuarlu siden|om \d+[.,]?\d* sapaatip-akunnera|for \d+[.,]?\d* minutsi siden|for \d+[.,]?\d* qaammat siden|for \d+[.,]?\d* sekundi siden|om \d+[.,]?\d* ulloq unnuarlu|for \d+[.,]?\d* ukioq siden|om \d+[.,]?\d* minutsi|om \d+[.,]?\d* qaammat|om \d+[.,]?\d* sekundi|om \d+[.,]?\d* ukioq)$`),
	KnownWords:      []string{"pingasunngorneq", "tallimanngorneq", "arfininngorneq", "ataasinngorneq", "sisamanngorneq", "marlunngorneq", "this minute", "last month", "next month", "septemberi", "this month", "augustusi", "decemberi", "last week", "last year", "next week", "next year", "novemberi", "this hour", "this week", "this year", "yesterday", "februari", "oktoberi", "tomorrow", "januari", "aprili", "martsi", "minute", "sabaat", "second", "month", "today", "hour", "juli", "juni", "maji", "week", "year", "apr", "arf", "ata", "aug", "day", "dec", "feb", "gmt", "jan", "jul", "jun", "maj", "mar", "nov", "now", "okt", "pin", "sab", "sep", "sis", "tal", "utc", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
