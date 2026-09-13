package dateparser

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/language"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/stretchr/testify/require"
)

func TestPythonCalendarPublicDefaults(t *testing.T) {
	cfg := &Configuration{CurrentTime: time.Date(2024, 3, 20, 12, 0, 0, 0, time.UTC)}
	for _, calendar := range []struct {
		name      string
		parse     func(*Configuration, string) (date.Date, error)
		ambiguous string
		clock     string
	}{
		{"jalali", ParseJalali, "2054-06-24", "2065-03-20 05:30"},
		{"hijri", ParseHijri, "2012-02-25", "2022-07-30 05:30"},
	} {
		t.Run(calendar.name, func(t *testing.T) {
			parsed, err := calendar.parse(cfg, "04-03-1433")
			require.NoError(t, err)
			require.Equal(t, calendar.ambiguous, parsed.Time.Format("2006-01-02"))
			parsed, err = calendar.parse(cfg, "1/1/1444 +05:30")
			require.NoError(t, err)
			require.Equal(t, calendar.clock, parsed.Time.Format("2006-01-02 15:04"))
			for _, input := range []string{"1/1/1444 UTC", "1/1/1444 GMT+05:30"} {
				_, err := calendar.parse(cfg, input)
				require.Error(t, err, input)
			}
		})
	}
	parsed, err := ParseJalali(cfg, "1 Farvardin 1403")
	require.NoError(t, err)
	require.Equal(t, "2024-03-20", parsed.Time.Format("2006-01-02"))
}

func TestPythonSearchChineseDates(t *testing.T) {
	input := "2024-02-29 12:34:56 UTC"
	for _, language := range []string{"zh", "zh-Hans", "zh-Hant"} {
		t.Run(language, func(t *testing.T) {
			parser := &Parser{}
			matches, err := parser.SearchWithLanguage(&Configuration{}, language, input)
			require.NoError(t, err)
			require.Len(t, matches, 1)
			require.Equal(t, input, matches[0].Text)
			require.Equal(t, "2024-02-29T12:34:56Z", matches[0].Date.Time.Format(time.RFC3339))
		})
	}
	require.NotPanics(t, func() {
		_, _ = (&Parser{}).SearchWithLanguage(&Configuration{}, "yue", input)
	})
}

func TestPythonSearchOverflowDoesNotWrap(t *testing.T) {
	for _, strategy := range []string{"split", "ngram"} {
		cfg := &Configuration{
			CurrentTime: time.Date(2024, 3, 20, 12, 0, 0, 0, time.UTC),
			Languages:   []string{"en"}, SearchStrategy: strategy, ReturnTimeSpan: true,
		}
		for _, count := range []string{"99999999999999999999", "999999999999999999", "9223372036854775807"} {
			_, matches, err := (&Parser{}).Search(cfg, "Messages received past "+count+" weeks")
			require.NoError(t, err)
			require.Empty(t, matches, strategy+": "+count)
		}
	}
}

func TestPythonLongRelativeClock(t *testing.T) {
	parsed, err := Parse(&Configuration{
		CurrentTime: time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC),
		Languages:   []string{"en"},
	}, "3504000 hours ago")
	require.NoError(t, err)
	require.Equal(t, "1600-04-07T12:00:00Z", parsed.Time.Format(time.RFC3339))
}

func TestPythonLanguageEvidence(t *testing.T) {
	for _, test := range []struct {
		code  string
		input string
		score [2]int
	}{
		{"en", "meeting on 2024-02-29", [2]int{0, 4}},
		{"nn", "meeting on 2024-02-29", [2]int{1, 3}},
		{"fr", "la reunion a eu lieu le 12 mars 2024.", [2]int{2, 2}},
		{"nn", "la reunion a eu lieu le 12 mars 2024.", [2]int{2, 2}},
	} {
		locale, exists := data.GetLocaleData(test.code)
		require.True(t, exists)
		words, skipped := language.CountApplicability(&setting.Configuration{SkipTokens: []string{"t"}}, locale, test.input, false)
		require.Equal(t, test.score, [2]int{words, skipped}, "%s %q", test.code, test.input)
	}
}

