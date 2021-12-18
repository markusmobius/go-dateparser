// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var az_Latn_Locale = merge(nil, LocaleData{
	Name:      "az-Latn",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"cərsənbə axsamı": "tuesday",
		"bazar ertəsi":    "monday",
		"kecən həftə":     "1 week ago",
		"gələn həftə":     "in 1 week",
		"cumə axsamı":     "thursday",
		"bu dəqiqə":       "0 minute ago",
		"bu həftə":        "0 week ago",
		"kecən ay":        "1 month ago",
		"kecən il":        "1 year ago",
		"gələn ay":        "in 1 month",
		"gələn il":        "in 1 year",
		"sentyabr":        "september",
		"cərsənbə":        "wednesday",
		"bu saat":         "0 hour ago",
		"oktyabr":         "october",
		"bu gun":          "0 day ago",
		"avqust":          "august",
		"dekabr":          "december",
		"fevral":          "february",
		"yanvar":          "january",
		"dəqiqə":          "minute",
		"noyabr":          "november",
		"saniyə":          "second",
		"bu ay":           "0 month ago",
		"bu il":           "0 year ago",
		"dunən":           "1 day ago",
		"aprel":           "april",
		"sabah":           "in 1 day",
		"sənbə":           "saturday",
		"bazar":           "sunday",
		"həftə":           "week",
		"indi":            "0 second ago",
		"cumə":            "friday",
		"saat":            "hour",
		"iyul":            "july",
		"iyun":            "june",
		"mart":            "march",
		"apr":             "april",
		"avq":             "august",
		"gun":             "day",
		"dek":             "december",
		"fev":             "february",
		"gmt":             "gmt",
		"yan":             "january",
		"iyl":             "july",
		"iyn":             "june",
		"mar":             "march",
		"may":             "may",
		"dəq":             "minute",
		"noy":             "november",
		"okt":             "october",
		"san":             "second",
		"sen":             "september",
		"utc":             "utc",
		"am":              "am",
		"be":              "monday",
		"ay":              "month",
		"pm":              "pm",
		"ca":              "thursday",
		"il":              "year",
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
		"c":               "friday",
		"s":               "saturday",
		"b":               "sunday",
		"z":               "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) dəqiqə ərzində`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) saniyə ərzində`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) həftə ərzində`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) saat ərzində`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) dəqiqə oncə`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) saniyə oncə`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) gun ərzində`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) həftə oncə`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ay ərzində`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) il ərzində`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) saat oncə`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) gun oncə`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ay oncə`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) il oncə`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(\d+ dəqiqə ərzində|\d+ saniyə ərzində|\d+ həftə ərzində|\d+ saat ərzində|\d+ dəqiqə oncə|\d+ gun ərzində|\d+ saniyə oncə|\d+ ay ərzində|\d+ həftə oncə|\d+ il ərzində|\d+ saat oncə|\d+ gun oncə|\d+ ay oncə|\d+ il oncə)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ dəqiqə ərzində|\d+ saniyə ərzində|\d+ həftə ərzində|\d+ saat ərzində|\d+ dəqiqə oncə|\d+ gun ərzində|\d+ saniyə oncə|\d+ ay ərzində|\d+ həftə oncə|\d+ il ərzində|\d+ saat oncə|\d+ gun oncə|\d+ ay oncə|\d+ il oncə)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(cərsənbə axsamı|bazar ertəsi|cumə axsamı|gələn həftə|kecən həftə|bu dəqiqə|bu həftə|cərsənbə|gələn ay|gələn il|kecən ay|kecən il|sentyabr|bu saat|oktyabr|avqust|bu gun|dekabr|dəqiqə|fevral|noyabr|saniyə|yanvar|aprel|bazar|bu ay|bu il|dunən|həftə|sabah|sənbə|cumə|indi|iyul|iyun|mart|saat|apr|avq|dek|dəq|fev|gmt|gun|iyl|iyn|mar|may|noy|okt|san|sen|utc|yan|\+|\.|\[|\]|\||am|ay|be|ca|il|pm| |'|,|-|/|:|;|@|b|c|s|z)((?:\z|\W|_|\d).*)$`),
})
