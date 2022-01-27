// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bg_Locale = merge(nil, LocaleData{
	Name:      "bg",
	DateOrder: "DMY 'Г'.",
	Charset:   []rune("+,-./;@[]cgtuz|абвгдезиклмнопрстуфхцчщъюя"),
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)един(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)два(\z|[^\pL\pM\d]|_)`), "${1}2${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)три(\z|[^\pL\pM\d]|_)`), "${1}3${2}"},
	},
	Translations: map[string]string{
		"понеделник": "monday",
		"септември":  "september",
		"четвъртък":  "thursday",
		"декември":   "december",
		"февруари":   "february",
		"октомври":   "october",
		"ноември":    "november",
		"секунда":    "second",
		"секунди":    "second",
		"вторник":    "tuesday",
		"седмица":    "week",
		"седмици":    "week",
		"август":     "august",
		"януари":     "january",
		"минута":     "minute",
		"минути":     "minute",
		"месеци":     "month",
		"събота":     "saturday",
		"неделя":     "sunday",
		"година":     "year",
		"години":     "year",
		"върху":      "",
		"около":      "",
		"преди":      "ago",
		"април":      "april",
		"петък":      "friday",
		"подир":      "in",
		"после":      "in",
		"месец":      "month",
		"септм":      "september",
		"сряда":      "wednesday",
		"проб":       "am",
		"дена":       "day",
		"часа":       "hour",
		"след":       "in",
		"март":       "march",
		"слоб":       "pm",
		"септ":       "september",
		"седм":       "week",
		"под":        "",
		"апр":        "april",
		"авг":        "august",
		"ден":        "day",
		"дни":        "day",
		"дек":        "december",
		"фев":        "february",
		"gmt":        "gmt",
		"час":        "hour",
		"яну":        "january",
		"юли":        "july",
		"юни":        "june",
		"маи":        "may",
		"мин":        "minute",
		"пон":        "monday",
		"мес":        "month",
		"ное":        "november",
		"окт":        "october",
		"сек":        "second",
		"сеп":        "september",
		"вто":        "tuesday",
		"utc":        "utc",
		"сря":        "wednesday",
		"сед":        "week",
		"год":        "year",
		"до":         "",
		"на":         "",
		"от":         "",
		"am":         "am",
		"ап":         "april",
		"фв":         "february",
		"пт":         "friday",
		"ян":         "january",
		"юл":         "july",
		"юн":         "june",
		"пн":         "monday",
		"pm":         "pm",
		"сб":         "saturday",
		"нд":         "sunday",
		"чт":         "thursday",
		"вт":         "tuesday",
		"ср":         "wednesday",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"д":          "day",
		"ч":          "hour",
		"м":          "month",
		"с":          "week",
		"г":          "year",
		"z":          "z",
	},
	RelativeType: map[string]string{
		"предходната седмица": "1 week ago",
		"следващата седмица":  "in 1 week",
		"след 1 десетилетие":  "in 10 year",
		"преди десетилетие":   "10 year ago",
		"следващата година":   "in 1 year",
		"миналата седмица":    "1 week ago",
		"предходен месец":     "1 month ago",
		"миналата година":     "1 year ago",
		"в тази минута":       "0 minute ago",
		"преди седмица":       "1 week ago",
		"следващ месец":       "in 1 month",
		"тази седмица":        "0 week ago",
		"преди година":        "1 year ago",
		"тази година":         "0 year ago",
		"в този час":          "0 hour ago",
		"този месец":          "0 month ago",
		"следв седм":          "in 1 week",
		"тази седм":           "0 week ago",
		"преди ден":           "1 day ago",
		"преди час":           "1 hour ago",
		"следв мес":           "in 1 month",
		"вдругиден":           "in 2 day",
		"този мес":            "0 month ago",
		"мин седм":            "1 week ago",
		"онзи ден":            "2 day ago",
		"след ден":            "in 1 day",
		"след час":            "in 1 hour",
		"мин мес":             "1 month ago",
		"сл седм":             "in 1 week",
		"следв г":             "in 1 year",
		"вчера":               "1 day ago",
		"снощи":               "1 day ago",
		"мин м":               "1 month ago",
		"мин г":               "1 year ago",
		"днес":                "0 day ago",
		"сега":                "0 second ago",
		"утре":                "in 1 day",
		"сл м":                "in 1 month",
		"сл г":                "in 1 year",
		"т м":                 "0 month ago",
		"т г":                 "0 year ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)преди (\d+) секунда`), "$1 second ago"},
		{regexp.MustCompile(`(?i)преди (\d+) секунди`), "$1 second ago"},
		{regexp.MustCompile(`(?i)преди (\d+) седмица`), "$1 week ago"},
		{regexp.MustCompile(`(?i)преди (\d+) седмици`), "$1 week ago"},
		{regexp.MustCompile(`(?i)преди (\d+) минута`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)преди (\d+) минути`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)преди (\d+) месеца`), "$1 month ago"},
		{regexp.MustCompile(`(?i)преди (\d+) година`), "$1 year ago"},
		{regexp.MustCompile(`(?i)преди (\d+) години`), "$1 year ago"},
		{regexp.MustCompile(`(?i)след (\d+) секунда`), "in $1 second"},
		{regexp.MustCompile(`(?i)след (\d+) секунди`), "in $1 second"},
		{regexp.MustCompile(`(?i)след (\d+) седмица`), "in $1 week"},
		{regexp.MustCompile(`(?i)след (\d+) седмици`), "in $1 week"},
		{regexp.MustCompile(`(?i)преди (\d+) месец`), "$1 month ago"},
		{regexp.MustCompile(`(?i)след (\d+) минута`), "in $1 minute"},
		{regexp.MustCompile(`(?i)след (\d+) минути`), "in $1 minute"},
		{regexp.MustCompile(`(?i)след (\d+) месеца`), "in $1 month"},
		{regexp.MustCompile(`(?i)след (\d+) година`), "in $1 year"},
		{regexp.MustCompile(`(?i)след (\d+) години`), "in $1 year"},
		{regexp.MustCompile(`(?i)преди (\d+) часа`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)преди (\d+) седм`), "$1 week ago"},
		{regexp.MustCompile(`(?i)след (\d+) месец`), "in $1 month"},
		{regexp.MustCompile(`(?i)преди (\d+) ден`), "$1 day ago"},
		{regexp.MustCompile(`(?i)преди (\d+) дни`), "$1 day ago"},
		{regexp.MustCompile(`(?i)преди (\d+) час`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)преди (\d+) мин`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)преди (\d+) сек`), "$1 second ago"},
		{regexp.MustCompile(`(?i)след (\d+) часа`), "in $1 hour"},
		{regexp.MustCompile(`(?i)след (\d+) седм`), "in $1 week"},
		{regexp.MustCompile(`(?i)след (\d+) ден`), "in $1 day"},
		{regexp.MustCompile(`(?i)след (\d+) дни`), "in $1 day"},
		{regexp.MustCompile(`(?i)след (\d+) час`), "in $1 hour"},
		{regexp.MustCompile(`(?i)след (\d+) мин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)след (\d+) сек`), "in $1 second"},
		{regexp.MustCompile(`(?i)преди (\d+) ч`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)преди (\d+) м`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пр (\d+) седм`), "$1 week ago"},
		{regexp.MustCompile(`(?i)преди (\d+) г`), "$1 year ago"},
		{regexp.MustCompile(`(?i)сл (\d+) седм`), "in $1 week"},
		{regexp.MustCompile(`(?i)пр (\d+) мин`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)пр (\d+) сек`), "$1 second ago"},
		{regexp.MustCompile(`(?i)след (\d+) ч`), "in $1 hour"},
		{regexp.MustCompile(`(?i)сл (\d+) мин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)след (\d+) м`), "in $1 month"},
		{regexp.MustCompile(`(?i)сл (\d+) сек`), "in $1 second"},
		{regexp.MustCompile(`(?i)след (\d+) г`), "in $1 year"},
		{regexp.MustCompile(`(?i)пр (\d+) д`), "$1 day ago"},
		{regexp.MustCompile(`(?i)пр (\d+) ч`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)пр (\d+) м`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пр (\d+) г`), "$1 year ago"},
		{regexp.MustCompile(`(?i)сл (\d+) д`), "in $1 day"},
		{regexp.MustCompile(`(?i)сл (\d+) ч`), "in $1 hour"},
		{regexp.MustCompile(`(?i)сл (\d+) м`), "in $1 month"},
		{regexp.MustCompile(`(?i)сл (\d+) г`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(преди \d+ седмица|преди \d+ седмици|преди \d+ секунда|преди \d+ секунди|преди \d+ година|преди \d+ години|преди \d+ месеца|преди \d+ минута|преди \d+ минути|след \d+ седмица|след \d+ седмици|след \d+ секунда|след \d+ секунди|преди \d+ месец|след \d+ година|след \d+ години|след \d+ месеца|след \d+ минута|след \d+ минути|преди \d+ седм|преди \d+ часа|след \d+ месец|преди \d+ ден|преди \d+ дни|преди \d+ мин|преди \d+ сек|преди \d+ час|след \d+ седм|след \d+ часа|след \d+ ден|след \d+ дни|след \d+ мин|след \d+ сек|след \d+ час|пр \d+ седм|преди \d+ г|преди \d+ м|преди \d+ ч|сл \d+ седм|пр \d+ мин|пр \d+ сек|сл \d+ мин|сл \d+ сек|след \d+ г|след \d+ м|след \d+ ч|пр \d+ г|пр \d+ д|пр \d+ м|пр \d+ ч|сл \d+ г|сл \d+ д|сл \d+ м|сл \d+ ч)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(преди \d+ седмица|преди \d+ седмици|преди \d+ секунда|преди \d+ секунди|преди \d+ година|преди \d+ години|преди \d+ месеца|преди \d+ минута|преди \d+ минути|след \d+ седмица|след \d+ седмици|след \d+ секунда|след \d+ секунди|преди \d+ месец|след \d+ година|след \d+ години|след \d+ месеца|след \d+ минута|след \d+ минути|преди \d+ седм|преди \d+ часа|след \d+ месец|преди \d+ ден|преди \d+ дни|преди \d+ мин|преди \d+ сек|преди \d+ час|след \d+ седм|след \d+ часа|след \d+ ден|след \d+ дни|след \d+ мин|след \d+ сек|след \d+ час|пр \d+ седм|преди \d+ г|преди \d+ м|преди \d+ ч|сл \d+ седм|пр \d+ мин|пр \d+ сек|сл \d+ мин|сл \d+ сек|след \d+ г|след \d+ м|след \d+ ч|пр \d+ г|пр \d+ д|пр \d+ м|пр \d+ ч|сл \d+ г|сл \d+ д|сл \d+ м|сл \d+ ч)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(предходната седмица|след 1 десетилетие|следващата седмица|преди десетилетие|следващата година|миналата седмица|миналата година|предходен месец|в тази минута|преди седмица|следващ месец|преди година|тази седмица|тази година|в този час|понеделник|следв седм|този месец|вдругиден|преди ден|преди час|септември|следв мес|тази седм|четвъртък|декември|мин седм|октомври|онзи ден|след ден|след час|този мес|февруари|вторник|мин мес|ноември|седмица|седмици|секунда|секунди|сл седм|следв г|август|година|години|месеци|минута|минути|неделя|събота|януари|април|вчера|върху|месец|мин г|мин м|около|петък|подир|после|преди|септм|снощи|сряда|дена|днес|март|проб|сега|седм|септ|сл г|сл м|след|слоб|утре|часа|gmt|utc|авг|апр|вто|год|дек|ден|дни|маи|мес|мин|ное|окт|под|пон|сед|сек|сеп|сря|т г|т м|фев|час|юли|юни|яну|\+|\.|\[|\]|\||am|pm|ап|вт|до|на|нд|от|пн|пт|сб|ср|фв|чт|юл|юн|ян| |'|,|-|/|:|;|@|z|г|д|м|с|ч)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
