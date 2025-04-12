package main

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
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
	Charset                   map[rune]struct{}    `json:",omitempty"`
	SkipWords                 []string             `json:",omitempty"`
	PertainWords              []string             `json:",omitempty"`
	Abbreviations             []string             `json:",omitempty"`
	Simplifications           []SimplificationData `json:",omitempty"`
	Translations              map[string][]string  `json:",omitempty"`
	RelativeType              map[string]string    `json:",omitempty"`
	RelativeTypeRegexes       map[string]string    `json:",omitempty"`
	CombinedRegexPattern      string               `json:",omitempty"`
	ExactCombinedRegexPattern string               `json:",omitempty"`
	KnownWords                []string             `json:",omitempty"`
}

type SimplificationData struct {
	Pattern     string
	Replacement string
}

func (ld *LocaleData) AddCharset(str string) {
	// Prepare tracker
	if ld.Charset == nil {
		ld.Charset = make(map[rune]struct{})
	}

	// Prepare strings to save
	strs := []string{
		cleanString(str),
		normalizeCharset(str),
	}

	for _, str := range strs {
		// Check if this string doesn't have to be saved
		if rxCharsetFilter.MatchString(str) {
			return
		}

		// Save each chars
		for _, r := range str {
			if !unicode.Is(commonChars, r) {
				ld.Charset[r] = struct{}{}
			}
		}
	}
}

