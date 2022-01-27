// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ksb_Locale = merge(nil, LocaleData{
	Name:      "ksb",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]bcdefghijklnorstuvwxyz|"),
	Translations: map[string]string{
		"jumaatatu": "monday",
		"jumaamosi": "saturday",
		"jumaatano": "wednesday",
		"febluali":  "february",
		"nyiaghuo":  "pm",
		"septemba":  "september",
		"jumaapii":  "sunday",
		"alhamisi":  "thursday",
		"desemba":   "december",
		"januali":   "january",
		"novemba":   "november",
		"sekunde":   "second",
		"jumaane":   "tuesday",
		"ng'waka":   "year",
		"aplili":    "april",
		"agosti":    "august",
		"ijumaa":    "friday",
		"dakika":    "minute",
		"ng'ezi":    "month",
		"oktoba":    "october",
		"makeo":     "am",
		"julai":     "july",
		"machi":     "march",
		"siku":      "day",
		"juni":      "june",
		"niki":      "week",
		"apr":       "april",
		"ago":       "august",
		"des":       "december",
		"feb":       "february",
		"iju":       "friday",
		"gmt":       "gmt",
		"saa":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mac":       "march",
		"mei":       "may",
		"jtt":       "monday",
		"nov":       "november",
		"okt":       "october",
		"jmo":       "saturday",
		"sep":       "september",
		"jpi":       "sunday",
		"alh":       "thursday",
		"jmn":       "tuesday",
		"utc":       "utc",
		"jtn":       "wednesday",
		"am":        "am",
		"pm":        "pm",
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
		"z":         "z",
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
		"evi eo":      "0 day ago",
		"keloi":       "in 1 day",
		"ghuo":        "1 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|jumaamosi|jumaatano|jumaatatu|last week|last year|next week|next year|this hour|this week|this year|alhamisi|febluali|jumaapii|nyiaghuo|septemba|desemba|januali|jumaane|ng'waka|novemba|sekunde|agosti|aplili|dakika|evi eo|ijumaa|ng'ezi|oktoba|julai|keloi|machi|makeo|ghuo|juni|niki|siku|ago|alh|apr|des|feb|gmt|iju|jan|jmn|jmo|jpi|jtn|jtt|jul|jun|mac|mei|nov|now|okt|saa|sep|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
