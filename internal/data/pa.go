// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pa_Locale = merge(nil, LocaleData{
	Name:      "pa",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzਅਆਇਈਐਕਗਘਚਛਜਟਣਤਦਧਨਪਫਬਭਮਰਲਵਸਹ਼ਾਿੀੁੂੇੈੋ੍ੰੱ`),
	Translations: map[string][]string{
		"ਸਨਿਚਰਵਾਰ": {"saturday"},
		"ਮਗਲਵਾਰ":   {"tuesday"},
		"ਵੀਰਵਾਰ":   {"thursday"},
		"ਸਕਰਵਾਰ":   {"friday"},
		"ਅਕਤਬਰ":    {"october"},
		"ਐਤਵਾਰ":    {"sunday"},
		"ਜਨਵਰੀ":    {"january"},
		"ਫਰਵਰੀ":    {"february"},
		"ਬਧਵਾਰ":    {"wednesday"},
		"ਮਹੀਨਾ":    {"month"},
		"ਸਨਿਚਰ":    {"saturday"},
		"ਸਮਵਾਰ":    {"monday"},
		"ਅਗਸਤ":     {"august"},
		"ਅਪਰਲ":     {"april"},
		"ਜਲਾਈ":     {"july"},
		"ਦਸਬਰ":     {"december"},
		"ਨਵਬਰ":     {"november"},
		"ਮਾਰਚ":     {"march"},
		"ਸਕਿਟ":     {"second"},
		"ਸਤਬਰ":     {"september"},
		"ਹਫਤਾ":     {"week"},
		"gmt":      {"gmt"},
		"utc":      {"utc"},
		"ਅਕਤ":      {"october"},
		"ਅਪਰ":      {"april"},
		"ਘਟਾ":      {"hour"},
		"ਜਲਾ":      {"july"},
		"ਦਿਨ":      {"day"},
		"ਬਾਦ":      {"pm"},
		"ਮਗਲ":      {"tuesday"},
		"ਮਿਟ":      {"minute"},
		"ਵੀਰ":      {"thursday"},
		"ਸਕਰ":      {"friday"},
		"ਸਾਲ":      {"year"},
		"am":       {"am"},
		"pm":       {"pm"},
		"ਅਗ":       {"august"},
		"ਐਤ":       {"sunday"},
		"ਜਨ":       {"june", "january"},
		"ਦਸ":       {"december"},
		"ਨਵ":       {"november"},
		"ਪਦ":       {"am"},
		"ਫਰ":       {"february"},
		"ਬਧ":       {"wednesday"},
		"ਮਈ":       {"may"},
		"ਸਤ":       {"september"},
		"ਸਮ":       {"monday"},
		" ":        {" "},
		"'":        {""},
		"+":        {"+"},
		",":        {""},
		"-":        {"-"},
		".":        {"."},
		"/":        {"/"},
		":":        {":"},
		";":        {""},
		"@":        {""},
		"[":        {""},
		"]":        {""},
		"z":        {"z"},
		"|":        {""},
		"ਘ":        {"hour"},
	},
	RelativeType: map[string]string{
		"ਪਿਛਲਾ ਮਹੀਨਾ": "1 month ago",
		"ਅਗਲਾ ਮਹੀਨਾ":  "in 1 month",
		"ਪਿਛਲਾ ਹਫਤਾ":  "1 week ago",
		"ਅਗਲਾ ਹਫਤਾ":   "in 1 week",
		"ਪਿਛਲਾ ਸਾਲ":   "1 year ago",
		"ਬੀਤਿਆ ਕਲਹ":   "1 day ago",
		"ਅਗਲਾ ਸਾਲ":    "in 1 year",
		"ਇਹ ਮਹੀਨਾ":    "0 month ago",
		"ਇਹ ਹਫਤਾ":     "0 week ago",
		"ਇਸ ਮਿਟ":      "0 minute ago",
		"ਇਹ ਸਾਲ":      "0 year ago",
		"ਇਸ ਘਟ":       "0 hour ago",
		"ਭਲਕ":         "in 1 day",
		"ਅਜ":          "0 day ago",
		"ਹਣ":          "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਹੀਨਾ ਪਹਿਲਾ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਹੀਨ ਪਹਿਲਾ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਹੀਨਿਆ ਵਿਚ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਕਿਟ ਪਹਿਲਾ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਹਫਤਾ ਪਹਿਲਾ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਘਟਾ ਪਹਿਲਾ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਦਿਨ ਪਹਿਲਾ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਿਟ ਪਹਿਲਾ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਕਿਟਾ ਵਿਚ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਾਲ ਪਹਿਲਾ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਹਫਤ ਪਹਿਲਾ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਹਫਤਿਆ ਵਿਚ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਘਟ ਪਹਿਲਾ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਘਟਿਆ ਵਿਚ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਦਿਨਾ ਵਿਚ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਹੀਨ ਵਿਚ`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਿਟਾ ਵਿਚ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਕਿਟ ਵਿਚ`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਾਲਾ ਵਿਚ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਦਿਨ ਵਿਚ`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਮਿਟ ਵਿਚ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਸਾਲ ਵਿਚ`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਹਫਤ ਵਿਚ`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ਘਟ ਵਿਚ`), "in $1 hour"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* ਮਹੀਨਾ ਪਹਿਲਾ|\d+[.,]?\d* ਮਹੀਨ ਪਹਿਲਾ|\d+[.,]?\d* ਮਹੀਨਿਆ ਵਿਚ|\d+[.,]?\d* ਸਕਿਟ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤਾ ਪਹਿਲਾ|\d+[.,]?\d* ਘਟਾ ਪਹਿਲਾ|\d+[.,]?\d* ਦਿਨ ਪਹਿਲਾ|\d+[.,]?\d* ਮਿਟ ਪਹਿਲਾ|\d+[.,]?\d* ਸਕਿਟਾ ਵਿਚ|\d+[.,]?\d* ਸਾਲ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤਿਆ ਵਿਚ|\d+[.,]?\d* ਘਟ ਪਹਿਲਾ|\d+[.,]?\d* ਘਟਿਆ ਵਿਚ|\d+[.,]?\d* ਦਿਨਾ ਵਿਚ|\d+[.,]?\d* ਮਹੀਨ ਵਿਚ|\d+[.,]?\d* ਮਿਟਾ ਵਿਚ|\d+[.,]?\d* ਸਕਿਟ ਵਿਚ|\d+[.,]?\d* ਸਾਲਾ ਵਿਚ|\d+[.,]?\d* ਦਿਨ ਵਿਚ|\d+[.,]?\d* ਮਿਟ ਵਿਚ|\d+[.,]?\d* ਸਾਲ ਵਿਚ|\d+[.,]?\d* ਹਫਤ ਵਿਚ|\d+[.,]?\d* ਘਟ ਵਿਚ)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* ਮਹੀਨਾ ਪਹਿਲਾ|\d+[.,]?\d* ਮਹੀਨ ਪਹਿਲਾ|\d+[.,]?\d* ਮਹੀਨਿਆ ਵਿਚ|\d+[.,]?\d* ਸਕਿਟ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤਾ ਪਹਿਲਾ|\d+[.,]?\d* ਘਟਾ ਪਹਿਲਾ|\d+[.,]?\d* ਦਿਨ ਪਹਿਲਾ|\d+[.,]?\d* ਮਿਟ ਪਹਿਲਾ|\d+[.,]?\d* ਸਕਿਟਾ ਵਿਚ|\d+[.,]?\d* ਸਾਲ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤ ਪਹਿਲਾ|\d+[.,]?\d* ਹਫਤਿਆ ਵਿਚ|\d+[.,]?\d* ਘਟ ਪਹਿਲਾ|\d+[.,]?\d* ਘਟਿਆ ਵਿਚ|\d+[.,]?\d* ਦਿਨਾ ਵਿਚ|\d+[.,]?\d* ਮਹੀਨ ਵਿਚ|\d+[.,]?\d* ਮਿਟਾ ਵਿਚ|\d+[.,]?\d* ਸਕਿਟ ਵਿਚ|\d+[.,]?\d* ਸਾਲਾ ਵਿਚ|\d+[.,]?\d* ਦਿਨ ਵਿਚ|\d+[.,]?\d* ਮਿਟ ਵਿਚ|\d+[.,]?\d* ਸਾਲ ਵਿਚ|\d+[.,]?\d* ਹਫਤ ਵਿਚ|\d+[.,]?\d* ਘਟ ਵਿਚ)$`),
	KnownWords:      []string{"ਪਿਛਲਾ ਮਹੀਨਾ", "ਅਗਲਾ ਮਹੀਨਾ", "ਪਿਛਲਾ ਹਫਤਾ", "ਅਗਲਾ ਹਫਤਾ", "ਪਿਛਲਾ ਸਾਲ", "ਬੀਤਿਆ ਕਲਹ", "ਅਗਲਾ ਸਾਲ", "ਇਹ ਮਹੀਨਾ", "ਸਨਿਚਰਵਾਰ", "ਇਹ ਹਫਤਾ", "ਇਸ ਮਿਟ", "ਇਹ ਸਾਲ", "ਮਗਲਵਾਰ", "ਵੀਰਵਾਰ", "ਸਕਰਵਾਰ", "ਅਕਤਬਰ", "ਇਸ ਘਟ", "ਐਤਵਾਰ", "ਜਨਵਰੀ", "ਫਰਵਰੀ", "ਬਧਵਾਰ", "ਮਹੀਨਾ", "ਸਨਿਚਰ", "ਸਮਵਾਰ", "ਅਗਸਤ", "ਅਪਰਲ", "ਜਲਾਈ", "ਦਸਬਰ", "ਨਵਬਰ", "ਮਾਰਚ", "ਸਕਿਟ", "ਸਤਬਰ", "ਹਫਤਾ", "gmt", "utc", "ਅਕਤ", "ਅਪਰ", "ਘਟਾ", "ਜਲਾ", "ਦਿਨ", "ਬਾਦ", "ਭਲਕ", "ਮਗਲ", "ਮਿਟ", "ਵੀਰ", "ਸਕਰ", "ਸਾਲ", "am", "pm", "ਅਗ", "ਅਜ", "ਐਤ", "ਜਨ", "ਦਸ", "ਨਵ", "ਪਦ", "ਫਰ", "ਬਧ", "ਮਈ", "ਸਤ", "ਸਮ", "ਹਣ", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ਘ"},
})
