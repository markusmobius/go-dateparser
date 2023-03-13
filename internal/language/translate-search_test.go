package language

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestTranslateSearch(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Language    string
		Text        string
		Translation string
	}

	tests := []testScenario{
		// English
		{"en", "Sep 03 2014", "september 03 2014"},
		{"en", "friday, 03 september 2014", "friday, 03 september 2014"},
		{"en", "Aug 06, 2018 05:05 PM CDT", "august 06, 2018 05:05 pm cdt"},
		// Chinese
		{"zh", "1年11个月", "1year11month"},
		{"zh", "1年11個月", "1year11month"},
		{"zh", "2015年04月08日10点05", "2015-04-08 10:05"},
		{"zh", "2015年04月08日10:05", "2015-04-08 10:05"},
		{"zh", "2013年04月08日", "2013-04-08"},
		{"zh", "周一", "monday"},
		{"zh", "礼拜一", "monday"},
		{"zh", "周二", "tuesday"},
		{"zh", "礼拜二", "tuesday"},
		{"zh", "周三", "wednesday"},
		{"zh", "礼拜三", "wednesday"},
		{"zh", "星期日 2015年04月08日10:05", "sunday 2015-04-08 10:05"},
		{"zh", "周六 2013年04月08日", "saturday 2013-04-08"},
		{"zh", "下午3:30", "3:30 pm"},
		{"zh", "凌晨3:30", "3:30 am"},
		{"zh", "中午", "12:00"},
		// French
		{"fr", "20 Février 2012", "20 february 2012"},
		{"fr", "Mercredi 19 Novembre 2013", "wednesday 19 november 2013"},
		{"fr", "18 octobre 2012 à 19 h 21 min", "18 october 2012 19:21"},
		// German
		{"de", "29. Juni 2007", "29. june 2007"},
		{"de", "Montag 5 Januar, 2015", "monday 5 january, 2015"},
		// Hungarian
		{"hu", "2016 augusztus 11", "2016 august 11"},
		{"hu", "2016-08-13 szombat 10:21", "2016-08-13 saturday 10:21"},
		{"hu", "2016. augusztus 14. vasárnap 10:21", "2016. august 14. sunday 10:21"},
		{"hu", "hétfő", "monday"},
		{"hu", "tegnapelőtt", "2 day ago"},
		{"hu", "ma", "0 day ago"},
		{"hu", "2 hónappal ezelőtt", "2 month ago"},
		{"hu", "2016-08-13 szombat 10:21 GMT", "2016-08-13 saturday 10:21 gmt"},
		// Spanish
		{"es", "Miércoles 31 Diciembre 2014", "wednesday 31 december 2014"},
		// Italian
		{"it", "Giovedi Maggio 29 2013", "thursday may 29 2013"},
		{"it", "19 Luglio 2013", "19 july 2013"},
		// Portuguese
		{"pt", "22 de dezembro de 2014 às 02:38", "22 december 2014 02:38"},
		// Russian
		{"ru", "5 августа 2014 г в 12:00", "5 august 2014 year 12:00"},
		// Real: {"ru", "5 августа 2014 г. в 12:00"},
		// Turkish
		{"tr", "2 Ocak 2015 Cuma, 16:49", "2 january 2015 friday, 16:49"},
		// Czech
		{"cs", "22. prosinec 2014 v 2:38", "22. december 2014 2:38"},
		// Dutch
		{"nl", "maandag 22 december 2014 om 2:38", "monday 22 december 2014 2:38"},
		// Romanian
		{"ro", "22 Decembrie 2014 la 02:38", "22 december 2014 02:38"},
		// Polish
		{"pl", "4 stycznia o 13:50", "4 january 13:50"},
		{"pl", "29 listopada 2014 o 08:40", "29 november 2014 08:40"},
		// Ukrainian
		{"uk", "30 листопада 2013 о 04:27", "30 november 2013 04:27"},
		// Belarusian
		{"be", "5 снежня 2015 г у 12:00", "5 december 2015 year 12:00"},
		// Real: {"be", "5 снежня 2015 г. у 12:00"), Issue: Abbreviation segmentation.
		{"be", "11 верасня 2015 г у 12:11", "11 september 2015 year 12:11"},
		// Real: {"be", "11 верасня 2015 г. у 12:11"},
		{"be", "3 стд 2015 г у 10:33", "3 january 2015 year 10:33"},
		// Real: {"be", "3 стд 2015 г. у 10:33"},
		// Arabic
		{"ar", "6 يناير، 2015، الساعة 05:16 مساءً", "6 january، 2015، 05:16 pm"},
		{"ar", "7 يناير، 2015، الساعة 11:00 صباحاً", "7 january، 2015، 11:00 am"},
		// Vietnamese
		// Disabled - wrong segmentation at "Thứ Năm"
		// {"vi", "Thứ Năm, ngày 8 tháng 1 năm 2015"},
		// Disabled - wrong segmentation at "Thứ Tư"
		// {"vi", "Thứ Tư, 07/01/2015 | 22:34"},
		{"vi", "9 Tháng 1 2015 lúc 15:08", "9 january 2015 15:08"},
		// Thai
		// Disabled - spacing differences
		// {"th", "เมื่อ กุมภาพันธ์ 09, 2015, 09:27:57 AM"},
		// {"th", "เมื่อ กรกฎาคม 05, 2012, 01:18:06 AM"},
		// Tagalog
		{"tl", "Biyernes Hulyo 3, 2015", "friday july 3, 2015"},
		{"tl", "Pebrero 5, 2015 7:00 pm", "february 5, 2015 7:00 pm"},
		// Indonesian
		{"id", "06 Sep 2015", "06 september 2015"},
		{"id", "07 Feb 2015 20:15", "07 february 2015 20:15"},
		// Miscellaneous
		{"en", "2014-12-12T12:33:39-08:00", "2014-12-12t12:33:39-08:00"},
		{"en", "2014-10-15T16:12:20+00:00", "2014-10-15t16:12:20+00:00"},
		{"en", "28 Oct 2014 16:39:01 +0000", "28 october 2014 16:39:01 +0000"},
		// Disabled - wrong split at "a las".
		// {"es", "13 Febrero 2015 a las 23:00"},
		// Danish
		{"da", "Sep 03 2014", "september 03 2014"},
		{"da", "fredag, 03 september 2014", "friday, 03 september 2014"},
		{"da", "fredag d. 3 september 2014", "friday 3 september 2014"},
		// Finnish
		{"fi", "maanantai tammikuu 16, 2015", "monday january 16, 2015"},
		{"fi", "ma tammi 16, 2015", "monday january 16, 2015"},
		{"fi", "tiistai helmikuu 16, 2015", "tuesday february 16, 2015"},
		{"fi", "ti helmi 16, 2015", "tuesday february 16, 2015"},
		{"fi", "keskiviikko maaliskuu 16, 2015", "wednesday march 16, 2015"},
		{"fi", "ke maalis 16, 2015", "wednesday march 16, 2015"},
		{"fi", "torstai huhtikuu 16, 2015", "thursday april 16, 2015"},
		{"fi", "to huhti 16, 2015", "thursday april 16, 2015"},
		{"fi", "perjantai toukokuu 16, 2015", "friday may 16, 2015"},
		{"fi", "pe touko 16, 2015", "friday may 16, 2015"},
		{"fi", "lauantai kesäkuu 16, 2015", "saturday june 16, 2015"},
		{"fi", "la kesä 16, 2015", "saturday june 16, 2015"},
		{"fi", "sunnuntai heinäkuu 16, 2015", "sunday july 16, 2015"},
		{"fi", "su heinä 16, 2015", "sunday july 16, 2015"},
		{"fi", "su elokuu 16, 2015", "sunday august 16, 2015"},
		{"fi", "su elo 16, 2015", "sunday august 16, 2015"},
		{"fi", "su syyskuu 16, 2015", "sunday september 16, 2015"},
		{"fi", "su syys 16, 2015", "sunday september 16, 2015"},
		{"fi", "su lokakuu 16, 2015", "sunday october 16, 2015"},
		{"fi", "su loka 16, 2015", "sunday october 16, 2015"},
		{"fi", "su marraskuu 16, 2015", "sunday november 16, 2015"},
		{"fi", "su marras 16, 2015", "sunday november 16, 2015"},
		{"fi", "su joulukuu 16, 2015", "sunday december 16, 2015"},
		{"fi", "su joulu 16, 2015", "sunday december 16, 2015"},
		{"fi", "1. tammikuuta, 2016", "1. january, 2016"},
		{"fi", "tiistaina, 27. lokakuuta 2015", "tuesday, 27. october 2015"},
		// Japanese
		{"ja", "午後3時", "pm3:00"},
		{"ja", "2時", "2:00"},
		{"ja", "11時42分", "11:42"},
		{"ja", "3ヶ月", "3month"},
		{"ja", "約53か月前", "53monthago"},
		{"ja", "3月", "3month"},
		{"ja", "十二月", "december"},
		{"ja", "2月10日", "2-10"},
		{"ja", "2013年2月", "2013year2month"},
		{"ja", "2013年04月08日", "2013-04-08"},
		{"ja", "2016年03月24日 木曜日 10時05分", "2016-03-24 thursday 10:05"},
		{"ja", "2016年3月20日 21時40分", "2016-3-20 21:40"},
		{"ja", "2016年03月21日 23時05分11秒", "2016-03-21 23:05:11"},
		{"ja", "2016年3月21日(月) 14時48分", "2016-3-21monday 14:48"},
		{"ja", "2016年3月20日(日) 21時40分", "2016-3-20sunday 21:40"},
		{"ja", "2016年3月20日 (日) 21時40分", "2016-3-20 sunday 21:40"},
		// Hebrew
		{"he", "20 לאפריל 2012", "20 april 2012"},
		{"he", "יום רביעי ה-19 בנובמבר 2013", "wednesday ה-19 november 2013"},
		{"he", "18 לאוקטובר 2012 בשעה 19:21", "18 october 2012 19:21"},
		// Disabled - wrong split at "יום ה'".
		// {"he", "יום ה' 6/10/2016"},
		{"he", "חצות", "12 am"},
		{"he", "1 אחר חצות", "1 am"},
		{"he", "3 לפנות בוקר", "3 am"},
		{"he", "3 בבוקר", "3 am"},
		{"he", "3 בצהריים", "3 pm"},
		{"he", "6 לפנות ערב", "6 pm"},
		{"he", "6 אחרי הצהריים", "6 pm"},
		{"he", "6 אחרי הצהרים", "6 pm"},
		// Bangla
		{"bn", "সেপ্টেম্বর 03 2014", "september 03 2014"},
		{"bn", "শুক্রবার, 03 সেপ্টেম্বর 2014", "friday, 03 september 2014"},
		// Hindi
		{"hi", "सोमवार 13 जून 1998", "monday 13 june 1998"},
		{"hi", "मंगल 16 1786 12:18", "tuesday 16 1786 12:18"},
		{"hi", "शनि 11 अप्रैल 2002 03:09", "saturday 11 april 2002 03:09"},
		// Swedish
		{"sv", "Sept 03 2014", "september 03 2014"},
		{"sv", "fredag, 03 september 2014", "friday, 03 september 2014"},
	}

	// Prepare config
	cfg := &setting.Configuration{SkipTokens: []string{"t"}}

	// Start test
	var nFailed int
	for _, test := range tests {
		// Prepare message
		message := fmt.Sprintf("%s \"%s\" => \"%s\"", test.Language, test.Text, test.Translation)

		// Translate text
		ld, _ := data.GetLocaleData(test.Language)
		translation, original := TranslateSearch(cfg, ld, test.Text)

		// Assert result
		passed := assert.NotZero(t, translation, message)
		if passed {
			passed = assert.Equal(t, len(original), len(translation), message)
		}
		if passed {
			passed = assert.Equal(t, test.Text, original[0], message)
		}
		if passed {
			passed = assert.Equal(t, test.Translation, translation[0], message)
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
