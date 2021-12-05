package data

import "regexp"

var kaLocale = LocaleData{
	Name:                  "ka",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "და", "დაახლოებით", "ზე", "ის", "，"},
	January:               []string{"იან", "იანვარი"},
	February:              []string{"თებ", "თებერვალი"},
	March:                 []string{"მარ", "მარტი"},
	April:                 []string{"აპრ", "აპრილი"},
	May:                   []string{"მაი", "მაისი"},
	June:                  []string{"ივნ", "ივნისი"},
	July:                  []string{"ივლ", "ივლისი"},
	August:                []string{"აგვ", "აგვისტო"},
	September:             []string{"სექ", "სექტემბერი"},
	October:               []string{"ოქტ", "ოქტომბერი"},
	November:              []string{"ნოე", "ნოემბერი"},
	December:              []string{"დეკ", "დეკემბერი"},
	Monday:                []string{"ორშ", "ორშაბათი"},
	Tuesday:               []string{"სამ", "სამშაბათი"},
	Wednesday:             []string{"ოთხ", "ოთხშაბათი"},
	Thursday:              []string{"ხუთ", "ხუთშაბათი"},
	Friday:                []string{"პარ", "პარასკევი"},
	Saturday:              []string{"შაბ", "შაბათი"},
	Sunday:                []string{"კვი", "კვირა"},
	AM:                    []string{"am"},
	PM:                    []string{"pm", "შუადღ შემდეგ"},
	Year:                  []string{"წ", "წელი", "წლის"},
	Month:                 []string{"თვე"},
	Week:                  []string{"კვ", "კვირა"},
	Day:                   []string{"დღე"},
	Hour:                  []string{"საათი", "სთ"},
	Minute:                []string{"წთ", "წუთი"},
	Second:                []string{"წამი", "წმ"},
	In:                    []string{"დღეიდან"},
	Ago:                   []string{"წინ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`დღეს`},
		`0 hour ago`:   {`ამ საათში`},
		`0 minute ago`: {`ამ წუთში`},
		`0 month ago`:  {`ამ თვეში`},
		`0 second ago`: {`ახლა`},
		`0 week ago`:   {`ამ კვირაში`},
		`0 year ago`:   {`ამ წელს`},
		`1 day ago`:    {`გუშინ`},
		`1 month ago`:  {`გასულ თვეს`},
		`1 week ago`:   {`გასულ კვირაში`},
		`1 year ago`:   {`გასულ წელს`},
		`in 1 day`:     {`ხვალ`},
		`in 1 month`:   {`მომავალ თვეს`},
		`in 1 week`:    {`მომავალ კვირაში`},
		`in 1 year`:    {`მომავალ წელს`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) დღის წინ`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) საათის წინ`),
			regexp.MustCompile(`(?i)(\d+) სთ წინ`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) წთ წინ`),
			regexp.MustCompile(`(?i)(\d+) წუთის წინ`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) თვის წინ`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) წამის წინ`),
			regexp.MustCompile(`(?i)(\d+) წმ წინ`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) კვ წინ`),
			regexp.MustCompile(`(?i)(\d+) კვირის წინ`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) წლის წინ`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) დღეში`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) საათში`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) წუთში`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) თვეში`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) წამში`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) კვირაში`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) წელიწადში`),
			regexp.MustCompile(`(?i)(\d+) წელში`),
		},
	},
	Simplifications: map[string]string{
		`ერთ`: `1`,
	},
}
