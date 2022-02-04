// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hu_Locale = merge(nil, LocaleData{
	Name:                  "hu",
	DateOrder:             "YMD.",
	Charset:               []rune(`-bcdefghijklnorstuvxzáéóöúüő`),
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)egy(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string]string{
		"masodperccel": "second",
		"masodperctol": "second",
		"masodpercek":  "second",
		"masodpercig":  "second",
		"masodperce":   "second",
		"szeptember":   "september",
		"augusztus":    "august",
		"masodperc":    "second",
		"csutortok":    "thursday",
		"december":     "december",
		"honappal":     "month",
		"november":     "november",
		"vasarnap":     "sunday",
		"ezelott":      "ago",
		"aprilis":      "april",
		"februar":      "february",
		"marcius":      "march",
		"perccel":      "minute",
		"perctol":      "minute",
		"honapja":      "month",
		"honapok":      "month",
		"oktober":      "october",
		"szombat":      "saturday",
		"nappal":       "day",
		"pentek":       "friday",
		"oratol":       "hour",
		"oraval":       "hour",
		"januar":       "january",
		"julius":       "july",
		"junius":       "june",
		"percek":       "minute",
		"percig":       "minute",
		"szerda":       "wednesday",
		"hettel":       "week",
		"-akor":        "",
		"-atol":        "",
		"-ekor":        "",
		"-etol":        "",
		"napja":        "day",
		"napok":        "day",
		"oraig":        "hour",
		"oraja":        "hour",
		"mulva":        "in",
		"majus":        "may",
		"perce":        "minute",
		"hetfo":        "monday",
		"honap":        "month",
		"szept":        "september",
		"hetek":        "week",
		"evvel":        "year",
		"-aig":         "",
		"-eig":         "",
		"-jei":         "",
		"-kor":         "",
		"-tol":         "",
		"viii":         "august",
		"febr":         "february",
		"orak":         "hour",
		"marc":         "march",
		"perc":         "minute",
		"kedd":         "tuesday",
		"hete":         "week",
		"evek":         "year",
		"-ai":          "",
		"-an":          "",
		"-ei":          "",
		"-en":          "",
		"-es":          "",
		"-ig":          "",
		"-je":          "",
		"-ji":          "",
		"-os":          "",
		"apr":          "april",
		"aug":          "august",
		"nap":          "day",
		"dec":          "december",
		"xii":          "december",
		"feb":          "february",
		"gmt":          "gmt",
		"ora":          "hour",
		"jan":          "january",
		"jul":          "july",
		"vii":          "july",
		"jun":          "june",
		"iii":          "march",
		"mar":          "march",
		"maj":          "may",
		"nov":          "november",
		"okt":          "october",
		"szo":          "saturday",
		"vas":          "sunday",
		"utc":          "utc",
		"sze":          "wednesday",
		"het":          "week",
		"eve":          "year",
		"-a":           "",
		"-e":           "",
		"-i":           "",
		"am":           "am",
		"de":           "am",
		"iv":           "april",
		"ii":           "february",
		"vi":           "june",
		"ho":           "month",
		"xi":           "november",
		"du":           "pm",
		"pm":           "pm",
		"mp":           "second",
		"ix":           "september",
		"cs":           "thursday",
		"ev":           "year",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"o":            "hour",
		"i":            "january",
		"v":            "may",
		"p":            "minute",
		"h":            "monday",
		"x":            "october",
		"k":            "tuesday",
		"z":            "z",
	},
	RelativeType: map[string]string{
		"ebben az oraban": "0 hour ago",
		"ebben a percben": "0 minute ago",
		"kovetkezo honap": "in 1 month",
		"kovetkezo het":   "in 1 week",
		"kovetkezo ev":    "in 1 year",
		"elozo honap":     "1 month ago",
		"tegnapelott":     "2 day ago",
		"ez a honap":      "0 month ago",
		"elozo het":       "1 week ago",
		"ez a het":        "0 week ago",
		"ez az ev":        "0 year ago",
		"elozo ev":        "1 year ago",
		"tegnap":          "1 day ago",
		"holnap":          "in 1 day",
		"most":            "0 second ago",
		"ma":              "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) masodperccel ezelott`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) honappal ezelott`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) perccel ezelott`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) masodperc mulva`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) nappal ezelott`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) oraval ezelott`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) hettel ezelott`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) evvel ezelott`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) honap mulva`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) perc mulva`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) nap mulva`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ora mulva`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) het mulva`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ev mulva`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) napja`), "$1 day ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ masodperccel ezelott|\d+ honappal ezelott|\d+ masodperc mulva|\d+ perccel ezelott|\d+ hettel ezelott|\d+ nappal ezelott|\d+ oraval ezelott|\d+ evvel ezelott|\d+ honap mulva|\d+ perc mulva|\d+ het mulva|\d+ nap mulva|\d+ ora mulva|\d+ ev mulva|\d+ napja)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ masodperccel ezelott|\d+ honappal ezelott|\d+ masodperc mulva|\d+ perccel ezelott|\d+ hettel ezelott|\d+ nappal ezelott|\d+ oraval ezelott|\d+ evvel ezelott|\d+ honap mulva|\d+ perc mulva|\d+ het mulva|\d+ nap mulva|\d+ ora mulva|\d+ ev mulva|\d+ napja)$`),
	KnownWords:      []string{"ebben a percben", "ebben az oraban", "kovetkezo honap", "kovetkezo het", "kovetkezo ev", "masodperccel", "masodperctol", "elozo honap", "masodpercek", "masodpercig", "tegnapelott", "ez a honap", "masodperce", "szeptember", "augusztus", "csutortok", "elozo het", "masodperc", "december", "elozo ev", "ez a het", "ez az ev", "honappal", "november", "vasarnap", "aprilis", "ezelott", "februar", "honapja", "honapok", "marcius", "oktober", "perccel", "perctol", "szombat", "hettel", "holnap", "januar", "julius", "junius", "nappal", "oratol", "oraval", "pentek", "percek", "percig", "szerda", "tegnap", "-akor", "-atol", "-ekor", "-etol", "evvel", "hetek", "hetfo", "honap", "majus", "mulva", "napja", "napok", "oraig", "oraja", "perce", "szept", "-aig", "-eig", "-jei", "-kor", "-tol", "evek", "febr", "hete", "kedd", "marc", "most", "orak", "perc", "viii", "-ai", "-an", "-ei", "-en", "-es", "-ig", "-je", "-ji", "-os", "apr", "aug", "dec", "eve", "feb", "gmt", "het", "iii", "jan", "jul", "jun", "maj", "mar", "nap", "nov", "okt", "ora", "sze", "szo", "utc", "vas", "vii", "xii", "-a", "-e", "-i", "am", "cs", "de", "du", "ev", "ho", "ii", "iv", "ix", "ma", "mp", "pm", "vi", "xi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "i", "k", "o", "p", "v", "x", "z", "|"},
})
