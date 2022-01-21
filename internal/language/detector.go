package language

import (
	"sync"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

type LocaleDetector struct {
	sync.Mutex
	Configuration       *setting.Configuration
	Languages           []string
	Locales             []string
	Region              string
	UseGivenOrder       bool
	UseExternalDetector bool
	ConfidenceThreshold float64

	usedLocales        []*data.LocaleData
	usedLocalesTracker map[string]struct{}
}

func (d *LocaleDetector) Detect(str string) ([]*data.LocaleData, error) {
	// Lock first
	d.Lock()
	defer d.Unlock()

	// Prepare results
	var results []*data.LocaleData
	resultTracker := map[string]struct{}{}

	// Normalize and prepare date strings
	str = strutil.NormalizeString(str)
	dateStrings := []string{str}
	if poppedTz, _ := timezone.PopTzOffset(str); poppedTz != str {
		dateStrings = append(dateStrings, poppedTz)
	}

	// Fetch previously used locales first
	for _, usedLocale := range d.usedLocales {
		for _, ds := range dateStrings {
			if d.isApplicable(usedLocale, ds) {
				results = append(results, usedLocale)
				resultTracker[usedLocale.Name] = struct{}{}
				break
			}
		}
	}

	// If necessary, use external detector to fetch languages
	languages := append([]string{}, d.Languages...)

	if d.UseExternalDetector && len(d.Locales) == 0 && len(languages) == 0 {
		detectionResults := externalLanguageDetector.ComputeLanguageConfidenceValues(str)
		for _, detectionResult := range detectionResults {
			if d.ConfidenceThreshold <= 0 || detectionResult.Value() >= d.ConfidenceThreshold {
				langCode := detectionResult.Language().IsoCode639_1().String()
				languages = append(languages, langCode)
			}
		}
	}

	// Load locales
	locales, err := GetLocales(languages, d.Locales, d.Region, d.UseGivenOrder, false)
	if err != nil && err != ErrNotFound {
		return nil, err
	}

	names := []string{}
	for _, l := range locales {
		names = append(names, l.Name)
	}

	for _, locale := range locales {
		// fmt.Println("XX:", locale.Name)
		// If locale already stored in result, skip
		if _, exist := resultTracker[locale.Name]; exist {
			continue
		}

		// Check if locale is applicable
		var isApplicable bool
		for _, ds := range dateStrings {
			if d.isApplicable(locale, ds) {
				isApplicable = true
				break
			}
		}

		if isApplicable {
			results = append(results, locale)
			resultTracker[locale.Name] = struct{}{}
		}
	}

	// Finally, try locales of default languages
	if d.Configuration != nil && len(d.Configuration.DefaultLanguages) > 0 {
		locales, _ := GetLocales(d.Configuration.DefaultLanguages, nil, d.Region, d.UseGivenOrder, false)
		for _, locale := range locales {
			if _, exist := resultTracker[locale.Name]; !exist {
				results = append(results, locale)
			}
		}
	}

	return results, nil
}

func (d *LocaleDetector) SaveUsedLocale(ld *data.LocaleData) {
	// Lock
	d.Lock()
	defer d.Unlock()

	// Prepare map
	if d.usedLocalesTracker == nil {
		d.usedLocalesTracker = map[string]struct{}{}
	}

	// Save last used locale
	if _, exist := d.usedLocalesTracker[ld.Name]; !exist {
		d.usedLocales = append([]*data.LocaleData{ld}, d.usedLocales...)
		d.usedLocalesTracker[ld.Name] = struct{}{}
	}
}

func (d *LocaleDetector) isApplicable(ld *data.LocaleData, s string) bool {
	return IsApplicable(d.Configuration, ld, s, false)
}
