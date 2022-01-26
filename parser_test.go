package dateparser

import (
	"fmt"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text         string
		ExpectedTime time.Time
	}

	tests := []testScenario{
		// === ABSOLUTE PARSER ===
		// English dates
		{"[Sept] 04, 2014.", tt(2014, 9, 4)},
		{"Tuesday Jul 22, 2014", tt(2014, 7, 22)},
		{"Tues 9th Aug, 2015", tt(2015, 8, 9)},
		{"10:04am", tt(2012, 11, 13, 10, 4)},
		{"Friday", tt(2012, 11, 9)},
		{"November 19, 2014 at noon", tt(2014, 11, 19, 12, 0)},
		{"December 13, 2014 at midnight", tt(2014, 12, 13, 0, 0)},
		{"Nov 25 2014 10:17 pm", tt(2014, 11, 25, 22, 17)},
		{"Wed Aug 05 12:00:00 2015", tt(2015, 8, 5, 12, 0)},
		{"April 9, 2013 at 6:11 a.m.", tt(2013, 4, 9, 6, 11)},
		{"Aug. 9, 2012 at 2:57 p.m.", tt(2012, 8, 9, 14, 57)},
		{"December 10, 2014, 11:02:21 pm", tt(2014, 12, 10, 23, 2, 21)},
		{"8:25 a.m. Dec. 12, 2014", tt(2014, 12, 12, 8, 25)},
		{"2:21 p.m., December 11, 2014", tt(2014, 12, 11, 14, 21)},
		{"Fri, 12 Dec 2014 10:55:50", tt(2014, 12, 12, 10, 55, 50)},
		{"20 Mar 2013 10h11", tt(2013, 3, 20, 10, 11)},
		{"10:06am Dec 11, 2014", tt(2014, 12, 11, 10, 6)},
		{"19 February 2013 year 09:10", tt(2013, 2, 19, 9, 10)},
		{"21 January 2012 13:11:23.678", tt(2012, 1, 21, 13, 11, 23, 678000)},
		{"1/1/16 9:02:43.1", tt(2016, 1, 1, 9, 2, 43, 100000)},
		{"29.02.2020 13.12", tt(2020, 2, 29, 13, 12)},
		{"Wednesday, 22nd June, 2016, 12.16 pm.", tt(2016, 6, 22, 12, 16)},
		{"16:00", tt(2012, 11, 13, 16, 0)},
		{"Monday 7:15 AM", tt(2012, 11, 12, 7, 15)},
		// French dates
		{"11 Mai 2014", tt(2014, 5, 11)},
		{"dimanche, 11 Mai 2014", tt(2014, 5, 11)},
		{"22 janvier 2015 à 14h40", tt(2015, 1, 22, 14, 40)},
		{"Dimanche 1er Février à 21:24", tt(2012, 2, 1, 21, 24)},
		{"vendredi, décembre 5 2014.", tt(2014, 12, 5, 0, 0)},
		{"le 08 Déc 2014 15:11", tt(2014, 12, 8, 15, 11)},
		{"Le 11 Décembre 2014 à 09:00", tt(2014, 12, 11, 9, 0)},
		{"fév 15, 2013", tt(2013, 2, 15, 0, 0)},
		{"Jeu 15:12", tt(2012, 11, 8, 15, 12)},
		// Spanish dates
		{"Martes 21 de Octubre de 2014", tt(2014, 10, 21)},
		{"Miércoles 20 de Noviembre de 2013", tt(2013, 11, 20)},
		{"12 de junio del 2012", tt(2012, 6, 12)},
		{"13 Ago, 2014", tt(2014, 8, 13)},
		{"13 Septiembre, 2014", tt(2014, 9, 13)},
		{"11 Marzo, 2014", tt(2014, 3, 11)},
		{"julio 5, 2015 en 1:04 pm", tt(2015, 7, 5, 13, 4)},
		{"Vi 17:15", tt(2012, 11, 9, 17, 15)},
		// Dutch dates
		{"11 augustus 2014", tt(2014, 8, 11)},
		{"14 januari 2014", tt(2014, 1, 14)},
		{"vr jan 24, 2014 12:49", tt(2014, 1, 24, 12, 49)},
		// Italian dates
		{"16 giu 2014", tt(2014, 6, 16)},
		{"26 gennaio 2014", tt(2014, 1, 26)},
		{"Ven 18:23", tt(2012, 11, 9, 18, 23)},
		// Portuguese dates
		{"sexta-feira, 10 de junho de 2014 14:52", tt(2014, 6, 10, 14, 52)},
		{"13 Setembro, 2014", tt(2014, 9, 13)},
		{"Sab 3:03", tt(2012, 11, 10, 3, 3)},
		// Russian dates
		{"10 мая", tt(2012, 5, 10)}, // forum.codenet.ru
		{"26 апреля", tt(2012, 4, 26)},
		{"20 ноября 2013", tt(2013, 11, 20)},
		{"28 октября 2014 в 07:54", tt(2014, 10, 28, 7, 54)},
		{"13 января 2015 г. в 13:34", tt(2015, 1, 13, 13, 34)},
		{"09 августа 2012", tt(2012, 8, 9, 0, 0)},
		{"Авг 26, 2015 15:12", tt(2015, 8, 26, 15, 12)},
		{"2 Декабрь 95 11:15", tt(1995, 12, 2, 11, 15)},
		{"13 янв. 2005 19:13", tt(2005, 1, 13, 19, 13)},
		{"13 авг. 2005 19:13", tt(2005, 8, 13, 19, 13)},
		{"13 авг. 2005г. 19:13", tt(2005, 8, 13, 19, 13)},
		{"13 авг. 2005 г. 19:13", tt(2005, 8, 13, 19, 13)},
		{"21 сентября 2021г., вторник", tt(2021, 9, 21, 0, 0)},
		// Turkish dates
		{"11 Ağustos, 2014", tt(2014, 8, 11)},
		{"08.Haziran.2014, 11:07", tt(2014, 6, 8, 11, 7)}, // forum.andronova.net
		{"17.Şubat.2014, 17:51", tt(2014, 2, 17, 17, 51)},
		{"14-Aralık-2012, 20:56", tt(2012, 12, 14, 20, 56)}, // forum.ceviz.net
		// Romanian dates
		{"13 iunie 2013", tt(2013, 6, 13)},
		{"14 aprilie 2014", tt(2014, 4, 14)},
		{"18 martie 2012", tt(2012, 3, 18)},
		{"12-Iun-2013", tt(2013, 6, 12)},
		// German dates
		{"21. Dezember 2013", tt(2013, 12, 21)},
		{"19. Februar 2012", tt(2012, 2, 19)},
		{"26. Juli 2014", tt(2014, 7, 26)},
		{"18.10.14 um 22:56 Uhr", tt(2014, 10, 18, 22, 56)},
		{"12-Mär-2014", tt(2014, 3, 12)},
		{"Mit 13:14", tt(2012, 11, 7, 13, 14)},
		{"23. März 18.37 Uhr", tt(2012, 3, 23, 18, 37)},
		// Czech dates
		{"pon 16. čer 2014 10:07:43", tt(2014, 6, 16, 10, 7, 43)},
		{"13 Srpen, 2014", tt(2014, 8, 13)},
		{"čtv 14. lis 2013 12:38:43", tt(2013, 11, 14, 12, 38, 43)},
		// Thai dates
		{"ธันวาคม 11, 2014, 08:55:08 PM", tt(2014, 12, 11, 20, 55, 8)},
		{"22 พฤษภาคม 2012, 22:12", tt(2012, 5, 22, 22, 12)},
		{"11 กุมภา 2020, 8:13 AM", tt(2020, 2, 11, 8, 13)},
		{"1 เดือนตุลาคม 2005, 1:00 AM", tt(2005, 10, 1, 1, 0)},
		{"11 ก.พ. 2020, 1:13 pm", tt(2020, 2, 11, 13, 13)},
		// Vietnamese dates
		{"Thứ năm", tt(2012, 11, 8)},                                 // Thursday
		{"Thứ sáu", tt(2012, 11, 9)},                                 // Friday
		{"Tháng Mười Hai 29, 2013, 14:14", tt(2013, 12, 29, 14, 14)}, // bpsosrcs.wordpress.com NOQA
		{"05 Tháng một 2015 - 03:54 AM", tt(2015, 1, 5, 3, 54)},
		// Belarusian dates
		{"11 траўня", tt(2012, 5, 11)},
		{"4 мая", tt(2012, 5, 4)},
		{"Чацвер 06 жніўня 2015", tt(2015, 8, 6)},
		{"Нд 14 сакавіка 2015 у 7 гадзін 10 хвілін", tt(2015, 3, 14, 7, 10)},
		{"5 жніўня 2015 года у 13:34", tt(2015, 8, 5, 13, 34)},
		// Ukrainian dates
		{"2015-кві-12", tt(2015, 4, 12)},
		{"21 чер 2013 3:13", tt(2013, 6, 21, 3, 13)},
		{"12 лютого 2012, 13:12:23", tt(2012, 2, 12, 13, 12, 23)},
		{"вів о 14:04", tt(2012, 11, 13, 14, 4)},
		// Tagalog dates
		{"12 Hulyo 2003 13:01", tt(2003, 7, 12, 13, 1)},
		{"1978, 1 Peb, 7:05 PM", tt(1978, 2, 1, 19, 5)},
		{"2 hun", tt(2012, 6, 2)},
		{"Lin 16:16", tt(2012, 11, 11, 16, 16)},
		// Japanese dates
		{"2016年3月20日(日) 21時40分", tt(2016, 3, 20, 21, 40)},
		{"2016年3月20日 21時40分", tt(2016, 3, 20, 21, 40)},
		// Numeric dates
		{"06-17-2014", tt(2014, 6, 17)},
		{"13/03/2014", tt(2014, 3, 13)},
		{"11. 12. 2014, 08:45:39", tt(2014, 11, 12, 8, 45, 39)},
		// Miscellaneous dates
		{"1 Ni 2015", tt(2015, 4, 1, 0, 0)},
		{"1 Mar 2015", tt(2015, 3, 1, 0, 0)},
		{"1 сер 2015", tt(2015, 8, 1, 0, 0)},
		// Chinese dates
		{"2015年04月08日10:05", tt(2015, 4, 8, 10, 5)},
		{"2012年12月20日10:35", tt(2012, 12, 20, 10, 35)},
		{"2016年06月30日09时30分", tt(2016, 6, 30, 9, 30)},
		{"2016年6月2911:30", tt(2016, 6, 29, 11, 30)},
		{"2016年6月29", tt(2016, 6, 29, 0, 0)},
		{"2016年 2月 5日", tt(2016, 2, 5, 0, 0)},
		{"2016年9月14日晚8:00", tt(2016, 9, 14, 20, 0)},
		// Bulgarian
		{"25 ян 2016", tt(2016, 1, 25, 0, 0)},
		{"23 декември 2013 15:10:01", tt(2013, 12, 23, 15, 10, 1)},
		// Bangla dates
		{"[সেপ্টেম্বর] 04, 2014.", tt(2014, 9, 4)},
		{"মঙ্গলবার জুলাই 22, 2014", tt(2014, 7, 22)},
		{"শুক্রবার", tt(2012, 11, 9)},
		{"শুক্র, 12 ডিসেম্বর 2014 10:55:50", tt(2014, 12, 12, 10, 55, 50)},
		{"1লা জানুয়ারী 2015", tt(2015, 1, 1)},
		{"25শে মার্চ 1971", tt(1971, 3, 25)},
		{"8ই মে 2002", tt(2002, 5, 8)},
		{"10:06am ডিসেম্বর 11, 2014", tt(2014, 12, 11, 10, 6)},
		{"19 ফেব্রুয়ারী 2013 সাল 09:10", tt(2013, 2, 19, 9, 10)},
		// Hindi dates
		{"11 जुलाई 1994, 11:12", tt(1994, 7, 11, 11, 12)},
		{"१७ अक्टूबर २०१८", tt(2018, 10, 17, 0, 0)},
		{"12 जनवरी  1997 11:08 अपराह्न", tt(1997, 1, 12, 23, 8)},
		// Georgian dates
		{"2011 წლის 17 მარტი, ოთხშაბათი", tt(2011, 3, 17, 0, 0)},
		{"2015 წ. 12 ივნ, 15:34", tt(2015, 6, 12, 15, 34)},
		// Finnish dates
		{"5.7.2018 5.45 ip.", tt(2018, 7, 5, 17, 45)},
		// Others (from `test_clean_api.py`)
		{"24 de Janeiro de 2014", tt(2014, 1, 24)},
		{"2 de Enero de 2013", tt(2013, 1, 2)},
		{"January 25, 2014", tt(2014, 1, 25)},
		{"May 5, 2000 13:00", tt(2000, 5, 5, 13, 0)},
		{"August 8, 2018 5 PM", tt(2018, 8, 8, 17, 0)},
		{"February 26, 1981 5 am UTC", tt(1981, 2, 26, 5, 0)},
		{"2014/11/17 14:56 EDT", tt(2014, 11, 17, 18, 56)},
		{"08-08-2014\xa018:29", tt(2014, 8, 8, 18, 29)},
		{"12 jan 1876", tt(1876, 1, 12)},
		{"02/09/16", tt(2016, 2, 9)},
		{"10 giu 2018", tt(2018, 6, 10)},

		// === CONTAINS TIMEZONE ===
		{"Sep 03 2014 | 4:32 pm EDT", tt(2014, 9, 3, 20, 32)},
		{"17th October, 2034 @ 01:08 am PDT", tt(2034, 10, 17, 8, 8)},
		{"15 May 2004 23:24 EDT", tt(2004, 5, 16, 3, 24)},
		{"08/17/14 17:00 (PDT)", tt(2014, 8, 18, 0, 0)},
		{"15 May 2004 16:10 -0400", tt(2004, 5, 15, 20, 10)},
		{"1999-12-31 19:00:00 -0500", tt(2000, 1, 1)},
		{"1999-12-31 19:00:00 +0500", tt(1999, 12, 31, 14, 0)},
		{"Fri, 09 Sep 2005 13:51:39 -0700", tt(2005, 9, 9, 20, 51, 39)},
		{"Fri, 09 Sep 2005 13:51:39 +0000", tt(2005, 9, 9, 13, 51, 39)},
		{"Mon, 25 Jun 2018 10:37:47 +0530 ", tt(2018, 6, 25, 5, 7, 47)}, // trailing spaces
		{"Fri Sep 23 2016 10:34:51 GMT+0800 (CST)", tt(2016, 9, 23, 2, 34, 51)},

		// === NOSPACE PARSER ===
		{"2016020417:10", tt(2016, 2, 4, 17, 10)},

		// === ISO DATESTAMP ===
		{"2015-05-02T10:20:19+0000", tt(2015, 5, 2, 10, 20, 19)},

		// === TIMESTAMP ===
		{"1484823450", tt(2017, 1, 19, 10, 57, 30)},
		{"1436745600000", tt(2015, 7, 13, 0, 0)},
		{"1015673450", tt(2002, 3, 9, 11, 30, 50)},
		{"2016-09-23T02:54:32.845Z", tt(2016, 9, 23, 2, 54, 32, 845000)},
	}

	// Prepare config
	cfg := Configuration{CurrentTime: tt(2012, 11, 13)}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" => \"%s\"", test.Text,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Parse text
		dt, _ := Parse(&cfg, test.Text)
		if passed := assertResult(t, dt, test.ExpectedTime, date.Day, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_hasPreferredDateSource(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		DateSource   PreferredDateSource
		Text         string
		ExpectedTime time.Time
		BaseTime     time.Time
	}

	ts := func(ds PreferredDateSource, text string, eTime time.Time, bTime ...time.Time) testScenario {
		var baseTime time.Time
		if len(bTime) == 0 {
			baseTime = tt(2015, 2, 15, 15, 30)
		} else {
			baseTime = bTime[0]
		}

		return testScenario{ds, text, eTime, baseTime}
	}

	tests := []testScenario{
		// Prefer past dates
		ts(Past, "10 December", tt(2014, 12, 10)),
		ts(Past, "March", tt(2014, 3, 15)),
		ts(Past, "Friday", tt(2015, 2, 13)),
		ts(Past, "Monday", tt(2015, 2, 9)),
		ts(Past, "Sunday", tt(2015, 2, 8)), // current day
		ts(Past, "10:00PM", tt(2015, 2, 14, 22, 0)),
		ts(Past, "16:10", tt(2015, 2, 14, 16, 10)),
		ts(Past, "14:05", tt(2015, 2, 15, 14, 5)),
		ts(Past, "15 february 15:00", tt(2015, 2, 15, 15, 0)),
		ts(Past, "3/3/50", tt(1950, 3, 3)),
		ts(Past, "3/3/94", tt(1994, 3, 3)),

		// Prefer future dates
		ts(Future, "10 December", tt(2015, 12, 10)),
		ts(Future, "March", tt(2015, 3, 15)),
		ts(Future, "Friday", tt(2015, 2, 20)),
		ts(Future, "Sunday", tt(2015, 2, 22)), // current day
		ts(Future, "Monday", tt(2015, 2, 16)),
		ts(Future, "10:00PM", tt(2015, 2, 15, 22, 0)),
		ts(Future, "16:10", tt(2015, 2, 15, 16, 10)),
		ts(Future, "14:05", tt(2015, 2, 16, 14, 5)),
		ts(Future, "3/3/50", tt(2050, 3, 3)),
		ts(Future, "3/3/94", tt(2094, 3, 3)),

		// Prefer current period
		ts(CurrentPeriod, "10 December", tt(2015, 12, 10)),
		ts(CurrentPeriod, "March", tt(2015, 3, 15)),
		ts(CurrentPeriod, "Friday", tt(2015, 2, 13)),
		ts(CurrentPeriod, "Sunday", tt(2015, 2, 15)), // current weekday
		ts(CurrentPeriod, "10:00PM", tt(2015, 2, 15, 22, 0)),
		ts(CurrentPeriod, "16:10", tt(2015, 2, 15, 16, 10)),
		ts(CurrentPeriod, "14:05", tt(2015, 2, 15, 14, 5)),

		// Prefer leap year from the past
		ts(Past, "29 Feb", tt(2016, 2, 29), tt(2020, 1, 1)),
		ts(Past, "29/02", tt(2020, 2, 29), tt(2020, 3, 30)),
		ts(Past, "02/29", tt(1696, 2, 29), tt(1702, 1, 1)),

		// Prefer leap year from the future
		ts(Future, "29 Feb", tt(2020, 2, 29), tt(2020, 1, 1)),
		ts(Future, "29/02", tt(2024, 2, 29), tt(2020, 3, 30)),
		ts(Future, "02/29", tt(1704, 2, 29), tt(1696, 3, 1)),

		// No preference for the leap year
		ts(CurrentPeriod, "29 Feb", tt(2020, 2, 29), tt(2020, 1, 1)),
		ts(CurrentPeriod, "29/02", tt(2020, 2, 29), tt(2020, 3, 30)),
		ts(CurrentPeriod, "29 Feb", tt(1704, 2, 29), tt(1702, 3, 1)),
		ts(CurrentPeriod, "02/29", tt(1696, 2, 29), tt(1699, 3, 1)),
	}

	// Prepare config
	cfg := Configuration{}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%v) => \"%s\"",
			test.Text, test.DateSource,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Prepare configuration
		cfg.CurrentTime = test.BaseTime
		cfg.PreferredDateSource = test.DateSource

		// Parse text
		dt, err := Parse(&cfg, test.Text)

		// Check the result
		passed := assert.Nil(t, err, message)
		if passed {
			diff := test.ExpectedTime.Sub(dt.Time)
			passed = assert.Zero(t, diff, message)
		}

		if !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_hasPreferredDayOfMonth(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		DayOfMonth   PreferredDayOfMonth
		Text         string
		BaseTime     time.Time
		ExpectedTime time.Time
	}

	tests := []testScenario{
		// Prefer current day of month
		{Current, "February 2015", tt(2015, 1, 31), tt(2015, 2, 28)},
		{Current, "February 2012", tt(2015, 1, 31), tt(2012, 2, 29)},
		{Current, "March 2015", tt(2015, 1, 25), tt(2015, 3, 25)},
		{Current, "April 2015", tt(2015, 1, 31), tt(2015, 4, 30)},
		{Current, "April 2015", tt(2015, 2, 28), tt(2015, 4, 28)},
		{Current, "December 2014", tt(2015, 2, 15), tt(2014, 12, 15)},

		// Prefer last day of month
		{Last, "February 2015", tt(2015, 1, 1), tt(2015, 2, 28)},
		{Last, "February 2012", tt(2015, 1, 1), tt(2012, 2, 29)},
		{Last, "March 2015", tt(2015, 1, 25), tt(2015, 3, 31)},
		{Last, "April 2015", tt(2015, 1, 15), tt(2015, 4, 30)},
		{Last, "April 2015", tt(2015, 2, 28), tt(2015, 4, 30)},
		{Last, "December 2014", tt(2015, 2, 15), tt(2014, 12, 31)},

		// Prefer first day of month
		{First, "February 2015", tt(2015, 1, 8), tt(2015, 2, 1)},
		{First, "February 2012", tt(2015, 1, 7), tt(2012, 2, 1)},
		{First, "March 2015", tt(2015, 1, 25), tt(2015, 3, 1)},
		{First, "April 2015", tt(2015, 1, 15), tt(2015, 4, 1)},
		{First, "April 2015", tt(2015, 2, 28), tt(2015, 4, 1)},
		{First, "December 2014", tt(2015, 2, 15), tt(2014, 12, 1)},

		// Preference doesn't affect date with explicit day
		{Last, "24 April 2012", tt(2015, 2, 12), tt(2012, 4, 24)},
		{First, "24 April 2012", tt(2015, 2, 12), tt(2012, 4, 24)},
		{Current, "24 April 2012", tt(2015, 2, 12), tt(2012, 4, 24)},
	}

	// Prepare config
	cfg := Configuration{}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%v), base: \"%s\" => \"%s\"",
			test.Text, test.DayOfMonth,
			test.BaseTime.Format("2006-01-02 15:04:05"),
			test.ExpectedTime.Format("2006-01-02 15:04:05"))

		// Prepare configuration
		cfg.CurrentTime = test.BaseTime
		cfg.PreferredDayOfMonth = test.DayOfMonth

		// Parse text
		dt, err := Parse(&cfg, test.Text)

		// Check the result
		passed := assert.Nil(t, err, message)
		if passed {
			diff := test.ExpectedTime.Sub(dt.Time)
			passed = assert.Zero(t, diff, message)
		}

		if !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_returnTimeAsPeriod(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text           string
		BaseTime       time.Time
		ExpectedTime   time.Time
		ExpectedPeriod date.Period
	}

	ts := func(text string, eTime time.Time, bTime ...time.Time) testScenario {
		var baseTime time.Time
		if len(bTime) > 0 {
			baseTime = bTime[0]
		}
		return testScenario{text, baseTime, eTime, date.Time}
	}

	tsd := func(text string, eTime time.Time, bTime ...time.Time) testScenario {
		test := ts(text, eTime, bTime...)
		test.ExpectedPeriod = date.Day
		return test
	}

	tests := []testScenario{
		// Timestamp
		ts("1484823450", tt(2017, 1, 19, 10, 57, 30)),
		ts("1436745600000", tt(2015, 7, 13, 0, 0)),
		ts("1015673450", tt(2002, 3, 9, 11, 30, 50)),
		ts("2016-09-23T02:54:32.845Z", tt(2016, 9, 23, 2, 54, 32, 845000)),

		// Basic text
		ts("12th December 2019 19:00", tt(2019, 12, 12, 19, 0)),
		ts("9 Jan 11 0:00", tt(2011, 1, 9, 0, 0)),

		// Basic text with base time
		ts("10:04am EDT", tt(2020, 7, 19, 14, 4), tt(2020, 7, 19)),
		ts("16:00", tt(2018, 12, 13, 16, 0), tt(2018, 12, 13, 15, 15)),
		ts("Monday 7:15 AM", tt(2018, 12, 10, 7, 15), tt(2018, 12, 13, 15, 15)),

		// Period is day when time is not present
		tsd("12th March 2010", tt(2010, 3, 12, 0, 0)),
		tsd("21-12-19", tt(2019, 12, 21, 0, 0)),
	}

	// Prepare config
	cfg := Configuration{ReturnTimeAsPeriod: true}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%v), base: \"%s\" => \"%s\"",
			test.Text, test.ExpectedPeriod,
			test.BaseTime.Format("2006-01-02 15:04:05"),
			test.ExpectedTime.Format("2006-01-02 15:04:05"))

		// Parse text
		cfg.CurrentTime = test.BaseTime
		dt, _ := Parse(&cfg, test.Text)
		if passed := assertResult(t, dt, test.ExpectedTime, test.ExpectedPeriod, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_checkPeriod(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text           string
		ExpectedTime   time.Time
		ExpectedPeriod date.Period
	}

	tests := []testScenario{
		{"10 December", tt(2015, 12, 10), date.Day},
		{"March", tt(2015, 3, 15), date.Month},
		{"April", tt(2015, 4, 15), date.Month},
		{"December", tt(2015, 12, 15), date.Month},
		{"Friday", tt(2015, 2, 13), date.Day},
		{"Monday", tt(2015, 2, 9), date.Day},
		{"10:00PM", tt(2015, 2, 15, 22, 00), date.Day},
		{"16:10", tt(2015, 2, 15, 16, 10), date.Day},
		{"2014", tt(2014, 2, 15), date.Year},
		{"2008", tt(2008, 2, 15), date.Year},
		// Subscript and superscript dates
		{"²⁰¹⁵", tt(2015, 2, 15), date.Year},
		{"²⁹/⁰⁵/²⁰¹⁵", tt(2015, 5, 29), date.Day},
		{"₁₅/₀₂/₂₀₂₀", tt(2020, 2, 15), date.Day},
		{"₃₁ December", tt(2015, 12, 31), date.Day},
	}

	// Prepare config
	cfg := Configuration{CurrentTime: tt(2015, 2, 15, 15, 30)}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%v) => \"%s\"",
			test.Text, test.ExpectedPeriod,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Parse text
		dt, _ := Parse(&cfg, test.Text)
		if passed := assertResult(t, dt, test.ExpectedTime, test.ExpectedPeriod, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_customOrder(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text         string
		ExpectedTime time.Time
		DateOrder    string
	}

	tests := []testScenario{
		// Basic text
		{"15-12-18 06:00", tt(2015, 12, 18, 6, 0), "YMD"},
		{"15-18-12 06:00", tt(2015, 12, 18, 6, 0), "YDM"},
		{"10-11-12 06:00", tt(2012, 10, 11, 6, 0), "MDY"},
		{"10-11-12 06:00", tt(2011, 10, 12, 6, 0), "MYD"},
		{"10-11-12 06:00", tt(2011, 12, 10, 6, 0), "DYM"},
		{"15-12-18 06:00", tt(2018, 12, 15, 6, 0), "DMY"},
		{"12/09/08 04:23:15.567", tt(2008, 9, 12, 4, 23, 15, 567000), "DMY"},
		{"10/9/1914 03:07:09.788888 pm", tt(1914, 10, 9, 15, 7, 9, 788888), "MDY"},
		{"1-8-09 07:12:49 AM", tt(2009, 1, 8, 7, 12, 49), "MDY"},
		{"2016 july 13.", tt(2016, 7, 13, 0, 0), "YMD"},
		{"16 july 13.", tt(2016, 7, 13, 0, 0), "YMD"},
		{"Sunday 23 May 1856 12:09:08 AM", tt(1856, 5, 23, 0, 9, 8), "DMY"},

		// Nospace
		{"201508", tt(2015, 8, 20, 0, 0), "DYM"},
		{"201508", tt(2020, 8, 15, 0, 0), "YDM"},
		{"201108", tt(2008, 11, 20, 0, 0), "DMY"},
	}

	// Prepare config
	cfg := Configuration{PreferConfigDateOrder: true}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\" (%s) => \"%s\"",
			test.Text, test.DateOrder,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Parse text
		cfg.DateOrder = test.DateOrder
		dt, _ := Parse(&cfg, test.Text)
		if passed := assertResult(t, dt, test.ExpectedTime, date.Day, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_onlySeparatorTokens(t *testing.T) {
	// Prepare scenarios
	tests := []string{
		"::",
		"..",
		"  ",
		"--",
		"//",
		"++",
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Parse text
		dt, _ := Parse(nil, test)
		if passed := assert.True(t, dt.IsZero()); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_successWhenSkipTokensSpecified(t *testing.T) {
	cfg := Configuration{
		CurrentTime: tt(2015, 2, 12),
		SkipTokens:  []string{"de"},
	}

	dt, err := Parse(&cfg, "24 April 2012 de")
	assert.NoError(t, err)
	assert.Equal(t, tt(2012, 4, 24), dt.Time)
}

func TestParser_Parse_format(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text         string
		ExpectedTime time.Time
		Format       string
	}

	tests := []testScenario{
		{"14 giu 13", tt(2014, 6, 13), "06 January 02"},
		{"14_luglio_15", tt(2014, 7, 15), "06_January_02"},
		{"14_LUGLIO_15", tt(2014, 7, 15), "06_January_02"},
		{"10.01.2016, 20:35", tt(2016, 1, 10, 20, 35), "02.01.2006 15:04"},
		{"2014/11/17 14:56 EDT", tt(2014, 11, 17, 18, 56), "2006/01/02 15:04 MST"},
	}

	// Prepare parser
	parser := Parser{ParserTypes: []ParserType{CustomFormat}}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\", format \"%s\" => \"%s\"",
			test.Text, test.Format,
			test.ExpectedTime.Format("2006-01-02 15:04:05.999999"))

		// Parse text
		dt, _ := parser.Parse(nil, test.Text, test.Format)
		if passed := assertResult(t, dt, test.ExpectedTime, date.Day, message); !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_detectLocale(t *testing.T) {
	// Prepare scenarios
	type testScenario struct {
		Text           string
		ExpectedLocale string
	}

	tests := []testScenario{
		{"11 Marzo, 2014", "es"},
		{"13 Septiembre, 2014", "es"},
		{"Сегодня", "ru"},
		{"Avant-hier", "fr"},
		{"Anteontem", "pt"},
		{"ธันวาคม 11, 2014, 08:55:08 PM", "th"},
		{"Anteontem", "pt"},
		{"14 aprilie 2014", "ro"},
		{"11 Ağustos, 2014", "tr"},
		{"pon 16. čer 2014 10:07:43", "cs"},
		{"24 януари 2015г.", "bg"},
		{"Mañana", "es"},
		{"2020-05-01", "en"},
		{"Понедельник", "ru"},
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("\"%s\"  => \"%s\"", test.Text, test.ExpectedLocale)

		// Parse text
		dt, _ := Parse(nil, test.Text)
		passed := assert.False(t, dt.IsZero(), message)
		if passed {
			passed = assert.Equal(t, test.ExpectedLocale, dt.Locale, message)
		}

		if !passed {
			fmt.Println("\t\t\tGOT:", dt)
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}

func TestParser_Parse_customParser(t *testing.T) {
	// Prepare parser
	parser := Parser{}
	var dt date.Date

	// Language is DE, config specified
	cfg := Configuration{PreferredDayOfMonth: First}
	parser.Languages = []string{"de"}
	dt, _ = parser.Parse(&cfg, "10.1.2019")
	assert.Equal(t, tt(2019, 1, 10), dt.Time)

	// Languages is DE, config not specified
	parser.Languages = []string{"de"}
	dt, _ = parser.Parse(nil, "10.1.2019")
	assert.Equal(t, tt(2019, 1, 10), dt.Time)

	// Languages not specified, date order in config specified
	parser.Languages = nil
	cfg = Configuration{DateOrder: "MDY"}
	dt, _ = parser.Parse(&cfg, "10.1.2019")
	assert.Equal(t, tt(2019, 10, 1), dt.Time)

	// Languages is TH, prefer date order in config
	parser.Languages = []string{"th"}
	cfg = Configuration{DateOrder: "MDY", PreferConfigDateOrder: true}

	dt, _ = parser.Parse(&cfg, "03/11/2559 05:13")
	assert.Equal(t, tt(2559, 3, 11, 5, 13), dt.Time)

	dt, _ = parser.Parse(&cfg, "03/15/2559 05:13")
	assert.Equal(t, tt(2559, 3, 15, 5, 13), dt.Time)

	// Languages set to PT and text is PT as well
	parser.Languages = []string{"pt"}
	dt, _ = parser.Parse(nil, "24 de Janeiro de 2014")
	assert.Equal(t, tt(2014, 1, 24), dt.Time)

	// Languages set to PT but text is EN
	parser.Languages = []string{"pt"}
	dt, _ = parser.Parse(nil, "24 January, 2014")
	assert.True(t, dt.IsZero())

	// Locales set to pt-TL and text is matched
	parser.Languages = nil
	parser.Locales = []string{"pt-TL"}
	dt, _ = parser.Parse(nil, "24 de Janeiro de 2014")
	assert.Equal(t, tt(2014, 1, 24), dt.Time)

	// Locales and text not matched
	parser.Locales = []string{"pt-AO"}
	dt, _ = parser.Parse(nil, "24 January, 2014")
	assert.True(t, dt.IsZero())

	// Restricted languages
	parser.Locales = nil
	parser.Languages = []string{"en", "de", "fr", "it", "pt", "nl", "ro", "es", "ru"}

	dt, _ = parser.Parse(nil, "07/07/2014") // any language
	assert.Equal(t, tt(2014, 7, 7), dt.Time)

	dt, _ = parser.Parse(nil, "07.jul.2014 | 12:52") // en, es, pt, nl
	assert.Equal(t, tt(2014, 7, 7, 12, 52), dt.Time)
	assert.Equal(t, "en", dt.Locale)

	dt, _ = parser.Parse(nil, "07.ago.2014 | 12:52") // es, it, pt
	assert.Equal(t, tt(2014, 8, 7, 12, 52), dt.Time)
	assert.Equal(t, "es", dt.Locale)

	dt, _ = parser.Parse(nil, "07.feb.2014 | 12:52") // en, de, es, it, nl, ro
	assert.Equal(t, tt(2014, 2, 7, 12, 52), dt.Time)
	assert.Equal(t, "en", dt.Locale)

	dt, _ = parser.Parse(nil, "07.ene.2014 | 12:52") // es
	assert.Equal(t, tt(2014, 1, 7, 12, 52), dt.Time)
	assert.Equal(t, "es", dt.Locale)

	// Try previous locales
	parser = Parser{TryPreviousLocales: true}
	dt, _ = parser.Parse(nil, "Mañana")
	assert.Equal(t, "es", dt.Locale)
	dt, _ = parser.Parse(nil, "2020-05-01")
	assert.Equal(t, "es", dt.Locale) // is en, but follow the previous locale
	dt, _ = parser.Parse(nil, "Понедельник")
	assert.Equal(t, "ru", dt.Locale)

	// Used locale order is deterministic
	usedLocaleNames := make([]string, len(parser.usedLocales))
	for i, ld := range parser.usedLocales {
		usedLocaleNames[i] = ld.Name
	}
	assert.Equal(t, []string{"es", "ru"}, usedLocaleNames)
}

func tt(Y, M, D int, times ...int) time.Time {
	nTimes := len(times)
	var H, m, s, ms int

	if nTimes >= 1 {
		H = times[0]
	}
	if nTimes >= 2 {
		m = times[1]
	}
	if nTimes >= 3 {
		s = times[2]
	}
	if nTimes >= 4 {
		ms = times[3]
	}

	return time.Date(Y, time.Month(M), D, H, m, s, ms*1000, time.UTC)
}

func assertResult(t *testing.T, dt date.Date, expectedTime time.Time, expectedPeriod date.Period, messages ...string) bool {
	var message string
	if len(messages) > 0 {
		message = messages[0]
	}

	passed := assert.False(t, dt.IsZero(), message)
	if passed {
		diff := expectedTime.Sub(dt.Time)
		passed = assert.Zero(t, diff, message)
	}

	if passed {
		passed = assert.Equal(t, expectedPeriod, dt.Period, message)
	}

	return passed
}
