// Code is generated by script; DO NOT EDIT.

package data

import "github.com/markusmobius/go-dateparser/internal/regexp"

var (
	hu_Locale LocaleData
)

func init() {
	hu_Locale = merge(nil, LocaleData{
		Name:                  "hu",
		DateOrder:             "YMD.",
		Charset:               []rune(`-bcdefghijklnorstuvxzáéóöúüő`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])egy(\z|[^\pL\pM\d_])`), "${1}1${2}"},
		},
		Translations: map[string][]string{
			"masodperccel": {"second"},
			"masodperctol": {"second"},
			"masodpercek":  {"second"},
			"masodpercig":  {"second"},
			"masodperce":   {"second"},
			"szeptember":   {"september"},
			"augusztus":    {"august"},
			"csutortok":    {"thursday"},
			"masodperc":    {"second"},
			"december":     {"december"},
			"honappal":     {"month"},
			"november":     {"november"},
			"vasarnap":     {"sunday"},
			"aprilis":      {"april"},
			"ezelott":      {"ago"},
			"februar":      {"february"},
			"honapja":      {"month"},
			"honapok":      {"month"},
			"marcius":      {"march"},
			"oktober":      {"october"},
			"perccel":      {"minute"},
			"perctol":      {"minute"},
			"szombat":      {"saturday"},
			"hettel":       {"week"},
			"januar":       {"january"},
			"julius":       {"july"},
			"junius":       {"june"},
			"nappal":       {"day"},
			"oratol":       {"hour"},
			"oraval":       {"hour"},
			"pentek":       {"friday"},
			"percek":       {"minute"},
			"percig":       {"minute"},
			"szerda":       {"wednesday"},
			"-akor":        {""},
			"-atol":        {""},
			"-ekor":        {""},
			"-etol":        {""},
			"evvel":        {"year"},
			"hetek":        {"week"},
			"hetfo":        {"monday"},
			"honap":        {"month"},
			"majus":        {"may"},
			"mulva":        {"in"},
			"napja":        {"day"},
			"napok":        {"day"},
			"oraig":        {"hour"},
			"oraja":        {"hour"},
			"perce":        {"minute"},
			"szept":        {"september"},
			"-aig":         {""},
			"-eig":         {""},
			"-jei":         {""},
			"-kor":         {""},
			"-tol":         {""},
			"evek":         {"year"},
			"febr":         {"february"},
			"hete":         {"week"},
			"kedd":         {"tuesday"},
			"marc":         {"march"},
			"orak":         {"hour"},
			"perc":         {"minute"},
			"viii":         {"august"},
			"-ai":          {""},
			"-an":          {""},
			"-ei":          {""},
			"-en":          {""},
			"-es":          {""},
			"-ig":          {""},
			"-je":          {""},
			"-ji":          {""},
			"-os":          {""},
			"apr":          {"april"},
			"aug":          {"august"},
			"dec":          {"december"},
			"eve":          {"year"},
			"feb":          {"february"},
			"gmt":          {"gmt"},
			"het":          {"week"},
			"iii":          {"march"},
			"jan":          {"january"},
			"jul":          {"july"},
			"jun":          {"june"},
			"maj":          {"may"},
			"mar":          {"march"},
			"nap":          {"day"},
			"nov":          {"november"},
			"okt":          {"october"},
			"ora":          {"hour"},
			"sze":          {"wednesday"},
			"szo":          {"saturday"},
			"utc":          {"utc"},
			"vas":          {"sunday"},
			"vii":          {"july"},
			"xii":          {"december"},
			"-a":           {""},
			"-e":           {""},
			"-i":           {""},
			"am":           {"am"},
			"cs":           {"thursday"},
			"de":           {"am"},
			"du":           {"pm"},
			"ev":           {"year"},
			"ho":           {"month"},
			"ii":           {"february"},
			"iv":           {"april"},
			"ix":           {"september"},
			"mp":           {"second"},
			"pm":           {"pm"},
			"vi":           {"june"},
			"xi":           {"november"},
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
			"h":            {"monday"},
			"i":            {"january"},
			"k":            {"tuesday"},
			"o":            {"hour"},
			"p":            {"friday", "minute"},
			"v":            {"may", "sunday"},
			"x":            {"october"},
			"z":            {"z"},
			"|":            {""},
		},
		RelativeType: map[string]string{
			"ebben a percben": "0 minute ago",
			"ebben az oraban": "0 hour ago",
			"kovetkezo honap": "in 1 month",
			"kovetkezo het":   "in 1 week",
			"kovetkezo ev":    "in 1 year",
			"elozo honap":     "1 month ago",
			"tegnapelott":     "2 day ago",
			"ez a honap":      "0 month ago",
			"elozo het":       "1 week ago",
			"elozo ev":        "1 year ago",
			"ez a het":        "0 week ago",
			"ez az ev":        "0 year ago",
			"holnap":          "in 1 day",
			"tegnap":          "1 day ago",
			"most":            "0 second ago",
			"ma":              "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) masodperccel ezelott`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) honappal ezelott`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) masodperc mulva`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) perccel ezelott`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hettel ezelott`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) nappal ezelott`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) oraval ezelott`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) evvel ezelott`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) honap mulva`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) perc mulva`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) het mulva`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) nap mulva`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ora mulva`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ev mulva`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) napja`), "$1 day ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* masodperccel ezelott|\d+[.,]?\d* honappal ezelott|\d+[.,]?\d* masodperc mulva|\d+[.,]?\d* perccel ezelott|\d+[.,]?\d* hettel ezelott|\d+[.,]?\d* nappal ezelott|\d+[.,]?\d* oraval ezelott|\d+[.,]?\d* evvel ezelott|\d+[.,]?\d* honap mulva|\d+[.,]?\d* perc mulva|\d+[.,]?\d* het mulva|\d+[.,]?\d* nap mulva|\d+[.,]?\d* ora mulva|\d+[.,]?\d* ev mulva|\d+[.,]?\d* napja)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* masodperccel ezelott|\d+[.,]?\d* honappal ezelott|\d+[.,]?\d* masodperc mulva|\d+[.,]?\d* perccel ezelott|\d+[.,]?\d* hettel ezelott|\d+[.,]?\d* nappal ezelott|\d+[.,]?\d* oraval ezelott|\d+[.,]?\d* evvel ezelott|\d+[.,]?\d* honap mulva|\d+[.,]?\d* perc mulva|\d+[.,]?\d* het mulva|\d+[.,]?\d* nap mulva|\d+[.,]?\d* ora mulva|\d+[.,]?\d* ev mulva|\d+[.,]?\d* napja)$`),
		KnownWords:      []string{"ebben a percben", "ebben az oraban", "kovetkezo honap", "kovetkezo het", "kovetkezo ev", "masodperccel", "masodperctol", "elozo honap", "masodpercek", "masodpercig", "tegnapelott", "ez a honap", "masodperce", "szeptember", "augusztus", "csutortok", "elozo het", "masodperc", "december", "elozo ev", "ez a het", "ez az ev", "honappal", "november", "vasarnap", "aprilis", "ezelott", "februar", "honapja", "honapok", "marcius", "oktober", "perccel", "perctol", "szombat", "hettel", "holnap", "januar", "julius", "junius", "nappal", "oratol", "oraval", "pentek", "percek", "percig", "szerda", "tegnap", "-akor", "-atol", "-ekor", "-etol", "evvel", "hetek", "hetfo", "honap", "majus", "mulva", "napja", "napok", "oraig", "oraja", "perce", "szept", "-aig", "-eig", "-jei", "-kor", "-tol", "evek", "febr", "hete", "kedd", "marc", "most", "orak", "perc", "viii", "-ai", "-an", "-ei", "-en", "-es", "-ig", "-je", "-ji", "-os", "apr", "aug", "dec", "eve", "feb", "gmt", "het", "iii", "jan", "jul", "jun", "maj", "mar", "nap", "nov", "okt", "ora", "sze", "szo", "utc", "vas", "vii", "xii", "-a", "-e", "-i", "am", "cs", "de", "du", "ev", "ho", "ii", "iv", "ix", "ma", "mp", "pm", "vi", "xi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "i", "k", "o", "p", "v", "x", "z", "|"},
	})
}
