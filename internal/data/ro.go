// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ro_Locale = LocaleData{
	Name:                  "ro",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "de", "la", "|"},
	January:               []string{"ian", "ianuarie"},
	February:              []string{"feb", "febr", "februarie"},
	March:                 []string{"mar", "mart", "martie"},
	April:                 []string{"apr", "aprilie"},
	May:                   []string{"mai"},
	June:                  []string{"iun", "iunie"},
	July:                  []string{"iul", "iulie"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "sept", "septembrie"},
	October:               []string{"oct", "octombrie"},
	November:              []string{"noiem", "noiembrie", "nov"},
	December:              []string{"dec", "decembrie"},
	Monday:                []string{"lun", "luni"},
	Tuesday:               []string{"mar", "marți"},
	Wednesday:             []string{"mi", "mie", "miercuri"},
	Thursday:              []string{"joi"},
	Friday:                []string{"vin", "vineri"},
	Saturday:              []string{"sâm", "sâmbătă"},
	Sunday:                []string{"dum", "duminică"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"an", "ani"},
	Month:                 []string{"luna", "luni", "lună"},
	Week:                  []string{"săpt", "săptămâni", "săptămână"},
	Day:                   []string{"zi", "zile"},
	Hour:                  []string{"h", "ore", "oră"},
	Minute:                []string{"m", "min", "minut", "minute"},
	Second:                []string{"s", "sec", "secunde", "secundă"},
	In:                    []string{"în"},
	Ago:                   []string{"în urmă"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`azi`},
		`0 hour ago`:   {`ora aceasta`},
		`0 minute ago`: {`minutul acesta`},
		`0 month ago`:  {`luna aceasta`},
		`0 second ago`: {`acum`},
		`0 week ago`:   {`săptămâna aceasta`},
		`0 year ago`:   {`anul acesta`},
		`1 day ago`:    {`ieri`},
		`1 month ago`:  {`luna trecută`},
		`1 week ago`:   {`săptămâna trecută`},
		`1 year ago`:   {`anul trecut`},
		`in 1 day`:     {`mâine`},
		`in 1 month`:   {`luna viitoare`},
		`in 1 week`:    {`săptămâna viitoare`},
		`in 1 year`:    {`anul viitor`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de zile`),
			regexp.MustCompile(`(?i)acum (\d+) zi`),
			regexp.MustCompile(`(?i)acum (\d+) zile`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de ore`),
			regexp.MustCompile(`(?i)acum (\d+) h`),
			regexp.MustCompile(`(?i)acum (\d+) oră`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de minute`),
			regexp.MustCompile(`(?i)acum (\d+) min`),
			regexp.MustCompile(`(?i)acum (\d+) minut`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de luni`),
			regexp.MustCompile(`(?i)acum (\d+) luni`),
			regexp.MustCompile(`(?i)acum (\d+) lună`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de secunde`),
			regexp.MustCompile(`(?i)acum (\d+) sec`),
			regexp.MustCompile(`(?i)acum (\d+) secundă`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de săptămâni`),
			regexp.MustCompile(`(?i)acum (\d+) săpt`),
			regexp.MustCompile(`(?i)acum (\d+) săptămână`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)acum (\d+) an`),
			regexp.MustCompile(`(?i)acum (\d+) de ani`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)peste (\d+) de zile`),
			regexp.MustCompile(`(?i)peste (\d+) zi`),
			regexp.MustCompile(`(?i)peste (\d+) zile`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)peste (\d+) de ore`),
			regexp.MustCompile(`(?i)peste (\d+) h`),
			regexp.MustCompile(`(?i)peste (\d+) oră`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)peste (\d+) de minute`),
			regexp.MustCompile(`(?i)peste (\d+) min`),
			regexp.MustCompile(`(?i)peste (\d+) minut`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)peste (\d+) de luni`),
			regexp.MustCompile(`(?i)peste (\d+) luni`),
			regexp.MustCompile(`(?i)peste (\d+) lună`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)peste (\d+) de secunde`),
			regexp.MustCompile(`(?i)peste (\d+) sec`),
			regexp.MustCompile(`(?i)peste (\d+) secundă`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)peste (\d+) de săptămâni`),
			regexp.MustCompile(`(?i)peste (\d+) săpt`),
			regexp.MustCompile(`(?i)peste (\d+) săptămână`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)peste (\d+) an`),
			regexp.MustCompile(`(?i)peste (\d+) ani`),
			regexp.MustCompile(`(?i)peste (\d+) de ani`),
		},
	},
}

