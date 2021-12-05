package data

import "regexp"

var sqLocale = LocaleData{
	Name:                  "sq",
	DateOrder:             "DMY",
	January:               []string{"jan", "janar"},
	February:              []string{"shk", "shkurt"},
	March:                 []string{"mar", "mars"},
	April:                 []string{"pri", "prill"},
	May:                   []string{"maj"},
	June:                  []string{"qer", "qershor"},
	July:                  []string{"kor", "korrik"},
	August:                []string{"gsh", "gusht"},
	September:             []string{"sht", "shtator"},
	October:               []string{"tet", "tetor"},
	November:              []string{"nën", "nëntor"},
	December:              []string{"dhj", "dhjetor"},
	Monday:                []string{"e hënë", "hën"},
	Tuesday:               []string{"e martë", "mar"},
	Wednesday:             []string{"e mërkurë", "mër"},
	Thursday:              []string{"e enjte", "enj"},
	Friday:                []string{"e premte", "pre"},
	Saturday:              []string{"e shtunë", "sht"},
	Sunday:                []string{"die", "e diel"},
	AM:                    []string{"e paradites", "paradite"},
	PM:                    []string{"e pasdites", "pasdite"},
	Year:                  []string{"vit"},
	Month:                 []string{"muaj"},
	Week:                  []string{"javë"},
	Day:                   []string{"ditë"},
	Hour:                  []string{"orë"},
	Minute:                []string{"min", "minutë"},
	Second:                []string{"sek", "sekondë"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`sot`},
		`0 hour ago`:   {`këtë orë`},
		`0 minute ago`: {`këtë minutë`},
		`0 month ago`:  {`këtë muaj`},
		`0 second ago`: {`tani`},
		`0 week ago`:   {`këtë javë`},
		`0 year ago`:   {`këtë vit`},
		`1 day ago`:    {`dje`},
		`1 month ago`:  {`muajin e kaluar`},
		`1 week ago`:   {`javën e kaluar`},
		`1 year ago`:   {`vitin e kaluar`},
		`in 1 day`:     {`nesër`},
		`in 1 month`:   {`muajin e ardhshëm`},
		`in 1 week`:    {`javën e ardhshme`},
		`in 1 year`:    {`vitin e ardhshëm`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ditë më parë`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) orë më parë`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min më parë`),
			regexp.MustCompile(`(?i)(\d+) minuta më parë`),
			regexp.MustCompile(`(?i)(\d+) minutë më parë`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) muaj më parë`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) sek më parë`),
			regexp.MustCompile(`(?i)(\d+) sekonda më parë`),
			regexp.MustCompile(`(?i)(\d+) sekondë më parë`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) javë më parë`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) vit më parë`),
			regexp.MustCompile(`(?i)(\d+) vjet më parë`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)pas (\d+) dite`),
			regexp.MustCompile(`(?i)pas (\d+) ditësh`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)pas (\d+) ore`),
			regexp.MustCompile(`(?i)pas (\d+) orësh`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)pas (\d+) min`),
			regexp.MustCompile(`(?i)pas (\d+) minutash`),
			regexp.MustCompile(`(?i)pas (\d+) minute`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)pas (\d+) muaji`),
			regexp.MustCompile(`(?i)pas (\d+) muajsh`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)pas (\d+) sek`),
			regexp.MustCompile(`(?i)pas (\d+) sekondash`),
			regexp.MustCompile(`(?i)pas (\d+) sekonde`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)pas (\d+) jave`),
			regexp.MustCompile(`(?i)pas (\d+) javësh`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)pas (\d+) viti`),
			regexp.MustCompile(`(?i)pas (\d+) vjetësh`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"sq-MK": sqMKLocale,
		"sq-XK": sqXKLocale,
	},
}

var sqMKLocale = LocaleData{
	Name:                  "sq-MK",
	DateOrder:             "",
}

var sqXKLocale = LocaleData{
	Name:                  "sq-XK",
	DateOrder:             "",
}
