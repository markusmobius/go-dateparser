// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lv_Locale = merge(nil, LocaleData{
	Name:      "lv",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuvzāēīļšū`),
	Translations: map[string]string{
		"priekspusdiena": "am",
		"ceturtdiena":    "thursday",
		"pecpusdiena":    "pm",
		"piektdiena":     "friday",
		"septembris":     "september",
		"decembris":      "december",
		"februaris":      "february",
		"novembris":      "november",
		"pirmdiena":      "monday",
		"sestdiena":      "saturday",
		"svetdiena":      "sunday",
		"tresdiena":      "wednesday",
		"janvaris":       "january",
		"oktobris":       "october",
		"otrdiena":       "tuesday",
		"sekundes":       "second",
		"aprilis":        "april",
		"augusts":        "august",
		"ceturtd":        "thursday",
		"menesis":        "month",
		"minutes":        "minute",
		"pecpusd":        "pm",
		"prieksp":        "am",
		"stundas":        "hour",
		"julijs":         "july",
		"junijs":         "june",
		"nedela":         "week",
		"piektd":         "friday",
		"diena":          "day",
		"maijs":          "may",
		"marts":          "march",
		"pirmd":          "monday",
		"sestd":          "saturday",
		"svetd":          "sunday",
		"tresd":          "wednesday",
		"febr":           "february",
		"gads":           "year",
		"janv":           "january",
		"otrd":           "tuesday",
		"pecp":           "pm",
		"sept":           "september",
		"apr":            "april",
		"aug":            "august",
		"dec":            "december",
		"gmt":            "gmt",
		"jul":            "july",
		"jun":            "june",
		"men":            "month",
		"min":            "minute",
		"ned":            "week",
		"nov":            "november",
		"okt":            "october",
		"sek":            "second",
		"utc":            "utc",
		"am":             "am",
		"pm":             "pm",
		"st":             "hour",
		" ":              " ",
		"'":              "",
		"+":              "+",
		",":              "",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"d":              "day",
		"g":              "year",
		"h":              "hour",
		"s":              "second",
		"z":              "z",
		"|":              "",
	},
	RelativeType: map[string]string{
		"pagajusaja menesi": "1 month ago",
		"pagajusaja nedela": "1 week ago",
		"nakamaja menesi":   "in 1 month",
		"nakamaja nedela":   "in 1 week",
		"pagajusaja gada":   "1 year ago",
		"nakamaja gada":     "in 1 year",
		"saja menesi":       "0 month ago",
		"saja minute":       "0 minute ago",
		"saja nedela":       "0 week ago",
		"saja stunda":       "0 hour ago",
		"saja gada":         "0 year ago",
		"sodien":            "0 day ago",
		"tagad":             "0 second ago",
		"vakar":             "1 day ago",
		"rit":               "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pirms (\d+) menesiem`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) sekundem`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) sekundes`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) minutem`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) minutes`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) nedelam`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) nedelas`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) stundam`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) stundas`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pec (\d+) menesiem`), "in $1 month"},
		{regexp.MustCompile(`(?i)pec (\d+) sekundem`), "in $1 second"},
		{regexp.MustCompile(`(?i)pec (\d+) sekundes`), "in $1 second"},
		{regexp.MustCompile(`(?i)pirms (\d+) dienam`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) dienas`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) gadiem`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) menesa`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pec (\d+) minutem`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pec (\d+) minutes`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pec (\d+) nedelam`), "in $1 week"},
		{regexp.MustCompile(`(?i)pec (\d+) nedelas`), "in $1 week"},
		{regexp.MustCompile(`(?i)pec (\d+) stundam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pec (\d+) stundas`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pec (\d+) dienam`), "in $1 day"},
		{regexp.MustCompile(`(?i)pec (\d+) dienas`), "in $1 day"},
		{regexp.MustCompile(`(?i)pec (\d+) gadiem`), "in $1 year"},
		{regexp.MustCompile(`(?i)pec (\d+) menesa`), "in $1 month"},
		{regexp.MustCompile(`(?i)pirms (\d+) gada`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) men`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) ned`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pec (\d+) gada`), "in $1 year"},
		{regexp.MustCompile(`(?i)pirms (\d+) st`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pec (\d+) men`), "in $1 month"},
		{regexp.MustCompile(`(?i)pec (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pec (\d+) ned`), "in $1 week"},
		{regexp.MustCompile(`(?i)pec (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)pirms (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) g`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pirms (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pec (\d+) st`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pec (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)pec (\d+) g`), "in $1 year"},
		{regexp.MustCompile(`(?i)pec (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)pec (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pirms \d+ menesiem|pirms \d+ sekundem|pirms \d+ sekundes|pirms \d+ minutem|pirms \d+ minutes|pirms \d+ nedelam|pirms \d+ nedelas|pirms \d+ stundam|pirms \d+ stundas|pec \d+ menesiem|pec \d+ sekundem|pec \d+ sekundes|pirms \d+ dienam|pirms \d+ dienas|pirms \d+ gadiem|pirms \d+ menesa|pec \d+ minutem|pec \d+ minutes|pec \d+ nedelam|pec \d+ nedelas|pec \d+ stundam|pec \d+ stundas|pec \d+ dienam|pec \d+ dienas|pec \d+ gadiem|pec \d+ menesa|pirms \d+ gada|pirms \d+ men|pirms \d+ min|pirms \d+ ned|pirms \d+ sek|pec \d+ gada|pirms \d+ st|pec \d+ men|pec \d+ min|pec \d+ ned|pec \d+ sek|pirms \d+ d|pirms \d+ g|pirms \d+ h|pirms \d+ s|pec \d+ st|pec \d+ d|pec \d+ g|pec \d+ h|pec \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pirms \d+ menesiem|pirms \d+ sekundem|pirms \d+ sekundes|pirms \d+ minutem|pirms \d+ minutes|pirms \d+ nedelam|pirms \d+ nedelas|pirms \d+ stundam|pirms \d+ stundas|pec \d+ menesiem|pec \d+ sekundem|pec \d+ sekundes|pirms \d+ dienam|pirms \d+ dienas|pirms \d+ gadiem|pirms \d+ menesa|pec \d+ minutem|pec \d+ minutes|pec \d+ nedelam|pec \d+ nedelas|pec \d+ stundam|pec \d+ stundas|pec \d+ dienam|pec \d+ dienas|pec \d+ gadiem|pec \d+ menesa|pirms \d+ gada|pirms \d+ men|pirms \d+ min|pirms \d+ ned|pirms \d+ sek|pec \d+ gada|pirms \d+ st|pec \d+ men|pec \d+ min|pec \d+ ned|pec \d+ sek|pirms \d+ d|pirms \d+ g|pirms \d+ h|pirms \d+ s|pec \d+ st|pec \d+ d|pec \d+ g|pec \d+ h|pec \d+ s)$`),
	KnownWords:      []string{"pagajusaja menesi", "pagajusaja nedela", "nakamaja menesi", "nakamaja nedela", "pagajusaja gada", "priekspusdiena", "nakamaja gada", "ceturtdiena", "pecpusdiena", "saja menesi", "saja minute", "saja nedela", "saja stunda", "piektdiena", "septembris", "decembris", "februaris", "novembris", "pirmdiena", "saja gada", "sestdiena", "svetdiena", "tresdiena", "janvaris", "oktobris", "otrdiena", "sekundes", "aprilis", "augusts", "ceturtd", "menesis", "minutes", "pecpusd", "prieksp", "stundas", "julijs", "junijs", "nedela", "piektd", "sodien", "diena", "maijs", "marts", "pirmd", "sestd", "svetd", "tagad", "tresd", "vakar", "febr", "gads", "janv", "otrd", "pecp", "sept", "apr", "aug", "dec", "gmt", "jul", "jun", "men", "min", "ned", "nov", "okt", "rit", "sek", "utc", "am", "pm", "st", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "g", "h", "s", "z", "|"},
})
