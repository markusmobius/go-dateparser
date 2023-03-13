package language

import (
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/digit"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
	"github.com/markusmobius/go-dateparser/internal/timezone"
)

func TranslateSearch(cfg *setting.Configuration, ld *data.LocaleData, str string) ([]string, []string) {
	// Parse config
	skippedTokens := mapSkippedTokens(cfg, ld)

	// Split string to sentences
	var original [][]string
	var translated [][]string
	sentences := SplitSentence(ld, str)
	joinUnusable := ld.Name == "zh" || ld.Name == "ja"

	for _, sentence := range sentences {
		var originalChunk []string
		var translatedChunk [][]string
		var skipNextToken bool

		originalTokens, simplifiedTokens := simplifySplitAlign(ld, sentence, skippedTokens)
		lastTokenIndex := len(simplifiedTokens) - 1

		for i, word := range simplifiedTokens {
			var nextWord string
			if i < lastTokenIndex {
				nextWord = simplifiedTokens[i+1]
			}

			currentAndNextJoined := joinChunk(ld, word, nextWord)
			if skipNextToken {
				skipNextToken = false
				continue
			}

			cleanedWord := strutil.TrimChars(word, `()"'{}[],.،`)

			if word == "" || word == " " {
				translatedChunk = append(translatedChunk, []string{word})
				originalChunk = append(originalChunk, originalTokens[i])
			} else if isInDictionary(ld, currentAndNextJoined) && !dashes.Contain(word) && !joinUnusable {
				translations, _ := translateWord(ld, currentAndNextJoined)
				joinedOriginalToken := joinChunk(ld, originalTokens[i], originalTokens[i+1])

				skipNextToken = true
				translatedChunk = append(translatedChunk, translations)
				originalChunk = append(originalChunk, joinedOriginalToken)
			} else if isInDictionary(ld, word) && !dashes.Contain(word) {
				translations, _ := translateWord(ld, word)
				translatedChunk = append(translatedChunk, translations)
				originalChunk = append(originalChunk, originalTokens[i])
			} else if isInDictionary(ld, cleanedWord) && !dashes.Contain(word) {
				punct := word[len(cleanedWord):]
				translations, _ := translateWord(ld, cleanedWord)
				for i, translation := range translations {
					translations[i] = translation + punct
				}

				translatedChunk = append(translatedChunk, translations)
				originalChunk = append(originalChunk, originalTokens[i])
			} else if tokenWithDigitsIsOk(ld, word) {
				translatedChunk = append(translatedChunk, []string{word})
				originalChunk = append(originalChunk, originalTokens[i])
			} else if len(translatedChunk) > 0 && timezone.WordIsTz(originalTokens[i]) {
				translatedChunk = append(translatedChunk, []string{word})
				originalChunk = append(originalChunk, originalTokens[i])
			} else if len(translatedChunk) > 0 {
				translatedChunkPermutations := createPermutation(translatedChunk)
				translated = append(translated, translatedChunkPermutations...)
				for i := 0; i < len(translatedChunkPermutations); i++ {
					original = append(original, originalChunk)
				}

				originalChunk, translatedChunk = []string{}, [][]string{}
			}
		}

		if len(translatedChunk) > 0 {
			translatedChunkPermutations := createPermutation(translatedChunk)
			translated = append(translated, translatedChunkPermutations...)
			for i := 0; i < len(translatedChunkPermutations); i++ {
				original = append(original, originalChunk)
			}
		}
	}

	// Merge translation
	var mergedOriginal []string
	var mergedTranslation []string

	for i := range translated {
		originalChunk := filterChunk(original[i]...)
		translatedChunk := filterChunk(clearFutureWords(translated[i])...)

		mergedOriginal = append(mergedOriginal, joinChunk(ld, originalChunk...))
		mergedTranslation = append(mergedTranslation, joinChunk(ld, translatedChunk...))
	}

	return mergedTranslation, mergedOriginal
}

func simplifySplitAlign(ld *data.LocaleData, str string, skippedTokens strutil.Dict) ([]string, []string) {
	normalizeToken := func(token string) string {
		token = strutil.NormalizeString(token)
		token = digit.NormalizeString(token)
		return token
	}

	// Split tokens
	simplifiedStr := Simplify(ld, normalizeToken(str))
	originalTokens := wordSplit(ld, str, skippedTokens)
	simplifiedTokens := wordSplit(ld, simplifiedStr, skippedTokens)

	// If both token have same count, return it
	if len(originalTokens) == len(simplifiedTokens) {
		return originalTokens, simplifiedTokens
	}

	if len(originalTokens) < len(simplifiedTokens) {
		var addEmpty bool
		for i, token := range simplifiedTokens {
			if i < len(originalTokens) {
				if token == normalizeToken(originalTokens[i]) {
					addEmpty = false
				} else {
					if !addEmpty {
						addEmpty = true
						continue
					} else {
						originalTokens = insertSlice(originalTokens, i, "")
					}
				}
			} else {
				originalTokens = append(originalTokens, "")
			}
		}
	} else {
		var addEmpty bool
		for i, token := range originalTokens {
			if i < len(simplifiedTokens) {
				if normalizeToken(token) == simplifiedTokens[i] {
					addEmpty = false
				} else {
					if !addEmpty {
						addEmpty = true
						continue
					} else {
						simplifiedTokens = insertSlice(simplifiedTokens, i, "")
					}
				}
			} else {
				simplifiedTokens = append(simplifiedTokens, "")
			}
		}
	}

	var removed bool
	for len(originalTokens) != len(simplifiedTokens) {
		if len(originalTokens) > len(simplifiedTokens) {
			originalTokens, removed = removeEmptySliceItem(originalTokens)
		} else {
			simplifiedTokens, removed = removeEmptySliceItem(simplifiedTokens)
		}

		if !removed {
			break
		}
	}

	return originalTokens, simplifiedTokens
}

func wordSplit(ld *data.LocaleData, str string, skippedTokens strutil.Dict) []string {
	if ld.NoWordSpacing {
		return simpleSplit(ld, str, true, skippedTokens)
	} else {
		return strings.Fields(str)
	}
}

func joinChunk(ld *data.LocaleData, chunk ...string) string {
	if ld.NoWordSpacing {
		return join(chunk, "")
	} else {
		str := strings.Join(chunk, " ")
		str = rxSpaces.ReplaceAllString(str, " ")
		return str
	}
}

func filterChunk(chunk ...string) []string {
	var filtered []string
	for _, c := range chunk {
		if c != "" {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func tokenWithDigitsIsOk(ld *data.LocaleData, token string) bool {
	if ld.NoWordSpacing {
		return rxTokenDigit.MatchString(token)
	} else {
		return rxNumeral.MatchString(token)
	}
}

func insertSlice(list []string, i int, s string) []string {
	if i < 0 {
		i = 0
	} else if i > len(list) {
		i = len(list)
	}

	list = append(list, "")
	copy(list[i+1:], list[i:])
	list[i] = s

	return list
}

func removeEmptySliceItem(list []string) ([]string, bool) {
	emptyIdx := -1
	for i := range list {
		if list[i] == "" {
			emptyIdx = i
			break
		}
	}

	if emptyIdx < 0 {
		return list, false
	}

	list = append(list[:emptyIdx], list[emptyIdx+1:]...)
	return list, true
}
