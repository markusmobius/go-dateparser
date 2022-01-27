// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ky_Locale = merge(nil, LocaleData{
	Name:      "ky",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]cgtuz|абвгдежзиклмнопрстуфчшыьэюяңүӊө"),
	Translations: map[string]string{
		"түштөн кииинки": "pm",
		"дүишөмбү":       "monday",
		"сентябрь":       "september",
		"жекшемби":       "sunday",
		"беишемби":       "thursday",
		"шеишемби":       "tuesday",
		"шаршемби":       "wednesday",
		"декабрь":        "december",
		"февраль":        "february",
		"октябрь":        "october",
		"апрель":         "april",
		"август":         "august",
		"январь":         "january",
		"ноябрь":         "november",
		"ишемби":         "saturday",
		"секунд":         "second",
		"таңкы":          "am",
		"мүнөт":          "minute",
		"жума":           "friday",
		"саат":           "hour",
		"июль":           "july",
		"июнь":           "june",
		"март":           "march",
		"беиш":           "thursday",
		"шеиш":           "tuesday",
		"шарш":           "wednesday",
		"апта":           "week",
		"апр":            "april",
		"авг":            "august",
		"күн":            "day",
		"дек":            "december",
		"фев":            "february",
		"gmt":            "gmt",
		"янв":            "january",
		"июл":            "july",
		"июн":            "june",
		"мар":            "march",
		"маи":            "may",
		"мүн":            "minute",
		"дүи":            "monday",
		"ноя":            "november",
		"окт":            "october",
		"ишм":            "saturday",
		"сек":            "second",
		"сен":            "september",
		"жек":            "sunday",
		"utc":            "utc",
		"апт":            "week",
		"жыл":            "year",
		"am":             "am",
		"тң":             "am",
		"ст":             "hour",
		"аи":             "month",
		"pm":             "pm",
		"тк":             "pm",
		"'":              "",
		",":              "",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"|":              "",
		" ":              " ",
		"+":              "+",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		"м":              "minute",
		"ж":              "year",
		"z":              "z",
	},
	RelativeType: map[string]string{
		"келерки аптада": "in 1 week",
		"ушул мүнөттө":   "0 minute ago",
		"өткөн аптада":   "1 week ago",
		"ушул саатта":    "0 hour ago",
		"ушул аптада":    "0 week ago",
		"эмдиги аида":    "in 1 month",
		"келерки апт":    "in 1 week",
		"эмдиги жылы":    "in 1 year",
		"өткөн аида":     "1 month ago",
		"өткөн апт":      "1 week ago",
		"бул аида":       "0 month ago",
		"ушул апт":       "0 week ago",
		"былтыр":         "1 year ago",
		"бүгүн":          "0 day ago",
		"быиыл":          "0 year ago",
		"кечээ":          "1 day ago",
		"эртеӊ":          "in 1 day",
		"азыр":           "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) секунддан кииин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) мүнөттөн кииин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) сааттан кииин`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) аптадан кииин`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) секунд мурун`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) күндөн кииин`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) жылдан кииин`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) мүнөт мурун`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) аидан кииин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) саат мурун`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) апта мурун`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) саат кииин`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) күн мурун`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) мүн мурун`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) сек мурун`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) апт мурун`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) жыл мурун`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) күн кииин`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) мүн кииин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) аид кииин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) сек кииин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) апт кииин`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) жыл кииин`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) мүн мурн`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) аи мурун`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) сек мурн`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) мүн киин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) аид киин`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) сек киин`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) аи мурн`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) с мурн`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) с киин`), "in $1 hour"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ секунддан кииин|\d+ мүнөттөн кииин|\d+ аптадан кииин|\d+ сааттан кииин|\d+ жылдан кииин|\d+ күндөн кииин|\d+ секунд мурун|\d+ аидан кииин|\d+ мүнөт мурун|\d+ апта мурун|\d+ саат кииин|\d+ саат мурун|\d+ аид кииин|\d+ апт кииин|\d+ апт мурун|\d+ жыл кииин|\d+ жыл мурун|\d+ күн кииин|\d+ күн мурун|\d+ мүн кииин|\d+ мүн мурун|\d+ сек кииин|\d+ сек мурун|\d+ аи мурун|\d+ аид киин|\d+ мүн киин|\d+ мүн мурн|\d+ сек киин|\d+ сек мурн|\d+ аи мурн|\d+ с киин|\d+ с мурн)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ секунддан кииин|\d+ мүнөттөн кииин|\d+ аптадан кииин|\d+ сааттан кииин|\d+ жылдан кииин|\d+ күндөн кииин|\d+ секунд мурун|\d+ аидан кииин|\d+ мүнөт мурун|\d+ апта мурун|\d+ саат кииин|\d+ саат мурун|\d+ аид кииин|\d+ апт кииин|\d+ апт мурун|\d+ жыл кииин|\d+ жыл мурун|\d+ күн кииин|\d+ күн мурун|\d+ мүн кииин|\d+ мүн мурун|\d+ сек кииин|\d+ сек мурун|\d+ аи мурун|\d+ аид киин|\d+ мүн киин|\d+ мүн мурн|\d+ сек киин|\d+ сек мурн|\d+ аи мурн|\d+ с киин|\d+ с мурн)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(келерки аптада|түштөн кииинки|ушул мүнөттө|өткөн аптада|келерки апт|ушул аптада|ушул саатта|эмдиги аида|эмдиги жылы|өткөн аида|өткөн апт|беишемби|бул аида|дүишөмбү|жекшемби|сентябрь|ушул апт|шаршемби|шеишемби|декабрь|октябрь|февраль|август|апрель|былтыр|ишемби|ноябрь|секунд|январь|быиыл|бүгүн|кечээ|мүнөт|таңкы|эртеӊ|азыр|апта|беиш|жума|июль|июнь|март|саат|шарш|шеиш|gmt|utc|авг|апр|апт|дек|дүи|жек|жыл|ишм|июл|июн|күн|маи|мар|мүн|ноя|окт|сек|сен|фев|янв|\+|\.|\[|\]|\||am|pm|аи|ст|тк|тң| |'|,|-|/|:|;|@|z|ж|м)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
