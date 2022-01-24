package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	// Helper functions
	sp := func(items ...string) []string { return items }
	tt := func(items ...TokenType) []TokenType { return items }

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

	tests := []testScenario{{
		"11 april 2010",
		sp("11", " ", "april", " ", "2010"),
		tt(D, O, L, O, D),
	}, {
		"Tuesday 11 april 2010",
		sp("Tuesday", " ", "11", " ", "april", " ", "2010"),
		tt(L, O, D, O, L, O, D),
	}, {
		"11/12-2013",
		sp("11", "/", "12", "-", "2013"),
		tt(D, O, D, O, D),
	}, {
		"11/12-2013",
		sp("11", "/", "12", "-", "2013"),
		tt(D, O, D, O, D),
	}, {
		"10:30:35 PM",
		sp("10:30:35", " ", "PM"),
		tt(D, O, L),
	}, {
		"18:50",
		sp("18:50"),
		tt(D),
	}, {
		"December 23, 2010, 16:50 pm",
		sp("December", " ", "23", ", ", "2010", ", ", "16:50", " ", "pm"),
		tt(L, O, D, O, D, O, D, O, L),
	}, {
		"Oct 1 2018 4:40 PM EST —",
		sp("Oct", " ", "1", " ", "2018", " ", "4:40", " ", "PM", " ", "EST", " —"),
		tt(L, O, D, O, D, O, D, O, L, O, L, O),
	}, {
		"0123456789",
		sp("0123456789"),
		tt(D),
	}, {
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		sp("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		tt(L),
	}, {
		"./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊", // unrecognized characters
		sp("./\\()\"',.;<>~!@#$%^&*|+=[]{}`~?-—–     😊"),
		tt(O),
	}}

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