type pythonFeatureConfiguration struct {
	CurrentTime           string               `json:"current_time"`
	DateOrder             string               `json:"date_order"`
	PreferredDateSource   PreferredDateSource  `json:"preferred_date_source"`
	PreferredDayOfMonth   PreferredDayOfMonth  `json:"preferred_day_of_month"`
	PreferredMonthOfYear  PreferredMonthOfYear `json:"preferred_month_of_year"`
	StrictParsing         bool                 `json:"strict_parsing"`
	RequiredParts         []string             `json:"required_parts"`
	ReturnTimeAsPeriod    bool                 `json:"return_time_as_period"`
	IgnoreSurroundingText bool                 `json:"ignore_surrounding_text"`
	Languages             []string             `json:"languages"`
	DefaultLanguages      []string             `json:"default_languages"`
	SkipTokens            []string             `json:"skip_tokens"`
	UseGivenOrder         bool                 `json:"use_given_order"`
	SearchStrategy        string               `json:"search_strategy"`
	ReturnTimeSpan        bool                 `json:"return_time_span"`
	DefaultStartOfWeek    string               `json:"default_start_of_week"`
	DefaultDaysInMonth    int                  `json:"default_days_in_month"`
	PreserveEndOfMonth    bool                 `json:"preserve_end_of_month"`
}

func (configuration pythonFeatureConfiguration) public(t *testing.T) *Configuration {
	t.Helper()
	reference, err := time.Parse(time.RFC3339Nano, configuration.CurrentTime)
	require.NoError(t, err)
	result := &Configuration{
		CurrentTime: reference, PreferredDateSource: configuration.PreferredDateSource,
		PreferredDayOfMonth: configuration.PreferredDayOfMonth, PreferredMonthOfYear: configuration.PreferredMonthOfYear,
		StrictParsing: configuration.StrictParsing, RequiredParts: configuration.RequiredParts,
		ReturnTimeAsPeriod: configuration.ReturnTimeAsPeriod, IgnoreSurroundingText: configuration.IgnoreSurroundingText,
		Languages: configuration.Languages, DefaultLanguages: configuration.DefaultLanguages,
		SkipTokens: configuration.SkipTokens, UseGivenOrder: configuration.UseGivenOrder,
		SearchStrategy: configuration.SearchStrategy, ReturnTimeSpan: configuration.ReturnTimeSpan,
		DefaultStartOfWeek: configuration.DefaultStartOfWeek, DefaultDaysInMonth: configuration.DefaultDaysInMonth,
		PreserveEndOfMonth: configuration.PreserveEndOfMonth,
	}
	if configuration.DateOrder != "" {
		result.DateOrder = func(string) string { return configuration.DateOrder }
	}
	return result
}

type pythonFeatureDate struct {
	UnixSeconds int64  `json:"unix_seconds"`
	Nanosecond  int    `json:"nanosecond"`
	Offset      int    `json:"offset"`
	Period      string `json:"period"`
}

type pythonFeature struct {
	ID                string                     `json:"id"`
	Stage             string                     `json:"stage"`
	Input             string                     `json:"input"`
	Language          string                     `json:"language"`
	Detected          string                     `json:"detected"`
	Configuration     pythonFeatureConfiguration `json:"configuration"`
	Error             string                     `json:"error"`
	KnownDifference   string                     `json:"known_difference"`
	DetectorLanguages *[]string                  `json:"detector_languages"`
	DetectionInputs   []string                   `json:"detection_inputs"`
	Matches           []struct {
		Text string            `json:"text"`
		Date pythonFeatureDate `json:"date"`
	} `json:"matches"`
}

func readPythonFeatures(t *testing.T) []pythonFeature {
	t.Helper()
	contents, err := os.ReadFile("testdata/python-features.json")
	require.NoError(t, err)
	var fixture struct {
		Reference struct {
			Versions map[string]string `json:"versions"`
		} `json:"reference"`
		CalendarDataSHA256 string          `json:"calendar_data_sha256"`
		Cases              []pythonFeature `json:"cases"`
	}
	require.NoError(t, json.Unmarshal(contents, &fixture))
	require.Equal(t, "1.4.3", fixture.Reference.Versions["dateparser"])
	calendarData, err := os.ReadFile("internal/parser/calendars/data.json")
	require.NoError(t, err)
	require.Equal(t, fixture.CalendarDataSHA256, fmt.Sprintf("%x", sha256.Sum256(calendarData)))
	return fixture.Cases
}

