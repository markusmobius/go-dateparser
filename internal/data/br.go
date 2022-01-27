// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var br_Locale = merge(nil, LocaleData{
	Name:      "br",
	DateOrder: "YMD",
	Charset:   []rune("+,-./;@[]bcdeghiklnorstuvwyz|"),
	Translations: map[string]string{
		"c'hwevrer": "february",
		"mezheven":  "june",
		"gwengolo":  "september",
		"merc'her":  "wednesday",
		"gwener":    "friday",
		"genver":    "january",
		"gouere":    "july",
		"meurzh":    "march",
		"sadorn":    "saturday",
		"eilenn":    "second",
		"sizhun":    "week",
		"ebrel":     "april",
		"kerzu":     "december",
		"c'hwe":     "february",
		"munut":     "minute",
		"bloaz":     "year",
		"eost":      "august",
		"deiz":      "day",
		"goue":      "july",
		"mezh":      "june",
		"meur":      "march",
		"here":      "october",
		"gwen":      "september",
		"yaou":      "thursday",
		"ebr":       "april",
		"ker":       "december",
		"kzu":       "december",
		"gwe":       "friday",
		"gmt":       "gmt",
		"eur":       "hour",
		"gen":       "january",
		"mae":       "may",
		"min":       "minute",
		"lun":       "monday",
		"miz":       "month",
		"sad":       "saturday",
		"sul":       "sunday",
		"meu":       "tuesday",
		"utc":       "utc",
		"mer":       "wednesday",
		"am":        "am",
		"du":        "november",
		"gm":        "pm",
		"pm":        "pm",
		"bl":        "year",
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
		"e":         "hour",
		"s":         "second",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"ar sizhun diaraok": "1 week ago",
		"ar sizhun a zeu":   "in 1 week",
		"ar miz diaraok":    "1 month ago",
		"ar bloaz a zeu":    "in 1 year",
		"ar sizhun-man":     "0 week ago",
		"ar miz a zeu":      "in 1 month",
		"this minute":       "0 minute ago",
		"ar bl a zeu":       "in 1 year",
		"ar miz-man":        "0 month ago",
		"warc'hoazh":        "in 1 day",
		"this hour":         "0 hour ago",
		"hevlene":           "0 year ago",
		"warlene":           "1 year ago",
		"breman":            "0 second ago",
		"hiziv":             "0 day ago",
		"dec'h":             "1 day ago",
		"brem":              "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)a-benn (\d+) eilenn`), "in $1 second"},
		{regexp.MustCompile(`(?i)a-benn (\d+) sizhun`), "in $1 week"},
		{regexp.MustCompile(`(?i)a-benn (\d+) munut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)a-benn (\d+) bloaz`), "in $1 year"},
		{regexp.MustCompile(`(?i)a-benn (\d+) vloaz`), "in $1 year"},
		{regexp.MustCompile(`(?i)a-benn (\d+) deiz`), "in $1 day"},
		{regexp.MustCompile(`(?i)a-benn (\d+) eur`), "in $1 hour"},
		{regexp.MustCompile(`(?i)a-benn (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)a-benn (\d+) miz`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) eilenn zo`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) sizhun zo`), "$1 week ago"},
		{regexp.MustCompile(`(?i)a-benn (\d+) bl`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) munut zo`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bloaz zo`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) vloaz zo`), "$1 year ago"},
		{regexp.MustCompile(`(?i)a-benn (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)a-benn (\d+) e`), "in $1 hour"},
		{regexp.MustCompile(`(?i)a-benn (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) deiz zo`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) eur zo`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) min zo`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) miz zo`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) bl zo`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) d zo`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) e zo`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) s zo`), "$1 second ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(a-benn \d+ eilenn|a-benn \d+ sizhun|a-benn \d+ bloaz|a-benn \d+ munut|a-benn \d+ vloaz|a-benn \d+ deiz|a-benn \d+ eur|a-benn \d+ min|a-benn \d+ miz|\d+ eilenn zo|\d+ sizhun zo|a-benn \d+ bl|\d+ bloaz zo|\d+ munut zo|\d+ vloaz zo|a-benn \d+ d|a-benn \d+ e|a-benn \d+ s|\d+ deiz zo|\d+ eur zo|\d+ min zo|\d+ miz zo|\d+ bl zo|\d+ d zo|\d+ e zo|\d+ s zo)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(a-benn \d+ eilenn|a-benn \d+ sizhun|a-benn \d+ bloaz|a-benn \d+ munut|a-benn \d+ vloaz|a-benn \d+ deiz|a-benn \d+ eur|a-benn \d+ min|a-benn \d+ miz|\d+ eilenn zo|\d+ sizhun zo|a-benn \d+ bl|\d+ bloaz zo|\d+ munut zo|\d+ vloaz zo|a-benn \d+ d|a-benn \d+ e|a-benn \d+ s|\d+ deiz zo|\d+ eur zo|\d+ min zo|\d+ miz zo|\d+ bl zo|\d+ d zo|\d+ e zo|\d+ s zo)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ar sizhun diaraok|ar sizhun a zeu|ar bloaz a zeu|ar miz diaraok|ar sizhun-man|ar miz a zeu|ar bl a zeu|this minute|ar miz-man|warc'hoazh|c'hwevrer|this hour|gwengolo|merc'her|mezheven|hevlene|warlene|breman|eilenn|genver|gouere|gwener|meurzh|sadorn|sizhun|bloaz|c'hwe|dec'h|ebrel|hiziv|kerzu|munut|brem|deiz|eost|goue|gwen|here|meur|mezh|yaou|ebr|eur|gen|gmt|gwe|ker|kzu|lun|mae|mer|meu|min|miz|sad|sul|utc|\+|\.|\[|\]|\||am|bl|du|gm|pm| |'|,|-|/|:|;|@|d|e|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
