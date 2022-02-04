// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var de_Locale = merge(nil, LocaleData{
	Name:                  "de",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdefghijklnorstuvwzäü`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)einer(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)einem(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ein(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)zwei(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)drei(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)vier(\z|[^\pL\pM\d]|_)`), "${1}4${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)funf(\z|[^\pL\pM\d]|_)`), "${1}5${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)sechs(\z|[^\pL\pM\d]|_)`), "${1}6${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)sieben(\z|[^\pL\pM\d]|_)`), "${1}7${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)acht(\z|[^\pL\pM\d]|_)`), "${1}8${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)neun(\z|[^\pL\pM\d]|_)`), "${1}9${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)zehn(\z|[^\pL\pM\d]|_)`), "${1}10${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)elf(\z|[^\pL\pM\d]|_)`), "${1}11${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)zwolf(\z|[^\pL\pM\d]|_)`), "${1}12${2}"},
	},
	Translations: map[string]string{
		"donnerstag": "thursday",
		"september":  "september",
		"dezember":   "december",
		"november":   "november",
		"sekunden":   "second",
		"dienstag":   "tuesday",
		"mittwoch":   "wednesday",
		"februar":    "february",
		"freitag":    "friday",
		"stunden":    "hour",
		"minuten":    "minute",
		"monaten":    "month",
		"oktober":    "october",
		"samstag":    "saturday",
		"sekunde":    "second",
		"sonntag":    "sunday",
		"august":     "august",
		"stunde":     "hour",
		"janner":     "january",
		"januar":     "january",
		"minute":     "minute",
		"montag":     "monday",
		"monate":     "month",
		"wochen":     "week",
		"jahren":     "year",
		"april":      "april",
		"tagen":      "day",
		"feber":      "february",
		"monat":      "month",
		"nachm":      "pm",
		"woche":      "week",
		"jahre":      "year",
		"etwa":       "",
		"vorm":       "am",
		"tage":       "day",
		"juli":       "july",
		"juni":       "june",
		"marz":       "march",
		"jahr":       "year",
		"uhr":        "",
		"und":        "",
		"vor":        "ago",
		"apr":        "april",
		"aug":        "august",
		"tag":        "day",
		"dez":        "december",
		"feb":        "february",
		"fre":        "friday",
		"gmt":        "gmt",
		"std":        "hour",
		"jan":        "january",
		"jul":        "july",
		"jun":        "june",
		"mar":        "march",
		"mrz":        "march",
		"mai":        "may",
		"min":        "minute",
		"mon":        "monday",
		"nov":        "november",
		"okt":        "october",
		"sam":        "saturday",
		"sek":        "second",
		"sep":        "september",
		"son":        "sunday",
		"don":        "thursday",
		"die":        "tuesday",
		"utc":        "utc",
		"mit":        "wednesday",
		"um":         "",
		"am":         "am",
		"fr":         "friday",
		"im":         "in",
		"in":         "in",
		"mo":         "monday",
		"pm":         "pm",
		"sa":         "saturday",
		"so":         "sunday",
		"do":         "thursday",
		"di":         "tuesday",
		"mi":         "wednesday",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"m":          "month",
		"w":          "week",
		"j":          "year",
		"z":          "z",
	},
	RelativeType: map[string]string{
		"in dieser stunde": "0 hour ago",
		"in dieser minute": "0 minute ago",
		"nachsten monat":   "in 1 month",
		"letzten monat":    "1 month ago",
		"nachste woche":    "in 1 week",
		"nachstes jahr":    "in 1 year",
		"diesen monat":     "0 month ago",
		"letzte woche":     "1 week ago",
		"letztes jahr":     "1 year ago",
		"diese woche":      "0 week ago",
		"dieses jahr":      "0 year ago",
		"vorgestern":       "2 day ago",
		"ubermorgen":       "in 2 day",
		"gestern":          "1 day ago",
		"morgen":           "in 1 day",
		"heute":            "0 day ago",
		"jetzt":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)vor (\d+) sekunden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)vor (\d+) stunden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+) minuten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+) monaten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)vor (\d+) sekunde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) sekunden`), "in $1 second"},
		{regexp.MustCompile(`(?i)vor (\d+) stunde`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+) minute`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+) wochen`), "$1 week ago"},
		{regexp.MustCompile(`(?i)vor (\d+) jahren`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) stunden`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) minuten`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) monaten`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) sekunde`), "in $1 second"},
		{regexp.MustCompile(`(?i)vor (\d+) tagen`), "$1 day ago"},
		{regexp.MustCompile(`(?i)vor (\d+) monat`), "$1 month ago"},
		{regexp.MustCompile(`(?i)vor (\d+) woche`), "$1 week ago"},
		{regexp.MustCompile(`(?i)in (\d+) stunde`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) wochen`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) jahren`), "in $1 year"},
		{regexp.MustCompile(`(?i)vor (\d+) jahr`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+) tagen`), "in $1 day"},
		{regexp.MustCompile(`(?i)in (\d+) monat`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+) woche`), "in $1 week"},
		{regexp.MustCompile(`(?i)vor (\d+) tag`), "$1 day ago"},
		{regexp.MustCompile(`(?i)vor (\d+) std`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+)\s*h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+)\s*m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)vor (\d+)\s*s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) jahr`), "in $1 year"},
		{regexp.MustCompile(`(?i)vor (\d+) wo`), "$1 week ago"},
		{regexp.MustCompile(`(?i)in (\d+) tag`), "in $1 day"},
		{regexp.MustCompile(`(?i)in (\d+) std`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)vor (\d+) m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+) wo`), "in $1 week"},
		{regexp.MustCompile(`(?i)in (\d+) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(vor \d+ sekunden|in \d+ sekunden|vor \d+ minuten|vor \d+ monaten|vor \d+ sekunde|vor \d+ stunden|in \d+ minuten|in \d+ monaten|in \d+ sekunde|in \d+ stunden|vor \d+ jahren|vor \d+ minute|vor \d+ stunde|vor \d+ wochen|in \d+ jahren|in \d+ minute|in \d+ stunde|in \d+ wochen|vor \d+ monat|vor \d+ tagen|vor \d+ woche|in \d+ monat|in \d+ tagen|in \d+ woche|vor \d+ jahr|in \d+ jahr|vor \d+ min|vor \d+ sek|vor \d+ std|vor \d+ tag|vor \d+\s*h|vor \d+\s*m|vor \d+\s*s|in \d+ min|in \d+ sek|in \d+ std|in \d+ tag|vor \d+ wo|in \d+ wo|vor \d+ m|vor \d+ s|in \d+ m|in \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(vor \d+ sekunden|in \d+ sekunden|vor \d+ minuten|vor \d+ monaten|vor \d+ sekunde|vor \d+ stunden|in \d+ minuten|in \d+ monaten|in \d+ sekunde|in \d+ stunden|vor \d+ jahren|vor \d+ minute|vor \d+ stunde|vor \d+ wochen|in \d+ jahren|in \d+ minute|in \d+ stunde|in \d+ wochen|vor \d+ monat|vor \d+ tagen|vor \d+ woche|in \d+ monat|in \d+ tagen|in \d+ woche|vor \d+ jahr|in \d+ jahr|vor \d+ min|vor \d+ sek|vor \d+ std|vor \d+ tag|vor \d+\s*h|vor \d+\s*m|vor \d+\s*s|in \d+ min|in \d+ sek|in \d+ std|in \d+ tag|vor \d+ wo|in \d+ wo|vor \d+ m|vor \d+ s|in \d+ m|in \d+ s)$`),
	KnownWords:      []string{"in dieser minute", "in dieser stunde", "nachsten monat", "letzten monat", "nachste woche", "nachstes jahr", "diesen monat", "letzte woche", "letztes jahr", "diese woche", "dieses jahr", "donnerstag", "ubermorgen", "vorgestern", "september", "dezember", "dienstag", "mittwoch", "november", "sekunden", "februar", "freitag", "gestern", "minuten", "monaten", "oktober", "samstag", "sekunde", "sonntag", "stunden", "august", "jahren", "janner", "januar", "minute", "monate", "montag", "morgen", "stunde", "wochen", "april", "feber", "heute", "jahre", "jetzt", "monat", "nachm", "tagen", "woche", "etwa", "jahr", "juli", "juni", "marz", "tage", "vorm", "apr", "aug", "dez", "die", "don", "feb", "fre", "gmt", "jan", "jul", "jun", "mai", "mar", "min", "mit", "mon", "mrz", "nov", "okt", "sam", "sek", "sep", "son", "std", "tag", "uhr", "und", "utc", "vor", "am", "di", "do", "fr", "im", "in", "mi", "mo", "pm", "sa", "so", "um", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "j", "m", "w", "z", "|"},
})

var de_AT_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-AT",
	DateOrder: "DMY",
})

var de_BE_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-BE",
	DateOrder: "DMY",
})

var de_CH_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-CH",
	DateOrder: "DMY",
})

var de_IT_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-IT",
	DateOrder: "DMY",
})

var de_LI_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-LI",
	DateOrder: "DMY",
})

var de_LU_Locale = merge(&de_Locale, LocaleData{
	Name:      "de-LU",
	DateOrder: "DMY",
})
