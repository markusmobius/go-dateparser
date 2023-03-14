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
	Translations: map[string][]string{
		"donnerstag": {"thursday"},
		"september":  {"september"},
		"dezember":   {"december"},
		"dienstag":   {"tuesday"},
		"mittwoch":   {"wednesday"},
		"november":   {"november"},
		"sekunden":   {"second"},
		"februar":    {"february"},
		"freitag":    {"friday"},
		"minuten":    {"minute"},
		"monaten":    {"month"},
		"oktober":    {"october"},
		"samstag":    {"saturday"},
		"sekunde":    {"second"},
		"sonntag":    {"sunday"},
		"stunden":    {"hour"},
		"august":     {"august"},
		"jahren":     {"year"},
		"janner":     {"january"},
		"januar":     {"january"},
		"minute":     {"minute"},
		"monate":     {"month"},
		"montag":     {"monday"},
		"stunde":     {"hour"},
		"wochen":     {"week"},
		"april":      {"april"},
		"feber":      {"february"},
		"jahre":      {"year"},
		"monat":      {"month"},
		"nachm":      {"pm"},
		"tagen":      {"day"},
		"woche":      {"week"},
		"etwa":       {""},
		"jahr":       {"year"},
		"juli":       {"july"},
		"juni":       {"june"},
		"marz":       {"march"},
		"tage":       {"day"},
		"vorm":       {"am"},
		"apr":        {"april"},
		"aug":        {"august"},
		"dez":        {"december"},
		"die":        {"tuesday"},
		"don":        {"thursday"},
		"feb":        {"february"},
		"fre":        {"friday"},
		"gmt":        {"gmt"},
		"jan":        {"january"},
		"jul":        {"july"},
		"jun":        {"june"},
		"mai":        {"may"},
		"mar":        {"march"},
		"min":        {"minute"},
		"mit":        {"wednesday"},
		"mon":        {"monday"},
		"mrz":        {"march"},
		"nov":        {"november"},
		"okt":        {"october"},
		"sam":        {"saturday"},
		"sek":        {"second"},
		"sep":        {"september"},
		"son":        {"sunday"},
		"std":        {"hour"},
		"tag":        {"day"},
		"uhr":        {""},
		"und":        {""},
		"utc":        {"utc"},
		"vor":        {"ago"},
		"am":         {"am"},
		"di":         {"tuesday"},
		"do":         {"thursday"},
		"fr":         {"friday"},
		"im":         {"in"},
		"in":         {"in"},
		"mi":         {"wednesday"},
		"mo":         {"monday"},
		"pm":         {"pm"},
		"sa":         {"saturday"},
		"so":         {"sunday"},
		"um":         {""},
		" ":          {" "},
		"'":          {""},
		"+":          {"+"},
		",":          {""},
		"-":          {"-"},
		".":          {"."},
		"/":          {"/"},
		":":          {":"},
		";":          {""},
		"@":          {""},
		"[":          {""},
		"]":          {""},
		"j":          {"year"},
		"m":          {"month"},
		"w":          {"week"},
		"z":          {"z"},
		"|":          {""},
	},
	RelativeType: map[string]string{
		"in dieser minute": "0 minute ago",
		"in dieser stunde": "0 hour ago",
		"nachsten monat":   "in 1 month",
		"letzten monat":    "1 month ago",
		"nachste woche":    "in 1 week",
		"nachstes jahr":    "in 1 year",
		"diesen monat":     "0 month ago",
		"letzte woche":     "1 week ago",
		"letztes jahr":     "1 year ago",
		"diese woche":      "0 week ago",
		"dieses jahr":      "0 year ago",
		"ubermorgen":       "in 2 day",
		"vorgestern":       "2 day ago",
		"gestern":          "1 day ago",
		"morgen":           "in 1 day",
		"heute":            "0 day ago",
		"jetzt":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) sekunden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) sekunden`), "in $1 second"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) minuten`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) monaten`), "$1 month ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) sekunde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) stunden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) minuten`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) monaten`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) sekunde`), "in $1 second"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) stunden`), "in $1 hour"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) jahren`), "$1 year ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) minute`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) stunde`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) wochen`), "$1 week ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) jahren`), "in $1 year"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) stunde`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) wochen`), "in $1 week"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) monat`), "$1 month ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) tagen`), "$1 day ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) woche`), "$1 week ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) monat`), "in $1 month"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) tagen`), "in $1 day"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) woche`), "in $1 week"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) jahr`), "$1 year ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) jahr`), "in $1 year"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) std`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) tag`), "$1 day ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*)\s*h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*)\s*m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*)\s*s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) std`), "in $1 hour"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) tag`), "in $1 day"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) wo`), "$1 week ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) wo`), "in $1 week"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(vor \d+[.,]?\d* sekunden|in \d+[.,]?\d* sekunden|vor \d+[.,]?\d* minuten|vor \d+[.,]?\d* monaten|vor \d+[.,]?\d* sekunde|vor \d+[.,]?\d* stunden|in \d+[.,]?\d* minuten|in \d+[.,]?\d* monaten|in \d+[.,]?\d* sekunde|in \d+[.,]?\d* stunden|vor \d+[.,]?\d* jahren|vor \d+[.,]?\d* minute|vor \d+[.,]?\d* stunde|vor \d+[.,]?\d* wochen|in \d+[.,]?\d* jahren|in \d+[.,]?\d* minute|in \d+[.,]?\d* stunde|in \d+[.,]?\d* wochen|vor \d+[.,]?\d* monat|vor \d+[.,]?\d* tagen|vor \d+[.,]?\d* woche|in \d+[.,]?\d* monat|in \d+[.,]?\d* tagen|in \d+[.,]?\d* woche|vor \d+[.,]?\d* jahr|in \d+[.,]?\d* jahr|vor \d+[.,]?\d* min|vor \d+[.,]?\d* sek|vor \d+[.,]?\d* std|vor \d+[.,]?\d* tag|vor \d+[.,]?\d*\s*h|vor \d+[.,]?\d*\s*m|vor \d+[.,]?\d*\s*s|in \d+[.,]?\d* min|in \d+[.,]?\d* sek|in \d+[.,]?\d* std|in \d+[.,]?\d* tag|vor \d+[.,]?\d* wo|in \d+[.,]?\d* wo|vor \d+[.,]?\d* m|vor \d+[.,]?\d* s|in \d+[.,]?\d* m|in \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(vor \d+[.,]?\d* sekunden|in \d+[.,]?\d* sekunden|vor \d+[.,]?\d* minuten|vor \d+[.,]?\d* monaten|vor \d+[.,]?\d* sekunde|vor \d+[.,]?\d* stunden|in \d+[.,]?\d* minuten|in \d+[.,]?\d* monaten|in \d+[.,]?\d* sekunde|in \d+[.,]?\d* stunden|vor \d+[.,]?\d* jahren|vor \d+[.,]?\d* minute|vor \d+[.,]?\d* stunde|vor \d+[.,]?\d* wochen|in \d+[.,]?\d* jahren|in \d+[.,]?\d* minute|in \d+[.,]?\d* stunde|in \d+[.,]?\d* wochen|vor \d+[.,]?\d* monat|vor \d+[.,]?\d* tagen|vor \d+[.,]?\d* woche|in \d+[.,]?\d* monat|in \d+[.,]?\d* tagen|in \d+[.,]?\d* woche|vor \d+[.,]?\d* jahr|in \d+[.,]?\d* jahr|vor \d+[.,]?\d* min|vor \d+[.,]?\d* sek|vor \d+[.,]?\d* std|vor \d+[.,]?\d* tag|vor \d+[.,]?\d*\s*h|vor \d+[.,]?\d*\s*m|vor \d+[.,]?\d*\s*s|in \d+[.,]?\d* min|in \d+[.,]?\d* sek|in \d+[.,]?\d* std|in \d+[.,]?\d* tag|vor \d+[.,]?\d* wo|in \d+[.,]?\d* wo|vor \d+[.,]?\d* m|vor \d+[.,]?\d* s|in \d+[.,]?\d* m|in \d+[.,]?\d* s)$`),
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
