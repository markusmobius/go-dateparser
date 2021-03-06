// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ms_Locale = merge(nil, LocaleData{
	Name:      "ms",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvz`),
	Translations: map[string]string{
		"september": "september",
		"disember":  "december",
		"februari":  "february",
		"november":  "november",
		"januari":   "january",
		"oktober":   "october",
		"jumaat":    "friday",
		"khamis":    "thursday",
		"minggu":    "week",
		"selasa":    "tuesday",
		"april":     "april",
		"bulan":     "month",
		"isnin":     "monday",
		"julai":     "july",
		"minit":     "minute",
		"sabtu":     "saturday",
		"tahun":     "year",
		"ahad":      "sunday",
		"hari":      "day",
		"ogos":      "august",
		"rabu":      "wednesday",
		"saat":      "second",
		"ahd":       "sunday",
		"apr":       "april",
		"bln":       "month",
		"dis":       "december",
		"feb":       "february",
		"gmt":       "gmt",
		"isn":       "monday",
		"jam":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jum":       "friday",
		"jun":       "june",
		"kha":       "thursday",
		"mac":       "march",
		"mei":       "may",
		"mgu":       "week",
		"min":       "minute",
		"nov":       "november",
		"ogo":       "august",
		"okt":       "october",
		"ptg":       "pm",
		"rab":       "wednesday",
		"sab":       "saturday",
		"sel":       "tuesday",
		"sep":       "september",
		"thn":       "year",
		"utc":       "utc",
		"am":        "am",
		"pg":        "am",
		"pm":        "pm",
		" ":         " ",
		"'":         "",
		"+":         "+",
		",":         "",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"z":         "z",
		"|":         "",
	},
	RelativeType: map[string]string{
		"pada minit ini": "0 minute ago",
		"minggu depan":   "in 1 week",
		"bulan depan":    "in 1 month",
		"minggu lalu":    "1 week ago",
		"tahun depan":    "in 1 year",
		"bulan lalu":     "1 month ago",
		"minggu ini":     "0 week ago",
		"tahun lalu":     "1 year ago",
		"bln depan":      "in 1 month",
		"bulan ini":      "0 month ago",
		"mng depan":      "in 1 week",
		"mng lepas":      "1 week ago",
		"tahun ini":      "0 year ago",
		"thn depan":      "in 1 year",
		"thn lepas":      "1 year ago",
		"bln lalu":       "1 month ago",
		"hari ini":       "0 day ago",
		"sekarang":       "0 second ago",
		"bln ini":        "0 month ago",
		"jam ini":        "0 hour ago",
		"mng ini":        "0 week ago",
		"semalam":        "1 day ago",
		"thn ini":        "0 year ago",
		"semlm":          "1 day ago",
		"esok":           "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)dalam (\d+) minggu`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) minggu lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) bulan`), "in $1 month"},
		{regexp.MustCompile(`(?i)dalam (\d+) minit`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) bulan lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) minit lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) tahun lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) hari`), "in $1 day"},
		{regexp.MustCompile(`(?i)dalam (\d+) saat`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) hari lalu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) saat lalu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) jam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dalam (\d+) thn`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) bln lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) jam lalu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) mgu lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) min lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) thn lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dlm (\d+) hari`), "in $1 day"},
		{regexp.MustCompile(`(?i)dlm (\d+) saat`), "in $1 second"},
		{regexp.MustCompile(`(?i)dlm (\d+) bln`), "in $1 month"},
		{regexp.MustCompile(`(?i)dlm (\d+) jam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dlm (\d+) mgu`), "in $1 week"},
		{regexp.MustCompile(`(?i)dlm (\d+) min`), "in $1 minute"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(dalam \d+ minggu|\d+ minggu lalu|dalam \d+ bulan|dalam \d+ minit|\d+ bulan lalu|\d+ minit lalu|\d+ tahun lalu|dalam \d+ hari|dalam \d+ saat|\d+ hari lalu|\d+ saat lalu|dalam \d+ jam|dalam \d+ thn|\d+ bln lalu|\d+ jam lalu|\d+ mgu lalu|\d+ min lalu|\d+ thn lalu|dlm \d+ hari|dlm \d+ saat|dlm \d+ bln|dlm \d+ jam|dlm \d+ mgu|dlm \d+ min)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dalam \d+ minggu|\d+ minggu lalu|dalam \d+ bulan|dalam \d+ minit|\d+ bulan lalu|\d+ minit lalu|\d+ tahun lalu|dalam \d+ hari|dalam \d+ saat|\d+ hari lalu|\d+ saat lalu|dalam \d+ jam|dalam \d+ thn|\d+ bln lalu|\d+ jam lalu|\d+ mgu lalu|\d+ min lalu|\d+ thn lalu|dlm \d+ hari|dlm \d+ saat|dlm \d+ bln|dlm \d+ jam|dlm \d+ mgu|dlm \d+ min)$`),
	KnownWords:      []string{"pada minit ini", "minggu depan", "bulan depan", "minggu lalu", "tahun depan", "bulan lalu", "minggu ini", "tahun lalu", "bln depan", "bulan ini", "mng depan", "mng lepas", "september", "tahun ini", "thn depan", "thn lepas", "bln lalu", "disember", "februari", "hari ini", "november", "sekarang", "bln ini", "jam ini", "januari", "mng ini", "oktober", "semalam", "thn ini", "jumaat", "khamis", "minggu", "selasa", "april", "bulan", "isnin", "julai", "minit", "sabtu", "semlm", "tahun", "ahad", "esok", "hari", "ogos", "rabu", "saat", "ahd", "apr", "bln", "dis", "feb", "gmt", "isn", "jam", "jan", "jul", "jum", "jun", "kha", "mac", "mei", "mgu", "min", "nov", "ogo", "okt", "ptg", "rab", "sab", "sel", "sep", "thn", "utc", "am", "pg", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})

var ms_BN_Locale = merge(&ms_Locale, LocaleData{
	Name:      "ms-BN",
	DateOrder: "DMY",
})

var ms_SG_Locale = merge(&ms_Locale, LocaleData{
	Name:      "ms-SG",
	DateOrder: "DMY",
})
