// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uz_Cyrl_Locale = merge(nil, LocaleData{
	Name:      "uz-Cyrl",
	DateOrder: "DMY",
	Charset:   []rune(`ceghinorstuzабвгдежзийклмнопрстуфчшэюяўқҳ`),
	Translations: map[string][]string{
		"паишанба": {"thursday"},
		"чоршанба": {"wednesday"},
		"душанба":  {"monday"},
		"сентябр":  {"september"},
		"сешанба":  {"tuesday"},
		"якшанба":  {"sunday"},
		"август":   {"august"},
		"дақиқа":   {"minute"},
		"декабр":   {"december"},
		"октябр":   {"october"},
		"феврал":   {"february"},
		"апрел":    {"april"},
		"ноябр":    {"november"},
		"сония":    {"second"},
		"шанба":    {"saturday"},
		"январ":    {"january"},
		"ҳафта":    {"week"},
		"жума":     {"friday"},
		"март":     {"march"},
		"соат":     {"hour"},
		"gmt":      {"gmt"},
		"utc":      {"utc"},
		"авг":      {"august"},
		"апр":      {"april"},
		"дек":      {"december"},
		"душ":      {"monday"},
		"жум":      {"friday"},
		"иил":      {"year"},
		"июл":      {"july"},
		"июн":      {"june"},
		"кун":      {"day"},
		"маи":      {"may"},
		"мар":      {"march"},
		"ноя":      {"november"},
		"окт":      {"october"},
		"паи":      {"thursday"},
		"сен":      {"september"},
		"сеш":      {"tuesday"},
		"фев":      {"february"},
		"чор":      {"wednesday"},
		"шан":      {"saturday"},
		"якш":      {"sunday"},
		"янв":      {"january"},
		"am":       {"am"},
		"pm":       {"pm"},
		"ои":       {"month"},
		"тк":       {"pm"},
		"то":       {"am"},
		" ":        {" "},
		"'":        {""},
		"+":        {"+"},
		",":        {""},
		"-":        {"-"},
		".":        {"."},
		"/":        {"/"},
		":":        {":"},
		";":        {""},
		"@":        {""},
		"[":        {""},
		"]":        {""},
		"z":        {"z"},
		"|":        {""},
	},
	RelativeType: map[string]string{
		"кеиинги ҳафта": "in 1 week",
		"this minute":   "0 minute ago",
		"кеиинги иил":   "in 1 year",
		"утган ҳафта":   "1 week ago",
		"кеиинги ои":    "in 1 month",
		"this hour":     "0 hour ago",
		"утган иил":     "1 year ago",
		"бу ҳафта":      "0 week ago",
		"утган ои":      "1 month ago",
		"бу иил":        "0 year ago",
		"эртага":        "in 1 day",
		"бу ои":         "0 month ago",
		"бугун":         "0 day ago",
		"ҳозир":         "0 second ago",
		"кеча":          "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дақиқадан сунг`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сониядан сунг`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ҳафтадан сунг`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) дақиқа олдин`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) соатдан сунг`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) иилдан сунг`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) кундан сунг`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сония олдин`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ҳафта олдин`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) оидан сунг`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) соат олдин`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) иил аввал`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) кун олдин`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ои аввал`), "$1 month ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* дақиқадан сунг|\d+[.,]?\d* сониядан сунг|\d+[.,]?\d* ҳафтадан сунг|\d+[.,]?\d* дақиқа олдин|\d+[.,]?\d* соатдан сунг|\d+[.,]?\d* иилдан сунг|\d+[.,]?\d* кундан сунг|\d+[.,]?\d* сония олдин|\d+[.,]?\d* ҳафта олдин|\d+[.,]?\d* оидан сунг|\d+[.,]?\d* соат олдин|\d+[.,]?\d* иил аввал|\d+[.,]?\d* кун олдин|\d+[.,]?\d* ои аввал)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* дақиқадан сунг|\d+[.,]?\d* сониядан сунг|\d+[.,]?\d* ҳафтадан сунг|\d+[.,]?\d* дақиқа олдин|\d+[.,]?\d* соатдан сунг|\d+[.,]?\d* иилдан сунг|\d+[.,]?\d* кундан сунг|\d+[.,]?\d* сония олдин|\d+[.,]?\d* ҳафта олдин|\d+[.,]?\d* оидан сунг|\d+[.,]?\d* соат олдин|\d+[.,]?\d* иил аввал|\d+[.,]?\d* кун олдин|\d+[.,]?\d* ои аввал)$`),
	KnownWords:      []string{"кеиинги ҳафта", "this minute", "кеиинги иил", "утган ҳафта", "кеиинги ои", "this hour", "утган иил", "бу ҳафта", "паишанба", "утган ои", "чоршанба", "душанба", "сентябр", "сешанба", "якшанба", "август", "бу иил", "дақиқа", "декабр", "октябр", "феврал", "эртага", "апрел", "бу ои", "бугун", "ноябр", "сония", "шанба", "январ", "ҳафта", "ҳозир", "жума", "кеча", "март", "соат", "gmt", "utc", "авг", "апр", "дек", "душ", "жум", "иил", "июл", "июн", "кун", "маи", "мар", "ноя", "окт", "паи", "сен", "сеш", "фев", "чор", "шан", "якш", "янв", "am", "pm", "ои", "тк", "то", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
