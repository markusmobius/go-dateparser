// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var az_Latn_Locale = LocaleData{
	Name:      "az-Latn",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) dəqiqə ərzində(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saniyə ərzində(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) həftə ərzində(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saat ərzində(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) gün ərzində(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saniyə öncə(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ay ərzində(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) həftə öncə(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) il ərzində(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saat öncə(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)çərşənbə axşamı(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) gün öncə(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ay öncə(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) il öncə(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)bazar ertəsi(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)cümə axşamı(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)gələn həftə(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)keçən həftə(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu dəqiqə(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu həftə(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)gələn ay(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)gələn il(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)keçən ay(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)keçən il(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)sentyabr(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)çərşənbə(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu saat(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)oktyabr(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)avqust(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu gün(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)dekabr(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)fevral(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)noyabr(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)saniyə(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)yanvar(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)aprel(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)bazar(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu ay(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu il(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)dünən(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)həftə(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)sabah(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)şənbə(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)cümə(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)indi(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyul(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyun(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mart(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)saat(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)apr(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)avq(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)dek(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)fev(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)gün(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyl(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)iyn(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)may(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)noy(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)okt(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)san(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)sen(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)yan(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)am(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)ay(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)be(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ca(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)il(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)pm(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)ça(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)b(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)c(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ç(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ş(\z|\W|_)`), "${1}saturday${2}"},
	},
}
