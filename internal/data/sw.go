// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sw_Locale = LocaleData{
	Name:      "sw",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)sekunde (\d+) zilizopita(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)sekunde (\d+) iliyopita(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya sekunde (\d+)(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya dakika (\d+)(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)miaka (\d+) iliyopita(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)miezi (\d+) iliyopita(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)siku (\d+) zilizopita(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)wiki (\d+) zilizopita(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya miaka (\d+)(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya miezi (\d+)(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya mwaka (\d+)(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya mwezi (\d+)(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)mwaka (\d+) uliopita(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi (\d+) uliopita(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)saa (\d+) zilizopita(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)siku (\d+) iliyopita(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)wiki (\d+) iliyopita(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya siku (\d+)(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya wiki (\d+)(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)saa (\d+) iliyopita(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)baada ya saa (\d+)(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)mwaka uliopita(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi uliopita(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)wiki iliyopita(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)dakika hii(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwaka ujao(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi ujao(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)wiki ijayo(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwaka huu(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi huu(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sasa hivi(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)alhamisi(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)februari(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumamosi(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumapili(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumatano(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumatatu(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)septemba(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)wiki hii(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)asubuhi(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)desemba(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)januari(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jumanne(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)novemba(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)saa hii(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekunde(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)agosti(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)aprili(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ijumaa(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mchana(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktoba(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)julai(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)kesho(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)machi(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwaka(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)mwezi(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)jana(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)juni(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)siku(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)wiki(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ago(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)des(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)leo(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mac(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mei(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)saa(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
	},
}

var sw_CD_Locale = LocaleData{
	Name:      "sw-CD",
	Parent:    &sw_Locale,
	DateOrder: "DMY",
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)juma(\z|\W|_)`), "${1}week${2}"},
	},
}

var sw_KE_Locale = LocaleData{
	Name:      "sw-KE",
	Parent:    &sw_Locale,
	DateOrder: "DMY",
}

var sw_UG_Locale = LocaleData{
	Name:      "sw-UG",
	Parent:    &sw_Locale,
	DateOrder: "DMY",
}
