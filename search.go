package dateparser

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"slices"

	"github.com/markusmobius/go-dateparser/date"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/language"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

var (
	splitters           = []string{",", "،", "——", "—", "–", ".", " "}
	relativeWords       = []string{"ago", "in", "from now", "tomorrow", "today", "yesterday"}
	rxRussianSearchFrom = regexp.MustCompile(`(^|[^\p{L}\p{N}_])с[\s\p{Z}]+(\p{Nd})`)
)

// SearchResult pairs a parsed date with matching text or a time-span boundary label.
type SearchResult struct {
	Date date.Date
	Text string
}

type parsedSearch struct {
	Date       date.Date
	IsRelative bool
}

type splitResult struct {
	EntryParts    []string
	OriginalParts []string
}

type splitRating struct {
	Index             int
	NSubString        int
	NotParsedRatio    float64
	WithoutDigitRatio float64
}

// Search detects a language and extracts dates using the configured search strategy.
// It returns the selected language code, matches, and any error.
func (p *Parser) Search(cfg *Configuration, text string) (string, []SearchResult, error) {
	// Prepare config
	cfg, err := p.initSearchConfig(cfg)
	if err != nil {
		return "", nil, err
	}

	if slices.Contains(cfg.Languages, "ru") {
		text = rxRussianSearchFrom.ReplaceAllString(text, "${1}[FROM] ${2}")
	}

	// Get list of used languages
	cfgLanguages := cfg.Languages
	if len(cfgLanguages) == 0 && p.DetectLanguagesFunction != nil {
		cfgLanguages = p.DetectLanguagesFunction(text)
	}

	languages, err := language.GetLanguages(cfg.Locales, cfgLanguages, cfg.UseGivenOrder)
	if err != nil {
		return "", nil, err
	}

	// Generate charsets
	uniqueCharsets := p.initUniqueCharsets(languages)

	// Detect language of the text
	iCfg := cfg.toInternalConfig()
	lang, detectionErr := language.DetectFullTextLanguage(iCfg, text, languages, uniqueCharsets)
	var candidates []string
	if lang != "" {
		candidates = append(candidates, lang)
	}

	if len(cfg.Languages) > 1 || len(cfg.Locales) > 1 {
		fallbacks, err := language.GetLanguages(cfg.Locales, cfg.Languages, true)
		if err != nil {
			return "", nil, err
		}
		for _, fallback := range fallbacks {
			if !slices.Contains(candidates, fallback) {
				candidates = append(candidates, fallback)
			}
		}
	}

	if cfg.SearchStrategy == "ngram" && len(candidates) > 0 {
		return lang, p.searchNgrams(cfg, candidates, lang, text), nil
	}

	for _, candidate := range candidates {
		result, err := p.SearchWithLanguage(cfg, candidate, text)
		if err != nil || len(result) > 0 {
			return candidate, result, err
		}
	}
	return lang, nil, detectionErr
}

// SearchWithLanguage extracts dates using the specified language or locale
// and the configured search strategy.
func (p *Parser) SearchWithLanguage(cfg *Configuration, lang string, text string) ([]SearchResult, error) {
	// Initiate config
	cfg, err := p.initSearchConfig(cfg)
	if err != nil {
		return nil, err
	}
	iCfg := cfg.toInternalConfig()

	// Find the locale data
	ld, exist := data.GetLocaleData(lang)
	if ld == nil || !exist {
		return nil, fmt.Errorf("unknown language: %s", lang)
	}
	if cfg.SearchStrategy == "ngram" {
		return p.searchNgrams(cfg, []string{lang}, lang, text), nil
	}

	// Translate the text
	translation, original := language.TranslateSearch(iCfg, ld, text)

	// Create languages for the temporary parser
	var languages []string
	if lang != "vi" && lang != "hu" {
		languages = []string{"en"}
	} else {
		languages = []string{ld.Name}
	}

	// Parse the text
	parsedSearch, subStrings := p.parseFoundObjects(iCfg, languages, translation, original)

	// Create final result
	result := make([]SearchResult, len(parsedSearch))
	for i, ps := range parsedSearch {
		ps.Date.Locale = lang
		result[i] = SearchResult{
			Date: ps.Date,
			Text: subStrings[i],
		}
	}

	if cfg.ReturnTimeSpan {
		result = append(result, searchTimeSpan(cfg, lang, text)...)
	}

	return result, nil
}

