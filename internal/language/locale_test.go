package language

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
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

	tests = []testScenario{
		{"ar", "6 يناير، 2015، الساعة 05:16 مساءً", "6 january 2015 05:16 pm"},
	}

	nFailed := 0
	for _, test := range tests {
		message := fmt.Sprintf("%s, \"%s\"", test.Locale, test.String)

		ld, err := GetLocale(test.Locale)
		assert.Nil(t, err, message)

		translation := Translate(*ld, test.String, false)
		passed := assert.Equal(t, test.Expected, translation, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
