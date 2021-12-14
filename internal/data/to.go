// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var to_Locale = LocaleData{
	Name:      "to",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)māhina 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)sekoni 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)'aho 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he miniti 'e (\d+)(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he māhina 'e (\d+)(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he sekoni 'e (\d+)(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)houa 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ta'u 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)uike 'e (\d+) kuo'osi(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he 'aho 'e (\d+)(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he houa 'e (\d+)(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he ta'u 'e (\d+)(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)'i he uike 'e (\d+)(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)māhina kuo'osi(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)māhina kaha'u(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)'apongipongi(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)ta'u kuo'osi(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)tu'apulelulu(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)uike kuo'osi(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ta'u kaha'u(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)uike kaha'u(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)hengihengi(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)māhiná ni(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)'epeleli(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)'okatopa(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pulelulu(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sepitema(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)taimí ni(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)tokonaki(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)'ahó ni(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)'aneafi(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)'aokosi(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)falaite(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fēpueli(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)sānuali(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)ta'ú ni(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)uiké ni(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)efiafi(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ma'asi(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)māhina(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)mōnite(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)nōvema(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekoni(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)siulai(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)sāpate(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tīsema(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)tūsite(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)'aho(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)'aok(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)'epe(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)'oka(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)houa(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ma'a(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)sune(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)ta'u(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)tu'a(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)uike(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)fal(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fēp(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)mōn(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)nōv(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)pul(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sep(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)siu(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)sun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)sān(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)sāp(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tok(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tīs(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)tūs(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ea(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)hh(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)mē(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
	},
}
