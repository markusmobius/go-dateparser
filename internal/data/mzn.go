// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mzn_Locale = LocaleData{
	Name:      "mzn",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) ثانیه دله(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ثانیه پیش(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) دقیقه دله(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ساعِت دله(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ساعِت پیش(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) دَقه دله(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ساعت دله(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ساعت پیش(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) هفته دله(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) هفته پیش(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) روز دله(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) روز پیش(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) سال دله(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) سال پیش(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ماه دله(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ماه پیش(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)بعدی هفته(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)قبلی هفته(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)این هفته(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سال دیگه(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماه ِبعد(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)این ماه(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سپتامبر(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماه قبل(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)اَمروز(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)دسامبر(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)نوامبر(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)پارسال(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ژانویه(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)آوریل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)امسال(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)اکتبر(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ثانیه(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)دیروز(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعِت(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)فوریه(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)فِردا(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ژوئیه(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ساعت(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)مارس(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)هفته(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ژوئن(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)fri(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mon(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sat(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sun(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)thu(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tue(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)wed(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)اوت(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)روز(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)سال(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماه(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)مه(\z|\W|_)`), "${1}may${2}"},
	},
}