func (p *Parser) initSearchConfig(cfg *Configuration) (*Configuration, error) {
	// Initiate and validate config
	if cfg == nil {
		cfg = &Configuration{}
	}

	cfgHasCurrentTime := !cfg.CurrentTime.IsZero()
	cfg = cfg.initiate()
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	if !cfgHasCurrentTime {
		cfg.CurrentTime = time.Time{}
	}

	return cfg, nil
}

func (p *Parser) initUniqueCharsets(languages []string) map[string][]rune {
	// Lock parser
	p.Lock()
	defer p.Unlock()

	// Prepare map
	if p.uniqueCharsets == nil {
		p.uniqueCharsets = map[string][]rune{}
	}

	// Check if charsets need to be generated
	mustGenerate := len(p.uniqueCharsets) != len(languages)
	if !mustGenerate {
		for _, lang := range languages {
			if _, exist := p.uniqueCharsets[lang]; !exist {
				mustGenerate = true
				break
			}
		}
	}

	if mustGenerate {
		p.uniqueCharsets = language.GetUniqueCharsets(languages)
	}
	return p.uniqueCharsets
}

func (p *Parser) parseFoundObjects(iCfg *setting.Configuration, languages, translation, original []string) ([]parsedSearch, []string) {
	// Specify entries to parse
	var parserEntries []string
	if languages[0] == "en" {
		parserEntries = slices.Clone(translation)
	} else {
		parserEntries = slices.Clone(original)
	}

	// Parse each entries
	var parsedList []parsedSearch
	var subStrings []string
	needRelativeBase := iCfg.CurrentTime.IsZero()

	for i, entry := range parserEntries {
		if utf8.RuneCountInString(entry) <= 2 {
			continue
		}

		parsedEntry, isRelative := p.parseEntry(iCfg, entry, translation[i], parsedList, needRelativeBase)
		if !parsedEntry.IsZero() {
			parsedList = append(parsedList, parsedSearch{parsedEntry, isRelative})
			subStrings = append(subStrings, strutil.TrimChars(original[i], ` .,:()[]-'`))
			continue
		}

		possibleSplits := splitIfNotParsed(entry, original[i])
		if len(possibleSplits) == 0 {
			continue
		}

		var possibleSubStrings [][]string
		var possibleParseResult [][]parsedSearch
		for _, split := range possibleSplits {
			var currentSubStrings []string
			var currentParseResult []parsedSearch

			if len(split.EntryParts) > 0 {
				for j, jEntry := range split.EntryParts {
					if utf8.RuneCountInString(jEntry) <= 2 {
						continue
					}

					parsedJEntry, jIsRelative := p.parseEntry(iCfg, jEntry,
						split.EntryParts[j], currentParseResult, needRelativeBase)

					jSubString := strutil.TrimChars(split.OriginalParts[j], ` .,:()[]-`)
					jParsedSearch := parsedSearch{parsedJEntry, jIsRelative}
					currentParseResult = append(currentParseResult, jParsedSearch)
					currentSubStrings = append(currentSubStrings, jSubString)
				}
			}

			possibleSubStrings = append(possibleSubStrings, currentSubStrings)
			possibleParseResult = append(possibleParseResult, currentParseResult)
		}

		bestParseResult, bestSubString := chooseBestSplit(possibleParseResult, possibleSubStrings)
		for k, parseResult := range bestParseResult {
			if !parseResult.Date.IsZero() {
				parsedList = append(parsedList, parseResult)
				subStrings = append(subStrings, bestSubString[k])
			}
		}
	}

	return parsedList, subStrings
}

