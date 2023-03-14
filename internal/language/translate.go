package language

import (
	"github.com/elliotchance/pie/v2"
	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

// Translate the date string `str` to its English equivalent using information from the locale data.
// If `keepFormatting` is set to true, retain formatting of the date string after translation.
func Translate(cfg *setting.Configuration, ld *data.LocaleData, str string, keepFormatting bool) []string {
	// Parse config
	skippedTokens := mapSkippedTokens(cfg, ld)

	// Normalize and simplify the string
	str = strutil.NormalizeString(str)
	str = digit.NormalizeString(str)
	str = Simplify(ld, str)

	// Split string to tokens
	tokens := Split(ld, str, keepFormatting, skippedTokens)

	// Translate each token
	translatedTokens := make([][]string, len(tokens))
	for i, token := range tokens {
		// Check if token skipped
		if skippedTokens.Contain(token) {
			translatedTokens[i] = []string{""}
			continue
		}

		// Try to use regex to translate relative type
		var translationFound bool
		for _, data := range ld.RelativeTypeRegexes {
			if data.Rx.MatchString(token) {
				translation := data.Rx.ReplaceAllString(token, data.Replacement)
				translatedTokens[i] = []string{translation}
				translationFound = true
				break
			}
		}

		// If not found, look in dictionary
		if !translationFound {
			if translations, exist := translateWord(ld, token); exist {
				// If keep formatting, empty translation is replaced
				// with current token
				for j, t := range translations {
					if t == "" && keepFormatting {
						translations[j] = token
					}
				}

				translatedTokens[i] = translations
				translationFound = true
			}
		}

		// If still not found, use token as is
		if !translationFound {
			translatedTokens[i] = []string{token}
		}
	}

	// Create token permutations from translated tokens
	tokenPermutations := createPermutation(translatedTokens)

	// Clean up and finalize translations
	var translations []string
	for _, tokens := range tokenPermutations {
		// Handle future words
		if pie.Contains(tokens, "in") {
			tokens = clearFutureWords(tokens)
		}

		// Exclude empty tokens
		tokens = pie.Filter(tokens, func(t string) bool {
			return t != ""
		})

		// Join the tokens to get final translations
		joinSeparator := ""
		if !keepFormatting {
			joinSeparator = " "
		}

		translation := join(tokens, joinSeparator)
		translations = append(translations, translation)
	}

	return translations
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
