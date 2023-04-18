package language

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/regexp"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

// Split splits the date string `str` using translations in locale data. If `keepFormatting` is
// set to true, it will retaing the formatting of date string.
func Split(ld *data.LocaleData, str string, keepFormatting bool, skippedTokens strutil.Dict) []string {
	// Split the strings
	if ld.RxCombined != nil {
		str = splitWithRegex(str, ld.RxCombined)
	}

	// Check if each token can be split further
	var tokens []string
	for _, token := range strings.Split(str, splitSeparator) {
		if ld.RxExactCombined != nil && ld.RxExactCombined.MatchString(token) {
			tokens = append(tokens, token)
		} else {
			tokens = append(tokens, splitByKnownWords(ld, token, keepFormatting)...)
		}
	}

	// Filter token to remove empty or skipped token
	var finalTokens []string
	for _, token := range tokens {
		if token == "" {
			continue
		}

		trimmedToken := strings.TrimSpace(token)
		if !skippedTokens.Contain(trimmedToken) {
			finalTokens = append(finalTokens, token)
		}
	}

	return finalTokens
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
		if useDigitAbbrs && strutil.IsNumberOnly(eos) {
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
	matches := findMatchingKnownWord(ld, str)
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

func findMatchingKnownWord(ld *data.LocaleData, str string) []string {
	if strings.TrimSpace(str) == "" {
		return nil
	}

	// Try each known word
	lastPos := -1
	var candidates []string
	for _, knownWord := range ld.KnownWords {
		wordLen := len(knownWord)
		pos := strings.Index(str, knownWord)
		if pos < 0 {
			continue
		}

		prefix := str[:pos]
		match := str[pos : pos+wordLen]
		suffix := str[pos+wordLen:]

		prefixMatch := true
		suffixMatch := true
		if !ld.NoWordSpacing {
			prefixMatch = checkKnownWordNeighbor(prefix, true)
			suffixMatch = checkKnownWordNeighbor(suffix, false)
		}

		if prefixMatch && suffixMatch {
			if lastPos == -1 || pos < lastPos {
				lastPos = pos
				candidates = []string{str, prefix, match, suffix}
			}

			if pos == 0 {
				return candidates
			}
		}
	}

	return candidates
}

func checkKnownWordNeighbor(s string, isPrefix bool) bool {
	if s == "" {
		return true
	}

	var r rune
	if isPrefix {
		r, _ = utf8.DecodeLastRuneInString(s)
	} else {
		r, _ = utf8.DecodeRuneInString(s)
	}

	return r == '_' || unicode.IsNumber(r) || !unicode.Is(disallowedKnownWordRunes, r)
}

func splitByNumerals(str string, keepFormatting bool) []string {
	// Splid string to tokens
	var tokens []string
	var currentGroup []rune
	var currentIsDigit bool

	for i, r := range str {
		rIsDigit := unicode.IsDigit(r)
		if i > 0 && currentIsDigit != rIsDigit {
			tokens = append(tokens, string(currentGroup))
			currentGroup = []rune{}
		}

		currentIsDigit = rIsDigit
		currentGroup = append(currentGroup, r)
	}

	tokens = append(tokens, string(currentGroup))

	// Filter tokens
	var validTokens []string
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

func simpleSplit(ld *data.LocaleData, str string, keepFormatting bool, skippedTokens strutil.Dict) []string {
	// Split by numerals
	var lastPos int
	var tokens []string
	for _, pos := range rxNumeral.FindAllStringSubmatchIndex(str, -1) {
		if beforeNumber := str[lastPos:pos[0]]; beforeNumber != "" {
			tokens = append(tokens, beforeNumber)
		}

		tokens = append(tokens, str[pos[0]:pos[1]])
		lastPos = pos[1]
	}

	if leftover := str[lastPos:]; leftover != "" {
		tokens = append(tokens, leftover)
	}

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
