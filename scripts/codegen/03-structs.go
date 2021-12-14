package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type LocaleData struct {
	Name                  string            `json:",omitempty"`
	Parent                string            `json:",omitempty"`
	DateOrder             string            `json:",omitempty"`
	NoWordSpacing         bool              `json:",omitempty"`
	SentenceSplitterGroup int               `json:",omitempty"`
	SkipWords             []string          `json:",omitempty"`
	PertainWords          []string          `json:",omitempty"`
	Simplifications       map[string]string `json:",omitempty"`
	Translations          []TranslationData `json:",omitempty"`

	patternTrackers map[string]struct{}
}

type TranslationData struct {
	Pattern     string
	Translation string
}

func (ld *LocaleData) AddSimplification(pattern string, replacement string) {
	if ld.Simplifications == nil {
		ld.Simplifications = make(map[string]string)
	}

	pattern = normalizeString(pattern)
	replacement = normalizeString(replacement)
	ld.Simplifications[pattern] = replacement
}

func (ld *LocaleData) AddTranslation(pattern string, translation string, cleanPattern bool) {
	// Prepare tracker
	if ld.patternTrackers == nil {
		ld.patternTrackers = make(map[string]struct{})
	}

	// Sanitize pattern
	if cleanPattern {
		pattern = cleanString(pattern)
	} else {
		pattern = normalizeString(pattern)
	}
	pattern = strings.ReplaceAll(pattern, `{0}`, `(\d+)`)

	// Sanitize translation
	translation = rxPythonCaptureGroup.ReplaceAllString(translation, "$${$1}")

	// Make sure pattern not empty
	if pattern == "" {
		return
	}

	// Make sure pattern never used
	if _, exist := ld.patternTrackers[pattern]; exist {
		return
	}

	// Save translation
	ld.patternTrackers[pattern] = struct{}{}
	ld.Translations = append(ld.Translations, TranslationData{
		Pattern:     pattern,
		Translation: translation,
	})
}

func (ld LocaleData) Clone() LocaleData {
	clone := LocaleData{
		Name:                  ld.Name,
		DateOrder:             ld.DateOrder,
		NoWordSpacing:         ld.NoWordSpacing,
		SentenceSplitterGroup: ld.SentenceSplitterGroup,
		SkipWords:             append([]string{}, ld.SkipWords...),
		PertainWords:          append([]string{}, ld.PertainWords...),
		Translations:          append([]TranslationData{}, ld.Translations...),
		Simplifications:       map[string]string{},
		patternTrackers:       map[string]struct{}{},
	}

	for pattern, replacement := range ld.Simplifications {
		clone.Simplifications[pattern] = replacement
	}

	for pattern := range ld.patternTrackers {
		clone.patternTrackers[pattern] = struct{}{}
	}

	return clone
}

func (ld LocaleData) Merge(input LocaleData) LocaleData {
	clone := ld.Clone()

	if clone.Name == "" {
		clone.Name = input.Name
	}

	if clone.DateOrder == "" {
		clone.DateOrder = input.DateOrder
	}

	clone.NoWordSpacing = clone.NoWordSpacing || input.NoWordSpacing
	clone.SentenceSplitterGroup = input.SentenceSplitterGroup
	clone.SkipWords = cleanList(false, append(clone.SkipWords, input.SkipWords...)...)
	clone.PertainWords = cleanList(false, append(clone.PertainWords, input.PertainWords...)...)

	for pattern, replacement := range input.Simplifications {
		clone.AddSimplification(pattern, replacement)
	}

	for _, trans := range input.Translations {
		clone.AddTranslation(trans.Pattern, trans.Translation, false)
	}

	return clone
}

func (ld LocaleData) Reduce(input LocaleData) LocaleData {
	// Helper function
	reduceStrings := func(base, input []string) []string {
		tracker := map[string]struct{}{}
		for _, s := range input {
			tracker[s] = struct{}{}
		}

		var reducedStrings []string
		for _, s := range base {
			if _, exist := tracker[s]; !exist {
				reducedStrings = append(reducedStrings, s)
			}
		}

		return reducedStrings
	}

	reduceTranslations := func(base, input []TranslationData) []TranslationData {
		tracker := map[string]struct{}{}
		for _, trans := range input {
			tracker[trans.Pattern] = struct{}{}
		}

		var reducedTranslations []TranslationData
		for _, trans := range base {
			if _, exist := tracker[trans.Pattern]; !exist {
				reducedTranslations = append(reducedTranslations, trans)
			}
		}

		return reducedTranslations
	}

	// Reduce data
	clone := ld.Clone()
	clone.SkipWords = reduceStrings(clone.SkipWords, input.SkipWords)
	clone.PertainWords = reduceStrings(clone.PertainWords, input.PertainWords)
	clone.Translations = reduceTranslations(clone.Translations, input.Translations)

	for pattern := range input.Simplifications {
		if _, exist := clone.Simplifications[pattern]; exist {
			delete(clone.Simplifications, pattern)
		}
	}

	return clone
}

