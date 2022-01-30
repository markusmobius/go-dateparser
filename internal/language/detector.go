package language

import (
	"fmt"
	"sort"
	"unicode"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func DetectLanguage(cfg *setting.Configuration, str string, languages []string, detectLanguagesFunction func(string) []string) (string, error) {
	// If custom detect function defined and there are no languages specified,
	// use the custom detect function
	if detectLanguagesFunction != nil && len(languages) == 0 {
		detectedLanguages := detectLanguagesFunction(str)
		if len(detectedLanguages) == 0 && cfg != nil {
			detectedLanguages = append(detectedLanguages, cfg.DefaultLanguages...)
		}

		if len(detectedLanguages) > 0 {
			lang := detectedLanguages[0]
			if _, exist := data.LanguageMap[lang]; !exist {
				return "", fmt.Errorf("detected language is not known by dateparser: %s", lang)
			}
			return lang, nil
		}
	}

	// If list of languages defined, validate it
	// If not, take from the known language order
	if len(languages) > 0 {
		for _, lang := range languages {
			if _, exist := data.LanguageMap[lang]; !exist {
				return "", fmt.Errorf("unknown language: %s", lang)
			}
		}
	} else {
		languages = make([]string, len(data.LanguageOrder))
		for lang, order := range data.LanguageOrder {
			languages[order] = lang
		}
	}

	// Fetch charset for all languages
	charsets := getLanguageCharsets(languages)
	uniqueCharsets := getUniqueCharsets(languages, charsets)

	// Fetch language candidates for the string
	candidates := getLanguageCandidates(str, languages, charsets, uniqueCharsets)
	if len(candidates) == 1 {
		return candidates[0], nil
	}

	// Normalize string
	str = strutil.NormalizeString(str)

	// Create the candidate score
	var finalCandidates []string
	candidateScores := map[string][]int{}
	for _, candidate := range candidates {
		ld := data.LocaleDataMap[candidate]
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

func getLanguageCharsets(languages []string) [][]rune {
	charsets := make([][]rune, len(languages))
	for i, language := range languages {
		locale := data.LocaleDataMap[language]
		charsets[i] = locale.Charset
	}
	return charsets
}

func getUniqueCharsets(languages []string, charsets [][]rune) [][]rune {
	// Prepare result
	uniqueCharsets := make([][]rune, len(charsets))

	// Process each charset
	for i, charset := range charsets {
		language := languages[i]

		// Initiate map to contain the unique chars for this language
		mapUniqueChar := map[rune]struct{}{}
		for _, char := range charset {
			mapUniqueChar[char] = struct{}{}
		}

		// Look for charset in the other languages
		for j, otherCharset := range charsets {
			otherLanguage := languages[j]
			if otherLanguage == language {
				continue
			}

			// If the charset of this language exist in map, remove it
			for _, char := range otherCharset {
				if _, exist := mapUniqueChar[char]; exist {
					delete(mapUniqueChar, char)
				}
			}
		}

		// Convert map to list
		var uniqueCharset []rune
		for char := range mapUniqueChar {
			uniqueCharset = append(uniqueCharset, char)
		}

		uniqueCharsets[i] = uniqueCharset
	}

	return uniqueCharsets
}

func getLanguageCandidates(str string, languages []string, charsets [][]rune, uniqueCharsets [][]rune) []string {
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
	for i, uniqueCharset := range uniqueCharsets {
		for _, char := range uniqueCharset {
			if _, exist := strCharset[char]; exist {
				return []string{languages[i]}
			}
		}
	}

	// Finally, use any languages whose chars used in string
	var usedLanguages []string
	for i, charset := range charsets {
		var charsetUsed bool
		for _, char := range charset {
			if _, exist := strCharset[char]; exist {
				charsetUsed = true
				break
			}
		}

		if charsetUsed {
			usedLanguages = append(usedLanguages, languages[i])
		}
	}

	return usedLanguages
}
