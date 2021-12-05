package data

import "regexp"

var mnLocale = LocaleData{
	Name:                  "mn",
	DateOrder:             "YMD",
	January:               []string{"1-р сар", "нэгдүгээр сар"},
	February:              []string{"2-р сар", "хоёрдугаар сар"},
	March:                 []string{"3-р сар", "гуравдугаар сар"},
	April:                 []string{"4-р сар", "дөрөвдүгээр сар"},
	May:                   []string{"5-р сар", "тавдугаар сар"},
	June:                  []string{"6-р сар", "зургадугаар сар"},
	July:                  []string{"7-р сар", "долдугаар сар"},
	August:                []string{"8-р сар", "наймдугаар сар"},
	September:             []string{"9-р сар", "есдүгээр сар"},
	October:               []string{"10-р сар", "аравдугаар сар"},
	November:              []string{"11-р сар", "арван нэгдүгээр сар"},
	December:              []string{"12-р сар", "арван хоёрдугаар сар"},
	Monday:                []string{"да", "даваа"},
	Tuesday:               []string{"мя", "мягмар"},
	Wednesday:             []string{"лх", "лхагва"},
	Thursday:              []string{"пү", "пүрэв"},
	Friday:                []string{"ба", "баасан"},
	Saturday:              []string{"бя", "бямба"},
	Sunday:                []string{"ня", "ням"},
	AM:                    []string{"үө"},
	PM:                    []string{"үх"},
	Year:                  []string{"жил"},
	Month:                 []string{"сар"},
	Week:                  []string{"7х", "долоо хоног"},
	Day:                   []string{"өдөр"},
	Hour:                  []string{"ц", "цаг"},
	Minute:                []string{"мин", "минут"},
	Second:                []string{"сек", "секунд"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`өнөөдөр`},
		`0 hour ago`:   {`энэ цаг`},
		`0 minute ago`: {`энэ минут`},
		`0 month ago`:  {`энэ сар`},
		`0 second ago`: {`одоо`},
		`0 week ago`:   {`энэ долоо хоног`},
		`0 year ago`:   {`энэ жил`},
		`1 day ago`:    {`өчигдөр`},
		`1 month ago`:  {`өнгөрсөн сар`},
		`1 week ago`:   {`өнгөрсөн долоо хоног`},
		`1 year ago`:   {`өнгөрсөн жил`},
		`in 1 day`:     {`маргааш`},
		`in 1 month`:   {`ирэх сар`},
		`in 1 week`:    {`ирэх долоо хоног`},
		`in 1 year`:    {`ирэх жил`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) өдрийн өмнө`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ц өмнө`),
			regexp.MustCompile(`(?i)(\d+) цагийн өмнө`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) мин өмнө`),
			regexp.MustCompile(`(?i)(\d+) минутын өмнө`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) сарын өмнө`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) сек өмнө`),
			regexp.MustCompile(`(?i)(\d+) секундын өмнө`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) 7х-ийн өмнө`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) жилийн өмнө`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) өдрийн дараа`),
			regexp.MustCompile(`(?i)(\d+) өдөрт`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ц дараа`),
			regexp.MustCompile(`(?i)(\d+) цагийн дараа`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) мин дараа`),
			regexp.MustCompile(`(?i)(\d+) минутын дараа`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) сарын дараа`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) сек дараа`),
			regexp.MustCompile(`(?i)(\d+) секундын дараа`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) 7х-ийн дараа`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) жилийн дараа`),
		},
	},
}
