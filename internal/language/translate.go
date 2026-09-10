package language

import (
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

var rxRussianNumberPair = regexp.MustCompile(`\b(\d+)[\s\p{Z}]+(\d+)\b`)

// Translate the date string `str` to its English equivalent using information from the locale data.
// If `keepFormatting` is set to true, retain formatting of the date string after translation.
func Translate(cfg *setting.Configuration, ld *data.LocaleData, str string, keepFormatting bool, ignoreSurroundingText ...bool) []string {
	// Parse config
	skippedTokens := mapSkippedTokens(cfg, ld)

	// Normalize and simplify the string
	str = strings.ToLower(strutil.NormalizeUnicode(str))
	str = digit.NormalizeString(str)
	str = Simplify(ld, str)

	// Split string to tokens
	tokens := Split(ld, str, keepFormatting, skippedTokens)
	if len(ignoreSurroundingText) > 0 && ignoreSurroundingText[0] {
		tokens = stripUnknownEdgeTokens(ld, tokens)
	}

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
					if t == "" && keepFormatting && strings.IndexFunc(token, func(character rune) bool {
						return !unicode.IsLetter(character)
					}) >= 0 {
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
		if slices.Contains(tokens, "in") {
			tokens = clearFutureWords(tokens)
		}

		// Remove empty tokens
		tokens = removeEmptyTokens(tokens)

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

	if ld.Name == "ru" {
		var translated strings.Builder
		var last int
		for _, position := range rxRussianNumberPair.FindAllStringIndex(str, -1) {
			pair := str[position[0]:position[1]]
			numbers := strings.Fields(pair)
			first, _ := strconv.Atoi(numbers[0])
			second, _ := strconv.Atoi(numbers[1])
			translated.WriteString(str[last:position[0]])
			if (first == 20 || first == 30) && second >= 1 && second <= 9 && first+second <= 31 {
				translated.WriteString(strconv.Itoa(first + second))
			} else {
				translated.WriteString(pair)
			}
			last = position[1]
		}
		translated.WriteString(str[last:])
		str = translated.String()
	}

	return str
}

func clearFutureWords(words []string) []string {
	// Check if words has freshness word
	var hasFreshness bool
	for _, word := range words {
		isFreshWord := freshnessWords.Contain(word)
		hasFreshness = hasFreshness || isFreshWord
	}

	if !hasFreshness {
		if index := slices.Index(words, "in"); index >= 0 {
			words[index] = ""
		}
	}

	return words
}
