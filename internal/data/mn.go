// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mn_Locale = merge(nil, LocaleData{
	Name:      "mn",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"өнгөрсөн долоо хоног": "1 week ago",
		"арван хоердугаар сар": "december",
		"арван нэгдүгээр сар":  "november",
		"ирэх долоо хоног":     "in 1 week",
		"энэ долоо хоног":      "0 week ago",
		"дөрөвдүгээр сар":      "april",
		"зургадугаар сар":      "june",
		"гуравдугаар сар":      "march",
		"наимдугаар сар":       "august",
		"хоердугаар сар":       "february",
		"аравдугаар сар":       "october",
		"нэгдүгээр сар":        "january",
		"долдугаар сар":        "july",
		"тавдугаар сар":        "may",
		"өнгөрсөн сар":         "1 month ago",
		"өнгөрсөн жил":         "1 year ago",
		"есдүгээр сар":         "september",
		"долоо хоног":          "week",
		"энэ минут":            "0 minute ago",
		"12-р сар":             "december",
		"ирэх сар":             "in 1 month",
		"ирэх жил":             "in 1 year",
		"11-р сар":             "november",
		"10-р сар":             "october",
		"өнөөдөр":              "0 day ago",
		"энэ цаг":              "0 hour ago",
		"энэ сар":              "0 month ago",
		"энэ жил":              "0 year ago",
		"өчигдөр":              "1 day ago",
		"4-р сар":              "april",
		"8-р сар":              "august",
		"2-р сар":              "february",
		"маргааш":              "in 1 day",
		"1-р сар":              "january",
		"7-р сар":              "july",
		"6-р сар":              "june",
		"3-р сар":              "march",
		"5-р сар":              "may",
		"9-р сар":              "september",
		"баасан":               "friday",
		"секунд":               "second",
		"мягмар":               "tuesday",
		"лхагва":               "wednesday",
		"минут":                "minute",
		"даваа":                "monday",
		"бямба":                "saturday",
		"пүрэв":                "thursday",
		"одоо":                 "0 second ago",
		"өдөр":                 "day",
		"gmt":                  "gmt",
		"цаг":                  "hour",
		"мин":                  "minute",
		"сар":                  "month",
		"сек":                  "second",
		"ням":                  "sunday",
		"utc":                  "utc",
		"жил":                  "year",
		"am":                   "am",
		"үө":                   "am",
		"ба":                   "friday",
		"да":                   "monday",
		"pm":                   "pm",
		"үх":                   "pm",
		"бя":                   "saturday",
		"ня":                   "sunday",
		"пү":                   "thursday",
		"мя":                   "tuesday",
		"лх":                   "wednesday",
		"7х":                   "week",
		"'":                    "",
		",":                    "",
		";":                    "",
		"@":                    "",
		"[":                    "",
		"]":                    "",
		"|":                    "",
		" ":                    " ",
		"+":                    "+",
		"-":                    "-",
		".":                    ".",
		"/":                    "/",
		":":                    ":",
		"ц":                    "hour",
		"z":                    "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секундын дараа`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) секундын өмнө`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) минутын дараа`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) минутын өмнө`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) өдриин дараа`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) цагиин дараа`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) 7х-иин дараа`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) жилиин дараа`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) өдриин өмнө`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) цагиин өмнө`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) 7х-иин өмнө`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) жилиин өмнө`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) сарын дараа`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) сарын өмнө`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) мин дараа`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) сек дараа`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) мин өмнө`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек өмнө`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) ц дараа`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ц өмнө`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) өдөрт`), "in $1 day"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(\d+ секундын дараа|\d+ минутын дараа|\d+ секундын өмнө|\d+ 7х-иин дараа|\d+ жилиин дараа|\d+ минутын өмнө|\d+ цагиин дараа|\d+ өдриин дараа|\d+ 7х-иин өмнө|\d+ жилиин өмнө|\d+ сарын дараа|\d+ цагиин өмнө|\d+ өдриин өмнө|\d+ сарын өмнө|\d+ мин дараа|\d+ сек дараа|\d+ мин өмнө|\d+ сек өмнө|\d+ ц дараа|\d+ ц өмнө|\d+ өдөрт)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секундын дараа|\d+ минутын дараа|\d+ секундын өмнө|\d+ 7х-иин дараа|\d+ жилиин дараа|\d+ минутын өмнө|\d+ цагиин дараа|\d+ өдриин дараа|\d+ 7х-иин өмнө|\d+ жилиин өмнө|\d+ сарын дараа|\d+ цагиин өмнө|\d+ өдриин өмнө|\d+ сарын өмнө|\d+ мин дараа|\d+ сек дараа|\d+ мин өмнө|\d+ сек өмнө|\d+ ц дараа|\d+ ц өмнө|\d+ өдөрт)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(арван хоердугаар сар|өнгөрсөн долоо хоног|арван нэгдүгээр сар|ирэх долоо хоног|гуравдугаар сар|дөрөвдүгээр сар|зургадугаар сар|энэ долоо хоног|аравдугаар сар|наимдугаар сар|хоердугаар сар|долдугаар сар|нэгдүгээр сар|тавдугаар сар|есдүгээр сар|өнгөрсөн жил|өнгөрсөн сар|долоо хоног|энэ минут|10-р сар|11-р сар|12-р сар|ирэх жил|ирэх сар|1-р сар|2-р сар|3-р сар|4-р сар|5-р сар|6-р сар|7-р сар|8-р сар|9-р сар|маргааш|энэ жил|энэ сар|энэ цаг|өнөөдөр|өчигдөр|баасан|лхагва|мягмар|секунд|бямба|даваа|минут|пүрэв|одоо|өдөр|gmt|utc|жил|мин|ням|сар|сек|цаг|7х|\+|\.|\[|\]|\||am|pm|ба|бя|да|лх|мя|ня|пү|үх|үө| |'|,|-|/|:|;|@|z|ц)((?:\z|\W|_|\d).*)$`),
})
