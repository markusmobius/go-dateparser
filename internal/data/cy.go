// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var cy_Locale = merge(nil, LocaleData{
	Name:      "cy",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghilnorstuwyz`),
	Translations: map[string][]string{
		"dydd mercher": {"wednesday"},
		"dydd gwener":  {"friday"},
		"dydd mawrth":  {"tuesday"},
		"dydd sadwrn":  {"saturday"},
		"gorffennaf":   {"july"},
		"dydd llun":    {"monday"},
		"blwyddyn":     {"year"},
		"chwefror":     {"february"},
		"dydd iau":     {"thursday"},
		"dydd sul":     {"sunday"},
		"tachwedd":     {"november"},
		"mehefin":      {"june"},
		"rhagfyr":      {"december"},
		"wythnos":      {"week"},
		"ebrill":       {"april"},
		"eiliad":       {"second"},
		"hydref":       {"october"},
		"ionawr":       {"january"},
		"mawrth":       {"march"},
		"chwef":        {"february"},
		"gorff":        {"july"},
		"munud":        {"minute"},
		"awst":         {"august"},
		"dydd":         {"day"},
		"gwen":         {"friday"},
		"llun":         {"monday"},
		"medi":         {"september"},
		"rhag":         {"december"},
		"tach":         {"november"},
		"awr":          {"hour"},
		"chw":          {"february"},
		"ebr":          {"april"},
		"gmt":          {"gmt"},
		"gor":          {"july"},
		"gwe":          {"friday"},
		"hyd":          {"october"},
		"iau":          {"thursday"},
		"ion":          {"january"},
		"mai":          {"may"},
		"maw":          {"march", "tuesday"},
		"meh":          {"june"},
		"mer":          {"wednesday"},
		"mis":          {"month"},
		"mun":          {"minute"},
		"sad":          {"saturday"},
		"sul":          {"sunday"},
		"utc":          {"utc"},
		"am":           {"am"},
		"bl":           {"year"},
		"pm":           {"pm"},
		"yb":           {"am"},
		"yh":           {"pm"},
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
		"wythnos ddiwethaf": "1 week ago",
		"blwyddyn nesaf":    "in 1 year",
		"yr wythnos hon":    "0 week ago",
		"wythnos nesaf":     "in 1 week",
		"mis diwethaf":      "1 month ago",
		"y funud hon":       "0 minute ago",
		"yr awr hon":        "0 hour ago",
		"mis nesaf":         "in 1 month",
		"y mis hwn":         "0 month ago",
		"llynedd":           "1 year ago",
		"heddiw":            "0 day ago",
		"eleni":             "0 year ago",
		"yfory":             "in 1 day",
		"ddoe":              "1 day ago",
		"nawr":              "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) o flynyddoedd yn ol`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) diwrnod yn ol`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) wythnos yn ol`), "$1 week ago"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) diwrnod`), "in $1 day"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) mlynedd`), "in $1 year"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) wythnos`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) eiliad yn ol`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) eiliad`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) munud yn ol`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) munud`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) awr yn ol`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mis yn ol`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mun yn ol`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) awr`), "in $1 hour"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) mis`), "in $1 month"},
		{regexp.MustCompile(`(?i)ymhen (\d+[.,]?\d*) mun`), "in $1 minute"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* o flynyddoedd yn ol|\d+[.,]?\d* diwrnod yn ol|\d+[.,]?\d* wythnos yn ol|ymhen \d+[.,]?\d* diwrnod|ymhen \d+[.,]?\d* mlynedd|ymhen \d+[.,]?\d* wythnos|\d+[.,]?\d* eiliad yn ol|ymhen \d+[.,]?\d* eiliad|\d+[.,]?\d* munud yn ol|ymhen \d+[.,]?\d* munud|\d+[.,]?\d* awr yn ol|\d+[.,]?\d* mis yn ol|\d+[.,]?\d* mun yn ol|ymhen \d+[.,]?\d* awr|ymhen \d+[.,]?\d* mis|ymhen \d+[.,]?\d* mun)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* o flynyddoedd yn ol|\d+[.,]?\d* diwrnod yn ol|\d+[.,]?\d* wythnos yn ol|ymhen \d+[.,]?\d* diwrnod|ymhen \d+[.,]?\d* mlynedd|ymhen \d+[.,]?\d* wythnos|\d+[.,]?\d* eiliad yn ol|ymhen \d+[.,]?\d* eiliad|\d+[.,]?\d* munud yn ol|ymhen \d+[.,]?\d* munud|\d+[.,]?\d* awr yn ol|\d+[.,]?\d* mis yn ol|\d+[.,]?\d* mun yn ol|ymhen \d+[.,]?\d* awr|ymhen \d+[.,]?\d* mis|ymhen \d+[.,]?\d* mun)$`),
	KnownWords:      []string{"wythnos ddiwethaf", "blwyddyn nesaf", "yr wythnos hon", "wythnos nesaf", "dydd mercher", "mis diwethaf", "dydd gwener", "dydd mawrth", "dydd sadwrn", "y funud hon", "gorffennaf", "yr awr hon", "dydd llun", "mis nesaf", "y mis hwn", "blwyddyn", "chwefror", "dydd iau", "dydd sul", "tachwedd", "llynedd", "mehefin", "rhagfyr", "wythnos", "ebrill", "eiliad", "heddiw", "hydref", "ionawr", "mawrth", "chwef", "eleni", "gorff", "munud", "yfory", "awst", "ddoe", "dydd", "gwen", "llun", "medi", "nawr", "rhag", "tach", "awr", "chw", "ebr", "gmt", "gor", "gwe", "hyd", "iau", "ion", "mai", "maw", "meh", "mer", "mis", "mun", "sad", "sul", "utc", "am", "bl", "pm", "yb", "yh", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
