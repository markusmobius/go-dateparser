// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sr_Latn_Locale = merge(nil, LocaleData{
	Name:      "sr-Latn",
	DateOrder: "DMY.",
	Charset:   []rune(`bcdefgijklnorstuvzćčš`),
	Translations: map[string]string{
		"ponedeljak": "monday",
		"pre podne":  "am",
		"septembar":  "september",
		"cetvrtak":   "thursday",
		"decembar":   "december",
		"novembar":   "november",
		"po podne":   "pm",
		"februar":    "february",
		"nedelja":    "week",
		"oktobar":    "october",
		"avgust":     "august",
		"godina":     "year",
		"januar":     "january",
		"sekund":     "second",
		"subota":     "saturday",
		"utorak":     "tuesday",
		"april":      "april",
		"mesec":      "month",
		"minut":      "minute",
		"petak":      "friday",
		"sreda":      "wednesday",
		"mart":       "march",
		"apr":        "april",
		"avg":        "august",
		"cet":        "thursday",
		"dan":        "day",
		"dec":        "december",
		"feb":        "february",
		"gmt":        "gmt",
		"god":        "year",
		"jan":        "january",
		"jul":        "july",
		"jun":        "june",
		"maj":        "may",
		"mar":        "march",
		"mes":        "month",
		"min":        "minute",
		"ned":        "week",
		"nov":        "november",
		"okt":        "october",
		"pet":        "friday",
		"pon":        "monday",
		"sat":        "hour",
		"sek":        "second",
		"sep":        "september",
		"sre":        "wednesday",
		"sub":        "saturday",
		"utc":        "utc",
		"uto":        "tuesday",
		"am":         "am",
		"pm":         "pm",
		" ":          " ",
		"'":          "",
		"+":          "+",
		",":          "",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"c":          "hour",
		"d":          "day",
		"g":          "year",
		"m":          "month",
		"n":          "week",
		"s":          "second",
		"z":          "z",
		"|":          "",
	},
	RelativeType: map[string]string{
		"sledece nedelje": "in 1 week",
		"sledeceg meseca": "in 1 month",
		"prosle nedelje":  "1 week ago",
		"proslog meseca":  "1 month ago",
		"sledece godine":  "in 1 year",
		"prosle godine":   "1 year ago",
		"ove nedelje":     "0 week ago",
		"ovog meseca":     "0 month ago",
		"ovog minuta":     "0 minute ago",
		"ove godine":      "0 year ago",
		"ovog sata":       "0 hour ago",
		"danas":           "0 day ago",
		"sutra":           "in 1 day",
		"juce":            "1 day ago",
		"sada":            "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pre (\d+) nedelja`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pre (\d+) nedelje`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pre (\d+) sekunde`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pre (\d+) sekundi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pre (\d+) godina`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pre (\d+) godine`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pre (\d+) meseca`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pre (\d+) meseci`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pre (\d+) minuta`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)za (\d+) nedelja`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) nedelju`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) sekundi`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) godina`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) godinu`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) meseci`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) minuta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pre (\d+) dana`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pre (\d+) sata`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pre (\d+) sati`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)za (\d+) mesec`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)pre (\d+) god`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pre (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pre (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pre (\d+) ned`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pre (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) dana`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) sati`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) dan`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) god`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) ned`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) sat`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)pre (\d+) c`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pre (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pre (\d+) g`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pre (\d+) m`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pre (\d+) n`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pre (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) c`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) g`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) m`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) n`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pre \d+ nedelja|pre \d+ nedelje|pre \d+ sekunde|pre \d+ sekundi|pre \d+ godina|pre \d+ godine|pre \d+ meseca|pre \d+ meseci|pre \d+ minuta|za \d+ nedelja|za \d+ nedelju|za \d+ sekundi|za \d+ sekundu|za \d+ godina|za \d+ godinu|za \d+ meseci|za \d+ minuta|pre \d+ dana|pre \d+ sata|pre \d+ sati|za \d+ mesec|za \d+ minut|pre \d+ god|pre \d+ mes|pre \d+ min|pre \d+ ned|pre \d+ sek|za \d+ dana|za \d+ sati|za \d+ dan|za \d+ god|za \d+ mes|za \d+ min|za \d+ ned|za \d+ sat|za \d+ sek|pre \d+ c|pre \d+ d|pre \d+ g|pre \d+ m|pre \d+ n|pre \d+ s|za \d+ c|za \d+ d|za \d+ g|za \d+ m|za \d+ n|za \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pre \d+ nedelja|pre \d+ nedelje|pre \d+ sekunde|pre \d+ sekundi|pre \d+ godina|pre \d+ godine|pre \d+ meseca|pre \d+ meseci|pre \d+ minuta|za \d+ nedelja|za \d+ nedelju|za \d+ sekundi|za \d+ sekundu|za \d+ godina|za \d+ godinu|za \d+ meseci|za \d+ minuta|pre \d+ dana|pre \d+ sata|pre \d+ sati|za \d+ mesec|za \d+ minut|pre \d+ god|pre \d+ mes|pre \d+ min|pre \d+ ned|pre \d+ sek|za \d+ dana|za \d+ sati|za \d+ dan|za \d+ god|za \d+ mes|za \d+ min|za \d+ ned|za \d+ sat|za \d+ sek|pre \d+ c|pre \d+ d|pre \d+ g|pre \d+ m|pre \d+ n|pre \d+ s|za \d+ c|za \d+ d|za \d+ g|za \d+ m|za \d+ n|za \d+ s)$`),
	KnownWords:      []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "novembar", "po podne", "februar", "nedelja", "oktobar", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
})

var sr_Latn_BA_Locale = merge(&sr_Latn_Locale, LocaleData{
	Name:      "sr-Latn-BA",
	DateOrder: "DMY.",
	Translations: map[string]string{
		"prije podne": "am",
		"nedjelja":    "sunday",
		"srijeda":     "wednesday",
		"sept":        "september",
		"sr":          "wednesday",
		"ut":          "tuesday",
	},
	KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "prije podne", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "po podne", "februar", "nedelja", "oktobar", "srijeda", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
})

var sr_Latn_ME_Locale = merge(&sr_Latn_Locale, LocaleData{
	Name:      "sr-Latn-ME",
	DateOrder: "DMY.",
	Translations: map[string]string{
		"prije podne": "am",
		"nedjelja":    "sunday",
		"srijeda":     "wednesday",
		"sept":        "september",
		"sr":          "wednesday",
		"ut":          "tuesday",
	},
	KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "prije podne", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "po podne", "februar", "nedelja", "oktobar", "srijeda", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
})

var sr_Latn_XK_Locale = merge(&sr_Latn_Locale, LocaleData{
	Name:      "sr-Latn-XK",
	DateOrder: "DMY.",
	Translations: map[string]string{
		"sept": "september",
		"sr":   "wednesday",
		"ut":   "tuesday",
	},
	KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "novembar", "po podne", "februar", "nedelja", "oktobar", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
})