func TestPythonCalendarFixture(t *testing.T) {
	differences, checked := 0, 0
	for _, test := range readPythonFeatures(t) {
		if test.Stage != "jalali" && test.Stage != "hijri" {
			continue
		}
		checked++
		parse := ParseJalali
		if test.Stage == "hijri" {
			parse = ParseHijri
		} else {
			require.Equal(t, "jalali", test.Stage)
		}
		parsed, parseErr := parse(test.Configuration.public(t), test.Input)
		matches := test.Error != "" && (parseErr != nil || parsed.IsZero())
		if test.Error == "" && parseErr == nil && len(test.Matches) == 1 {
			expected := test.Matches[0].Date
			period := parsed.Period.String()
			periodMatches := period == expected.Period || expected.Period == "Time" && (period == "Hour" || period == "Minute" || period == "Second")
			matches = !parsed.IsZero() && parsed.Time.Unix() == expected.UnixSeconds && parsed.Time.Nanosecond() == expected.Nanosecond && periodMatches
		}
		if !matches {
			differences++
			if differences <= 20 {
				t.Logf("%s %q: got %v (%v), Python matches=%+v error=%q", test.ID, test.Input, parsed, parseErr, test.Matches, test.Error)
			}
		}
	}
	if differences != 0 {
		t.Fatalf("%d of %d Python calendar cases differed", differences, checked)
	}
	t.Logf("Matched %d Python calendar cases", checked)
}

func TestPythonSearchFixture(t *testing.T) {
	differences, checked, safety := 0, 0, 0
	for _, test := range readPythonFeatures(t) {
		if test.Stage != "search" && test.Stage != "search_with_language" {
			continue
		}
		parser := &Parser{}
		var detectionInputs []string
		if test.DetectorLanguages != nil {
			parser.DetectLanguagesFunction = func(input string) []string {
				detectionInputs = append(detectionInputs, input)
				return *test.DetectorLanguages
			}
		}
		var detected string
		var matches []SearchResult
		var err error
		if test.Stage == "search" {
			detected, matches, err = parser.Search(test.Configuration.public(t), test.Input)
		} else {
			detected = test.Language
			matches, err = parser.SearchWithLanguage(test.Configuration.public(t), test.Language, test.Input)
		}
		if test.KnownDifference == "python-search-exception" {
			safety++
			continue
		}
		if test.KnownDifference == "python-relative-range" {
			safety++
			require.Empty(t, matches, test.ID)
			continue
		}
		checked++
		valid := len(matches) == len(test.Matches)
		if len(test.Matches) > 0 {
			valid = valid && err == nil && detected == test.Detected
		}
		if valid {
			for index, match := range matches {
				expected := test.Matches[index]
				_, offset := match.Date.Time.Zone()
				valid = valid && match.Text == expected.Text && match.Date.Time.Unix() == expected.Date.UnixSeconds &&
					match.Date.Time.Nanosecond() == expected.Date.Nanosecond && offset == expected.Date.Offset
			}
		}
		valid = valid && fmt.Sprint(detectionInputs) == fmt.Sprint(test.DetectionInputs)
		if !valid {
			differences++
			if differences <= 20 {
				t.Logf("%s %q: got %s %+v (%v), callbacks=%q; Python %s %+v error=%q callbacks=%q", test.ID, test.Input, detected, matches, err, detectionInputs, test.Detected, test.Matches, test.Error, test.DetectionInputs)
			}
		}
	}
	if differences != 0 {
		t.Fatalf("%d of %d ordinary Python search cases differed; %d Python exception inputs checked for safety", differences, checked, safety)
	}
	t.Logf("Matched %d Python search cases; %d Python exception inputs checked for safety", checked, safety)
}
