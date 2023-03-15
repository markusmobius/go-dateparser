package language

import (
	"unicode/utf8"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

// IsApplicable checks the specified locale data is applicable to translate the date string `str`.
// The `str` parameter is a string representing date and/or time in a recognizably valid format.
// If `stripTimezone` set to true, timezone will be stripped and ignored.
func IsApplicable(cfg *setting.Configuration, ld *data.LocaleData, str string, stripTimezone bool) bool {
	// Parse config
	skippedTokens := mapSkippedTokens(cfg, ld)

	// Strip timezone if needed
	if stripTimezone {
		str, _ = timezone.PopTzOffset(str)
	}

	// Normalize string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)
	str = Simplify(ld, str)

	// Generate tokens
	tokens := Split(ld, str, false, skippedTokens)

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

		var exactCombinedMatch bool
		if ld.RxExactCombined != nil {
			exactCombinedMatch = ld.RxExactCombined.MatchString(token)
		}

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
	for token := range tokens {
		translations, exist := ld.Translations[token]
		if exist && utf8.RuneCountInString(token) >= 2 {
			if len(translations) > 0 {
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
