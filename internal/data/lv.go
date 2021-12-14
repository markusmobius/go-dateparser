// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lv_Locale = LocaleData{
	Name:      "lv",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) mēnešiem(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) sekundes(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) sekundēm(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) nedēļas(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) nedēļām(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) stundas(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) stundām(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) dienas(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) dienām(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) gadiem(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) mēneša(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) mēnešiem(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) sekundes(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) sekundēm(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)pagājušajā mēnesī(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)pagājušajā nedēļā(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) minūtes(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) minūtēm(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) nedēļas(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) nedēļām(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) stundas(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) stundām(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) gada(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) dienas(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) dienām(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) gadiem(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) mēneša(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)nākamajā mēnesī(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)nākamajā nedēļā(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)pagājušajā gadā(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) mēn(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) ned(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) sek(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) st(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)priekšpusdiena(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)priekšpusdienā(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) gada(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)nākamajā gadā(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) d(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) g(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) h(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pirms (\d+) s(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) min(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) mēn(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) ned(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) sek(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) st(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)ceturtdiena(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) d(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) g(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) h(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēc (\d+) s(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)pēcpusdiena(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēcpusdienā(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)šajā minūtē(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)šajā mēnesī(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)šajā nedēļā(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)šajā stundā(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)piektdiena(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)septembris(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)decembris(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)februāris(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)novembris(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)pirmdiena(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sestdiena(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)svētdiena(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)trešdiena(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)šajā gadā(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)janvāris(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktobris(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)otrdiena(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekundes(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)aprīlis(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)augusts(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ceturtd(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mēnesis(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)priekšp(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēcpusd(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)stundas(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)jūlijs(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jūnijs(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)nedēļa(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)piektd(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)šodien(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)diena(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)maijs(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)marts(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)pirmd(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sestd(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)svētd(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tagad(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)trešd(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vakar(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)febr(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)gads(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)janv(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)otrd(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)pēcp(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)sept(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dec(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)jūl(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)jūn(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mēn(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ned(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)rīt(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)st(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)d(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)g(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`), "${1}second${2}"},
	},
}
