// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var cs_Locale = merge(nil, LocaleData{
	Name:                  "cs",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdeghijklnorstuvyzáíúýčěřšů`),
	SentenceSplitterGroup: 1,
	Translations: map[string]string{
		"priblizne": "",
		"listopadu": "november",
		"prosince":  "december",
		"prosinec":  "december",
		"hodinami":  "hour",
		"cervence":  "july",
		"cervenec":  "july",
		"listopad":  "november",
		"pondeli":   "monday",
		"sekunda":   "second",
		"sekundy":   "second",
		"vterina":   "second",
		"vteriny":   "second",
		"ctvrtek":   "thursday",
		"hodina":    "hour",
		"hodinu":    "hour",
		"hodiny":    "hour",
		"cerven":    "june",
		"cervna":    "june",
		"brezen":    "march",
		"brezna":    "march",
		"kveten":    "may",
		"kvetna":    "may",
		"minuta":    "minute",
		"minuty":    "minute",
		"mesice":    "month",
		"mesicu":    "month",
		"sobota":    "saturday",
		"sobotu":    "saturday",
		"sekund":    "second",
		"vterin":    "second",
		"nedele":    "sunday",
		"nedeli":    "sunday",
		"streda":    "wednesday",
		"stredu":    "wednesday",
		"duben":     "april",
		"dubna":     "april",
		"srpen":     "august",
		"srpna":     "august",
		"unora":     "february",
		"patek":     "friday",
		"hodin":     "hour",
		"leden":     "january",
		"ledna":     "january",
		"minut":     "minute",
		"mesic":     "month",
		"rijen":     "october",
		"rijna":     "october",
		"utery":     "tuesday",
		"tyden":     "week",
		"tydnu":     "week",
		"tydny":     "week",
		"pred":      "ago",
		"unor":      "february",
		"zari":      "september",
		"roku":      "year",
		"roky":      "year",
		"dop":       "am",
		"dub":       "april",
		"srp":       "august",
		"den":       "day",
		"dnu":       "day",
		"dny":       "day",
		"pro":       "december",
		"uno":       "february",
		"pat":       "friday",
		"gmt":       "gmt",
		"led":       "january",
		"crc":       "july",
		"cvc":       "july",
		"cer":       "june",
		"cvn":       "june",
		"bre":       "march",
		"kve":       "may",
		"min":       "minute",
		"pon":       "monday",
		"mes":       "month",
		"lis":       "november",
		"rij":       "october",
		"odp":       "pm",
		"sob":       "saturday",
		"zar":       "september",
		"ned":       "sunday",
		"ctv":       "thursday",
		"ute":       "tuesday",
		"utc":       "utc",
		"str":       "wednesday",
		"tyd":       "week",
		"rok":       "year",
		"am":        "am",
		"pa":        "friday",
		"ve":        "in",
		"po":        "monday",
		"pm":        "pm",
		"so":        "saturday",
		"ne":        "sunday",
		"ct":        "thursday",
		"ut":        "tuesday",
		"st":        "wednesday",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"h":         "hour",
		"v":         "in",
		"s":         "second",
		"r":         "year",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"minuly mesic": "1 month ago",
		"minuly tyden": "1 week ago",
		"pristi mesic": "in 1 month",
		"pristi tyden": "in 1 week",
		"tuto hodinu":  "0 hour ago",
		"tuto minutu":  "0 minute ago",
		"tento mesic":  "0 month ago",
		"tento tyden":  "0 week ago",
		"predevcirem":  "2 day ago",
		"minuly tyd":   "1 week ago",
		"minuly rok":   "1 year ago",
		"pristi tyd":   "in 1 week",
		"pristi rok":   "in 1 year",
		"tento tyd":    "0 week ago",
		"tento rok":    "0 year ago",
		"vcera":        "1 day ago",
		"zitra":        "in 1 day",
		"dnes":         "0 day ago",
		"nyni":         "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)pred (\d+) sekundami`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinami`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutami`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sekundou`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodinou`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutou`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesicem`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesici`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydnem`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydny`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) rokem`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) dnem`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) lety`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodinu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) mesicu`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) dny`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tyd`), "$1 week ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodin`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) mesic`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) tyden`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) tydnu`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) l`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) r`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) den`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) dni`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) tyd`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) let`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) rok`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) l`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) r`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pred \d+ sekundami|pred \d+ hodinami|pred \d+ minutami|pred \d+ sekundou|pred \d+ hodinou|pred \d+ mesicem|pred \d+ minutou|pred \d+ mesici|pred \d+ tydnem|pred \d+ rokem|pred \d+ tydny|za \d+ sekundu|pred \d+ dnem|pred \d+ lety|za \d+ hodinu|za \d+ mesicu|za \d+ minutu|za \d+ sekund|pred \d+ dny|pred \d+ mes|pred \d+ min|pred \d+ tyd|za \d+ hodin|za \d+ mesic|za \d+ minut|za \d+ tyden|za \d+ tydnu|pred \d+ h|pred \d+ l|pred \d+ r|pred \d+ s|za \d+ den|za \d+ dni|za \d+ let|za \d+ mes|za \d+ min|za \d+ rok|za \d+ tyd|za \d+ h|za \d+ l|za \d+ r|za \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+ sekundami|pred \d+ hodinami|pred \d+ minutami|pred \d+ sekundou|pred \d+ hodinou|pred \d+ mesicem|pred \d+ minutou|pred \d+ mesici|pred \d+ tydnem|pred \d+ rokem|pred \d+ tydny|za \d+ sekundu|pred \d+ dnem|pred \d+ lety|za \d+ hodinu|za \d+ mesicu|za \d+ minutu|za \d+ sekund|pred \d+ dny|pred \d+ mes|pred \d+ min|pred \d+ tyd|za \d+ hodin|za \d+ mesic|za \d+ minut|za \d+ tyden|za \d+ tydnu|pred \d+ h|pred \d+ l|pred \d+ r|pred \d+ s|za \d+ den|za \d+ dni|za \d+ let|za \d+ mes|za \d+ min|za \d+ rok|za \d+ tyd|za \d+ h|za \d+ l|za \d+ r|za \d+ s)$`),
	KnownWords:      []string{"minuly mesic", "minuly tyden", "pristi mesic", "pristi tyden", "predevcirem", "tento mesic", "tento tyden", "tuto hodinu", "tuto minutu", "minuly rok", "minuly tyd", "pristi rok", "pristi tyd", "listopadu", "priblizne", "tento rok", "tento tyd", "cervence", "cervenec", "hodinami", "listopad", "prosince", "prosinec", "ctvrtek", "pondeli", "sekunda", "sekundy", "vterina", "vteriny", "brezen", "brezna", "cerven", "cervna", "hodina", "hodinu", "hodiny", "kveten", "kvetna", "mesice", "mesicu", "minuta", "minuty", "nedele", "nedeli", "sekund", "sobota", "sobotu", "streda", "stredu", "vterin", "duben", "dubna", "hodin", "leden", "ledna", "mesic", "minut", "patek", "rijen", "rijna", "srpen", "srpna", "tyden", "tydnu", "tydny", "unora", "utery", "vcera", "zitra", "dnes", "nyni", "pred", "roku", "roky", "unor", "zari", "bre", "cer", "crc", "ctv", "cvc", "cvn", "den", "dnu", "dny", "dop", "dub", "gmt", "kve", "led", "lis", "mes", "min", "ned", "odp", "pat", "pon", "pro", "rij", "rok", "sob", "srp", "str", "tyd", "uno", "utc", "ute", "zar", "am", "ct", "ne", "pa", "pm", "po", "so", "st", "ut", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "r", "s", "v", "z", "|"},
})
