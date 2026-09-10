package dateparser

import (
	"slices"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/regexp"
)

var (
	rxNgramToken        = regexp.MustCompile(`[^\s\p{Z}\x{0085},|()@]+`)
	rxNgramBadCandidate = regexp.MustCompile(`^(\p{Nd}{1,3}|#\p{Nd}+|[-/.]+|[\p{L}\p{N}_]\.?|an)$`)
)

func (p *Parser) searchNgrams(cfg *Configuration, locales []string, lang, text string) []SearchResult {
	parseConfig := cfg.Clone()
	parseConfig.Locales = slices.Clone(locales)
	parseConfig.Languages = nil
	parseConfig.UseGivenOrder = true
	parseConfig.TryPreviousLocales = false

	var tokens [][]int
	for _, position := range rxNgramToken.FindAllStringIndex(text, -1) {
		switch text[position[0]:position[1]] {
		case "on", "at", "of", "a":
			continue
		}
		tokens = append(tokens, position)
	}

	var results []SearchResult
	for index := 0; index < len(tokens); {
		matched := false
		for size := min(7, len(tokens)-index); size > 0; size-- {
			positions := tokens[index : index+size]
			words := make([]string, size)
			for offset, position := range positions {
				words[offset] = text[position[0]:position[1]]
			}
			candidate := strings.Join(words, " ")
			if rxNgramBadCandidate.MatchString(candidate) {
				continue
			}
			parsed, err := p.Parse(parseConfig, candidate)
			if err != nil || parsed.IsZero() {
				continue
			}
			parsed.Locale = lang
			substring := text[positions[0][0]:positions[len(positions)-1][1]]
			results = append(results, SearchResult{Date: parsed, Text: strings.Trim(substring, " .,:()[]-'")})
			index += size
			matched = true
			break
		}
		if !matched {
			index++
		}
	}
	if cfg.ReturnTimeSpan {
		results = append(results, searchTimeSpan(cfg, lang, text)...)
	}
	return results
}
