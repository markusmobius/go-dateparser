// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ee_Locale = merge(nil, LocaleData{
	Name:      "ee",
	DateOrder: "MDY",
	Charset:   []rune("+,-./;@[]bcdefghiklnorstuvwxyz|ŋƒɔɖɛɣ"),
	Translations: map[string]string{
		"adeɛmekpɔxe": "november",
		"kɔsiɖa ɖeka": "week",
		"deasiamime":  "august",
		"aɖabaƒoƒo":   "minute",
		"siamlɔm":     "july",
		"memleɖa":     "saturday",
		"anyɔnyɔ":     "september",
		"afɔfie":      "april",
		"dzodze":      "february",
		"gaƒoƒo":      "hour",
		"tedoxe":      "march",
		"sekend":      "second",
		"kɔsiɖa":      "sunday",
		"yawoɖa":      "thursday",
		"ŋkeke":       "day",
		"dzome":       "december",
		"dzove":       "january",
		"dzoɖa":       "monday",
		"ɣleti":       "month",
		"ɣetrɔ":       "pm",
		"blaɖa":       "tuesday",
		"fiɖa":        "friday",
		"masa":        "june",
		"dama":        "may",
		"kele":        "october",
		"kuɖa":        "wednesday",
		"ŋdi":         "am",
		"afɔ":         "april",
		"dea":         "august",
		"dzm":         "december",
		"dzd":         "february",
		"fiɖ":         "friday",
		"gmt":         "gmt",
		"dzv":         "january",
		"sia":         "july",
		"mas":         "june",
		"ted":         "march",
		"dam":         "may",
		"dzo":         "monday",
		"ade":         "november",
		"kel":         "october",
		"mem":         "saturday",
		"any":         "september",
		"kɔs":         "sunday",
		"yaw":         "thursday",
		"bla":         "tuesday",
		"utc":         "utc",
		"kuɖ":         "wednesday",
		"am":          "am",
		"pm":          "pm",
		"ƒe":          "year",
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
	RelativeType: map[string]string{
		"kɔsiɖa si gbɔ na": "in 1 week",
		"kɔsiɖa si va yi":  "1 week ago",
		"ɣleti si gbɔ na":  "in 1 month",
		"ɣleti si va yi":   "1 month ago",
		"etsɔ si va yi":    "1 day ago",
		"etsɔ si gbɔna":    "in 1 day",
		"ƒe si gbɔ na":     "in 1 year",
		"this minute":      "0 minute ago",
		"ƒe si va yi":      "1 year ago",
		"kɔsiɖa sia":       "0 week ago",
		"this hour":        "0 hour ago",
		"ɣleti sia":        "0 month ago",
		"ƒe sia":           "0 year ago",
		"egbe":             "0 day ago",
		"fifi":             "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)aɖabaƒoƒo (\d+) si wo va yi`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)gaƒoƒo (\d+) si wo va yi`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)aɖabaƒoƒo (\d+) si va yi`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)sekend (\d+) si wo va yi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)kɔsiɖa (\d+) si wo va yi`), "$1 week ago"},
		{regexp.MustCompile(`(?i)le aɖabaƒoƒo (\d+) wo me`), "in $1 minute"},
		{regexp.MustCompile(`(?i)ŋkeke (\d+) si wo va yi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ɣleti (\d+) si wo va yi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)le ƒe (\d+) si va yi me`), "$1 year ago"},
		{regexp.MustCompile(`(?i)le ƒe (\d+) si gbɔna me`), "in $1 year"},
		{regexp.MustCompile(`(?i)gaƒoƒo (\d+) si va yi`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)sekend (\d+) si va yi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)kɔsiɖa (\d+) si va yi`), "$1 week ago"},
		{regexp.MustCompile(`(?i)le gaƒoƒo (\d+) wo me`), "in $1 hour"},
		{regexp.MustCompile(`(?i)le aɖabaƒoƒo (\d+) me`), "in $1 minute"},
		{regexp.MustCompile(`(?i)le sekend (\d+) wo me`), "in $1 second"},
		{regexp.MustCompile(`(?i)le kɔsiɖa (\d+) wo me`), "in $1 week"},
		{regexp.MustCompile(`(?i)ŋkeke (\d+) si va yi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ɣleti (\d+) si va yi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)ƒe (\d+) si va yi me`), "$1 year ago"},
		{regexp.MustCompile(`(?i)ƒe (\d+) si wo va yi`), "$1 year ago"},
		{regexp.MustCompile(`(?i)le ŋkeke (\d+) wo me`), "in $1 day"},
		{regexp.MustCompile(`(?i)le ɣleti (\d+) wo me`), "in $1 month"},
		{regexp.MustCompile(`(?i)le gaƒoƒo (\d+) me`), "in $1 hour"},
		{regexp.MustCompile(`(?i)le sekend (\d+) me`), "in $1 second"},
		{regexp.MustCompile(`(?i)le kɔsiɖa (\d+) me`), "in $1 week"},
		{regexp.MustCompile(`(?i)ƒe (\d+) si va yi`), "$1 year ago"},
		{regexp.MustCompile(`(?i)le ŋkeke (\d+) me`), "in $1 day"},
		{regexp.MustCompile(`(?i)le ɣleti (\d+) me`), "in $1 month"},
		{regexp.MustCompile(`(?i)le ƒe (\d+) me`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(aɖabaƒoƒo \d+ si wo va yi|aɖabaƒoƒo \d+ si va yi|gaƒoƒo \d+ si wo va yi|kɔsiɖa \d+ si wo va yi|le aɖabaƒoƒo \d+ wo me|sekend \d+ si wo va yi|le ƒe \d+ si gbɔna me|le ƒe \d+ si va yi me|ŋkeke \d+ si wo va yi|ɣleti \d+ si wo va yi|gaƒoƒo \d+ si va yi|kɔsiɖa \d+ si va yi|le aɖabaƒoƒo \d+ me|le gaƒoƒo \d+ wo me|le kɔsiɖa \d+ wo me|le sekend \d+ wo me|sekend \d+ si va yi|le ŋkeke \d+ wo me|le ɣleti \d+ wo me|ŋkeke \d+ si va yi|ƒe \d+ si va yi me|ƒe \d+ si wo va yi|ɣleti \d+ si va yi|le gaƒoƒo \d+ me|le kɔsiɖa \d+ me|le sekend \d+ me|le ŋkeke \d+ me|le ɣleti \d+ me|ƒe \d+ si va yi|le ƒe \d+ me)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(aɖabaƒoƒo \d+ si wo va yi|aɖabaƒoƒo \d+ si va yi|gaƒoƒo \d+ si wo va yi|kɔsiɖa \d+ si wo va yi|le aɖabaƒoƒo \d+ wo me|sekend \d+ si wo va yi|le ƒe \d+ si gbɔna me|le ƒe \d+ si va yi me|ŋkeke \d+ si wo va yi|ɣleti \d+ si wo va yi|gaƒoƒo \d+ si va yi|kɔsiɖa \d+ si va yi|le aɖabaƒoƒo \d+ me|le gaƒoƒo \d+ wo me|le kɔsiɖa \d+ wo me|le sekend \d+ wo me|sekend \d+ si va yi|le ŋkeke \d+ wo me|le ɣleti \d+ wo me|ŋkeke \d+ si va yi|ƒe \d+ si va yi me|ƒe \d+ si wo va yi|ɣleti \d+ si va yi|le gaƒoƒo \d+ me|le kɔsiɖa \d+ me|le sekend \d+ me|le ŋkeke \d+ me|le ɣleti \d+ me|ƒe \d+ si va yi|le ƒe \d+ me)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(kɔsiɖa si gbɔ na|kɔsiɖa si va yi|ɣleti si gbɔ na|ɣleti si va yi|etsɔ si gbɔna|etsɔ si va yi|ƒe si gbɔ na|adeɛmekpɔxe|kɔsiɖa ɖeka|this minute|ƒe si va yi|deasiamime|kɔsiɖa sia|aɖabaƒoƒo|this hour|ɣleti sia|anyɔnyɔ|memleɖa|siamlɔm|afɔfie|dzodze|gaƒoƒo|kɔsiɖa|sekend|tedoxe|yawoɖa|ƒe sia|blaɖa|dzome|dzove|dzoɖa|ŋkeke|ɣetrɔ|ɣleti|dama|egbe|fifi|fiɖa|kele|kuɖa|masa|ade|afɔ|any|bla|dam|dea|dzd|dzm|dzo|dzv|fiɖ|gmt|kel|kuɖ|kɔs|mas|mem|sia|ted|utc|yaw|ŋdi|\+|\.|\[|\]|\||am|pm|ƒe| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ee_TG_Locale = merge(&ee_Locale, LocaleData{
	Name:            "ee-TG",
	DateOrder:       "MDY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(aɖabaƒoƒo \d+ si wo va yi|aɖabaƒoƒo \d+ si va yi|gaƒoƒo \d+ si wo va yi|kɔsiɖa \d+ si wo va yi|le aɖabaƒoƒo \d+ wo me|sekend \d+ si wo va yi|le ƒe \d+ si gbɔna me|le ƒe \d+ si va yi me|ŋkeke \d+ si wo va yi|ɣleti \d+ si wo va yi|gaƒoƒo \d+ si va yi|kɔsiɖa \d+ si va yi|le aɖabaƒoƒo \d+ me|le gaƒoƒo \d+ wo me|le kɔsiɖa \d+ wo me|le sekend \d+ wo me|sekend \d+ si va yi|le ŋkeke \d+ wo me|le ɣleti \d+ wo me|ŋkeke \d+ si va yi|ƒe \d+ si va yi me|ƒe \d+ si wo va yi|ɣleti \d+ si va yi|le gaƒoƒo \d+ me|le kɔsiɖa \d+ me|le sekend \d+ me|le ŋkeke \d+ me|le ɣleti \d+ me|ƒe \d+ si va yi|le ƒe \d+ me)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(aɖabaƒoƒo \d+ si wo va yi|aɖabaƒoƒo \d+ si va yi|gaƒoƒo \d+ si wo va yi|kɔsiɖa \d+ si wo va yi|le aɖabaƒoƒo \d+ wo me|sekend \d+ si wo va yi|le ƒe \d+ si gbɔna me|le ƒe \d+ si va yi me|ŋkeke \d+ si wo va yi|ɣleti \d+ si wo va yi|gaƒoƒo \d+ si va yi|kɔsiɖa \d+ si va yi|le aɖabaƒoƒo \d+ me|le gaƒoƒo \d+ wo me|le kɔsiɖa \d+ wo me|le sekend \d+ wo me|sekend \d+ si va yi|le ŋkeke \d+ wo me|le ɣleti \d+ wo me|ŋkeke \d+ si va yi|ƒe \d+ si va yi me|ƒe \d+ si wo va yi|ɣleti \d+ si va yi|le gaƒoƒo \d+ me|le kɔsiɖa \d+ me|le sekend \d+ me|le ŋkeke \d+ me|le ɣleti \d+ me|ƒe \d+ si va yi|le ƒe \d+ me)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(kɔsiɖa si gbɔ na|kɔsiɖa si va yi|ɣleti si gbɔ na|ɣleti si va yi|etsɔ si gbɔna|etsɔ si va yi|ƒe si gbɔ na|adeɛmekpɔxe|kɔsiɖa ɖeka|this minute|ƒe si va yi|deasiamime|kɔsiɖa sia|aɖabaƒoƒo|this hour|ɣleti sia|anyɔnyɔ|memleɖa|siamlɔm|afɔfie|dzodze|gaƒoƒo|kɔsiɖa|sekend|tedoxe|yawoɖa|ƒe sia|blaɖa|dzome|dzove|dzoɖa|ŋkeke|ɣetrɔ|ɣleti|dama|egbe|fifi|fiɖa|kele|kuɖa|masa|ade|afɔ|any|bla|dam|dea|dzd|dzm|dzo|dzv|fiɖ|gmt|kel|kuɖ|kɔs|mas|mem|sia|ted|utc|yaw|ŋdi|\+|\.|\[|\]|\||am|pm|ƒe| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
