package language

import (
	"fmt"
	"testing"

	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestIsApplicable(t *testing.T) {
	type testScenario struct {
		Locale        string
		Text          string
		Expected      bool
		StripTimezone bool
	}

	ts := func(locale string, text string, expected bool, stripTimezone ...bool) testScenario {
		s := testScenario{Locale: locale, Text: text, Expected: expected}
		if len(stripTimezone) > 0 {
			s.StripTimezone = stripTimezone[0]
		}
		return s
	}

	stripTimezone := true
	tests := []testScenario{
		// Should be applicable
		ts("en", "17th October, 2034 @ 01:08 am PDT", true, stripTimezone),
		ts("en", "#@Sept#04#2014", true),
		ts("en", "2014-12-13T00:11:00Z", true),
		ts("de", "Donnerstag, 8. Januar 2015 um 07:17", true),
		ts("da", "Torsdag, 8. januar 2015 kl. 07:17", true),
		ts("ru", "8 января 2015 г. в 9:10", true),
		ts("cs", "Pondělí v 22:29", true),
		ts("nl", "woensdag 7 januari om 21:32", true),
		ts("ro", "8 Ianuarie 2015 la 13:33", true),
		ts("ar", "ساعتين", true),
		ts("tr", "3 hafta", true),
		ts("th", "17 เดือนมิถุนายน", true),
		ts("pl", "przedwczoraj", true),
		ts("fa", "ژانویه 8, 2015، ساعت 15:46", true),
		ts("vi", "2 tuần 3 ngày", true),
		ts("tl", "Hulyo 3, 2015 7:00 pm", true),
		ts("be", "3 верасня 2015 г. у 11:10", true),
		ts("id", "01 Agustus 2015 18:23", true),
		ts("he", "6 לדצמבר 1973", true),
		ts("bn", "3 সপ্তাহ", true),

		// Should be NOT applicable
		ts("ru", "08.haziran.2014, 11:07", false),
		ts("ar", "6 دقیقه", false),
		ts("fa", "ساعتين", false),
		ts("cs", "3 hafta", false),
	}

	// Prepare config
	cfg := &setting.Configuration{
		SkipTokens: []string{"t"},
	}

	nFailed := 0
	for _, test := range tests {
		// Prepare log message
		message := fmt.Sprintf("%s, \"%s\"", test.Locale, test.Text)

		// Load locale
		ld, err := GetLocale(test.Locale)
		assert.Nil(t, err, message)

		// Make sure it's applicable
		isApplicable := IsApplicable(cfg, ld, test.Text, test.StripTimezone)
		passed := assert.Equal(t, test.Expected, isApplicable, message)
		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
