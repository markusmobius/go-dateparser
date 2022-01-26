package tokenizer

import (
	"unicode"

	"golang.org/x/text/unicode/rangetable"
)

var (
	digits  = rangetable.New([]rune("1234567890:")...)
	letters = rangetable.New([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
)

type TokenType uint8

const (
	Null TokenType = iota
	Digit
	Letter
	Other
)

type Token struct {
	Text string
	Type TokenType
}

func (t Token) IsEmpty() bool {
	return t.Text == "" && t.Type == 0
}

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
	r1IsDigit := unicode.Is(digits, r1) || r1 == ':'
	r2IsDigit := unicode.Is(digits, r2) || r2 == ':'
	r1IsLetter := unicode.Is(letters, r1)
	r2IsLetter := unicode.Is(letters, r2)

	if r1IsDigit {
		return Digit, !r2IsDigit
	}

	if r1IsLetter {
		return Letter, !r2IsLetter
	}

	return Other, r2IsDigit || r2IsLetter
}
