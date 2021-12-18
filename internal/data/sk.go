// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sk_Locale = merge(nil, LocaleData{
	Name:      "sk",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"v tejto hodine": "0 hour ago",
		"v tejto minute": "0 minute ago",
		"minuly mesiac":  "1 month ago",
		"minuly tyzden":  "1 week ago",
		"buduci mesiac":  "in 1 month",
		"buduci tyzden":  "in 1 week",
		"tento mesiac":   "0 month ago",
		"tento tyzden":   "0 week ago",
		"minuly rok":     "1 year ago",
		"buduci rok":     "in 1 year",
		"tento rok":      "0 year ago",
		"september":      "september",
		"septembra":      "september",
		"december":       "december",
		"decembra":       "december",
		"februara":       "february",
		"pondelok":       "monday",
		"november":       "november",
		"novembra":       "november",
		"augusta":        "august",
		"februar":        "february",
		"januara":        "january",
		"oktober":        "october",
		"oktobra":        "october",
		"sekunda":        "second",
		"stvrtok":        "thursday",
		"aprila":         "april",
		"august":         "august",
		"piatok":         "friday",
		"hodina":         "hour",
		"zajtra":         "in 1 day",
		"januar":         "january",
		"minuta":         "minute",
		"mesiac":         "month",
		"sobota":         "saturday",
		"nedela":         "sunday",
		"utorok":         "tuesday",
		"streda":         "wednesday",
		"tyzden":         "week",
		"teraz":          "0 second ago",
		"vcera":          "1 day ago",
		"april":          "april",
		"marca":          "march",
		"marec":          "march",
		"dnes":           "0 day ago",
		"jula":           "july",
		"juna":           "june",
		"maja":           "may",
		"apr":            "april",
		"aug":            "august",
		"den":            "day",
		"dec":            "december",
		"feb":            "february",
		"gmt":            "gmt",
		"jan":            "january",
		"jul":            "july",
		"jun":            "june",
		"mar":            "march",
		"maj":            "may",
		"min":            "minute",
		"mes":            "month",
		"nov":            "november",
		"okt":            "october",
		"sep":            "september",
		"utc":            "utc",
		"tyz":            "week",
		"rok":            "year",
		"am":             "am",
		"pi":             "friday",
		"po":             "monday",
		"pm":             "pm",
		"so":             "saturday",
		"ne":             "sunday",
		"st":             "thursday",
		"ut":             "tuesday",
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
		"d":              "day",
		"h":              "hour",
		"s":              "second",
		"r":              "year",
		"z":              "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pred (\d+) sekundami`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinami`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutami`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesiacmi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesiacom`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sekundou`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tyzdnami`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinou`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutou`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tyzdnom`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) dnami`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) rokmi`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) rokom`), "$1 year ago"},
		{regexp.MustCompile(`(?i)o (\d+) mesiacov`), "in $1 month"},
		{regexp.MustCompile(`(?i)pred (\d+) dnom`), "$1 day ago"},
		{regexp.MustCompile(`(?i)o (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)o (\d+) tyzdnov`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tyz`), "$1 week ago"},
		{regexp.MustCompile(`(?i)o (\d+) hodinu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)o (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)o (\d+) mesiac`), "in $1 month"},
		{regexp.MustCompile(`(?i)o (\d+) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)o (\d+) tyzden`), "in $1 week"},
		{regexp.MustCompile(`(?i)o (\d+) hodin`), "in $1 hour"},
		{regexp.MustCompile(`(?i)o (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)o (\d+) rokov`), "in $1 year"},
		{regexp.MustCompile(`(?i)pred (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) r`), "$1 year ago"},
		{regexp.MustCompile(`(?i)o (\d+) den`), "in $1 day"},
		{regexp.MustCompile(`(?i)o (\d+) dni`), "in $1 day"},
		{regexp.MustCompile(`(?i)o (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)o (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)o (\d+) tyz`), "in $1 week"},
		{regexp.MustCompile(`(?i)o (\d+) rok`), "in $1 year"},
		{regexp.MustCompile(`(?i)o (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)o (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)o (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)o (\d+) r`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(pred \d+ sekundami|pred \d+ hodinami|pred \d+ mesiacmi|pred \d+ mesiacom|pred \d+ minutami|pred \d+ sekundou|pred \d+ tyzdnami|pred \d+ hodinou|pred \d+ minutou|pred \d+ tyzdnom|o \d+ mesiacov|pred \d+ dnami|pred \d+ rokmi|pred \d+ rokom|o \d+ sekundu|o \d+ tyzdnov|pred \d+ dnom|o \d+ hodinu|o \d+ mesiac|o \d+ minutu|o \d+ sekund|o \d+ tyzden|pred \d+ mes|pred \d+ min|pred \d+ tyz|o \d+ hodin|o \d+ minut|o \d+ rokov|pred \d+ d|pred \d+ h|pred \d+ r|pred \d+ s|o \d+ den|o \d+ dni|o \d+ mes|o \d+ min|o \d+ rok|o \d+ tyz|o \d+ d|o \d+ h|o \d+ r|o \d+ s)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+ sekundami|pred \d+ hodinami|pred \d+ mesiacmi|pred \d+ mesiacom|pred \d+ minutami|pred \d+ sekundou|pred \d+ tyzdnami|pred \d+ hodinou|pred \d+ minutou|pred \d+ tyzdnom|o \d+ mesiacov|pred \d+ dnami|pred \d+ rokmi|pred \d+ rokom|o \d+ sekundu|o \d+ tyzdnov|pred \d+ dnom|o \d+ hodinu|o \d+ mesiac|o \d+ minutu|o \d+ sekund|o \d+ tyzden|pred \d+ mes|pred \d+ min|pred \d+ tyz|o \d+ hodin|o \d+ minut|o \d+ rokov|pred \d+ d|pred \d+ h|pred \d+ r|pred \d+ s|o \d+ den|o \d+ dni|o \d+ mes|o \d+ min|o \d+ rok|o \d+ tyz|o \d+ d|o \d+ h|o \d+ r|o \d+ s)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(v tejto hodine|v tejto minute|buduci mesiac|buduci tyzden|minuly mesiac|minuly tyzden|tento mesiac|tento tyzden|buduci rok|minuly rok|september|septembra|tento rok|december|decembra|februara|november|novembra|pondelok|augusta|februar|januara|oktober|oktobra|sekunda|stvrtok|aprila|august|hodina|januar|mesiac|minuta|nedela|piatok|sobota|streda|tyzden|utorok|zajtra|april|marca|marec|teraz|vcera|dnes|jula|juna|maja|apr|aug|dec|den|feb|gmt|jan|jul|jun|maj|mar|mes|min|nov|okt|rok|sep|tyz|utc|\+|\.|\[|\]|\||am|ne|pi|pm|po|so|st|ut| |'|,|-|/|:|;|@|d|h|r|s|z)((?:\z|\W|_|\d).*)$`),
})
