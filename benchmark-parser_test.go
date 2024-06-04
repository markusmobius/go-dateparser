package dateparser_test

import (
	"fmt"
	"testing"

	dps "github.com/markusmobius/go-dateparser"
)

func BenchmarkParser_Parse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, text := range benchmarkParserTexts {
			_, err := dps.Parse(nil, text)
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

var benchmarkParserTexts = []string{
	// === ABSOLUTE PARSER ===
	// English dates
	"[Sept] 04, 2014.",
	"Tuesday Jul 22, 2014",
	"Tues 9th Aug, 2015",
	"10:04am",
	"Friday",
	"November 19, 2014 at noon",
	"December 13, 2014 at midnight",
	"Nov 25 2014 10:17 pm",
	"Wed Aug 05 12:00:00 2015",
	"April 9, 2013 at 6:11 a.m.",
	"Aug. 9, 2012 at 2:57 p.m.",
	"December 10, 2014, 11:02:21 pm",
	"8:25 a.m. Dec. 12, 2014",
	"2:21 p.m., December 11, 2014",
	"Fri, 12 Dec 2014 10:55:50",
	"20 Mar 2013 10h11",
	"10:06am Dec 11, 2014",
	"19 February 2013 year 09:10",
	"21 January 2012 13:11:23.678",
	"1/1/16 9:02:43.1",
	"29.02.2020 13.12",
	"Wednesday, 22nd June, 2016, 12.16 pm.",
	"16:00",
	"Monday 7:15 AM",
	// French dates
	"11 Mai 2014",
	"11 sept. 2014",
	"dimanche, 11 Mai 2014",
	"22 janvier 2015 à 14h40",
	"Dimanche 1er Février à 21:24",
	"vendredi, décembre 5 2014.",
	"le 08 Déc 2014 15:11",
	"Le 11 Décembre 2014 à 09:00",
	"fév 15, 2013",
	"Jeu 15:12",
	// Spanish dates
	"Martes 21 de Octubre de 2014",
	"Miércoles 20 de Noviembre de 2013",
	"12 de junio del 2012",
	"13 Ago, 2014",
	"13 Septiembre, 2014",
	"11 Marzo, 2014",
	"julio 5, 2015 en 1:04 pm",
	"Vi 17:15",
	// Dutch dates
	"11 augustus 2014",
	"14 januari 2014",
	"vr jan 24, 2014 12:49",
	// Italian dates
	"16 giu 2014",
	"26 gennaio 2014",
	"Ven 18:23",
	// Portuguese dates
	"sexta-feira, 10 de junho de 2014 14:52",
	"13 Setembro, 2014",
	"Sab 3:03",
	// Russian dates
	"10 мая",
	"26 апреля",
	"20 ноября 2013",
	"28 октября 2014 в 07:54",
	"13 января 2015 г. в 13:34",
	"09 августа 2012",
	"Авг 26, 2015 15:12",
	"2 Декабрь 95 11:15",
	"13 янв. 2005 19:13",
	"13 авг. 2005 19:13",
	"13 авг. 2005г. 19:13",
	"13 авг. 2005 г. 19:13",
	"21 сентября 2021г., вторник",
	// Turkish dates
	"11 Ağustos, 2014",
	"08.Haziran.2014, 11:07",
	"17.Şubat.2014, 17:51",
	"14-Aralık-2012, 20:56",
	// Romanian dates
	"13 iunie 2013",
	"14 aprilie 2014",
	"18 martie 2012",
	"12-Iun-2013",
	// German dates
	"21. Dezember 2013",
	"19. Februar 2012",
	"26. Juli 2014",
	"18.10.14 um 22:56 Uhr",
	"12-Mär-2014",
	"Mit 13:14",
	"23. März 18.37 Uhr",
	// Czech dates
	"pon 16. čer 2014 10:07:43",
	"13 Srpen, 2014",
	"čtv 14. lis 2013 12:38:43",
	// Thai dates
	"ธันวาคม 11, 2014, 08:55:08 PM",
	"22 พฤษภาคม 2012, 22:12",
	"11 กุมภา 2020, 8:13 AM",
	"1 เดือนตุลาคม 2005, 1:00 AM",
	"11 ก.พ. 2020, 1:13 pm",
	// Vietnamese dates
	"Thứ năm",
	"Thứ sáu",
	"Tháng Mười Hai 29, 2013, 14:14",
	"05 Tháng một 2015 - 03:54 AM",
	// Belarusian dates
	"11 траўня",
	"4 мая",
	"Чацвер 06 жніўня 2015",
	"Нд 14 сакавіка 2015 у 7 гадзін 10 хвілін",
	"5 жніўня 2015 года у 13:34",
	// Ukrainian dates
	"2015-кві-12",
	"21 чер 2013 3:13",
	"12 лютого 2012, 13:12:23",
	"вів о 14:04",
	// Tagalog dates
	"12 Hulyo 2003 13:01",
	"1978, 1 Peb, 7:05 PM",
	"2 hun",
	"Lin 16:16",
	// Japanese dates
	"2016年3月20日(日) 21時40分",
	"2016年3月20日 21時40分",
	// Numeric dates
	"06-17-2014",
	"13/03/2014",
	"11. 12. 2014, 08:45:39",
	// Miscellaneous dates
	"1 Ni 2015",
	"1 Mar 2015",
	"1 сер 2015",
	// Chinese dates
	"2015年04月08日10:05",
	"2012年12月20日10:35",
	"2016年06月30日09时30分",
	"2016年6月2911:30",
	"2016年6月29",
	"2016年 2月 5日",
	"2016年9月14日晚8:00",
	// Bulgarian
	"25 ян 2016",
	"23 декември 2013 15:10:01",
	// Bangla dates
	"[সেপ্টেম্বর] 04, 2014.",
	"মঙ্গলবার জুলাই 22, 2014",
	"শুক্রবার",
	"শুক্র, 12 ডিসেম্বর 2014 10:55:50",
	"1লা জানুয়ারী 2015",
	"25শে মার্চ 1971",
	"8ই মে 2002",
	"10:06am ডিসেম্বর 11, 2014",
	"19 ফেব্রুয়ারী 2013 সাল 09:10",
	// Hindi dates
	"11 जुलाई 1994, 11:12",
	"१७ अक्टूबर २०१८",
	"12 जनवरी  1997 11:08 अपराह्न",
	// Georgian dates
	"2011 წლის 17 მარტი, ოთხშაბათი",
	"2015 წ. 12 ივნ, 15:34",
	// Finnish dates
	"5.7.2018 5.45 ip.",
	// Others (from `test_clean_api.py`)
	"24 de Janeiro de 2014",
	"2 de Enero de 2013",
	"January 25, 2014",
	"May 5, 2000 13:00",
	"August 8, 2018 5 PM",
	"February 26, 1981 5 am UTC",
	"2014/11/17 14:56 EDT",
	"08-08-2014\xa018:29",
	"12 jan 1876",
	"02/09/16",
	"10 giu 2018",

	// === CONTAINS TIMEZONE ===
	"Sep 03 2014 | 4:32 pm EDT",
	"17th October, 2034 @ 01:08 am PDT",
	"15 May 2004 23:24 EDT",
	"08/17/14 17:00 (PDT)",
	"15 May 2004 16:10 -0400",
	"1999-12-31 19:00:00 -0500",
	"1999-12-31 19:00:00 +0500",
	"Fri, 09 Sep 2005 13:51:39 -0700",
	"Fri, 09 Sep 2005 13:51:39 +0000",
	"Mon, 25 Jun 2018 10:37:47 +0530 ",
	"Fri Sep 23 2016 10:34:51 GMT+0800 (CST)",

	// === NOSPACE PARSER ===
	"2016020417:10",

	// === ISO DATESTAMP ===
	"2015-05-02T10:20:19+0000",

	// === TIMESTAMP ===
	"1484823450",
	"1436745600000",
	"1015673450",
	"2016-09-23T02:54:32.845Z",

	// === RELATIVE PAST TIMES ===
	// Mixed temporal nouns
	"today",
	"Today",
	"TODAY",
	"Сегодня",
	"Hoje",
	"Oggi",
	"yesterday",
	" Yesterday \n",
	"Ontem",
	"Ieri",
	"the day before yesterday",
	"The DAY before Yesterday",
	"Anteontem",
	"Avant-hier",
	"вчера",
	"снощи",

	// English dates
	"yesterday",
	"yesterday at 11:30",
	"1 decade",
	"1 decade 2 years",
	"1 decade 12 months",
	"1 decade and 11 months",
	"last decade",
	"a decade ago",
	"100 decades",
	"yesterday",
	"the day before yesterday",
	"today",
	"an hour ago",
	"about an hour ago",
	"a day ago",
	"a week ago",
	"2 hours ago",
	"about 23 hours ago",
	"1 year 2 months",
	"1 year, 09 months,01 weeks",
	"1 year 11 months",
	"1 year 12 months",
	"15 hr",
	"15 hrs",
	"2 min",
	"2 mins",
	"3 sec",
	"1000 years ago",
	"2013 years ago",
	"5000 months ago",
	fmt.Sprintf("%d months ago", 2013*12+8),
	"1 year, 1 month, 1 week, 1 day, 1 hour and 1 minute ago",
	"just now",
	// Fix for #291, work till one to twelve only
	"nine hours ago",
	"three week ago",
	"eight months ago",
	"six days ago",
	"five years ago",

	// French dates
	"Aujourd'hui",
	"Aujourd’hui",
	"Aujourdʼhui",
	"Aujourdʻhui",
	"Aujourd՚hui",
	"Aujourdꞌhui",
	"Aujourd＇hui",
	"Aujourd′hui",
	"Aujourd‵hui",
	"Aujourdʹhui",
	"Aujourd＇hui",
	"moins de 21s",
	"moins de 21m",
	"moins de 21h",
	"moins de 21 minute",
	"moins de 21 heure",
	"Hier",
	"Avant-hier",
	"Il ya un jour",
	"Il ya une heure",
	"Il ya 2 heures",
	"Il ya environ 23 heures",
	"1 an 2 mois",
	"1 année, 09 mois, 01 semaines",
	"1 an 11 mois",
	"Il ya 1 an, 1 mois, 1 semaine, 1 jour, 1 heure et 1 minute",
	"Il y a 40 min",

	// German dates
	"Heute",
	"Gestern",
	"vorgestern",
	"vor einem Tag",
	"vor einer Stunden",
	"Vor 2 Stunden",
	"vor etwa 23 Stunden",
	"1 Jahr 2 Monate",
	"1 Jahr, 09 Monate, 01 Wochen",
	"1 Jahr 11 Monate",
	"vor 29h",
	"vor 29m",
	"1 Jahr, 1 Monat, 1 Woche, 1 Tag, 1 Stunde und 1 Minute",

	// Italian dates
	"oggi",
	"ieri",
	"2 ore fa",
	"circa 23 ore fa",
	"1 anno 2 mesi",
	"1 anno, 09 mesi, 01 settimane",
	"1 anno 11 mesi",
	"1 anno, 1 mese, 1 settimana, 1 giorno, 1 ora e 1 minuto fa",

	// Portuguese dates
	"ontem",
	"anteontem",
	"hoje",
	"uma hora atrás",
	"1 segundo atrás",
	"um dia atrás",
	"uma semana atrás",
	"2 horas atrás",
	"cerca de 23 horas atrás",
	"1 ano 2 meses",
	"1 ano, 09 meses, 01 semanas",
	"1 ano 11 meses",
	"1 ano, 1 mês, 1 semana, 1 dia, 1 hora e 1 minuto atrás",

	// Turkish dates
	"Dün",
	"Bugün",
	"2 saat önce",
	"yaklaşık 23 saat önce",
	"1 yıl 2 ay",
	"1 yıl, 09 ay, 01 hafta",
	"1 yıl 11 ay",
	"1 yıl, 1 ay, 1 hafta, 1 gün, 1 saat ve 1 dakika önce",

	// Russian dates
	"сегодня",
	"Вчера в",
	"вчера",
	"2 часа назад",
	"час назад",
	"минуту назад",
	"2 ч. 21 мин. назад",
	"около 23 часов назад",
	"1 год 2 месяца",
	"1 год, 09 месяцев, 01 недель",
	"1 год 11 месяцев",
	"1 год, 1 месяц, 1 неделя, 1 день, 1 час и 1 минуту назад",

	// Czech dates
	"Dnes",
	"Včera",
	"Předevčírem",
	"Před 2 hodinami",
	"před přibližně 23 hodin",
	"1 rok 2 měsíce",
	"1 rok, 09 měsíců, 01 týdnů",
	"1 rok 11 měsíců",
	"3 dny",
	"3 hodiny",
	"2 roky, 2 týdny, 1 den, 1 hodinu, 5 vteřin před",
	"1 rok, 1 měsíc, 1 týden, 1 den, 1 hodina, 1 minuta před",

	// Spanish dates
	"anteayer",
	"ayer",
	"hoy",
	"hace una hora",
	"Hace un día",
	"Hace una semana",
	"Hace 2 horas",
	"Hace cerca de 23 horas",
	"1 año 2 meses",
	"1 año, 09 meses, 01 semanas",
	"1 año 11 meses",
	"Hace 1 año, 1 mes, 1 semana, 1 día, 1 hora y 1 minuto",

	// Chinese dates
	"昨天",
	"前天",
	"2小时前",
	"约23小时前",
	"1年2个月",
	"1年2個月",
	"1年11个月",
	"1年11個月",
	"1年，1月，1周，1天，1小时，1分钟前",

	// Arabic dates
	"اليوم",
	"يوم أمس",
	"منذ يومين",
	"منذ 3 أيام",
	"منذ 21 أيام",
	"1 عام, 1 شهر, 1 أسبوع, 1 يوم, 1 ساعة, 1 دقيقة",

	// Thai dates
	"วันนี้",
	"เมื่อวานนี้",
	"2 วัน",
	"2 ชั่วโมง",
	"23 ชม.",
	"2 สัปดาห์ 3 วัน",
	"1 ปี 9 เดือน 1 สัปดาห์",
	"1 ปี 1 เดือน 1 สัปดาห์ 1 วัน 1 ชั่วโมง 1 นาที",

	// Vietnamese dates
	"Hôm nay",
	"Hôm qua",
	"2 giờ",
	"2 tuần 3 ngày",
	// Following test unsupported, refer to discussion at:
	// http://github.com/scrapinghub/dateparser/issues/33
	// {"1 năm 1 tháng 1 tuần 1 ngày 1 giờ 1 chút",

	// Belarusian dates
	"сёння",
	"учора ў",
	"ўчора",
	"пазаўчора",
	"2 гадзіны таму назад",
	"2 гадзіны таму",
	"гадзіну назад",
	"хвіліну таму",
	"2 гадзіны 21 хвіл. назад",
	"каля 23 гадзін назад",
	"1 год 2 месяцы",
	"1 год, 09 месяцаў, 01 тыдзень",
	"2 гады 3 месяцы",
	"5 гадоў, 1 месяц, 6 тыдняў, 3 дні, 5 гадзін 1 хвіліну і 3 секунды таму назад",

	// Polish dates
	"wczoraj",
	"1 godz. 2 minuty temu",
	"2 lata, 3 miesiące, 1 tydzień, 2 dni, 4 godziny, 15 minut i 25 sekund temu",
	"2 minuty temu",
	"15 minut temu",

	// Bulgarian dates
	"преди 3 дни",
	"преди час",
	"преди година",
	"вчера",
	"онзи ден",
	"днес",
	"преди час",
	"преди един ден",
	"преди седмица",
	"преди 2 часа",
	"преди около 23 часа",

	// Bangla dates
	// {"গতকাল",
	// {"আজ",
	"1 ঘন্টা আগে",
	"প্রায় 1 ঘন্টা আগে",
	"1 দিন আগে",
	"1 সপ্তাহ আগে",
	"2 ঘন্টা আগে",
	"প্রায় 23 ঘন্টা আগে",
	"1 বছর 2 মাস",
	"1 বছর, 09 মাস,01 সপ্তাহ",
	"1 বছর 11 মাস",
	"1 বছর 12 মাস",
	"15 ঘন্টা",
	"2 মিনিট",
	"3 সেকেন্ড",
	"1000 বছর আগে",
	"5000 মাস আগে",
	fmt.Sprintf("%d মাস আগে", 2013*12+8),
	"1 বছর, 1 মাস, 1 সপ্তাহ, 1 দিন, 1 ঘন্টা এবং 1 মিনিট আগে",
	// {"এখন",

	// Hindi dates
	"1 घंटे पहले",
	"15 मिनट पहले",
	"25 सेकंड पूर्व",
	"1 वर्ष, 8 महीने, 2 सप्ताह",
	"1 वर्ष 7 महीने",
	"आज",

	// af
	"2 uur gelede",
	"verlede maand",
	// agq
	"ā zūɛɛ",
	// ak
	"ndeda",
	// am
	"ከ8 ወራት በፊት",
	"ያለፈው ሳምንት",
	// as
	"কালি",
	"আজি",
	// asa
	"ighuo",
	// ast
	"hai 2 selmanes hai 3 díes",
	"l'añu pas el mes pasáu",
	// az-Latn
	"1 il öncə 2 ay öncə 3 həftə öncə",
	"6 saat öncə 5 dəqiqə öncə 4 saniyə öncə",
	// az
	"2 gün öncə 23 saat öncə",
	"5 dəqiqə öncə 27 saniyə öncə",
	// be
	"2 гадзіны таму 10 хвіліны таму",
	// bg
	"преди 3 месеца преди 2 седм",
	// bn
	"8 মিনিট আগে 42 সেকেন্ড পূর্বে",
	// br
	"7 eur zo 15 min zo 25 s zo",
	"14 sizhun zo 3 deiz zo",
	// bs-Cyrl
	"пре 5 година пре 7 месеци",
	"пре 5 сати пре 25 секунди",
	// bs-Latn
	"prije 20 sat 5 minuta",
	"prije 13 godina prije 3 sed",
	// bs
	"prije 3 mjeseci prije 12 sati",
	"prije 3 god 4 mj 5 sed 7 dan",
	// ca
	"fa 4 setmanes fa 5 segon",
	"fa 1 hora 2 minut 3 segon",
	// ce
	"10 кӏир хьалха 3 д хьалха",
	"12 сахь 30 мин 30 сек хьалха",
	// chr
	"ᎾᎿ 10 ᏒᎾᏙᏓᏆᏍᏗ ᏥᎨᏒ 5 ᎢᎦ ᏥᎨᏒ",
	// cs
	"před 3 rok 4 měsíc 5 den",
	"před 2 minutou před 45 sekundou",
	// cy
	"5 wythnos yn ôl",
	"7 munud 8 eiliad yn ôl",
	// dsb
	"pśed 15 góźinu 10 minuta 5 sekunda",
	"pśed 5 lětom, pśed 7 mjasecom",
	// ee
	"ŋkeke 12 si wo va yi",
	"ƒe 6 si va yi ɣleti 5 si va yi",
	// el
	"πριν από 5 ώρα 6 λεπτό 7 δευτερόλεπτο",
	"προηγούμενος μήνας",
	// es
	"hace 5 hora 2 minuto 3 segundo",
	// et
	"5 minut 12 sekundi eest",
	"11 aasta 11 kuu eest",
	// eu
	"duela 3 minutu",
	// fil
	"10 oras ang nakalipas",
	// fo
	"3 tími 12 minutt síðan",
	// fur
	"10 setemane 12 zornade indaûr",
	// fy
	"6 moannen lyn",
	// ga
	"12 uair an chloig ó shin",
	// gd
	"15 mhionaid air ais",
	// gl
	"hai 5 ano 7 mes",
	// gu
	"5 કલાક પહેલાં",
	// hr
	"prije 3 tjedana",
	// hsb
	"před 12 lětom 15 měsac",
	// hy
	"15 րոպե առաջ",
	// is
	"fyrir 3 ári fyrir 2 mánuði",
	// it
	"5 settimana fa",
	// jgo
	"ɛ́ gɛ́ mɔ́ pɛsaŋ 3",
	// ka
	"4 წლის წინ",
	// kk
	"5 сағат бұрын",
	// kl
	"for 6 ulloq unnuarlu siden",
	// km
	"11 សប្ដាហ៍​មុន",
	// kn
	"15 ಸೆಕೆಂಡುಗಳ ಹಿಂದೆ",
	// ko
	"12개월 전",
	// ksh
	"vör 15 johre",
	// ky
	"18 секунд мурун",
	// lb
	"virun 15 stonn",
	// lkt
	"hékta 8-čháŋ k'uŋ héhaŋ",
	// lt
	"prieš 20 minučių",
	// lv
	"pirms 10 gadiem",
	// mk
	"пред 13 часа",
	// ml
	"3 ആഴ്ച മുമ്പ്",
	// mn
	"15 секундын өмнө",
	// mr
	"25 वर्षापूर्वी",
	// ms
	"10 minit lalu",
	// my
	"ပြီးခဲ့သည့် 15 နှစ်",
	// nb
	"for 12 måneder siden",
	// ne
	"23 हप्ता पहिले",
	// nl
	"32 minuten geleden",
	// nn
	"for 15 sekunder siden",
	// os
	"35 сахаты размӕ",
	// pa-Guru
	"23 ਹਫ਼ਤੇ ਪਹਿਲਾਂ",
	// pa
	"7 ਸਾਲ ਪਹਿਲਾਂ",
	// ro
	"acum 56 de secunde",
	// sah
	"2 нэдиэлэ анараа өттүгэр",
	// se
	"8 jahkki árat",
	// si
	"මිනිත්තු 6කට පෙර",
	// sk
	"pred 20 hodinami 45 min",
	// sl
	"pred 15 tednom 10 dan",
	// sq
	"11 minutë më parë",
	// sr-Cyrl
	"пре 8 године 3 месец",
	// sr-Latn
	"pre 2 nedelje",
	// sr
	"пре 59 секунди",
	// sw
	"mwezi 2 uliopita",
	// ta
	"41 நிமிடங்களுக்கு முன்",
	// te
	"36 వారాల క్రితం",
	// to
	"houa 'e 7 kuo'osi",
	// tr
	"32 dakika önce",
	// uk
	"3 року тому",
	// uz-Cyrl
	"10 ҳафта олдин",
	// uz-Latn
	"3 oy oldin",
	// uz
	"45 soniya oldin",
	// vi
	"23 ngày trước",
	// wae
	"vor 15 stunde",
	// yue
	"5 分鐘前",
	// zh-Hans
	"3周前",
	// zh-Hant
	"2 年前",
	// zu
	"21 izinsuku ezedlule",

	// === RELATIVE FUTURE TIMES ===
	// English dates
	"in 1 decade 2 months",
	"in 100 decades",
	"in 1 decade 12 years",
	"next decade",
	"in a decade",
	"tomorrow",
	"today",
	"in an hour",
	"in about an hour",
	"in 1 day",
	"in a week",
	"in 2 hours",
	"in about 23 hours",
	"in 1 year 2 months",
	"in 1 year, 09 months,01 weeks",
	"in 1 year 11 months",
	"in 1 year 12 months",
	"in 15 hr",
	"in 15 hrs",
	"in 2 min",
	"in 2 mins",
	"in 3 sec",
	"in 1000 years",
	"in 5000 months",
	fmt.Sprintf("in %d months", 2013*12+8),
	"in 1 year, 1 month, 1 week, 1 day, 1 hour and 1 minute",
	"just now",

	// French dates
	"Aujourd'hui",
	"Dans un jour",
	"Dans une heure",
	"En 2 heures",
	"Dans environ 23 heures",
	"Dans 1 an 2 mois",
	"En 1 année, 09 mois, 01 semaines",
	"Dans 1 an 11 mois",
	"Dans 1 année, 1 mois, 1 semaine, 1 jour, 1 heure et 1 minute",
	"Dans 40 min",

	// German dates
	"Heute",
	"Morgen",
	"in einem Tag",
	"in einer Stunde",
	"in 2 Stunden",
	"in etwa 23 Stunden",
	"im 1 Jahr 2 Monate",
	"im 1 Jahr, 09 Monate, 01 Wochen",
	"im 1 Jahr 11 Monate",
	"im 1 Jahr, 1 Monat, 1 Woche, 1 Tag, 1 Stunde und 1 Minute",

	// Polish dates
	"jutro",
	"pojutrze",
	"za 2 lata, 3 miesiące, 1 tydzień, 2 dni, 4 godziny, 15 minut i 25 sekund",
	"za 2 minuty",
	"za 15 minut",

	// Turkish dates
	"yarın",
	"2 gün içerisinde",
	"4 ay içerisinde",
	"3 gün sonra",
	"2 ay sonra",
	"5 yıl 3 gün sonra",
	"5 gün içinde",
	"6 ay içinde",
	"5 yıl içinde",
	"5 sene içinde",
	"haftaya",
	"gelecek hafta",
	"gelecek ay",
	"gelecek yıl",

	// Hindi dates
	"1 वर्ष 10 महीने में",
	"15 घंटे बाद",
	"2 मिनट में",
	"17 सेकंड बाद",
	"1 वर्ष, 5 महीने, 1 सप्ताह में",

	// af
	"oor 10 jaar",
	"oor 5 min 3 sek",
	// am
	"በ2 ሳምንታት ውስጥ",
	"በ16 ቀናት ውስጥ",
	// ast
	"en 15 años",
	"en 20 minutos",
	// az-Latn
	"5 saniyə ərzində",
	"10 saat 20 dəqiqə ərzində",
	// az
	"15 il 6 ay ərzində",
	// be
	"праз 5 гадзіны 6 хвіліны",
	// bg
	"след 12 мин 18 сек",
	// bn
	"10 সেকেন্ডে",
	// br
	"a-benn 20 vloaz",
	"a-benn 15 deiz 20 eur",
	// bs-Cyrl
	"за 5 минут 10 секунд",
	"за 10 годину 11 месец",
	// bs-Latn
	"za 7 mjeseci",
	"za 6 dan 23 sat",
	// bs
	"za 15 sedmica",
	// ca
	"d'aquí a 10 anys",
	"d'aquí a 15 minut 53 segon",
	// ce
	"20 кӏира даьлча",
	"10 минот 25 секунд яьлча",
	// chr
	"ᎾᎿ 10 ᎧᎸᎢ",
	"ᎾᎿ 24 ᎢᏳᏟᎶᏓ",
	// cs
	"za 12 rok",
	"za 10 den 5 hodin",
	// cy
	"ymhen 15 mis",
	"ymhen 10 munud 8 eiliad",
	// da
	"om 10 minut 54 sekund",
	// de
	"in 15 jahren 10 monat",
	// dsb
	"za 10 mjasec",
	"za 30 min 50 sek",
	// dz
	"ལོ་འཁོར་ 4 ནང་",
	"སྐར་ཆ་ 20 ནང་",
	// ee
	"le ƒe 15 si gbɔna me",
	"le ŋkeke 2 wo me",
	// el
	"σε 5 ώρες",
	"σε 4 λεπτό 45 δευτ",
	// et
	"5 aasta 10 kuu pärast",
	"10 nädala pärast",
	// eu
	"15 hilabete barru",
	"20 egun barru",
	// fil
	"sa 8 segundo",
	"sa 2 oras 24 min",
	// fo
	"um 12 mánaðir",
	"um 10 tímar",
	// fur
	"ca di 15 setemanis",
	"ca di 15 minût 20 secont",
	// fy
	"oer 10 jier",
	"oer 22 deien",
	// ga
	"i gceann 23 bliain",
	"i gceann 12 scht",
	// gd
	"an ceann 10 bliadhna",
	"an ceann 18 latha",
	// gl
	"en 5 anos 26 mes",
	"en 14 semanas",
	// gu
	"10 મહિનામાં",
	"8 કલાકમાં",
	// hr
	"za 12 dana",
	"za 10 sat 43 min",
	// hsb
	"za 6 měsacow",
	"za 1 dźeń 12 hodź",
	// hy
	"7 ր-ից",
	"51 շաբաթից",
	// id
	"dalam 12 detik",
	"dalam 10 hari",
	// is
	"eftir 11 mínútur",
	"eftir 12 klukkustundir",
	// it
	"tra 5 minuto",
	"tra 16 settimane",
	// jgo
	// {"nǔu ŋgu" 10",
	"nǔu ŋgap-mbi 11",
	// ka
	"5 საათში",
	"3 კვირაში",
	// kea
	"di li 10 anu",
	"di li 43 minutu",
	// kk
	"10 сағаттан кейін",
	"18 айдан кейін",
	// kl
	"om 15 sapaatip-akunnera",
	"om 23 nalunaaquttap-akunnera",
	// km
	"2 នាទីទៀត",
	"5 សប្ដាហ៍ទៀត",
	// kn
	"10 ವಾರದಲ್ಲಿ",
	"15 ನಿಮಿಷಗಳಲ್ಲಿ",
	// ko
	"5초 후",
	"7개월 후",
	// ksh
	"en 8 johre",
	// ky
	"15 мүнөттөн кийин",
	"11 айд кийин",
	// lb
	"an 30 dag",
	"an 10 minutt 15 sekonn",
	// lkt
	"letáŋhaŋ okó 20 kiŋháŋ",
	"letáŋhaŋ ómakȟa 11 kiŋháŋ",
	// lo
	"ໃນອີກ 25 ຊົ່ວໂມງ",
	"ໃນອີກ 13 ອາທິດ",
	// lt
	"po 7 valandos",
	"po 5 min 5 sek",
	// lv
	"pēc 15 sekundēm",
	"pēc 10 mēneša",
	// mk
	"за 16 седмици",
	"за 2 месеци",
	// ml
	"5 ആഴ്ചയിൽ",
	"8 മിനിറ്റിൽ",
	// mn
	"10 сарын дараа",
	"15 цагийн дараа",
	// mr
	"2 महिन्यांमध्ये",
	"15 मिनि मध्ये",
	// ms
	"dalam 6 jam",
	"dalam 11 thn",
	// my
	"12 လအတွင်း",
	"8 နာရီအတွင်း",
	// nb
	"om 1 måneder",
	"om 5 minutter",
	// ne
	"10 वर्षमा",
	"15 घण्टामा",
	// nl
	"over 3 weken",
	"over 12 seconden",
	// nn
	"om 7 uker",
	"om 2 timer",
	// os
	"10 сахаты фӕстӕ",
	// pa-Guru
	"3 ਸਾਲਾਂ ਵਿੱਚ",
	"7 ਦਿਨਾਂ ਵਿੱਚ",
	// pa
	"8 ਘੰਟਿਆਂ ਵਿੱਚ",
	"16 ਸਕਿੰਟ ਵਿੱਚ",
	// pl
	"za 12 sekundy",
	"za 22 tygodnia",
	// pt
	"dentro de 11 minuto",
	"dentro de 8 meses",
	// ro
	"peste 12 de săptămâni",
	"peste 6 de ore",
	// sah
	"15 нэдиэлэннэн",
	"12 мүнүүтэннэн",
	// se
	"3 mánotbadji maŋŋilit",
	"10 sekunda maŋŋilit",
	// si
	"මිනිත්තු 10කින්",
	"දින 3න්",
	// sk
	"o 23 týždňov",
	// sl
	"čez 7 leto",
	"čez 8 minut 22 sek",
	// sq
	"pas 2 muajsh",
	"pas 15 ditësh",
	// sr-Cyrl
	"за 3 годину",
	"за 10 мин 20 сек",
	// sr-Latn
	"za 2 god 6 mes",
	"za 14 nedelja",
	// sr
	"за 18 недеља",
	"за 5 месеци",
	// sv
	"om 7 veckor",
	"om 10 timmar",
	// sw
	"baada ya saa 21",
	"baada ya sekunde 16",
	// ta
	"4 மாதங்களில்",
	"14 நாட்களில்",
	// te
	"3 వారాల్లో",
	"15 గంలో",
	// th
	"ในอีก 6 นาที",
	"ในอีก 3 ปี",
	// to
	"'i he māhina 'e 5",
	"'i he houa 'e 11",
	// tr
	"15 saniye sonra",
	"45 saat 234 dakika sonra",
	// uk
	"через 8 хвилини",
	"через 10 тижня",
	// uz-Cyrl
	"12 кундан сўнг",
	"10 дақиқадан сўнг",
	// uz-Latn
	"3 yildan keyin",
	"5 haftadan keyin",
	// uz
	"12 kundan keyin",
	"50 daqiqadan keyin",
	// vi
	"sau 5 năm nữa",
	"sau 2 phút nữa",
	// wae
	"i 3 stunde",
	"i 5 täg",
	// yue
	"3 個星期後",
	"6 年後",
	// zh-Hans
	"5个月后",
	"7天后",
	// zh-Hant
	"2 分鐘後",
	"4 週後",
}
