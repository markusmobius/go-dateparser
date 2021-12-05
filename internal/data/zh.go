package data

import "regexp"

var zhLocale = LocaleData{
	Name:                  "zh",
	DateOrder:             "YMD",
	NoWordSpacing:         true,
	SentenceSplitterGroup: 4,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "约", "，"},
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
	Monday:                []string{"周一", "星期一", "礼拜一"},
	Tuesday:               []string{"周二", "星期二", "礼拜二"},
	Wednesday:             []string{"周三", "星期三", "礼拜三"},
	Thursday:              []string{"周四", "星期四", "礼拜四"},
	Friday:                []string{"周五", "星期五", "礼拜五"},
	Saturday:              []string{"周六", "星期六", "礼拜六"},
	Sunday:                []string{"周日", "星期天", "星期日", "礼拜天", "礼拜日"},
	AM:                    []string{"上午"},
	PM:                    []string{"下午"},
	Year:                  []string{"年"},
	Month:                 []string{"个月", "個月", "月"},
	Week:                  []string{"周", "星期"},
	Day:                   []string{"天", "日"},
	Hour:                  []string{"小时"},
	Minute:                []string{"分", "分钟"},
	Second:                []string{"秒"},
	In:                    []string{"在"},
	Ago:                   []string{"前"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`今天`},
		`0 hour ago`:   {`这一时间 / 此时`},
		`0 minute ago`: {`此刻`},
		`0 month ago`:  {`本月`},
		`0 second ago`: {`刚刚`, `现在`},
		`0 week ago`:   {`本周`},
		`0 year ago`:   {`今年`},
		`1 day ago`:    {`昨天`},
		`1 month ago`:  {`上个月`},
		`1 week ago`:   {`上周`},
		`1 year ago`:   {`去年`},
		`2 day ago`:    {`前天`},
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
	Simplifications: map[string]string{
		`(?:上午|早上|凌晨)(?:\s*)(\d+)(?:\s*):(?:\s+|:)?(\d+)`:                              `\1:\2 am`,
		`(?:中午|下午|(?:晚上?))(?:\s*)(\d+)(?:\s*):(?:\s+|:)?(\d+)`:                         `\1:\2 pm`,
		`(\d+)年(?:\s+)?(\d+)月(?:\s+)?(\d+)日(?:\s+)?(\d+)时(?:\s+)?(\d+)分`:               `\1-\2-\3 \4:\5`,
		`(\d+)年(?:\s+)?(\d+)月(?:\s+)?(\d{1,2})(?:日)?`:                                  `\1-\2-\3`,
		`(\d+)年(?:\s+)?(\d+)月(?:\s+)?(\d{1,2})(?:日)?(?:\s+)?(\d{1,2})(?:点|:)(\d{1,2})`: `\1-\2-\3 \4:\5`,
		`(\d+)月(?=.*[前后])`: `\1 月`,
		`中午`:               `12:00`,
		`半小时前`:             `30分前`,
	},
}
