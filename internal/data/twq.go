// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var twq_Locale = merge(nil, LocaleData{
	Name:      "twq",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"zaarikay b":  "pm",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"deesanbur":   "december",
		"feewiriye":   "february",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"noowanbur":   "november",
		"sektanbur":   "september",
		"alhamiisa":   "thursday",
		"subbaahi":    "am",
		"oktoobur":    "october",
		"atalaata":    "tuesday",
		"zanwiye":     "january",
		"awiril":      "april",
		"alzuma":      "friday",
		"miniti":      "minute",
		"atinni":      "monday",
		"asibti":      "saturday",
		"alhadi":      "sunday",
		"alarba":      "wednesday",
		"zaari":       "day",
		"guuru":       "hour",
		"zuyye":       "july",
		"zuweŋ":       "june",
		"marsi":       "march",
		"handu":       "month",
		"jiiri":       "year",
		"suba":        "in 1 day",
		"miti":        "second",
		"hebu":        "week",
		"hoo":         "0 day ago",
		"now":         "0 second ago",
		"awi":         "april",
		"dee":         "december",
		"fee":         "february",
		"alz":         "friday",
		"gmt":         "gmt",
		"zan":         "january",
		"zuy":         "july",
		"zuw":         "june",
		"mar":         "march",
		"ati":         "monday",
		"noo":         "november",
		"okt":         "october",
		"asi":         "saturday",
		"sek":         "september",
		"alh":         "sunday",
		"alm":         "thursday",
		"ata":         "tuesday",
		"utc":         "utc",
		"ala":         "wednesday",
		"bi":          "1 day ago",
		"am":          "am",
		"ut":          "august",
		"me":          "may",
		"pm":          "pm",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"|":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"z":           "z",
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)()(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^()$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(this minute|last month|next month|this month|zaarikay b|alhamiisa|deesanbur|feewiriye|last week|last year|next week|next year|noowanbur|sektanbur|this hour|this week|this year|atalaata|oktoobur|subbaahi|zanwiye|alarba|alhadi|alzuma|asibti|atinni|awiril|miniti|guuru|handu|jiiri|marsi|zaari|zuweŋ|zuyye|hebu|miti|suba|ala|alh|alm|alz|asi|ata|ati|awi|dee|fee|gmt|hoo|mar|noo|now|okt|sek|utc|zan|zuw|zuy|\+|\.|\[|\]|\||am|bi|me|pm|ut| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
