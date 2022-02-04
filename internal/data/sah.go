// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sah_Locale = merge(nil, LocaleData{
	Name:      "sah",
	DateOrder: "YMD",
	Charset:   []rune(`ceghinorstuzабгдийклмнопрстухчыьэҕҥүһө`),
	Translations: map[string]string{
		"атырдьых ыиын": "august",
		"атырдьых ыиа":  "august",
		"бэнидиэнньик":  "monday",
		"балаҕан ыиын":  "september",
		"баскыһыанньа":  "sunday",
		"оптуорунньук":  "tuesday",
		"кулун тутар":   "march",
		"балаҕан ыиа":   "september",
		"муус устар":    "april",
		"тохсунньу":     "january",
		"ахсынньы":      "december",
		"бээтиҥсэ":      "friday",
		"бэс ыиын":      "june",
		"ыам ыиын":      "may",
		"сэтинньи":      "november",
		"алтынньы":      "october",
		"сөкүүндэ":      "second",
		"олунньу":       "february",
		"от ыиын":       "july",
		"бэс ыиа":       "june",
		"ыам ыиа":       "may",
		"мүнүүтэ":       "minute",
		"субуота":       "saturday",
		"чэппиэр":       "thursday",
		"нэдиэлэ":       "week",
		"от ыиа":        "july",
		"сэрэдэ":        "wednesday",
		"олун":          "february",
		"чаас":          "hour",
		"тохс":          "january",
		"мсу":           "april",
		"атр":           "august",
		"күн":           "day",
		"ахс":           "december",
		"gmt":           "gmt",
		"оти":           "july",
		"бэс":           "june",
		"клн":           "march",
		"ыам":           "may",
		"сэт":           "november",
		"алт":           "october",
		"блҕ":           "september",
		"utc":           "utc",
		"сыл":           "year",
		"am":            "am",
		"эи":            "am",
		"бэ":            "friday",
		"бн":            "monday",
		"ыи":            "month",
		"pm":            "pm",
		"эк":            "pm",
		"сб":            "saturday",
		"бс":            "sunday",
		"чп":            "thursday",
		"оп":            "tuesday",
		"сэ":            "wednesday",
		"'":             "",
		",":             "",
		";":             "",
		"@":             "",
		"[":             "",
		"]":             "",
		"|":             "",
		" ":             " ",
		"+":             "+",
		"-":             "-",
		".":             ".",
		"/":             "/",
		":":             ":",
		"z":             "z",
	},
	RelativeType: map[string]string{
		"ааспыт нэдиэлэ": "1 week ago",
		"кэлэр нэдиэлэ":  "in 1 week",
		"this minute":    "0 minute ago",
		"аныгыскы ыи":    "in 1 month",
		"бу нэдиэлэ":     "0 week ago",
		"this hour":      "0 hour ago",
		"ааспыт ыи":      "1 month ago",
		"былырыын":       "1 year ago",
		"билигин":        "0 second ago",
		"бэҕэһээ":        "1 day ago",
		"сарсын":         "in 1 day",
		"бүгүн":          "0 day ago",
		"бу ыи":          "0 month ago",
		"быиыл":          "0 year ago",
		"эһиил":          "in 1 year",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) сөкүүндэ ынараа өттүгэр`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) мүнүүтэ ынараа өттүгэр`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) нэдиэлэ анараа өттүгэр`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) чаас ынараа өттүгэр`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) күн ынараа өттүгэр`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) сөк анараа өттүгэр`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) сыл ынараа өттүгэр`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ыи ынараа өттүгэр`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) сөкүүндэннэн`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) мүнүүтэннэн`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) нэдиэлэннэн`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) чааһынан`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) күнүнэн`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) сылынан`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ыиынан`), "in $1 month"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ сөкүүндэ ынараа өттүгэр|\d+ мүнүүтэ ынараа өттүгэр|\d+ нэдиэлэ анараа өттүгэр|\d+ чаас ынараа өттүгэр|\d+ күн ынараа өттүгэр|\d+ сыл ынараа өттүгэр|\d+ сөк анараа өттүгэр|\d+ ыи ынараа өттүгэр|\d+ сөкүүндэннэн|\d+ мүнүүтэннэн|\d+ нэдиэлэннэн|\d+ чааһынан|\d+ күнүнэн|\d+ сылынан|\d+ ыиынан)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ сөкүүндэ ынараа өттүгэр|\d+ мүнүүтэ ынараа өттүгэр|\d+ нэдиэлэ анараа өттүгэр|\d+ чаас ынараа өттүгэр|\d+ күн ынараа өттүгэр|\d+ сыл ынараа өттүгэр|\d+ сөк анараа өттүгэр|\d+ ыи ынараа өттүгэр|\d+ сөкүүндэннэн|\d+ мүнүүтэннэн|\d+ нэдиэлэннэн|\d+ чааһынан|\d+ күнүнэн|\d+ сылынан|\d+ ыиынан)$`),
	KnownWords:      []string{"ааспыт нэдиэлэ", "атырдьых ыиын", "кэлэр нэдиэлэ", "атырдьых ыиа", "балаҕан ыиын", "баскыһыанньа", "бэнидиэнньик", "оптуорунньук", "this minute", "аныгыскы ыи", "балаҕан ыиа", "кулун тутар", "бу нэдиэлэ", "муус устар", "this hour", "ааспыт ыи", "тохсунньу", "алтынньы", "ахсынньы", "былырыын", "бэс ыиын", "бээтиҥсэ", "сэтинньи", "сөкүүндэ", "ыам ыиын", "билигин", "бэс ыиа", "бэҕэһээ", "мүнүүтэ", "нэдиэлэ", "олунньу", "от ыиын", "субуота", "чэппиэр", "ыам ыиа", "от ыиа", "сарсын", "сэрэдэ", "бу ыи", "быиыл", "бүгүн", "эһиил", "олун", "тохс", "чаас", "gmt", "utc", "алт", "атр", "ахс", "блҕ", "бэс", "клн", "күн", "мсу", "оти", "сыл", "сэт", "ыам", "am", "pm", "бн", "бс", "бэ", "оп", "сб", "сэ", "чп", "ыи", "эи", "эк", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
