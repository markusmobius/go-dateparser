package data

import "regexp"

var koLocale = LocaleData{
	Name:                  "ko",
	DateOrder:             "YMD.",
	January:               []string{"1월"},
	February:              []string{"2월"},
	March:                 []string{"3월"},
	April:                 []string{"4월"},
	May:                   []string{"5월"},
	June:                  []string{"6월"},
	July:                  []string{"7월"},
	August:                []string{"8월"},
	September:             []string{"9월"},
	October:               []string{"10월"},
	November:              []string{"11월"},
	December:              []string{"12월"},
	Monday:                []string{"월", "월요일"},
	Tuesday:               []string{"화", "화요일"},
	Wednesday:             []string{"수", "수요일"},
	Thursday:              []string{"목", "목요일"},
	Friday:                []string{"금", "금요일"},
	Saturday:              []string{"토", "토요일"},
	Sunday:                []string{"일", "일요일"},
	AM:                    []string{"am", "오전"},
	PM:                    []string{"pm", "오후"},
	Year:                  []string{"년"},
	Month:                 []string{"월"},
	Week:                  []string{"주"},
	Day:                   []string{"일"},
	Hour:                  []string{"시"},
	Minute:                []string{"분"},
	Second:                []string{"초"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`오늘`},
		`0 hour ago`:   {`현재 시간`},
		`0 minute ago`: {`현재 분`},
		`0 month ago`:  {`이번 달`},
		`0 second ago`: {`지금`},
		`0 week ago`:   {`이번 주`},
		`0 year ago`:   {`올해`},
		`1 day ago`:    {`어제`},
		`1 month ago`:  {`지난달`},
		`1 week ago`:   {`지난주`},
		`1 year ago`:   {`작년`},
		`in 1 day`:     {`내일`},
		`in 1 month`:   {`다음 달`},
		`in 1 week`:    {`다음 주`},
		`in 1 year`:    {`내년`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+)일 전`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+)시간 전`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+)분 전`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+)개월 전`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+)초 전`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+)주 전`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+)년 전`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+)일 후`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+)시간 후`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+)분 후`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+)개월 후`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+)초 후`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+)주 후`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+)년 후`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"ko-KP": koKPLocale,
	},
}

var koKPLocale = LocaleData{
	Name:                  "ko-KP",
	DateOrder:             "",
}
