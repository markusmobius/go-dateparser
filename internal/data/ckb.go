// Code is generated by script; DO NOT EDIT.

package data

var (
	ckb_Locale    LocaleData
	ckb_IR_Locale LocaleData
)

func init() {
	ckb_Locale = merge(nil, LocaleData{
		Name:      "ckb",
		DateOrder: "YMD",
		Charset:   []rune(`cdeghiklnorstuwxyzئابتجحدرزسشلمنويپچکھیێە`),
		Translations: map[string][]string{
			"کانوونی دووەم": {"january"},
			"تشرینی دووەم":  {"november"},
			"تشرینی یەکەم":  {"october"},
			"کانونی یەکەم":  {"december"},
			"پێنجشەممە":     {"thursday"},
			"چوارشەممە":     {"wednesday"},
			"حوزەیران":      {"june"},
			"دووشەممە":      {"monday"},
			"یەکشەممە":      {"sunday"},
			"سێشەممە":       {"tuesday"},
			"يەیلوول":       {"september"},
			"minute":        {"minute"},
			"second":        {"second"},
			"تەمووز":        {"july"},
			"month":         {"month"},
			"شوبات":         {"february"},
			"شەممە":         {"saturday"},
			"نیسان":         {"april"},
			"يازار":         {"march"},
			"يایار":         {"may"},
			"ھەینی":         {"friday"},
			"hour":          {"hour"},
			"week":          {"week"},
			"year":          {"year"},
			"day":           {"day"},
			"gmt":           {"gmt"},
			"utc":           {"utc"},
			"ياب":           {"august"},
			"am":            {"am"},
			"pm":            {"pm"},
			"بن":            {"am"},
			"دن":            {"pm"},
			" ":             {" "},
			"'":             {""},
			"+":             {"+"},
			",":             {""},
			"-":             {"-"},
			".":             {"."},
			"/":             {"/"},
			":":             {":"},
			";":             {""},
			"@":             {""},
			"[":             {""},
			"]":             {""},
			"z":             {"z"},
			"|":             {""},
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
			"yesterday":   "1 day ago",
			"tomorrow":    "in 1 day",
			"today":       "0 day ago",
			"now":         "0 second ago",
		},
		KnownWords: []string{"کانوونی دووەم", "تشرینی دووەم", "تشرینی یەکەم", "کانونی یەکەم", "this minute", "last month", "next month", "this month", "last week", "last year", "next week", "next year", "this hour", "this week", "this year", "yesterday", "پێنجشەممە", "چوارشەممە", "tomorrow", "حوزەیران", "دووشەممە", "یەکشەممە", "سێشەممە", "يەیلوول", "minute", "second", "تەمووز", "month", "today", "شوبات", "شەممە", "نیسان", "يازار", "يایار", "ھەینی", "hour", "week", "year", "day", "gmt", "now", "utc", "ياب", "am", "pm", "بن", "دن", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})

	ckb_IR_Locale = merge(&ckb_Locale, LocaleData{
		Name:      "ckb-IR",
		DateOrder: "YMD",
	})
}
