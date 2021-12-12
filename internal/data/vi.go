// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var vi_Locale = LocaleData{
	Name:                  "vi",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "，"},
	PertainWords:          []string{"lúc"},
	January:               []string{"thg 1", "tháng 1", "tháng một"},
	February:              []string{"thg 2", "tháng 2", "tháng hai"},
	March:                 []string{"thg 3", "tháng 3", "tháng ba"},
	April:                 []string{"thg 4", "tháng 4", "tháng tư"},
	May:                   []string{"thg 5", "tháng 5", "tháng năm"},
	June:                  []string{"thg 6", "tháng 6", "tháng sáu"},
	July:                  []string{"thg 7", "tháng 7", "tháng bảy"},
	August:                []string{"thg 8", "tháng 8", "tháng tám"},
	September:             []string{"thg 9", "tháng 9", "tháng chín"},
	October:               []string{"thg 10", "tháng 10", "tháng mười"},
	November:              []string{"thg 11", "tháng 11", "tháng mười một"},
	December:              []string{"thg 12", "tháng 12", "tháng mười hai"},
	Monday:                []string{"th 2", "thứ 2", "thứ hai"},
	Tuesday:               []string{"th 3", "thứ 3", "thứ ba"},
	Wednesday:             []string{"th 4", "thứ 4", "thứ tư"},
	Thursday:              []string{"th 5", "thứ 5", "thứ năm"},
	Friday:                []string{"th 6", "thứ 6", "thứ sáu"},
	Saturday:              []string{"th 7", "thứ 7", "thứ bảy"},
	Sunday:                []string{"chủ nhật", "cn", "thứ 1"},
	AM:                    []string{"sa"},
	PM:                    []string{"ch"},
	Year:                  []string{"năm"},
	Month:                 []string{"thang", "thg", "tháng"},
	Week:                  []string{"tuần", "tuần lể"},
	Day:                   []string{"ban ngày", "buổi", "ngày"},
	Hour:                  []string{"giờ"},
	Minute:                []string{"chút", "lát", "nguyên bản", "phút"},
	Second:                []string{"giây", "giây đồng hồ", "hạng nhì"},
	In:                    []string{"trong"},
	Ago:                   []string{"cách đây", "trước", "trước đây", "trước"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hôm nay`},
		`0 hour ago`:   {`giờ này`},
		`0 minute ago`: {`phút này`},
		`0 month ago`:  {`tháng này`},
		`0 second ago`: {`bây giờ`},
		`0 week ago`:   {`tuần này`},
		`0 year ago`:   {`năm nay`},
		`1 day ago`:    {`hôm qua`},
		`1 month ago`:  {`tháng trước`},
		`1 week ago`:   {`tuần trước`},
		`1 year ago`:   {`năm ngoái`},
		`in 1 day`:     {`ngày mai`},
		`in 1 month`:   {`tháng sau`},
		`in 1 week`:    {`tuần sau`},
		`in 1 year`:    {`năm sau`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ngày trước`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) giờ trước`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) phút trước`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) tháng trước`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) giây trước`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) tuần trước`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) năm trước`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)sau (\d+) ngày nữa`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)sau (\d+) giờ nữa`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)sau (\d+) phút nữa`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)sau (\d+) tháng nữa`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)sau (\d+) giây nữa`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)sau (\d+) tuần nữa`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)sau (\d+) năm nữa`),
		},
	},
	Simplifications: map[string]string{
		`(?:ngày|năm)\s(\d+)`: `\1`,
	},
}