package data

import "regexp"

var bnLocale = LocaleData{
	Name:                  "bn",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 3,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "ই", "এবং", "টা", "প্রায়", "লা", "শে", "，"},
	January:               []string{"জানু", "জানুয়ারী", "জানুয়ারি"},
	February:              []string{"ফেব", "ফেব্রুয়ারী", "ফেব্রুয়ারি"},
	March:                 []string{"মার্চ"},
	April:                 []string{"এপ্রিল"},
	May:                   []string{"মে"},
	June:                  []string{"জুন"},
	July:                  []string{"জুলাই"},
	August:                []string{"আগস্ট"},
	September:             []string{"সেপ্টেম্বর"},
	October:               []string{"অক্টোবর"},
	November:              []string{"নভেম্বর"},
	December:              []string{"ডিসেম্বর"},
	Monday:                []string{"সোম", "সোমবার"},
	Tuesday:               []string{"মঙ্গল", "মঙ্গলবার"},
	Wednesday:             []string{"বুধ", "বুধবার"},
	Thursday:              []string{"বৃহষ্পতিবার", "বৃহস্পতি", "বৃহস্পতিবার"},
	Friday:                []string{"শুক্র", "শুক্রবার"},
	Saturday:              []string{"শনি", "শনিবার"},
	Sunday:                []string{"রবি", "রবিবার"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"বছর", "বৎসর", "সাল"},
	Month:                 []string{"মাস"},
	Week:                  []string{"সপ্তাহ"},
	Day:                   []string{"দিন"},
	Hour:                  []string{"ঘন্টা"},
	Minute:                []string{"মিনিট"},
	Second:                []string{"সেকেন্ড"},
	In:                    []string{"মধ্যে"},
	Ago:                   []string{"আগে"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`আজ`},
		`0 hour ago`:   {`এই ঘণ্টায়`},
		`0 minute ago`: {`এই মিনিট`},
		`0 month ago`:  {`এই মাস`},
		`0 second ago`: {`এখন`},
		`0 week ago`:   {`এই সপ্তাহ`},
		`0 year ago`:   {`এই বছর`},
		`1 day ago`:    {`গতকাল`},
		`1 month ago`:  {`গত মাস`},
		`1 week ago`:   {`গত সপ্তাহ`},
		`1 year ago`:   {`গত বছর`},
		`2 day ago`:    {`গত পরশু`},
		`in 1 day`:     {`আগামীকাল`},
		`in 1 month`:   {`আগামী মাস`, `পরের মাস`},
		`in 1 week`:    {`আগামী সপ্তাহ`, `পরের সপ্তাহ`},
		`in 1 year`:    {`আগামী বছর`, `পরের বছর`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) দিন আগে`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ঘন্টা আগে`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) মিনিট আগে`),
			regexp.MustCompile(`(?i)(\d+) মিনিট পূর্বে`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) মাস আগে`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) সেকেন্ড আগে`),
			regexp.MustCompile(`(?i)(\d+) সেকেন্ড পূর্বে`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) সপ্তাহ আগে`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) বছর পূর্বে`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) দিনের মধ্যে`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ঘন্টায়`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) মিনিটে`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) মাসে`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) সেকেন্ডে`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) সপ্তাহে`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) বছরে`),
		},
	},
	Simplifications: map[string]string{
		`মধ্যরাত`:  `00:00`,
		`মধ্যাহ্ন`: `12:00`,
	},
	LocaleSpecific: map[string]LocaleData{
		"bn-IN": bnINLocale,
	},
}

var bnINLocale = LocaleData{
	Name:                  "bn-IN",
	DateOrder:             "",
}
