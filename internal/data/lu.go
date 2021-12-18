// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lu_Locale = merge(nil, LocaleData{
	Name:      "lu",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"this minute": "0 minute ago",
		"kabalashipu": "july",
		"kasunsukusu": "second",
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
		"lumungulu":   "may",
		"kaswekese":   "november",
		"lutongolo":   "september",
		"makelela":    "1 day ago",
		"tshidimu":    "year",
		"lushika":     "august",
		"lufuimi":     "june",
		"kasunsu":     "minute",
		"lungudi":     "october",
		"lumingu":     "sunday",
		"lubingu":     "week",
		"dituku":      "day",
		"luishi":      "february",
		"ngovya":      "friday",
		"malaba":      "in 1 day",
		"ciongo":      "january",
		"lusolo":      "march",
		"nkodya":      "monday",
		"ngondo":      "month",
		"dilolo":      "pm",
		"ndaaya":      "tuesday",
		"ndangu":      "wednesday",
		"dinda":       "am",
		"muuya":       "april",
		"ciswa":       "december",
		"njowa":       "thursday",
		"lelu":        "0 day ago",
		"lush":        "august",
		"diba":        "hour",
		"now":         "0 second ago",
		"muu":         "april",
		"cis":         "december",
		"lui":         "february",
		"ngv":         "friday",
		"gmt":         "gmt",
		"cio":         "january",
		"kab":         "july",
		"luf":         "june",
		"lus":         "march",
		"nko":         "monday",
		"kas":         "november",
		"lun":         "october",
		"lub":         "saturday",
		"lut":         "september",
		"lum":         "sunday",
		"njw":         "thursday",
		"ndy":         "tuesday",
		"utc":         "utc",
		"ndg":         "wednesday",
		"am":          "am",
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
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(kabalashipu|kasunsukusu|this minute|last month|next month|this month|kaswekese|last week|last year|lumungulu|lutongolo|next week|next year|this hour|this week|this year|makelela|tshidimu|kasunsu|lubingu|lufuimi|lumingu|lungudi|lushika|ciongo|dilolo|dituku|luishi|lusolo|malaba|ndaaya|ndangu|ngondo|ngovya|nkodya|ciswa|dinda|muuya|njowa|diba|lelu|lush|cio|cis|gmt|kab|kas|lub|luf|lui|lum|lun|lus|lut|muu|ndg|ndy|ngv|njw|nko|now|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
