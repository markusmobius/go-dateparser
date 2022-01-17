package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	// Helper function
	fnTexts := func(tokens []Token) []string {
		var texts []string
		for i := range tokens {
			texts = append(texts, tokens[i].Text)
		}
		return texts
	}

	fnTypes := func(tokens []Token) []TokenType {
		var types []TokenType
		for i := range tokens {
			types = append(types, tokens[i].Type)
		}
		return types
	}

	// Prepare scenarios
	D, L, O := Digit, Letter, Other
	type testScenario struct {
		String        string
		ExpectedTexts []string
		ExpectedTypes []TokenType
	}

	tests := []testScenario{
		{
			"11 april 2010",
			[]string{"11", " ", "april", " ", "2010"},
			[]TokenType{D, O, L, O, D},
		}, {
			"Tuesday 11 april 2010",
			[]string{"Tuesday", " ", "11", " ", "april", " ", "2010"},
			[]TokenType{L, O, D, O, L, O, D},
		}, {
			"11/12-2013",
			[]string{"11", "/", "12", "-", "2013"},
			[]TokenType{D, O, D, O, D},
		}, {
			"11/12-2013",
			[]string{"11", "/", "12", "-", "2013"},
			[]TokenType{D, O, D, O, D},
		}, {
			"10:30:35 PM",
			[]string{"10:30:35", " ", "PM"},
			[]TokenType{D, O, L},
		}, {
			"18:50",
			[]string{"18:50"},
			[]TokenType{D},
		}, {
			"December 23, 2010, 16:50 pm",
			[]string{"December", " ", "23", ", ", "2010", ", ", "16:50", " ", "pm"},
			[]TokenType{L, O, D, O, D, O, D, O, L},
		}, {
			"Oct 1 2018 4:40 PM EST —",
			[]string{"Oct", " ", "1", " ", "2018", " ", "4:40", " ", "PM", " ", "EST", " —"},
			[]TokenType{L, O, D, O, D, O, D, O, L, O, L, O},
		}, {
			"0123456789",
			[]string{"0123456789"},
			[]TokenType{D},
		}, {
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			[]string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			[]TokenType{L},
		}, {
			"./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊", // unrecognized characters
			[]string{"./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊"},
			[]TokenType{O},
		},
	}

	// Start tests
	nFailed := 0
	for _, test := range tests {
		tokens := Tokenize(test.String)
		passed := assert.Equal(t, test.ExpectedTexts, fnTexts(tokens), test.String)
		if passed {
			passed = assert.Equal(t, test.ExpectedTypes, fnTypes(tokens), test.String)
		}

		if !passed {
			nFailed++
		}
	}

	if nFailed > 0 {
		fmt.Printf("Failed %d from %d tests\n", nFailed, len(tests))
	}
}
