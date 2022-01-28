package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type LocaleData struct {
	Name                      string               `json:",omitempty"`
	Parent                    string               `json:",omitempty"`
	DateOrder                 string               `json:",omitempty"`
	NoWordSpacing             bool                 `json:",omitempty"`
	SentenceSplitterGroup     int                  `json:",omitempty"`
	Charset                   []rune               `json:",omitempty"`
	SkipWords                 []string             `json:",omitempty"`
	PertainWords              []string             `json:",omitempty"`
	Simplifications           []SimplificationData `json:",omitempty"`
	Translations              map[string]string    `json:",omitempty"`
	RelativeType              map[string]string    `json:",omitempty"`
	RelativeTypeRegexes       map[string]string    `json:",omitempty"`
	CombinedRegexPattern      string               `json:",omitempty"`
	ExactCombinedRegexPattern string               `json:",omitempty"`
	KnownWordsPattern         string               `json:",omitempty"`
}

type SimplificationData struct {
	Pattern     string
	Replacement string
}

func (ld *LocaleData) AddSimplification(pattern string, replacement string) {
	// Sanitize pattern
	pattern = normalizeString(pattern)
	pattern = strings.ReplaceAll(pattern, `{0}`, `(\d+)`)
	if !ld.NoWordSpacing {
		pattern = `(\A|[^\pL\pM\d]|_)` + pattern + `(\z|[^\pL\pM\d]|_)`
	}

	// Sanitize replacement
	replacement = normalizeString(replacement)
	replacement = rxPythonCaptureGroup.ReplaceAllString(replacement, "$${$1}")
	if !ld.NoWordSpacing {
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
		replacement = fmt.Sprintf("${1}%s${%d}", replacement, maxNumber+1)
	}

	// Save simplification if pattern not empty
	if pattern != "" {
		ld.Simplifications = append(ld.Simplifications, SimplificationData{
			Pattern:     pattern,
			Replacement: replacement,
		})
	}
}

func (ld *LocaleData) AddTranslation(word string, translation string, cleanWord bool) {
	// Prepare map
	if ld.Translations == nil {
		ld.Translations = map[string]string{}
	}

	// Sanitize word
	if cleanWord {
		word = cleanString(word)
	} else {
		word = normalizeString(word)
	}

	// Save translation if word not empty
	if word != "" {
		ld.Translations[word] = normalizeString(translation)
	}
}

func (ld *LocaleData) AddRelativeType(pattern string, translation string, cleanPattern bool) {
	// Prepare maps
	if ld.RelativeType == nil {
		ld.RelativeType = map[string]string{}
	}

	if ld.RelativeTypeRegexes == nil {
		ld.RelativeTypeRegexes = map[string]string{}
	}

	// Sanitize pattern
	if cleanPattern {
		pattern = cleanString(pattern)
	} else {
		pattern = normalizeString(pattern)
	}
	pattern = strings.ReplaceAll(pattern, `{0}`, `(\d+)`)

	// Specify target map
	var targetMap map[string]string
	if strings.Contains(pattern, `(\d+)`) {
		targetMap = ld.RelativeTypeRegexes
	} else {
		targetMap = ld.RelativeType
	}

	// Save if pattern not empty
	if _, exist := targetMap[pattern]; !exist && pattern != "" {
		translation = normalizeString(translation)
		translation = rxPythonCaptureGroup.ReplaceAllString(translation, "$${$1}")
		targetMap[pattern] = translation
	}
}

func (ld *LocaleData) CombineRegexPatterns() {
	// Fetch all regex patterns
	var regexPatterns []string
	for pattern := range ld.RelativeTypeRegexes {
		pattern = rxParentheses.ReplaceAllString(pattern, "")
		regexPatterns = append(regexPatterns, pattern)
	}

	// Sort patterns by the longest
	sort.Slice(regexPatterns, func(a, b int) bool {
		patternA, patternB := regexPatterns[a], regexPatterns[b]
		lenA := utf8.RuneCountInString(patternA)
		lenB := utf8.RuneCountInString(patternB)
		if lenA != lenB {
			return lenA > lenB
		}
		return patternA < patternB
	})

	// Combine the patterns
	ld.CombinedRegexPattern = ``
	ld.ExactCombinedRegexPattern = ``

	if len(regexPatterns) > 0 {
		combined := strings.Join(regexPatterns, "|")

		if ld.NoWordSpacing {
			ld.CombinedRegexPattern = `(` + combined + `)`
		} else {
			ld.CombinedRegexPattern = `(\A|[^\pL\pM\d]|_)(` + combined + `)(\z|[^\pL\pM\d]|_)`
		}

		ld.ExactCombinedRegexPattern = "^(" + combined + ")$"
	}
}

