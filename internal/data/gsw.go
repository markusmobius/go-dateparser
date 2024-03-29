// Code is generated by script; DO NOT EDIT.

package data

var (
	gsw_Locale    LocaleData
	gsw_FR_Locale LocaleData
	gsw_LI_Locale LocaleData
)

func init() {
	gsw_Locale = merge(nil, LocaleData{
		Name:      "gsw",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvwxyzäü`),
		Translations: map[string][]string{
			"am vormittag": {"am"},
			"am namittag":  {"pm"},
			"dunschtig":    {"thursday"},
			"samschtig":    {"saturday"},
			"septamber":    {"september"},
			"vormittag":    {"am"},
			"ziischtig":    {"tuesday"},
			"auguscht":     {"august"},
			"dezamber":     {"december"},
			"mittwuch":     {"wednesday"},
			"namittag":     {"pm"},
			"novamber":     {"november"},
			"oktoober":     {"october"},
			"februar":      {"february"},
			"friitig":      {"friday"},
			"maantig":      {"monday"},
			"minuute":      {"minute"},
			"schtund":      {"hour"},
			"sekunde":      {"second"},
			"sunntig":      {"sunday"},
			"januar":       {"january"},
			"april":        {"april"},
			"monet":        {"month"},
			"wuche":        {"week"},
			"jaar":         {"year"},
			"juli":         {"july"},
			"juni":         {"june"},
			"marz":         {"march"},
			"vorm":         {"am"},
			"apr":          {"april"},
			"aug":          {"august"},
			"dez":          {"december"},
			"feb":          {"february"},
			"gmt":          {"gmt"},
			"jan":          {"january"},
			"jul":          {"july"},
			"jun":          {"june"},
			"mai":          {"may"},
			"mar":          {"march"},
			"nam":          {"pm"},
			"nov":          {"november"},
			"okt":          {"october"},
			"sep":          {"september"},
			"tag":          {"day"},
			"utc":          {"utc"},
			"am":           {"am"},
			"du":           {"thursday"},
			"fr":           {"friday"},
			"ma":           {"monday"},
			"mi":           {"wednesday"},
			"pm":           {"pm"},
			"sa":           {"saturday"},
			"su":           {"sunday"},
			"zi":           {"tuesday"},
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
			"z":            {"z"},
			"|":            {""},
		},
		RelativeType: map[string]string{
			"this minute": "0 minute ago",
			"last month":  "1 month ago",
			"next month":  "in 1 month",
			"this month":  "0 month ago",
			"last week":   "1 week ago",
			"last year":   "1 year ago",
			"next week":   "in 1 week",
			"next year":   "in 1 year",
			"this hour":   "0 hour ago",
			"this week":   "0 week ago",
			"this year":   "0 year ago",
			"geschter":    "1 day ago",
			"moorn":       "in 1 day",
			"hut":         "0 day ago",
			"now":         "0 second ago",
		},
		KnownWords: []string{"am vormittag", "am namittag", "this minute", "last month", "next month", "this month", "dunschtig", "last week", "last year", "next week", "next year", "samschtig", "septamber", "this hour", "this week", "this year", "vormittag", "ziischtig", "auguscht", "dezamber", "geschter", "mittwuch", "namittag", "novamber", "oktoober", "februar", "friitig", "maantig", "minuute", "schtund", "sekunde", "sunntig", "januar", "april", "monet", "moorn", "wuche", "jaar", "juli", "juni", "marz", "vorm", "apr", "aug", "dez", "feb", "gmt", "hut", "jan", "jul", "jun", "mai", "mar", "nam", "nov", "now", "okt", "sep", "tag", "utc", "am", "du", "fr", "ma", "mi", "pm", "sa", "su", "zi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})

	gsw_FR_Locale = merge(&gsw_Locale, LocaleData{
		Name:      "gsw-FR",
		DateOrder: "DMY",
	})

	gsw_LI_Locale = merge(&gsw_Locale, LocaleData{
		Name:      "gsw-LI",
		DateOrder: "DMY",
	})
}
