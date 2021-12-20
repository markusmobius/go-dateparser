package language

import (
	"fmt"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
)

func Translate(ld data.LocaleData, str string, keepFormatting bool, config *setting.Configuration) string {
	// str = "2 months ago, friday, in 10 days, 03 september 2014"
	str = normalizeString(str)
	str = digit.NormalizeString(str)

	// Simplify the string
	for _, data := range ld.Simplifications {
		if data.Rx.MatchString(str) {
			str = data.Rx.ReplaceAllString(str, data.Replacement)
		}
	}

	// Split string to tokens
	inInTokens := false
	tokens := Split(ld, str, keepFormatting)

	// Translate
	skippedTokens := map[string]struct{}{}
	if config != nil {
		for _, token := range config.SkipTokens {
			skippedTokens[token] = struct{}{}
		}
	}

	for i, token := range tokens {
		// Check if token skipped
		if _, skipped := skippedTokens[token]; skipped {
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

		// If not found, use dictionary of relative type
		if !translationFound {
			if translation, exist := ld.RelativeType[token]; exist {
				token = translation
				translationFound = true
			}
		}

		// If still not found, use words dictionary
		if !translationFound {
			if translation, exist := ld.Translations[token]; exist {
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

func Split(ld data.LocaleData, str string, keepFormatting bool) []string {
	// Split the strings
	if ld.RxCombined != nil {
		str = ld.RxCombined.ReplaceAllStringFunc(str, func(s string) string {
			parts := ld.RxCombined.FindStringSubmatch(s)
			switch len(parts) {
			case 2:
				return fmt.Sprintf("%s%s%s",
					splitSeparator,
					parts[1],
					splitSeparator,
				)

			case 4:
				return fmt.Sprintf("%s%s%s%s%s",
					parts[1],
					splitSeparator,
					parts[2],
					splitSeparator,
					parts[3],
				)

			default:
				return fmt.Sprintf("%s%s%s",
					splitSeparator,
					s,
					splitSeparator,
				)
			}
		})
	}

	var tokens []string
	for _, token := range strings.Split(str, splitSeparator) {
		if ld.RxExactCombined != nil && ld.RxExactCombined.MatchString(token) {
			tokens = append(tokens, token)
		} else {
			tokens = append(tokens, splitByKnownWords(ld, token, keepFormatting)...)
		}
	}

	return tokens
}

func splitByKnownWords(ld data.LocaleData, str string, keepFormatting bool) []string {
	var matches []string
	if ld.RxKnownWords != nil {
		matches = ld.RxKnownWords.FindStringSubmatch(str)
	}

	if len(matches) == 0 {
		if tokenShouldBeCaptured(str, keepFormatting) {
			return splitByNumerals(str, keepFormatting)
		} else {
			return nil
		}
	}

	var splitted []string
	unparsed, known, unknown := matches[1], matches[2], matches[3]

	if tokenShouldBeCaptured(known, keepFormatting) {
		splitted = []string{known}
	}

	if unparsed != "" && tokenShouldBeCaptured(unparsed, keepFormatting) {
		splitted = append(splitByNumerals(unparsed, keepFormatting), splitted...)
	}

	if unknown != "" {
		splitted = append(splitted, splitByKnownWords(ld, unknown, keepFormatting)...)
	}

	return splitted
}

func splitByNumerals(str string, keepFormatting bool) []string {
	str = rxNumeral.ReplaceAllStringFunc(str, func(s string) string {
		return splitSeparator + s + splitSeparator
	})

	var validTokens []string
	tokens := strings.Split(str, splitSeparator)
	for _, token := range tokens {
		if tokenShouldBeCaptured(token, keepFormatting) {
			validTokens = append(validTokens, token)
		}
	}

	return validTokens
}

func tokenShouldBeCaptured(token string, keepFormatting bool) bool {
	_, alwaysKept := alwaysKeptTokens[token]
	return keepFormatting || alwaysKept ||
		rxKeepToken1.MatchString(token) ||
		rxKeepToken2.MatchString(token)
}

func clearFutureWords(words []string) []string {
	// Check if words has freshness word
	var hasFreshness bool
	var wordsWithoutIn []string
	for _, word := range words {
		_, isFreshWord := freshnessWords[word]
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

func join(tokens []string, separator string) string {
	if len(tokens) == 0 {
		return ""
	}

	joined := tokens[0]
	for i := 1; i < len(tokens); i++ {
		left, right := tokens[i-1], tokens[i]
		_, leftAlwaysKept := alwaysKeptTokens[left]
		_, rightAlwaysKept := alwaysKeptTokens[right]
		if !leftAlwaysKept && !rightAlwaysKept {
			joined += separator
		}

		joined += right
	}

	joined = strings.TrimSpace(joined)
	return joined
}
