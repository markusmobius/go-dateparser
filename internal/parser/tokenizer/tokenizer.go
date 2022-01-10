package tokenizer

import (
	"unicode"
)

type TokenType uint8

type Token struct {
	Text string
	Type TokenType
}

const (
	Digit TokenType = iota
	Letter
	Other
)

func Tokenize(str string) []Token {
	chars := []rune(str)
	switch len(chars) {
	case 0:
		return nil
	case 1:
		tokenType, _ := compareChars(chars[0], 0)
		return []Token{{str, tokenType}}
	}

	var tokens []Token
	text := []rune{chars[0]}
	for i := 1; i < len(chars); i++ {
		prev, current := chars[i-1], chars[i]
		tokenType, changed := compareChars(prev, current)
		if changed {
			tokens = append(tokens, Token{
				Text: string(text),
				Type: tokenType,
			})
			text = []rune{current}
		} else {
			text = append(text, current)
		}
	}

	if len(text) > 0 {
		tokenType, _ := compareChars(text[0], 0)
		tokens = append(tokens, Token{
			Text: string(text),
			Type: tokenType,
		})
	}

	return tokens
}

func compareChars(r1, r2 rune) (TokenType, bool) {
	r1IsDigit := unicode.IsDigit(r1) || r1 == ':'
	r2IsDigit := unicode.IsDigit(r2) || r2 == ':'
	r1IsLetter := unicode.IsLetter(r1)
	r2IsLetter := unicode.IsLetter(r2)

	if r1IsDigit {
		return Digit, !r2IsDigit
	}

	if r1IsLetter {
		return Letter, !r2IsLetter
	}

	return Other, r2IsDigit || r2IsLetter
}
