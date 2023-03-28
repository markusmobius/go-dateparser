// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	bs_Cyrl_Locale LocaleData
)

func init() {
	bs_Cyrl_Locale = merge(nil, LocaleData{
		Name:      "bs-Cyrl",
		DateOrder: "DMY.",
		Charset:   []rune(`ceghinorstuwzабвгдеиклмнопрстуфцчшјљћ`),
		Translations: map[string][]string{
			"понедељак": {"monday"},
			"пре подне": {"am"},
			"септембар": {"september"},
			"децембар":  {"december"},
			"новембар":  {"november"},
			"четвртак":  {"thursday"},
			"октобар":   {"october"},
			"поподне":   {"pm"},
			"сриједа":   {"wednesday"},
			"фебруар":   {"february"},
			"август":    {"august"},
			"година":    {"year"},
			"недеља":    {"week", "sunday"},
			"секунд":    {"second"},
			"субота":    {"saturday"},
			"уторак":    {"tuesday"},
			"јануар":    {"january"},
			"април":     {"april"},
			"месец":     {"month"},
			"минут":     {"minute"},
			"петак":     {"friday"},
			"март":      {"march"},
			"јули":      {"july"},
			"јуни":      {"june"},
			"gmt":       {"gmt"},
			"utc":       {"utc"},
			"авг":       {"august"},
			"апр":       {"april"},
			"дан":       {"day"},
			"дец":       {"december"},
			"мар":       {"march"},
			"мај":       {"may"},
			"нед":       {"sunday"},
			"нов":       {"november"},
			"окт":       {"october"},
			"пет":       {"friday"},
			"пон":       {"monday"},
			"сеп":       {"september"},
			"сри":       {"wednesday"},
			"суб":       {"saturday"},
			"уто":       {"tuesday"},
			"феб":       {"february"},
			"час":       {"hour"},
			"чет":       {"thursday"},
			"јан":       {"january"},
			"јул":       {"july"},
			"јун":       {"june"},
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
			"следећег месеца": "in 1 month",
			"прошлог месеца":  "1 month ago",
			"следеће године":  "in 1 year",
			"следеће недеље":  "in 1 week",
			"прошле године":   "1 year ago",
			"прошле недеље":   "1 week ago",
			"this minute":     "0 minute ago",
			"овог месеца":     "0 month ago",
			"ове године":      "0 year ago",
			"ове недеље":      "0 week ago",
			"this hour":       "0 hour ago",
			"данас":           "0 day ago",
			"сутра":           "in 1 day",
			"јуче":            "1 day ago",
			"now":             "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) секунди`), "$1 second ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунди`), "in $1 second"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) година`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) годину`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) месеци`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) минута`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) недеља`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) недељу`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) секунд`), "$1 second ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) година`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) годину`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месеци`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минута`), "in $1 minute"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) недеља`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) недељу`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунд`), "in $1 second"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) месец`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) минут`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месец`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минут`), "in $1 minute"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) дана`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) сати`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) дана`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) сати`), "in $1 hour"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) дан`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пре (\d+[.,]?\d*) сат`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) дан`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) сат`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(пре \d+[.,]?\d* секунди|за \d+[.,]?\d* секунди|пре \d+[.,]?\d* година|пре \d+[.,]?\d* годину|пре \d+[.,]?\d* месеци|пре \d+[.,]?\d* минута|пре \d+[.,]?\d* недеља|пре \d+[.,]?\d* недељу|пре \d+[.,]?\d* секунд|за \d+[.,]?\d* година|за \d+[.,]?\d* годину|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* недеља|за \d+[.,]?\d* недељу|за \d+[.,]?\d* секунд|пре \d+[.,]?\d* месец|пре \d+[.,]?\d* минут|за \d+[.,]?\d* месец|за \d+[.,]?\d* минут|пре \d+[.,]?\d* дана|пре \d+[.,]?\d* сати|за \d+[.,]?\d* дана|за \d+[.,]?\d* сати|пре \d+[.,]?\d* дан|пре \d+[.,]?\d* сат|за \d+[.,]?\d* дан|за \d+[.,]?\d* сат)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(пре \d+[.,]?\d* секунди|за \d+[.,]?\d* секунди|пре \d+[.,]?\d* година|пре \d+[.,]?\d* годину|пре \d+[.,]?\d* месеци|пре \d+[.,]?\d* минута|пре \d+[.,]?\d* недеља|пре \d+[.,]?\d* недељу|пре \d+[.,]?\d* секунд|за \d+[.,]?\d* година|за \d+[.,]?\d* годину|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* недеља|за \d+[.,]?\d* недељу|за \d+[.,]?\d* секунд|пре \d+[.,]?\d* месец|пре \d+[.,]?\d* минут|за \d+[.,]?\d* месец|за \d+[.,]?\d* минут|пре \d+[.,]?\d* дана|пре \d+[.,]?\d* сати|за \d+[.,]?\d* дана|за \d+[.,]?\d* сати|пре \d+[.,]?\d* дан|пре \d+[.,]?\d* сат|за \d+[.,]?\d* дан|за \d+[.,]?\d* сат)$`),
		KnownWords:      []string{"следећег месеца", "прошлог месеца", "следеће године", "следеће недеље", "прошле године", "прошле недеље", "this minute", "овог месеца", "ове године", "ове недеље", "this hour", "понедељак", "пре подне", "септембар", "децембар", "новембар", "четвртак", "октобар", "поподне", "сриједа", "фебруар", "август", "година", "недеља", "секунд", "субота", "уторак", "јануар", "април", "данас", "месец", "минут", "петак", "сутра", "март", "јули", "јуни", "јуче", "gmt", "now", "utc", "авг", "апр", "дан", "дец", "мар", "мај", "нед", "нов", "окт", "пет", "пон", "сеп", "сри", "суб", "уто", "феб", "час", "чет", "јан", "јул", "јун", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
