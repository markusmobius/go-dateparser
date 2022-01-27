// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pa_Arab_Locale = merge(nil, LocaleData{
	Name:      "pa-Arab",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]cdeghiklnorstuwxyz|ابتجدرسعفلمنويٹپچکگھہی"),
	Translations: map[string]string{
		"جولايی": "july",
		"اکتوبر": "october",
		"second": "second",
		"جمعرات": "thursday",
		"اپریل":  "april",
		"دسمبر":  "december",
		"فروری":  "february",
		"گھنٹا":  "hour",
		"جنوری":  "january",
		"مہينا":  "month",
		"نومبر":  "november",
		"ستمبر":  "september",
		"اتوار":  "sunday",
		"اگست":   "august",
		"جمعہ":   "friday",
		"مارچ":   "march",
		"منگل":   "tuesday",
		"ہفتہ":   "week",
		"ورھا":   "year",
		"دين":    "day",
		"gmt":    "gmt",
		"جون":    "june",
		"منٹ":    "minute",
		"پیر":    "monday",
		"utc":    "utc",
		"بدھ":    "wednesday",
		"am":     "am",
		"مي":     "may",
		"pm":     "pm",
		"'":      "",
		",":      "",
		";":      "",
		"@":      "",
		"[":      "",
		"]":      "",
		"|":      "",
		" ":      " ",
		"+":      "+",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		"z":      "z",
	},
	RelativeType: map[string]string{
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
		"today":       "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|tomorrow|second|اکتوبر|جمعرات|جولايی|today|اتوار|اپریل|جنوری|دسمبر|ستمبر|فروری|مہينا|نومبر|گھنٹا|اگست|جمعہ|مارچ|منگل|ورھا|ہفتہ|gmt|now|utc|بدھ|جون|دين|منٹ|پیر|\+|\.|\[|\]|\||am|pm|مي| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
