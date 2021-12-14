// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var et_Locale = LocaleData{
	Name:      "et",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekundi pärast(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) minuti pärast(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) nädala pärast(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) aasta pärast(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) päeva pärast(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sekundi eest(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) tunni pärast(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) nädala eest(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)praegusel minutil(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) aasta eest(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) kuu pärast(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) min pärast(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) näd pärast(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) päeva eest(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sek pärast(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) tunni eest(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)praegusel tunnil(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) a pärast(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) k pärast(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) kuu eest(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) näd eest(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) p pärast(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) s pärast(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sek eest(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) t pärast(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)järgmine aasta(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)järgmine nädal(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)käesolev aasta(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)käesolev nädal(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eelmine aasta(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eelmine nädal(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) a eest(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) k eest(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) p eest(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) s eest(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) t eest(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)järgmine kuu(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)käesolev kuu(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eelmine kuu(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)detsember(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)esmaspäev(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)kolmapäev(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)neljapäev(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)september(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)teisipäev(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)november(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktoober(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pühapäev(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)veebruar(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)jaanuar(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)laupäev(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)aprill(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)august(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)sekund(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)aasta(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)homme(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)juuli(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)juuni(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)märts(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)nädal(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)reede(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)veebr(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)dets(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)eile(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)jaan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)nüüd(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)päev(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)sept(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)tund(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)täna(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)aug(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)kuu(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)mai(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)nov(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)näd(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)sek(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)a(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)e(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)k(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)l(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)n(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)p(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)r(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)t(\z|\W|_)`), "${1}tuesday${2}"},
	},
}
