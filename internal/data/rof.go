// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var rof_Locale = merge(nil, LocaleData{
	Name:      "rof",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]bcdefghijklnorstuwxyz|"),
	Translations: map[string]string{
		"mweri wa ikumi na mbili": "december",
		"mweri wa ikumi na moja":  "november",
		"mweri wa kwanza":         "january",
		"mweri wa katatu":         "march",
		"mweri wa kaana":          "april",
		"mweri wa kaili":          "february",
		"mweri wa ikumi":          "october",
		"mweri wa nane":           "august",
		"mweri wa saba":           "july",
		"mweri wa sita":           "june",
		"mweri wa tanu":           "may",
		"mweri wa tisa":           "september",
		"ijumatatu":               "monday",
		"ijumamosi":               "saturday",
		"ijumapili":               "sunday",
		"ijumatano":               "wednesday",
		"kang'ama":                "am",
		"alhamisi":                "thursday",
		"ijumanne":                "tuesday",
		"kingoto":                 "pm",
		"sekunde":                 "second",
		"ijumaa":                  "friday",
		"dakika":                  "minute",
		"mfiri":                   "day",
		"mweri":                   "month",
		"iwiki":                   "week",
		"muaka":                   "year",
		"isaa":                    "hour",
		"ijtn":                    "wednesday",
		"m12":                     "december",
		"iju":                     "friday",
		"gmt":                     "gmt",
		"ijt":                     "monday",
		"m11":                     "november",
		"m10":                     "october",
		"ijm":                     "saturday",
		"ijp":                     "sunday",
		"alh":                     "thursday",
		"ijn":                     "tuesday",
		"utc":                     "utc",
		"am":                      "am",
		"m4":                      "april",
		"m8":                      "august",
		"m2":                      "february",
		"m1":                      "january",
		"m7":                      "july",
		"m6":                      "june",
		"m3":                      "march",
		"m5":                      "may",
		"pm":                      "pm",
		"m9":                      "september",
		"'":                       "",
		",":                       "",
		";":                       "",
		"@":                       "",
		"[":                       "",
		"]":                       "",
		"|":                       "",
		" ":                       " ",
		"+":                       "+",
		"-":                       "-",
		".":                       ".",
		"/":                       "/",
		":":                       ":",
		"z":                       "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"ng'ama":      "in 1 day",
		"linu":        "0 day ago",
		"hiyo":        "1 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(mweri wa ikumi na mbili|mweri wa ikumi na moja|mweri wa katatu|mweri wa kwanza|mweri wa ikumi|mweri wa kaana|mweri wa kaili|mweri wa nane|mweri wa saba|mweri wa sita|mweri wa tanu|mweri wa tisa|this minute|last month|next month|this month|ijumamosi|ijumapili|ijumatano|ijumatatu|last week|last year|next week|next year|this hour|this week|this year|alhamisi|ijumanne|kang'ama|kingoto|sekunde|dakika|ijumaa|ng'ama|iwiki|mfiri|muaka|mweri|hiyo|ijtn|isaa|linu|alh|gmt|ijm|ijn|ijp|ijt|iju|m10|m11|m12|now|utc|\+|\.|\[|\]|\||am|m1|m2|m3|m4|m5|m6|m7|m8|m9|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