func (p *Parser) parseEntry(iCfg *setting.Configuration, entry, translation string, parsedList []parsedSearch, needRelativeBase bool) (date.Date, bool) {
	// Normalize entry
	entry = strings.ReplaceAll(entry, "ngày", "")
	entry = strings.ReplaceAll(entry, "am", "")

	// Parse entry
	cfg := configFromInternal(iCfg)
	parsedEntry, _ := p.Parse(cfg, entry)

	// If needed, generate relative base and parse the entry
	var relativeBase time.Time
	if needRelativeBase {
		relativeBase = getRelativeBase(parsedList)
	}

	if !relativeBase.IsZero() {
		cfg.CurrentTime = relativeBase
		iCfg.CurrentTime = relativeBase
		parsedEntry, _ = p.Parse(cfg, entry)
	}

	// Check if this date is relative
	var isRelative bool
	for _, rw := range relativeWords {
		if strings.Contains(translation, rw) {
			isRelative = true
			break
		}
	}

	return parsedEntry, isRelative
}

func getRelativeBase(parsedList []parsedSearch) time.Time {
	if len(parsedList) == 0 {
		return time.Time{}
	}

	i := len(parsedList) - 1
	for parsedList[i].IsRelative {
		i--
		if i == -1 {
			return time.Time{}
		}
	}

	relativeBase := parsedList[i].Date.Time
	return relativeBase
}

func splitIfNotParsed(entry, original string) []splitResult {
	var possibleSplits []splitResult

	for _, splitter := range splitters {
		if strings.Contains(entry, splitter) &&
			strings.Count(entry, splitter) == strings.Count(original, splitter) {
			possibleSplits = append(possibleSplits, splitBy(entry, original, splitter)...)
		}
	}

	return possibleSplits
}

func splitBy(entry, original, splitter string) []splitResult {
	entryParts := strings.Split(entry, splitter)
	originalParts := strings.Split(original, splitter)
	allSplits := []splitResult{{
		EntryParts:    entryParts,
		OriginalParts: originalParts,
	}}

	if strings.Count(entry, splitter) <= 2 {
		return allSplits
	}

	nEntryParts := len(entryParts)
	for i := 2; i < 4; i++ {
		var entryPartialParts []string
		var originalPartialParts []string

		for j := 0; j < nEntryParts; j += i {
			end := min(j+i, nEntryParts)
			entryJoin := strings.Join(entryParts[j:end], splitter)
			originalJoin := strings.Join(originalParts[j:end], splitter)
			entryPartialParts = append(entryPartialParts, entryJoin)
			originalPartialParts = append(originalPartialParts, originalJoin)
		}

		allSplits = append(allSplits, splitResult{
			EntryParts:    entryPartialParts,
			OriginalParts: originalPartialParts,
		})
	}

	return allSplits
}

func chooseBestSplit(possibleParseResult [][]parsedSearch, possibleSubStrings [][]string) ([]parsedSearch, []string) {
	if len(possibleParseResult) == 0 {
		return nil, nil
	}

	var rating []splitRating
	for i := range possibleParseResult {
		nSubStrings := len(possibleSubStrings[i])

		var nNotParsed int
		var nSubStringsWithoutDigits int
		for j, item := range possibleParseResult[i] {
			if item.Date.IsZero() {
				nNotParsed++
			}

			var hasDigit bool
			for _, r := range possibleSubStrings[i][j] {
				if unicode.IsDigit(r) {
					hasDigit = true
					break
				}
			}

			if !hasDigit {
				nSubStringsWithoutDigits++
			}
		}

		var notParsedRatio, withoutDigitRatio float64
		if nSubStrings > 0 {
			notParsedRatio = float64(nNotParsed) / float64(nSubStrings)
			withoutDigitRatio = float64(nSubStringsWithoutDigits) / float64(nSubStrings)
		}

		rating = append(rating, splitRating{
			Index:             i,
			NSubString:        nSubStrings,
			NotParsedRatio:    notParsedRatio,
			WithoutDigitRatio: withoutDigitRatio,
		})
	}

	// Look for the best rating
	sort.SliceStable(rating, func(a, b int) bool {
		ra, rb := rating[a], rating[b]

		if ra.NotParsedRatio != rb.NotParsedRatio {
			return ra.NotParsedRatio < rb.NotParsedRatio
		} else if ra.NSubString != rb.NSubString {
			return ra.NSubString < rb.NSubString
		} else if ra.WithoutDigitRatio != rb.WithoutDigitRatio {
			return ra.WithoutDigitRatio < rb.WithoutDigitRatio
		}

		return false
	})

	bestIdx := rating[0].Index
	return possibleParseResult[bestIdx], possibleSubStrings[bestIdx]
}
