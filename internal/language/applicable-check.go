package language

import (
	"slices"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

var applicabilityDictionaries sync.Map

func applicabilityDictionary(ld *data.LocaleData) map[string]bool {
	if dictionary, exists := applicabilityDictionaries.Load(ld); exists {
		return dictionary.(map[string]bool)
	}
	dictionary := map[string]bool{}
	addWords := func(phrase string, meaningful bool) {
		words := []string{phrase}
		if strings.Contains(phrase, " ") {
			words = strings.Fields(phrase)
		}
		for _, word := range words {
			dictionary[word] = dictionary[word] || meaningful
		}
	}
	for phrase, translations := range ld.Translations {
		addWords(phrase, slices.ContainsFunc(translations, func(translation string) bool { return translation != "" }))
	}
	for phrase, translation := range ld.RelativeType {
		addWords(phrase, translation != "")
	}
	stored, _ := applicabilityDictionaries.LoadOrStore(ld, dictionary)
	return stored.(map[string]bool)
}

// IsApplicable checks the specified locale data is applicable to translate the date string `str`.
// The `str` parameter is a string representing date and/or time in a recognizably valid format.
// If `stripTimezone` set to true, timezone will be stripped and ignored.
func IsApplicable(cfg *setting.Configuration, ld *data.LocaleData, str string, stripTimezone bool, ignoreSurroundingText ...bool) bool {
	// Strip timezone if needed
	if stripTimezone {
		str, _ = timezone.PopTzOffset(str)
	}

	// Normalize string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)
	return IsApplicablePrepared(cfg, ld, str, ignoreSurroundingText...)
}

func IsApplicablePrepared(cfg *setting.Configuration, ld *data.LocaleData, str string, ignoreSurroundingText ...bool) bool {
	skippedTokens := mapSkippedTokens(cfg, ld)
	str = Simplify(ld, str)

	// Generate tokens
	tokens := Split(ld, str, false, skippedTokens)
	if len(ignoreSurroundingText) > 0 && ignoreSurroundingText[0] {
		tokens = stripUnknownEdgeTokens(ld, tokens)
	}

	// Check if tokens valid
	// First check if tokens only consist of tokens that must be kept
	var nKeptTokens int
	for _, token := range tokens {
		if alwaysKeptTokens.Contain(token) {
			nKeptTokens++
		}
	}

	if nKeptTokens == len(tokens) {
		return false
	}

	// Check if token exist in locale data
	for _, token := range tokens {
		isSkipped := skippedTokens.Contain(token)
		isNumberOnly := strutil.IsNumberOnly(token)
		inDictionary := isInDictionary(ld, token)

		exactCombinedMatch := ld.MatchExactCombined(token)

		if isNumberOnly || inDictionary || isSkipped || exactCombinedMatch {
			continue
		}

		return false
	}

	return true
}

func CountApplicability(cfg *setting.Configuration, ld *data.LocaleData, str string, stripTimezone bool) (int, int) {
	// Parse config
	skippedTokens := strutil.NewDict()
	if cfg != nil {
		skippedTokens.Add(cfg.SkipTokens...)
	}

	// Strip timezone if needed
	if stripTimezone {
		str, _ = timezone.PopTzOffset(str)
	}

	// Split string to sentences
	str = Simplify(ld, str)
	sentences := SplitSentence(ld, str)

	// Extract tokens from sentences
	tokens := strutil.NewDict()
	for _, sentence := range sentences {
		sentenceTokens := simpleSplit(ld, sentence, false, skippedTokens)
		tokens.Add(sentenceTokens...)
	}

	// Count token that exist in dictionary
	var nExist, nSkipped int
	dictionary := applicabilityDictionary(ld)
	for token := range tokens {
		meaningful, exist := dictionary[token]
		if exist && utf8.RuneCountInString(token) >= 2 {
			if meaningful {
				nExist++
			} else {
				nSkipped++
			}
		} else if strutil.IsNumberOnly(token) {
			nSkipped++
		}
	}

	return nExist, nSkipped
}