var ro_MD_Locale = LocaleData{
	Name:                  "ro-MD",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "de", "la", "|"},
	January:               []string{"ian", "ianuarie"},
	February:              []string{"feb", "febr", "februarie"},
	March:                 []string{"mar", "mart", "martie"},
	April:                 []string{"apr", "aprilie"},
	May:                   []string{"mai"},
	June:                  []string{"iun", "iunie"},
	July:                  []string{"iul", "iulie"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "sept", "septembrie"},
	October:               []string{"oct", "octombrie"},
	November:              []string{"noiem", "noiembrie", "nov"},
	December:              []string{"dec", "decembrie"},
	Monday:                []string{"lun", "luni"},
	Tuesday:               []string{"mar", "marți"},
	Wednesday:             []string{"mi", "mie", "miercuri"},
	Thursday:              []string{"joi"},
	Friday:                []string{"vin", "vineri"},
	Saturday:              []string{"sâm", "sâmbătă"},
	Sunday:                []string{"dum", "duminică"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"an", "ani"},
	Month:                 []string{"luna", "luni", "lună"},
	Week:                  []string{"săpt", "săptămâni", "săptămână"},
	Day:                   []string{"zi", "zile"},
	Hour:                  []string{"h", "ore", "oră"},
	Minute:                []string{"m", "min", "minut", "minute"},
	Second:                []string{"s", "sec", "secunde", "secundă"},
	In:                    []string{"în"},
	Ago:                   []string{"în urmă"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`azi`},
		`0 hour ago`:   {`ora aceasta`},
		`0 minute ago`: {`minutul acesta`},
		`0 month ago`:  {`luna aceasta`},
		`0 second ago`: {`acum`},
		`0 week ago`:   {`săptămâna aceasta`},
		`0 year ago`:   {`anul acesta`},
		`1 day ago`:    {`ieri`},
		`1 month ago`:  {`luna trecută`},
		`1 week ago`:   {`săptămâna trecută`},
		`1 year ago`:   {`anul trecut`},
		`in 1 day`:     {`mâine`},
		`in 1 month`:   {`luna viitoare`},
		`in 1 week`:    {`săptămâna viitoare`},
		`in 1 year`:    {`anul viitor`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de zile`),
			regexp.MustCompile(`(?i)acum (\d+) zi`),
			regexp.MustCompile(`(?i)acum (\d+) zile`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de ore`),
			regexp.MustCompile(`(?i)acum (\d+) h`),
			regexp.MustCompile(`(?i)acum (\d+) oră`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de minute`),
			regexp.MustCompile(`(?i)acum (\d+) min`),
			regexp.MustCompile(`(?i)acum (\d+) minut`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de luni`),
			regexp.MustCompile(`(?i)acum (\d+) luni`),
			regexp.MustCompile(`(?i)acum (\d+) lună`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de secunde`),
			regexp.MustCompile(`(?i)acum (\d+) sec`),
			regexp.MustCompile(`(?i)acum (\d+) secundă`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)acum (\d+) de săptămâni`),
			regexp.MustCompile(`(?i)acum (\d+) săpt`),
			regexp.MustCompile(`(?i)acum (\d+) săptămână`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)acum (\d+) an`),
			regexp.MustCompile(`(?i)acum (\d+) de ani`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)peste (\d+) de zile`),
			regexp.MustCompile(`(?i)peste (\d+) zi`),
			regexp.MustCompile(`(?i)peste (\d+) zile`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)peste (\d+) de ore`),
			regexp.MustCompile(`(?i)peste (\d+) h`),
			regexp.MustCompile(`(?i)peste (\d+) oră`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)peste (\d+) de minute`),
			regexp.MustCompile(`(?i)peste (\d+) min`),
			regexp.MustCompile(`(?i)peste (\d+) minut`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)peste (\d+) de luni`),
			regexp.MustCompile(`(?i)peste (\d+) luni`),
			regexp.MustCompile(`(?i)peste (\d+) lună`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)peste (\d+) de secunde`),
			regexp.MustCompile(`(?i)peste (\d+) sec`),
			regexp.MustCompile(`(?i)peste (\d+) secundă`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)peste (\d+) de săptămâni`),
			regexp.MustCompile(`(?i)peste (\d+) săpt`),
			regexp.MustCompile(`(?i)peste (\d+) săptămână`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)peste (\d+) an`),
			regexp.MustCompile(`(?i)peste (\d+) ani`),
			regexp.MustCompile(`(?i)peste (\d+) de ani`),
		},
	},
}
