package language

import (
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

func join(tokens []string, separator string) string {
	if len(tokens) == 0 {
		return ""
	}

	joined := tokens[0]
	for i := 1; i < len(tokens); i++ {
		left, right := tokens[i-1], tokens[i]
		leftAlwaysKept := alwaysKeptTokens.Contain(left)
		rightAlwaysKept := alwaysKeptTokens.Contain(right)
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
		clone := append([]string{}, translations...)
		return clone, true
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
		for permutationIdx := 0; permutationIdx < nPermutation; permutationIdx++ {
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