func (ld *LocaleData) Validate() error {
	// Change patterns and its replacement depending on if the language uses
	// white space between words or not
	var patternChanger func(string) string
	var replacementChanger func(string) string

	if ld.NoWordSpacing {
		patternChanger = func(s string) string {
			return s
		}

		replacementChanger = func(s string) string {
			return " " + s + " "
		}
	} else {
		patternChanger = func(pattern string) string {
			return `(\A|\W|_)` + pattern + `(\z|\W|_)`
		}

		replacementChanger = func(replacement string) string {
			maxNumber := 1
			replacement = rxGoCaptureGroup.ReplaceAllStringFunc(replacement, func(s string) string {
				parts := rxGoCaptureGroup.FindStringSubmatch(s)
				number, _ := strconv.Atoi(parts[1])
				number++

				if number > maxNumber {
					maxNumber = number
				}

				return fmt.Sprintf("${%d}", number)
			})

			return fmt.Sprintf("${1}%s${%d}", replacement, maxNumber+1)
		}
	}

	newSimplifications := map[string]string{}
	for pattern, replacement := range ld.Simplifications {
		pattern = patternChanger(pattern)
		replacement = replacementChanger(replacement)
		newSimplifications[pattern] = replacement
	}
	ld.Simplifications = newSimplifications

	ld.patternTrackers = map[string]struct{}{}
	for i, trans := range ld.Translations {
		newPattern := patternChanger(trans.Pattern)
		newTranslation := replacementChanger(trans.Translation)

		ld.patternTrackers[newPattern] = struct{}{}
		ld.Translations[i] = TranslationData{newPattern, newTranslation}
	}

	// Validate patterns
	var err error
	for pattern := range ld.Simplifications {
		_, err = regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("simplification pattern error for %s: %w", ld.Name, err)
		}
	}

	for _, trans := range ld.Translations {
		_, err = regexp.Compile(trans.Pattern)
		if err != nil {
			return fmt.Errorf("translation pattern error for %s: %w", ld.Name, err)
		}
	}

	// Sort translations
	sort.Slice(ld.Translations, func(a, b int) bool {
		patternA := ld.Translations[a].Pattern
		patternB := ld.Translations[b].Pattern
		lenA := utf8.RuneCountInString(patternA)
		lenB := utf8.RuneCountInString(patternB)

		if lenA != lenB {
			return lenA > lenB
		}

		return patternA < patternB
	})

	return nil
}

type CldrTerritoryData struct {
	Supplemental struct {
		TerritoryInfo map[string]struct {
			Gdp                string `json:"_gdp"`
			LiteracyPercent    string `json:"_literacyPercent"`
			Population         string `json:"_population"`
			LanguagePopulation map[string]struct {
				PopulationPercent string `json:"_populationPercent"`
				OfficialStatus    string `json:"_officialStatus"`
			} `json:"languagePopulation"`
		} `json:"territoryInfo"`
	} `json:"supplemental"`
}

type CldrGregorianData struct {
	Main map[string]struct {
		Dates struct {
			Calendars struct {
				Gregorian struct {
					Months      CldrGregorianDataPart `json:"months"`
					Days        CldrGregorianDataPart `json:"days"`
					DayPeriods  CldrGregorianDataPart `json:"dayPeriods"`
					DateFormats struct {
						ShortItf interface{} `json:"short"`
						Short    string      `json:"-"`
					} `json:"dateFormats"`
				} `json:"gregorian"`
			} `json:"calendars"`
		} `json:"dates"`
	} `json:"main"`
}

type CldrGregorianDataPart struct {
	Format struct {
		Abbreviated map[string]string `json:"abbreviated"`
		Wide        map[string]string `json:"wide"`
	} `json:"format"`
	StandAlone struct {
		Abbreviated map[string]string `json:"abbreviated"`
		Wide        map[string]string `json:"wide"`
	} `json:"stand-alone"`
}

