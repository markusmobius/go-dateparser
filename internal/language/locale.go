package language

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
)

func Translate(ld data.LocaleData, str string, keepFormatting bool) string {
	// str = "2 months ago, friday, in 10 days, 03 september 2014"
	str = normalizeString(str)

	// Simplify the string
	for _, data := range ld.Simplifications {
		str = data.Rx.ReplaceAllString(str, data.Replacement)
	}

	// Split string to tokens
	inInTokens := false
	tokens := Split(ld, str, keepFormatting)

	for i, token := range tokens {
		if translation, exist := ld.Translations[token]; exist {
			token = translation
		} else {
			for _, data := range ld.TranslationRegexes {
				if data.Rx.MatchString(token) {
					token = data.Rx.ReplaceAllString(token, data.Replacement)
					break
				}
			}
		}

		inInTokens = inInTokens || token == "in"
		tokens[i] = token
	}

	// Handle future words
	fmt.Println("IN IN TOKENS:", inInTokens)
	fmt.Println("FIRST TRANSLATION:", strJson(tokens))

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

	translation := join(validTokens, joinSeparator)
	fmt.Println("JOINED TRANSLATION:", translation)
	return translation
}

func Split(ld data.LocaleData, str string, keepFormatting bool) []string {
	fmt.Println("INITIAL STR:", str)
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

	// str = strings.Trim(str, splitSeparator)

	fmt.Println("RELATIVE:", ld.RxCombined)
	fmt.Println("EXACT MATCH:", ld.RxExactCombined)
	fmt.Println("STR SPLIT:", str)
	fmt.Println("INITIAL SPLIT:", strJson(strings.Split(str, splitSeparator)))

	var tokens []string
	for _, token := range strings.Split(str, splitSeparator) {
		if ld.RxExactCombined != nil && ld.RxExactCombined.MatchString(token) {
			tokens = append(tokens, token)
		} else {
			tokens = append(tokens, splitByKnownWords(ld, token, keepFormatting)...)
		}
	}

	fmt.Println("FINAL SPLIT:", strJson(tokens))

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

func strJson(v interface{}) string {
	bt, _ := json.Marshal(v)
	return string(bt)
}
