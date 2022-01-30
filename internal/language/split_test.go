package language

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	// Helper function
	sp := func(items ...string) []string { return items }

	// Prepare scenarios
	type testScenario struct {
		Locale   string
		Text     string
		Expected []string
	}

	tests := []testScenario{
		{"pt", "sexta-feira, 10 de junho de 2014 14:52", sp("sexta-feira", " ", "10", " ", "de", " ", "junho", " ", "de", " ", "2014", " ", "14", ":", "52")},
		{"it", "14_luglio_15", sp("14", "luglio", "15")},
		{"zh", "1年11个月", sp("1", "年", "11", "个月")},
		{"zh", "1年11個月", sp("1", "年", "11", "個月")},
		{"tr", "2 saat önce", sp("2 saat once")},
		{"fr", "il ya environ 23 heures'", sp("il ya", " ", "environ", " ", "23", " ", "heures")},
		{"de", "Gestern um 04:41", sp("gestern", " ", "um", " ", "04", ":", "41")},
		{"de", "Donnerstag, 8. Januar 2015 um 07:17", sp("donnerstag", " ", "8", ".", " ", "januar", " ", "2015", " ", "um", " ", "07", ":", "17")},
		{"ru", "8 января 2015 г. в 9:10", sp("8", " ", "января", " ", "2015", " ", "г", ".", " ", "в", " ", "9", ":", "10")},
		{"cs", "6. leden 2015 v 22:29", sp("6", ".", " ", "leden", " ", "2015", " ", "v", " ", "22", ":", "29")},
		{"nl", "woensdag 7 januari 2015 om 21:32", sp("woensdag", " ", "7", " ", "januari", " ", "2015", " ", "om", " ", "21", ":", "32")},
		{"ro", "8 Ianuarie 2015 la 13:33", sp("8", " ", "ianuarie", " ", "2015", " ", "la", " ", "13", ":", "33")},
		{"ar", "8 يناير، 2015، الساعة 10:01 صباحاً", sp("8", " ", "يناير", " ", "2015", "الساعة", " ", "10", ":", "01", " ", "صباحا")},
		{"th", "8 มกราคม 2015 เวลา 12:22 น.", sp("8", " ", "มกราคม", " ", "2015", " ", "เวลา", " ", "12", ":", "22", " ", "น.")},
		{"pl", "8 stycznia 2015 o 10:19", sp("8", " ", "stycznia", " ", "2015", " ", "o", " ", "10", ":", "19")},
		{"vi", "Thứ Năm, ngày 8 tháng 1 năm 2015", sp("thu nam", " ", "ngay", " ", "8", " ", "thang 1", " ", "nam", " ", "2015")},
		{"tl", "Biyernes Hulyo 3 2015", sp("biyernes", " ", "hulyo", " ", "3", " ", "2015")},
		{"be", "3 верасня 2015 г. у 11:10", sp("3", " ", "верасня", " ", "2015", " ", "г", ".", " ", "у", " ", "11", ":", "10")},
		{"id", "3 Juni 2015 13:05:46", sp("3", " ", "juni", " ", "2015", " ", "13", ":", "05", ":", "46")},
		{"he", "ה-21 לאוקטובר 2016 ב-15:00", sp("ה-", "21", " ", "לאוקטובר", " ", "2016", " ", "ב-", "15", ":", "00")},
		{"bn", "3 জুন 2015 13:05:46", sp("3", " ", "জন", " ", "2015", " ", "13", ":", "05", ":", "46")},
		{"hi", "13 मार्च 2013 11:15:09", sp("13", " ", "मारच", " ", "2013", " ", "11", ":", "15", ":", "09")},
		{"mgo", "aneg 5 12 iməg àdùmbə̀ŋ 2001 09:14 pm", sp("aneg 5", " ", "12", " ", "iməg adumbəŋ", " ", "2001", " ", "09", ":", "14", " ", "pm")},
		{"qu", "2 kapaq raymi 1998 domingo", sp("2", " ", "kapaq raymi", " ", "1998", " ", "domingo")},
		{"os", "24 сахаты размӕ 10:09 ӕмбисбоны размӕ", sp("24 сахаты размӕ", " ", "10", ":", "09", " ", "ӕмбисбоны размӕ")},
		{"pa", "25 ਘੰਟੇ ਪਹਿਲਾਂ 10:08 ਬਾਦੁ", sp("25 ਘਟ ਪਹਿਲਾ", " ", "10", ":", "08", " ", "ਬਾਦ")},
		{"en", "25_April_2008", sp("25", "april", "2008")},
		{"af", "hierdie uur 10:19 vm", sp("hierdie uur", " ", "10", ":", "19", " ", "vm")},
		{"rof", "7 mweri wa kaana 1998 12:09 kang'ama", sp("7", " ", "mweri wa kaana", " ", "1998", " ", "12", ":", "09", " ", "kang'ama")},
		{"saq", "14 lapa le tomon obo 2098 ong", sp("14", " ", "lapa le tomon obo", " ", "2098", " ", "ong")},
		{"wae", "cor 6 wučä 09:19 pm", sp("cor 6 wuca", " ", "09", ":", "19", " ", "pm")},
		{"naq", "13 ǃkhanǀgôab 1887", sp("13", " ", "ǃkhanǀgoab", " ", "1887")},
	}

	// Start test
	nFailed := 0
	skippedTokens := strutil.NewDict("t")
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("%s, \"%s\"", test.Locale, test.Text)

		// Load locale
		ld, err := GetLocale(test.Locale)
		assert.Nil(t, err, message)

		// Normalize text
		str := strutil.NormalizeString(test.Text)
		str = digit.NormalizeString(str)

		// Split text
		result := Split(ld, str, false, skippedTokens)
		passed := assert.Equal(t, test.Expected, result, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
