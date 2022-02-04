// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var vi_Locale = merge(nil, LocaleData{
	Name:                  "vi",
	DateOrder:             "DMY",
	Charset:               []rune(`bceghilnorstuyzàáâêìíôúăđưạảầậểồổộớờủứ`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(?:ngay|nam)\s(\d+)(\z|[^\pL\pM\d]|_)`), "${1}${2}${3}"},
	},
	Translations: map[string]string{
		"thang muoi hai": "december",
		"thang muoi mot": "november",
		"giay đong ho":   "second",
		"nguyen ban":     "minute",
		"thang muoi":     "october",
		"thang chin":     "september",
		"truoc đay":      "ago",
		"thang tam":      "august",
		"thang hai":      "february",
		"thang mot":      "january",
		"thang bay":      "july",
		"thang sau":      "june",
		"thang nam":      "may",
		"cach đay":       "ago",
		"thang tu":       "april",
		"ban ngay":       "day",
		"thang 12":       "december",
		"thang ba":       "march",
		"thang 11":       "november",
		"thang 10":       "october",
		"hang nhi":       "second",
		"chu nhat":       "sunday",
		"thang 4":        "april",
		"thang 8":        "august",
		"thang 2":        "february",
		"thu sau":        "friday",
		"thang 1":        "january",
		"thang 7":        "july",
		"thang 6":        "june",
		"thang 3":        "march",
		"thang 5":        "may",
		"thu hai":        "monday",
		"thu bay":        "saturday",
		"thang 9":        "september",
		"thu nam":        "thursday",
		"tuan le":        "week",
		"thg 12":         "december",
		"thg 11":         "november",
		"thg 10":         "october",
		"thu ba":         "tuesday",
		"thu tu":         "wednesday",
		"truoc":          "ago",
		"thg 4":          "april",
		"thg 8":          "august",
		"thg 2":          "february",
		"thu 6":          "friday",
		"trong":          "in",
		"thg 1":          "january",
		"thg 7":          "july",
		"thg 6":          "june",
		"thg 3":          "march",
		"thg 5":          "may",
		"thu 2":          "monday",
		"thang":          "month",
		"thu 7":          "saturday",
		"thg 9":          "september",
		"thu 1":          "sunday",
		"thu 5":          "thursday",
		"thu 3":          "tuesday",
		"thu 4":          "wednesday",
		"buoi":           "day",
		"ngay":           "day",
		"th 6":           "friday",
		"chut":           "minute",
		"phut":           "minute",
		"th 2":           "monday",
		"th 7":           "saturday",
		"giay":           "second",
		"th 5":           "thursday",
		"th 3":           "tuesday",
		"th 4":           "wednesday",
		"tuan":           "week",
		"luc":            "",
		"gmt":            "gmt",
		"gio":            "hour",
		"lat":            "minute",
		"thg":            "month",
		"utc":            "utc",
		"nam":            "year",
		"am":             "am",
		"sa":             "am",
		"ch":             "pm",
		"pm":             "pm",
		"cn":             "sunday",
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
	RelativeType: map[string]string{
		"thang truoc": "1 month ago",
		"tuan truoc":  "1 week ago",
		"thang nay":   "0 month ago",
		"nam ngoai":   "1 year ago",
		"thang sau":   "in 1 month",
		"phut nay":    "0 minute ago",
		"tuan nay":    "0 week ago",
		"ngay mai":    "in 1 day",
		"tuan sau":    "in 1 week",
		"hom nay":     "0 day ago",
		"gio nay":     "0 hour ago",
		"bay gio":     "0 second ago",
		"nam nay":     "0 year ago",
		"hom qua":     "1 day ago",
		"nam sau":     "in 1 year",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)sau (\d+) thang nua`), "in $1 month"},
		{regexp.MustCompile(`(?i)sau (\d+) ngay nua`), "in $1 day"},
		{regexp.MustCompile(`(?i)sau (\d+) phut nua`), "in $1 minute"},
		{regexp.MustCompile(`(?i)sau (\d+) giay nua`), "in $1 second"},
		{regexp.MustCompile(`(?i)sau (\d+) tuan nua`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) thang truoc`), "$1 month ago"},
		{regexp.MustCompile(`(?i)sau (\d+) gio nua`), "in $1 hour"},
		{regexp.MustCompile(`(?i)sau (\d+) nam nua`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ngay truoc`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) phut truoc`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) giay truoc`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) tuan truoc`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) gio truoc`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) nam truoc`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(sau \d+ thang nua|sau \d+ giay nua|sau \d+ ngay nua|sau \d+ phut nua|sau \d+ tuan nua|\d+ thang truoc|sau \d+ gio nua|sau \d+ nam nua|\d+ giay truoc|\d+ ngay truoc|\d+ phut truoc|\d+ tuan truoc|\d+ gio truoc|\d+ nam truoc)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(sau \d+ thang nua|sau \d+ giay nua|sau \d+ ngay nua|sau \d+ phut nua|sau \d+ tuan nua|\d+ thang truoc|sau \d+ gio nua|sau \d+ nam nua|\d+ giay truoc|\d+ ngay truoc|\d+ phut truoc|\d+ tuan truoc|\d+ gio truoc|\d+ nam truoc)$`),
	KnownWords:      []string{"thang muoi hai", "thang muoi mot", "giay đong ho", "thang truoc", "nguyen ban", "thang chin", "thang muoi", "tuan truoc", "nam ngoai", "thang bay", "thang hai", "thang mot", "thang nam", "thang nay", "thang sau", "thang sau", "thang tam", "truoc đay", "ban ngay", "cach đay", "chu nhat", "hang nhi", "ngay mai", "phut nay", "thang 10", "thang 11", "thang 12", "thang ba", "thang tu", "tuan nay", "tuan sau", "bay gio", "gio nay", "hom nay", "hom qua", "nam nay", "nam sau", "thang 1", "thang 2", "thang 3", "thang 4", "thang 5", "thang 6", "thang 7", "thang 8", "thang 9", "thu bay", "thu hai", "thu nam", "thu sau", "tuan le", "thg 10", "thg 11", "thg 12", "thu ba", "thu tu", "thang", "thg 1", "thg 2", "thg 3", "thg 4", "thg 5", "thg 6", "thg 7", "thg 8", "thg 9", "thu 1", "thu 2", "thu 3", "thu 4", "thu 5", "thu 6", "thu 7", "trong", "truoc", "buoi", "chut", "giay", "ngay", "phut", "th 2", "th 3", "th 4", "th 5", "th 6", "th 7", "tuan", "gio", "gmt", "lat", "luc", "nam", "thg", "utc", "am", "ch", "cn", "pm", "sa", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
