// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	nb_Locale    LocaleData
	nb_SJ_Locale LocaleData
)

func init() {
	nb_Locale = merge(nil, LocaleData{
		Name:      "nb",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefgijklnorstuvzåø`),
		Translations: map[string][]string{
			"september": {"september"},
			"desember":  {"december"},
			"november":  {"november"},
			"februar":   {"february"},
			"oktober":   {"october"},
			"tirsdag":   {"tuesday"},
			"torsdag":   {"thursday"},
			"august":    {"august"},
			"fredag":    {"friday"},
			"januar":    {"january"},
			"lørdag":    {"saturday"},
			"mandag":    {"monday"},
			"minutt":    {"minute"},
			"onsdag":    {"wednesday"},
			"sekund":    {"second"},
			"søndag":    {"sunday"},
			"april":     {"april"},
			"dager":     {"day"},
			"maned":     {"month"},
			"siden":     {"ago"},
			"timer":     {"hour"},
			"juli":      {"july"},
			"juni":      {"june"},
			"mars":      {"march"},
			"time":      {"hour"},
			"uker":      {"week"},
			"apr":       {"april"},
			"aug":       {"august"},
			"dag":       {"day"},
			"des":       {"december"},
			"feb":       {"february"},
			"fre":       {"friday"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jul":       {"july"},
			"jun":       {"june"},
			"lør":       {"saturday"},
			"mai":       {"may"},
			"man":       {"monday"},
			"mar":       {"march"},
			"min":       {"minute"},
			"mnd":       {"month"},
			"nov":       {"november"},
			"okt":       {"october"},
			"ons":       {"wednesday"},
			"sek":       {"second"},
			"sep":       {"september"},
			"søn":       {"sunday"},
			"tir":       {"tuesday"},
			"tor":       {"thursday"},
			"uke":       {"week"},
			"utc":       {"utc"},
			"am":        {"am"},
			"ar":        {"year"},
			"md":        {"month"},
			"om":        {"in"},
			"pm":        {"pm"},
			" ":         {" "},
			"'":         {""},
			"+":         {"+"},
			",":         {""},
			"-":         {"-"},
			".":         {"."},
			"/":         {"/"},
			":":         {":"},
			";":         {""},
			"@":         {""},
			"[":         {""},
			"]":         {""},
			"d":         {"day"},
			"m":         {"minute"},
			"s":         {"second"},
			"t":         {"hour"},
			"u":         {"week"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"dette minuttet": "0 minute ago",
			"denne maneden":  "0 month ago",
			"forrige maned":  "1 month ago",
			"denne timen":    "0 hour ago",
			"forrige uke":    "1 week ago",
			"neste maned":    "in 1 month",
			"denne uken":     "0 week ago",
			"forrige md":     "1 month ago",
			"neste uke":      "in 1 week",
			"denne md":       "0 month ago",
			"i morgen":       "in 1 day",
			"neste ar":       "in 1 year",
			"neste md":       "in 1 month",
			"i fjor":         "1 year ago",
			"i dag":          "0 day ago",
			"i gar":          "1 day ago",
			"i ar":           "0 year ago",
			"na":             "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minutter siden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekunder siden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) maneder siden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minutt siden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekund siden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) dager siden`), "${1} day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) maned siden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) timer siden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) døgn siden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) time siden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) uker siden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) min siden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sek siden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) uke siden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) ar siden`), "$1 year ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) md siden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) d siden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) t siden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) u siden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minutter`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekunder`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) maneder`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minutt`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekund`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) maned`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) timer`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) døgn`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) time`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) uker`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) uke`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) ar`), "in $1 year"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) md`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) t`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) u`), "in $1 week"},
			{regexp.MustCompile(`(?i)–(\d+[.,]?\d*) ar`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+[.,]?\d* minutter siden|for \d+[.,]?\d* sekunder siden|for \d+[.,]?\d* maneder siden|for \d+[.,]?\d* minutt siden|for \d+[.,]?\d* sekund siden|for \d+[.,]?\d* dager siden|for \d+[.,]?\d* maned siden|for \d+[.,]?\d* timer siden|for \d+[.,]?\d* døgn siden|for \d+[.,]?\d* time siden|for \d+[.,]?\d* uker siden|for \d+[.,]?\d* min siden|for \d+[.,]?\d* sek siden|for \d+[.,]?\d* uke siden|for \d+[.,]?\d* ar siden|for \d+[.,]?\d* md siden|for \d+[.,]?\d* d siden|for \d+[.,]?\d* t siden|for \d+[.,]?\d* u siden|om \d+[.,]?\d* minutter|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* maneder|om \d+[.,]?\d* minutt|om \d+[.,]?\d* sekund|om \d+[.,]?\d* maned|om \d+[.,]?\d* timer|om \d+[.,]?\d* døgn|om \d+[.,]?\d* time|om \d+[.,]?\d* uker|om \d+[.,]?\d* min|om \d+[.,]?\d* sek|om \d+[.,]?\d* uke|om \d+[.,]?\d* ar|om \d+[.,]?\d* md|om \d+[.,]?\d* d|om \d+[.,]?\d* t|om \d+[.,]?\d* u|–\d+[.,]?\d* ar)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(for \d+[.,]?\d* minutter siden|for \d+[.,]?\d* sekunder siden|for \d+[.,]?\d* maneder siden|for \d+[.,]?\d* minutt siden|for \d+[.,]?\d* sekund siden|for \d+[.,]?\d* dager siden|for \d+[.,]?\d* maned siden|for \d+[.,]?\d* timer siden|for \d+[.,]?\d* døgn siden|for \d+[.,]?\d* time siden|for \d+[.,]?\d* uker siden|for \d+[.,]?\d* min siden|for \d+[.,]?\d* sek siden|for \d+[.,]?\d* uke siden|for \d+[.,]?\d* ar siden|for \d+[.,]?\d* md siden|for \d+[.,]?\d* d siden|for \d+[.,]?\d* t siden|for \d+[.,]?\d* u siden|om \d+[.,]?\d* minutter|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* maneder|om \d+[.,]?\d* minutt|om \d+[.,]?\d* sekund|om \d+[.,]?\d* maned|om \d+[.,]?\d* timer|om \d+[.,]?\d* døgn|om \d+[.,]?\d* time|om \d+[.,]?\d* uker|om \d+[.,]?\d* min|om \d+[.,]?\d* sek|om \d+[.,]?\d* uke|om \d+[.,]?\d* ar|om \d+[.,]?\d* md|om \d+[.,]?\d* d|om \d+[.,]?\d* t|om \d+[.,]?\d* u|–\d+[.,]?\d* ar)$`),
		KnownWords:      []string{"dette minuttet", "denne maneden", "forrige maned", "denne timen", "forrige uke", "neste maned", "denne uken", "forrige md", "neste uke", "september", "denne md", "desember", "i morgen", "neste ar", "neste md", "november", "februar", "oktober", "tirsdag", "torsdag", "august", "fredag", "i fjor", "januar", "lørdag", "mandag", "minutt", "onsdag", "sekund", "søndag", "april", "dager", "i dag", "i gar", "maned", "siden", "timer", "i ar", "juli", "juni", "mars", "time", "uker", "apr", "aug", "dag", "des", "feb", "fre", "gmt", "jan", "jul", "jun", "lør", "mai", "man", "mar", "min", "mnd", "nov", "okt", "ons", "sek", "sep", "søn", "tir", "tor", "uke", "utc", "am", "ar", "md", "na", "om", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "m", "s", "t", "u", "z", "|"},
	})

	nb_SJ_Locale = merge(&nb_Locale, LocaleData{
		Name:      "nb-SJ",
		DateOrder: "DMY",
	})
}
