package godateparser

import (
	"fmt"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
	"github.com/stretchr/testify/assert"
)

var (
	relativeTestNow    = tt(2014, 9, 1, 10, 30)
	relativeTestConfig = Configuration{
		CurrentTime:        relativeTestNow,
		Parsers:            []ParserType{RelativeTime},
		ReturnTimeAsPeriod: true,
	}
)

func TestParser_Parse_relative_pastAndFutureDates(t *testing.T) {
	// Prepare structs
	type pfpDiff map[string]int
	type testScenario struct {
		String string
		Diff   pfpDiff
		Period date.Period
	}

	// Helper functions
	expectedDate := func(d pfpDiff) time.Time {
		return relativeTestNow.AddDate(d["year"], d["month"], d["day"]+d["week"]*7).
			Add(time.Duration(d["hour"]) * time.Hour).
			Add(time.Duration(d["minute"]) * time.Minute).
			Add(time.Duration(d["second"]) * time.Second)
	}

	// Prepare scenarios
	Day, Month, Year, Time := date.Day, date.Month, date.Year, date.Time
	pastTimes := []testScenario{
		// Mixed temporal nouns
		{"today", pfpDiff{"day": -0}, date.Day},
		{"Today", pfpDiff{"day": -0}, date.Day},
		{"TODAY", pfpDiff{"day": -0}, date.Day},
		{"Сегодня", pfpDiff{"day": -0}, date.Day},
		{"Hoje", pfpDiff{"day": -0}, date.Day},
		{"Oggi", pfpDiff{"day": -0}, date.Day},
		{"yesterday", pfpDiff{"day": -1}, date.Day},
		{" Yesterday \n", pfpDiff{"day": -1}, date.Day},
		{"Ontem", pfpDiff{"day": -1}, date.Day},
		{"Ieri", pfpDiff{"day": -1}, date.Day},
		{"the day before yesterday", pfpDiff{"day": -2}, date.Day},
		{"The DAY before Yesterday", pfpDiff{"day": -2}, date.Day},
		{"Anteontem", pfpDiff{"day": -2}, date.Day},
		{"Avant-hier", pfpDiff{"day": -2}, date.Day},
		{"вчера", pfpDiff{"day": -1}, date.Day},
		{"снощи", pfpDiff{"day": -1}, date.Day},

		// English dates
		{"yesterday", pfpDiff{"day": -1}, Day},
		{"yesterday at 11:30", pfpDiff{"hour": -23}, Time},
		{"1 decade", pfpDiff{"year": -10}, Year},
		{"1 decade 2 years", pfpDiff{"year": -12}, Year},
		{"1 decade 12 months", pfpDiff{"year": -10, "month": -12}, Month},
		{"1 decade and 11 months", pfpDiff{"year": -10, "month": -11}, Month},
		{"last decade", pfpDiff{"year": -10}, Year},
		{"a decade ago", pfpDiff{"year": -10}, Year},
		{"100 decades", pfpDiff{"year": -1000}, Year},
		{"yesterday", pfpDiff{"day": -1}, Day},
		{"the day before yesterday", pfpDiff{"day": -2}, Day},
		{"today", pfpDiff{"day": -0}, Day},
		{"an hour ago", pfpDiff{"hour": -1}, Day},
		{"about an hour ago", pfpDiff{"hour": -1}, Day},
		{"a day ago", pfpDiff{"day": -1}, Day},
		{"a week ago", pfpDiff{"week": -1}, Day},
		{"2 hours ago", pfpDiff{"hour": -2}, Day},
		{"about 23 hours ago", pfpDiff{"hour": -23}, Day},
		{"1 year 2 months", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 year, 09 months,01 weeks", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 year 11 months", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 year 12 months", pfpDiff{"year": -1, "month": -12}, Month},
		{"15 hr", pfpDiff{"hour": -15}, Day},
		{"15 hrs", pfpDiff{"hour": -15}, Day},
		{"2 min", pfpDiff{"minute": -2}, Day},
		{"2 mins", pfpDiff{"minute": -2}, Day},
		{"3 sec", pfpDiff{"second": -3}, Day},
		{"1000 years ago", pfpDiff{"year": -1000}, Year},
		{"2013 years ago", pfpDiff{"year": -2013}, Year},
		{"5000 months ago", pfpDiff{"year": -416, "month": -8}, Month},
		{fmt.Sprintf("%d months ago", 2013*12+8), pfpDiff{"year": -2013, "month": -8}, Month},
		{"1 year, 1 month, 1 week, 1 day, 1 hour and 1 minute ago", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},
		{"just now", pfpDiff{"second": -0}, Day},
		// Fix for #291, work till one to twelve only
		{"nine hours ago", pfpDiff{"hour": -9}, Day},
		{"three week ago", pfpDiff{"week": -3}, Day},
		{"eight months ago", pfpDiff{"month": -8}, Month},
		{"six days ago", pfpDiff{"day": -6}, Day},
		{"five years ago", pfpDiff{"year": -5}, Year},

		// French dates
		{"Aujourd'hui", nil, Day},
		{"Aujourd’hui", nil, Day},
		{"Aujourdʼhui", nil, Day},
		{"Aujourdʻhui", nil, Day},
		{"Aujourd՚hui", nil, Day},
		{"Aujourdꞌhui", nil, Day},
		{"Aujourd＇hui", nil, Day},
		{"Aujourd′hui", nil, Day},
		{"Aujourd‵hui", nil, Day},
		{"Aujourdʹhui", nil, Day},
		{"Aujourd＇hui", nil, Day},
		{"moins de 21s", pfpDiff{"second": -21}, Day},
		{"moins de 21m", pfpDiff{"minute": -21}, Day},
		{"moins de 21h", pfpDiff{"hour": -21}, Day},
		{"moins de 21 minute", pfpDiff{"minute": -21}, Day},
		{"moins de 21 heure", pfpDiff{"hour": -21}, Day},
		{"Hier", pfpDiff{"day": -1}, Day},
		{"Avant-hier", pfpDiff{"day": -2}, Day},
		{"Il ya un jour", pfpDiff{"day": -1}, Day},
		{"Il ya une heure", pfpDiff{"hour": -1}, Day},
		{"Il ya 2 heures", pfpDiff{"hour": -2}, Day},
		{"Il ya environ 23 heures", pfpDiff{"hour": -23}, Day},
		{"1 an 2 mois", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 année, 09 mois, 01 semaines", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 an 11 mois", pfpDiff{"year": -1, "month": -11}, Month},
		{"Il ya 1 an, 1 mois, 1 semaine, 1 jour, 1 heure et 1 minute", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},
		{"Il y a 40 min", pfpDiff{"minute": -40}, Day},

		// German dates
		{"Heute", pfpDiff{"day": -0}, Day},
		{"Gestern", pfpDiff{"day": -1}, Day},
		{"vorgestern", pfpDiff{"day": -2}, Day},
		{"vor einem Tag", pfpDiff{"day": -1}, Day},
		{"vor einer Stunden", pfpDiff{"hour": -1}, Day},
		{"Vor 2 Stunden", pfpDiff{"hour": -2}, Day},
		{"vor etwa 23 Stunden", pfpDiff{"hour": -23}, Day},
		{"1 Jahr 2 Monate", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 Jahr, 09 Monate, 01 Wochen", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 Jahr 11 Monate", pfpDiff{"year": -1, "month": -11}, Month},
		{"vor 29h", pfpDiff{"hour": -29}, Day},
		{"vor 29m", pfpDiff{"minute": -29}, Day},
		{"1 Jahr, 1 Monat, 1 Woche, 1 Tag, 1 Stunde und 1 Minute", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Italian dates
		{"oggi", pfpDiff{"day": -0}, Day},
		{"ieri", pfpDiff{"day": -1}, Day},
		{"2 ore fa", pfpDiff{"hour": -2}, Day},
		{"circa 23 ore fa", pfpDiff{"hour": -23}, Day},
		{"1 anno 2 mesi", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 anno, 09 mesi, 01 settimane", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 anno 11 mesi", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 anno, 1 mese, 1 settimana, 1 giorno, 1 ora e 1 minuto fa", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Portuguese dates
		{"ontem", pfpDiff{"day": -1}, Day},
		{"anteontem", pfpDiff{"day": -2}, Day},
		{"hoje", pfpDiff{"day": -0}, Day},
		{"uma hora atrás", pfpDiff{"hour": -1}, Day},
		{"1 segundo atrás", pfpDiff{"second": -1}, Day},
		{"um dia atrás", pfpDiff{"day": -1}, Day},
		{"uma semana atrás", pfpDiff{"week": -1}, Day},
		{"2 horas atrás", pfpDiff{"hour": -2}, Day},
		{"cerca de 23 horas atrás", pfpDiff{"hour": -23}, Day},
		{"1 ano 2 meses", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 ano, 09 meses, 01 semanas", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 ano 11 meses", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 ano, 1 mês, 1 semana, 1 dia, 1 hora e 1 minuto atrás", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Turkish dates
		{"Dün", pfpDiff{"day": -1}, Day},
		{"Bugün", pfpDiff{"day": -0}, Day},
		{"2 saat önce", pfpDiff{"hour": -2}, Day},
		{"yaklaşık 23 saat önce", pfpDiff{"hour": -23}, Day},
		{"1 yıl 2 ay", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 yıl, 09 ay, 01 hafta", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 yıl 11 ay", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 yıl, 1 ay, 1 hafta, 1 gün, 1 saat ve 1 dakika önce", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Russian dates
		{"сегодня", pfpDiff{"day": -0}, Day},
		{"Вчера в", pfpDiff{"day": -1}, Day},
		{"вчера", pfpDiff{"day": -1}, Day},
		{"2 часа назад", pfpDiff{"hour": -2}, Day},
		{"час назад", pfpDiff{"hour": -1}, Day},
		{"минуту назад", pfpDiff{"minute": -1}, Day},
		{"2 ч. 21 мин. назад", pfpDiff{"hour": -2, "minute": -21}, Day},
		{"около 23 часов назад", pfpDiff{"hour": -23}, Day},
		{"1 год 2 месяца", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 год, 09 месяцев, 01 недель", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 год 11 месяцев", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 год, 1 месяц, 1 неделя, 1 день, 1 час и 1 минуту назад", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Czech dates
		{"Dnes", pfpDiff{"day": -0}, Day},
		{"Včera", pfpDiff{"day": -1}, Day},
		{"Předevčírem", pfpDiff{"day": -2}, Day},
		{"Před 2 hodinami", pfpDiff{"hour": -2}, Day},
		{"před přibližně 23 hodin", pfpDiff{"hour": -23}, Day},
		{"1 rok 2 měsíce", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 rok, 09 měsíců, 01 týdnů", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 rok 11 měsíců", pfpDiff{"year": -1, "month": -11}, Month},
		{"3 dny", pfpDiff{"day": -3}, Day},
		{"3 hodiny", pfpDiff{"hour": -3}, Day},
		{"2 roky, 2 týdny, 1 den, 1 hodinu, 5 vteřin před", pfpDiff{"year": -2, "week": -2, "day": -1, "hour": -1, "second": -5}, Day},
		{"1 rok, 1 měsíc, 1 týden, 1 den, 1 hodina, 1 minuta před", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Spanish dates
		{"anteayer", pfpDiff{"day": -2}, Day},
		{"ayer", pfpDiff{"day": -1}, Day},
		{"hoy", pfpDiff{"day": -0}, Day},
		{"hace una hora", pfpDiff{"hour": -1}, Day},
		{"Hace un día", pfpDiff{"day": -1}, Day},
		{"Hace una semana", pfpDiff{"week": -1}, Day},
		{"Hace 2 horas", pfpDiff{"hour": -2}, Day},
		{"Hace cerca de 23 horas", pfpDiff{"hour": -23}, Day},
		{"1 año 2 meses", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 año, 09 meses, 01 semanas", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 año 11 meses", pfpDiff{"year": -1, "month": -11}, Month},
		{"Hace 1 año, 1 mes, 1 semana, 1 día, 1 hora y 1 minuto", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Chinese dates
		{"昨天", pfpDiff{"day": -1}, Day},
		{"前天", pfpDiff{"day": -2}, Day},
		{"2小时前", pfpDiff{"hour": -2}, Day},
		{"约23小时前", pfpDiff{"hour": -23}, Day},
		{"1年2个月", pfpDiff{"year": -1, "month": -2}, Month},
		{"1年2個月", pfpDiff{"year": -1, "month": -2}, Month},
		{"1年11个月", pfpDiff{"year": -1, "month": -11}, Month},
		{"1年11個月", pfpDiff{"year": -1, "month": -11}, Month},
		{"1年，1月，1周，1天，1小时，1分钟前", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Arabic dates
		{"اليوم", nil, Day},
		{"يوم أمس", pfpDiff{"day": -1}, Day},
		{"منذ يومين", pfpDiff{"day": -2}, Day},
		{"منذ 3 أيام", pfpDiff{"day": -3}, Day},
		{"منذ 21 أيام", pfpDiff{"day": -21}, Day},
		{"1 عام, 1 شهر, 1 أسبوع, 1 يوم, 1 ساعة, 1 دقيقة", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Thai dates
		{"วันนี้", nil, Day},
		{"เมื่อวานนี้", pfpDiff{"day": -1}, Day},
		{"2 วัน", pfpDiff{"day": -2}, Day},
		{"2 ชั่วโมง", pfpDiff{"hour": -2}, Day},
		{"23 ชม.", pfpDiff{"hour": -23}, Day},
		{"2 สัปดาห์ 3 วัน", pfpDiff{"week": -2, "day": -3}, Day},
		{"1 ปี 9 เดือน 1 สัปดาห์", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 ปี 1 เดือน 1 สัปดาห์ 1 วัน 1 ชั่วโมง 1 นาที", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Vietnamese dates
		{"Hôm nay", nil, Day},
		{"Hôm qua", pfpDiff{"day": -1}, Day},
		{"2 giờ", pfpDiff{"hour": -2}, Day},
		{"2 tuần 3 ngày", pfpDiff{"week": -2, "day": -3}, Day},
		// Following test unsupported, refer to discussion at:
		// http://github.com/scrapinghub/dateparser/issues/33
		// {"1 năm 1 tháng 1 tuần 1 ngày 1 giờ 1 chút", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},

		// Belarusian dates
		{"сёння", nil, Day},
		{"учора ў", pfpDiff{"day": -1}, Day},
		{"ўчора", pfpDiff{"day": -1}, Day},
		{"пазаўчора", pfpDiff{"day": -2}, Day},
		{"2 гадзіны таму назад", pfpDiff{"hour": -2}, Day},
		{"2 гадзіны таму", pfpDiff{"hour": -2}, Day},
		{"гадзіну назад", pfpDiff{"hour": -1}, Day},
		{"хвіліну таму", pfpDiff{"minute": -1}, Day},
		{"2 гадзіны 21 хвіл. назад", pfpDiff{"hour": -2, "minute": -21}, Day},
		{"каля 23 гадзін назад", pfpDiff{"hour": -23}, Day},
		{"1 год 2 месяцы", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 год, 09 месяцаў, 01 тыдзень", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"2 гады 3 месяцы", pfpDiff{"year": -2, "month": -3}, Month},
		{"5 гадоў, 1 месяц, 6 тыдняў, 3 дні, 5 гадзін 1 хвіліну і 3 секунды таму назад", pfpDiff{"year": -5, "month": -1, "week": -6, "day": -3, "hour": -5, "minute": -1, "second": -3}, Day},

		// Polish dates
		{"wczoraj", pfpDiff{"day": -1}, Day},
		{"1 godz. 2 minuty temu", pfpDiff{"hour": -1, "minute": -2}, Day},
		{"2 lata, 3 miesiące, 1 tydzień, 2 dni, 4 godziny, 15 minut i 25 sekund temu", pfpDiff{"year": -2, "month": -3, "week": -1, "day": -2, "hour": -4, "minute": -15, "second": -25}, Day},
		{"2 minuty temu", pfpDiff{"minute": -2}, Day},
		{"15 minut temu", pfpDiff{"minute": -15}, Day},

		// Bulgarian dates
		{"преди 3 дни", pfpDiff{"day": -3}, Day},
		{"преди час", pfpDiff{"hour": -1}, Day},
		{"преди година", pfpDiff{"year": -1}, Year},
		{"вчера", pfpDiff{"day": -1}, Day},
		{"онзи ден", pfpDiff{"day": -2}, Day},
		{"днес", nil, Day},
		{"преди час", pfpDiff{"hour": -1}, Day},
		{"преди един ден", pfpDiff{"day": -1}, Day},
		{"преди седмица", pfpDiff{"week": -1}, Day},
		{"преди 2 часа", pfpDiff{"hour": -2}, Day},
		{"преди около 23 часа", pfpDiff{"hour": -23}, Day},

		// Bangla dates
		// {"গতকাল", pfpDiff{"day": -1}, Day},
		// {"আজ", nil, Day},
		{"1 ঘন্টা আগে", pfpDiff{"hour": -1}, Day},
		{"প্রায় 1 ঘন্টা আগে", pfpDiff{"hour": -1}, Day},
		{"1 দিন আগে", pfpDiff{"day": -1}, Day},
		{"1 সপ্তাহ আগে", pfpDiff{"week": -1}, Day},
		{"2 ঘন্টা আগে", pfpDiff{"hour": -2}, Day},
		{"প্রায় 23 ঘন্টা আগে", pfpDiff{"hour": -23}, Day},
		{"1 বছর 2 মাস", pfpDiff{"year": -1, "month": -2}, Month},
		{"1 বছর, 09 মাস,01 সপ্তাহ", pfpDiff{"year": -1, "month": -9, "week": -1}, Day},
		{"1 বছর 11 মাস", pfpDiff{"year": -1, "month": -11}, Month},
		{"1 বছর 12 মাস", pfpDiff{"year": -1, "month": -12}, Month},
		{"15 ঘন্টা", pfpDiff{"hour": -15}, Day},
		{"2 মিনিট", pfpDiff{"minute": -2}, Day},
		{"3 সেকেন্ড", pfpDiff{"second": -3}, Day},
		{"1000 বছর আগে", pfpDiff{"year": -1000}, Year},
		{"5000 মাস আগে", pfpDiff{"year": -416, "month": -8}, Month},
		{fmt.Sprintf("%d মাস আগে", 2013*12+8), pfpDiff{"year": -2013, "month": -8}, Month},
		{"1 বছর, 1 মাস, 1 সপ্তাহ, 1 দিন, 1 ঘন্টা এবং 1 মিনিট আগে", pfpDiff{"year": -1, "month": -1, "week": -1, "day": -1, "hour": -1, "minute": -1}, Day},
		// {"এখন", nil, Day},

		// Hindi dates
		{"1 घंटे पहले", pfpDiff{"hour": -1}, Day},
		{"15 मिनट पहले", pfpDiff{"minute": -15}, Day},
		{"25 सेकंड पूर्व", pfpDiff{"second": -25}, Day},
		{"1 वर्ष, 8 महीने, 2 सप्ताह", pfpDiff{"year": -1, "month": -8, "week": -2}, Day},
		{"1 वर्ष 7 महीने", pfpDiff{"year": -1, "month": -7}, Month},
		{"आज", nil, Day},

		// af
		{"2 uur gelede", pfpDiff{"hour": -2}, Day},
		{"verlede maand", pfpDiff{"month": -1}, Month},
		// agq
		{"ā zūɛɛ", pfpDiff{"day": -1}, Day},
		// ak
		{"ndeda", pfpDiff{"day": -1}, Day},
		// am
		{"ከ8 ወራት በፊት", pfpDiff{"month": -8}, Month},
		{"ያለፈው ሳምንት", pfpDiff{"week": -1}, Day},
		// as
		{"কালি", pfpDiff{"day": -1}, Day},
		{"আজি", nil, Day},
		// asa
		{"ighuo", pfpDiff{"day": -1}, Day},
		// ast
		{"hai 2 selmanes hai 3 díes", pfpDiff{"week": -2, "day": -3}, Day},
		{"l'añu pas el mes pasáu", pfpDiff{"year": -1, "month": -1}, Month},
		// az-Latn
		{"1 il öncə 2 ay öncə 3 həftə öncə", pfpDiff{"year": -1, "month": -2, "week": -3}, Day},
		{"6 saat öncə 5 dəqiqə öncə 4 saniyə öncə", pfpDiff{"hour": -6, "minute": -5, "second": -4}, Day},
		// az
		{"2 gün öncə 23 saat öncə", pfpDiff{"day": -2, "hour": -23}, Day},
		{"5 dəqiqə öncə 27 saniyə öncə", pfpDiff{"minute": -5, "second": -27}, Day},
		// be
		{"2 гадзіны таму 10 хвіліны таму", pfpDiff{"hour": -2, "minute": -10}, Day},
		// bg
		{"преди 3 месеца преди 2 седм", pfpDiff{"month": -3, "week": -2}, Day},
		// bn
		{"8 মিনিট আগে 42 সেকেন্ড পূর্বে", pfpDiff{"minute": -8, "second": -42}, Day},
		// br
		{"7 eur zo 15 min zo 25 s zo", pfpDiff{"hour": -7, "minute": -15, "second": -25}, Day},
		{"14 sizhun zo 3 deiz zo", pfpDiff{"week": -14, "day": -3}, Day},
		// bs-Cyrl
		{"пре 5 година пре 7 месеци", pfpDiff{"year": -5, "month": -7}, Month},
		{"пре 5 сати пре 25 секунди", pfpDiff{"hour": -5, "second": -25}, Day},
		// bs-Latn
		{"prije 20 sat 5 minuta", pfpDiff{"hour": -20, "minute": -5}, Day},
		{"prije 13 godina prije 3 sed", pfpDiff{"year": -13, "week": -3}, Day},
		// bs
		{"prije 3 mjeseci prije 12 sati", pfpDiff{"month": -3, "hour": -12}, Month},
		{"prije 3 god 4 mj 5 sed 7 dan", pfpDiff{"year": -3, "month": -4, "week": -5, "day": -7}, Day},
		// ca
		{"fa 4 setmanes fa 5 segon", pfpDiff{"week": -4, "second": -5}, Day},
		{"fa 1 hora 2 minut 3 segon", pfpDiff{"hour": -1, "minute": -2, "second": -3}, Day},
		// ce
		{"10 кӏир хьалха 3 д хьалха", pfpDiff{"week": -10, "day": -3}, Day},
		{"12 сахь 30 мин 30 сек хьалха", pfpDiff{"hour": -12, "minute": -30, "second": -30}, Day},
		// chr
		{"ᎾᎿ 10 ᏒᎾᏙᏓᏆᏍᏗ ᏥᎨᏒ 5 ᎢᎦ ᏥᎨᏒ", pfpDiff{"week": -10, "day": -5}, Day},
		// cs
		{"před 3 rok 4 měsíc 5 den", pfpDiff{"year": -3, "month": -4, "day": -5}, Day},
		{"před 2 minutou před 45 sekundou", pfpDiff{"minute": -2, "second": -45}, Day},
		// cy
		{"5 wythnos yn ôl", pfpDiff{"week": -5}, Day},
		{"7 munud 8 eiliad yn ôl", pfpDiff{"minute": -7, "second": -8}, Day},
		// dsb
		{"pśed 15 góźinu 10 minuta 5 sekunda", pfpDiff{"hour": -15, "minute": -10, "second": -5}, Day},
		{"pśed 5 lětom, pśed 7 mjasecom", pfpDiff{"year": -5, "month": -7}, Month},
		// ee
		{"ŋkeke 12 si wo va yi", pfpDiff{"day": -12}, Day},
		{"ƒe 6 si va yi ɣleti 5 si va yi", pfpDiff{"year": -6, "month": -5}, Month},
		// el
		{"πριν από 5 ώρα 6 λεπτό 7 δευτερόλεπτο", pfpDiff{"hour": -5, "minute": -6, "second": -7}, Day},
		{"προηγούμενος μήνας", pfpDiff{"month": -1}, Month},
		// es
		{"hace 5 hora 2 minuto 3 segundo", pfpDiff{"hour": -5, "minute": -2, "second": -3}, Day},
		// et
		{"5 minut 12 sekundi eest", pfpDiff{"minute": -5, "second": -12}, Day},
		{"11 aasta 11 kuu eest", pfpDiff{"year": -11, "month": -11}, Month},
		// eu
		{"duela 3 minutu", pfpDiff{"minute": -3}, Day},
		// fil
		{"10 oras ang nakalipas", pfpDiff{"hour": -10}, Day},
		// fo
		{"3 tími 12 minutt síðan", pfpDiff{"hour": -3, "minute": -12}, Day},
		// fur
		{"10 setemane 12 zornade indaûr", pfpDiff{"week": -10, "day": -12}, Day},
		// fy
		{"6 moannen lyn", pfpDiff{"month": -6}, Month},
		// ga
		{"12 uair an chloig ó shin", pfpDiff{"hour": -12}, Day},
		// gd
		{"15 mhionaid air ais", pfpDiff{"minute": -15}, Day},
		// gl
		{"hai 5 ano 7 mes", pfpDiff{"year": -5, "month": -7}, Month},
		// gu
		{"5 કલાક પહેલાં", pfpDiff{"hour": -5}, Day},
		// hr
		{"prije 3 tjedana", pfpDiff{"week": -3}, Day},
		// hsb
		{"před 12 lětom 15 měsac", pfpDiff{"year": -12, "month": -15}, Month},
		// hy
		{"15 րոպե առաջ", pfpDiff{"minute": -15}, Day},
		// is
		{"fyrir 3 ári fyrir 2 mánuði", pfpDiff{"year": -3, "month": -2}, Month},
		// it
		{"5 settimana fa", pfpDiff{"week": -5}, Day},
		// jgo
		{"ɛ́ gɛ́ mɔ́ pɛsaŋ 3", pfpDiff{"month": -3}, Month},
		// ka
		{"4 წლის წინ", pfpDiff{"year": -4}, Year},
		// kk
		{"5 сағат бұрын", pfpDiff{"hour": -5}, Day},
		// kl
		{"for 6 ulloq unnuarlu siden", pfpDiff{"day": -6}, Day},
		// km
		{"11 សប្ដាហ៍​មុន", pfpDiff{"week": -11}, Day},
		// kn
		{"15 ಸೆಕೆಂಡುಗಳ ಹಿಂದೆ", pfpDiff{"second": -15}, Day},
		// ko
		{"12개월 전", pfpDiff{"month": -12}, Month},
		// ksh
		{"vör 15 johre", pfpDiff{"year": -15}, Year},
		// ky
		{"18 секунд мурун", pfpDiff{"second": -18}, Day},
		// lb
		{"virun 15 stonn", pfpDiff{"hour": -15}, Day},
		// lkt
		{"hékta 8-čháŋ k'uŋ héhaŋ", pfpDiff{"day": -8}, Day},
		// lt
		{"prieš 20 minučių", pfpDiff{"minute": -20}, Day},
		// lv
		{"pirms 10 gadiem", pfpDiff{"year": -10}, Year},
		// mk
		{"пред 13 часа", pfpDiff{"hour": -13}, Day},
		// ml
		{"3 ആഴ്ച മുമ്പ്", pfpDiff{"week": -3}, Day},
		// mn
		{"15 секундын өмнө", pfpDiff{"second": -15}, Day},
		// mr
		{"25 वर्षापूर्वी", pfpDiff{"year": -25}, Year},
		// ms
		{"10 minit lalu", pfpDiff{"minute": -10}, Day},
		// my
		{"ပြီးခဲ့သည့် 15 နှစ်", pfpDiff{"year": -15}, Year},
		// nb
		{"for 12 måneder siden", pfpDiff{"month": -12}, Month},
		// ne
		{"23 हप्ता पहिले", pfpDiff{"week": -23}, Day},
		// nl
		{"32 minuten geleden", pfpDiff{"minute": -32}, Day},
		// nn
		{"for 15 sekunder siden", pfpDiff{"second": -15}, Day},
		// os
		{"35 сахаты размӕ", pfpDiff{"hour": -35}, Day},
		// pa-Guru
		{"23 ਹਫ਼ਤੇ ਪਹਿਲਾਂ", pfpDiff{"week": -23}, Day},
		// pa
		{"7 ਸਾਲ ਪਹਿਲਾਂ", pfpDiff{"year": -7}, Year},
		// ro
		{"acum 56 de secunde", pfpDiff{"second": -56}, Day},
		// sah
		{"2 нэдиэлэ анараа өттүгэр", pfpDiff{"week": -2}, Day},
		// se
		{"8 jahkki árat", pfpDiff{"year": -8}, Year},
		// si
		{"මිනිත්තු 6කට පෙර", pfpDiff{"minute": -6}, Day},
		// sk
		{"pred 20 hodinami 45 min", pfpDiff{"hour": -20, "minute": -45}, Day},
		// sl
		{"pred 15 tednom 10 dan", pfpDiff{"week": -15, "day": -10}, Day},
		// sq
		{"11 minutë më parë", pfpDiff{"minute": -11}, Day},
		// sr-Cyrl
		{"пре 8 године 3 месец", pfpDiff{"year": -8, "month": -3}, Month},
		// sr-Latn
		{"pre 2 nedelje", pfpDiff{"week": -2}, Day},
		// sr
		{"пре 59 секунди", pfpDiff{"second": -59}, Day},
		// sw
		{"mwezi 2 uliopita", pfpDiff{"month": -2}, Month},
		// ta
		{"41 நிமிடங்களுக்கு முன்", pfpDiff{"minute": -41}, Day},
		// te
		{"36 వారాల క్రితం", pfpDiff{"week": -36}, Day},
		// to
		{"houa 'e 7 kuo'osi", pfpDiff{"hour": -7}, Day},
		// tr
		{"32 dakika önce", pfpDiff{"minute": -32}, Day},
		// uk
		{"3 року тому", pfpDiff{"year": -3}, Year},
		// uz-Cyrl
		{"10 ҳафта олдин", pfpDiff{"week": -10}, Day},
		// uz-Latn
		{"3 oy oldin", pfpDiff{"month": -3}, Month},
		// uz
		{"45 soniya oldin", pfpDiff{"second": -45}, Day},
		// vi
		{"23 ngày trước", pfpDiff{"day": -23}, Day},
		// wae
		{"vor 15 stunde", pfpDiff{"hour": -15}, Day},
		// yue
		{"5 分鐘前", pfpDiff{"minute": -5}, Day},
		// zh-Hans
		{"3周前", pfpDiff{"week": -3}, Day},
		// zh-Hant
		{"2 年前", pfpDiff{"year": -2}, Year},
		// zu
		{"21 izinsuku ezedlule", pfpDiff{"day": -21}, Day},
	}

	futureTimes := []testScenario{
		// English dates
		{"in 1 decade 2 months", pfpDiff{"year": 10, "month": 2}, Month},
		{"in 100 decades", pfpDiff{"year": 1000}, Year},
		{"in 1 decade 12 years", pfpDiff{"year": 22}, Year},
		{"next decade", pfpDiff{"year": 10}, Year},
		{"in a decade", pfpDiff{"year": 10}, Year},
		{"tomorrow", pfpDiff{"day": 1}, Day},
		{"today", pfpDiff{"day": 0}, Day},
		{"in an hour", pfpDiff{"hour": 1}, Day},
		{"in about an hour", pfpDiff{"hour": 1}, Day},
		{"in 1 day", pfpDiff{"day": 1}, Day},
		{"in a week", pfpDiff{"week": 1}, Day},
		{"in 2 hours", pfpDiff{"hour": 2}, Day},
		{"in about 23 hours", pfpDiff{"hour": 23}, Day},
		{"in 1 year 2 months", pfpDiff{"year": 1, "month": 2}, Month},
		{"in 1 year, 09 months,01 weeks", pfpDiff{"year": 1, "month": 9, "week": 1}, Day},
		{"in 1 year 11 months", pfpDiff{"year": 1, "month": 11}, Month},
		{"in 1 year 12 months", pfpDiff{"year": 1, "month": 12}, Month},
		{"in 15 hr", pfpDiff{"hour": 15}, Day},
		{"in 15 hrs", pfpDiff{"hour": 15}, Day},
		{"in 2 min", pfpDiff{"minute": 2}, Day},
		{"in 2 mins", pfpDiff{"minute": 2}, Day},
		{"in 3 sec", pfpDiff{"second": 3}, Day},
		{"in 1000 years", pfpDiff{"year": 1000}, Year},
		{"in 5000 months", pfpDiff{"year": 416, "month": 8}, Month},
		{fmt.Sprintf("in %d months", 2013*12+8), pfpDiff{"year": 2013, "month": 8}, Month},
		{"in 1 year, 1 month, 1 week, 1 day, 1 hour and 1 minute", pfpDiff{"year": 1, "month": 1, "week": 1, "day": 1, "hour": 1, "minute": 1}, Day},
		{"just now", pfpDiff{"second": 0}, Day},

		// French dates
		{"Aujourd'hui", pfpDiff{"day": 0}, Day},
		{"Dans un jour", pfpDiff{"day": 1}, Day},
		{"Dans une heure", pfpDiff{"hour": 1}, Day},
		{"En 2 heures", pfpDiff{"hour": 2}, Day},
		{"Dans environ 23 heures", pfpDiff{"hour": 23}, Day},
		{"Dans 1 an 2 mois", pfpDiff{"year": 1, "month": 2}, Month},
		{"En 1 année, 09 mois, 01 semaines", pfpDiff{"year": 1, "month": 9, "week": 1}, Day},
		{"Dans 1 an 11 mois", pfpDiff{"year": 1, "month": 11}, Month},
		{"Dans 1 année, 1 mois, 1 semaine, 1 jour, 1 heure et 1 minute", pfpDiff{"year": 1, "month": 1, "week": 1, "day": 1, "hour": 1, "minute": 1}, Day},
		{"Dans 40 min", pfpDiff{"minute": 40}, Day},

		// German dates
		{"Heute", pfpDiff{"day": 0}, Day},
		{"Morgen", pfpDiff{"day": 1}, Day},
		{"in einem Tag", pfpDiff{"day": 1}, Day},
		{"in einer Stunde", pfpDiff{"hour": 1}, Day},
		{"in 2 Stunden", pfpDiff{"hour": 2}, Day},
		{"in etwa 23 Stunden", pfpDiff{"hour": 23}, Day},
		{"im 1 Jahr 2 Monate", pfpDiff{"year": 1, "month": 2}, Month},
		{"im 1 Jahr, 09 Monate, 01 Wochen", pfpDiff{"year": 1, "month": 9, "week": 1}, Day},
		{"im 1 Jahr 11 Monate", pfpDiff{"year": 1, "month": 11}, Month},
		{"im 1 Jahr, 1 Monat, 1 Woche, 1 Tag, 1 Stunde und 1 Minute", pfpDiff{"year": 1, "month": 1, "week": 1, "day": 1, "hour": 1, "minute": 1}, Day},

		// Polish dates
		{"jutro", pfpDiff{"day": 1}, Day},
		{"pojutrze", pfpDiff{"day": 2}, Day},
		{"za 2 lata, 3 miesiące, 1 tydzień, 2 dni, 4 godziny, 15 minut i 25 sekund", pfpDiff{"year": 2, "month": 3, "week": 1, "day": 2, "hour": 4, "minute": 15, "second": 25}, Day},
		{"za 2 minuty", pfpDiff{"minute": 2}, Day},
		{"za 15 minut", pfpDiff{"minute": 15}, Day},

		// Turkish dates
		{"yarın", pfpDiff{"day": 1}, Day},
		{"2 gün içerisinde", pfpDiff{"day": 2}, Day},
		{"4 ay içerisinde", pfpDiff{"month": 4}, Month},
		{"3 gün sonra", pfpDiff{"day": 3}, Day},
		{"2 ay sonra", pfpDiff{"month": 2}, Month},
		{"5 yıl 3 gün sonra", pfpDiff{"year": 5, "day": 3}, Day},
		{"5 gün içinde", pfpDiff{"day": 5}, Day},
		{"6 ay içinde", pfpDiff{"month": 6}, Month},
		{"5 yıl içinde", pfpDiff{"year": 5}, Year},
		{"5 sene içinde", pfpDiff{"year": 5}, Year},
		{"haftaya", pfpDiff{"week": 1}, Day},
		{"gelecek hafta", pfpDiff{"week": 1}, Day},
		{"gelecek ay", pfpDiff{"month": 1}, Month},
		{"gelecek yıl", pfpDiff{"year": 1}, Year},

		// Hindi dates
		{"1 वर्ष 10 महीने में", pfpDiff{"year": 1, "month": 10}, Month},
		{"15 घंटे बाद", pfpDiff{"hour": 15}, Day},
		{"2 मिनट में", pfpDiff{"minute": 2}, Day},
		{"17 सेकंड बाद", pfpDiff{"second": 17}, Day},
		{"1 वर्ष, 5 महीने, 1 सप्ताह में", pfpDiff{"year": 1, "month": 5, "week": 1}, Day},

		// af
		{"oor 10 jaar", pfpDiff{"year": 10}, Year},
		{"oor 5 min 3 sek", pfpDiff{"minute": 5, "second": 3}, Day},
		// am
		{"በ2 ሳምንታት ውስጥ", pfpDiff{"week": 2}, Day},
		{"በ16 ቀናት ውስጥ", pfpDiff{"day": 16}, Day},
		// ast
		{"en 15 años", pfpDiff{"year": 15}, Year},
		{"en 20 minutos", pfpDiff{"minute": 20}, Day},
		// az-Latn
		{"5 saniyə ərzində", pfpDiff{"second": 5}, Day},
		{"10 saat 20 dəqiqə ərzində", pfpDiff{"hour": 10, "minute": 20}, Day},
		// az
		{"15 il 6 ay ərzində", pfpDiff{"year": 15, "month": 6}, Month},
		// be
		{"праз 5 гадзіны 6 хвіліны", pfpDiff{"hour": 5, "minute": 6}, Day},
		// bg
		{"след 12 мин 18 сек", pfpDiff{"minute": 12, "second": 18}, Day},
		// bn
		{"10 সেকেন্ডে", pfpDiff{"second": 10}, Day},
		// br
		{"a-benn 20 vloaz", pfpDiff{"year": 20}, Year},
		{"a-benn 15 deiz 20 eur", pfpDiff{"day": 15, "hour": 20}, Day},
		// bs-Cyrl
		{"за 5 минут 10 секунд", pfpDiff{"minute": 5, "second": 10}, Day},
		{"за 10 годину 11 месец", pfpDiff{"year": 10, "month": 11}, Month},
		// bs-Latn
		{"za 7 mjeseci", pfpDiff{"month": 7}, Month},
		{"za 6 dan 23 sat", pfpDiff{"day": 6, "hour": 23}, Day},
		// bs
		{"za 15 sedmica", pfpDiff{"week": 15}, Day},
		// ca
		{"d'aquí a 10 anys", pfpDiff{"year": 10}, Year},
		{"d'aquí a 15 minut 53 segon", pfpDiff{"minute": 15, "second": 53}, Day},
		// ce
		{"20 кӏира даьлча", pfpDiff{"week": 20}, Day},
		{"10 минот 25 секунд яьлча", pfpDiff{"minute": 10, "second": 25}, Day},
		// chr
		{"ᎾᎿ 10 ᎧᎸᎢ", pfpDiff{"month": 10}, Month},
		{"ᎾᎿ 24 ᎢᏳᏟᎶᏓ", pfpDiff{"hour": 24}, Day},
		// cs
		{"za 12 rok", pfpDiff{"year": 12}, Year},
		{"za 10 den 5 hodin", pfpDiff{"day": 10, "hour": 5}, Day},
		// cy
		{"ymhen 15 mis", pfpDiff{"month": 15}, Month},
		{"ymhen 10 munud 8 eiliad", pfpDiff{"minute": 10, "second": 8}, Day},
		// da
		{"om 10 minut 54 sekund", pfpDiff{"minute": 10, "second": 54}, Day},
		// de
		{"in 15 jahren 10 monat", pfpDiff{"year": 15, "month": 10}, Month},
		// dsb
		{"za 10 mjasec", pfpDiff{"month": 10}, Month},
		{"za 30 min 50 sek", pfpDiff{"minute": 30, "second": 50}, Day},
		// dz
		{"ལོ་འཁོར་ 4 ནང་", pfpDiff{"year": 4}, Year},
		{"སྐར་ཆ་ 20 ནང་", pfpDiff{"second": 20}, Day},
		// ee
		{"le ƒe 15 si gbɔna me", pfpDiff{"year": 15}, Year},
		{"le ŋkeke 2 wo me", pfpDiff{"day": 2}, Day},
		// el
		{"σε 5 ώρες", pfpDiff{"hour": 5}, Day},
		{"σε 4 λεπτό 45 δευτ", pfpDiff{"minute": 4, "second": 45}, Day},
		// et
		{"5 aasta 10 kuu pärast", pfpDiff{"year": 5, "month": 10}, Month},
		{"10 nädala pärast", pfpDiff{"week": 10}, Day},
		// eu
		{"15 hilabete barru", pfpDiff{"month": 15}, Month},
		{"20 egun barru", pfpDiff{"day": 20}, Day},
		// fil
		{"sa 8 segundo", pfpDiff{"second": 8}, Day},
		{"sa 2 oras 24 min", pfpDiff{"hour": 2, "minute": 24}, Day},
		// fo
		{"um 12 mánaðir", pfpDiff{"month": 12}, Month},
		{"um 10 tímar", pfpDiff{"hour": 10}, Day},
		// fur
		{"ca di 15 setemanis", pfpDiff{"week": 15}, Day},
		{"ca di 15 minût 20 secont", pfpDiff{"minute": 15, "second": 20}, Day},
		// fy
		{"oer 10 jier", pfpDiff{"year": 10}, Year},
		{"oer 22 deien", pfpDiff{"day": 22}, Day},
		// ga
		{"i gceann 23 bliain", pfpDiff{"year": 23}, Year},
		{"i gceann 12 scht", pfpDiff{"week": 12}, Day},
		// gd
		{"an ceann 10 bliadhna", pfpDiff{"year": 10}, Year},
		{"an ceann 18 latha", pfpDiff{"day": 18}, Day},
		// gl
		{"en 5 anos 26 mes", pfpDiff{"year": 5, "month": 26}, Month},
		{"en 14 semanas", pfpDiff{"week": 14}, Day},
		// gu
		{"10 મહિનામાં", pfpDiff{"month": 10}, Month},
		{"8 કલાકમાં", pfpDiff{"hour": 8}, Day},
		// hr
		{"za 12 dana", pfpDiff{"day": 12}, Day},
		{"za 10 sat 43 min", pfpDiff{"hour": 10, "minute": 43}, Day},
		// hsb
		{"za 6 měsacow", pfpDiff{"month": 6}, Month},
		{"za 1 dźeń 12 hodź", pfpDiff{"day": 1, "hour": 12}, Day},
		// hy
		{"7 ր-ից", pfpDiff{"minute": 7}, Day},
		{"51 շաբաթից", pfpDiff{"week": 51}, Day},
		// id
		{"dalam 12 detik", pfpDiff{"second": 12}, Day},
		{"dalam 10 hari", pfpDiff{"day": 10}, Day},
		// is
		{"eftir 11 mínútur", pfpDiff{"minute": 11}, Day},
		{"eftir 12 klukkustundir", pfpDiff{"hour": 12}, Day},
		// it
		{"tra 5 minuto", pfpDiff{"minute": 5}, Day},
		{"tra 16 settimane", pfpDiff{"week": 16}, Day},
		// jgo
		// {"nǔu ŋgu" 10", pfpDiff{"year": 10}, Year},
		{"nǔu ŋgap-mbi 11", pfpDiff{"week": 11}, Day},
		// ka
		{"5 საათში", pfpDiff{"hour": 5}, Day},
		{"3 კვირაში", pfpDiff{"week": 3}, Day},
		// kea
		{"di li 10 anu", pfpDiff{"year": 10}, Year},
		{"di li 43 minutu", pfpDiff{"minute": 43}, Day},
		// kk
		{"10 сағаттан кейін", pfpDiff{"hour": 10}, Day},
		{"18 айдан кейін", pfpDiff{"month": 18}, Month},
		// kl
		{"om 15 sapaatip-akunnera", pfpDiff{"week": 15}, Day},
		{"om 23 nalunaaquttap-akunnera", pfpDiff{"hour": 23}, Day},
		// km
		{"2 នាទីទៀត", pfpDiff{"minute": 2}, Day},
		{"5 សប្ដាហ៍ទៀត", pfpDiff{"week": 5}, Day},
		// kn
		{"10 ವಾರದಲ್ಲಿ", pfpDiff{"week": 10}, Day},
		{"15 ನಿಮಿಷಗಳಲ್ಲಿ", pfpDiff{"minute": 15}, Day},
		// ko
		{"5초 후", pfpDiff{"second": 5}, Day},
		{"7개월 후", pfpDiff{"month": 7}, Month},
		// ksh
		{"en 8 johre", pfpDiff{"year": 8}, Year},
		// ky
		{"15 мүнөттөн кийин", pfpDiff{"minute": 15}, Day},
		{"11 айд кийин", pfpDiff{"month": 11}, Month},
		// lb
		{"an 30 dag", pfpDiff{"day": 30}, Day},
		{"an 10 minutt 15 sekonn", pfpDiff{"minute": 10, "second": 15}, Day},
		// lkt
		{"letáŋhaŋ okó 20 kiŋháŋ", pfpDiff{"week": 20}, Day},
		{"letáŋhaŋ ómakȟa 11 kiŋháŋ", pfpDiff{"year": 11}, Year},
		// lo
		{"ໃນອີກ 25 ຊົ່ວໂມງ", pfpDiff{"hour": 25}, Day},
		{"ໃນອີກ 13 ອາທິດ", pfpDiff{"week": 13}, Day},
		// lt
		{"po 7 valandos", pfpDiff{"hour": 7}, Day},
		{"po 5 min 5 sek", pfpDiff{"minute": 5, "second": 5}, Day},
		// lv
		{"pēc 15 sekundēm", pfpDiff{"second": 15}, Day},
		{"pēc 10 mēneša", pfpDiff{"month": 10}, Month},
		// mk
		{"за 16 седмици", pfpDiff{"week": 16}, Day},
		{"за 2 месеци", pfpDiff{"month": 2}, Month},
		// ml
		{"5 ആഴ്ചയിൽ", pfpDiff{"week": 5}, Day},
		{"8 മിനിറ്റിൽ", pfpDiff{"minute": 8}, Day},
		// mn
		{"10 сарын дараа", pfpDiff{"month": 10}, Month},
		{"15 цагийн дараа", pfpDiff{"hour": 15}, Day},
		// mr
		{"2 महिन्यांमध्ये", pfpDiff{"month": 2}, Month},
		{"15 मिनि मध्ये", pfpDiff{"minute": 15}, Day},
		// ms
		{"dalam 6 jam", pfpDiff{"hour": 6}, Day},
		{"dalam 11 thn", pfpDiff{"year": 11}, Year},
		// my
		{"12 လအတွင်း", pfpDiff{"month": 12}, Month},
		{"8 နာရီအတွင်း", pfpDiff{"hour": 8}, Day},
		// nb
		{"om 1 måneder", pfpDiff{"month": 1}, Month},
		{"om 5 minutter", pfpDiff{"minute": 5}, Day},
		// ne
		{"10 वर्षमा", pfpDiff{"year": 10}, Year},
		{"15 घण्टामा", pfpDiff{"hour": 15}, Day},
		// nl
		{"over 3 weken", pfpDiff{"week": 3}, Day},
		{"over 12 seconden", pfpDiff{"second": 12}, Day},
		// nn
		{"om 7 uker", pfpDiff{"week": 7}, Day},
		{"om 2 timer", pfpDiff{"hour": 2}, Day},
		// os
		{"10 сахаты фӕстӕ", pfpDiff{"hour": 10}, Day},
		// pa-Guru
		{"3 ਸਾਲਾਂ ਵਿੱਚ", pfpDiff{"year": 3}, Year},
		{"7 ਦਿਨਾਂ ਵਿੱਚ", pfpDiff{"day": 7}, Day},
		// pa
		{"8 ਘੰਟਿਆਂ ਵਿੱਚ", pfpDiff{"hour": 8}, Day},
		{"16 ਸਕਿੰਟ ਵਿੱਚ", pfpDiff{"second": 16}, Day},
		// pl
		{"za 12 sekundy", pfpDiff{"second": 12}, Day},
		{"za 22 tygodnia", pfpDiff{"week": 22}, Day},
		// pt
		{"dentro de 11 minuto", pfpDiff{"minute": 11}, Day},
		{"dentro de 8 meses", pfpDiff{"month": 8}, Month},
		// ro
		{"peste 12 de săptămâni", pfpDiff{"week": 12}, Day},
		{"peste 6 de ore", pfpDiff{"hour": 6}, Day},
		// sah
		{"15 нэдиэлэннэн", pfpDiff{"week": 15}, Day},
		{"12 мүнүүтэннэн", pfpDiff{"minute": 12}, Day},
		// se
		{"3 mánotbadji maŋŋilit", pfpDiff{"month": 3}, Month},
		{"10 sekunda maŋŋilit", pfpDiff{"second": 10}, Day},
		// si
		{"මිනිත්තු 10කින්", pfpDiff{"minute": 10}, Day},
		{"දින 3න්", pfpDiff{"day": 3}, Day},
		// sk
		{"o 23 týždňov", pfpDiff{"week": 23}, Day},
		// sl
		{"čez 7 leto", pfpDiff{"year": 7}, Year},
		{"čez 8 minut 22 sek", pfpDiff{"minute": 8, "second": 22}, Day},
		// sq
		{"pas 2 muajsh", pfpDiff{"month": 2}, Month},
		{"pas 15 ditësh", pfpDiff{"day": 15}, Day},
		// sr-Cyrl
		{"за 3 годину", pfpDiff{"year": 3}, Year},
		{"за 10 мин 20 сек", pfpDiff{"minute": 10, "second": 20}, Day},
		// sr-Latn
		{"za 2 god 6 mes", pfpDiff{"year": 2, "month": 6}, Month},
		{"za 14 nedelja", pfpDiff{"week": 14}, Day},
		// sr
		{"за 18 недеља", pfpDiff{"week": 18}, Day},
		{"за 5 месеци", pfpDiff{"month": 5}, Month},
		// sv
		{"om 7 veckor", pfpDiff{"week": 7}, Day},
		{"om 10 timmar", pfpDiff{"hour": 10}, Day},
		// sw
		{"baada ya saa 21", pfpDiff{"hour": 21}, Day},
		{"baada ya sekunde 16", pfpDiff{"second": 16}, Day},
		// ta
		{"4 மாதங்களில்", pfpDiff{"month": 4}, Month},
		{"14 நாட்களில்", pfpDiff{"day": 14}, Day},
		// te
		{"3 వారాల్లో", pfpDiff{"week": 3}, Day},
		{"15 గంలో", pfpDiff{"hour": 15}, Day},
		// th
		{"ในอีก 6 นาที", pfpDiff{"minute": 6}, Day},
		{"ในอีก 3 ปี", pfpDiff{"year": 3}, Year},
		// to
		{"'i he māhina 'e 5", pfpDiff{"month": 5}, Month},
		{"'i he houa 'e 11", pfpDiff{"hour": 11}, Day},
		// tr
		{"15 saniye sonra", pfpDiff{"second": 15}, Day},
		{"45 saat 234 dakika sonra", pfpDiff{"hour": 45, "minute": 234}, Day},
		// uk
		{"через 8 хвилини", pfpDiff{"minute": 8}, Day},
		{"через 10 тижня", pfpDiff{"week": 10}, Day},
		// uz-Cyrl
		{"12 кундан сўнг", pfpDiff{"day": 12}, Day},
		{"10 дақиқадан сўнг", pfpDiff{"minute": 10}, Day},
		// uz-Latn
		{"3 yildan keyin", pfpDiff{"year": 3}, Year},
		{"5 haftadan keyin", pfpDiff{"week": 5}, Day},
		// uz
		{"12 kundan keyin", pfpDiff{"day": 12}, Day},
		{"50 daqiqadan keyin", pfpDiff{"minute": 50}, Day},
		// vi
		{"sau 5 năm nữa", pfpDiff{"year": 5}, Year},
		{"sau 2 phút nữa", pfpDiff{"minute": 2}, Day},
		// wae
		{"i 3 stunde", pfpDiff{"hour": 3}, Day},
		{"i 5 täg", pfpDiff{"day": 5}, Day},
		// yue
		{"3 個星期後", pfpDiff{"week": 3}, Day},
		{"6 年後", pfpDiff{"year": 6}, Year},
		// zh-Hans
		{"5个月后", pfpDiff{"month": 5}, Month},
		{"7天后", pfpDiff{"day": 7}, Day},
		// zh-Hant
		{"2 分鐘後", pfpDiff{"minute": 2}, Day},
		{"4 週後", pfpDiff{"week": 4}, Day},
	}

	crazyTimes := []testScenario{
		{"5000 years ago", pfpDiff{"year": -5000}, Year},
		{"2014 years ago", pfpDiff{"year": -2014}, Year},
		{"123456789 hour", pfpDiff{"hour": -123456789}, Day},
		{"123456789123 hour", pfpDiff{"hour": -123456789123}, Day},
		{"1234567 days", pfpDiff{"day": -1234567}, Day},
		{"1234567891 days", pfpDiff{"day": -1234567891}, Day},
		{"12345678912 days", pfpDiff{"day": -12345678912}, Day},
		{"123455678976543 month", pfpDiff{"month": -123455678976543}, Month},
	}

	pastSeconds := []testScenario{
		{"несколько секунд назад", pfpDiff{"second": -44}, Day},
		{"há alguns segundos", pfpDiff{"second": -44}, Day},
	}

	// Prepare parser
	parser := Parser{Config: &relativeTestConfig}

	// Start tests
	tests := []testScenario{}
	tests = append(tests, pastTimes...)
	tests = append(tests, futureTimes...)
	tests = append(tests, crazyTimes...)
	tests = append(tests, pastSeconds...)

	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => %v (%d)", test.String, strutil.Jsonify(&test.Diff), test.Period)

		// Parse date time
		dt, err := parser.Parse(test.String)
		passed := assert.Nil(t, err, message)
		if passed {
			passed = assert.Equal(t, test.Period, dt.Period, message)
		}
		if passed {
			passed = assert.Equal(t, expectedDate(test.Diff), dt.Time, message)
		}

		// Track failure
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		nTest := len(pastTimes) + len(futureTimes)
		fmt.Printf("Failed %d from %d tests\n", nFailed, nTest)
	}
}

func TestParser_Parse_relative_invalidDates(t *testing.T) {
	parser := Parser{Config: &relativeTestConfig}
	dt, err := parser.Parse("15th of Aug, 2014 Diane Bennett")
	assert.Error(t, err)
	assert.True(t, dt.IsZero())
}

func TestParser_Parse_relative_hasSpecificTime(t *testing.T) {
	// Prepare struct
	type testScenario struct {
		Text        string
		CurrentTime time.Time
		Expected    time.Time
	}

	// Helper function
	tt := func(Y, M, D, h, m int, locName ...string) time.Time {
		tz := time.UTC
		if len(locName) > 0 {
			tz = timezone.MustLoad(locName[0])
		}
		return time.Date(Y, time.Month(M), D, h, m, 0, 0, tz)
	}

	ts := func(text string, expected time.Time) testScenario {
		return testScenario{Text: text, Expected: expected}
	}

	tzTs := func(text, tzName string, e time.Time) testScenario {
		tz := timezone.MustLoad(tzName)
		e = time.Date(e.Year(), e.Month(), e.Day(),
			e.Hour(), e.Minute(), e.Second(), e.Nanosecond(), tz)
		return testScenario{Text: text, Expected: e, CurrentTime: relativeTestNow.In(tz)}
	}

	ctTs := func(text string, ct, e time.Time) testScenario {
		return testScenario{Text: text, Expected: e, CurrentTime: ct}
	}

	// Prepare scenarios
	ct := tt(2010, 6, 4, 13, 15)
	ctUsEastern := tt(2014, 9, 1, 10, 30, "US/Eastern")
	ctUsMountain := tt(2014, 9, 1, 10, 30, "US/Mountain")

	tests := []testScenario{
		// No timezone in string
		ts("Today at 9 pm", tt(2014, 9, 1, 21, 0)),
		ts("Today at 11:20 am", tt(2014, 9, 1, 11, 20)),
		ts("Yesterday 1:20 pm", tt(2014, 8, 31, 13, 20)),
		ts("Yesterday by 13:20", tt(2014, 8, 31, 13, 20)),
		ts("the day before yesterday 16:50", tt(2014, 8, 30, 16, 50)),
		ts("2 Tage 18:50", tt(2014, 8, 30, 18, 50)),
		ts("1 day ago at 2 PM", tt(2014, 8, 31, 14, 0)),
		ts("one day ago at 2 PM", tt(2014, 8, 31, 14, 0)),
		ts("Dnes v 12:40", tt(2014, 9, 1, 12, 40)),
		ts("1 week ago at 12:00 am", tt(2014, 8, 25, 0, 0)),
		ts("one week ago at 12:00 am", tt(2014, 8, 25, 0, 0)),
		ts("tomorrow at 2 PM", tt(2014, 9, 2, 14, 0)),

		// Has timezone in string
		ts("tomorrow 8:30 CST", tt(2014, 9, 2, 8, 30, "CST")),

		// Use IANA timezone
		tzTs("2 hours ago", "Asia/Karachi", tt(2014, 9, 1, 13, 30)),
		tzTs("3 hours ago", "Europe/Paris", tt(2014, 9, 1, 9, 30)),
		tzTs("5 hours ago", "US/Eastern", tt(2014, 9, 1, 1, 30)),
		tzTs("Today at 9 pm", "Asia/Karachi", tt(2014, 9, 1, 21, 0)),

		// Use timezone abbreviations
		tzTs("2 hours ago", "PKT", tt(2014, 9, 1, 13, 30)),
		tzTs("5 hours ago", "EST", tt(2014, 9, 1, 0, 30)),
		tzTs("3 hours ago", "MET", tt(2014, 9, 1, 8, 30)),

		// Use UTC offsets
		tzTs("2 hours ago", "+05:00", tt(2014, 9, 1, 13, 30)),
		tzTs("5 hours ago", "-05:00", tt(2014, 9, 1, 0, 30)),
		tzTs("3 hours ago", "+01:00", tt(2014, 9, 1, 8, 30)),

		// Use custom current time
		ctTs("Today at 4:25 pm", ctUsMountain, tt(2014, 9, 1, 22, 25)),
		ctTs("Yesterday at 4:25 pm", ctUsMountain, tt(2014, 8, 31, 22, 25)),
		ctTs("Yesterday", ctUsMountain, tt(2014, 8, 31, 16, 30)),
		ctTs("Today", ctUsMountain, tt(2014, 9, 1, 16, 30)),
		ctTs("1 minute ago", ctUsEastern, tt(2014, 9, 1, 14, 29)),
		ctTs("Today at 9 pm", ct, tt(2010, 6, 4, 21, 0)),
		ctTs("Today at 11:20 am", ct, tt(2010, 6, 4, 11, 20)),
		ctTs("Yesterday 1:20 pm", ct, tt(2010, 6, 3, 13, 20)),
		ctTs("the day before yesterday 16:50", ct, tt(2010, 6, 2, 16, 50)),
		ctTs("2 Tage 18:50", ct, tt(2010, 6, 2, 18, 50)),
		ctTs("1 day ago at 2 PM", ct, tt(2010, 6, 3, 14, 0)),
		ctTs("Dnes v 12:40", ct, tt(2010, 6, 4, 12, 40)),
		ctTs("1 week ago at 12:00 am", ct, tt(2010, 5, 28, 0, 0)),
		ctTs("yesterday", ct, tt(2010, 6, 3, 13, 15)),
		ctTs("the day before yesterday", ct, tt(2010, 6, 2, 13, 15)),
		ctTs("today", ct, tt(2010, 6, 4, 13, 15)),
		ctTs("an hour ago", ct, tt(2010, 6, 4, 12, 15)),
		ctTs("about an hour ago", ct, tt(2010, 6, 4, 12, 15)),
		ctTs("a day ago", ct, tt(2010, 6, 3, 13, 15)),
		ctTs("a week ago", ct, tt(2010, 5, 28, 13, 15)),
		ctTs("2 hours ago", ct, tt(2010, 6, 4, 11, 15)),
		ctTs("about 23 hours ago", ct, tt(2010, 6, 3, 14, 15)),
		ctTs("1 year 2 months", ct, tt(2009, 4, 4, 13, 15)),
		ctTs("1 year, 09 months,01 weeks", ct, tt(2008, 8, 28, 13, 15)),
		ctTs("1 year 11 months", ct, tt(2008, 7, 4, 13, 15)),
		ctTs("1 year 12 months", ct, tt(2008, 6, 4, 13, 15)),
		ctTs("15 hr", ct, tt(2010, 6, 3, 22, 15)),
		ctTs("15 hrs", ct, tt(2010, 6, 3, 22, 15)),
		ctTs("2 min", ct, tt(2010, 6, 4, 13, 13)),
		ctTs("2 mins", ct, tt(2010, 6, 4, 13, 13)),
		ctTs("3 sec", ct, time.Date(2010, 6, 4, 13, 14, 57, 0, time.UTC)),
		ctTs("1000 years ago", ct, tt(1010, 6, 4, 13, 15)),
		ctTs("2008 years ago", ct, tt(2, 6, 4, 13, 15)),
		ctTs("5000 months ago", ct, tt(1593, 10, 4, 13, 15)),
		ctTs(fmt.Sprintf("%d months ago", 2008*12+8), ct, tt(1, 10, 4, 13, 15)),
		ctTs("1 year, 1 month, 1 week, 1 day, 1 hour and 1 minute ago", ct, tt(2009, 4, 26, 12, 14)),
		ctTs("just now", ct, tt(2010, 6, 4, 13, 15)),
	}

	// Prepare parser
	cfg := relativeTestConfig.Clone()
	parser := Parser{Config: &cfg}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\", currently \"%s\" => \"%s\"",
			test.Text,
			test.CurrentTime.Format("2006-01-02 15:04:05 MST -0700"),
			test.Expected.Format("2006-01-02 15:04:05 MST -0700"))

		// Prepare config
		cfg.CurrentTime = relativeTestNow
		if !test.CurrentTime.IsZero() {
			cfg.CurrentTime = test.CurrentTime
		}

		// Parse date time
		dt, err := parser.Parse(test.Text)
		passed := assert.Nil(t, err, message)
		if passed {
			passed = assert.Zero(t, test.Expected.Sub(dt.Time).Seconds(), message)
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_relative_hasPreferredTimes(t *testing.T) {
	// Prepare struct
	type testScenario struct {
		Text         string
		Expected     time.Time
		PreferFuture bool
	}

	// Helper function
	tt := func(Y, M, D, h, m int) time.Time {
		return time.Date(Y, time.Month(M), D, h, m, 0, 0, time.UTC)
	}

	// Prepare parser
	cfg := relativeTestConfig.Clone()
	cfg.CurrentTime = tt(2010, 6, 4, 13, 15)
	parser := Parser{Config: &cfg}

	// Prepare scenarios
	tests := []testScenario{
		{"3 days", tt(2010, 6, 1, 13, 15), false},
		{"2 years", tt(2008, 6, 4, 13, 15), false},
		{"3 days", tt(2010, 6, 7, 13, 15), true},
		{"2 years", tt(2012, 6, 4, 13, 15), true},
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\", future: %v => \"%s\"",
			test.Text, test.PreferFuture, test.Expected.Format("2006-01-02 15:04:05 MST -0700"))

		// Prepare config
		if test.PreferFuture {
			cfg.PreferredDateSource = Future
		} else {
			cfg.PreferredDateSource = Past
		}

		// Parse date time
		dt, err := parser.Parse(test.Text)
		passed := assert.Nil(t, err, message)
		if passed {
			passed = assert.Zero(t, test.Expected.Sub(dt.Time).Seconds(), message)
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
