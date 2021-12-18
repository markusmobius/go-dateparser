// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uz_Arab_Locale = merge(nil, LocaleData{
	Name:      "uz-Arab",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"yesterday":   "1 day ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"tomorrow":    "in 1 day",
		"چهارشنبه":    "wednesday",
		"پنجشنبه":     "thursday",
		"سه‌شنبه":     "tuesday",
		"فبروری":      "february",
		"minute":      "minute",
		"دوشنبه":      "monday",
		"اکتوبر":      "october",
		"second":      "second",
		"سپتمبر":      "september",
		"یکشنبه":      "sunday",
		"today":       "0 day ago",
		"اپریل":       "april",
		"دسمبر":       "december",
		"جنوری":       "january",
		"جولای":       "july",
		"month":       "month",
		"نومبر":       "november",
		"اگست":        "august",
		"جمعه":        "friday",
		"hour":        "hour",
		"مارچ":        "march",
		"شنبه":        "saturday",
		"week":        "week",
		"year":        "year",
		"now":         "0 second ago",
		"اپر":         "april",
		"اگس":         "august",
		"day":         "day",
		"دسم":         "december",
		"فبر":         "february",
		"gmt":         "gmt",
		"جنو":         "january",
		"جول":         "july",
		"جون":         "june",
		"مار":         "march",
		"نوم":         "november",
		"اکت":         "october",
		"سپت":         "september",
		"utc":         "utc",
		"am":          "am",
		"می":          "may",
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
		"ج":           "friday",
		"د":           "monday",
		"ش":           "saturday",
		"ی":           "sunday",
		"پ":           "thursday",
		"س":           "tuesday",
		"چ":           "wednesday",
		"z":           "z",
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)()(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^()$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|tomorrow|چهارشنبه|سه‌شنبه|پنجشنبه|minute|second|اکتوبر|دوشنبه|سپتمبر|فبروری|یکشنبه|month|today|اپریل|جنوری|جولای|دسمبر|نومبر|hour|week|year|اگست|جمعه|شنبه|مارچ|day|gmt|now|utc|اپر|اکت|اگس|جنو|جول|جون|دسم|سپت|فبر|مار|نوم|\+|\.|\[|\]|\||am|pm|می| |'|,|-|/|:|;|@|z|ج|د|س|ش|پ|چ|ی)((?:\z|\W|_|\d).*)$`),
})
