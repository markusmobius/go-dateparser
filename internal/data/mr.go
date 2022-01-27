// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mr_Locale = merge(nil, LocaleData{
	Name:      "mr",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]cgtuz|आउएऑकगचजटठडढतदधनपफबमयरलळवशषसहािीो"),
	Translations: map[string]string{
		"फबरवारी": "february",
		"जानवारी": "january",
		"शकरवार":  "friday",
		"सोमवार":  "monday",
		"नोवहबर":  "november",
		"ऑकटोबर":  "october",
		"शनिवार":  "saturday",
		"रविवार":  "sunday",
		"मगळवार":  "tuesday",
		"एपरिल":   "april",
		"डिसबर":   "december",
		"मिनिट":   "minute",
		"महिना":   "month",
		"सपटबर":   "september",
		"गरवार":   "thursday",
		"बधवार":   "wednesday",
		"आठवडा":   "week",
		"एपरि":    "april",
		"ऑगसट":    "august",
		"दिवस":    "day",
		"मारच":    "march",
		"नोवह":    "november",
		"ऑकटो":    "october",
		"डिस":     "december",
		"फबर":     "february",
		"शकर":     "friday",
		"gmt":     "gmt",
		"तास":     "hour",
		"जान":     "january",
		"सोम":     "monday",
		"शनि":     "saturday",
		"सकद":     "second",
		"सपट":     "september",
		"रवि":     "sunday",
		"मगळ":     "tuesday",
		"utc":     "utc",
		"वरष":     "year",
		"am":      "am",
		"मप":      "am",
		"ऑग":      "august",
		"जल":      "july",
		"जन":      "june",
		"मि":      "minute",
		"pm":      "pm",
		"मउ":      "pm",
		"गर":      "thursday",
		"बध":      "wednesday",
		"'":       "",
		",":       "",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"|":       "",
		" ":       " ",
		"+":       "+",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		"म":       "may",
		"स":       "second",
		"z":       "z",
	},
	RelativeType: map[string]string{
		"मागील महिना": "1 month ago",
		"मागील आठवडा": "1 week ago",
		"या मिनिटात":  "0 minute ago",
		"पढील महिना":  "in 1 month",
		"पढील आठवडा":  "in 1 week",
		"मागील वरष":   "1 year ago",
		"हा महिना":    "0 month ago",
		"हा आठवडा":    "0 week ago",
		"पढील वरष":    "in 1 year",
		"तासात":       "0 hour ago",
		"ह वरष":       "0 year ago",
		"आतता":        "0 second ago",
		"उदया":        "in 1 day",
		"काल":         "1 day ago",
		"आज":          "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) मिनिटापरवी`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) महिनयापरवी`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) आठवडयापरवी`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) दिवसापरवी`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) मिनि परवी`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) मिनिटामधय`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) महिनयामधय`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) आठवडयामधय`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) तासापरवी`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) सकदापरवी`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) वरषापरवी`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) दिवसामधय`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) मिनि मधय`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) तासामधय`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) सकदामधय`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) वरषामधय`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) स परवी`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) स मधय`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ आठवडयापरवी|\d+ महिनयापरवी|\d+ मिनिटापरवी|\d+ आठवडयामधय|\d+ दिवसापरवी|\d+ महिनयामधय|\d+ मिनि परवी|\d+ मिनिटामधय|\d+ तासापरवी|\d+ दिवसामधय|\d+ मिनि मधय|\d+ वरषापरवी|\d+ सकदापरवी|\d+ तासामधय|\d+ वरषामधय|\d+ सकदामधय|\d+ स परवी|\d+ स मधय)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ आठवडयापरवी|\d+ महिनयापरवी|\d+ मिनिटापरवी|\d+ आठवडयामधय|\d+ दिवसापरवी|\d+ महिनयामधय|\d+ मिनि परवी|\d+ मिनिटामधय|\d+ तासापरवी|\d+ दिवसामधय|\d+ मिनि मधय|\d+ वरषापरवी|\d+ सकदापरवी|\d+ तासामधय|\d+ वरषामधय|\d+ सकदामधय|\d+ स परवी|\d+ स मधय)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(मागील आठवडा|मागील महिना|पढील आठवडा|पढील महिना|या मिनिटात|मागील वरष|पढील वरष|हा आठवडा|हा महिना|जानवारी|फबरवारी|ऑकटोबर|नोवहबर|मगळवार|रविवार|शकरवार|शनिवार|सोमवार|आठवडा|एपरिल|गरवार|डिसबर|तासात|बधवार|महिना|मिनिट|सपटबर|ह वरष|आतता|उदया|एपरि|ऑकटो|ऑगसट|दिवस|नोवह|मारच|gmt|utc|काल|जान|डिस|तास|फबर|मगळ|रवि|वरष|शकर|शनि|सकद|सपट|सोम|\+|\.|\[|\]|\||am|pm|आज|ऑग|गर|जन|जल|बध|मउ|मप|मि| |'|,|-|/|:|;|@|z|म|स)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
