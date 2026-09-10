package language

import (
	"fmt"
	"maps"
	"math/rand/v2"
	"regexp/syntax"
	"slices"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/regexp"
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

func TestExactCombinedMatchers(t *testing.T) {
	random := rand.New(rand.NewPCG(1, 2))
	seen := map[string]bool{}
	var comparisons int
	for _, name := range slices.Sorted(maps.Keys(data.LocaleOrder)) {
		locale, exists := data.GetLocaleData(name)
		if !exists {
			t.Fatalf("missing locale %s", name)
		}
		if locale.RxExactCombined == nil {
			if locale.ExactCombinedMatcher != nil || locale.MatchExactCombined("today") {
				t.Fatalf("unexpected matcher for %s", name)
			}
			continue
		}
		if locale.ExactCombinedMatcher == nil {
			t.Fatalf("missing generated or inherited matcher for %s", name)
		}
		pattern := locale.RxExactCombined.String()
		if seen[pattern] {
			continue
		}
		seen[pattern] = true
		parsed, err := syntax.Parse(pattern, syntax.Perl)
		if err != nil {
			t.Fatal(err)
		}
		check := func(input string) {
			t.Helper()
			comparisons++
			expected, actual := locale.RxExactCombined.MatchString(input), locale.MatchExactCombined(input)
			if actual != expected {
				t.Fatalf("%s input=%q: regex=%v generated=%v", name, input, expected, actual)
			}
		}
		for _, input := range []string{"", "0", "123", "hello", "today", "in 3 days", "12 days ago", "\x00", "\xff", "\xc0\x80", "\ufffd", "\u212a", "\u017f", "\u03c2", "\r\n", " ", "2025-09-15"} {
			check(input)
		}
		for word := range locale.Translations {
			check(word)
		}
		for _, word := range locale.KnownWords {
			check(word)
		}
		for range 100 {
			input := sampleExactPattern(t, parsed, random)
			check(input)
			check(strings.ToUpper(input))
			check(strings.ToLower(input))
			for _, edge := range []string{"x", "\n", " ", "\x00", "\xff"} {
				check(edge + input)
				check(input + edge)
			}
			if len(input) > 0 {
				check(input[1:])
				check(input[:len(input)-1])
				position := random.IntN(len(input))
				check(input[:position] + "\x00" + input[position+1:])
				check(input[:position] + "\xff" + input[position+1:])
			}
		}
	}
	t.Logf("%d comparisons across %d unique generated matchers", comparisons, len(seen))
}

func TestExactCombinedMatcherFallback(t *testing.T) {
	locale := &data.LocaleData{RxExactCombined: regexp.MustCompile(`^\x{fffd}$`)}
	assert.True(t, locale.MatchExactCombined("\ufffd"))
	locale.ExactCombinedMatcher = func(input string) bool {
		if !utf8.ValidString(input) {
			t.Fatal("malformed UTF-8 must use the regex fallback")
		}
		return input == "\ufffd"
	}
	for _, input := range []string{"", "\ufffd", "\xff", "\xc0\x80", "\x00"} {
		assert.Equal(t, locale.RxExactCombined.MatchString(input), locale.MatchExactCombined(input), "%q", input)
	}
}

func sampleExactPattern(t *testing.T, parsed *syntax.Regexp, random *rand.Rand) string {
	t.Helper()
	switch parsed.Op {
	case syntax.OpBeginText, syntax.OpEndText, syntax.OpEmptyMatch:
		return ""
	case syntax.OpLiteral:
		values := slices.Clone(parsed.Rune)
		if parsed.Flags&syntax.FoldCase != 0 {
			for index := range values {
				for remaining := random.IntN(4); remaining > 0; remaining-- {
					values[index] = unicode.SimpleFold(values[index])
				}
			}
		}
		return string(values)
	case syntax.OpCharClass:
		index := random.IntN(len(parsed.Rune)/2) * 2
		return string(parsed.Rune[index] + rune(random.IntN(int(parsed.Rune[index+1]-parsed.Rune[index])+1)))
	case syntax.OpCapture:
		return sampleExactPattern(t, parsed.Sub[0], random)
	case syntax.OpAlternate:
		return sampleExactPattern(t, parsed.Sub[random.IntN(len(parsed.Sub))], random)
	case syntax.OpConcat:
		var result strings.Builder
		for _, child := range parsed.Sub {
			result.WriteString(sampleExactPattern(t, child, random))
		}
		return result.String()
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest, syntax.OpRepeat:
		minimum, maximum := 0, 4
		switch parsed.Op {
		case syntax.OpPlus:
			minimum = 1
		case syntax.OpQuest:
			maximum = 1
		case syntax.OpRepeat:
			minimum, maximum = parsed.Min, parsed.Max
			if maximum < 0 {
				maximum = minimum + 4
			}
		}
		count := minimum + random.IntN(maximum-minimum+1)
		var result strings.Builder
		for range count {
			result.WriteString(sampleExactPattern(t, parsed.Sub[0], random))
		}
		return result.String()
	default:
		t.Fatalf("unsupported sample operation %s", parsed.Op)
		return ""
	}
}
