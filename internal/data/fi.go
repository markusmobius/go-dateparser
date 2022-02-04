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
		"huhtikuussa":   "april",
		"joulukuussa":   "december",
		"helmikuussa":   "february",
		"perjantaina":   "friday",
		"tammikuussa":   "january",
		"heinakuussa":   "july",
		"maaliskuuta":   "march",
		"toukokuussa":   "may",
		"maanantaina":   "monday",
		"marraskuuta":   "november",
		"sunnuntaina":   "sunday",
		"keskiviikko":   "wednesday",
		"huhtikuuta":    "april",
		"joulukuuta":    "december",
		"helmikuuta":    "february",
		"tammikuuta":    "january",
		"heinakuuta":    "july",
		"kesakuussa":    "june",
		"toukokuuta":    "may",
		"lokakuussa":    "october",
		"lauantaina":    "saturday",
		"syyskuussa":    "september",
		"elokuussa":     "august",
		"perjantai":     "friday",
		"kesakuuta":     "june",
		"maaliskuu":     "march",
		"minuuttia":     "minute",
		"maanantai":     "monday",
		"kuukauden":     "month",
		"kuukautta":     "month",
		"marraskuu":     "november",
		"lokakuuta":     "october",
		"sekunttia":     "second",
		"syyskuuta":     "september",
		"sunnuntai":     "sunday",
		"torstaina":     "thursday",
		"tiistaina":     "tuesday",
		"huhtikuu":      "april",
		"elokuuta":      "august",
		"joulukuu":      "december",
		"helmikuu":      "february",
		"kuluttua":      "in",
		"tammikuu":      "january",
		"heinakuu":      "july",
		"toukokuu":      "may",
		"minuutin":      "minute",
		"minuutti":      "minute",
		"kuukausi":      "month",
		"lauantai":      "saturday",
		"sekunnin":      "second",
		"sekuntia":      "second",
		"sekuntin":      "second",
		"sekuntti":      "second",
		"kesakuu":       "june",
		"maalisk":       "march",
		"marrask":       "november",
		"lokakuu":       "october",
		"sekunti":       "second",
		"syyskuu":       "september",
		"torstai":       "thursday",
		"tiistai":       "tuesday",
		"viikkoa":       "week",
		"sitten":        "ago",
		"huhtik":        "april",
		"elokuu":        "august",
		"paivaa":        "day",
		"paivan":        "day",
		"jouluk":        "december",
		"helmik":        "february",
		"tunnin":        "hour",
		"tuntia":        "hour",
		"paasta":        "in",
		"tammik":        "january",
		"heinak":        "july",
		"maalis":        "march",
		"toukok":        "may",
		"marras":        "november",
		"viikko":        "week",
		"viikon":        "week",
		"vuoden":        "year",
		"vuonna":        "year",
		"vuotta":        "year",
		"huhti":         "april",
		"paiva":         "day",
		"joulu":         "december",
		"helmi":         "february",
		"tunti":         "hour",
		"tammi":         "january",
		"heina":         "july",
		"kesak":         "june",
		"touko":         "may",
		"lokak":         "october",
		"syysk":         "september",
		"vuosi":         "year",
		"elok":          "august",
		"pvaa":          "day",
		"kesa":          "june",
		"loka":          "october",
		"syys":          "september",
		"elo":           "august",
		"pva":           "day",
		"gmt":           "gmt",
		"min":           "minute",
		"utc":           "utc",
		"vko":           "week",
		":n":            "",
		"am":            "am",
		"ap":            "am",
		"pv":            "day",
		"pe":            "friday",
		"ma":            "monday",
		"kk":            "month",
		"ip":            "pm",
		"pm":            "pm",
		"la":            "saturday",
		"su":            "sunday",
		"to":            "thursday",
		"ti":            "tuesday",
		"ke":            "wednesday",
		"vk":            "week",
		"vv":            "year",
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
		"p":             "day",
		"t":             "hour",
		"s":             "second",
		"v":             "year",
		"z":             "z",
	},
	RelativeType: map[string]string{
		"taman minuutin aikana": "0 minute ago",
		"taman tunnin aikana":   "0 hour ago",
		"minuutin sisalla":      "0 minute ago",
		"toissa viikolla":       "2 week ago",
		"tunnin sisalla":        "0 hour ago",
		"talla viikolla":        "0 week ago",
		"viime viikolla":        "1 week ago",
		"toissa paivana":        "2 day ago",
		"toissa kuussa":         "2 month ago",
		"toissa vuonna":         "2 year ago",
		"ensi viikolla":         "in 1 week",
		"tassa kuussa":          "0 month ago",
		"viime kuussa":          "1 month ago",
		"viime vuonna":          "1 year ago",
		"tana vuonna":           "0 year ago",
		"ensi kuussa":           "in 1 month",
		"ensi vuonna":           "in 1 year",
		"tassa kk":              "0 month ago",
		"talla vk":              "0 week ago",
		"viime kk":              "1 month ago",
		"viime vk":              "1 week ago",
		"huomenna":              "in 1 day",
		"viime v":               "1 year ago",
		"ensi kk":               "in 1 month",
		"ensi vk":               "in 1 week",
		"tanaan":                "0 day ago",
		"tana v":                "0 year ago",
		"ensi v":                "in 1 year",
		"eilen":                 "1 day ago",
		"huom":                  "in 1 day",
		"nyt":                   "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) minuuttia sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) kuukautta sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) kuukauden paasta`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) minuutti sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) kuukausi sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) sekuntia sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) minuutin paasta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) sekunnin paasta`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) sekunti sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) viikkoa sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) paivaa sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) tuntia sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) viikko sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) vuotta sitten`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) paivan paasta`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) tunnin paasta`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) viikon paasta`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) vuoden paasta`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) paiva sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) tunti sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) vuosi sitten`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) min sitten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) min paasta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) pv sitten`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) kk sitten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) vk sitten`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) pv paasta`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) kk paasta`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) vk paasta`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) t sitten`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) s sitten`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) v sitten`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) t paasta`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) s paasta`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) v paasta`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ kuukauden paasta|\d+ kuukautta sitten|\d+ minuuttia sitten|\d+ kuukausi sitten|\d+ minuutin paasta|\d+ minuutti sitten|\d+ sekunnin paasta|\d+ sekuntia sitten|\d+ sekunti sitten|\d+ viikkoa sitten|\d+ paivaa sitten|\d+ paivan paasta|\d+ tunnin paasta|\d+ tuntia sitten|\d+ viikko sitten|\d+ viikon paasta|\d+ vuoden paasta|\d+ vuotta sitten|\d+ paiva sitten|\d+ tunti sitten|\d+ vuosi sitten|\d+ min paasta|\d+ min sitten|\d+ kk paasta|\d+ kk sitten|\d+ pv paasta|\d+ pv sitten|\d+ vk paasta|\d+ vk sitten|\d+ s paasta|\d+ s sitten|\d+ t paasta|\d+ t sitten|\d+ v paasta|\d+ v sitten)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ kuukauden paasta|\d+ kuukautta sitten|\d+ minuuttia sitten|\d+ kuukausi sitten|\d+ minuutin paasta|\d+ minuutti sitten|\d+ sekunnin paasta|\d+ sekuntia sitten|\d+ sekunti sitten|\d+ viikkoa sitten|\d+ paivaa sitten|\d+ paivan paasta|\d+ tunnin paasta|\d+ tuntia sitten|\d+ viikko sitten|\d+ viikon paasta|\d+ vuoden paasta|\d+ vuotta sitten|\d+ paiva sitten|\d+ tunti sitten|\d+ vuosi sitten|\d+ min paasta|\d+ min sitten|\d+ kk paasta|\d+ kk sitten|\d+ pv paasta|\d+ pv sitten|\d+ vk paasta|\d+ vk sitten|\d+ s paasta|\d+ s sitten|\d+ t paasta|\d+ t sitten|\d+ v paasta|\d+ v sitten)$`),
	KnownWords:      []string{"taman minuutin aikana", "taman tunnin aikana", "minuutin sisalla", "toissa viikolla", "talla viikolla", "toissa paivana", "tunnin sisalla", "viime viikolla", "ensi viikolla", "keskiviikkona", "toissa kuussa", "toissa vuonna", "maaliskuussa", "marraskuussa", "tassa kuussa", "viime kuussa", "viime vuonna", "ensi kuussa", "ensi vuonna", "heinakuussa", "helmikuussa", "huhtikuussa", "joulukuussa", "keskiviikko", "maaliskuuta", "maanantaina", "marraskuuta", "perjantaina", "sunnuntaina", "tammikuussa", "tana vuonna", "toukokuussa", "heinakuuta", "helmikuuta", "huhtikuuta", "joulukuuta", "kesakuussa", "lauantaina", "lokakuussa", "syyskuussa", "tammikuuta", "toukokuuta", "elokuussa", "kesakuuta", "kuukauden", "kuukautta", "lokakuuta", "maaliskuu", "maanantai", "marraskuu", "minuuttia", "perjantai", "sekunttia", "sunnuntai", "syyskuuta", "tiistaina", "torstaina", "elokuuta", "heinakuu", "helmikuu", "huhtikuu", "huomenna", "joulukuu", "kuluttua", "kuukausi", "lauantai", "minuutin", "minuutti", "sekunnin", "sekuntia", "sekuntin", "sekuntti", "talla vk", "tammikuu", "tassa kk", "toukokuu", "viime kk", "viime vk", "ensi kk", "ensi vk", "kesakuu", "lokakuu", "maalisk", "marrask", "sekunti", "syyskuu", "tiistai", "torstai", "viikkoa", "viime v", "elokuu", "ensi v", "heinak", "helmik", "huhtik", "jouluk", "maalis", "marras", "paasta", "paivaa", "paivan", "sitten", "tammik", "tana v", "tanaan", "toukok", "tunnin", "tuntia", "viikko", "viikon", "vuoden", "vuonna", "vuotta", "eilen", "heina", "helmi", "huhti", "joulu", "kesak", "lokak", "paiva", "syysk", "tammi", "touko", "tunti", "vuosi", "elok", "huom", "kesa", "loka", "pvaa", "syys", "elo", "gmt", "min", "nyt", "pva", "utc", "vko", ":n", "am", "ap", "ip", "ke", "kk", "la", "ma", "pe", "pm", "pv", "su", "ti", "to", "vk", "vv", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "p", "s", "t", "v", "z", "|"},
})
