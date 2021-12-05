package data

import "regexp"

var guLocale = LocaleData{
	Name:                  "gu",
	DateOrder:             "DMY",
	January:               []string{"જાન્યુ", "જાન્યુઆરી"},
	February:              []string{"ફેબ્રુ", "ફેબ્રુઆરી"},
	March:                 []string{"માર્ચ"},
	April:                 []string{"એપ્રિલ"},
	May:                   []string{"મે"},
	June:                  []string{"જૂન"},
	July:                  []string{"જુલાઈ"},
	August:                []string{"ઑગસ્ટ"},
	September:             []string{"સપ્ટે", "સપ્ટેમ્બર"},
	October:               []string{"ઑક્ટો", "ઑક્ટોબર"},
	November:              []string{"નવે", "નવેમ્બર"},
	December:              []string{"ડિસે", "ડિસેમ્બર"},
	Monday:                []string{"સોમ", "સોમવાર"},
	Tuesday:               []string{"મંગળ", "મંગળવાર"},
	Wednesday:             []string{"બુધ", "બુધવાર"},
	Thursday:              []string{"ગુરુ", "ગુરુવાર"},
	Friday:                []string{"શુક્ર", "શુક્રવાર"},
	Saturday:              []string{"શનિ", "શનિવાર"},
	Sunday:                []string{"રવિ", "રવિવાર"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"વ", "વર્ષ"},
	Month:                 []string{"મ", "મહિનો"},
	Week:                  []string{"અઠ", "અઠવાડિયું"},
	Day:                   []string{"દિવસ"},
	Hour:                  []string{"ક", "કલાક"},
	Minute:                []string{"મિ", "મિનિટ"},
	Second:                []string{"સે", "સેકન્ડ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`આજે`},
		`0 hour ago`:   {`આ કલાક`},
		`0 minute ago`: {`આ મિનિટ`},
		`0 month ago`:  {`આ મહિને`},
		`0 second ago`: {`હમણાં`},
		`0 week ago`:   {`આ અઠવાડિયે`},
		`0 year ago`:   {`આ વર્ષે`},
		`1 day ago`:    {`ગઈકાલે`},
		`1 month ago`:  {`ગયા મહિને`},
		`1 week ago`:   {`ગયા અઠવાડિયે`},
		`1 year ago`:   {`ગયા વર્ષે`},
		`in 1 day`:     {`આવતીકાલે`},
		`in 1 month`:   {`આવતા મહિને`},
		`in 1 week`:    {`આવતા અઠવાડિયે`},
		`in 1 year`:    {`આવતા વર્ષે`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) દિવસ પહેલાં`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) કલાક પહેલાં`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) મિનિટ પહેલાં`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) મહિના પહેલાં`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) સેકંડ પહેલાં`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) અઠ પહેલાં`),
			regexp.MustCompile(`(?i)(\d+) અઠવાડિયા પહેલાં`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) વર્ષ પહેલા`),
			regexp.MustCompile(`(?i)(\d+) વર્ષ પહેલાં`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) દિવસમાં`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) કલાકમાં`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) મિનિટમાં`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) મહિનામાં`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) સેકંડમાં`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) અઠ માં`),
			regexp.MustCompile(`(?i)(\d+) અઠવાડિયામાં`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) વર્ષમાં`),
		},
	},
}
