package dateparser

import (
	"fmt"
	"strings"
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
	// NegativeTimestamp is parser to parse Unix timestamp in negative value.
	NegativeTimestamp
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

// Parse parses string representing date and/or time in recognizable localized formats.
// Supports parsing multiple languages.
func (p *Parser) Parse(cfg *Configuration, str string, formats ...string) (date.Date, error) {
	// Lock mutex
	p.Lock()
	defer p.Unlock()

	// Validate and initiate parsers
	for _, parser := range p.ParserTypes {
		if parser > NoSpacesTime {
			return date.Date{}, fmt.Errorf("invalid parser type: %d", parser)
		}
	}

	if len(p.ParserTypes) == 0 {
		p.ParserTypes = []ParserType{
			Timestamp,
			NegativeTimestamp,
			RelativeTime,
			CustomFormat,
			AbsoluteTime,
			NoSpacesTime,
		}
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
	locales, err := p.getApplicableLocales(cfg, iCfg, str)
	if err != nil {
		return date.Date{}, err
	}

	// Process each locale
	for _, locale := range locales {
		// Create date order for this locale
		dateOrder := locale.DateOrder
		if cfg.DateOrder != nil {
			do := cfg.DateOrder(locale.Name)
			if do, valid := validateDateOrder(do); valid {
				dateOrder = do
			}
		}

		// Create locale specific config
		lCfg := iCfg.Clone()
		lCfg.DateOrder = dateOrder

		// Translate string
		translations := language.Translate(lCfg, locale, str, false)
		translationsWithFormat := language.Translate(lCfg, locale, str, true)

		for _, parserType := range p.ParserTypes {
			switch parserType {
			case Timestamp:
				dt = timestamp.Parse(lCfg, str, false)
			case NegativeTimestamp:
				dt = timestamp.Parse(lCfg, str, true)
			case RelativeTime:
				dt = p.tryRelativeTime(lCfg, translations)
			case CustomFormat:
				dt = p.tryCustomFormat(lCfg, translationsWithFormat, formats...)
			case AbsoluteTime:
				dt = p.tryAbsoluteTime(lCfg, translations)
			case NoSpacesTime:
				dt = p.tryNoSpacesTime(lCfg, translations)
			}

			if !dt.IsZero() {
				if cfg.TryPreviousLocales {
					p.saveUsedLocale(locale)
				}

				dt.Locale = locale.Name
				return dt, nil
			}
		}
	}

	return date.Date{}, fmt.Errorf("failed to parse \"%s\": unknown format", originalStr)
}

func (p *Parser) tryRelativeTime(iCfg *setting.Configuration, translations []string) date.Date {
	for _, translation := range translations {
		dt := relative.Parse(iCfg, translation)
		if !dt.IsZero() {
			return dt
		}
	}
	return date.Date{}
}

func (p *Parser) tryCustomFormat(iCfg *setting.Configuration, translations []string, formats ...string) date.Date {
	for _, translation := range translations {
		dt := formatted.Parse(iCfg, translation, formats...)
		if !dt.IsZero() {
			return dt
		}
	}
	return date.Date{}
}

func (p *Parser) tryAbsoluteTime(iCfg *setting.Configuration, translations []string) date.Date {
	for _, translation := range translations {
		if t, tz := p.stripBracesAndTimezones(translation); t != "" {
			dt, _ := absolute.Parse(iCfg, t, tz)
			dt = p.applyTimezone(iCfg, dt, tz)
			if !dt.IsZero() {
				return dt
			}
		}
	}
	return date.Date{}
}

func (p *Parser) tryNoSpacesTime(iCfg *setting.Configuration, translations []string) date.Date {
	for _, translation := range translations {
		if t, tz := p.stripBracesAndTimezones(translation); t != "" {
			dt, _ := nospace.Parse(iCfg, t)
			dt = p.applyTimezone(iCfg, dt, tz)
			if !dt.IsZero() {
				return dt
			}
		}
	}
	return date.Date{}
}

func (p *Parser) getApplicableLocales(cfg *Configuration, iCfg *setting.Configuration, str string) ([]*data.LocaleData, error) {
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
	if cfg.TryPreviousLocales {
		ld := p.checkPreviousLocales(iCfg, dateStrings)
		if ld != nil {
			results = append(results, ld)
			resultTracker.Add(ld.Name)
		}
	}

	// If specified, use external detector to fetch languages
	languages := append([]string{}, cfg.Languages...)
	if p.DetectLanguagesFunction != nil && len(cfg.Locales) == 0 && len(languages) == 0 {
		detectionResults := p.DetectLanguagesFunction(str)
		languages = append(languages, detectionResults...)
	}

	// Load locales
	locales, err := language.GetLocales(cfg.Locales, languages, cfg.Region, cfg.UseGivenOrder, false)
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
		locales, _ := language.GetLocales(nil, cfg.DefaultLanguages, cfg.Region, cfg.UseGivenOrder, false)
		for _, locale := range locales {
			if !resultTracker.Contain(locale.Name) {
				results = append(results, locale)
			}
		}
	}

	return results, nil
}

func (p *Parser) checkPreviousLocales(iCfg *setting.Configuration, dateStrings []string) *data.LocaleData {
	for _, usedLocale := range p.usedLocales {
		for _, ds := range dateStrings {
			if p.localeIsApplicable(iCfg, usedLocale, ds) {
				return usedLocale
			}
		}
	}

	return nil
}

func (p *Parser) saveUsedLocale(ld *data.LocaleData) {
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

func (p *Parser) stripBracesAndTimezones(s string) (string, timezone.OffsetData) {
	s = strutil.StripBraces(s)
	return timezone.PopTzOffset(s)
}

func (p *Parser) applyTimezone(iCfg *setting.Configuration, dt date.Date, tz timezone.OffsetData) date.Date {
	if dt.IsZero() || (tz.IsZero() && iCfg.DefaultTimezone == nil) {
		return dt
	}

	var tzName string
	var tzOffset int

	if !tz.IsZero() {
		tzName, tzOffset = tz.Name, tz.Offset
	} else {
		tzName, tzOffset = dt.Time.In(iCfg.DefaultTimezone).Zone()
	}

	dt.Time = time.Date(dt.Time.Year(), dt.Time.Month(), dt.Time.Day(),
		dt.Time.Hour(), dt.Time.Minute(), dt.Time.Second(), dt.Time.Nanosecond(),
		time.FixedZone(tzName, tzOffset))

	return dt
}

func validateDateOrder(do string) (string, bool) {
	if len(do) != 3 {
		return do, false
	}

	do = strings.ToUpper(do)
	mapChars := map[rune]struct{}{}
	for _, r := range do {
		if r == 'D' || r == 'M' || r == 'Y' {
			mapChars[r] = struct{}{}
		} else {
			return do, false
		}
	}

	if len(mapChars) != 3 {
		return do, false
	}

	return do, true
}
