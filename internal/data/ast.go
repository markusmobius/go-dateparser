// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ast_Locale = LocaleData{
	Name:      "ast",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []ReplacementData{
		{regexp.MustCompile(`(\A|\W|_)la selmana viniente(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) segundos(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) selmanes(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) segundos(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) selmanes(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) segundu(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) selmana(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)la selmana pasada(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) minutos(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) segundu(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) selmana(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)el mes viniente(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) minutu(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) hores(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) meses(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) hores(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) meses(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) años(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) díes(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) hora(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) selm(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)l'añu viniente(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) años(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) díes(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) hora(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) selm(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) añu(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) día(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) mes(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) seg(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)selm viniente(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)de la mañana(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)de setiembre(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)el mes pasáu(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) añu(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) día(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) mes(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) min(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) seg(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)esta selmana(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) se(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)de la tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) se(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)esti minutu(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) a(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) d(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) h(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) m(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)hai (\d+) s(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)l'añu pasáu(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)selm pasada(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)de febreru(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)de payares(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) a(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) d(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) h(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) m(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)en (\d+) s(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)d'avientu(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)d'ochobre(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)de xineru(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)de xunetu(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)esta hora(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)esta selm(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)l'añu pas(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)l'añu vin(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)miércoles(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)setiembre(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)d'agostu(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)de marzu(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)esti añu(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)esti mes(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)esti min(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)selm pas(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)selm vin(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)avientu(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)añu pas(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)añu vin(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)d'abril(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)de mayu(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)de xunu(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)domingu(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)febreru(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)mes pas(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)mes vin(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ochobre(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)payares(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)segundu(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)selmana(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)vienres(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)agostu(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)esta h(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)llunes(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)martes(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mañana(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)sábadu(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)xineru(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)xueves(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)xunetu(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)abril(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)agora(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ayeri(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)marzu(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)tarde(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)güei(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)hora(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)mayu(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)xunu(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)abr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)ago(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)avi(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)añu(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)dom(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)día(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)feb(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)llu(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)may(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)mañ(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)mes(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)mié(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)och(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)pay(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)sel(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)set(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)sáb(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)vie(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)xin(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)xnt(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)xue(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)xun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)h(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)s(\z|\W|_)`), "${1}second${2}"},
	},
}
