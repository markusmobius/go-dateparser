// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kea_Locale = merge(nil, LocaleData{
	Name:      "kea",
	DateOrder: "DMY",
	Charset:   []rune(`-bcdefghijklnorstuvwzáó`),
	Translations: map[string][]string{
		"sigunda-fera": {"monday"},
		"kuarta-fera":  {"wednesday"},
		"kinta-fera":   {"thursday"},
		"sesta-fera":   {"friday"},
		"tersa-fera":   {"tuesday"},
		"dizenbru":     {"december"},
		"nuvenbru":     {"november"},
		"setenbru":     {"september"},
		"dumingu":      {"sunday"},
		"febreru":      {"february"},
		"sigundu":      {"second"},
		"agostu":       {"august"},
		"janeru":       {"january"},
		"minutu":       {"minute"},
		"otubru":       {"october"},
		"sabadu":       {"saturday"},
		"simana":       {"week"},
		"abril":        {"april"},
		"julhu":        {"july"},
		"junhu":        {"june"},
		"marsu":        {"march"},
		"maiu":         {"may"},
		"abr":          {"april"},
		"ago":          {"august"},
		"anu":          {"year"},
		"dia":          {"day"},
		"diz":          {"december"},
		"dum":          {"sunday"},
		"feb":          {"february"},
		"gmt":          {"gmt"},
		"jan":          {"january"},
		"jul":          {"july"},
		"jun":          {"june"},
		"kin":          {"thursday"},
		"kua":          {"wednesday"},
		"mai":          {"may"},
		"mar":          {"march"},
		"mes":          {"month"},
		"min":          {"minute"},
		"nuv":          {"november"},
		"ora":          {"hour"},
		"otu":          {"october"},
		"sab":          {"saturday"},
		"ses":          {"friday"},
		"set":          {"september"},
		"sig":          {"monday", "second"},
		"sim":          {"week"},
		"ter":          {"tuesday"},
		"utc":          {"utc"},
		"am":           {"am"},
		"pm":           {"pm"},
		" ":            {" "},
		"'":            {""},
		"+":            {"+"},
		",":            {""},
		"-":            {"-"},
		".":            {"."},
		"/":            {"/"},
		":":            {":"},
		";":            {""},
		"@":            {""},
		"[":            {""},
		"]":            {""},
		"h":            {"hour"},
		"m":            {"minute"},
		"s":            {"second"},
		"z":            {"z"},
		"|":            {""},
	},
	RelativeType: map[string]string{
		"prosimu simana": "in 1 week",
		"simana pasadu":  "1 week ago",
		"es simana li":   "0 week ago",
		"prosimu anu":    "in 1 year",
		"prosimu mes":    "in 1 month",
		"this minute":    "0 minute ago",
		"anu pasadu":     "1 year ago",
		"mes pasadu":     "1 month ago",
		"es anu li":      "0 year ago",
		"es mes li":      "0 month ago",
		"this hour":      "0 hour ago",
		"manha":          "in 1 day",
		"onti":           "1 day ago",
		"now":            "0 second ago",
		"oji":            "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) sigundu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) sigundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) simana`), "$1 week ago"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) simana`), "in $1 week"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) anu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) dia`), "$1 day ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) ora`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) sig`), "$1 second ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) sim`), "$1 week ago"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) anu`), "in $1 year"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) dia`), "in $1 day"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) ora`), "in $1 hour"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) sig`), "in $1 second"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) sim`), "in $1 week"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)a ten (\d+[.,]?\d*) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)di li (\d+[.,]?\d*) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(a ten \d+[.,]?\d* sigundu|di li \d+[.,]?\d* sigundu|a ten \d+[.,]?\d* minutu|a ten \d+[.,]?\d* simana|di li \d+[.,]?\d* minutu|di li \d+[.,]?\d* simana|a ten \d+[.,]?\d* anu|a ten \d+[.,]?\d* dia|a ten \d+[.,]?\d* mes|a ten \d+[.,]?\d* min|a ten \d+[.,]?\d* ora|a ten \d+[.,]?\d* sig|a ten \d+[.,]?\d* sim|di li \d+[.,]?\d* anu|di li \d+[.,]?\d* dia|di li \d+[.,]?\d* mes|di li \d+[.,]?\d* min|di li \d+[.,]?\d* ora|di li \d+[.,]?\d* sig|di li \d+[.,]?\d* sim|a ten \d+[.,]?\d* m|a ten \d+[.,]?\d* s|di li \d+[.,]?\d* m|di li \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(a ten \d+[.,]?\d* sigundu|di li \d+[.,]?\d* sigundu|a ten \d+[.,]?\d* minutu|a ten \d+[.,]?\d* simana|di li \d+[.,]?\d* minutu|di li \d+[.,]?\d* simana|a ten \d+[.,]?\d* anu|a ten \d+[.,]?\d* dia|a ten \d+[.,]?\d* mes|a ten \d+[.,]?\d* min|a ten \d+[.,]?\d* ora|a ten \d+[.,]?\d* sig|a ten \d+[.,]?\d* sim|di li \d+[.,]?\d* anu|di li \d+[.,]?\d* dia|di li \d+[.,]?\d* mes|di li \d+[.,]?\d* min|di li \d+[.,]?\d* ora|di li \d+[.,]?\d* sig|di li \d+[.,]?\d* sim|a ten \d+[.,]?\d* m|a ten \d+[.,]?\d* s|di li \d+[.,]?\d* m|di li \d+[.,]?\d* s)$`),
	KnownWords:      []string{"prosimu simana", "simana pasadu", "es simana li", "sigunda-fera", "kuarta-fera", "prosimu anu", "prosimu mes", "this minute", "anu pasadu", "kinta-fera", "mes pasadu", "sesta-fera", "tersa-fera", "es anu li", "es mes li", "this hour", "dizenbru", "nuvenbru", "setenbru", "dumingu", "febreru", "sigundu", "agostu", "janeru", "minutu", "otubru", "sabadu", "simana", "abril", "julhu", "junhu", "manha", "marsu", "maiu", "onti", "abr", "ago", "anu", "dia", "diz", "dum", "feb", "gmt", "jan", "jul", "jun", "kin", "kua", "mai", "mar", "mes", "min", "now", "nuv", "oji", "ora", "otu", "sab", "ses", "set", "sig", "sim", "ter", "utc", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "m", "s", "z", "|"},
})
