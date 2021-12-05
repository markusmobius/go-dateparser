package data

import "regexp"

var cyLocale = LocaleData{
	Name:                  "cy",
	DateOrder:             "DMY",
	January:               []string{"ion", "ionawr"},
	February:              []string{"chw", "chwef", "chwefror"},
	March:                 []string{"maw", "mawrth"},
	April:                 []string{"ebr", "ebrill"},
	May:                   []string{"mai"},
	June:                  []string{"meh", "mehefin"},
	July:                  []string{"gor", "gorff", "gorffennaf"},
	August:                []string{"awst"},
	September:             []string{"medi"},
	October:               []string{"hyd", "hydref"},
	November:              []string{"tach", "tachwedd"},
	December:              []string{"rhag", "rhagfyr"},
	Monday:                []string{"dydd llun", "llun"},
	Tuesday:               []string{"dydd mawrth", "maw"},
	Wednesday:             []string{"dydd mercher", "mer"},
	Thursday:              []string{"dydd iau", "iau"},
	Friday:                []string{"dydd gwener", "gwe", "gwen"},
	Saturday:              []string{"dydd sadwrn", "sad"},
	Sunday:                []string{"dydd sul", "sul"},
	AM:                    []string{"yb"},
	PM:                    []string{"yh"},
	Year:                  []string{"bl", "blwyddyn"},
	Month:                 []string{"mis"},
	Week:                  []string{"wythnos"},
	Day:                   []string{"dydd"},
	Hour:                  []string{"awr"},
	Minute:                []string{"mun", "munud"},
	Second:                []string{"eiliad"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`heddiw`},
		`0 hour ago`:   {`yr awr hon`},
		`0 minute ago`: {`y funud hon`},
		`0 month ago`:  {`y mis hwn`},
		`0 second ago`: {`nawr`},
		`0 week ago`:   {`yr wythnos hon`},
		`0 year ago`:   {`eleni`},
		`1 day ago`:    {`ddoe`},
		`1 month ago`:  {`mis diwethaf`},
		`1 week ago`:   {`wythnos ddiwethaf`},
		`1 year ago`:   {`llynedd`},
		`in 1 day`:     {`yfory`},
		`in 1 month`:   {`mis nesaf`},
		`in 1 week`:    {`wythnos nesaf`},
		`in 1 year`:    {`blwyddyn nesaf`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) diwrnod yn ôl`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) awr yn ôl`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) mun yn ôl`),
			regexp.MustCompile(`(?i)(\d+) munud yn ôl`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mis yn ôl`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) eiliad yn ôl`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) wythnos yn ôl`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) o flynyddoedd yn ôl`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)ymhen (\d+) diwrnod`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)ymhen (\d+) awr`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)ymhen (\d+) mun`),
			regexp.MustCompile(`(?i)ymhen (\d+) munud`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)ymhen (\d+) mis`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)ymhen (\d+) eiliad`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)ymhen (\d+) wythnos`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)ymhen (\d+) mlynedd`),
		},
	},
}