type CldrDateFieldsData struct {
	Main map[string]struct {
		Dates struct {
			Fields map[string]CldrDateFieldsPart `json:"fields"`
		} `json:"dates"`
	} `json:"main"`
}

type CldrDateFieldsPart struct {
	DisplayName            string                       `json:"displayName"`
	RelativeTypeMin1       string                       `json:"relative-type--1"`
	RelativeType0          string                       `json:"relative-type-0"`
	RelativeType1          string                       `json:"relative-type-1"`
	RelativeTimeTypeFuture CldrDateFieldsRelTimePattern `json:"relativeTime-type-future"`
	RelativeTimeTypePast   CldrDateFieldsRelTimePattern `json:"relativeTime-type-past"`
}

type CldrDateFieldsRelTimePattern struct {
	CountOne   string `json:"relativeTimePattern-count-one"`
	CountOther string `json:"relativeTimePattern-count-other"`
}

type SupplementaryData struct {
	SkipWords             []string            `json:"skip,omitempty"                    yaml:"skip,omitempty"`
	PertainWords          []string            `json:"pertain,omitempty"                 yaml:"pertain,omitempty"`
	NoWordSpacing         bool                `json:"no_word_spacing,omitempty"         yaml:"no_word_spacing,omitempty"`
	SentenceSplitterGroup int                 `json:"sentence_splitter_group,omitempty" yaml:"sentence_splitter_group,omitempty"`
	January               []string            `json:"january,omitempty"                 yaml:"january,omitempty"`
	February              []string            `json:"february,omitempty"                yaml:"february,omitempty"`
	March                 []string            `json:"march,omitempty"                   yaml:"march,omitempty"`
	April                 []string            `json:"april,omitempty"                   yaml:"april,omitempty"`
	May                   []string            `json:"may,omitempty"                     yaml:"may,omitempty"`
	June                  []string            `json:"june,omitempty"                    yaml:"june,omitempty"`
	July                  []string            `json:"july,omitempty"                    yaml:"july,omitempty"`
	August                []string            `json:"august,omitempty"                  yaml:"august,omitempty"`
	September             []string            `json:"september,omitempty"               yaml:"september,omitempty"`
	October               []string            `json:"october,omitempty"                 yaml:"october,omitempty"`
	November              []string            `json:"november,omitempty"                yaml:"november,omitempty"`
	December              []string            `json:"december,omitempty"                yaml:"december,omitempty"`
	Monday                []string            `json:"monday,omitempty"                  yaml:"monday,omitempty"`
	Tuesday               []string            `json:"tuesday,omitempty"                 yaml:"tuesday,omitempty"`
	Wednesday             []string            `json:"wednesday,omitempty"               yaml:"wednesday,omitempty"`
	Thursday              []string            `json:"thursday,omitempty"                yaml:"thursday,omitempty"`
	Friday                []string            `json:"friday,omitempty"                  yaml:"friday,omitempty"`
	Saturday              []string            `json:"saturday,omitempty"                yaml:"saturday,omitempty"`
	Sunday                []string            `json:"sunday,omitempty"                  yaml:"sunday,omitempty"`
	AM                    []string            `json:"am,omitempty"                      yaml:"am,omitempty"`
	PM                    []string            `json:"pm,omitempty"                      yaml:"pm,omitempty"`
	Decade                []string            `json:"decade,omitempty"                  yaml:"decade,omitempty"`
	Year                  []string            `json:"year,omitempty"                    yaml:"year,omitempty"`
	Month                 []string            `json:"month,omitempty"                   yaml:"month,omitempty"`
	Week                  []string            `json:"week,omitempty"                    yaml:"week,omitempty"`
	Day                   []string            `json:"day,omitempty"                     yaml:"day,omitempty"`
	Hour                  []string            `json:"hour,omitempty"                    yaml:"hour,omitempty"`
	Minute                []string            `json:"minute,omitempty"                  yaml:"minute,omitempty"`
	Second                []string            `json:"second,omitempty"                  yaml:"second,omitempty"`
	In                    []string            `json:"in,omitempty"                      yaml:"in,omitempty"`
	Ago                   []string            `json:"ago,omitempty"                     yaml:"ago,omitempty"`
	RelativeType          map[string][]string `json:"relative-type,omitempty"           yaml:"relative-type,omitempty"`
	RelativeTypeRegex     map[string][]string `json:"relative-type-regex,omitempty"     yaml:"relative-type-regex,omitempty"`
	Simplifications       []map[string]string `json:"simplifications,omitempty"         yaml:"simplifications,omitempty"`
}
