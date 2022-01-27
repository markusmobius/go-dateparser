// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kn_Locale = merge(nil, LocaleData{
	Name:      "kn",
	DateOrder: "DMY",
	Charset:   []rune("+,-./;@[]cgtuz|ಂಅಆಇಈಏಕಗಚಜಟಡತದಧನಪಫಬಭಮರಲಳವಶಷಸಹಾುೂೕೖ"),
	Translations: map[string]string{
		"ಪೂರವಾಹನ": "am",
		"ಶುಕರವಾರ": "friday",
		"ಸೂೕಮವಾರ": "monday",
		"ಅಕಟೂೕಬರ": "october",
		"ಭಾನುವಾರ": "sunday",
		"ಗುರುವಾರ": "thursday",
		"ಮಂಗಳವಾರ": "tuesday",
		"ಅಪರಾಹನ":  "pm",
		"ಸಪಟಂಬರ":  "september",
		"ಬುಧವಾರ":  "wednesday",
		"ಡಸಂಬರ":   "december",
		"ಫಬರವರ":   "february",
		"ತಂಗಳು":   "month",
		"ನವಂಬರ":   "november",
		"ಅಕಟೂೕ":   "october",
		"ಶನವಾರ":   "saturday",
		"ಏಪರಲ":    "april",
		"ಆಗಸಟ":    "august",
		"ಶುಕರ":    "friday",
		"ಜನವರ":    "january",
		"ಜುಲೖ":    "july",
		"ಮಾರಚ":    "march",
		"ಸೂೕಮ":    "monday",
		"ಸಕಂಡ":    "second",
		"ಸಪಟಂ":    "september",
		"ಭಾನು":    "sunday",
		"ಗುರು":    "thursday",
		"ಮಂಗಳ":    "tuesday",
		"ಏಪರ":     "april",
		"ಡಸಂ":     "december",
		"ಫಬರ":     "february",
		"gmt":     "gmt",
		"ಗಂಟ":     "hour",
		"ಜೂನ":     "june",
		"ನಮಷ":     "minute",
		"ನವಂ":     "november",
		"utc":     "utc",
		"ಬುಧ":     "wednesday",
		"ವಾರ":     "week",
		"ವರಷ":     "year",
		"am":      "am",
		"ಆಗ":      "august",
		"ದನ":      "day",
		"ಜನ":      "january",
		"ಮೕ":      "may",
		"pm":      "pm",
		"ಶನ":      "saturday",
		"'":       "",
		",":       "",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"|":       "",
		" ":       " ",
		"+":       "+",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		"z":       "z",
	},
	RelativeType: map[string]string{
		"ಮುಂದನ ತಂಗಳು": "in 1 month",
		"ಕಳದ ತಂಗಳು":   "1 month ago",
		"ಮುಂದನ ವಾರ":   "in 1 week",
		"ಮುಂದನ ವರಷ":   "in 1 year",
		"ಹಂದನ ವರಷ":    "1 year ago",
		"ಈ ತಂಗಳು":     "0 month ago",
		"ಕಳದ ವಾರ":     "1 week ago",
		"ಕಳದ ವರಷ":     "1 year ago",
		"ಈ ಗಂಟ":       "0 hour ago",
		"ಈ ನಮಷ":       "0 minute ago",
		"ಈ ವಾರ":       "0 week ago",
		"ಈ ವರಷ":       "0 year ago",
		"ಇಂದು":        "0 day ago",
		"ನನನ":         "1 day ago",
		"ನಾಳ":         "in 1 day",
		"ಈಗ":          "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) ತಂಗಳುಗಳ ಹಂದ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ಸಕಂಡುಗಳ ಹಂದ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) ಗಂಟಗಳ ಹಂದ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ನಮಷಗಳ ಹಂದ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ತಂಗಳು ಹಂದ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ವಾರಗಳ ಹಂದ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ವರಷಗಳ ಹಂದ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ತಂಗಳುಗಳಲಲ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ಸಕಂಡ‌ಗಳಲಲ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ದನಗಳ ಹಂದ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ನಮಷದ ಹಂದ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ತಂಗಳ ಹಂದ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ಸಕಂಡ ಹಂದ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) ವಾರದ ಹಂದ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ವರಷದ ಹಂದ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) ಸಕಂಡ‌ನಲಲ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ದನದ ಹಂದ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ಗಂಟ ಹಂದ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ಗಂಟಗಳಲಲ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ನಮಷಗಳಲಲ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ವಾರಗಳಲಲ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ವರಷಗಳಲಲ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ದನಗಳಲಲ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ಗಂಟಯಲಲ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ನಮಷದಲಲ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ತಂಗಳಲಲ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ವಾರದಲಲ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ವರಷದಲಲ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ದನದಲಲ`), "in $1 day"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ತಂಗಳುಗಳ ಹಂದ|\d+ ಸಕಂಡುಗಳ ಹಂದ|\d+ ಗಂಟಗಳ ಹಂದ|\d+ ತಂಗಳು ಹಂದ|\d+ ತಂಗಳುಗಳಲಲ|\d+ ನಮಷಗಳ ಹಂದ|\d+ ವರಷಗಳ ಹಂದ|\d+ ವಾರಗಳ ಹಂದ|\d+ ಸಕಂಡ‌ಗಳಲಲ|\d+ ತಂಗಳ ಹಂದ|\d+ ದನಗಳ ಹಂದ|\d+ ನಮಷದ ಹಂದ|\d+ ವರಷದ ಹಂದ|\d+ ವಾರದ ಹಂದ|\d+ ಸಕಂಡ ಹಂದ|\d+ ಸಕಂಡ‌ನಲಲ|\d+ ಗಂಟ ಹಂದ|\d+ ಗಂಟಗಳಲಲ|\d+ ದನದ ಹಂದ|\d+ ನಮಷಗಳಲಲ|\d+ ವರಷಗಳಲಲ|\d+ ವಾರಗಳಲಲ|\d+ ಗಂಟಯಲಲ|\d+ ತಂಗಳಲಲ|\d+ ದನಗಳಲಲ|\d+ ನಮಷದಲಲ|\d+ ವರಷದಲಲ|\d+ ವಾರದಲಲ|\d+ ದನದಲಲ)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ತಂಗಳುಗಳ ಹಂದ|\d+ ಸಕಂಡುಗಳ ಹಂದ|\d+ ಗಂಟಗಳ ಹಂದ|\d+ ತಂಗಳು ಹಂದ|\d+ ತಂಗಳುಗಳಲಲ|\d+ ನಮಷಗಳ ಹಂದ|\d+ ವರಷಗಳ ಹಂದ|\d+ ವಾರಗಳ ಹಂದ|\d+ ಸಕಂಡ‌ಗಳಲಲ|\d+ ತಂಗಳ ಹಂದ|\d+ ದನಗಳ ಹಂದ|\d+ ನಮಷದ ಹಂದ|\d+ ವರಷದ ಹಂದ|\d+ ವಾರದ ಹಂದ|\d+ ಸಕಂಡ ಹಂದ|\d+ ಸಕಂಡ‌ನಲಲ|\d+ ಗಂಟ ಹಂದ|\d+ ಗಂಟಗಳಲಲ|\d+ ದನದ ಹಂದ|\d+ ನಮಷಗಳಲಲ|\d+ ವರಷಗಳಲಲ|\d+ ವಾರಗಳಲಲ|\d+ ಗಂಟಯಲಲ|\d+ ತಂಗಳಲಲ|\d+ ದನಗಳಲಲ|\d+ ನಮಷದಲಲ|\d+ ವರಷದಲಲ|\d+ ವಾರದಲಲ|\d+ ದನದಲಲ)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ಮುಂದನ ತಂಗಳು|ಕಳದ ತಂಗಳು|ಮುಂದನ ವರಷ|ಮುಂದನ ವಾರ|ಹಂದನ ವರಷ|ಅಕಟೂೕಬರ|ಈ ತಂಗಳು|ಕಳದ ವರಷ|ಕಳದ ವಾರ|ಗುರುವಾರ|ಪೂರವಾಹನ|ಭಾನುವಾರ|ಮಂಗಳವಾರ|ಶುಕರವಾರ|ಸೂೕಮವಾರ|ಅಪರಾಹನ|ಬುಧವಾರ|ಸಪಟಂಬರ|ಅಕಟೂೕ|ಈ ಗಂಟ|ಈ ನಮಷ|ಈ ವರಷ|ಈ ವಾರ|ಡಸಂಬರ|ತಂಗಳು|ನವಂಬರ|ಫಬರವರ|ಶನವಾರ|ಆಗಸಟ|ಇಂದು|ಏಪರಲ|ಗುರು|ಜನವರ|ಜುಲೖ|ಭಾನು|ಮಂಗಳ|ಮಾರಚ|ಶುಕರ|ಸಕಂಡ|ಸಪಟಂ|ಸೂೕಮ|gmt|utc|ಏಪರ|ಗಂಟ|ಜೂನ|ಡಸಂ|ನನನ|ನಮಷ|ನವಂ|ನಾಳ|ಫಬರ|ಬುಧ|ವರಷ|ವಾರ|\+|\.|\[|\]|\||am|pm|ಆಗ|ಈಗ|ಜನ|ದನ|ಮೕ|ಶನ| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
