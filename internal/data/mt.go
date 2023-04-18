// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	mt_Locale LocaleData
)

func init() {
	mt_Locale = merge(nil, LocaleData{
		Name:      "mt",
		DateOrder: "DMY",
		Charset:   []rune(`-bcdefghijklnorstuvwxzċġħ`),
		Translations: map[string][]string{
			"il-gimgħa": {"friday"},
			"it-tlieta": {"tuesday"},
			"settembru": {"september"},
			"dicembru":  {"december"},
			"il-ħamis":  {"thursday"},
			"it-tnejn":  {"monday"},
			"l-erbgħa":  {"wednesday"},
			"novembru":  {"november"},
			"awwissu":   {"august"},
			"il-ħadd":   {"sunday"},
			"is-sibt":   {"saturday"},
			"ottubru":   {"october"},
			"sekonda":   {"second"},
			"gimgħa":    {"week"},
			"jannar":    {"january"},
			"minuta":    {"minute"},
			"siegħa":    {"hour"},
			"april":     {"april"},
			"gunju":     {"june"},
			"lulju":     {"july"},
			"marzu":     {"march"},
			"mejju":     {"may"},
			"xahar":     {"month"},
			"frar":      {"february"},
			"sena":      {"year"},
			"apr":       {"april"},
			"aww":       {"august"},
			"dic":       {"december"},
			"erb":       {"wednesday"},
			"fra":       {"february"},
			"gim":       {"friday"},
			"gmt":       {"gmt"},
			"gun":       {"june"},
			"jan":       {"january"},
			"jum":       {"day"},
			"lul":       {"july"},
			"mar":       {"march"},
			"mej":       {"may"},
			"nov":       {"november"},
			"ott":       {"october"},
			"set":       {"september"},
			"sib":       {"saturday"},
			"tli":       {"tuesday"},
			"tne":       {"monday"},
			"utc":       {"utc"},
			"ħad":       {"sunday"},
			"ħam":       {"thursday"},
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
			"il-gimgħa li għaddiet": "1 week ago",
			"is-sena li għaddiet":   "1 year ago",
			"il-gimgħa d-dieħla":    "in 1 week",
			"ix-xahar id-dieħel":    "in 1 month",
			"ix-xahar li għadda":    "1 month ago",
			"is-sena d-dieħla":      "in 1 year",
			"din il-gimgħa":         "0 week ago",
			"dan ix-xahar":          "0 month ago",
			"din is-sena":           "0 year ago",
			"this minute":           "0 minute ago",
			"this hour":             "0 hour ago",
			"ilbieraħ":              "1 day ago",
			"għada":                 "in 1 day",
			"illum":                 "0 day ago",
			"now":                   "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sena ilu`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) snin ilu`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* sena ilu|\d+[.,]?\d* snin ilu)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* sena ilu|\d+[.,]?\d* snin ilu)$`),
		KnownWords:      []string{"il-gimgħa li għaddiet", "is-sena li għaddiet", "il-gimgħa d-dieħla", "ix-xahar id-dieħel", "ix-xahar li għadda", "is-sena d-dieħla", "din il-gimgħa", "dan ix-xahar", "din is-sena", "this minute", "il-gimgħa", "it-tlieta", "settembru", "this hour", "dicembru", "il-ħamis", "ilbieraħ", "it-tnejn", "l-erbgħa", "novembru", "awwissu", "il-ħadd", "is-sibt", "ottubru", "sekonda", "gimgħa", "jannar", "minuta", "siegħa", "april", "gunju", "għada", "illum", "lulju", "marzu", "mejju", "xahar", "frar", "sena", "apr", "aww", "dic", "erb", "fra", "gim", "gmt", "gun", "jan", "jum", "lul", "mar", "mej", "nov", "now", "ott", "set", "sib", "tli", "tne", "utc", "ħad", "ħam", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
