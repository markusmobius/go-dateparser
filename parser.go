package dateparser

import (
	"fmt"
	"iter"
	"strings"
	"sync"
	"time"

	"slices"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
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

// Parser detects languages, translates date strings, and parses dates and times.
// It supports concurrent Parse and Search calls. Configure its exported fields
// before use, provide concurrency-safe callbacks, and do not copy it after use.
type Parser struct {
	sync.Mutex

	// DetectLanguagesFunction returns candidate language codes for input text.
	// Parse uses it when Languages and Locales are empty. Search uses it when
	// Languages is empty; explicit Locales still restrict the candidate locales.
	DetectLanguagesFunction func(string) []string

	// ParserTypes selects the parsers and their order. The default order is Timestamp,
	// NegativeTimestamp, RelativeTime, CustomFormat, AbsoluteTime, and NoSpacesTime.
	// Explicit Go layouts are also tried before locale detection.
	ParserTypes []ParserType

	usedLocales        []*data.LocaleData
	usedLocalesTracker strutil.Dict
	uniqueCharsets     map[string][]rune
}

// ParserType identifies a parsing strategy.
type ParserType uint8

const (
	// Timestamp parses nonnegative Unix timestamps.
	Timestamp ParserType = iota
	// NegativeTimestamp parses negative Unix timestamps.
	NegativeTimestamp
	// RelativeTime parses relative dates such as
	// "1 year, 2 months ago" and "3 hours, 50 minutes ago".
	RelativeTime
	// CustomFormat parses translated dates using caller-supplied Go layouts.
	CustomFormat
	// AbsoluteTime parses absolute dates such as
	// "12 August 2021" and "23 January, 15:10:01".
	AbsoluteTime
	// NoSpacesTime parses compact dates, such as 20211011 for 2021-10-11.
	NoSpacesTime
)

// Parse parses a localized date or time. A nil configuration uses the defaults.
// Optional formats are Go time layouts, tried before locale-based parsing.
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

	dt, err = p.parseUsingLocales(cfg, iCfg, str, false, formats...)
	if err == nil && dt.IsZero() && cfg.IgnoreSurroundingText {
		dt, err = p.parseUsingLocales(cfg, iCfg, str, true, formats...)
	}
	if err != nil {
		return date.Date{}, err
	}
	if !dt.IsZero() {
		return dt, nil
	}
	return date.Date{}, fmt.Errorf("failed to parse \"%s\": unknown format", originalStr)
}

func (p *Parser) parseUsingLocales(cfg *Configuration, iCfg *setting.Configuration, str string, ignoreSurroundingText bool, formats ...string) (date.Date, error) {
	var dt date.Date
	// Find the suitable locales for this string
	locales, err := p.getApplicableLocales(cfg, iCfg, str, ignoreSurroundingText)
	if err != nil {
		return date.Date{}, err
	}

	// Process each locale
	for locale := range locales {
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
		dateOrders := []string{dateOrder}
		if cfg.DateOrder == nil && slices.Contains(cfg.RequiredParts, "year") && !slices.Contains(cfg.RequiredParts, "day") {
			for _, order := range []string{"MYD", "YMD"} {
				if !slices.Contains(dateOrders, order) {
					dateOrders = append(dateOrders, order)
				}
			}
		}

		// Translate string
		translations := language.Translate(lCfg, locale, str, false, ignoreSurroundingText)
		translationsWithFormat := language.Translate(lCfg, locale, str, true, ignoreSurroundingText)

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
			case AbsoluteTime, NoSpacesTime:
				for _, order := range dateOrders {
					lCfg.DateOrder = order
					if parserType == AbsoluteTime {
						dt = p.tryAbsoluteTime(lCfg, translations)
					} else {
						dt = p.tryNoSpacesTime(lCfg, translations)
					}
					if !dt.IsZero() {
						break
					}
				}
				lCfg.DateOrder = dateOrder
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

	return date.Date{}, nil
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

func (p *Parser) getApplicableLocales(cfg *Configuration, iCfg *setting.Configuration, str string, ignoreSurroundingText bool) (iter.Seq[*data.LocaleData], error) {
	// Normalize and prepare date strings
	str = strutil.NormalizeString(str)
	dateStrings := []string{digit.NormalizeString(strutil.NormalizeString(str))}
	if poppedTz, _ := timezone.PopTzOffset(str); poppedTz != str {
		dateStrings = append(dateStrings, digit.NormalizeString(strutil.NormalizeString(poppedTz)))
	}

	// Fetch previously used locales first
	var previous *data.LocaleData
	if cfg.TryPreviousLocales {
		previous = p.checkPreviousLocales(iCfg, dateStrings, ignoreSurroundingText)
	}

	// If specified, use external detector to fetch languages
	languages := slices.Clone(cfg.Languages)
	if p.DetectLanguagesFunction != nil && len(cfg.Locales) == 0 && len(languages) == 0 {
		detectionResults := p.DetectLanguagesFunction(str)
		languages = append(languages, detectionResults...)
	}

	// Load locales
	locales, err := language.GetLocales(cfg.Locales, languages, cfg.Region, cfg.UseGivenOrder, false)
	if err != nil && err != language.ErrNotFound {
		return nil, err
	}

	var defaults []*data.LocaleData
	if len(iCfg.DefaultLanguages) > 0 {
		defaults, _ = language.GetLocales(nil, cfg.DefaultLanguages, cfg.Region, cfg.UseGivenOrder, false)
	}

	return func(yield func(*data.LocaleData) bool) {
		resultTracker := strutil.NewDict()
		if previous != nil {
			resultTracker.Add(previous.Name)
			if !yield(previous) {
				return
			}
		}

		for _, locale := range locales {
			if resultTracker.Contain(locale.Name) {
				continue
			}

			var isApplicable bool
			for _, dateString := range dateStrings {
				if p.localeIsApplicable(iCfg, locale, dateString, ignoreSurroundingText) {
					isApplicable = true
					break
				}
			}
			if isApplicable {
				resultTracker.Add(locale.Name)
				if !yield(locale) {
					return
				}
			}
		}

		for _, locale := range defaults {
			if !resultTracker.Contain(locale.Name) && !yield(locale) {
				return
			}
		}
	}, nil
}

func (p *Parser) checkPreviousLocales(iCfg *setting.Configuration, dateStrings []string, ignoreSurroundingText bool) *data.LocaleData {
	for _, usedLocale := range p.usedLocales {
		for _, ds := range dateStrings {
			if p.localeIsApplicable(iCfg, usedLocale, ds, ignoreSurroundingText) {
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

func (p *Parser) localeIsApplicable(iCfg *setting.Configuration, ld *data.LocaleData, s string, ignoreSurroundingText bool) bool {
	return language.IsApplicablePrepared(iCfg, ld, s, ignoreSurroundingText)
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
