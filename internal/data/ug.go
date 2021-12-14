// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ug_Locale = LocaleData{
	Name:      "ug",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) سېكۇنتتىن كېيىن(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) سائەتتىن كېيىن(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) سېكۇنت ئىلگىرى(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) مىنۇتتىن كېيىن(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ھەپتىدىن كېيىن(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) سائەت ئىلگىرى(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ھەپتە ئىلگىرى(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ئايدىن كېيىن(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) كۈندىن كېيىن(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) يىلدىن كېيىن(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ئاي ئىلگىرى(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) كۈن ئىلگىرى(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) يىل ئىلگىرى(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ئۆتكەن ھەپتە(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)چۈشتىن بۇرۇن(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)چۈشتىن كېيىن(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)كېلەر ھەپتە(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئۆتكەن ئاي(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئۆتكەن يىل(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)كېلەر ئاي(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)كېلەر يىل(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئۆكتەبىر(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)بۇ ھەپتە(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سېنتەبىر(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)سەيشەنبە(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)يەكشەنبە(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)پەيشەنبە(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)چارشەنبە(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئاۋغۇست(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)تۈنۈگۈن(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)دۈشەنبە(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)دېكابىر(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)نويابىر(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئاپرېل(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)بۇ ئاي(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)بۇ يىل(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سېكۇنت(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)فېۋرال(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)يانۋار(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئىيۇل(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئىيۇن(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)بۈگۈن(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)سائەت(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)شەنبە(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ھەپتە(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئەتە(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)جۈمە(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)مارت(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ئاي(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)كۈن(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ماي(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)يىل(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)جۈ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)دۈ(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)سە(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)شە(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)يە(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)پە(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)چا(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)چب(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)چك(\z|\W|_)`), "${1}pm${2}"},
	},
}
