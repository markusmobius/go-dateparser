package dateparser

import (
	"fmt"
	"sync"
	"time"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/language"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
	"github.com/markusmobius/go-dateparser/internal/parser/formatted"
	"github.com/markusmobius/go-dateparser/internal/parser/nospace"
	"github.com/markusmobius/go-dateparser/internal/parser/relative"
	"github.com/markusmobius/go-dateparser/internal/parser/timestamp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

// Parser is object that handles language detection, translation and subsequent
// generic parsing of string representing date and/or time.
type Parser struct {
	sync.Mutex

	// Locales is a list of locale codes, e.g. ['fr-PF', 'qu-EC', 'af-NA'].
	// The parser uses only these locales to translate date string.
	Locales []string
	// Languages is a list of language codes, e.g. ['en', 'es', 'zh-Hant']. If
	// locales are not given, languages and region are used to construct locales
	// for translation.
	Languages []string
	// Region is a region code, e.g. 'IN', '001', 'NE'. If locales are not given,
	// languages and region are used to construct locales for translation.
	Region string
	// If true, locales previously used to translate date are tried first.
	TryPreviousLocales bool
	// If true, locales are tried for translation of date string in the order in
	// which they are given.
	UseGivenOrder bool
	// DetectLanguagesFunction is a function for language detection that takes
	// as input a `text` and returns a list of detected language codes. Note:
	// this function is only used if `languages` and `locales` are not provided.
	DetectLanguagesFunction func(string) []string
	// ParserTypes is a list of types of parsers to try, allowing to customize which parsers are tried
	// against the input date string, and in which order they are tried. By default it will use
	// all parser in following order: `Timestamp`, `RelativeTime`, `CustomFormat`, `AbsoluteTime`,
	// and finally `NoSpacesTime`.
	ParserTypes []ParserType

	usedLocales        []*data.LocaleData
	usedLocalesTracker strutil.Dict
	uniqueCharsets     map[string][]rune
}

// ParserType is the variable to specify which type of parser that will be used.
type ParserType uint8

const (
	// Timestamp is parser to parse Unix timestamp.
	Timestamp ParserType = iota
	// RelativeTime is parser to parse date string with relative value like
	// "1 year, 2 months ago" and "3 hours, 50 minutes ago".
	RelativeTime
	// CustomFormat is parser to parse a date string with custom formats.
	CustomFormat
	// AbsoluteTime is parser to parse date string with absolute value like
	// "12 August 2021" and "23 January, 15:10:01".
	AbsoluteTime
	// NoSpacesTime is parser to parse date string that written without spaces,
	// for example 2021-10-11 that written as 20211011.
	NoSpacesTime
)

// Parse parses string representing date and/or time in recognizable localized formats,
// using the default Parser. Useful for quick use.
func Parse(cfg *Configuration, str string, formats ...string) (date.Date, error) {
	return (&Parser{}).Parse(cfg, str, formats...)
}

// Parse parses string representing date and/or time in recognizable localized formats.
// Supports parsing multiple languages.
func (p *Parser) Parse(cfg *Configuration, str string, formats ...string) (date.Date, error) {
	// Lock mutex
	p.Lock()
	defer p.Unlock()

	// Validate and initiate parsers
	for _, parser := range p.ParserTypes {
		if parser < 0 || parser > NoSpacesTime {
			return date.Date{}, fmt.Errorf("invalid parser type: %d", parser)
		}
	}

	if len(p.ParserTypes) == 0 {
		p.ParserTypes = []ParserType{Timestamp, RelativeTime, CustomFormat, AbsoluteTime, NoSpacesTime}
	}

	// Initiate and validate config
	if cfg == nil {
		cfg = &Configuration{}
	}

	cfg = cfg.initiate()
	err := cfg.validate()
	if err != nil {
		return date.Date{}, fmt.Errorf("config error: %w", err)
	}

	// Convert config to internal config
	iCfg := cfg.toInternalConfig()

	// Try to parse with specified formats
	var dt date.Date
	if dt = formatted.Parse(iCfg, str, formats...); !dt.IsZero() {
		return dt, nil
	}

	// Sanitize string
	originalStr := str
	str = strutil.SanitizeDate(str)

	// Find the suitable locales for this string
	locales, err := p.getApplicableLocales(iCfg, str)
	if err != nil {
		return date.Date{}, err
	}

	// Process each locale
	for _, locale := range locales {
		// Create locale specific config
		lCfg := iCfg.Clone()
		if lCfg.DateOrder == "" || !lCfg.PreferConfigDateOrder {
			lCfg.DateOrder = locale.DateOrder
		}

		// Translate string
		translation := language.Translate(&lCfg, locale, str, false)
		translationWithFormat := language.Translate(&lCfg, locale, str, true)

		for _, parserType := range p.ParserTypes {
			switch parserType {
			case Timestamp:
				dt = timestamp.Parse(&lCfg, str)
			case RelativeTime:
				dt = relative.Parse(&lCfg, translation)
			case CustomFormat:
				dt = formatted.Parse(&lCfg, translationWithFormat, formats...)
			case AbsoluteTime:
				if t, tz := p.stripBracesAndTimezones(translation); t != "" {
					dt, _ = absolute.Parse(&lCfg, t)
					dt = p.applyTimezone(dt, tz)
				}
			case NoSpacesTime:
				if t, tz := p.stripBracesAndTimezones(translation); t != "" {
					dt, _ = nospace.Parse(&lCfg, t)
					dt = p.applyTimezone(dt, tz)
				}
			}

			if !dt.IsZero() {
				dt.Locale = locale.Name
				p.saveUsedLocale(locale)
				return dt, nil
			}
		}
	}

	return date.Date{}, fmt.Errorf("failed to parse \"%s\": unknown format", originalStr)
}

func (p *Parser) getApplicableLocales(iCfg *setting.Configuration, str string) ([]*data.LocaleData, error) {
	// Prepare results
	var results []*data.LocaleData
	resultTracker := strutil.NewDict()

	// Normalize and prepare date strings
	str = strutil.NormalizeString(str)
	dateStrings := []string{str}
	if poppedTz, _ := timezone.PopTzOffset(str); poppedTz != str {
		dateStrings = append(dateStrings, poppedTz)
	}

	// Fetch previously used locales first
	if p.TryPreviousLocales {
		for _, usedLocale := range p.usedLocales {
			for _, ds := range dateStrings {
				if p.localeIsApplicable(iCfg, usedLocale, ds) {
					results = append(results, usedLocale)
					resultTracker.Add(usedLocale.Name)
					break
				}
			}
		}
	}

	// If specified, use external detector to fetch languages
	languages := append([]string{}, p.Languages...)
	if p.DetectLanguagesFunction != nil && len(p.Locales) == 0 && len(languages) == 0 {
		detectionResults := p.DetectLanguagesFunction(str)
		languages = append(languages, detectionResults...)
	}

	// Load locales
	locales, err := language.GetLocales(p.Locales, languages, p.Region, p.UseGivenOrder, false)
	if err != nil && err != language.ErrNotFound {
		return nil, err
	}

	for _, locale := range locales {
		if resultTracker.Contain(locale.Name) {
			continue
		}

		// Check if locale is applicable
		var isApplicable bool
		for _, ds := range dateStrings {
			if p.localeIsApplicable(iCfg, locale, ds) {
				isApplicable = true
				break
			}
		}

		if isApplicable {
			results = append(results, locale)
			resultTracker.Add(locale.Name)
		}
	}

	// Finally, append locales of default languages
	if len(iCfg.DefaultLanguages) > 0 {
		locales, _ := language.GetLocales(nil, iCfg.DefaultLanguages, p.Region, p.UseGivenOrder, false)
		for _, locale := range locales {
			if !resultTracker.Contain(locale.Name) {
				results = append(results, locale)
			}
		}
	}

	return results, nil
}

func (p *Parser) saveUsedLocale(ld *data.LocaleData) {
	if !p.TryPreviousLocales {
		return
	}

	if p.usedLocalesTracker == nil {
		p.usedLocalesTracker = strutil.NewDict()
	}

	if !p.usedLocalesTracker.Contain(ld.Name) {
		p.usedLocalesTracker.Add(ld.Name)
		p.usedLocales = append(p.usedLocales, ld)
	}
}

func (p *Parser) localeIsApplicable(iCfg *setting.Configuration, ld *data.LocaleData, s string) bool {
	return language.IsApplicable(iCfg, ld, s, false)
}

func (p *Parser) stripBracesAndTimezones(s string) (string, timezone.TimezoneOffsetData) {
	s = strutil.StripBraces(s)
	return timezone.PopTzOffset(s)
}

func (p *Parser) applyTimezone(dt date.Date, tz timezone.TimezoneOffsetData) date.Date {
	if dt.IsZero() || tz.IsZero() {
		return dt
	}

	dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(),
		dt.Time.Hour(), dt.Time.Minute(), dt.Time.Second(), dt.Time.Nanosecond(),
		time.FixedZone(tz.Name, tz.Offset))

	return dt
}
