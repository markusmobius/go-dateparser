// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var jgo_Locale = LocaleData{
	Name:      "jgo",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)ɛ́ gɛ́ mɔ (\d+) ŋgap-mbi(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɛ́ gɛ́ mɔ́ pɛsaŋ (\d+)(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɛ́ gɛ́ mɔ́ lɛ́ꞌ (\d+)(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɛ́ gɛ mɔ́ (\d+) háwa(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ɛ́gɛ́ mɔ́ ŋguꞌ (\d+)(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu ŋgap-mbi (\d+)(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́nɛ́ntúkú(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́nɛ́pfúꞌú(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu (\d+) minút(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ ntsɔ̌pmɔ́(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́nɛ́fɔm(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́nɛ́kwa(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu háwa (\d+)(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu lɛ́ꞌ (\d+)(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu ŋguꞌ (\d+)(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ ntsɔ̌ppá(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)nǔu (\d+) saŋ(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)ŋka mbɔ́t nji(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ nɛgɛ́m(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́tát(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ saambá(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)nduŋmbi saŋ(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pataa(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)pɛsaŋ pɛ́pá(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ápta mɔ́ndi(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)wɛ́nɛsɛdɛ(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)yesterday(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)fɛlâyɛdɛ(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mba'mba'(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)tomorrow(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)tɔ́sɛdɛ(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mɔ́ndi(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)second(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)sásidɛ(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sɔ́ndi(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)month(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)hour(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)lɔꞌɔ(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)week(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)year(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)day(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
	},
}
