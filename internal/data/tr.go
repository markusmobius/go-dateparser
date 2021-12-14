// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var tr_Locale = LocaleData{
	Name:      "tr",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "ve", "yaklaşık", "|"},
	Translations: []TranslationData{
		{regexp.MustCompile(`(\A|\W|_)(\d+) dakika sonra(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saniye sonra(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hafta sonra(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saniye önce(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hafta önce(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saat sonra(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)önümüzdeki hafta(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) gün sonra(\z|\W|_)`), "${1}in ${2} day${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) saat önce(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) yıl sonra(\z|\W|_)`), "${1}in ${2} year${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ay sonra(\z|\W|_)`), "${1}in ${2} month${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) dk sonra(\z|\W|_)`), "${1}in ${2} minute${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) gün önce(\z|\W|_)`), "${1}${2} day ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hf sonra(\z|\W|_)`), "${1}in ${2} week${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sa sonra(\z|\W|_)`), "${1}in ${2} hour${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sn sonra(\z|\W|_)`), "${1}in ${2} second${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) yıl önce(\z|\W|_)`), "${1}${2} year ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)önümüzdeki gün(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)önümüzdeki yıl(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) ay önce(\z|\W|_)`), "${1}${2} month ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) hf önce(\z|\W|_)`), "${1}${2} week ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sa önce(\z|\W|_)`), "${1}${2} hour ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)(\d+) sn önce(\z|\W|_)`), "${1}${2} second ago${3}"},
		{regexp.MustCompile(`(\A|\W|_)gelecek hafta(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)önümüzdeki ay(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)gelecek yıl(\z|\W|_)`), "${1}in 1 year${2}"},
		{regexp.MustCompile(`(\A|\W|_)geçen hafta(\z|\W|_)`), "${1}1 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)gelecek ay(\z|\W|_)`), "${1}in 1 month${2}"},
		{regexp.MustCompile(`(\A|\W|_)içerisinde(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu dakika(\z|\W|_)`), "${1}0 minute ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)cumartesi(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)geçen gün(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)geçen yıl(\z|\W|_)`), "${1}1 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)pazartesi(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu hafta(\z|\W|_)`), "${1}0 week ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)geçen ay(\z|\W|_)`), "${1}1 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)perşembe(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)çarşamba(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ağustos(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu saat(\z|\W|_)`), "${1}0 hour ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)haftaya(\z|\W|_)`), "${1}in 1 week${2}"},
		{regexp.MustCompile(`(\A|\W|_)haziran(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)aralık(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu yıl(\z|\W|_)`), "${1}0 year ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)içinde(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)saniye(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)temmuz(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)bu ay(\z|\W|_)`), "${1}0 month ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)bugün(\z|\W|_)`), "${1}0 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eylül(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)hafta(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)kasım(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)mayıs(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)nisan(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)pazar(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sonra(\z|\W|_)`), "${1}in${2}"},
		{regexp.MustCompile(`(\A|\W|_)yarın(\z|\W|_)`), "${1}in 1 day${2}"},
		{regexp.MustCompile(`(\A|\W|_)şimdi(\z|\W|_)`), "${1}0 second ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)şubat(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)cuma(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)ekim(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)mart(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)ocak(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)saat(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)salı(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sene(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)önce(\z|\W|_)`), "${1}ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)ara(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ağu(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)cmt(\z|\W|_)`), "${1}saturday${2}"},
		{regexp.MustCompile(`(\A|\W|_)cum(\z|\W|_)`), "${1}friday${2}"},
		{regexp.MustCompile(`(\A|\W|_)dün(\z|\W|_)`), "${1}1 day ago${2}"},
		{regexp.MustCompile(`(\A|\W|_)eki(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)eyl(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)gün(\z|\W|_)`), "${1}day${2}"},
		{regexp.MustCompile(`(\A|\W|_)haz(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)kas(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)mar(\z|\W|_)`), "${1}march${2}"},
		{regexp.MustCompile(`(\A|\W|_)may(\z|\W|_)`), "${1}may${2}"},
		{regexp.MustCompile(`(\A|\W|_)nis(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)oca(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)paz(\z|\W|_)`), "${1}sunday${2}"},
		{regexp.MustCompile(`(\A|\W|_)per(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)prs(\z|\W|_)`), "${1}thursday${2}"},
		{regexp.MustCompile(`(\A|\W|_)pzt(\z|\W|_)`), "${1}monday${2}"},
		{regexp.MustCompile(`(\A|\W|_)sal(\z|\W|_)`), "${1}tuesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)tem(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)yıl(\z|\W|_)`), "${1}year${2}"},
		{regexp.MustCompile(`(\A|\W|_)çar(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)çrs(\z|\W|_)`), "${1}wednesday${2}"},
		{regexp.MustCompile(`(\A|\W|_)şub(\z|\W|_)`), "${1}february${2}"},
		{regexp.MustCompile(`(\A|\W|_)ar(\z|\W|_)`), "${1}december${2}"},
		{regexp.MustCompile(`(\A|\W|_)ay(\z|\W|_)`), "${1}month${2}"},
		{regexp.MustCompile(`(\A|\W|_)ağ(\z|\W|_)`), "${1}august${2}"},
		{regexp.MustCompile(`(\A|\W|_)ek(\z|\W|_)`), "${1}october${2}"},
		{regexp.MustCompile(`(\A|\W|_)ey(\z|\W|_)`), "${1}september${2}"},
		{regexp.MustCompile(`(\A|\W|_)ha(\z|\W|_)`), "${1}june${2}"},
		{regexp.MustCompile(`(\A|\W|_)hf(\z|\W|_)`), "${1}week${2}"},
		{regexp.MustCompile(`(\A|\W|_)ka(\z|\W|_)`), "${1}november${2}"},
		{regexp.MustCompile(`(\A|\W|_)ni(\z|\W|_)`), "${1}april${2}"},
		{regexp.MustCompile(`(\A|\W|_)oc(\z|\W|_)`), "${1}january${2}"},
		{regexp.MustCompile(`(\A|\W|_)sa(\z|\W|_)`), "${1}hour${2}"},
		{regexp.MustCompile(`(\A|\W|_)sn(\z|\W|_)`), "${1}second${2}"},
		{regexp.MustCompile(`(\A|\W|_)te(\z|\W|_)`), "${1}july${2}"},
		{regexp.MustCompile(`(\A|\W|_)ös(\z|\W|_)`), "${1}pm${2}"},
		{regexp.MustCompile(`(\A|\W|_)öö(\z|\W|_)`), "${1}am${2}"},
		{regexp.MustCompile(`(\A|\W|_)şu(\z|\W|_)`), "${1}february${2}"},
	},
}

var tr_CY_Locale = LocaleData{
	Name:      "tr-CY",
	Parent:    &tr_Locale,
	DateOrder: "DMY",
}
