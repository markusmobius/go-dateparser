// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fur_Locale = LocaleData{
	Name:      "fur",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) setemanis indaûr(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) setemane indaûr(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) zornadis indaûr(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) setemanis(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) seconts indaûr(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) zornade indaûr(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) setemane(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) zornadis(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) secont indaûr(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) seconts(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) zornade(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) minûts(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) secont(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) agns indaûr(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) oris indaûr(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) minût(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) mês indaûr(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ore indaûr(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) agns(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) oris(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) an indaûr(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) mês(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) ore(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)ca di (\d+) an(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
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
		{regexp.MustCompile(`(\A|\W|_)dicembar(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)novembar(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)setemane(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)setembar(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)domenie(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)martars(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)miercus(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fevrâr(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)otubar(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sabide(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)secont(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)vinars(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)avost(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)avrîl(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)doman(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)joibe(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lunis(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)zenâr(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)jugn(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)març(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)avo(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)avr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)dic(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)dom(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)fev(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)joi(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)jug(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)lui(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)lun(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mai(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)mie(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mês(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ore(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)otu(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sab(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)set(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)vin(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vuê(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)zen(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)an(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)dì(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)îr(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)a(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)p(\z|\W|_)`), "${1}pm${2}"},
	},
}
