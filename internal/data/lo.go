// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lo_Locale = merge(nil, LocaleData{
	Name:      "lo",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzກຄງຈຊຍດຕຖທນປພມລວສຫອະັາິີຶືຸົຼຽເແໂ່້ໍ`),
	Translations: map[string]string{
		"ວນອງຄານ": "tuesday",
		"ກອນທຽງ":  "am",
		"ພດສະພາ":  "may",
		"ວນພະຫດ":  "thursday",
		"ວນອາທດ":  "sunday",
		"ກລະກດ":   "july",
		"ຊວໂມງ":   "hour",
		"ມງກອນ":   "january",
		"ວນເສາ":   "saturday",
		"ຫງທຽງ":   "pm",
		"ອງຄານ":   "tuesday",
		"ກນຍາ":    "september",
		"ກມພາ":    "february",
		"ທນວາ":    "december",
		"ພະຈກ":    "november",
		"ພະຫດ":    "thursday",
		"ມຖນາ":    "june",
		"ວນຈນ":    "monday",
		"ວນພດ":    "wednesday",
		"ວນສກ":    "friday",
		"ວນາທ":    "second",
		"ສງຫາ":    "august",
		"ອາທດ":    "week",
		"ເດອນ":    "month",
		"ເມສາ":    "april",
		"gmt":     "gmt",
		"utc":     "utc",
		"ຕລາ":     "october",
		"ນາທ":     "minute",
		"ມນາ":     "march",
		"ເສາ":     "saturday",
		"am":      "am",
		"pm":      "pm",
		"ກຍ":      "september",
		"ກພ":      "february",
		"ກລ":      "july",
		"ຈນ":      "monday",
		"ຊມ":      "hour",
		"ຕລ":      "october",
		"ທວ":      "december",
		"ນທ":      "minute",
		"ພຈ":      "november",
		"ພດ":      "wednesday",
		"ພພ":      "may",
		"ມກ":      "january",
		"ມຖ":      "june",
		"ມນ":      "march",
		"ມສ":      "april",
		"ສກ":      "friday",
		"ສຫ":      "august",
		" ":       " ",
		"'":       "",
		"+":       "+",
		",":       "",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"z":       "z",
		"|":       "",
		"ດ":       "month",
		"ປ":       "year",
		"ມ":       "day",
		"ວ":       "second",
		"ອ":       "week",
	},
	RelativeType: map[string]string{
		"ອາທດຫນາ": "in 1 week",
		"ອາທດແລວ": "1 week ago",
		"ເດອນຫນາ": "in 1 month",
		"ເດອນແລວ": "1 month ago",
		"ຊວໂມງນ":  "0 hour ago",
		"ອາທດນ":   "0 week ago",
		"ເດອນນ":   "0 month ago",
		"ຕອນນ":    "0 second ago",
		"ນາທນ":    "0 minute ago",
		"ປກາຍ":    "1 year ago",
		"ປຫນາ":    "in 1 year",
		"ມວານ":    "1 day ago",
		"ມອນ":     "in 1 day",
		"ປນ":      "0 year ago",
		"ມນ":      "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) ໃນອກ 0 ນາທ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ຊວໂມງ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ວນາທ`), "in $1 second"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ອາທດ`), "in $1 week"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ເດອນ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ຊວໂມງກອນ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ວນາທກອນ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) ອາທດກອນ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ເດອນກອນ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ຊມ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ອທ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ຊມ ກອນ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ນທ ກອນ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ນາທກອນ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ອທ ກອນ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ດ`), "in $1 month"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ປ`), "in $1 year"},
		{regexp.MustCompile(`(?i)ໃນອກ (\d+) ມ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ດ ກອນ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ວ ກອນ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ໃນ (\d+) ນທ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ປກອນ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ມກອນ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ໃນ (\d+) ວ`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ໃນອກ 0 ນາທ|ໃນອກ \d+ ຊວໂມງ|ໃນອກ \d+ ວນາທ|ໃນອກ \d+ ອາທດ|ໃນອກ \d+ ເດອນ|\d+ ຊວໂມງກອນ|\d+ ວນາທກອນ|\d+ ອາທດກອນ|\d+ ເດອນກອນ|ໃນອກ \d+ ຊມ|ໃນອກ \d+ ອທ|\d+ ຊມ ກອນ|\d+ ນທ ກອນ|\d+ ນາທກອນ|\d+ ອທ ກອນ|ໃນອກ \d+ ດ|ໃນອກ \d+ ປ|ໃນອກ \d+ ມ|\d+ ດ ກອນ|\d+ ວ ກອນ|ໃນ \d+ ນທ|\d+ ປກອນ|\d+ ມກອນ|ໃນ \d+ ວ)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ໃນອກ 0 ນາທ|ໃນອກ \d+ ຊວໂມງ|ໃນອກ \d+ ວນາທ|ໃນອກ \d+ ອາທດ|ໃນອກ \d+ ເດອນ|\d+ ຊວໂມງກອນ|\d+ ວນາທກອນ|\d+ ອາທດກອນ|\d+ ເດອນກອນ|ໃນອກ \d+ ຊມ|ໃນອກ \d+ ອທ|\d+ ຊມ ກອນ|\d+ ນທ ກອນ|\d+ ນາທກອນ|\d+ ອທ ກອນ|ໃນອກ \d+ ດ|ໃນອກ \d+ ປ|ໃນອກ \d+ ມ|\d+ ດ ກອນ|\d+ ວ ກອນ|ໃນ \d+ ນທ|\d+ ປກອນ|\d+ ມກອນ|ໃນ \d+ ວ)$`),
	KnownWords:      []string{"ວນອງຄານ", "ອາທດຫນາ", "ອາທດແລວ", "ເດອນຫນາ", "ເດອນແລວ", "ກອນທຽງ", "ຊວໂມງນ", "ພດສະພາ", "ວນພະຫດ", "ວນອາທດ", "ກລະກດ", "ຊວໂມງ", "ມງກອນ", "ວນເສາ", "ຫງທຽງ", "ອງຄານ", "ອາທດນ", "ເດອນນ", "ກນຍາ", "ກມພາ", "ຕອນນ", "ທນວາ", "ນາທນ", "ປກາຍ", "ປຫນາ", "ພະຈກ", "ພະຫດ", "ມຖນາ", "ມວານ", "ວນຈນ", "ວນພດ", "ວນສກ", "ວນາທ", "ສງຫາ", "ອາທດ", "ເດອນ", "ເມສາ", "gmt", "utc", "ຕລາ", "ນາທ", "ມນາ", "ມອນ", "ເສາ", "am", "pm", "ກຍ", "ກພ", "ກລ", "ຈນ", "ຊມ", "ຕລ", "ທວ", "ນທ", "ປນ", "ພຈ", "ພດ", "ພພ", "ມກ", "ມຖ", "ມນ", "ມນ", "ມສ", "ສກ", "ສຫ", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ດ", "ປ", "ມ", "ວ", "ອ"},
})
