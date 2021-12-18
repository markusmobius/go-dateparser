// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bez_Locale = merge(nil, LocaleData{
	Name:      "bez",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"pa mwedzi gwa kumi na mbili": "december",
		"pa mwedzi gwa kumi na moja":  "november",
		"pa mwedzi gwa wuvili":        "february",
		"pa mwedzi gwa hutala":        "january",
		"pa mwedzi gwa wudatu":        "march",
		"pa mwedzi gwa wuhanu":        "may",
		"pa mwedzi gwa wutai":         "april",
		"pa mwedzi gwa nane":          "august",
		"pa mwedzi gwa saba":          "july",
		"pa mwedzi gwa sita":          "june",
		"pa mwedzi gwa kumi":          "october",
		"pa mwedzi gwa tisa":          "september",
		"pa shahulembela":             "saturday",
		"pa shahuviluha":              "monday",
		"mlungu gumamfu":              "week",
		"this minute":                 "0 minute ago",
		"this month":                  "0 month ago",
		"last month":                  "1 month ago",
		"next month":                  "in 1 month",
		"pa mulungu":                  "sunday",
		"neng'u ni":                   "0 day ago",
		"this hour":                   "0 hour ago",
		"this week":                   "0 week ago",
		"this year":                   "0 year ago",
		"last week":                   "1 week ago",
		"last year":                   "1 year ago",
		"pa hihanu":                   "friday",
		"next week":                   "in 1 week",
		"next year":                   "in 1 year",
		"pa hitayi":                   "thursday",
		"pa hivili":                   "tuesday",
		"pa hidatu":                   "wednesday",
		"pamilau":                     "am",
		"pamunyi":                     "pm",
		"sekunde":                     "second",
		"hilawu":                      "in 1 day",
		"dakika":                      "minute",
		"mwedzi":                      "month",
		"igolo":                       "1 day ago",
		"mwaha":                       "year",
		"sihu":                        "day",
		"now":                         "0 second ago",
		"tai":                         "april",
		"nan":                         "august",
		"kmb":                         "december",
		"hih":                         "friday",
		"gmt":                         "gmt",
		"saa":                         "hour",
		"hut":                         "january",
		"sab":                         "july",
		"sit":                         "june",
		"dat":                         "march",
		"han":                         "may",
		"vil":                         "monday",
		"kmj":                         "november",
		"kum":                         "october",
		"lem":                         "saturday",
		"tis":                         "september",
		"mul":                         "sunday",
		"hit":                         "thursday",
		"hiv":                         "tuesday",
		"utc":                         "utc",
		"hid":                         "wednesday",
		"am":                          "am",
		"pm":                          "pm",
		"'":                           "",
		",":                           "",
		";":                           "",
		"@":                           "",
		"[":                           "",
		"]":                           "",
		"|":                           "",
		" ":                           " ",
		"+":                           "+",
		"-":                           "-",
		".":                           ".",
		"/":                           "/",
		":":                           ":",
		"z":                           "z",
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)()(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^()$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(pa mwedzi gwa kumi na mbili|pa mwedzi gwa kumi na moja|pa mwedzi gwa hutala|pa mwedzi gwa wudatu|pa mwedzi gwa wuhanu|pa mwedzi gwa wuvili|pa mwedzi gwa wutai|pa mwedzi gwa kumi|pa mwedzi gwa nane|pa mwedzi gwa saba|pa mwedzi gwa sita|pa mwedzi gwa tisa|pa shahulembela|mlungu gumamfu|pa shahuviluha|this minute|last month|next month|pa mulungu|this month|last week|last year|neng'u ni|next week|next year|pa hidatu|pa hihanu|pa hitayi|pa hivili|this hour|this week|this year|pamilau|pamunyi|sekunde|dakika|hilawu|mwedzi|igolo|mwaha|sihu|dat|gmt|han|hid|hih|hit|hiv|hut|kmb|kmj|kum|lem|mul|nan|now|saa|sab|sit|tai|tis|utc|vil|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
