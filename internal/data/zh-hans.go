package data

import "regexp"

var zhHansLocale = LocaleData{
	Name:                  "zh-Hans",
	DateOrder:             "YMD",
	NoWordSpacing:         true,
	SentenceSplitterGroup: 4,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "，"},
	January:               []string{"1月", "一月"},
	February:              []string{"2月", "二月"},
	March:                 []string{"3月", "三月"},
	April:                 []string{"4月", "四月"},
	May:                   []string{"5月", "五月"},
	June:                  []string{"6月", "六月"},
	July:                  []string{"7月", "七月"},
	August:                []string{"8月", "八月"},
	September:             []string{"9月", "九月"},
	October:               []string{"10月", "十月"},
	November:              []string{"11月", "十一月"},
	December:              []string{"12月", "十二月"},
	Monday:                []string{"周一", "星期一"},
	Tuesday:               []string{"周二", "星期二"},
	Wednesday:             []string{"周三", "星期三"},
	Thursday:              []string{"周四", "星期四"},
	Friday:                []string{"周五", "星期五"},
	Saturday:              []string{"周六", "星期六"},
	Sunday:                []string{"周日", "星期日"},
	AM:                    []string{"上午"},
	PM:                    []string{"下午"},
	Year:                  []string{"年"},
	Month:                 []string{"月"},
	Week:                  []string{"周"},
	Day:                   []string{"日"},
	Hour:                  []string{"小时"},
	Minute:                []string{"分", "分钟"},
	Second:                []string{"秒"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`今天`},
		`0 hour ago`:   {`这一时间 / 此时`},
		`0 minute ago`: {`此刻`},
		`0 month ago`:  {`本月`},
		`0 second ago`: {`现在`},
		`0 week ago`:   {`本周`},
		`0 year ago`:   {`今年`},
		`1 day ago`:    {`昨天`},
		`1 month ago`:  {`上个月`},
		`1 week ago`:   {`上周`},
		`1 year ago`:   {`去年`},
		`in 1 day`:     {`明天`},
		`in 1 month`:   {`下个月`},
		`in 1 week`:    {`下周`},
		`in 1 year`:    {`明年`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+)天前`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+)小时前`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+)分钟前`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+)个月前`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+)秒前`),
			regexp.MustCompile(`(?i)(\d+)秒钟前`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+)周前`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+)年前`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+)天后`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+)小时后`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+)分钟后`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+)个月后`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+)秒后`),
			regexp.MustCompile(`(?i)(\d+)秒钟后`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+)周后`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+)年后`),
		},
	},
	LocaleSpecific: map[string]LocaleData{
		"zh-Hans-HK": zhHansHKLocale,
		"zh-Hans-MO": zhHansMOLocale,
		"zh-Hans-SG": zhHansSGLocale,
	},
}

var zhHansSGLocale = LocaleData{
	Name:                  "zh-Hans-SG",
	DateOrder:             "DMY",
}

var zhHansHKLocale = LocaleData{
	Name:                  "zh-Hans-HK",
	DateOrder:             "DMY",
}

var zhHansMOLocale = LocaleData{
	Name:                  "zh-Hans-MO",
	DateOrder:             "DMY",
}
