// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var se_Locale = LocaleData{
	Name:      "se",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) mánotbadji maŋŋilit(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) minuhtta maŋŋilit(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekundda maŋŋilit(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) diibmur maŋŋilit(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jándora maŋŋilit(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) minuhta maŋŋilit(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekunda maŋŋilit(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) diibmu maŋŋilit(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jahkki maŋŋilit(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jándor maŋŋilit(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) mánotbadji árat(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) vahkku maŋŋilit(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jahki maŋŋilit(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) vahku maŋŋilit(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekundda árat(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) diibmur árat(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jándora árat(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekunda árat(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) diibmu árat(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jahkki árat(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jándor árat(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) vahkku árat(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jahki árat(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) vahku árat(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)eahketbeaivet(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ođđajagemánnu(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)eahketbeaivi(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)golggotmánnu(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)suoidnemánnu(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)gaskavahkku(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)geassemánnu(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)guovvamánnu(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)iđitbeaivet(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)juovlamánnu(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)miessemánnu(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)njukčamánnu(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)skábmamánnu(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)sotnabeaivi(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)this minute(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)borgemánnu(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)cuoŋománnu(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)iđitbeaivi(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)last month(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)maŋŋebárga(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)next month(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)this month(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)čakčamánnu(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)bearjadat(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)duorasdat(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)last week(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)last year(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)lávvardat(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)next week(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)next year(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)this hour(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this week(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)this year(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)vuossárga(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekunda(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)beaivi(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)diibmu(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)ihttin(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)váhkku(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)jáhki(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)mánnu(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)bear(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)borg(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)duor(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)gask(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)geas(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)golg(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)guov(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ikte(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)juov(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)mies(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)njuk(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)odne(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ođđj(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)skáb(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)sotn(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)suoi(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)vuos(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)čakč(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)cuo(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)láv(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)maŋ(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)now(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eb(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ib(\z|\W|_)`), "${1}am${2}"},
	},
}

var se_FI_Locale = LocaleData{
	Name:      "se-FI",
	Parent:    &se_Locale,
	DateOrder: "YMD",
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) jagi siste(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) jagi árat(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)boahtte jagi(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)mannan jagi(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)maŋŋebárgga(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)bearjadaga(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)duorastaga(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)gaskavahku(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)lávvardaga(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vuossárgga(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dán jagi(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)vahkku(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)jahki(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)j(\z|\W|_)`), "${1}year${2}"},
	},
}

var se_SE_Locale = LocaleData{
	Name:      "se-SE",
	Parent:    &se_Locale,
	DateOrder: "YMD",
}
