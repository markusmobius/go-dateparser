// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fi_Locale = merge(nil, LocaleData{
	Name:                  "fi",
	DateOrder:             "DMY",
	Charset:               []rune(`cdeghijklnorstuvyzä`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+) (sekunnin|sekuntin|minuutin|tunnin|paivan|viikon|kuukauden|vuoden) (paasta|kuluttua)(\z|[^\pL\pM\d]|_)`), "${1}${4} ${2} ${3}${5}"},
	},
	Translations: map[string]string{
		"keskiviikkona": "wednesday",
		"maaliskuussa":  "march",
		"marraskuussa":  "november",
		"heinakuussa":   "july",
		"helmikuussa":   "february",
		"huhtikuussa":   "april",
		"joulukuussa":   "december",
		"keskiviikko":   "wednesday",
		"maaliskuuta":   "march",
		"maanantaina":   "monday",
		"marraskuuta":   "november",
		"perjantaina":   "friday",
		"sunnuntaina":   "sunday",
		"tammikuussa":   "january",
		"toukokuussa":   "may",
		"heinakuuta":    "july",
		"helmikuuta":    "february",
		"huhtikuuta":    "april",
		"joulukuuta":    "december",
		"kesakuussa":    "june",
		"lauantaina":    "saturday",
		"lokakuussa":    "october",
		"syyskuussa":    "september",
		"tammikuuta":    "january",
		"toukokuuta":    "may",
		"elokuussa":     "august",
		"kesakuuta":     "june",
		"kuukauden":     "month",
		"kuukautta":     "month",
		"lokakuuta":     "october",
		"maaliskuu":     "march",
		"maanantai":     "monday",
		"marraskuu":     "november",
		"minuuttia":     "minute",
		"perjantai":     "friday",
		"sekunttia":     "second",
		"sunnuntai":     "sunday",
		"syyskuuta":     "september",
		"tiistaina":     "tuesday",
		"torstaina":     "thursday",
		"elokuuta":      "august",
		"heinakuu":      "july",
		"helmikuu":      "february",
		"huhtikuu":      "april",
		"joulukuu":      "december",
		"kuluttua":      "in",
		"kuukausi":      "month",
		"lauantai":      "saturday",
		"minuutin":      "minute",
		"minuutti":      "minute",
		"sekunnin":      "second",
		"sekuntia":      "second",
		"sekuntin":      "second",
		"sekuntti":      "second",
		"tammikuu":      "january",
		"toukokuu":      "may",
		"kesakuu":       "june",
		"lokakuu":       "october",
		"maalisk":       "march",
		"marrask":       "november",
		"sekunti":       "second",
		"syyskuu":       "september",
		"tiistai":       "tuesday",
		"torstai":       "thursday",
		"viikkoa":       "week",
		"elokuu":        "august",
		"heinak":        "july",
		"helmik":        "february",
		"huhtik":        "april",
		"jouluk":        "december",
		"maalis":        "march",
		"marras":        "november",
		"paasta":        "in",
		"paivaa":        "day",
		"paivan":        "day",
		"sitten":        "ago",
		"tammik":        "january",
		"toukok":        "may",
		"tunnin":        "hour",
		"tuntia":        "hour",
		"viikko":        "week",
		"viikon":        "week",
		"vuoden":        "year",
		"vuonna":        "year",
		"vuotta":        "year",
		"heina":         "july",
		"helmi":         "february",
		"huhti":         "april",
		"joulu":         "december",
		"kesak":         "june",
		"lokak":         "october",
		"paiva":         "day",
		"syysk":         "september",
		"tammi":         "january",
		"touko":         "may",
		"tunti":         "hour",
		"vuosi":         "year",
		"elok":          "august",
		"kesa":          "june",
		"loka":          "october",
		"pvaa":          "day",
		"syys":          "september",
		"elo":           "august",
		"gmt":           "gmt",
		"min":           "minute",
		"pva":           "day",
		"utc":           "utc",
		"vko":           "week",
		":n":            "",
		"am":            "am",
		"ap":            "am",
		"ip":            "pm",
		"ke":            "wednesday",
		"kk":            "month",
		"la":            "saturday",
		"ma":            "monday",
		"pe":            "friday",
		"pm":            "pm",
		"pv":            "day",
		"su":            "sunday",
		"ti":            "tuesday",
		"to":            "thursday",
		"vk":            "week",
		"vv":            "year",
		" ":             " ",
		"'":             "",
		"+":             "+",
		",":             "",
		"-":             "-",
		".":             ".",
		"/":             "/",
		":":             ":",
		";":             "",
		"@":             "",
		"[":             "",
		"]":             "",
		"p":             "day",
		"s":             "second",
		"t":             "hour",
		"v":             "year",
		"z":             "z",
		"|":             "",
	},
	RelativeType: map[string]string{
		"taman minuutin aikana": "0 minute ago",
		"taman tunnin aikana":   "0 hour ago",
		"minuutin sisalla":      "0 minute ago",
		"toissa viikolla":       "2 week ago",
		"talla viikolla":        "0 week ago",
		"toissa paivana":        "2 day ago",
		"tunnin sisalla":        "0 hour ago",
		"viime viikolla":        "1 week ago",
		"ensi viikolla":         "in 1 week",
		"toissa kuussa":         "2 month ago",
		"toissa vuonna":         "2 year ago",
		"tassa kuussa":          "0 month ago",
		"viime kuussa":          "1 month ago",
		"viime vuonna":          "1 year ago",
		"ensi kuussa":           "in 1 month",
		"ensi vuonna":           "in 1 year",
		"tana vuonna":           "0 year ago",
		"huomenna":              "in 1 day",
		"talla vk":              "0 week ago",
		"tassa kk":              "0 month ago",
		"viime kk":              "1 month ago",
		"viime vk":              "1 week ago",
		"ensi kk":               "in 1 month",
		"ensi vk":               "in 1 week",
		"viime v":               "1 year ago",
		"ensi v":                "in 1 year",
		"tana v":                "0 year ago",
		"tanaan":                "0 day ago",
		"eilen":                 "1 day ago",
		"huom":                  "in 1 day",
		"nyt":                   "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) kuukauden paasta`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) kuukautta sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) minuuttia sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) kuukausi sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) minuutin paasta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) minuutti sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sekunnin paasta`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) sekuntia sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) sekunti sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) viikkoa sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) paivaa sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) paivan paasta`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) tunnin paasta`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) tuntia sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) viikko sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) viikon paasta`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) vuoden paasta`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) vuotta sitten`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) paiva sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) tunti sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) vuosi sitten`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) min paasta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) min sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) kk paasta`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) kk sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) pv paasta`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) pv sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) vk paasta`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) vk sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) s paasta`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) s sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) t paasta`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) t sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) v paasta`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) v sitten`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ kuukauden paasta|\d+ kuukautta sitten|\d+ minuuttia sitten|\d+ kuukausi sitten|\d+ minuutin paasta|\d+ minuutti sitten|\d+ sekunnin paasta|\d+ sekuntia sitten|\d+ sekunti sitten|\d+ viikkoa sitten|\d+ paivaa sitten|\d+ paivan paasta|\d+ tunnin paasta|\d+ tuntia sitten|\d+ viikko sitten|\d+ viikon paasta|\d+ vuoden paasta|\d+ vuotta sitten|\d+ paiva sitten|\d+ tunti sitten|\d+ vuosi sitten|\d+ min paasta|\d+ min sitten|\d+ kk paasta|\d+ kk sitten|\d+ pv paasta|\d+ pv sitten|\d+ vk paasta|\d+ vk sitten|\d+ s paasta|\d+ s sitten|\d+ t paasta|\d+ t sitten|\d+ v paasta|\d+ v sitten)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ kuukauden paasta|\d+ kuukautta sitten|\d+ minuuttia sitten|\d+ kuukausi sitten|\d+ minuutin paasta|\d+ minuutti sitten|\d+ sekunnin paasta|\d+ sekuntia sitten|\d+ sekunti sitten|\d+ viikkoa sitten|\d+ paivaa sitten|\d+ paivan paasta|\d+ tunnin paasta|\d+ tuntia sitten|\d+ viikko sitten|\d+ viikon paasta|\d+ vuoden paasta|\d+ vuotta sitten|\d+ paiva sitten|\d+ tunti sitten|\d+ vuosi sitten|\d+ min paasta|\d+ min sitten|\d+ kk paasta|\d+ kk sitten|\d+ pv paasta|\d+ pv sitten|\d+ vk paasta|\d+ vk sitten|\d+ s paasta|\d+ s sitten|\d+ t paasta|\d+ t sitten|\d+ v paasta|\d+ v sitten)$`),
	KnownWords:      []string{"taman minuutin aikana", "taman tunnin aikana", "minuutin sisalla", "toissa viikolla", "talla viikolla", "toissa paivana", "tunnin sisalla", "viime viikolla", "ensi viikolla", "keskiviikkona", "toissa kuussa", "toissa vuonna", "maaliskuussa", "marraskuussa", "tassa kuussa", "viime kuussa", "viime vuonna", "ensi kuussa", "ensi vuonna", "heinakuussa", "helmikuussa", "huhtikuussa", "joulukuussa", "keskiviikko", "maaliskuuta", "maanantaina", "marraskuuta", "perjantaina", "sunnuntaina", "tammikuussa", "tana vuonna", "toukokuussa", "heinakuuta", "helmikuuta", "huhtikuuta", "joulukuuta", "kesakuussa", "lauantaina", "lokakuussa", "syyskuussa", "tammikuuta", "toukokuuta", "elokuussa", "kesakuuta", "kuukauden", "kuukautta", "lokakuuta", "maaliskuu", "maanantai", "marraskuu", "minuuttia", "perjantai", "sekunttia", "sunnuntai", "syyskuuta", "tiistaina", "torstaina", "elokuuta", "heinakuu", "helmikuu", "huhtikuu", "huomenna", "joulukuu", "kuluttua", "kuukausi", "lauantai", "minuutin", "minuutti", "sekunnin", "sekuntia", "sekuntin", "sekuntti", "talla vk", "tammikuu", "tassa kk", "toukokuu", "viime kk", "viime vk", "ensi kk", "ensi vk", "kesakuu", "lokakuu", "maalisk", "marrask", "sekunti", "syyskuu", "tiistai", "torstai", "viikkoa", "viime v", "elokuu", "ensi v", "heinak", "helmik", "huhtik", "jouluk", "maalis", "marras", "paasta", "paivaa", "paivan", "sitten", "tammik", "tana v", "tanaan", "toukok", "tunnin", "tuntia", "viikko", "viikon", "vuoden", "vuonna", "vuotta", "eilen", "heina", "helmi", "huhti", "joulu", "kesak", "lokak", "paiva", "syysk", "tammi", "touko", "tunti", "vuosi", "elok", "huom", "kesa", "loka", "pvaa", "syys", "elo", "gmt", "min", "nyt", "pva", "utc", "vko", ":n", "am", "ap", "ip", "ke", "kk", "la", "ma", "pe", "pm", "pv", "su", "ti", "to", "vk", "vv", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "p", "s", "t", "v", "z", "|"},
})
