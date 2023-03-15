// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lt_Locale = merge(nil, LocaleData{
	Name:      "lt",
	DateOrder: "YMD",
	Charset:   []rune(`bcdeghijklnorstuvyząčėęįšūž`),
	Translations: map[string][]string{
		"ketvirtadienis": {"thursday"},
		"penktadienis":   {"friday"},
		"treciadienis":   {"wednesday"},
		"antradienis":    {"tuesday"},
		"pirmadienis":    {"monday"},
		"sekmadienis":    {"sunday"},
		"sestadienis":    {"saturday"},
		"balandzio":      {"april"},
		"lapkricio":      {"november"},
		"lapkritis":      {"november"},
		"priespiet":      {"am"},
		"rugpjucio":      {"august"},
		"rugpjutis":      {"august"},
		"balandis":       {"april"},
		"birzelio":       {"june"},
		"birzelis":       {"june"},
		"gruodzio":       {"december"},
		"rugsejis":       {"september"},
		"geguzes":        {"may"},
		"gruodis":        {"december"},
		"rugsejo":        {"september"},
		"savaite":        {"week"},
		"sekunde":        {"second"},
		"valanda":        {"hour"},
		"vasario":        {"february"},
		"vasaris":        {"february"},
		"geguze":         {"may"},
		"liepos":         {"july"},
		"minute":         {"minute"},
		"popiet":         {"pm"},
		"sausio":         {"january"},
		"sausis":         {"january"},
		"spalio":         {"october"},
		"spalis":         {"october"},
		"diena":          {"day"},
		"gruod":          {"december"},
		"kovas":          {"march"},
		"lapkr":          {"november"},
		"liepa":          {"july"},
		"menuo":          {"month"},
		"metai":          {"year"},
		"birz":           {"june"},
		"kovo":           {"march"},
		"liep":           {"july"},
		"rugp":           {"august"},
		"rugs":           {"september"},
		"saus":           {"january"},
		"spal":           {"october"},
		"bal":            {"april"},
		"geg":            {"may"},
		"gmt":            {"gmt"},
		"kov":            {"march"},
		"men":            {"month"},
		"min":            {"minute"},
		"sav":            {"week"},
		"sek":            {"second"},
		"utc":            {"utc"},
		"val":            {"hour"},
		"vas":            {"february"},
		"am":             {"am"},
		"an":             {"tuesday"},
		"kt":             {"thursday"},
		"pm":             {"pm"},
		"pn":             {"friday"},
		"pr":             {"monday"},
		"sk":             {"sunday"},
		"st":             {"saturday"},
		"tr":             {"wednesday"},
		" ":              {" "},
		"'":              {""},
		"+":              {"+"},
		",":              {""},
		"-":              {"-"},
		".":              {"."},
		"/":              {"/"},
		":":              {":"},
		";":              {""},
		"@":              {""},
		"[":              {""},
		"]":              {""},
		"d":              {"day"},
		"h":              {"hour"},
		"m":              {"year"},
		"s":              {"second"},
		"z":              {"z"},
		"|":              {""},
	},
	RelativeType: map[string]string{
		"praejusiais metais": "1 year ago",
		"praejusia savaite":  "1 week ago",
		"praejusi menesi":    "1 month ago",
		"kitais metais":      "in 1 year",
		"kita savaite":       "in 1 week",
		"siais metais":       "0 year ago",
		"kita menesi":        "in 1 month",
		"sia savaite":        "0 week ago",
		"sia valanda":        "0 hour ago",
		"sia minute":         "0 minute ago",
		"si menesi":          "0 month ago",
		"siandien":           "0 day ago",
		"dabar":              "0 second ago",
		"rytoj":              "in 1 day",
		"vakar":              "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) sekundziu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) savaiciu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) menesiu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) minuciu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) savaite`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) sekunde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) valanda`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) valandu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) sekundziu`), "in $1 second"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) menesi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) minute`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) savaiciu`), "in $1 week"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) savaites`), "in $1 week"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) sekundes`), "in $1 second"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) valandos`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) diena`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) dienu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) metus`), "$1 year ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) menesio`), "in $1 month"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) menesiu`), "in $1 month"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) minuciu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) minutes`), "in $1 minute"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) valandu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) metu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) dienos`), "in $1 day"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) men`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) sav`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) val`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) dienu`), "in $1 day"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) metu`), "in $1 year"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) m`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pries (\d+[.,]?\d*) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) men`), "in $1 month"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) sav`), "in $1 week"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) val`), "in $1 hour"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) m`), "in $1 year"},
		{regexp.MustCompile(`(?i)po (\d+[.,]?\d*) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pries \d+[.,]?\d* sekundziu|pries \d+[.,]?\d* savaiciu|pries \d+[.,]?\d* menesiu|pries \d+[.,]?\d* minuciu|pries \d+[.,]?\d* savaite|pries \d+[.,]?\d* sekunde|pries \d+[.,]?\d* valanda|pries \d+[.,]?\d* valandu|po \d+[.,]?\d* sekundziu|pries \d+[.,]?\d* menesi|pries \d+[.,]?\d* minute|po \d+[.,]?\d* savaiciu|po \d+[.,]?\d* savaites|po \d+[.,]?\d* sekundes|po \d+[.,]?\d* valandos|pries \d+[.,]?\d* diena|pries \d+[.,]?\d* dienu|pries \d+[.,]?\d* metus|po \d+[.,]?\d* menesio|po \d+[.,]?\d* menesiu|po \d+[.,]?\d* minuciu|po \d+[.,]?\d* minutes|po \d+[.,]?\d* valandu|pries \d+[.,]?\d* metu|po \d+[.,]?\d* dienos|pries \d+[.,]?\d* men|pries \d+[.,]?\d* min|pries \d+[.,]?\d* sav|pries \d+[.,]?\d* sek|pries \d+[.,]?\d* val|po \d+[.,]?\d* dienu|po \d+[.,]?\d* metu|pries \d+[.,]?\d* d|pries \d+[.,]?\d* m|pries \d+[.,]?\d* s|po \d+[.,]?\d* men|po \d+[.,]?\d* min|po \d+[.,]?\d* sav|po \d+[.,]?\d* sek|po \d+[.,]?\d* val|po \d+[.,]?\d* d|po \d+[.,]?\d* m|po \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pries \d+[.,]?\d* sekundziu|pries \d+[.,]?\d* savaiciu|pries \d+[.,]?\d* menesiu|pries \d+[.,]?\d* minuciu|pries \d+[.,]?\d* savaite|pries \d+[.,]?\d* sekunde|pries \d+[.,]?\d* valanda|pries \d+[.,]?\d* valandu|po \d+[.,]?\d* sekundziu|pries \d+[.,]?\d* menesi|pries \d+[.,]?\d* minute|po \d+[.,]?\d* savaiciu|po \d+[.,]?\d* savaites|po \d+[.,]?\d* sekundes|po \d+[.,]?\d* valandos|pries \d+[.,]?\d* diena|pries \d+[.,]?\d* dienu|pries \d+[.,]?\d* metus|po \d+[.,]?\d* menesio|po \d+[.,]?\d* menesiu|po \d+[.,]?\d* minuciu|po \d+[.,]?\d* minutes|po \d+[.,]?\d* valandu|pries \d+[.,]?\d* metu|po \d+[.,]?\d* dienos|pries \d+[.,]?\d* men|pries \d+[.,]?\d* min|pries \d+[.,]?\d* sav|pries \d+[.,]?\d* sek|pries \d+[.,]?\d* val|po \d+[.,]?\d* dienu|po \d+[.,]?\d* metu|pries \d+[.,]?\d* d|pries \d+[.,]?\d* m|pries \d+[.,]?\d* s|po \d+[.,]?\d* men|po \d+[.,]?\d* min|po \d+[.,]?\d* sav|po \d+[.,]?\d* sek|po \d+[.,]?\d* val|po \d+[.,]?\d* d|po \d+[.,]?\d* m|po \d+[.,]?\d* s)$`),
	KnownWords:      []string{"praejusiais metais", "praejusia savaite", "praejusi menesi", "ketvirtadienis", "kitais metais", "kita savaite", "penktadienis", "siais metais", "treciadienis", "antradienis", "kita menesi", "pirmadienis", "sekmadienis", "sestadienis", "sia savaite", "sia valanda", "sia minute", "balandzio", "lapkricio", "lapkritis", "priespiet", "rugpjucio", "rugpjutis", "si menesi", "balandis", "birzelio", "birzelis", "gruodzio", "rugsejis", "siandien", "geguzes", "gruodis", "rugsejo", "savaite", "sekunde", "valanda", "vasario", "vasaris", "geguze", "liepos", "minute", "popiet", "sausio", "sausis", "spalio", "spalis", "dabar", "diena", "gruod", "kovas", "lapkr", "liepa", "menuo", "metai", "rytoj", "vakar", "birz", "kovo", "liep", "rugp", "rugs", "saus", "spal", "bal", "geg", "gmt", "kov", "men", "min", "sav", "sek", "utc", "val", "vas", "am", "an", "kt", "pm", "pn", "pr", "sk", "st", "tr", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "m", "s", "z", "|"},
})
