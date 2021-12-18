package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	CLDR_VERSION            = "31.0.1"
	RAW_DIR                 = "./raw-data"
	SUPPLEMENTARY_DIR       = "./data-supplementary/date-translation"
	SUPPLEMENTARY_BASE_PATH = "./data-supplementary/base_data.yaml"
	JSON_DIR                = "./data-json"
	GO_CODE_DIR             = "./internal/data"
)

var (
	unicodeTransformer = transform.Chain(norm.NFKD, runes.Remove(runes.In(unicode.Mn)), norm.NFKC)

	apostropheLookAlikeChars = []string{
		`\x{2019}`, // right single quotation mark
		`\x{02bc}`, // modifier letter apostrophe
		`\x{02bb}`, // modifier letter turned comma
		`\x{055a}`, // armenian apostrophe
		`\x{a78c}`, // latin small letter saltillo
		`\x{2032}`, // prime
		`\x{2035}`, // reversed prime
		`\x{02b9}`, // modifier letter prime
		`\x{ff07}`, // fullwidth apostrophe
	}

	numberChars = []string{
		`0-9`,
		`\x{0660}-\x{0669}`, // arabic
		`\x{0f20}-\x{0f29}`, // tibetan
	}

	rxLocale              = regexp.MustCompile(`-[A-Z0-9]+$`)
	rxLocaleCleaner       = regexp.MustCompile(`-\w+`)
	rxDateOrderPattern    = regexp.MustCompile(`([DMY])+\x{200f}*[-/. \t]*([DMY])+\x{200f}*[-/. \t]*([DMY])+`)
	rxNotRelativePattern  = regexp.MustCompile(`[\+\-]\s*\{0\}`)
	rxDefaultMonthPattern = regexp.MustCompile(`^m?[` + strings.Join(numberChars, "") + `]+$`)
	rxSanitizeAposthrope  = regexp.MustCompile(strings.Join(apostropheLookAlikeChars, "|"))
	rxSanitizeDateField   = regexp.MustCompile(`\$?\d+|^in|ago$`)
	rxAmPmPattern         = regexp.MustCompile(`^\s*[AaPp]\s*\.?\s*[Mm]\s*\.?\s*$`)
	rxParenthesisPattern  = regexp.MustCompile(`[\(\)]`)
	rxGoEmptyField        = regexp.MustCompile(`(?m)^.*\{\s*\},?$\n*`)
	rxGoZeroField         = regexp.MustCompile(`(?m)^.*(false|0|nil),?$\n*`)
	rxGoRegexImport       = regexp.MustCompile(`(?m)^\s*import "regexp"\s*$`)
	rxPythonCaptureGroup  = regexp.MustCompile(`\\(\d+)`)
	rxGoCaptureGroup      = regexp.MustCompile(`\$\{?(\d+)\}?`)
	rxParentheses         = regexp.MustCompile(`[\(\)]`)

	importantTokens = []string{"+", ":", ".", " ", "-", "/", "am", "pm", "utc", "gmt", "z"}

	// Languages with insufficient translation data are excluded
	excludedLanguages = map[string]struct{}{
		"cu":       {},
		"kkj":      {},
		"nds":      {},
		"prg":      {},
		"tk":       {},
		"vai":      {},
		"vai-Latn": {},
		"vai-Vaii": {},
		"vo":       {},
	}

	// List of most common locales from https://w3techs.com/technologies/overview/content_language
	// Last updated on 30.10.2020
	mostCommonLocales = map[string]int{
		"en": 1,
		"ru": 2,
		"es": 3,
		"tr": 4,
		"fa": 5,
		"fr": 6,
		"de": 7,
		"ja": 8,
		"pt": 9,
		"vi": 10,
		"zh": 11,
		"ar": 12,
		"it": 13,
		"pl": 14,
		"id": 15,
		"el": 16,
		"nl": 17,
		"ko": 18,
		"th": 19,
		"he": 20,
		"uk": 21,
		"cs": 22,
		"sv": 23,
		"ro": 24,
		"hu": 25,
		"da": 26,
		"sr": 27,
		"sk": 28,
		"fi": 29,
		"bg": 30,
		"hr": 31,
		"lt": 32,
		"hi": 33,
		"nb": 34,
		"sl": 35,
		"nn": 36,
		"et": 37,
		"lv": 38,
	}

	enMonthNames = map[int]string{
		1:  "january",
		2:  "february",
		3:  "march",
		4:  "april",
		5:  "may",
		6:  "june",
		7:  "july",
		8:  "august",
		9:  "september",
		10: "october",
		11: "november",
		12: "december",
	}

	enDayNames = map[int]string{
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
		7: "sunday",
	}

	enDayPeriods = []string{"am", "pm"}

	enDateFields = []string{
		"year",
		"month",
		"week",
		"day",
		"hour",
		"minute",
		"second",
		"1 year ago",
		"0 year ago",
		"in 1 year",
		"1 month ago",
		"0 month ago",
		"in 1 month",
		"1 week ago",
		"0 week ago",
		"in 1 week",
		"1 day ago",
		"0 day ago",
		"in 1 day",
		"0 hour ago",
		"0 minute ago",
		"0 second ago",
		"in $1 year",
		"$1 year ago",
		"in $1 month",
		"$1 month ago",
		"in $1 week",
		"$1 week ago",
		"in $1 day",
		"$1 day ago",
		"in $1 hour",
		"$1 hour ago",
		"in $1 minute",
		"$1 minute ago",
		"in $1 second",
		"$1 second ago",
	}
)
