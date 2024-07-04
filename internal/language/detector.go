package language

import (
	"fmt"
	"sort"
	"unicode"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func DetectFullTextLanguage(cfg *setting.Configuration, str string,
	languages []string, uniqueCharsets map[string][]rune) (string, error) {
	// Fetch language candidates for the string
	candidates := getLanguageCandidates(str, languages, uniqueCharsets)
	if len(candidates) == 1 {
		return candidates[0], nil
	}

	// Normalize string
	str = strutil.NormalizeString(str)

	// Create the candidate score
	var finalCandidates []string
	candidateScores := map[string][]int{}
	for _, candidate := range candidates {
		ld, exist := data.GetLocaleData(candidate)
		if !exist {
			continue
		}

		nWord, nSkipped := CountApplicability(cfg, ld, str, false)
		if nWord > 0 || nSkipped > 0 {
			finalCandidates = append(finalCandidates, candidate)
			candidateScores[candidate] = []int{nWord, nSkipped}
		} else {
			nWord, nSkipped = CountApplicability(cfg, ld, str, true)
			if nWord > 0 || nSkipped > 0 {
				finalCandidates = append(finalCandidates, candidate)
				candidateScores[candidate] = []int{nWord, nSkipped}
			}
		}
	}

	if len(finalCandidates) == 0 {
		// Fallback to default language
		if len(cfg.DefaultLanguages) > 0 {
			return cfg.DefaultLanguages[0], nil
		}

		return "", fmt.Errorf("detector failed to find the suitable language")
	}

	// Fetch candidate with the best score
	sort.SliceStable(finalCandidates, func(a, b int) bool {
		cA, cB := finalCandidates[a], finalCandidates[b]
		scoreA, scoreB := candidateScores[cA], candidateScores[cB]

		if scoreA[0] != scoreB[0] {
			return scoreA[0] > scoreB[0]
		}

		if scoreA[1] != scoreB[1] {
			return scoreA[1] > scoreB[1]
		}

		return false
	})

	return finalCandidates[0], nil
}

func getLanguageCandidates(str string, languages []string, uniqueCharsets map[string][]rune) []string {
	// Normalize charset
	str = strutil.NormalizeCharset(str)

	// Convert string to charset (runes)
	strCharset := map[rune]struct{}{}
	for _, r := range str {
		strCharset[r] = struct{}{}
	}

	// If this string only uses common chars, any language will do
	onlyUseCommon := true
	for char := range strCharset {
		if !unicode.Is(commonCharset, char) {
			onlyUseCommon = false
			break
		}
	}

	if onlyUseCommon {
		return []string{languages[0]}
	}

	// Check if this string use unique charsets
	for _, language := range languages {
		for _, char := range uniqueCharsets[language] {
			if _, exist := strCharset[char]; exist {
				return []string{language}
			}
		}
	}

	// Finally, use any languages whose chars used in string
	var usedLanguages []string
	for _, language := range languages {
		ld, exist := data.GetLocaleData(language)
		if !exist {
			continue
		}

		var charsetUsed bool
		for _, char := range ld.Charset {
			if _, exist := strCharset[char]; exist {
				charsetUsed = true
				break
			}
		}

		if charsetUsed {
			usedLanguages = append(usedLanguages, language)
		}
	}

	return usedLanguages
}
