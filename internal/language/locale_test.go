package language

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestTranslation(t *testing.T) {
	type testScenario struct {
		Locale   string
		String   string
		Expected string
	}

	tests := []testScenario{
		// English
		{"en", "Sep 03 2014", "september 03 2014"},
		{"en", "friday, 03 september 2014", "friday 03 september 2014"},
		// Chinese
		{"zh", "1年11个月", "1 year 11 month"},
		{"zh", "1年11個月", "1 year 11 month"},
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
		{"fr", "18 octobre 2012 à 19 h 21 min", "18 october 2012  19:21"},
		// German
		{"de", "29. Juni 2007", "29. june 2007"},
		{"de", "Montag 5 Januar, 2015", "monday 5 january 2015"},
		{"de", "vor einer Woche", "1 week ago"},
		{"de", "in zwei Monaten", "in 2 month"},
		{"de", "übermorgen", "in 2 day"},
		{"de", "3 mrz 1999", "3 march 1999"},
		// Hungarian
		{"hu", "2016 augusztus 11.", "2016 august 11."},
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
		{"pt", "22 de dezembro de 2014 às 02:38", "22  december  2014  02:38"},
		// Russian
		{"ru", "5 августа 2014 г. в 12:00", "5 august 2014 year.  12:00"},
		// Turkish
		{"tr", "2 Ocak 2015 Cuma, 16:49", "2 january 2015 friday 16:49"},
		// Czech
		{"cs", "22. prosinec 2014 v 2:38", "22. december 2014  2:38"},
		// Dutch
		{"nl", "maandag 22 december 2014 om 2:38", "monday 22 december 2014  2:38"},
		// Romanian
		{"ro", "22 Decembrie 2014 la 02:38", "22 december 2014  02:38"},
		// Polish
		{"pl", "4 stycznia o 13:50", "4 january  13:50"},
		{"pl", "29 listopada 2014 o 08:40", "29 november 2014  08:40"},
		// Ukrainian
		{"uk", "30 листопада 2013 о 04:27", "30 november 2013  04:27"},
		// Belarusian
		{"be", "5 снежня 2015 г. у 12:00", "5 december 2015 year.  12:00"},
		{"be", "11 верасня 2015 г. у 12:11", "11 september 2015 year.  12:11"},
		{"be", "3 стд 2015 г. у 10:33", "3 january 2015 year.  10:33"},
		// Arabic
		{"ar", "6 يناير، 2015، الساعة 05:16 مساءً", "6 january 2015 05:16 pm"},
		{"ar", "7 يناير، 2015، الساعة 11:00 صباحاً", "7 january 2015 11:00 am"},
		// Vietnamese
		{"vi", "Thứ Năm, ngày 8 tháng 1 năm 2015", "thursday 8 january 2015"},
		{"vi", "Thứ Tư, 07/01/2015 | 22:34", "wednesday 07/01/2015  22:34"},
		{"vi", "9 Tháng 1 2015 lúc 15:08", "9 january 2015  15:08"},
		// Thai
		{"th", "เมื่อ กุมภาพันธ์ 09, 2015, 09:27:57 AM", "february 09 2015 09:27:57 am"},
		{"th", "เมื่อ กรกฎาคม 05, 2012, 01:18:06 AM", "july 05 2012 01:18:06 am"},
		{"th", "วันเสาร์ที่ 16 ธันวาคม 2560 7:00 pm", "saturday 16 december 2560 7:00 pm"},
		{"th", "วันอาทิตย์ที่ 17 ธันวาคม 2560 6:00 pm", "sunday 17 december 2560 6:00 pm"},
		// Tagalog
		{"tl", "Biyernes Hulyo 3, 2015", "friday july 3 2015"},
		{"tl", "Pebrero 5, 2015 7:00 pm", "february 5 2015 7:00 pm"},
		// Indonesian
		{"id", "06 Sep 2015", "06 september 2015"},
		{"id", "07 Feb 2015 20:15", "07 february 2015 20:15"},
		// Miscellaneous
		{"en", "2014-12-12T12:33:39-08:00", "2014-12-12 12:33:39-08:00"},
		{"en", "2014-10-15T16:12:20+00:00", "2014-10-15 16:12:20+00:00"},
		{"en", "28 Oct 2014 16:39:01 +0000", "28 october 2014 16:39:01 +0000"},
		{"es", "13 Febrero 2015 a las 23:00", "13 february 2015  23:00"},
		// Danish
		{"da", "Sep 03 2014", "september 03 2014"},
		{"da", "fredag, 03 september 2014", "friday 03 september 2014"},
		{"da", "fredag d. 3 september 2014", "friday  3 september 2014"},
		// Finnish
		{"fi", "maanantai tammikuu 16, 2015", "monday january 16 2015"},
		{"fi", "ma tammi 16, 2015", "monday january 16 2015"},
		{"fi", "tiistai helmikuu 16, 2015", "tuesday february 16 2015"},
		{"fi", "ti helmi 16, 2015", "tuesday february 16 2015"},
		{"fi", "keskiviikko maaliskuu 16, 2015", "wednesday march 16 2015"},
		{"fi", "ke maalis 16, 2015", "wednesday march 16 2015"},
		{"fi", "torstai huhtikuu 16, 2015", "thursday april 16 2015"},
		{"fi", "to huhti 16, 2015", "thursday april 16 2015"},
		{"fi", "perjantai toukokuu 16, 2015", "friday may 16 2015"},
		{"fi", "pe touko 16, 2015", "friday may 16 2015"},
		{"fi", "lauantai kesäkuu 16, 2015", "saturday june 16 2015"},
		{"fi", "la kesä 16, 2015", "saturday june 16 2015"},
		{"fi", "sunnuntai heinäkuu 16, 2015", "sunday july 16 2015"},
		{"fi", "su heinä 16, 2015", "sunday july 16 2015"},
		{"fi", "su elokuu 16, 2015", "sunday august 16 2015"},
		{"fi", "su elo 16, 2015", "sunday august 16 2015"},
		{"fi", "su syyskuu 16, 2015", "sunday september 16 2015"},
		{"fi", "su syys 16, 2015", "sunday september 16 2015"},
		{"fi", "su lokakuu 16, 2015", "sunday october 16 2015"},
		{"fi", "su loka 16, 2015", "sunday october 16 2015"},
		{"fi", "su marraskuu 16, 2015", "sunday november 16 2015"},
		{"fi", "su marras 16, 2015", "sunday november 16 2015"},
		{"fi", "su joulukuu 16, 2015", "sunday december 16 2015"},
		{"fi", "su joulu 16, 2015", "sunday december 16 2015"},
		{"fi", "1. tammikuuta, 2016", "1. january 2016"},
		{"fi", "tiistaina, 27. lokakuuta 2015", "tuesday 27. october 2015"},
		// Japanese
		{"ja", "午後3時", "pm 3:00"},
		{"ja", "2時", "2:00"},
		{"ja", "11時42分", "11:42"},
		{"ja", "3ヶ月", "3 month"},
		{"ja", "約53か月前", "53 month ago"},
		{"ja", "3月", "march"},
		{"ja", "十二月", "december"},
		{"ja", "2月10日", "2-10"},
		{"ja", "2013年2月", "2013 year february"},
		{"ja", "2013年04月08日", "2013-04-08"},
		{"ja", "2016年03月24日 木曜日 10時05分", "2016-03-24 thursday 10:05"},
		{"ja", "2016年3月20日 21時40分", "2016-3-20 21:40"},
		{"ja", "2016年03月21日 23時05分11秒", "2016-03-21 23:05:11"},
		{"ja", "2016年3月21日(月) 14時48分", "2016-3-21 monday 14:48"},
		{"ja", "2016年3月20日(日) 21時40分", "2016-3-20 sunday 21:40"},
		{"ja", "2016年3月20日 (日) 21時40分", "2016-3-20 sunday 21:40"},
		// Hebrew
		{"he", "20 לאפריל 2012", "20 april 2012"},
		{"he", "יום רביעי ה-19 בנובמבר 2013", "wednesday 19 november 2013"},
		{"he", "18 לאוקטובר 2012 בשעה 19:21", "18 october 2012  19:21"},
		{"he", "יום ה' 6/10/2016", "thursday 6/10/2016"},
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
		{"bn", "শুক্রবার, 03 সেপ্টেম্বর 2014", "friday 03 september 2014"},
		// Hindi
		{"hi", "सोमवार 13 जून 1998", "monday 13 june 1998"},
		{"hi", "मंगल 16 1786 12:18", "tuesday 16 1786 12:18"},
		{"hi", "शनि 11 अप्रैल 2002 03:09", "saturday 11 april 2002 03:09"},
		// Swedish
		{"sv", "Sept 03 2014", "september 03 2014"},
		{"sv", "fredag, 03 september 2014", "friday 03 september 2014"},
		// af
		{"af", "5 Mei 2017", "5 may 2017"},
		{"af", "maandag, Augustus 15 2005 10 vm", "monday august 15 2005 10 am"},
		// agq
		{"agq", "12 ndzɔ̀ŋɔ̀tɨ̀fʉ̀ghàdzughù 1999", "12 september 1999"},
		{"agq", "tsuʔndzɨkɔʔɔ 14 see 10 ak", "saturday 14 may 10 pm"},
		// ak
		{"ak", "esusow aketseaba-kɔtɔnimba", "may"},
		{"ak", "8 mumu-ɔpɛnimba ben", "8 december tuesday"},
		// am
		{"am", "ፌብሩወሪ 22 8:00 ጥዋት", "february 22 8:00 am"},
		{"am", "ኖቬም 10", "november 10"},
		// as
		{"as", "17 জানুৱাৰী 1885", "17 january 1885"},
		{"as", "বৃহষ্পতিবাৰ 1 জুলাই 2009", "thursday 1 july 2009"},
		// asa
		{"asa", "12 julai 1879 08:00 ichamthi", "12 july 1879 08:00 pm"},
		{"asa", "jpi 2 desemba 2007 01:00 icheheavo", "sunday 2 december 2007 01:00 am"},
		// ast
		{"ast", "d'ochobre 11, 11:00 de la mañana", "october 11 11:00 am"},
		{"ast", "vienres 19 payares 1 tarde", "friday 19 november 1 pm"},
		// az-Cyrl
		{"az-Cyrl", "7 феврал 1788 05:30 пм", "7 february 1788 05:30 pm"},
		{"az-Cyrl", "чәршәнбә ахшамы ијл 14", "tuesday july 14"},
		// az-Latn
		{"az-Latn", "yanvar 13 şənbə", "january 13 saturday"},
		{"az-Latn", "b noy 12", "sunday november 12"},
		// az
		{"az", "17 iyn 2000 cümə axşamı", "17 june 2000 thursday"},
		{"az", "22 sentyabr 2003 bazar ertəsi", "22 september 2003 monday"},
		// bas
		{"bas", "1906 6 hìlòndɛ̀ ŋgwà njaŋgumba", "1906 6 june monday"},
		{"bas", "ŋgwà kɔɔ, 11 màtùmb 5 i ɓugajɔp", "friday 11 march 5 pm"},
		// be
		{"be", "13 лютага 1913", "13 february 1913"},
		{"be", "жнівень 12, чацвер", "august 12 thursday"},
		// bem
		{"bem", "palichimo 12 machi 2015 11:00 uluchelo", "monday 12 march 2015 11:00 am"},
		{"bem", "5 epreo 2000 pa mulungu", "5 april 2000 sunday"},
		// bez
		{"bez", "1 pa mwedzi gwa hutala 1889 10:00 pamilau", "1 january 1889 10:00 am"},
		{"bez", "31 pa mwedzi gwa kumi na mbili hit", "31 december thursday"},
		// bm
		{"bm", "12 ɔkutɔburu 2001 araba", "12 october 2001 wednesday"},
		{"bm", "alamisa 15 uti 1998", "thursday 15 august 1998"},
		// bo
		{"bo", "ཟླ་བ་བཅུ་གཅིག་པ་ 18", "november 18"},
		{"bo", "གཟའ་ཕུར་བུ་ 12 ཟླ་བ་བཅུ་པ་ 1879 10:15 ཕྱི་དྲོ་", "thursday 12 october 1879 10:15 pm"},
		// br
		{"br", "merc'her c'hwevrer 12 07:32 gm", "wednesday february 12 07:32 pm"},
		{"br", "10 gwengolo 2002 sadorn", "10 september 2002 saturday"},
		// brx
		{"brx", "6 अखथबर 2019 10:00 बेलासे", "6 october 2019 10:00 pm"},
		{"brx", "बिसथि 8 फेब्रुवारी", "thursday 8 february"},
		// bs-Cyrl
		{"bs-Cyrl", "2 септембар 2000, четвртак", "2 september 2000 thursday"},
		{"bs-Cyrl", "1 јули 1987 9:25 поподне", "1 july 1987 9:25 pm"},
		// bs-Latn
		{"bs-Latn", "23 septembar 1879, petak", "23 september 1879 friday"},
		{"bs-Latn", "subota 1 avg 2009 02:27 popodne", "saturday 1 august 2009 02:27 pm"},
		// bs
		{"bs", "10 maj 2020 utorak", "10 may 2020 tuesday"},
		{"bs", "ponedjeljak, 1989 2 januar", "monday 1989 2 january"},
		// ca
		{"ca", "14 d'abril 1980 diumenge", "14 april 1980 sunday"},
		{"ca", "3 de novembre 2004 dj", "3 november 2004 thursday"},
		// ce
		{"ce", "6 январь 1987 пӏераскан де", "6 january 1987 friday"},
		{"ce", "оршотан де 3 июль 1890", "monday 3 july 1890"},
		// cgg
		{"cgg", "20 okwakataana 2027 orwamukaaga", "20 may 2027 saturday"},
		{"cgg", "okwaikumi na ibiri 12 oks", "december 12 wednesday"},
		// chr
		{"chr", "ᎤᎾᏙᏓᏉᏅᎯ 16 ᏕᎭᎷᏱ 1562 11:16 ᏒᎯᏱᎢᏗᏢ", "monday 16 june 1562 11:16 pm"},
		{"chr", "13 ᎠᏂᏍᎬᏘ ᎤᎾᏙᏓᏈᏕᎾ 8:00 ᏌᎾᎴ", "13 may saturday 8:00 am"},
		// cy
		{"cy", "dydd sadwrn 27 chwefror 1990 9 yb", "saturday 27 february 1990 9 am"},
		{"cy", "19 gorff 2000 dydd gwener", "19 july 2000 friday"},
		// dav
		{"dav", "mori ghwa kawi 24 kuramuka kana", "february 24 thursday"},
		{"dav", "11 ike 4 luma lwa p", "11 september 4 pm"},
		// dje
		{"dje", "2 žuweŋ 2030 alz 11 zaarikay b", "2 june 2030 friday 11 pm"},
		{"dje", "sektanbur 12 alarba", "september 12 wednesday"},
		// dsb
		{"dsb", "njeźela julija 15 2 wótpołdnja", "sunday july 15 2 pm"},
		{"dsb", "awgusta 10 sob", "august 10 saturday"},
		// dua
		{"dua", "madiɓɛ́díɓɛ́ 15 ɗónɛsú 7 idiɓa", "july 15 friday 7 am"},
		{"dua", "éti 12 tiníní", "sunday 12 november"},
		// dyo
		{"dyo", "mee 1 2000 talata", "may 1 2000 tuesday"},
		{"dyo", "arjuma de 10", "friday december 10"},
		// dz
		{"dz", "ཟླ་བཅུ་གཅིག་པ་ 10 གཟའ་ཉི་མ་", "november 10 saturday"},
		{"dz", "མིར་ 2 སྤྱི་ཟླ་དྲུག་པ 2009 2 ཕྱི་ཆ་", "monday 2 june 2009 2 pm"},
		// ebu
		{"ebu", "mweri wa gatantatũ 11 maa 08:05 ut", "june 11 friday 08:05 pm"},
		{"ebu", "2 igi 1998 njumamothii", "2 december 1998 saturday"},
		// ee
		{"ee", "5 afɔfĩe 2009 05:05 ɣetrɔ kɔsiɖa", "5 april 2009 05:05 pm sunday"},
		{"ee", "yawoɖa 1890 deasiamime 23", "thursday 1890 august 23"},
		// el
		{"el", "απρίλιος 13 09:09 μμ", "april 13 09:09 pm"},
		{"el", "1 ιούνιος 2002 07:17 πμ", "1 june 2002 07:17 am"},
		// eo
		{"eo", "12 aŭgusto 1887 06:06 atm", "12 august 1887 06:06 am"},
		{"eo", "vendredo 10 sep 1957", "friday 10 september 1957"},
		// et
		{"et", "3 juuni 2001 neljapäev 07:09 pm", "3 june 2001 thursday 07:09 pm"},
		{"et", "7 veebr 2004", "7 february 2004"},
		// eu
		{"eu", "1 urtarrila 1990 asteazkena", "1 january 1990 wednesday"},
		{"eu", "ig 30 martxoa 1905", "sunday 30 march 1905"},
		// ewo
		{"ewo", "ngɔn lála 13 08:07 ngəgógəle", "march 13 08:07 pm"},
		{"ewo", "séradé ngad 12 1915 2:00 ngəgógəle", "saturday november 12 1915 2:00 pm"},
		// ff
		{"ff", "1 colte 1976 hoore-biir 04:15 subaka", "1 february 1976 saturday 04:15 am"},
		{"ff", "naasaande 3 yar 02:00 kikiiɗe", "thursday 3 october 02:00 pm"},
		// fil
		{"fil", "2 setyembre 1880 biyernes", "2 september 1880 friday"},
		{"fil", "15 ago 1909 lun", "15 august 1909 monday"},
		// fo
		{"fo", "mánadagur 30 januar 1976", "monday 30 january 1976"},
		{"fo", "2 apríl 1890 fríggjadagur", "2 april 1890 friday"},
		// fur
		{"fur", "12 avost 1990 domenie", "12 august 1990 sunday"},
		{"fur", "miercus 5 fev 1990 10:10 p", "wednesday 5 february 1990 10:10 pm"},
		// fy
		{"fy", "febrewaris 2 1987 freed", "february 2 1987 friday"},
		{"fy", "to 20 maaie 2010", "thursday 20 may 2010"},
		// ga
		{"ga", "1 bealtaine 2019 dé céadaoin", "1 may 2019 wednesday"},
		{"ga", "deireadh fómhair 12 aoine 10:09 pm", "october 12 friday 10:09 pm"},
		// gd
		{"gd", "2 am faoilleach 1890 diardaoin 02:13 m", "2 january 1890 thursday 02:13 am"},
		{"gd", "did an t-ògmhios 15 1876 08:15 f", "sunday june 15 1876 08:15 pm"},
		// gl
		{"gl", "1 xullo 2009 sáb", "1 july 2009 saturday"},
		{"gl", "martes 15 setembro 1980", "tuesday 15 september 1980"},
		// gsw
		{"gsw", "5 auguscht 1856 10:08 am namittag", "5 august 1856 10:08 pm"},
		{"gsw", "ziischtig 13 dezämber 03:12 vormittag", "tuesday 13 december 03:12 am"},
		// gu
		{"gu", "10 સપ્ટેમ્બર 2005 ગુરુવાર", "10 september 2005 thursday"},
		{"gu", "સોમવાર 1 જુલાઈ 1980", "monday 1 july 1980"},
		// guz
		{"guz", "apiriri 2 1789 chumatano", "april 2 1789 wednesday"},
		{"guz", "esabato 11 cul 2000 10:19 ma", "saturday 11 july 2000 10:19 am"},
		// gv
		{"gv", "3 toshiaght-arree 2023 jeh", "3 february 2023 friday"},
		{"gv", "1 m-souree 1999 jedoonee", "1 june 1999 sunday"},
		// ha
		{"ha", "18 yuni 1920 laraba", "18 june 1920 wednesday"},
		{"ha", "2 afi 1908 lit", "2 april 1908 monday"},
		// haw
		{"haw", "1 'apelila 1968 p6", "1 april 1968 saturday"},
		{"haw", "po'alima 29 'ok 1899", "friday 29 october 1899"},
		// hr
		{"hr", "2 ožujak 1980 pet", "2 march 1980 friday"},
		{"hr", "nedjelja 3 lis 1879", "sunday 3 october 1879"},
		// hsb
		{"hsb", "5 měrc 1789 póndźela 11:13 popołdnju", "5 march 1789 monday 11:13 pm"},
		{"hsb", "štwórtk 2000 awg 14", "thursday 2000 august 14"},
		// hy
		{"hy", "2 դեկտեմբերի 2006 շբթ 02:00 կա", "2 december 2006 saturday 02:00 am"},
		{"hy", "չորեքշաբթի մյս 17, 2009", "wednesday may 17 2009"},
		// ig
		{"ig", "1 ọgọọst 2001 wenezdee", "1 august 2001 wednesday"},
		{"ig", "mbọsị ụka 23 epr 1980", "sunday 23 april 1980"},
		// ii
		{"ii", "ꆏꊂꇖ 12 ꌕꆪ 1980", "thursday 12 march 1980"},
		{"ii", "ꉆꆪ 1 02:05 ꁯꋒ", "august 1 02:05 pm"},
		// is
		{"is", "þriðjudagur 15 júlí 2001", "tuesday 15 july 2001"},
		{"is", "fös 10 desember 08:17 fh", "friday 10 december 08:17 am"},
		// jgo
		{"jgo", "pɛsaŋ pɛ́nɛ́pfúꞌú 15 10:16 ŋka mbɔ́t nji", "september 15 10:16 pm"},
		{"jgo", "ápta mɔ́ndi 10 nduŋmbi saŋ 2009", "tuesday 10 january 2009"},
		// jmc
		{"jmc", "2 aprilyi 2015 jumapilyi 03:10 kyiukonyi", "2 april 2015 sunday 03:10 pm"},
		{"jmc", "alh 11 julyai 1987", "thursday 11 july 1987"},
		// kab
		{"kab", "3 meɣres 1999 kuẓass 11:16 n tmeddit", "3 march 1999 wednesday 11:16 pm"},
		{"kab", "1 yennayer 2004 sḍis", "1 january 2004 friday"},
		// kam
		{"kam", "mwai wa katatũ 12 wa katano 09:18 ĩyawĩoo", "march 12 friday 09:18 pm"},
		{"kam", "1 mwai wa ĩkumi na ilĩ 1789 wth", "1 december 1789 saturday"},
		// kde
		{"kde", "mwedi wa nnyano na umo 12 1907 liduva litandi", "june 12 1907 saturday"},
		{"kde", "2 mei 11:10 chilo ll6", "2 may 11:10 pm thursday"},
		// kea
		{"kea", "sigunda-fera 12 julhu 1902", "monday 12 july 1902"},
		{"kea", "2 diz 2005 kua", "2 december 2005 wednesday"},
		// khq
		{"khq", "1 žanwiye 2019 ati 01:09 adduha", "1 january 2019 monday 01:09 am"},
		{"khq", "alhamiisa 12 noowanbur 1908", "thursday 12 november 1908"},
		// ki
		{"ki", "1 mwere wa gatano 1980 09:12 hwaĩ-inĩ", "1 may 1980 09:12 pm"},
		{"ki", "njumatana 2 wmw 2000 01:12 kiroko", "wednesday 2 november 2000 01:12 am"},
		// kk
		{"kk", "3 маусым 1956 дс", "3 june 1956 monday"},
		{"kk", "жексенбі 12 қыркүйек 1890", "sunday 12 september 1890"},
		// kl
		{"kl", "2 martsi 2001 ataasinngorneq", "2 march 2001 monday"},
		{"kl", "pin 1 oktoberi 1901", "wednesday 1 october 1901"},
		// kln
		{"kln", "3 ng'atyaato koang'wan 10:09 kooskoliny", "3 february thursday 10:09 pm"},
		{"kln", "kipsuunde nebo aeng' 14 2009 kos", "december 14 2009 wednesday"},
		// kok
		{"kok", "1 नोव्हेंबर 2000 आदित्यवार 01:19 मनं", "1 november 2000 sunday 01:19 pm"},
		{"kok", "मंगळार 2 फेब्रुवारी 2003", "tuesday 2 february 2003"},
		// ksb
		{"ksb", "jumaamosi 1 ago 09:19 makeo", "saturday 1 august 09:19 am"},
		{"ksb", "3 febluali 1980 jmn", "3 february 1980 tuesday"},
		// ksf
		{"ksf", "ŋwíí a ntɔ́ntɔ 3 1990 09:15 cɛɛ́nko", "january 3 1990 09:15 pm"},
		{"ksf", "2 ŋ3 1789 jǝǝdí", "2 march 1789 thursday"},
		// ksh
		{"ksh", "mohndaach 12 fäbrowa 2001 12:18 nm", "monday 12 february 2001 12:18 pm"},
		{"ksh", "5 oujoß 12:17 uhr vörmiddaachs", "5 august 12:17 am"},
		// kw
		{"kw", "14 mis metheven 1980 dy yow", "14 june 1980 thursday"},
		{"kw", "mis kevardhu 2019 1 sad", "december 2019 1 saturday"},
		// ky
		{"ky", "22 февраль 2025 01:12 түштөн кийинки", "22 february 2025 01:12 pm"},
		{"ky", "шаршемби 11 авг 1908", "wednesday 11 august 1908"},
		// lag
		{"lag", "17 kʉvɨɨrɨ 2018 ijumáa", "17 august 2018 friday"},
		{"lag", "táatu 16 kwiinyi 1978", "monday 16 october 1978"},
		// lb
		{"lb", "2 mäerz 2034 don 02:19 moies", "2 march 2034 thursday 02:19 am"},
		{"lb", "samschdeg 15 abrëll", "saturday 15 april"},
		// lg
		{"lg", "sebuttemba 17 1980 lw6", "september 17 1980 saturday"},
		{"lg", "2 okitobba 2010 lwakusatu", "2 october 2010 wednesday"},
		// lkt
		{"lkt", "18 čhaŋwápetȟo wí 2013 owáŋgyužažapi", "18 may 2013 saturday"},
		{"lkt", "1 tȟahékapšuŋ wí 1978 aŋpétuzaptaŋ", "1 december 1978 friday"},
		// ln
		{"ln", "23 yan 2001 mokɔlɔ mwa mísáto", "23 january 2001 wednesday"},
		{"ln", "mtn 17 sánzá ya zómi na míbalé 09:17 ntɔ́ngɔ́", "friday 17 december 09:17 am"},
		// lo
		{"lo", "18 ພຶດສະພາ 1908 ວັນອາທິດ", "18 may 1908 sunday"},
		{"lo", "8 ກໍລະກົດ 2003 03:03 ຫຼັງທ່ຽງ", "8 july 2003 03:03 pm"},
		// lt
		{"lt", "15 gegužės 1970 trečiadienis", "15 may 1970 wednesday"},
		{"lt", "an 2 rugsėjo 09:18 priešpiet", "tuesday 2 september 09:18 am"},
		// lu
		{"lu", "2 ciongo 2016 njw 02:16 dilolo", "2 january 2016 thursday 02:16 pm"},
		{"lu", "16 lùshìkà 2009", "16 august 2009"},
		// luo
		{"luo", "15 dwe mar adek 1908 tan", "15 march 1908 thursday"},
		{"luo", "jumapil 3 dao 2008 01:12 ot", "sunday 3 july 2008 01:12 pm"},
		// luy
		{"luy", "23 juni 1970 murwa wa kanne", "23 june 1970 thursday"},
		{"luy", "jumatano, 5 aprili 1998", "wednesday 5 april 1998"},
		// lv
		{"lv", "14 jūnijs 2010 10:10 priekšpusdienā", "14 june 2010 10:10 am"},
		{"lv", "24 okt 2000 piektdiena 11:11 pēcpusd", "24 october 2000 friday 11:11 pm"},
		// mas
		{"mas", "2 olodoyíóríê inkókúâ 1954 08:16 ɛnkakɛnyá", "2 april 1954 08:16 am"},
		{"mas", "15 ɔɛn 2032 alaámisi 02:13 ɛndámâ", "15 march 2032 thursday 02:13 pm"},
		// mer
		{"mer", "1 mĩĩ 1924 wetano 10:05 ũg", "1 may 1924 friday 10:05 pm"},
		{"mer", "6 njuraĩ 1895 muramuko", "6 july 1895 monday"},
		// mfe
		{"mfe", "27 zilye 1988 merkredi", "27 july 1988 wednesday"},
		{"mfe", "lindi 3 oktob 1897", "monday 3 october 1897"},
		// mg
		{"mg", "1 martsa 1789 alakamisy", "1 march 1789 thursday"},
		{"mg", "5 aogositra 1911 zoma", "5 august 1911 friday"},
		// mgh
		{"mgh", "sabato mweri wo unethanu 15 07:18 wichishu", "sunday may 15 07:18 am"},
		{"mgh", "2 tis 1789 jumamosi 08:17 mchochil'l", "2 september 1789 saturday 08:17 pm"},
		// mgo
		{"mgo", "5 iməg mbegtug aneg 5", "5 january thursday"},
		{"mgo", "aneg 2 iməg zò 17 1908", "monday november 17 1908"},
		// mk
		{"mk", "4 септември 2009 09:18 претпл", "4 september 2009 09:18 am"},
		{"mk", "вторник 10 август 1777 01:12 попл", "tuesday 10 august 1777 01:12 pm"},
		// mn
		{"mn", "дөрөвдүгээр сар 15 баасан 10:10 үө", "april 15 friday 10:10 am"},
		{"mn", "12 9-р сар 2019 пүрэв", "12 september 2019 thursday"},
		// mr
		{"mr", "16 फेब्रुवारी 1908 गुरु 02:03 मउ", "16 february 1908 thursday 02:03 pm"},
		{"mr", "शनिवार 15 सप्टें 1888", "saturday 15 september 1888"},
		// ms
		{"ms", "4 mei 1768 jumaat 09:09 ptg", "4 may 1768 friday 09:09 pm"},
		{"ms", "isnin 14 disember 2001 11:09 pg", "monday 14 december 2001 11:09 am"},
		// mt
		{"mt", "3 frar 1998 il-ħadd", "3 february 1998 sunday"},
		{"mt", "16 awwissu 2019 erb", "16 august 2019 wednesday"},
		// mua
		{"mua", "1 cokcwaklaŋne 2014 comzyiiɗii", "1 february 2014 tuesday"},
		{"mua", "fĩi yuru 17 1908 cze 10:08 lilli", "december 17 1908 saturday 10:08 pm"},
		// naq
		{"naq", "20 ǂkhoesaob 1934 wunstaxtsees", "20 july 1934 wednesday"},
		{"naq", "do 12 gamaǀaeb 1999 05:12 ǃuias", "thursday 12 june 1999 05:12 pm"},
		// nb
		{"nb", "2 mars 1998 torsdag", "2 march 1998 thursday"},
		{"nb", "ons 15 des 2001", "wednesday 15 december 2001"},
		// nd
		{"nd", "19 nkwenkwezi 1967 mgqibelo", "19 may 1967 saturday"},
		{"nd", "sit 1 zibandlela 2011", "wednesday 1 january 2011"},
		// ne
		{"ne", "1 फेब्रुअरी 2003 बिहिबार 09:18 अपराह्न्", "1 february 2003 thursday 09:18 pm"},
		{"ne", "आइत् 4 अक्टोबर 1957", "sunday 4 october 1957"},
		// nl
		{"nl", "4 augustus 1678 zaterdag", "4 august 1678 saturday"},
		{"nl", "vr 27 juni 1997", "friday 27 june 1997"},
		// nmg
		{"nmg", "5 ngwɛn ńna 1897 sɔ́ndɔ mafú málal", "5 april 1897 wednesday"},
		{"nmg", "mɔ́ndɔ 1 ng11 1678 04:15 kugú", "monday 1 november 1678 04:15 pm"},
		// nn
		{"nn", "tysdag 2 september 1897 01:12 fm", "tuesday 2 september 1897 01:12 am"},
		{"nn", "2 mai 2000 søndag 09:18 ettermiddag", "2 may 2000 sunday 09:18 pm"},
		// nnh
		{"nnh", "1 saŋ tsɛ̀ɛ cÿó màga lyɛ̌' 08:18 ncwònzém", "1 may saturday 08:18 pm"},
		{"nnh", "3 saŋ lepyè shúm 1789 mvfò lyɛ̌'", "3 march 1789 monday"},
		// nus
		{"nus", "7 kornyoot 2006 diɔ̱k lätni 01:12 tŋ", "7 june 2006 wednesday 01:12 pm"},
		{"nus", "bäkɛl, 12 tio̱p in di̱i̱t 2008", "saturday 12 december 2008"},
		// nyn
		{"nyn", "5 okwakashatu 1980 okt", "5 march 1980 friday"},
		{"nyn", "2 kms 2087 sande", "2 july 2087 sunday"},
		// om
		{"om", "15 bitooteessa 1997 02:23 wb", "15 march 1997 02:23 pm"},
		{"om", "jimaata 13 gur 01:12 wd", "friday 13 february 01:12 am"},
		// os
		{"os", "хуыцаубон 1998 апрелы 12", "sunday 1998 april 12"},
		{"os", "1 ноя 1990 ӕртыццӕг", "1 november 1990 wednesday"},
		// pa-Guru
		{"pa-Guru", "ਸ਼ਨਿੱਚਰਵਾਰ 4 ਫ਼ਰਵਰੀ 1989 01:12 ਬਾਦ", "saturday 4 february 1989 01:12 pm"},
		{"pa-Guru", "2 ਅਕਤੂਬਰ 2015 ਸੋਮਵਾਰ", "2 october 2015 monday"},
		// pa
		{"pa", "2 ਅਗਸਤ 1682 ਸ਼ਨਿੱਚਰ", "2 august 1682 saturday"},
		{"pa", "12 ਅਕਤੂ 11:08 ਪੂਦੁ", "12 october 11:08 am"},
		// qu
		{"qu", "5 pauqar waray 1878 miércoles", "5 march 1878 wednesday"},
		{"qu", "6 int 2009 domingo", "6 june 2009 sunday"},
		// rm
		{"rm", "1 schaner 1890 venderdi", "1 january 1890 friday"},
		{"rm", "me 6 avust 2009", "wednesday 6 august 2009"},
		// rn
		{"rn", "11 ntwarante 2008 12:34 zmw", "11 march 2008 12:34 pm"},
		{"rn", "7 nze 1999 ku wa kabiri", "7 september 1999 tuesday"},
		// rof
		{"rof", "13 mweri wa tisa ijtn 12:56 kingoto", "13 september wednesday 12:56 pm"},
		{"rof", "ijumanne 2 mweri wa saba 1890", "tuesday 2 july 1890"},
		// rw
		{"rw", "16 kamena 2001 kuwa gatanu", "16 june 2001 friday"},
		{"rw", "3 ukuboza 2013 gnd", "3 december 2013 saturday"},
		// rwk
		{"rwk", "3 aprilyi 2009 ijumaa", "3 april 2009 friday"},
		{"rwk", "jumamosi 2 januari 02:13 utuko", "saturday 2 january 02:13 am"},
		// sah
		{"sah", "16 тохсунньу 2003 сэрэдэ 09:59 эк", "16 january 2003 wednesday 09:59 pm"},
		{"sah", "баскыһыанньа 14 балаҕан ыйа 1998", "sunday 14 september 1998"},
		// saq
		{"saq", "1 lapa le okuni 1980 kun 10:45 tesiran", "1 march 1980 monday 10:45 am"},
		{"saq", "mderot ee inet 12 lapa le ong'wan 1824", "wednesday 12 april 1824"},
		// sbp
		{"sbp", "1 mupalangulwa mulungu 08:15 lwamilawu", "1 january sunday 08:15 am"},
		{"sbp", "jtn 17 mokhu 2001", "wednesday 17 october 2001"},
		// se
		{"se", "láv 22 cuoŋománnu 10:08 iđitbeaivi", "saturday 22 april 10:08 am"},
		{"se", "duorasdat 11 borgemánnu 1978 12:09 eb", "thursday 11 august 1978 12:09 pm"},
		// seh
		{"seh", "12 fevreiro 2005 sha", "12 february 2005 friday"},
		{"seh", "chiposi 2 decembro 1987", "monday 2 december 1987"},
		// ses
		{"ses", "18 žuyye 2009 atalaata 03:12 aluula", "18 july 2009 tuesday 03:12 pm"},
		{"ses", "asibti 2 awi 1987", "saturday 2 april 1987"},
		// sg
		{"sg", "5 ngubùe 1890 bïkua-ûse 12:08 lk", "5 april 1890 monday 12:08 pm"},
		{"sg", "bk3 23 föndo 2001", "tuesday 23 june 2001"},
		// shi-Latn
		{"shi-Latn", "6 bṛayṛ 2014 akṛas 07:06 tifawt", "6 february 2014 wednesday 07:06 am"},
		{"shi-Latn", "asamas 15 ɣuct 2045", "sunday 15 august 2045"},
		// sk
		{"sk", "15 marec 1987 utorok", "15 march 1987 tuesday"},
		{"sk", "streda 17 mája 2003", "wednesday 17 may 2003"},
		// sl
		{"sl", "12 junij 2003 petek 10:09 pop", "12 june 2003 friday 10:09 pm"},
		{"sl", "ponedeljek 15 okt 1997 09:07 dopoldne", "monday 15 october 1997 09:07 am"},
		// smn
		{"smn", "1 njuhčâmáánu 2008 majebaargâ 08:08 ip", "1 march 2008 tuesday 08:08 am"},
		{"smn", "láv 23 roovvâd 1897", "saturday 23 october 1897"},
		// sn
		{"sn", "11 chikumi 1998 chipiri", "11 june 1998 tuesday"},
		{"sn", "china 2 mbudzi 1890", "thursday 2 november 1890"},
		// so
		{"so", "sab 5 bisha saddexaad 1765 11:08 gn", "saturday 5 march 1765 11:08 pm"},
		{"so", "16 lit 2008 axd", "16 december 2008 sunday"},
		// sq
		{"sq", "2 qershor 1997 e mërkurë 10:08 pasdite", "2 june 1997 wednesday 10:08 pm"},
		{"sq", "pre 15 gusht 1885 04:54 e paradites", "friday 15 august 1885 04:54 am"},
		// sr-Cyrl
		{"sr-Cyrl", "16 април 2016 суб 03:46 по подне", "16 april 2016 saturday 03:46 pm"},
		{"sr-Cyrl", "уторак 3 новембар 1999", "tuesday 3 november 1999"},
		// sr-Latn
		{"sr-Latn", "4 septembar 2000 četvrtak", "4 september 2000 thursday"},
		{"sr-Latn", "uto 18 maj 2004 11:15 pre podne", "tuesday 18 may 2004 11:15 am"},
		// sr
		{"sr", "3 децембар 2005 уто 10:15 по подне", "3 december 2005 tuesday 10:15 pm"},
		{"sr", "петак 12 август 2001", "friday 12 august 2001"},
		// sv
		{"sv", "4 augusti 2007 lördag 02:44 fm", "4 august 2007 saturday 02:44 am"},
		{"sv", "onsdag 16 mars 08:15 eftermiddag", "wednesday 16 march 08:15 pm"},
		// sw
		{"sw", "5 mei 1994 jumapili 10:17 asubuhi", "5 may 1994 sunday 10:17 am"},
		{"sw", "jumanne 2 desemba 2003", "tuesday 2 december 2003"},
		// ta
		{"ta", "6 ஏப்ரல் 1997 செவ்வாய் 02:09 முற்பகல்", "6 april 1997 tuesday 02:09 am"},
		{"ta", "ஞாயி 1 ஜூன் 1998", "sunday 1 june 1998"},
		// te
		{"te", "సోమవారం 3 నవంబర 1887", "monday 3 november 1887"},
		{"te", "5 మార్చి 2001 శుక్రవారం", "5 march 2001 friday"},
		// teo
		{"teo", "2 omodok'king'ol 1996 nakaare", "2 june 1996 tuesday"},
		{"teo", "nakasabiti 4 jol 2001 01:12 ebongi", "saturday 4 july 2001 01:12 pm"},
		// to
		{"to", "5 fēpueli 2007 mōn 02:17 efiafi", "5 february 2007 monday 02:17 pm"},
		{"to", "falaite 14 'okatopa 2015 09:48 hh", "friday 14 october 2015 09:48 am"},
		// twq
		{"twq", "17 feewiriye 2023 11:12 zaarikay b", "17 february 2023 11:12 pm"},
		{"twq", "alzuma 11 sektanbur 2019", "friday 11 september 2019"},
		// tzm
		{"tzm", "2 yulyuz 2002 akwas 01:16 ḍeffir aza", "2 july 2002 thursday 01:16 pm"},
		{"tzm", "asa 13 nwanbir 2005", "sunday 13 november 2005"},
		// uz-Cyrl
		{"uz-Cyrl", "пайшанба 24 ноябр 1957 01:18 то", "thursday 24 november 1957 01:18 am"},
		{"uz-Cyrl", "4 авг 1887 чоршанба", "4 august 1887 wednesday"},
		// uz-Latn
		{"uz-Latn", "3 iyul 1997 payshanba 08:17 tk", "3 july 1997 thursday 08:17 pm"},
		{"uz-Latn", "shan 15 sentabr 2008", "saturday 15 september 2008"},
		// uz
		{"uz", "1 fevral 1776 dushanba 09:17 to", "1 february 1776 monday 09:17 am"},
		{"uz", "juma 18 aprel 2027", "friday 18 april 2027"},
		// vun
		{"vun", "2 aprilyi 1956 jumatatuu", "2 april 1956 monday"},
		{"vun", "jumamosi 12 oktoba 02:16 kyiukonyi", "saturday 12 october 02:16 pm"},
		// wae
		{"wae", "zištag 16 abrille 2002", "tuesday 16 april 2002"},
		{"wae", "27 öigšte 1669 fritag", "27 august 1669 friday"},
		// xog
		{"xog", "21 marisi 2001 owokubili", "21 march 2001 tuesday"},
		{"xog", "kuta 30 okitobba 1955 02:17 eigulo", "friday 30 october 1955 02:17 pm"},
		// yav
		{"yav", "12 imɛŋ i puɔs 1998 metúkpíápɛ", "12 september 1998 wednesday"},
		{"yav", "5 o10 2001 séselé 12:07 kiɛmɛ́ɛm", "5 october 2001 saturday 12:07 am"},
		// yo
		{"yo", "5 èrèlè 2005 ọjọ́rú 10:07 àárọ̀", "5 february 2005 wednesday 10:07 am"},
		{"yo", "ọjọ́ àbámẹ́ta 2 oṣù ẹ̀bibi 1896", "saturday 2 may 1896"},
		// zu
		{"zu", "3 mashi 2007 ulwesibili 10:08", "3 march 2007 tuesday 10:08"},
		{"zu", "son 23 umasingana 1996", "sunday 23 january 1996"},
	}

	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("%s, \"%s\"", test.Locale, test.String)

		// Load locale
		ld, err := GetLocale(test.Locale)
		assert.Nil(t, err, message)

		// Translate string
		cfg := &setting.DefaultConfig
		translation := Translate(*ld, test.String, false, cfg)
		passed := assert.Equal(t, test.Expected, translation, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestFreshnessTranslation(t *testing.T) {
	type testScenario struct {
		Locale   string
		String   string
		Expected string
	}

	tests := []testScenario{
		// English
		{"en", "yesterday", "1 day ago"},
		{"en", "today", "0 day ago"},
		{"en", "day before yesterday", "2 day ago"},
		{"en", "last month", "1 month ago"},
		{"en", "less than a minute ago", "45 second ago"},
		// German
		{"de", "vorgestern", "2 day ago"},
		{"de", "heute", "0 day ago"},
		{"de", "vor 3 Stunden", "3 hour ago"},
		{"de", "vor 2 Monaten", "2 month ago"},
		{"de", "vor 2 Monaten, 2 Wochen", "2 month ago 2 week"},
		// French
		{"fr", "avant-hier", "2 day ago"},
		{"fr", "hier", "1 day ago"},
		{"fr", "aujourd'hui", "0 day ago"},
		{"fr", "après dix ans", "in 10 year"},
		// Spanish
		{"es", "anteayer", "2 day ago"},
		{"es", "ayer", "1 day ago"},
		{"es", "ayer a las", "1 day ago"},
		{"es", "hoy", "0 day ago"},
		{"es", "hace un horas", "1 hour ago"},
		{"es", "2 semanas", "2 week"},
		{"es", "2 año", "2 year"},
		// Italian
		{"it", "altro ieri", "2 day ago"},
		{"it", "ieri", "1 day ago"},
		{"it", "oggi", "0 day ago"},
		{"it", "2 settimana fa", "2 week ago"},
		{"it", "2 anno fa", "2 year ago"},
		// Portuguese
		{"pt", "anteontem", "2 day ago"},
		{"pt", "ontem", "1 day ago"},
		{"pt", "hoje", "0 day ago"},
		{"pt", "56 minutos", "56 minute"},
		{"pt", "12 dias", "12 day"},
		{"pt", "há 14 min.", "14 minute ago."},
		{"pt", "1 segundo atrás", "1 second ago"},
		// Russian
		{"ru", "9 месяцев", "9 month"},
		{"ru", "8 недель", "8 week"},
		{"ru", "7 лет", "7 year"},
		{"ru", "позавчера", "2 day ago"},
		{"ru", "сейчас", "0 second ago"},
		{"ru", "спустя 2 дня", "in 2 day"},
		{"ru", "вчера", "1 day ago"},
		{"ru", "сегодня", "0 day ago"},
		{"ru", "завтра", "in 1 day"},
		{"ru", "послезавтра", "in 2 day"},
		{"ru", "во вторник", "tuesday"},
		{"ru", "в воскресенье", "sunday"},
		{"ru", "несколько секунд", "44 second"},
		{"ru", "через пару секунд", "in 2 second"},
		{"ru", "одну минуту назад", "1 minute ago"},
		{"ru", "через полчаса", "in 30 minute"},
		{"ru", "сорок минут назад", "40 minute ago"},
		{"ru", "в течение пары часов", "in 2 hour"},
		{"ru", "через четыре часа", "in 4 hour"},
		{"ru", "в течение суток", "in 1 day"},
		{"ru", "двое суток назад", "2 day ago"},
		{"ru", "неделю назад", "1 week ago"},
		{"ru", "две недели назад", "2 week ago"},
		{"ru", "три месяца назад", "3 month ago"},
		{"ru", "спустя полгода", "in 6 month"},
		{"ru", "через год", "in 1 year"},
		{"ru", "через полтора года", "in 18 month"},
		// Turkish
		{"tr", "dün", "1 day ago"},
		{"tr", "22 dakika", "22 minute"},
		{"tr", "12 hafta", "12 week"},
		{"tr", "13 yıl", "13 year"},
		// Czech
		{"cs", "40 sekunda", "40 second"},
		{"cs", "4 týden", "4 week"},
		{"cs", "14 roků", "14 year"},
		// Chinese
		{"zh", "昨天", "1 day ago"},
		{"zh", "前天", "2 day ago"},
		{"zh", "50 秒", "50 second"},
		{"zh", "7 周", "7 week"},
		{"zh", "12 年", "12 year"},
		{"zh", "半小时前", "30 minute ago"},
		// Danish
		{"da", "i går", "1 day ago"},
		{"da", "i dag", "0 day ago"},
		{"da", "sidste måned", "1 month ago"},
		{"da", "mindre end et minut siden", "45  seconds"},
		// Dutch
		{"nl", "17 uur geleden", "17 hour ago"},
		{"nl", "27 jaar geleden", "27 year ago"},
		{"nl", "45 minuten", "45 minute"},
		{"nl", "nu", "0 second ago"},
		{"nl", "eergisteren", "2 day ago"},
		{"nl", "volgende maand", "in 1 month"},
		// Romanian
		{"ro", "23 săptămâni în urmă", "23 week ago"},
		{"ro", "23 săptămâni", "23 week"},
		{"ro", "13 oră", "13 hour"},
		// Arabic
		{"ar", "يومين", "2 day"},
		{"ar", "أمس", "1 day ago"},
		{"ar", "4 عام", "4 year"},
		{"ar", "منذ 2 ساعات", "ago 2 hour"},
		{"ar", "منذ ساعتين", "ago 2 hour"},
		{"ar", "اليوم السابق", "1 day ago"},
		{"ar", "اليوم", "0 day ago"},
		// Polish
		{"pl", "2 godz.", "2 hour."},
		{"pl", "Wczoraj o 07:40", "1 day ago  07:40"},
		// Vietnamese
		{"vi", "2 tuần 3 ngày", "2 week 3 day"},
		{"vi", "21 giờ trước", "21 hour ago"},
		{"vi", "Hôm qua 08:16", "1 day ago 08:16"},
		{"vi", "Hôm nay 15:39", "0 day ago 15:39"},
		// French
		{"fr", "maintenant", "0 second ago"},
		{"fr", "demain", "in 1 day"},
		{"fr", "Il y a moins d'une minute", "1 minute ago"},
		{"fr", "Il y a moins de 30s", "30 second ago"},
		// Tagalog
		{"tl", "kahapon", "1 day ago"},
		{"tl", "ngayon", "0 second ago"},
		// Ukrainian
		{"uk", "позавчора", "2 day ago"},
		// Belarusian
		{"be", "9 месяцаў", "9 month"},
		{"be", "8 тыдняў", "8 week"},
		{"be", "1 тыдзень", "1 week"},
		{"be", "2 года", "2 year"},
		{"be", "3 гады", "3 year"},
		{"be", "11 секунд", "11 second"},
		{"be", "учора", "1 day ago"},
		{"be", "пазаўчора", "2 day ago"},
		{"be", "сёння", "0 day ago"},
		{"be", "некалькі хвілін", "2 minute"},
		// Indonesian
		{"id", "baru saja", "0 second ago"},
		{"id", "hari ini", "0 day ago"},
		{"id", "kemarin", "1 day ago"},
		{"id", "kemarin lusa", "2 day ago"},
		{"id", "sehari yang lalu", "1 day  ago"},
		{"id", "seminggu yang lalu", "1 week  ago"},
		{"id", "sebulan yang lalu", "1 month  ago"},
		{"id", "setahun yang lalu", "1 year  ago"},
		// Finnish
		{"fi", "1 vuosi sitten", "1 year ago"},
		{"fi", "2 vuotta sitten", "2 year ago"},
		{"fi", "3 v sitten", "3 year ago"},
		{"fi", "4 v. sitten", "4 year. ago"},
		{"fi", "5 vv. sitten", "5 year. ago"},
		{"fi", "1 kuukausi sitten", "1 month ago"},
		{"fi", "2 kuukautta sitten", "2 month ago"},
		{"fi", "3 kk sitten", "3 month ago"},
		{"fi", "1 viikko sitten", "1 week ago"},
		{"fi", "2 viikkoa sitten", "2 week ago"},
		{"fi", "3 vk sitten", "3 week ago"},
		{"fi", "4 vko sitten", "4 week ago"},
		{"fi", "1 päivä sitten", "1 day ago"},
		{"fi", "2 päivää sitten", "2 day ago"},
		{"fi", "8 pvää sitten", "8 day ago"},
		{"fi", "3 pv sitten", "3 day ago"},
		{"fi", "4 p. sitten", "4 day. ago"},
		{"fi", "5 pvä sitten", "5 day ago"},
		{"fi", "1 tunti sitten", "1 hour ago"},
		{"fi", "2 tuntia sitten", "2 hour ago"},
		{"fi", "3 t sitten", "3 hour ago"},
		{"fi", "1 minuutti sitten", "1 minute ago"},
		{"fi", "2 minuuttia sitten", "2 minute ago"},
		{"fi", "3 min sitten", "3 minute ago"},
		{"fi", "1 sekunti sitten", "1 second ago"},
		{"fi", "2 sekuntia sitten", "2 second ago"},
		{"fi", "1 sekuntti sitten", "1 second ago"},
		{"fi", "2 sekunttia sitten", "2 second ago"},
		{"fi", "3 s sitten", "3 second ago"},
		{"fi", "eilen", "1 day ago"},
		{"fi", "tänään", "0 day ago"},
		{"fi", "huomenna", "in 1 day"},
		{"fi", "nyt", "0 second ago"},
		{"fi", "ensi viikolla", "in 1 week"},
		{"fi", "viime viikolla", "1 week ago"},
		{"fi", "toissa vuonna", "2 year ago"},
		{"fi", "9 kuukautta sitten", "9 month ago"},
		{"fi", "3 viikon päästä", "in 3 week"},
		{"fi", "10 tunnin kuluttua", "in 10 hour"},
		// Japanese
		{"ja", "今年", "0 year ago"},
		{"ja", "去年", "1 year ago"},
		{"ja", "17年前", "17 year ago"},
		{"ja", "今月", "0 month ago"},
		{"ja", "先月", "1 month ago"},
		{"ja", "1ヶ月前", "1 month ago"},
		{"ja", "2ヶ月前", "2 month ago"},
		{"ja", "今週", "0 week ago"},
		{"ja", "先週", "1 week ago"},
		{"ja", "先々週", "2 week ago"},
		{"ja", "2週間前", "2 week ago"},
		{"ja", "3週間", "3 week"},
		{"ja", "今日", "0 day ago"},
		{"ja", "昨日", "1 day ago"},
		{"ja", "一昨日", "2 day ago"},
		{"ja", "3日前", "3 day ago"},
		{"ja", "1時間", "1 hour"},
		{"ja", "23時間前", "23 hour ago"},
		{"ja", "30分", "30 minute"},
		{"ja", "3分間", "3 minute"},
		{"ja", "60秒", "60 second"},
		{"ja", "3秒前", "3 second ago"},
		{"ja", "現在", "0 second ago"},
		// Hebrew
		{"he", "אתמול", "1 day ago"},
		{"he", "אתמול בשעה 3", "1 day ago  3"},
		{"he", "היום", "0 day ago"},
		{"he", "לפני יומיים", "2 day ago"},
		{"he", "לפני שבועיים", "2 week ago"},
		// Bulgarian
		{"bg", "вдругиден", "in 2 day"},
		{"bg", "утре", "in 1 day"},
		{"bg", "след 5 дни", "in 5 day"},
		{"bg", "вчера", "1 day ago"},
		{"bg", "преди 9 дни", "9 day ago"},
		{"bg", "преди 10 минути", "10 minute ago"},
		{"bg", "преди час", "1 hour ago"},
		{"bg", "преди 4 години", "4 year ago"},
		{"bg", "преди десетилетие", "10 year ago"},
		// Bangla
		{"bn", "গতকাল", "1 day ago"},
		{"bn", "আজ", "0 day ago"},
		{"bn", "গত মাস", "1 month ago"},
		{"bn", "আগামী সপ্তাহ", "in 1 week"},
		// Hindi
		{"hi", "१ सप्ताह", "1 week"},
		{"hi", "२४ मिनट पहले", "24 minute ago"},
		{"hi", "5 वर्ष", "5 year"},
		{"hi", "५३ सप्ताह बाद", "53 week in"},
		{"hi", "12 सेकंड पूर्व", "12 second ago"},
		// Swedish
		{"sv", "igår", "1 day ago"},
		{"sv", "idag", "0 day ago"},
		{"sv", "förrgår", "2 day ago"},
		{"sv", "förra månaden", "1 month ago"},
		{"sv", "nästa månad", "in 1 month"},
		// Georgian
		{"ka", "გუშინ", "1 day ago"},
		{"ka", "დღეს", "0 day ago"},
		{"ka", "ერთ თვე", "1 month"},
		{"ka", "დღეიდან ერთ კვირა", "in 1 week"},
		// af
		{"af", "3 week gelede", "3 week ago"},
		{"af", "volgende jaar 10:08 nm", "in 1 year 10:08 pm"},
		{"af", "oor 45 sekondes", "in 45 second"},
		// am
		{"am", "የሚቀጥለው ሳምንት", "in 1 week"},
		{"am", "በ10 ሳምንት ውስጥ", "in 10 week"},
		// as
		{"as", "কাইলৈ", "in 1 day"},
		{"as", "আজি", "0 day ago"},
		// asa
		{"asa", "ighuo", "1 day ago"},
		{"asa", "yavo 09:27 ichamthi", "in 1 day 09:27 pm"},
		// ast
		{"ast", "el mes viniente 02:17 tarde", "in 1 month 02:17 pm"},
		{"ast", "hai 22 selmana", "22 week ago"},
		{"ast", "en 5 minutos", "in 5 minute"},
		// az-Latn
		{"az-Latn", "keçən həftə", "1 week ago"},
		{"az-Latn", "10 ay ərzində", "in 10 month"},
		{"az-Latn", "22 saniyə öncə", "22 second ago"},
		// az
		{"az", "12 saat ərzində", "in 12 hour"},
		{"az", "15 həftə öncə", "15 week ago"},
		// bas
		{"bas", "lɛ̀n 12:08 i ɓugajɔp", "0 day ago 12:08 pm"},
		{"bas", "yààni", "1 day ago"},
		// bem
		{"bem", "lelo", "0 day ago"},
		{"bem", "17 umweshi", "17 month"},
		// bez
		{"bez", "hilawu 08:44 pamilau", "in 1 day 08:44 am"},
		{"bez", "neng'u ni", "0 day ago"},
		{"bez", "12 mlungu gumamfu", "12 week"},
		// bm
		{"bm", "sini 01:18 am", "in 1 day 01:18 am"},
		{"bm", "kunu", "1 day ago"},
		{"bm", "22 dɔgɔkun", "22 week"},
		// bo
		{"bo", "ཁས་ས་", "1 day ago"},
		{"bo", "སང་ཉིན་", "in 1 day"},
		// br
		{"br", "ar sizhun diaraok", "1 week ago"},
		{"br", "ar bloaz a zeu 02:19 gm", "in 1 year 02:19 pm"},
		{"br", "10 deiz", "10 day"},
		// brx
		{"brx", "गाबोन", "in 1 day"},
		{"brx", "मैया 11:58 फुं", "1 day ago 11:58 am"},
		{"brx", "17 मिनिथ", "17 minute"},
		// bs-Cyrl
		{"bs-Cyrl", "следећег месеца", "in 1 month"},
		{"bs-Cyrl", "прошле године 10:05 пре подне", "1 year ago 10:05 am"},
		{"bs-Cyrl", "пре 28 недеља", "28 week ago"},
		// bs-Latn
		{"bs-Latn", "sljedeće godine", "in 1 year"},
		{"bs-Latn", "prije 4 mjeseci", "4 month ago"},
		{"bs-Latn", "za 36 sati", "in 36 hour"},
		// bs
		{"bs", "prije 12 sekundi", "12 second ago"},
		{"bs", "za 5 godinu", "in 5 year"},
		{"bs", "ovaj sat", "0 hour ago"},
		// ca
		{"ca", "d'aquí a 22 hores", "in 22 hour"},
		{"ca", "fa 17 anys", "17 year ago"},
		{"ca", "el mes passat", "1 month ago"},
		{"ca", "la pròxima setmana", "in 1 week"},
		{"ca", "despús-ahir", "2 day ago"},
		{"ca", "en un dia", "in 1 day"},
		{"ca", "demà passat", "in 2 day"},
		// ce
		{"ce", "72 сахьт даьлча", "in 72 hour"},
		{"ce", "42 шо хьалха", "42 year ago"},
		{"ce", "рогӏерчу баттахь", "in 1 month"},
		// cgg
		{"cgg", "nyenkyakare", "in 1 day"},
		{"cgg", "nyomwabazyo", "1 day ago"},
		{"cgg", "5 omwaka", "5 year"},
		// chr
		{"chr", "ᎯᎠ ᎢᏯᏔᏬᏍᏔᏅ", "0 minute ago"},
		{"chr", "ᎾᎿ 8 ᎧᎸᎢ ᏥᎨᏒ", "8 month ago"},
		{"chr", "ᎾᎿ 22 ᎢᏯᏔᏬᏍᏔᏅ", "in 22 minute"},
		// cs
		{"cs", "za 3 rok", "in 3 year"},
		{"cs", "před 11 měsícem", "11 month ago"},
		{"cs", "tento měsíc", "0 month ago"},
		// cy
		{"cy", "wythnos ddiwethaf", "1 week ago"},
		{"cy", "25 o flynyddoedd yn ôl", "25 year ago"},
		{"cy", "ymhen 4 awr", "in 4 hour"},
		// da
		{"da", "for 15 måneder siden", "15 month ago"},
		{"da", "om 60 sekunder", "in 60 second"},
		{"da", "sidste måned", "1 month ago"},
		// dav
		{"dav", "iguo", "1 day ago"},
		{"dav", "kesho 02:12 luma lwa p", "in 1 day 02:12 pm"},
		{"dav", "15 juma", "15 week"},
		// de
		{"de", "nächstes jahr", "in 1 year"},
		{"de", "letzte woche 04:25 nachm", "1 week ago 04:25 pm"},
		// dje
		{"dje", "hõo 08:08 subbaahi", "0 day ago 08:08 am"},
		{"dje", "suba", "in 1 day"},
		{"dje", "7 handu", "7 month"},
		// dsb
		{"dsb", "pśed 10 góźinami", "10 hour ago"},
		{"dsb", "za 43 minutow", "in 43 minute"},
		{"dsb", "pśiducy tyźeń", "in 1 week"},
		// dua
		{"dua", "kíɛlɛ nítómb́í", "1 day ago"},
		{"dua", "12 ŋgandɛ", "12 hour"},
		{"dua", "wɛ́ŋgɛ̄", "0 day ago"},
		// dyo
		{"dyo", "fucen", "1 day ago"},
		{"dyo", "kajom", "in 1 day"},
		{"dyo", "6 fuleeŋ", "6 month"},
		// dz
		{"dz", "ནངས་པ་", "in 1 day"},
		{"dz", "སྐར་མ་ 3 ནང་", "in 3 minute"},
		{"dz", "ལོ་འཁོར་ 21 ཧེ་མ་", "21 year ago"},
		// ebu
		{"ebu", "ĩgoro", "1 day ago"},
		{"ebu", "2 ndagĩka", "2 minute"},
		{"ebu", "rũciũ", "in 1 day"},
		// ee
		{"ee", "ɣleti si va yi", "1 month ago"},
		{"ee", "ƒe 24 si wo va yi", "24 year ago"},
		{"ee", "le sekend 20 me", "in 20 second"},
		// el
		{"el", "πριν από 45 λεπτό", "45 minute ago"},
		{"el", "σε 22 μήνες", "in 22 month"},
		{"el", "επόμενη εβδομάδα 12:09 μμ", "in 1 week 12:09 pm"},
		// et
		{"et", "eelmine nädal", "1 week ago"},
		{"et", "1 a pärast", "in 1 year"},
		{"et", "4 tunni eest", "4 hour ago"},
		// eu
		{"eu", "aurreko hilabetea", "1 month ago"},
		{"eu", "duela 15 segundo", "15 second ago"},
		{"eu", "2 hilabete barru", "in 2 month"},
		// ewo
		{"ewo", "okírí", "in 1 day"},
		{"ewo", "angogé 10:15 kíkíríg", "1 day ago 10:15 am"},
		{"ewo", "5 m̀bú", "5 year"},
		// ff
		{"ff", "hannde", "0 day ago"},
		{"ff", "haŋki 01:14 subaka", "1 day ago 01:14 am"},
		{"ff", "2 yontere", "2 week"},
		// fil
		{"fil", "22 min ang nakalipas", "22 minute ago"},
		{"fil", "sa 5 taon", "in 5 year"},
		{"fil", "nakalipas na linggo", "1 week ago"},
		// fo
		{"fo", "seinasta mánað", "1 month ago"},
		{"fo", "um 3 viku", "in 3 week"},
		{"fo", "7 tímar síðan", "7 hour ago"},
		// fur
		{"fur", "ca di 16 setemanis", "in 16 week"},
		{"fur", "15 secont indaûr", "15 second ago"},
		{"fur", "doman", "in 1 day"},
		// fy
		{"fy", "folgjende moanne", "in 1 month"},
		{"fy", "oer 24 oere", "in 24 hour"},
		{"fy", "2 deien lyn", "2 day ago"},
		// ga
		{"ga", "i gceann 6 nóiméad", "in 6 minute"},
		{"ga", "12 seachtain ó shin", "12 week ago"},
		{"ga", "an bhliain seo chugainn", "in 1 year"},
		// gd
		{"gd", "an ceann 2 mhìosa", "in 2 month"},
		{"gd", "15 uair a thìde air ais", "15 hour ago"},
		{"gd", "am mìos seo chaidh", "1 month ago"},
		// gl
		{"gl", "hai 25 semanas", "25 week ago"},
		{"gl", "en 2 horas", "in 2 hour"},
		{"gl", "o ano pasado", "1 year ago"},
		// gsw
		{"gsw", "moorn", "in 1 day"},
		{"gsw", "geschter", "1 day ago"},
		// gu
		{"gu", "2 વર્ષ પહેલા", "2 year ago"},
		{"gu", "આવતા મહિને", "in 1 month"},
		{"gu", "22 કલાક પહેલાં", "22 hour ago"},
		// guz
		{"guz", "mambia", "in 1 day"},
		{"guz", "igoro", "1 day ago"},
		// ha
		{"ha", "gobe", "in 1 day"},
		{"ha", "jiya", "1 day ago"},
		// hr
		{"hr", "prije 3 dana", "3 day ago"},
		{"hr", "sljedeći mjesec", "in 1 month"},
		{"hr", "za 2 sati", "in 2 hour"},
		// hsb
		{"hsb", "před 5 tydźenjemi", "5 week ago"},
		{"hsb", "za 60 sekundow", "in 60 second"},
		{"hsb", "lětsa", "0 year ago"},
		// hy
		{"hy", "հաջորդ ամիս", "in 1 month"},
		{"hy", "2 վայրկյան առաջ", "2 second ago"},
		{"hy", "3 տարուց", "in 3 year"},
		// id
		{"id", "5 tahun yang lalu", "5 year ago"},
		{"id", "dalam 43 menit", "in 43 minute"},
		{"id", "dlm 23 dtk", "in 23 second"},
		// ig
		{"ig", "nnyaafụ", "1 day ago"},
		{"ig", "taata", "0 day ago"},
		// is
		{"is", "í næstu viku", "in 1 week"},
		{"is", "fyrir 3 mánuðum", "3 month ago"},
		{"is", "eftir 2 klst", "in 2 hour"},
		// it
		{"it", "tra 3 minuti", "in 3 minute"},
		{"it", "5 giorni fa", "5 day ago"},
		{"it", "anno prossimo", "in 1 year"},
		// jgo
		{"jgo", "ɛ́ gɛ mɔ́ 20 háwa", "20 hour ago"},
		{"jgo", "ɛ́ gɛ́ mɔ́ pɛsaŋ 5", "5 month ago"},
		{"jgo", "nǔu ŋguꞌ 2", "in 2 year"},
		// jmc
		{"jmc", "ngama", "in 1 day"},
		{"jmc", "ukou", "1 day ago"},
		// ka
		{"ka", "ამ საათში", "0 hour ago"},
		{"ka", "7 კვირის წინ", "7 week ago"},
		{"ka", "6 წუთში", "in 6 minute"},
		// kab
		{"kab", "iḍelli", "1 day ago"},
		{"kab", "ass-a", "0 day ago"},
		// kam
		{"kam", "ũmũnthĩ", "0 day ago"},
		{"kam", "ũnĩ", "in 1 day"},
		// kde
		{"kde", "nundu", "in 1 day"},
		{"kde", "lido", "1 day ago"},
		// kea
		{"kea", "es simana li", "0 week ago"},
		{"kea", "di li 2 min", "in 2 minute"},
		{"kea", "a ten 6 anu", "6 year ago"},
		// khq
		{"khq", "hõo", "0 day ago"},
		{"khq", "bi", "1 day ago"},
		// ki
		{"ki", "rũciũ", "in 1 day"},
		{"ki", "ira", "1 day ago"},
		// kk
		{"kk", "3 апта бұрын", "3 week ago"},
		{"kk", "5 секундтан кейін", "in 5 second"},
		{"kk", "өткен ай", "1 month ago"},
		// kl
		{"kl", "om 8 sapaatip-akunnera", "in 8 week"},
		{"kl", "for 6 ukioq siden", "6 year ago"},
		{"kl", "om 56 minutsi", "in 56 minute"},
		// kln
		{"kln", "raini", "0 day ago"},
		{"kln", "mutai", "in 1 day"},
		// km
		{"km", "ម៉ោងនេះ", "0 hour ago"},
		{"km", "19 ខែមុន", "19 month ago"},
		{"km", "ក្នុង​រយៈ​ពេល 23 ម៉ោង", "in 23 hour"},
		// kn
		{"kn", "18 ತಿಂಗಳುಗಳ ಹಿಂದೆ", "18 month ago"},
		{"kn", "26 ಸೆಕೆಂಡ್‌ನಲ್ಲಿ", "in 26 second"},
		{"kn", "ಈ ನಿಮಿಷ", "0 minute ago"},
		// ko
		{"ko", "2분 후", "in 2 minute"},
		{"ko", "5년 전", "5 year ago"},
		{"ko", "다음 달", "in 1 month"},
		// ksb
		{"ksb", "keloi", "in 1 day"},
		{"ksb", "evi eo", "0 day ago"},
		// ksf
		{"ksf", "ridúrǝ́", "in 1 day"},
		{"ksf", "rinkɔɔ́", "1 day ago"},
		// ksh
		{"ksh", "nächste woche", "in 1 week"},
		{"ksh", "en 3 johre", "in 3 year"},
		{"ksh", "diese mohnd", "0 month ago"},
		// ky
		{"ky", "ушул мүнөттө", "0 minute ago"},
		{"ky", "6 айд кийин", "in 6 month"},
		{"ky", "5 мүнөт мурун", "5 minute ago"},
		// lag
		{"lag", "lamʉtoondo", "in 1 day"},
		{"lag", "niijo", "1 day ago"},
		// lb
		{"lb", "virun 2 stonn", "2 hour ago"},
		{"lb", "an 5 joer", "in 5 year"},
		{"lb", "leschte mount", "1 month ago"},
		// lg
		{"lg", "nkya", "in 1 day"},
		{"lg", "ggulo", "1 day ago"},
		// lkt
		{"lkt", "hékta wíyawapi 8 k'uŋ héhaŋ", "8 month ago"},
		{"lkt", "tȟokáta okó kiŋháŋ", "in 1 week"},
		{"lkt", "letáŋhaŋ owápȟe 4 kiŋháŋ", "in 4 hour"},
		// ln
		{"ln", "lóbi elékí", "1 day ago"},
		{"ln", "lóbi ekoyâ", "in 1 day"},
		// lo
		{"lo", "ໃນອີກ 5 ຊົ່ວໂມງ", "in 5 hour"},
		{"lo", "3 ປີກ່ອນ", "3 year ago"},
		// lt
		{"lt", "praėjusią savaitę", "1 week ago"},
		{"lt", "prieš 12 mėnesį", "12 month ago"},
		{"lt", "po 2 valandų", "in 2 hour"},
		// lu
		{"lu", "lelu", "0 day ago"},
		{"lu", "makelela", "1 day ago"},
		// luo
		{"luo", "nyoro", "1 day ago"},
		{"luo", "kiny", "in 1 day"},
		// luy
		{"luy", "mgorova", "1 day ago"},
		{"luy", "lero", "0 day ago"},
		// lv
		{"lv", "pēc 67 minūtes", "in 67 minute"},
		{"lv", "pirms 5 nedēļām", "5 week ago"},
		{"lv", "nākamajā gadā", "in 1 year"},
		// mas
		{"mas", "tááisérè", "in 1 day"},
		{"mas", "ŋolé", "1 day ago"},
		// mer
		{"mer", "ĩgoro", "1 day ago"},
		{"mer", "narua", "0 day ago"},
		// mfe
		{"mfe", "zordi", "0 day ago"},
		{"mfe", "demin", "in 1 day"},
		// mg
		{"mg", "rahampitso", "in 1 day"},
		{"mg", "omaly", "1 day ago"},
		// mgh
		{"mgh", "lel'lo", "0 day ago"},
		{"mgh", "n'chana", "1 day ago"},
		// mgo
		{"mgo", "ikwiri", "1 day ago"},
		{"mgo", "isu", "in 1 day"},
		// mk
		{"mk", "пред 4 минута", "4 minute ago"},
		{"mk", "за 6 месеци", "in 6 month"},
		{"mk", "минатата година", "1 year ago"},
		// ml
		{"ml", "ഈ മിനിറ്റിൽ", "0 minute ago"},
		{"ml", "7 മണിക്കൂറിൽ", "in 7 hour"},
		{"ml", "2 വർഷം മുമ്പ്", "2 year ago"},
		// mn
		{"mn", "5 цагийн өмнө", "5 hour ago"},
		{"mn", "10 жилийн дараа", "in 10 year"},
		{"mn", "өнгөрсөн долоо хоног", "1 week ago"},
		// mr
		{"mr", "2 मिनिटांमध्ये", "in 2 minute"},
		{"mr", "5 महिन्यापूर्वी", "5 month ago"},
		{"mr", "हे वर्ष", "0 year ago"},
		// ms
		{"ms", "dalam 7 hari", "in 7 day"},
		{"ms", "3 thn lalu", "3 year ago"},
		{"ms", "bulan depan", "in 1 month"},
		// mt
		{"mt", "ix-xahar li għadda", "1 month ago"},
		{"mt", "2 sena ilu", "2 year ago"},
		{"mt", "il-ġimgħa d-dieħla", "in 1 week"},
		// mua
		{"mua", "tǝ'nahko", "0 day ago"},
		{"mua", "tǝ'nane", "in 1 day"},
		// my
		{"my", "ပြီးခဲ့သည့် 7 မိနစ်", "7 minute ago"},
		{"my", "12 လအတွင်း", "in 12 month"},
		{"my", "ယခု သီတင်းပတ်", "0 week ago"},
		// nb
		{"nb", "om 6 timer", "in 6 hour"},
		{"nb", "om 2 måneder", "in 2 month"},
		{"nb", "forrige uke", "1 week ago"},
		{"nb", "for 3 dager siden", "3 day ago"},
		{"nb", "for 3 timer siden", "3 hour ago"},
		{"nb", "3 dager siden", "3 day ago"},
		{"nb", "3 mnd siden", "3 month ago"},
		{"nb", "2 uker siden", "2 week ago"},
		{"nb", "1 uke siden", "1 week ago"},
		{"nb", "10 timer siden", "10 hour ago"},
		// nd
		{"nd", "kusasa", "in 1 day"},
		{"nd", "izolo", "1 day ago"},
		// ne
		{"ne", "5 वर्ष अघि", "5 year ago"},
		{"ne", "35 मिनेटमा", "in 35 minute"},
		{"ne", "यो हप्ता", "0 week ago"},
		// nl
		{"nl", "15 dgn geleden", "15 day ago"},
		{"nl", "over 2 maand", "in 2 month"},
		{"nl", "vorige jaar", "1 year ago"},
		// nmg
		{"nmg", "nakugú", "1 day ago"},
		{"nmg", "namáná", "in 1 day"},
		// nn
		{"nn", "for 5 minutter siden", "5 minute ago"},
		{"nn", "om 3 uker", "in 3 week"},
		{"nn", "i morgon", "in 1 day"},
		// nnh
		{"nnh", "jǔɔ gẅie à ne ntóo", "in 1 day"},
		{"nnh", "jǔɔ gẅie à ka tɔ̌g", "1 day ago"},
		// nus
		{"nus", "ruun", "in 1 day"},
		{"nus", "walɛ 06:23 tŋ", "0 day ago 06:23 pm"},
		// nyn
		{"nyn", "nyomwabazyo", "1 day ago"},
		{"nyn", "erizooba", "0 day ago"},
		// os
		{"os", "3 боны размӕ", "3 day ago"},
		{"os", "47 сахаты фӕстӕ", "in 47 hour"},
		{"os", "знон", "1 day ago"},
		// pa-Guru
		{"pa-Guru", "ਅਗਲਾ ਹਫ਼ਤਾ", "in 1 week"},
		{"pa-Guru", "5 ਮਹੀਨੇ ਪਹਿਲਾਂ", "5 month ago"},
		{"pa-Guru", "22 ਮਿੰਟਾਂ ਵਿੱਚ", "in 22 minute"},
		// pa
		{"pa", "15 ਘੰਟੇ ਪਹਿਲਾਂ", "15 hour ago"},
		{"pa", "16 ਸਕਿੰਟ ਵਿੱਚ", "in 16 second"},
		{"pa", "ਅਗਲਾ ਸਾਲ", "in 1 year"},
		// pl
		{"pl", "6 tygodni temu", "6 week ago"},
		{"pl", "za 8 lat", "in 8 year"},
		{"pl", "ta minuta", "0 minute ago"},
		// rm
		{"rm", "damaun", "in 1 day"},
		{"rm", "ier", "1 day ago"},
		// ro
		{"ro", "acum 2 de ore", "2 hour ago"},
		{"ro", "peste 5 de ani", "in 5 year"},
		{"ro", "săptămâna trecută", "1 week ago"},
		// rof
		{"rof", "linu", "0 day ago"},
		{"rof", "ng'ama", "in 1 day"},
		// ru
		{"ru", "12 секунд назад", "12 second ago"},
		{"ru", "через 8 месяцев", "in 8 month"},
		{"ru", "в прошлом году", "1 year ago"},
		// rwk
		{"rwk", "ukou", "1 day ago"},
		{"rwk", "ngama", "in 1 day"},
		// sah
		{"sah", "20 чаас ынараа өттүгэр", "20 hour ago"},
		{"sah", "50 сылынан", "in 50 year"},
		{"sah", "ааспыт нэдиэлэ", "1 week ago"},
		// saq
		{"saq", "duo", "0 day ago"},
		{"saq", "taisere", "in 1 day"},
		// sbp
		{"sbp", "pamulaawu", "in 1 day"},
		{"sbp", "ineng'uni", "0 day ago"},
		// se
		{"se", "51 minuhta maŋŋilit", "in 51 minute"},
		{"se", "3 jahkki árat", "3 year ago"},
		{"se", "ihttin", "in 1 day"},
		// seh
		{"seh", "manguana", "in 1 day"},
		{"seh", "zuro", "1 day ago"},
		// ses
		{"ses", "suba", "in 1 day"},
		{"ses", "hõo", "0 day ago"},
		// sg
		{"sg", "bîrï", "1 day ago"},
		{"sg", "lâsô", "0 day ago"},
		// shi-Latn
		{"shi-Latn", "iḍlli", "1 day ago"},
		{"shi-Latn", "askka 06:15 tifawt", "in 1 day 06:15 am"},
		// shi-Tfng
		{"shi-Tfng", "ⴰⵙⴽⴽⴰ", "in 1 day"},
		{"shi-Tfng", "ⴰⵙⵙⴰ", "0 day ago"},
		// shi
		{"shi", "ⵉⴹⵍⵍⵉ", "1 day ago"},
		// si
		{"si", "තත්පර 14කින්", "in 14 second"},
		{"si", "වසර 2කට පෙර", "2 year ago"},
		{"si", "මෙම සතිය", "0 week ago"},
		// sk
		{"sk", "pred 11 týždňami", "11 week ago"},
		{"sk", "o 25 rokov", "in 25 year"},
		{"sk", "v tejto hodine", "0 hour ago"},
		// sl
		{"sl", "pred 4 dnevom", "4 day ago"},
		{"sl", "čez 76 leto", "in 76 year"},
		{"sl", "naslednji mesec", "in 1 month"},
		// sn
		{"sn", "mangwana", "in 1 day"},
		{"sn", "nhasi", "0 day ago"},
		// so
		{"so", "berri", "in 1 day"},
		{"so", "shalay", "1 day ago"},
		// sq
		{"sq", "pas 6 muajsh", "in 6 month"},
		{"sq", "72 orë më parë", "72 hour ago"},
		{"sq", "javën e ardhshme", "in 1 week"},
		// sr-Cyrl
		{"sr-Cyrl", "пре 5 година", "5 year ago"},
		{"sr-Cyrl", "за 52 нед", "in 52 week"},
		{"sr-Cyrl", "данас", "0 day ago"},
		// sr-Latn
		{"sr-Latn", "za 120 sekundi", "in 120 second"},
		{"sr-Latn", "pre 365 dana", "365 day ago"},
		{"sr-Latn", "prošle nedelje", "1 week ago"},
		// sr
		{"sr", "пре 40 сати", "40 hour ago"},
		{"sr", "за 100 год", "in 100 year"},
		{"sr", "овог месеца", "0 month ago"},
		// sv
		{"sv", "för 15 vecka sedan", "15 week ago"},
		{"sv", "om 2 sekunder", "in 2 second"},
		{"sv", "förra året", "1 year ago"},
		// sw
		{"sw", "sekunde 25 zilizopita", "25 second ago"},
		{"sw", "miezi 5 iliyopita", "5 month ago"},
		{"sw", "mwaka uliopita", "1 year ago"},
		// ta
		{"ta", "7 நாட்களுக்கு முன்", "7 day ago"},
		{"ta", "45 ஆண்டுகளில்", "in 45 year"},
		{"ta", "இப்போது", "0 second ago"},
		// te
		{"te", "12 గంటల క్రితం", "12 hour ago"},
		{"te", "25 సంవత్సరాల్లో", "in 25 year"},
		{"te", "గత వారం", "1 week ago"},
		// teo
		{"teo", "moi", "in 1 day"},
		{"teo", "lolo", "0 day ago"},
		// to
		{"to", "miniti 'e 5 kuo'osi", "5 minute ago"},
		{"to", "'i he ta'u 'e 6", "in 6 year"},
		{"to", "'aneafi", "1 day ago"},
		// tr
		{"tr", "11 saat önce", "11 hour ago"},
		{"tr", "10 yıl sonra", "in 10 year"},
		{"tr", "geçen ay", "1 month ago"},
		// twq
		{"twq", "hõo", "0 day ago"},
		{"twq", "suba", "in 1 day"},
		// tzm
		{"tzm", "assenaṭ", "1 day ago"},
		{"tzm", "asekka", "in 1 day"},
		// uk
		{"uk", "18 хвилину тому", "18 minute ago"},
		{"uk", "через 22 року", "in 22 year"},
		{"uk", "цього тижня", "0 week ago"},
		// uz-Cyrl
		{"uz-Cyrl", "кейинги ой", "in 1 month"},
		{"uz-Cyrl", "30 йил аввал", "30 year ago"},
		{"uz-Cyrl", "59 сониядан сўнг", "in 59 second"},
		// uz-Latn
		{"uz-Latn", "3 haftadan keyin", "in 3 week"},
		{"uz-Latn", "5 soat oldin", "5 hour ago"},
		{"uz-Latn", "shu yil", "0 year ago"},
		// uz
		{"uz", "25 soat oldin", "25 hour ago"},
		{"uz", "8 yildan keyin", "in 8 year"},
		{"uz", "bugun", "0 day ago"},
		// vi
		{"vi", "sau 22 giờ nữa", "in 22 hour"},
		{"vi", "15 tháng trước", "15 month ago"},
		{"vi", "tuần sau", "in 1 week"},
		// vun
		{"vun", "ngama", "in 1 day"},
		{"vun", "ukou", "1 day ago"},
		// wae
		{"wae", "vor 11 minüta", "11 minute ago"},
		{"wae", "i 15 wuča", "in 15 week"},
		{"wae", "hitte", "0 day ago"},
		// xog
		{"xog", "enkyo", "in 1 day"},
		{"xog", "edho", "1 day ago"},
		// yav
		{"yav", "ínaan", "0 day ago"},
		{"yav", "nakinyám", "in 1 day"},
		// yo
		{"yo", "ọ̀la", "in 1 day"},
		{"yo", "òní", "0 day ago"},
		// yue
		{"yue", "13 個星期後", "in 13 week"},
		{"yue", "2 小時前", "2 hour ago"},
		{"yue", "上個月", "1 month ago"},
		// zgh
		{"zgh", "ⴰⵙⵙⴰ", "0 day ago"},
		{"zgh", "ⵉⴹⵍⵍⵉ", "1 day ago"},
		// zh-Hans
		{"zh-Hans", "3秒钟后", "in 3 second"},
		{"zh-Hans", "4年后", "in 4 year"},
		{"zh-Hans", "上周", "1 week ago"},
		// zh-Hant
		{"zh-Hant", "7 分鐘後", "in 7 minute"},
		{"zh-Hant", "12 個月前", "12 month ago"},
		{"zh-Hant", "這一小時", "0 hour ago"},
		// zu
		{"zu", "10 amaminithi edlule", "10 minute ago"},
		{"zu", "20 unyaka odlule", "20 year ago"},
		{"zu", "manje", "0 second ago"},
	}

	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("%s, \"%s\"", test.Locale, test.String)

		// Load locale
		ld, err := GetLocale(test.Locale)
		assert.Nil(t, err, message)

		// Prepare config: finnish language use "t" as hour, so empty SKIP_TOKENS.
		var cfg *setting.Configuration
		if test.Locale != "fi" {
			cfg = &setting.DefaultConfig
		}

		// Translate string
		translation := Translate(*ld, test.String, false, cfg)
		passed := assert.Equal(t, test.Expected, translation, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
