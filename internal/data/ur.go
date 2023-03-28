// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	ur_Locale    LocaleData
	ur_IN_Locale LocaleData
)

func init() {
	ur_Locale = merge(nil, LocaleData{
		Name:      "ur",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzآئابتجدرزسشعفلمنويٹپچڈکگھہیے`),
		Translations: map[string][]string{
			"اکتوبر": {"october"},
			"جمعرات": {"thursday"},
			"جولايی": {"july"},
			"سوموار": {"monday"},
			"اتوار":  {"sunday"},
			"اپریل":  {"april"},
			"جنوری":  {"january"},
			"دسمبر":  {"december"},
			"ستمبر":  {"september"},
			"سیکنڈ":  {"second"},
			"فروری":  {"february"},
			"مہینہ":  {"month"},
			"نومبر":  {"november"},
			"گھنٹہ":  {"hour"},
			"اگست":   {"august"},
			"جمعہ":   {"friday"},
			"مارچ":   {"march"},
			"منگل":   {"tuesday"},
			"ہفتہ":   {"week", "saturday"},
			"gmt":    {"gmt"},
			"utc":    {"utc"},
			"بدھ":    {"wednesday"},
			"جون":    {"june"},
			"سال":    {"year"},
			"ماہ":    {"month"},
			"منٹ":    {"minute"},
			"ميی":    {"may"},
			"am":     {"am"},
			"pm":     {"pm"},
			"دن":     {"day"},
			" ":      {" "},
			"'":      {""},
			"+":      {"+"},
			",":      {""},
			"-":      {"-"},
			".":      {"."},
			"/":      {"/"},
			":":      {":"},
			";":      {""},
			"@":      {""},
			"[":      {""},
			"]":      {""},
			"z":      {"z"},
			"|":      {""},
		},
		RelativeType: map[string]string{
			"پچھلے مہینہ": "1 month ago",
			"اگلے مہینہ":  "in 1 month",
			"پچھلے ہفتہ":  "1 week ago",
			"اگلے ہفتہ":   "in 1 week",
			"گزشتہ سال":   "1 year ago",
			"اس مہینہ":    "0 month ago",
			"اس گھنٹے":    "0 hour ago",
			"ايندہ کل":    "in 1 day",
			"اگلے سال":    "in 1 year",
			"گزشتہ کل":    "1 day ago",
			"اس ہفتہ":     "0 week ago",
			"اس سال":      "0 year ago",
			"اس منٹ":      "0 minute ago",
			"اب":          "0 second ago",
			"اج":          "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سیکنڈ پہلے`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) مہینہ پہلے`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) مہینے پہلے`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹوں میں`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹہ پہلے`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹے پہلے`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دنوں پہلے`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سیکنڈ میں`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) مہینہ میں`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) مہینے میں`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹہ میں`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹے میں`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتہ پہلے`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتے پہلے`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دنوں میں`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سال پہلے`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ماہ پہلے`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) منٹ پہلے`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتہ میں`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتے میں`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دن پہلے`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سال میں`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ماہ قبل`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ماہ میں`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) منٹ میں`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دن میں`), "in $1 day"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* سیکنڈ پہلے|\d+[.,]?\d* مہینہ پہلے|\d+[.,]?\d* مہینے پہلے|\d+[.,]?\d* گھنٹوں میں|\d+[.,]?\d* گھنٹہ پہلے|\d+[.,]?\d* گھنٹے پہلے|\d+[.,]?\d* دنوں پہلے|\d+[.,]?\d* سیکنڈ میں|\d+[.,]?\d* مہینہ میں|\d+[.,]?\d* مہینے میں|\d+[.,]?\d* گھنٹہ میں|\d+[.,]?\d* گھنٹے میں|\d+[.,]?\d* ہفتہ پہلے|\d+[.,]?\d* ہفتے پہلے|\d+[.,]?\d* دنوں میں|\d+[.,]?\d* سال پہلے|\d+[.,]?\d* ماہ پہلے|\d+[.,]?\d* منٹ پہلے|\d+[.,]?\d* ہفتہ میں|\d+[.,]?\d* ہفتے میں|\d+[.,]?\d* دن پہلے|\d+[.,]?\d* سال میں|\d+[.,]?\d* ماہ قبل|\d+[.,]?\d* ماہ میں|\d+[.,]?\d* منٹ میں|\d+[.,]?\d* دن میں)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* سیکنڈ پہلے|\d+[.,]?\d* مہینہ پہلے|\d+[.,]?\d* مہینے پہلے|\d+[.,]?\d* گھنٹوں میں|\d+[.,]?\d* گھنٹہ پہلے|\d+[.,]?\d* گھنٹے پہلے|\d+[.,]?\d* دنوں پہلے|\d+[.,]?\d* سیکنڈ میں|\d+[.,]?\d* مہینہ میں|\d+[.,]?\d* مہینے میں|\d+[.,]?\d* گھنٹہ میں|\d+[.,]?\d* گھنٹے میں|\d+[.,]?\d* ہفتہ پہلے|\d+[.,]?\d* ہفتے پہلے|\d+[.,]?\d* دنوں میں|\d+[.,]?\d* سال پہلے|\d+[.,]?\d* ماہ پہلے|\d+[.,]?\d* منٹ پہلے|\d+[.,]?\d* ہفتہ میں|\d+[.,]?\d* ہفتے میں|\d+[.,]?\d* دن پہلے|\d+[.,]?\d* سال میں|\d+[.,]?\d* ماہ قبل|\d+[.,]?\d* ماہ میں|\d+[.,]?\d* منٹ میں|\d+[.,]?\d* دن میں)$`),
		KnownWords:      []string{"پچھلے مہینہ", "اگلے مہینہ", "پچھلے ہفتہ", "اگلے ہفتہ", "گزشتہ سال", "اس مہینہ", "اس گھنٹے", "ايندہ کل", "اگلے سال", "گزشتہ کل", "اس ہفتہ", "اس سال", "اس منٹ", "اکتوبر", "جمعرات", "جولايی", "سوموار", "اتوار", "اپریل", "جنوری", "دسمبر", "ستمبر", "سیکنڈ", "فروری", "مہینہ", "نومبر", "گھنٹہ", "اگست", "جمعہ", "مارچ", "منگل", "ہفتہ", "gmt", "utc", "بدھ", "جون", "سال", "ماہ", "منٹ", "ميی", "am", "pm", "اب", "اج", "دن", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})

	ur_IN_Locale = merge(&ur_Locale, LocaleData{
		Name:      "ur-IN",
		DateOrder: "DMY",
		Translations: map[string][]string{
			"پیر": {"monday"},
		},
		RelativeType: map[string]string{
			"گزشتہ ہفتہ": "1 week ago",
			"گزشتہ ماہ":  "1 month ago",
			"اگلے ماہ":   "in 1 month",
			"اس ماہ":     "0 month ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سالوں پہلے`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سالوں میں`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) سیکنڈ قبل`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹہ قبل`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) گھنٹے قبل`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتوں میں`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتہ قبل`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ہفتے قبل`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) منٹ قبل`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) دن قبل`), "$1 day ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* سالوں پہلے|\d+[.,]?\d* سیکنڈ پہلے|\d+[.,]?\d* مہینہ پہلے|\d+[.,]?\d* مہینے پہلے|\d+[.,]?\d* گھنٹوں میں|\d+[.,]?\d* گھنٹہ پہلے|\d+[.,]?\d* گھنٹے پہلے|\d+[.,]?\d* دنوں پہلے|\d+[.,]?\d* سالوں میں|\d+[.,]?\d* سیکنڈ قبل|\d+[.,]?\d* سیکنڈ میں|\d+[.,]?\d* مہینہ میں|\d+[.,]?\d* مہینے میں|\d+[.,]?\d* گھنٹہ قبل|\d+[.,]?\d* گھنٹہ میں|\d+[.,]?\d* گھنٹے قبل|\d+[.,]?\d* گھنٹے میں|\d+[.,]?\d* ہفتوں میں|\d+[.,]?\d* ہفتہ پہلے|\d+[.,]?\d* ہفتے پہلے|\d+[.,]?\d* دنوں میں|\d+[.,]?\d* سال پہلے|\d+[.,]?\d* ماہ پہلے|\d+[.,]?\d* منٹ پہلے|\d+[.,]?\d* ہفتہ قبل|\d+[.,]?\d* ہفتہ میں|\d+[.,]?\d* ہفتے قبل|\d+[.,]?\d* ہفتے میں|\d+[.,]?\d* دن پہلے|\d+[.,]?\d* سال میں|\d+[.,]?\d* ماہ قبل|\d+[.,]?\d* ماہ میں|\d+[.,]?\d* منٹ قبل|\d+[.,]?\d* منٹ میں|\d+[.,]?\d* دن قبل|\d+[.,]?\d* دن میں)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* سالوں پہلے|\d+[.,]?\d* سیکنڈ پہلے|\d+[.,]?\d* مہینہ پہلے|\d+[.,]?\d* مہینے پہلے|\d+[.,]?\d* گھنٹوں میں|\d+[.,]?\d* گھنٹہ پہلے|\d+[.,]?\d* گھنٹے پہلے|\d+[.,]?\d* دنوں پہلے|\d+[.,]?\d* سالوں میں|\d+[.,]?\d* سیکنڈ قبل|\d+[.,]?\d* سیکنڈ میں|\d+[.,]?\d* مہینہ میں|\d+[.,]?\d* مہینے میں|\d+[.,]?\d* گھنٹہ قبل|\d+[.,]?\d* گھنٹہ میں|\d+[.,]?\d* گھنٹے قبل|\d+[.,]?\d* گھنٹے میں|\d+[.,]?\d* ہفتوں میں|\d+[.,]?\d* ہفتہ پہلے|\d+[.,]?\d* ہفتے پہلے|\d+[.,]?\d* دنوں میں|\d+[.,]?\d* سال پہلے|\d+[.,]?\d* ماہ پہلے|\d+[.,]?\d* منٹ پہلے|\d+[.,]?\d* ہفتہ قبل|\d+[.,]?\d* ہفتہ میں|\d+[.,]?\d* ہفتے قبل|\d+[.,]?\d* ہفتے میں|\d+[.,]?\d* دن پہلے|\d+[.,]?\d* سال میں|\d+[.,]?\d* ماہ قبل|\d+[.,]?\d* ماہ میں|\d+[.,]?\d* منٹ قبل|\d+[.,]?\d* منٹ میں|\d+[.,]?\d* دن قبل|\d+[.,]?\d* دن میں)$`),
		KnownWords:      []string{"پچھلے مہینہ", "اگلے مہینہ", "پچھلے ہفتہ", "گزشتہ ہفتہ", "اگلے ہفتہ", "گزشتہ سال", "گزشتہ ماہ", "اس مہینہ", "اس گھنٹے", "ايندہ کل", "اگلے سال", "اگلے ماہ", "گزشتہ کل", "اس ہفتہ", "اس سال", "اس ماہ", "اس منٹ", "اکتوبر", "جمعرات", "جولايی", "سوموار", "اتوار", "اپریل", "جنوری", "دسمبر", "ستمبر", "سیکنڈ", "فروری", "مہینہ", "نومبر", "گھنٹہ", "اگست", "جمعہ", "مارچ", "منگل", "ہفتہ", "gmt", "utc", "بدھ", "جون", "سال", "ماہ", "منٹ", "ميی", "پیر", "am", "pm", "اب", "اج", "دن", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
