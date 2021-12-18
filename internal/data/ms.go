// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ms_Locale = merge(nil, LocaleData{
	Name:      "ms",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"pada minit ini": "0 minute ago",
		"minggu depan":   "in 1 week",
		"minggu lalu":    "1 week ago",
		"bulan depan":    "in 1 month",
		"tahun depan":    "in 1 year",
		"minggu ini":     "0 week ago",
		"bulan lalu":     "1 month ago",
		"tahun lalu":     "1 year ago",
		"bulan ini":      "0 month ago",
		"tahun ini":      "0 year ago",
		"mng lepas":      "1 week ago",
		"thn lepas":      "1 year ago",
		"bln depan":      "in 1 month",
		"mng depan":      "in 1 week",
		"thn depan":      "in 1 year",
		"september":      "september",
		"hari ini":       "0 day ago",
		"sekarang":       "0 second ago",
		"bln lalu":       "1 month ago",
		"disember":       "december",
		"februari":       "february",
		"november":       "november",
		"jam ini":        "0 hour ago",
		"bln ini":        "0 month ago",
		"mng ini":        "0 week ago",
		"thn ini":        "0 year ago",
		"semalam":        "1 day ago",
		"januari":        "january",
		"oktober":        "october",
		"jumaat":         "friday",
		"khamis":         "thursday",
		"selasa":         "tuesday",
		"minggu":         "week",
		"semlm":          "1 day ago",
		"april":          "april",
		"julai":          "july",
		"minit":          "minute",
		"isnin":          "monday",
		"bulan":          "month",
		"sabtu":          "saturday",
		"tahun":          "year",
		"ogos":           "august",
		"hari":           "day",
		"esok":           "in 1 day",
		"saat":           "second",
		"ahad":           "sunday",
		"rabu":           "wednesday",
		"apr":            "april",
		"ogo":            "august",
		"dis":            "december",
		"feb":            "february",
		"jum":            "friday",
		"gmt":            "gmt",
		"jam":            "hour",
		"jan":            "january",
		"jul":            "july",
		"jun":            "june",
		"mac":            "march",
		"mei":            "may",
		"min":            "minute",
		"isn":            "monday",
		"bln":            "month",
		"nov":            "november",
		"okt":            "october",
		"ptg":            "pm",
		"sab":            "saturday",
		"sep":            "september",
		"ahd":            "sunday",
		"kha":            "thursday",
		"sel":            "tuesday",
		"utc":            "utc",
		"rab":            "wednesday",
		"mgu":            "week",
		"thn":            "year",
		"am":             "am",
		"pg":             "am",
		"pm":             "pm",
		"'":              "",
		",":              "",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"|":              "",
		" ":              " ",
		"+":              "+",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		"z":              "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)dalam (\d+) minggu`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) minggu lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) minit`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dalam (\d+) bulan`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) minit lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bulan lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) tahun lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) hari`), "in $1 day"},
		{regexp.MustCompile(`(?i)dalam (\d+) saat`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) hari lalu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) saat lalu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) jam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dalam (\d+) thn`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) jam lalu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) min lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bln lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) mgu lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) thn lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dlm (\d+) hari`), "in $1 day"},
		{regexp.MustCompile(`(?i)dlm (\d+) saat`), "in $1 second"},
		{regexp.MustCompile(`(?i)dlm (\d+) jam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dlm (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dlm (\d+) bln`), "in $1 month"},
		{regexp.MustCompile(`(?i)dlm (\d+) mgu`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(dalam \d+ minggu|\d+ minggu lalu|dalam \d+ bulan|dalam \d+ minit|\d+ bulan lalu|\d+ minit lalu|\d+ tahun lalu|dalam \d+ hari|dalam \d+ saat|\d+ hari lalu|\d+ saat lalu|dalam \d+ jam|dalam \d+ thn|\d+ bln lalu|\d+ jam lalu|\d+ mgu lalu|\d+ min lalu|\d+ thn lalu|dlm \d+ hari|dlm \d+ saat|dlm \d+ bln|dlm \d+ jam|dlm \d+ mgu|dlm \d+ min)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dalam \d+ minggu|\d+ minggu lalu|dalam \d+ bulan|dalam \d+ minit|\d+ bulan lalu|\d+ minit lalu|\d+ tahun lalu|dalam \d+ hari|dalam \d+ saat|\d+ hari lalu|\d+ saat lalu|dalam \d+ jam|dalam \d+ thn|\d+ bln lalu|\d+ jam lalu|\d+ mgu lalu|\d+ min lalu|\d+ thn lalu|dlm \d+ hari|dlm \d+ saat|dlm \d+ bln|dlm \d+ jam|dlm \d+ mgu|dlm \d+ min)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(pada minit ini|minggu depan|bulan depan|minggu lalu|tahun depan|bulan lalu|minggu ini|tahun lalu|bln depan|bulan ini|mng depan|mng lepas|september|tahun ini|thn depan|thn lepas|bln lalu|disember|februari|hari ini|november|sekarang|bln ini|jam ini|januari|mng ini|oktober|semalam|thn ini|jumaat|khamis|minggu|selasa|april|bulan|isnin|julai|minit|sabtu|semlm|tahun|ahad|esok|hari|ogos|rabu|saat|ahd|apr|bln|dis|feb|gmt|isn|jam|jan|jul|jum|jun|kha|mac|mei|mgu|min|nov|ogo|okt|ptg|rab|sab|sel|sep|thn|utc|\+|\.|\[|\]|\||am|pg|pm| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})

var ms_BN_Locale = merge(&ms_Locale, LocaleData{
	Name:            "ms-BN",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)()(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^()$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))()((?:\z|\W|_|\d).*)$`),
})

var ms_SG_Locale = merge(&ms_Locale, LocaleData{
	Name:            "ms-SG",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)()(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^()$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))()((?:\z|\W|_|\d).*)$`),
})
