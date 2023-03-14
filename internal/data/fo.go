// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fo_Locale = merge(nil, LocaleData{
	Name:      "fo",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvyzáæíðóøúý`),
	Translations: map[string][]string{
		"friggjadagur": {"friday"},
		"leygardagur":  {"saturday"},
		"sunnudagur":   {"sunday"},
		"manadagur":    {"monday"},
		"mikudagur":    {"wednesday"},
		"september":    {"september"},
		"desember":     {"december"},
		"hosdagur":     {"thursday"},
		"minuttur":     {"minute"},
		"november":     {"november"},
		"tysdagur":     {"tuesday"},
		"februar":      {"february"},
		"manaður":      {"month"},
		"oktober":      {"october"},
		"august":       {"august"},
		"januar":       {"january"},
		"sekund":       {"second"},
		"april":        {"april"},
		"dagur":        {"day"},
		"juli":         {"july"},
		"juni":         {"june"},
		"mars":         {"march"},
		"timi":         {"hour"},
		"vika":         {"week"},
		"apr":          {"april"},
		"aug":          {"august"},
		"des":          {"december"},
		"feb":          {"february"},
		"fri":          {"friday"},
		"gmt":          {"gmt"},
		"hos":          {"thursday"},
		"jan":          {"january"},
		"jul":          {"july"},
		"jun":          {"june"},
		"ley":          {"saturday"},
		"mai":          {"may"},
		"man":          {"monday"},
		"mar":          {"march"},
		"mik":          {"wednesday"},
		"min":          {"minute"},
		"mnð":          {"month"},
		"nov":          {"november"},
		"okt":          {"october"},
		"sek":          {"second"},
		"sep":          {"september"},
		"sun":          {"sunday"},
		"tys":          {"tuesday"},
		"utc":          {"utc"},
		"am":           {"am"},
		"ar":           {"year"},
		"da":           {"day"},
		"pm":           {"pm"},
		"vi":           {"week"},
		" ":            {" "},
		"'":            {""},
		"+":            {"+"},
		",":            {""},
		"-":            {"-"},
		".":            {"."},
		"/":            {"/"},
		":":            {":"},
		";":            {""},
		"@":            {""},
		"[":            {""},
		"]":            {""},
		"d":            {"day"},
		"m":            {"minute"},
		"s":            {"second"},
		"t":            {"hour"},
		"v":            {"week"},
		"z":            {"z"},
		"|":            {""},
	},
	RelativeType: map[string]string{
		"hendan minuttin": "0 minute ago",
		"seinasta manað":  "1 month ago",
		"henda manaðin":   "0 month ago",
		"seinastu viku":   "1 week ago",
		"hendan timan":    "0 hour ago",
		"næsta manað":     "in 1 month",
		"næstu viku":      "in 1 week",
		"hesu viku":       "0 week ago",
		"i morgin":        "in 1 day",
		"næsta ar":        "in 1 year",
		"i fjør":          "1 year ago",
		"i gjar":          "1 day ago",
		"i dag":           "0 day ago",
		"i ar":            "0 year ago",
		"nu":              "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuttir siðan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) manaðir siðan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minutt siðan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sekund siðan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dagar siðan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dagur siðan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) manað siðan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) timar siðan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vikur siðan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) minuttir`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) timi siðan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vika siðan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) manaðir`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min siðan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mnð siðan`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sek siðan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) minutt`), "in $1 minute"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ar siðan`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) da siðan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) vi siðan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) dagar`), "in $1 day"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) manað`), "in $1 month"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) timar`), "in $1 hour"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) vikur`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) d siðan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) m siðan`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) s siðan`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) t siðan`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) v siðan`), "$1 week ago"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) tima`), "in $1 hour"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) viku`), "in $1 week"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) mnð`), "in $1 month"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) ar`), "in $1 year"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) da`), "in $1 day"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) vi`), "in $1 week"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) t`), "in $1 hour"},
		{regexp.MustCompile(`(?i)um (\d+[.,]?\d*) v`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* minuttir siðan|\d+[.,]?\d* manaðir siðan|\d+[.,]?\d* minutt siðan|\d+[.,]?\d* sekund siðan|\d+[.,]?\d* dagar siðan|\d+[.,]?\d* dagur siðan|\d+[.,]?\d* manað siðan|\d+[.,]?\d* timar siðan|\d+[.,]?\d* vikur siðan|um \d+[.,]?\d* minuttir|\d+[.,]?\d* timi siðan|\d+[.,]?\d* vika siðan|um \d+[.,]?\d* manaðir|\d+[.,]?\d* min siðan|\d+[.,]?\d* mnð siðan|\d+[.,]?\d* sek siðan|um \d+[.,]?\d* minutt|um \d+[.,]?\d* sekund|\d+[.,]?\d* ar siðan|\d+[.,]?\d* da siðan|\d+[.,]?\d* vi siðan|um \d+[.,]?\d* dagar|um \d+[.,]?\d* manað|um \d+[.,]?\d* timar|um \d+[.,]?\d* vikur|\d+[.,]?\d* d siðan|\d+[.,]?\d* m siðan|\d+[.,]?\d* s siðan|\d+[.,]?\d* t siðan|\d+[.,]?\d* v siðan|um \d+[.,]?\d* tima|um \d+[.,]?\d* viku|um \d+[.,]?\d* dag|um \d+[.,]?\d* min|um \d+[.,]?\d* mnð|um \d+[.,]?\d* sek|um \d+[.,]?\d* ar|um \d+[.,]?\d* da|um \d+[.,]?\d* vi|um \d+[.,]?\d* d|um \d+[.,]?\d* m|um \d+[.,]?\d* s|um \d+[.,]?\d* t|um \d+[.,]?\d* v)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* minuttir siðan|\d+[.,]?\d* manaðir siðan|\d+[.,]?\d* minutt siðan|\d+[.,]?\d* sekund siðan|\d+[.,]?\d* dagar siðan|\d+[.,]?\d* dagur siðan|\d+[.,]?\d* manað siðan|\d+[.,]?\d* timar siðan|\d+[.,]?\d* vikur siðan|um \d+[.,]?\d* minuttir|\d+[.,]?\d* timi siðan|\d+[.,]?\d* vika siðan|um \d+[.,]?\d* manaðir|\d+[.,]?\d* min siðan|\d+[.,]?\d* mnð siðan|\d+[.,]?\d* sek siðan|um \d+[.,]?\d* minutt|um \d+[.,]?\d* sekund|\d+[.,]?\d* ar siðan|\d+[.,]?\d* da siðan|\d+[.,]?\d* vi siðan|um \d+[.,]?\d* dagar|um \d+[.,]?\d* manað|um \d+[.,]?\d* timar|um \d+[.,]?\d* vikur|\d+[.,]?\d* d siðan|\d+[.,]?\d* m siðan|\d+[.,]?\d* s siðan|\d+[.,]?\d* t siðan|\d+[.,]?\d* v siðan|um \d+[.,]?\d* tima|um \d+[.,]?\d* viku|um \d+[.,]?\d* dag|um \d+[.,]?\d* min|um \d+[.,]?\d* mnð|um \d+[.,]?\d* sek|um \d+[.,]?\d* ar|um \d+[.,]?\d* da|um \d+[.,]?\d* vi|um \d+[.,]?\d* d|um \d+[.,]?\d* m|um \d+[.,]?\d* s|um \d+[.,]?\d* t|um \d+[.,]?\d* v)$`),
	KnownWords:      []string{"hendan minuttin", "seinasta manað", "henda manaðin", "seinastu viku", "friggjadagur", "hendan timan", "leygardagur", "næsta manað", "næstu viku", "sunnudagur", "hesu viku", "manadagur", "mikudagur", "september", "desember", "hosdagur", "i morgin", "minuttur", "november", "næsta ar", "tysdagur", "februar", "manaður", "oktober", "august", "i fjør", "i gjar", "januar", "sekund", "april", "dagur", "i dag", "i ar", "juli", "juni", "mars", "timi", "vika", "apr", "aug", "des", "feb", "fri", "gmt", "hos", "jan", "jul", "jun", "ley", "mai", "man", "mar", "mik", "min", "mnð", "nov", "okt", "sek", "sep", "sun", "tys", "utc", "am", "ar", "da", "nu", "pm", "vi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "m", "s", "t", "v", "z", "|"},
})

var fo_DK_Locale = merge(&fo_Locale, LocaleData{
	Name:      "fo-DK",
	DateOrder: "DMY",
})
