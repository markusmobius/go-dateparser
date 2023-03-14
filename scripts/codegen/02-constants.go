package main

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/unicode/rangetable"
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
	nfkcTransformer    = transform.Chain(norm.NFKD, norm.NFKC)
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
	rxCharsetFilter       = regexp.MustCompile(`^[\P{L}\p{N}_]+$`)
	rxNbsp                = regexp.MustCompile(`\x{a0}`)

	importantTokens = []string{"+", ":", ".", " ", "-", "/", "am", "pm", "utc", "gmt", "z"}
	commonChars     = rangetable.New([]rune("0123456789:(){}'.qamp ")...)

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
