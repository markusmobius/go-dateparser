// Code is generated by script; DO NOT EDIT.

package data

var (
	ha_Locale    LocaleData
	ha_GH_Locale LocaleData
	ha_NE_Locale LocaleData
)

func init() {
	ha_Locale = merge(nil, LocaleData{
		Name:      "ha",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuwxyzƙ`),
		Translations: map[string][]string{
			"faburairu": {"february"},
			"afirilu":   {"april"},
			"alhamis":   {"thursday"},
			"disamba":   {"december"},
			"janairu":   {"january"},
			"jumma'a":   {"friday"},
			"litinin":   {"monday"},
			"nuwamba":   {"november"},
			"satumba":   {"september"},
			"shekara":   {"year"},
			"agusta":    {"august"},
			"asabar":    {"saturday"},
			"daƙiƙa":    {"second"},
			"lahadi":    {"sunday"},
			"laraba":    {"wednesday"},
			"oktoba":    {"october"},
			"talata":    {"tuesday"},
			"kwana":     {"day"},
			"maris":     {"march"},
			"minti":     {"minute"},
			"mako":      {"week"},
			"mayu":      {"may"},
			"wata":      {"month"},
			"yuli":      {"july"},
			"yuni":      {"june"},
			"afi":       {"april"},
			"agu":       {"august"},
			"alh":       {"thursday"},
			"asa":       {"saturday"},
			"awa":       {"hour"},
			"dis":       {"december"},
			"fab":       {"february"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jum":       {"friday"},
			"lah":       {"sunday"},
			"lar":       {"wednesday"},
			"lit":       {"monday"},
			"mar":       {"march"},
			"may":       {"may"},
			"nuw":       {"november"},
			"okt":       {"october"},
			"sat":       {"september"},
			"tal":       {"tuesday"},
			"utc":       {"utc"},
			"yul":       {"july"},
			"yun":       {"june"},
			"am":        {"am"},
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
			"z":         {"z"},
			"|":         {""},
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
			"gobe":        "in 1 day",
			"jiya":        "1 day ago",
			"now":         "0 second ago",
			"yau":         "0 day ago",
		},
		KnownWords: []string{"this minute", "last month", "next month", "this month", "faburairu", "last week", "last year", "next week", "next year", "this hour", "this week", "this year", "afirilu", "alhamis", "disamba", "janairu", "jumma'a", "litinin", "nuwamba", "satumba", "shekara", "agusta", "asabar", "daƙiƙa", "lahadi", "laraba", "oktoba", "talata", "kwana", "maris", "minti", "gobe", "jiya", "mako", "mayu", "wata", "yuli", "yuni", "afi", "agu", "alh", "asa", "awa", "dis", "fab", "gmt", "jan", "jum", "lah", "lar", "lit", "mar", "may", "now", "nuw", "okt", "sat", "tal", "utc", "yau", "yul", "yun", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})

	ha_GH_Locale = merge(&ha_Locale, LocaleData{
		Name:      "ha-GH",
		DateOrder: "DMY",
	})

	ha_NE_Locale = merge(&ha_Locale, LocaleData{
		Name:      "ha-NE",
		DateOrder: "DMY",
	})
}
