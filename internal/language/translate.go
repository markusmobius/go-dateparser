package language

import (
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

// Translate the date string `str` to its English equivalent using information from the locale data.
// If `keepFormatting` is set to true, retain formatting of the date string after translation.
func Translate(cfg *setting.Configuration, ld *data.LocaleData, str string, keepFormatting bool) string {
	// Parse config
	skippedTokens := mapSkippedTokens(cfg, ld)

	// Normalize and simplify the string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)
	str = Simplify(ld, str)

	// Split string to tokens
	inInTokens := false
	tokens := Split(ld, str, keepFormatting, skippedTokens)

	// Translate
	for i, token := range tokens {
		// Check if token skipped
		if skippedTokens.Contain(token) {
			tokens[i] = ""
			continue
		}

		// Try to use regex to translate relative type
		var translationFound bool
		for _, data := range ld.RelativeTypeRegexes {
			if data.Rx.MatchString(token) {
				token = data.Rx.ReplaceAllString(token, data.Replacement)
				translationFound = true
				break
			}
		}

		// If not found, look in dictionary
		if !translationFound {
			if translation, exist := translateWord(ld, token); exist {
				token = translation
			}
		}

		inInTokens = inInTokens || token == "in"
		tokens[i] = token
	}

	// Handle future words
	if inInTokens {
		tokens = clearFutureWords(tokens)
	}

	// Join the tokens
	var validTokens []string
	for _, token := range tokens {
		if token != "" {
			validTokens = append(validTokens, token)
		}
	}

	joinSeparator := ""
	if !keepFormatting {
		joinSeparator = " "
	}

	return join(validTokens, joinSeparator)
}

func Simplify(ld *data.LocaleData, str string) string {
	for _, data := range ld.Simplifications {
		if data.Rx.MatchString(str) {
			str = data.Rx.ReplaceAllString(str, data.Replacement)
		}
	}

	return str
}

func clearFutureWords(words []string) []string {
	// Check if words has freshness word
	var hasFreshness bool
	var wordsWithoutIn []string
	for _, word := range words {
		isFreshWord := freshnessWords.Contain(word)
		hasFreshness = hasFreshness || isFreshWord

		if word != "in" {
			wordsWithoutIn = append(wordsWithoutIn, word)
		}
	}

	if !hasFreshness {
		return wordsWithoutIn
	}

	return words
}
