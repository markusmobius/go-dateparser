package godateparser

import (
	"fmt"
	"sync"
	"time"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/language"
	"github.com/markusmobius/go-dateparser/internal/parser/absolute"
	"github.com/markusmobius/go-dateparser/internal/parser/date"
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

	// Config is the advanced settings to set the parser behavior.
	Config *Configuration
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

	usedLocales []*data.LocaleData
}

// Parse parses string representing date and/or time in recognizable localized formats.
// Supports parsing multiple languages.
func (p *Parser) Parse(str string, formats ...string) (Date, error) {
	// Initiate and validate config
	cfg := p.Config
	if cfg == nil {
		cfg = &Configuration{}
		p.Config = cfg
	}

	cfg.initiate()
	err := cfg.validate()
	if err != nil {
		return Date{}, err
	}

	// Convert config to internal config
	iCfg := cfg.toInternalConfig()

	// Try to parse with specified formats
	var iDt date.Date
	if iDt = formatted.Parse(&iCfg, str, formats...); !iDt.IsZero() {
		return dateFromInternal(iDt), nil
	}

	// Sanitize string
	originalStr := str
	str = strutil.SanitizeDate(str)

	// Find the suitable locales for this string
	locales, err := p.getApplicableLocales(&iCfg, str)
	if err != nil {
		return Date{}, err
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

		for _, parserType := range p.Config.Parsers {
			switch parserType {
			case Timestamp:
				iDt = timestamp.Parse(&lCfg, str)
			case RelativeTime:
				iDt = relative.Parse(&lCfg, translation)
			case CustomFormat:
				iDt = formatted.Parse(&lCfg, translationWithFormat, formats...)
			case AbsoluteTime:
				if t, tz := p.stripBracesAndTimezones(translation); t != "" {
					iDt, _ = absolute.Parse(&lCfg, t)
					iDt = p.applyTimezone(iDt, tz)
				}
			case NoSpacesTime:
				if t, tz := p.stripBracesAndTimezones(translation); t != "" {
					iDt, _ = nospace.Parse(&lCfg, t)
					iDt = p.applyTimezone(iDt, tz)
				}
			}

			if !iDt.IsZero() {
				iDt.Locale = locale.Name
				return dateFromInternal(iDt), nil
			}
		}
	}

	return Date{}, fmt.Errorf("failed to parse \"%s\": unknown format", originalStr)
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
	if len(p.Config.DefaultLanguages) > 0 {
		locales, _ := language.GetLocales(nil, p.Config.DefaultLanguages, p.Region, p.UseGivenOrder, false)
		for _, locale := range locales {
			if !resultTracker.Contain(locale.Name) {
				results = append(results, locale)
			}
		}
	}

	return results, nil
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