func (ld *LocaleData) GenerateKnownWordPattern() {
	// Fetch all texts
	var texts []string

	for text := range ld.Translations {
		text = regexp.QuoteMeta(text)
		texts = append(texts, text)
	}

	for text := range ld.RelativeType {
		text = regexp.QuoteMeta(text)
		texts = append(texts, text)
	}

	// Sort texts by the longest
	sort.Slice(texts, func(a, b int) bool {
		textA, textB := texts[a], texts[b]
		lenA := utf8.RuneCountInString(textA)
		lenB := utf8.RuneCountInString(textB)
		if lenA != lenB {
			return lenA > lenB
		}
		return textA < textB
	})

	// Combine the texts
	ld.KnownWordsPattern = ""

	if len(texts) > 0 {
		combined := strings.Join(texts, "|")

		if ld.NoWordSpacing {
			ld.KnownWordsPattern = `^(.*?)(` + combined + `)(.*)$`
		} else {
			ld.KnownWordsPattern = `^(.*?(?:\A|[^\pL\pM\d]|_|\d))(` + combined + `)((?:\z|[^\pL\pM\d]|_|\d).*)$`
		}
	}
}

func (ld *LocaleData) GenerateCharset() {
	// Prepare variables and helper function
	var chars []rune
	tracker := map[rune]struct{}{}
	fnSaveChars := func(s string) {
		if rxCharsetFilter.MatchString(s) {
			return
		}

		for _, r := range s {
			if unicode.Is(commonChars, r) {
				continue
			}

			if _, exist := tracker[r]; !exist {
				tracker[r] = struct{}{}
				chars = append(chars, r)
			}
		}
	}

	// Fetch all chars
	for text := range ld.Translations {
		fnSaveChars(text)
	}

	for text := range ld.RelativeType {
		fnSaveChars(text)
	}

	// Sort the chars
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	ld.Charset = chars
}

func (ld LocaleData) Clone() LocaleData {
	cloneMap := func(m map[string]string) map[string]string {
		newMap := map[string]string{}
		for k, v := range m {
			newMap[k] = v
		}
		return newMap
	}

	return LocaleData{
		Name:                      ld.Name,
		DateOrder:                 ld.DateOrder,
		NoWordSpacing:             ld.NoWordSpacing,
		SentenceSplitterGroup:     ld.SentenceSplitterGroup,
		Charset:                   append([]rune{}, ld.Charset...),
		SkipWords:                 append([]string{}, ld.SkipWords...),
		PertainWords:              append([]string{}, ld.PertainWords...),
		Simplifications:           append([]SimplificationData{}, ld.Simplifications...),
		Translations:              cloneMap(ld.Translations),
		RelativeType:              cloneMap(ld.RelativeType),
		RelativeTypeRegexes:       cloneMap(ld.RelativeTypeRegexes),
		CombinedRegexPattern:      ld.CombinedRegexPattern,
		ExactCombinedRegexPattern: ld.ExactCombinedRegexPattern,
		KnownWordsPattern:         ld.KnownWordsPattern,
	}
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
	clone.Simplifications = cleanSimplifications(append(clone.Simplifications, input.Simplifications...)...)

	for word, translation := range input.Translations {
		clone.Translations[word] = translation
	}

	for pattern, translation := range input.RelativeType {
		clone.RelativeType[pattern] = translation
	}

	for pattern, translation := range input.RelativeTypeRegexes {
		clone.RelativeTypeRegexes[pattern] = translation
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

	reduceSimplifications := func(base, input []SimplificationData) []SimplificationData {
		tracker := map[string]struct{}{}
		for _, data := range input {
			key := data.Pattern + "==" + data.Replacement
			tracker[key] = struct{}{}
		}

		var reducedSimplifications []SimplificationData
		for _, data := range base {
			key := data.Pattern + "==" + data.Replacement
			if _, exist := tracker[key]; !exist {
				reducedSimplifications = append(reducedSimplifications, data)
			}
		}

		return reducedSimplifications
	}

	reduceMap := func(base, input map[string]string) map[string]string {
		newMap := map[string]string{}
		for k, v := range base {
			if _, exist := input[k]; !exist {
				newMap[k] = v
			}
		}
		return newMap
	}

	// Reduce data
	clone := ld.Clone()
	clone.SkipWords = reduceStrings(clone.SkipWords, input.SkipWords)
	clone.PertainWords = reduceStrings(clone.PertainWords, input.PertainWords)
	clone.Simplifications = reduceSimplifications(clone.Simplifications, input.Simplifications)
	clone.Translations = reduceMap(clone.Translations, input.Translations)
	clone.RelativeType = reduceMap(clone.RelativeType, input.RelativeType)
	clone.RelativeTypeRegexes = reduceMap(clone.RelativeTypeRegexes, input.RelativeTypeRegexes)

	return clone
}

func (ld *LocaleData) Validate() error {
	// Validate patterns
	var err error

	for _, data := range ld.Simplifications {
		_, err = regexp.Compile(data.Pattern)
		if err != nil {
			return fmt.Errorf("simplification pattern error for %s: %w", ld.Name, err)
		}
	}

	for pattern := range ld.RelativeTypeRegexes {
		_, err = regexp.Compile(pattern)
		if err != nil {
			return fmt.Errorf("relative pattern error for %s: %w", ld.Name, err)
		}
	}

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

type MapEntry struct {
	Key   string
	Value string
}
