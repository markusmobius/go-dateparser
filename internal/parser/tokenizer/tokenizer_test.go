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
	type testScenario struct {
		String        string
		ExpectedTexts []string
		ExpectedTypes []TokenType
	}

	tests := []testScenario{
		{
			"11 april 2010",
			[]string{"11", " ", "april", " ", "2010"},
			[]TokenType{0, 2, 1, 2, 0},
		}, {
			"Tuesday 11 april 2010",
			[]string{"Tuesday", " ", "11", " ", "april", " ", "2010"},
			[]TokenType{1, 2, 0, 2, 1, 2, 0},
		}, {
			"11/12-2013",
			[]string{"11", "/", "12", "-", "2013"},
			[]TokenType{0, 2, 0, 2, 0},
		}, {
			"11/12-2013",
			[]string{"11", "/", "12", "-", "2013"},
			[]TokenType{0, 2, 0, 2, 0},
		}, {
			"10:30:35 PM",
			[]string{"10:30:35", " ", "PM"},
			[]TokenType{0, 2, 1},
		}, {
			"18:50",
			[]string{"18:50"},
			[]TokenType{0},
		}, {
			"December 23, 2010, 16:50 pm",
			[]string{"December", " ", "23", ", ", "2010", ", ", "16:50", " ", "pm"},
			[]TokenType{1, 2, 0, 2, 0, 2, 0, 2, 1},
		}, {
			"Oct 1 2018 4:40 PM EST —",
			[]string{"Oct", " ", "1", " ", "2018", " ", "4:40", " ", "PM", " ", "EST", " —"},
			[]TokenType{1, 2, 0, 2, 0, 2, 0, 2, 1, 2, 1, 2},
		}, {
			"0123456789",
			[]string{"0123456789"},
			[]TokenType{0},
		}, {
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			[]string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			[]TokenType{1},
		}, {
			"./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊", // unrecognized characters
			[]string{"./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊"},
			[]TokenType{2},
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
