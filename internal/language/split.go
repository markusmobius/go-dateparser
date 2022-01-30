package language

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

// Split splits the date string `str` using translations in locale data. If `keepFormatting` is
// set to true, it will retaing the formatting of date string.
func Split(ld *data.LocaleData, str string, keepFormatting bool, skippedTokens strutil.Dict) []string {
	// Split the strings
	if ld.RxCombined != nil {
		str = splitWithRegex(str, ld.RxCombined)
	}

	var tokens []string
	for _, token := range strings.Split(str, splitSeparator) {
		if ld.RxExactCombined != nil && ld.RxExactCombined.MatchString(token) {
			tokens = append(tokens, token)
		} else {
			// HEYY
			tokens = append(tokens, splitByKnownWords(ld, token, keepFormatting)...)
		}
	}

	if len(skippedTokens) > 0 {
		var finalTokens []string
		for _, token := range tokens {
			trimmedToken := strings.TrimSpace(token)
			if !skippedTokens.Contain(trimmedToken) {
				finalTokens = append(finalTokens, token)
			}
		}
		tokens = finalTokens
	}

	return tokens
}

func SplitSentence(ld *data.LocaleData, str string) []string {
	// Prepare helper function
	isAbbreviation := func(s string) bool {
		for i := range ld.Abbreviations {
			if s == ld.Abbreviations[i] {
				return true
			}
		}
		return false
	}

	// Get splitter
	splitterGroup := ld.SentenceSplitterGroup
	if splitterGroup < 1 || splitterGroup > 6 {
		splitterGroup = 1
	}
	rxSplitter := rxSentenceSplitters[splitterGroup]

	// Split sentences
	var sentences []string
	var lastPosition int
	useDigitAbbrs := langWithDigitAbbrs.Contain(ld.Name)
	for _, pos := range rxSplitter.FindAllStringSubmatchIndex(str, -1) {
		// Fetch end of sentence from this match
		eos := str[pos[2]:pos[3]]

		// If end of sentence is abbreviation, continue
		if isAbbreviation(eos) {
			continue
		}

		// If this language use digit abbreviation, make sure end of sentence is not number
		if useDigitAbbrs && rxNumberOnly.MatchString(eos) {
			continue
		}

		// Extract this sentence
		sentence := strings.TrimSpace(str[lastPosition:pos[3]])
		if sentence != "" {
			sentences = append(sentences, sentence)
		}

		lastPosition = pos[1]
	}

	// Fetch last sentence (if exist)
	lastSentence := strings.TrimSpace(str[lastPosition:])
	if lastSentence != "" {
		sentences = append(sentences, lastSentence)
	}

	return sentences
}

func SimpleSplit(ld *data.LocaleData, str string, keepFormatting bool, skippedTokens strutil.Dict) []string {
	// Split by numerals
	var lastPos int
	var tokens []string
	for _, pos := range rxNumeral.FindAllStringSubmatchIndex(str, -1) {
		tokens = append(tokens, str[lastPos:pos[0]])
		tokens = append(tokens, str[pos[0]:pos[1]])
		lastPos = pos[1]
	}
	tokens = append(tokens, str[lastPos:])

	// Split by known word
	var finalTokens []string
	for _, token := range tokens {
		if token != "" {
			subtokens := Split(ld, token, keepFormatting, skippedTokens)
			finalTokens = append(finalTokens, subtokens...)
		}
	}

	return finalTokens
}

func splitWithRegex(s string, rx *regexp.Regexp) string {
	var tmp string
	for {
		pos := rx.FindStringSubmatchIndex(s)
		lenPos := len(pos)
		if lenPos == 0 {
			break
		}

		tmp += s[:pos[0]]

		switch lenPos {
		case 2 * 2:
			tmp += fmt.Sprintf("%s%s%s",
				splitSeparator,
				s[pos[2]:pos[3]],
				splitSeparator)

		case 4 * 2:
			tmp += fmt.Sprintf("%s%s%s%s%s",
				s[pos[2]:pos[3]],
				splitSeparator,
				s[pos[4]:pos[5]],
				splitSeparator,
				s[pos[6]:pos[7]])

		default:
			tmp += fmt.Sprintf("%s%s%s",
				splitSeparator,
				s[pos[0]:pos[1]],
				splitSeparator)
		}

		s = s[pos[1]:]
	}

	s = tmp + s
	return s
}

func splitByKnownWords(ld *data.LocaleData, str string, keepFormatting bool) []string {
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
	return keepFormatting ||
		alwaysKeptTokens.Contain(token) ||
		rxKeepToken1.MatchString(token) ||
		rxKeepToken2.MatchString(token)
}
