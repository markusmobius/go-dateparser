// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var is_Locale = merge(nil, LocaleData{
	Name:      "is",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvzáæíðóöúþ`),
	Translations: map[string]string{
		"miðvikudagur": "wednesday",
		"klukkustund":  "hour",
		"laugardagur":  "saturday",
		"fimmtudagur":  "thursday",
		"þriðjudagur":  "tuesday",
		"fostudagur":   "friday",
		"sunnudagur":   "sunday",
		"manudagur":    "monday",
		"september":    "september",
		"desember":     "december",
		"november":     "november",
		"februar":      "february",
		"manuður":      "month",
		"oktober":      "october",
		"sekunda":      "second",
		"januar":       "january",
		"minuta":       "minute",
		"april":        "april",
		"agust":        "august",
		"dagur":        "day",
		"klst":         "hour",
		"juli":         "july",
		"juni":         "june",
		"mars":         "march",
		"vika":         "week",
		"apr":          "april",
		"agu":          "august",
		"des":          "december",
		"feb":          "february",
		"fos":          "friday",
		"gmt":          "gmt",
		"jan":          "january",
		"jul":          "july",
		"jun":          "june",
		"mar":          "march",
		"mai":          "may",
		"min":          "minute",
		"man":          "month",
		"nov":          "november",
		"okt":          "october",
		"lau":          "saturday",
		"sek":          "second",
		"sep":          "september",
		"sun":          "sunday",
		"fim":          "thursday",
		"þri":          "tuesday",
		"utc":          "utc",
		"mið":          "wednesday",
		"am":           "am",
		"fh":           "am",
		"eh":           "pm",
		"pm":           "pm",
		"ar":           "year",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"d":            "day",
		"v":            "week",
		"z":            "z",
	},
	RelativeType: map[string]string{
		"i siðasta manuði": "1 month ago",
		"i þessum manuði":  "0 month ago",
		"i þessari viku":   "0 week ago",
		"i siðustu viku":   "1 week ago",
		"i næsta manuði":   "in 1 month",
		"i siðasta man":    "1 month ago",
		"a siðasta ari":    "1 year ago",
		"i þessum man":     "0 month ago",
		"i næstu viku":     "in 1 week",
		"this minute":      "0 minute ago",
		"a þessu ari":      "0 year ago",
		"i næsta man":      "in 1 month",
		"a næsta ari":      "in 1 year",
		"this hour":        "0 hour ago",
		"a morgun":         "in 1 day",
		"i dag":            "0 day ago",
		"i gær":            "1 day ago",
		"nuna":             "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)fyrir (\d+) klukkustundum`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) klukkustundir`), "in $1 hour"},
		{regexp.MustCompile(`(?i)fyrir (\d+) klukkustund`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) klukkustund`), "in $1 hour"},
		{regexp.MustCompile(`(?i)fyrir (\d+) sekundum`), "$1 second ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) sekundur`), "in $1 second"},
		{regexp.MustCompile(`(?i)fyrir (\d+) minutum`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) manuðum`), "$1 month ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) sekundu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) minutur`), "in $1 minute"},
		{regexp.MustCompile(`(?i)eftir (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)fyrir (\d+) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) manuði`), "$1 month ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)eftir (\d+) manuði`), "in $1 month"},
		{regexp.MustCompile(`(?i)fyrir (\d+) dogum`), "$1 day ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) vikum`), "$1 week ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) manuð`), "in $1 month"},
		{regexp.MustCompile(`(?i)eftir (\d+) vikur`), "in $1 week"},
		{regexp.MustCompile(`(?i)fyrir (\d+) degi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) klst`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) viku`), "$1 week ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) arum`), "$1 year ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) daga`), "in $1 day"},
		{regexp.MustCompile(`(?i)eftir (\d+) klst`), "in $1 hour"},
		{regexp.MustCompile(`(?i)eftir (\d+) viku`), "in $1 week"},
		{regexp.MustCompile(`(?i)fyrir (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) man`), "$1 month ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)fyrir (\d+) ari`), "$1 year ago"},
		{regexp.MustCompile(`(?i)eftir (\d+) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)eftir (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)eftir (\d+) man`), "in $1 month"},
		{regexp.MustCompile(`(?i)eftir (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)eftir (\d+) ar`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(eftir \d+ klukkustundir|fyrir \d+ klukkustundum|eftir \d+ klukkustund|fyrir \d+ klukkustund|eftir \d+ sekundur|fyrir \d+ sekundum|eftir \d+ minutur|eftir \d+ sekundu|fyrir \d+ manuðum|fyrir \d+ minutum|fyrir \d+ sekundu|eftir \d+ manuði|eftir \d+ minutu|fyrir \d+ manuði|fyrir \d+ minutu|eftir \d+ manuð|eftir \d+ vikur|fyrir \d+ dogum|fyrir \d+ vikum|eftir \d+ daga|eftir \d+ klst|eftir \d+ viku|fyrir \d+ arum|fyrir \d+ degi|fyrir \d+ klst|fyrir \d+ viku|eftir \d+ dag|eftir \d+ man|eftir \d+ min|eftir \d+ sek|fyrir \d+ ari|fyrir \d+ man|fyrir \d+ min|fyrir \d+ sek|eftir \d+ ar)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(eftir \d+ klukkustundir|fyrir \d+ klukkustundum|eftir \d+ klukkustund|fyrir \d+ klukkustund|eftir \d+ sekundur|fyrir \d+ sekundum|eftir \d+ minutur|eftir \d+ sekundu|fyrir \d+ manuðum|fyrir \d+ minutum|fyrir \d+ sekundu|eftir \d+ manuði|eftir \d+ minutu|fyrir \d+ manuði|fyrir \d+ minutu|eftir \d+ manuð|eftir \d+ vikur|fyrir \d+ dogum|fyrir \d+ vikum|eftir \d+ daga|eftir \d+ klst|eftir \d+ viku|fyrir \d+ arum|fyrir \d+ degi|fyrir \d+ klst|fyrir \d+ viku|eftir \d+ dag|eftir \d+ man|eftir \d+ min|eftir \d+ sek|fyrir \d+ ari|fyrir \d+ man|fyrir \d+ min|fyrir \d+ sek|eftir \d+ ar)$`),
	KnownWords:      []string{"i siðasta manuði", "i þessum manuði", "i næsta manuði", "i siðustu viku", "i þessari viku", "a siðasta ari", "i siðasta man", "i næstu viku", "i þessum man", "miðvikudagur", "a næsta ari", "a þessu ari", "fimmtudagur", "i næsta man", "klukkustund", "laugardagur", "this minute", "þriðjudagur", "fostudagur", "sunnudagur", "manudagur", "september", "this hour", "a morgun", "desember", "november", "februar", "manuður", "oktober", "sekunda", "januar", "minuta", "agust", "april", "dagur", "i dag", "i gær", "juli", "juni", "klst", "mars", "nuna", "vika", "agu", "apr", "des", "feb", "fim", "fos", "gmt", "jan", "jul", "jun", "lau", "mai", "man", "mar", "min", "mið", "nov", "okt", "sek", "sep", "sun", "utc", "þri", "am", "ar", "eh", "fh", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "v", "z", "|"},
})