func (ld *LocaleData) AddSimplification(pattern string, replacement string) {
	// Sanitize pattern
	pattern = normalizeString(pattern)
	pattern = strings.ReplaceAll(pattern, `{0}`, `(\d+)`)
	pattern = strings.ReplaceAll(pattern, `\w`, `[\pL\pM\d_]`)
	pattern = strings.ReplaceAll(pattern, `\W`, `[^\pL\pM\d_]`)
	pattern = strings.ReplaceAll(pattern, `\b`, `(?:\z|[^\pL\pM\d_])`)
	if !ld.NoWordSpacing {
		pattern = `(\A|[^\pL\pM\d_])` + pattern + `(\z|[^\pL\pM\d_])`
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
		ld.Translations = map[string][]string{}
	}

	// Save the charset before word normalized
	ld.AddCharset(word)

	// Sanitize word
	if cleanWord {
		word = cleanString(word)
	} else {
		word = normalizeString(word)
	}

	// Save translation if word not empty
	if word != "" {
		translation := normalizeString(translation)
		if !slices.Contains(ld.Translations[word], translation) {
			ld.Translations[word] = append(ld.Translations[word], translation)
		}
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

	// If not regex, save the pattern's charset
	if !strings.Contains(pattern, `\d+`) && !strings.Contains(pattern, `{0}`) {
		ld.AddCharset(pattern)
	}

	// Sanitize pattern
	if cleanPattern {
		pattern = cleanString(pattern)
	} else {
		pattern = normalizeString(pattern)
	}

	// Convert CLDR pattern to regex pattern
	pattern = strings.ReplaceAll(pattern, `{0}`, `(\d+[.,]?\d*)`)

	// Specify target map
	var targetMap map[string]string
	if strings.Contains(pattern, `\d+`) {
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

func (ld *LocaleData) GenerateKnownWords() {
	// Fetch all words
	translationWords := slices.Collect(maps.Keys(ld.Translations))
	relativeWords := slices.Collect(maps.Keys(ld.RelativeType))

	var words []string
	words = append(words, translationWords...)
	words = append(words, relativeWords...)

	// Make words unique
	slices.Sort(words)
	words = slices.Compact(words)

	// Sort words by the longest
	sort.Slice(words, func(a, b int) bool {
		wordA, wordB := words[a], words[b]
		lenA := utf8.RuneCountInString(wordA)
		lenB := utf8.RuneCountInString(wordB)
		if lenA != lenB {
			return lenA > lenB
		}
		return wordA < wordB
	})

	ld.KnownWords = words
}

func (ld *LocaleData) GenerateAbbreviations() {
	// Extract words
	translationWords := slices.Collect(maps.Keys(ld.Translations))
	relativeWords := slices.Collect(maps.Keys(ld.RelativeType))

	var words []string
	words = append(words, translationWords...)
	words = append(words, relativeWords...)

	// Filter abbreviations from words
	var abbreviations []string
	for _, word := range words {
		trimmedWord := strings.TrimSuffix(word, ".")
		if trimmedWord != word && trimmedWord != "" {
			abbreviations = append(abbreviations, trimmedWord)
		}
	}

	// Sort the abbreviations
	slices.Sort(abbreviations)
	words = slices.Compact(abbreviations)

	ld.Abbreviations = abbreviations
}

func (ld LocaleData) Clone() LocaleData {
	clone := LocaleData{
		Name:                      ld.Name,
		DateOrder:                 ld.DateOrder,
		NoWordSpacing:             ld.NoWordSpacing,
		SentenceSplitterGroup:     ld.SentenceSplitterGroup,
		Charset:                   cloneMap(ld.Charset),
		SkipWords:                 slices.Clone(ld.SkipWords),
		PertainWords:              slices.Clone(ld.PertainWords),
		Abbreviations:             slices.Clone(ld.Abbreviations),
		Simplifications:           slices.Clone(ld.Simplifications),
		Translations:              cloneMapSlice(ld.Translations),
		RelativeType:              cloneMap(ld.RelativeType),
		RelativeTypeRegexes:       cloneMap(ld.RelativeTypeRegexes),
		CombinedRegexPattern:      ld.CombinedRegexPattern,
		ExactCombinedRegexPattern: ld.ExactCombinedRegexPattern,
		KnownWords:                slices.Clone(ld.KnownWords),
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

	if input.SentenceSplitterGroup > 0 {
		clone.SentenceSplitterGroup = input.SentenceSplitterGroup
	}

	clone.NoWordSpacing = clone.NoWordSpacing || input.NoWordSpacing
	clone.SkipWords = cleanList(false, append(clone.SkipWords, input.SkipWords...)...)
	clone.PertainWords = cleanList(false, append(clone.PertainWords, input.PertainWords...)...)
	clone.Abbreviations = cleanList(false, append(clone.Abbreviations, input.Abbreviations...)...)
	clone.Simplifications = cleanSimplifications(append(clone.Simplifications, input.Simplifications...)...)

	for word, translations := range input.Translations {
		merged := append(clone.Translations[word], translations...)

		slices.Sort(merged)
		merged = slices.Compact(merged)

		clone.Translations[word] = merged
	}

	maps.Copy(clone.RelativeType, input.RelativeType)
	maps.Copy(clone.RelativeTypeRegexes, input.RelativeTypeRegexes)

	for r := range input.Charset {
		clone.Charset[r] = struct{}{}
	}

	return clone
}

func (ld LocaleData) Reduce(input LocaleData) LocaleData {
	// Helper function
	reducePattern := func(base, input string) string {
		if base == input {
			return ""
		}
		return base
	}

	reduceKnownWords := func(base, input []string) []string {
		if len(base) != len(input) {
			return base
		}

		for i := range base {
			if base[i] != input[i] {
				return base
			}
		}

		return nil
	}

	// Reduce common data
	clone := ld.Clone()
	clone.SkipWords = reduceSlice(clone.SkipWords, input.SkipWords)
	clone.PertainWords = reduceSlice(clone.PertainWords, input.PertainWords)
	clone.Abbreviations = reduceSlice(clone.Abbreviations, input.Abbreviations)
	clone.Simplifications = reduceSimplifications(clone.Simplifications, input.Simplifications)
	clone.Translations = reduceMapSlice(clone.Translations, input.Translations)
	clone.RelativeType = reduceMap(clone.RelativeType, input.RelativeType)
	clone.RelativeTypeRegexes = reduceMap(clone.RelativeTypeRegexes, input.RelativeTypeRegexes)
	clone.CombinedRegexPattern = reducePattern(clone.CombinedRegexPattern, input.CombinedRegexPattern)
	clone.ExactCombinedRegexPattern = reducePattern(clone.ExactCombinedRegexPattern, input.ExactCombinedRegexPattern)

	// Merge known words
	clone.KnownWords = reduceKnownWords(clone.KnownWords, input.KnownWords)

	// Reduce charset
	clone.Charset = reduceMap(clone.Charset, input.Charset)

	// Apply sentence splitter group
	if clone.SentenceSplitterGroup == input.SentenceSplitterGroup {
		clone.SentenceSplitterGroup = 0
	}

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
						ShortItf any    `json:"short"`
						Short    string `json:"-"`
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

// SupplementaryData is struct that contains YAML data for `data-supplementary/date-translation`.
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

type MapSliceEntry struct {
	Key    string
	Values []string
}

type OrderEntry struct {
	Key   string
	Order int
}
