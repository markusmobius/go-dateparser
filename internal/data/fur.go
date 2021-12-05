package data

import "regexp"

var furLocale = LocaleData{
	Name:                  "fur",
	DateOrder:             "DMY",
	January:               []string{"zen", "zenâr"},
	February:              []string{"fev", "fevrâr"},
	March:                 []string{"mar", "març"},
	April:                 []string{"avr", "avrîl"},
	May:                   []string{"mai"},
	June:                  []string{"jug", "jugn"},
	July:                  []string{"lui"},
	August:                []string{"avo", "avost"},
	September:             []string{"set", "setembar"},
	October:               []string{"otu", "otubar"},
	November:              []string{"nov", "novembar"},
	December:              []string{"dic", "dicembar"},
	Monday:                []string{"lun", "lunis"},
	Tuesday:               []string{"mar", "martars"},
	Wednesday:             []string{"mie", "miercus"},
	Thursday:              []string{"joi", "joibe"},
	Friday:                []string{"vin", "vinars"},
	Saturday:              []string{"sab", "sabide"},
	Sunday:                []string{"dom", "domenie"},
	AM:                    []string{"a"},
	PM:                    []string{"p"},
	Year:                  []string{"an"},
	Month:                 []string{"mês"},
	Week:                  []string{"setemane"},
	Day:                   []string{"dì"},
	Hour:                  []string{"ore"},
	Minute:                []string{"minût"},
	Second:                []string{"secont"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`vuê`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`this month`},
		`0 second ago`: {`now`},
		`0 week ago`:   {`this week`},
		`0 year ago`:   {`this year`},
		`1 day ago`:    {`îr`},
		`1 month ago`:  {`last month`},
		`1 week ago`:   {`last week`},
		`1 year ago`:   {`last year`},
		`in 1 day`:     {`doman`},
		`in 1 month`:   {`next month`},
		`in 1 week`:    {`next week`},
		`in 1 year`:    {`next year`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) zornade indaûr`),
			regexp.MustCompile(`(?i)(\d+) zornadis indaûr`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ore indaûr`),
			regexp.MustCompile(`(?i)(\d+) oris indaûr`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) minût indaûr`),
			regexp.MustCompile(`(?i)(\d+) minûts indaûr`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mês indaûr`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) secont indaûr`),
			regexp.MustCompile(`(?i)(\d+) seconts indaûr`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) setemane indaûr`),
			regexp.MustCompile(`(?i)(\d+) setemanis indaûr`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) agns indaûr`),
			regexp.MustCompile(`(?i)(\d+) an indaûr`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)ca di (\d+) zornade`),
			regexp.MustCompile(`(?i)ca di (\d+) zornadis`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)ca di (\d+) ore`),
			regexp.MustCompile(`(?i)ca di (\d+) oris`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)ca di (\d+) minût`),
			regexp.MustCompile(`(?i)ca di (\d+) minûts`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)ca di (\d+) mês`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)ca di (\d+) secont`),
			regexp.MustCompile(`(?i)ca di (\d+) seconts`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)ca di (\d+) setemane`),
			regexp.MustCompile(`(?i)ca di (\d+) setemanis`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)ca di (\d+) agns`),
			regexp.MustCompile(`(?i)ca di (\d+) an`),
		},
	},
}
