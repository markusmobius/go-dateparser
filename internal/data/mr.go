// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mr_Locale = merge(nil, LocaleData{
	Name:      "mr",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzंआउएऑकगचजटठडढतदधनपफबमयरलळवशषसहािीुूेैो्`),
	Translations: map[string][]string{
		"जानवारी": {"january"},
		"फबरवारी": {"february"},
		"ऑकटोबर":  {"october"},
		"नोवहबर":  {"november"},
		"मगळवार":  {"tuesday"},
		"रविवार":  {"sunday"},
		"शकरवार":  {"friday"},
		"शनिवार":  {"saturday"},
		"सोमवार":  {"monday"},
		"आठवडा":   {"week"},
		"एपरिल":   {"april"},
		"गरवार":   {"thursday"},
		"डिसबर":   {"december"},
		"बधवार":   {"wednesday"},
		"महिना":   {"month"},
		"मिनिट":   {"minute"},
		"सपटबर":   {"september"},
		"एपरि":    {"april"},
		"ऑकटो":    {"october"},
		"ऑगसट":    {"august"},
		"दिवस":    {"day"},
		"नोवह":    {"november"},
		"मारच":    {"march"},
		"gmt":     {"gmt"},
		"utc":     {"utc"},
		"जान":     {"january"},
		"डिस":     {"december"},
		"तास":     {"hour"},
		"फबर":     {"february"},
		"मगळ":     {"tuesday"},
		"रवि":     {"sunday"},
		"वरष":     {"year"},
		"शकर":     {"friday"},
		"शनि":     {"saturday"},
		"सकद":     {"second"},
		"सपट":     {"september"},
		"सोम":     {"monday"},
		"am":      {"am"},
		"pm":      {"pm"},
		"ऑग":      {"august"},
		"गर":      {"thursday"},
		"जन":      {"june"},
		"जल":      {"july"},
		"बध":      {"wednesday"},
		"मउ":      {"pm"},
		"मप":      {"am"},
		"मि":      {"minute"},
		" ":       {" "},
		"'":       {""},
		"+":       {"+"},
		",":       {""},
		"-":       {"-"},
		".":       {"."},
		"/":       {"/"},
		":":       {":"},
		";":       {""},
		"@":       {""},
		"[":       {""},
		"]":       {""},
		"z":       {"z"},
		"|":       {""},
		"म":       {"may"},
		"स":       {"second"},
	},
	RelativeType: map[string]string{
		"मागील आठवडा": "1 week ago",
		"मागील महिना": "1 month ago",
		"पढील आठवडा":  "in 1 week",
		"पढील महिना":  "in 1 month",
		"या मिनिटात":  "0 minute ago",
		"मागील वरष":   "1 year ago",
		"पढील वरष":    "in 1 year",
		"हा आठवडा":    "0 week ago",
		"हा महिना":    "0 month ago",
		"तासात":       "0 hour ago",
		"ह वरष":       "0 year ago",
		"आतता":        "0 second ago",
		"उदया":        "in 1 day",
		"काल":         "1 day ago",
		"आज":          "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) आठवडयापरवी`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) महिनयापरवी`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) मिनिटापरवी`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) आठवडयामधय`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) दिवसापरवी`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) महिनयामधय`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) मिनि परवी`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) मिनिटामधय`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) तासापरवी`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) दिवसामधय`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) मिनि मधय`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) वरषापरवी`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) सकदापरवी`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) तासामधय`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) वरषामधय`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) सकदामधय`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) स परवी`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) स मधय`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* आठवडयापरवी|\d+[.,]?\d* महिनयापरवी|\d+[.,]?\d* मिनिटापरवी|\d+[.,]?\d* आठवडयामधय|\d+[.,]?\d* दिवसापरवी|\d+[.,]?\d* महिनयामधय|\d+[.,]?\d* मिनि परवी|\d+[.,]?\d* मिनिटामधय|\d+[.,]?\d* तासापरवी|\d+[.,]?\d* दिवसामधय|\d+[.,]?\d* मिनि मधय|\d+[.,]?\d* वरषापरवी|\d+[.,]?\d* सकदापरवी|\d+[.,]?\d* तासामधय|\d+[.,]?\d* वरषामधय|\d+[.,]?\d* सकदामधय|\d+[.,]?\d* स परवी|\d+[.,]?\d* स मधय)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* आठवडयापरवी|\d+[.,]?\d* महिनयापरवी|\d+[.,]?\d* मिनिटापरवी|\d+[.,]?\d* आठवडयामधय|\d+[.,]?\d* दिवसापरवी|\d+[.,]?\d* महिनयामधय|\d+[.,]?\d* मिनि परवी|\d+[.,]?\d* मिनिटामधय|\d+[.,]?\d* तासापरवी|\d+[.,]?\d* दिवसामधय|\d+[.,]?\d* मिनि मधय|\d+[.,]?\d* वरषापरवी|\d+[.,]?\d* सकदापरवी|\d+[.,]?\d* तासामधय|\d+[.,]?\d* वरषामधय|\d+[.,]?\d* सकदामधय|\d+[.,]?\d* स परवी|\d+[.,]?\d* स मधय)$`),
	KnownWords:      []string{"मागील आठवडा", "मागील महिना", "पढील आठवडा", "पढील महिना", "या मिनिटात", "मागील वरष", "पढील वरष", "हा आठवडा", "हा महिना", "जानवारी", "फबरवारी", "ऑकटोबर", "नोवहबर", "मगळवार", "रविवार", "शकरवार", "शनिवार", "सोमवार", "आठवडा", "एपरिल", "गरवार", "डिसबर", "तासात", "बधवार", "महिना", "मिनिट", "सपटबर", "ह वरष", "आतता", "उदया", "एपरि", "ऑकटो", "ऑगसट", "दिवस", "नोवह", "मारच", "gmt", "utc", "काल", "जान", "डिस", "तास", "फबर", "मगळ", "रवि", "वरष", "शकर", "शनि", "सकद", "सपट", "सोम", "am", "pm", "आज", "ऑग", "गर", "जन", "जल", "बध", "मउ", "मप", "मि", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "म", "स"},
})
