package language

import (
	"slices"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func mapSkippedTokens(cfg *setting.Configuration, ld *data.LocaleData) strutil.Dict {
	skippedTokens := strutil.NewDict()
	if cfg != nil {
		skippedTokens.Add(cfg.SkipTokens...)
	}

	if ld != nil && ld.Name == "fi" {
		skippedTokens.Remove("t")
	}

	return skippedTokens
}

func isSpaceToken(token string) bool {
	return token != "" && strings.Trim(token, " ") == ""
}

func removeEmptyTokens(tokens []string) []string {
	filtered := make([]string, 0, len(tokens))
	for index := 0; index < len(tokens); index++ {
		if tokens[index] != "" {
			filtered = append(filtered, tokens[index])
			continue
		}

		previousStart := len(filtered)
		previousSpaces := 0
		for previousStart > 0 && isSpaceToken(filtered[previousStart-1]) {
			previousStart--
			previousSpaces += len(filtered[previousStart])
		}
		next := index + 1
		nextSpaces := 0
		for next < len(tokens) && isSpaceToken(tokens[next]) {
			nextSpaces += len(tokens[next])
			next++
		}
		if previousSpaces > 0 && nextSpaces > 0 {
			filtered = filtered[:previousStart]
			filtered = append(filtered, strings.Repeat(" ", max(previousSpaces, nextSpaces)))
			index = next - 1
		}
	}
	return filtered
}

func join(tokens []string, separator string) string {
	if len(tokens) == 0 {
		return ""
	}

	joined := tokens[0]
	for i := 1; i < len(tokens); i++ {
		left, right := tokens[i-1], tokens[i]
		leftAlwaysKept := alwaysKeptTokens.Contain(left) || isSpaceToken(left)
		rightAlwaysKept := alwaysKeptTokens.Contain(right) || isSpaceToken(right)
		if !leftAlwaysKept && !rightAlwaysKept {
			joined += separator
		}

		joined += right
	}

	joined = strings.TrimSpace(joined)
	return joined
}

func translateWord(ld *data.LocaleData, word string) ([]string, bool) {
	if translation, exist := ld.RelativeType[word]; exist {
		return []string{translation}, true
	}

	if translations, exist := ld.Translations[word]; exist {
		return slices.Clone(translations), true
	}

	return nil, false
}

func isInDictionary(ld *data.LocaleData, word string) bool {
	_, exist := translateWord(ld, word)
	return exist
}

func createPermutation[T comparable](input [][]T) [][]T {
	// Calculate count of possible permutation
	nPermutation := 1
	for _, entries := range input {
		nPermutation *= len(entries)
	}

	// Create result container
	inputSize := len(input)
	results := make([][]T, nPermutation)
	for i := range results {
		results[i] = make([]T, inputSize)
	}

	// Fill the permutation
	currentTotal := nPermutation
	for i, entries := range input {
		nEntry := len(entries)
		maxEntrySubmission := currentTotal / nEntry

		var entryIdx, nEntrySubmitted int
		for permutationIdx := range nPermutation {
			if nEntrySubmitted >= maxEntrySubmission {
				entryIdx++
				nEntrySubmitted = 0
			}

			if entryIdx >= nEntry {
				entryIdx = 0
			}

			results[permutationIdx][i] = entries[entryIdx]
			nEntrySubmitted++
		}

		currentTotal /= nEntry
	}

	return results
}
