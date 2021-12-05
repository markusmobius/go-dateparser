package data

import "regexp"

var kmLocale = LocaleData{
	Name:                  "km",
	DateOrder:             "DMY",
	January:               []string{"មករា"},
	February:              []string{"កុម្ភៈ"},
	March:                 []string{"មីនា"},
	April:                 []string{"មេសា"},
	May:                   []string{"ឧសភា"},
	June:                  []string{"មិថុនា"},
	July:                  []string{"កក្កដា"},
	August:                []string{"សីហា"},
	September:             []string{"កញ្ញា"},
	October:               []string{"តុលា"},
	November:              []string{"វិច្ឆិកា"},
	December:              []string{"ធ្នូ"},
	Monday:                []string{"ច័ន្ទ"},
	Tuesday:               []string{"អង្គារ"},
	Wednesday:             []string{"ពុធ"},
	Thursday:              []string{"ព្រហស្បតិ៍"},
	Friday:                []string{"សុក្រ"},
	Saturday:              []string{"សៅរ៍"},
	Sunday:                []string{"អាទិត្យ"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ឆ្នាំ"},
	Month:                 []string{"ខែ"},
	Week:                  []string{"សប្ដាហ៍"},
	Day:                   []string{"ថ្ងៃ"},
	Hour:                  []string{"ម៉ោង"},
	Minute:                []string{"នាទី"},
	Second:                []string{"វិនាទី"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ថ្ងៃ​នេះ`},
		`0 hour ago`:   {`ម៉ោងនេះ`},
		`0 minute ago`: {`នាទីនេះ`},
		`0 month ago`:  {`ខែ​នេះ`},
		`0 second ago`: {`ឥឡូវ`},
		`0 week ago`:   {`សប្ដាហ៍​នេះ`},
		`0 year ago`:   {`ឆ្នាំ​នេះ`},
		`1 day ago`:    {`ម្សិលមិញ`},
		`1 month ago`:  {`ខែ​មុន`},
		`1 week ago`:   {`សប្ដាហ៍​មុន`},
		`1 year ago`:   {`ឆ្នាំ​មុន`},
		`in 1 day`:     {`ថ្ងៃស្អែក`, `ថ្ងៃ​ស្អែក`},
		`in 1 month`:   {`ខែ​ក្រោយ`},
		`in 1 week`:    {`សប្ដាហ៍​ក្រោយ`},
		`in 1 year`:    {`ឆ្នាំ​ក្រោយ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ថ្ងៃ​មុន`),
			regexp.MustCompile(`(?i)(\d+) ថ្ងៃ​​មុន`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ម៉ោង​មុន`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) នាទី​មុន`),
			regexp.MustCompile(`(?i)(\d+) នាទី​​មុន`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ខែមុន`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) វិនាទី​មុន`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) សប្ដាហ៍​មុន`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ឆ្នាំ​មុន`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) ថ្ងៃទៀត`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ម៉ោងទៀត`),
			regexp.MustCompile(`(?i)ក្នុង​រយៈ​ពេល (\d+) ម៉ោង`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) នាទីទៀត`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ខែទៀត`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) វិនាទីទៀត`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) សប្ដាហ៍ទៀត`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) ឆ្នាំទៀត`),
		},
	},
}
