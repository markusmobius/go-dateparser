package main

import (
	"strings"
	"testing"
)

func TestExactMatcherRule(t *testing.T) {
	for _, test := range []struct {
		Pattern string
		Parts   []string
	}{
		{`^(k)$`, []string{`[\x4B\x6B\u212A]`}},
		{`^(s)$`, []string{`[\x53\x73\u017F]`}},
		{`^(\x{03c3})$`, []string{`[\u03A3\u03C2\u03C3]`}},
		{`^(\d+[.,]?\d*)$`, []string{`[\x30-\x39]`, "+", "?", "*"}},
		{`^(?:ab|cd){2,3}$`, []string{" | ", "{2,3}"}},
		{`^(?:ab){2,}$`, []string{"{2,}"}},
		{`^(\x00|\x{1f600})$`, []string{`\x00`, `\U0001F600`}},
		{`^()$`, []string{`""`}},
		{`^(?i:a)(?-i:b)$`, []string{`[\x41\x61]`, `[\x62]`}},
	} {
		t.Run(test.Pattern, func(t *testing.T) {
			expression, err := exactMatcherRule(test.Pattern)
			if err != nil {
				t.Fatal(err)
			}
			for _, part := range test.Parts {
				if !strings.Contains(expression, part) {
					t.Errorf("expression %q does not contain %q", expression, part)
				}
			}
		})
	}
}

func TestExactMatcherRuleRejectsUnsupported(t *testing.T) {
	for _, pattern := range []string{`[`, `abc`, `^abc`, `abc$`, `(?m)^abc$`, `^(a\bb)$`, `^(.)$`} {
		if _, err := exactMatcherRule(pattern); err == nil {
			t.Errorf("accepted unsupported pattern %q", pattern)
		}
	}
}

func TestExactMatcherName(t *testing.T) {
	if exactMatcherName("") != "nil" {
		t.Fatal("empty patterns must inherit the parent's matcher")
	}
	first := exactMatcherName("^(first)$")
	if first != exactMatcherName("^(first)$") || first == exactMatcherName("^(second)$") {
		t.Fatal("matcher names must be stable and distinguish patterns")
	}
	if !strings.HasPrefix(first, "matchExact") {
		t.Fatal("matcher name is not a Go identifier")
	}
}

func TestExactMatcherFieldPreserved(t *testing.T) {
	input := "\tExactCombinedMatcher: matchExact0,\n\tCount: 0,\n\tRxCombined: nil,\n\tNoWordSpacing: false,\n"
	expected := "\tExactCombinedMatcher: matchExact0,\n"
	if output := rxGoZeroField.ReplaceAllString(input, ""); output != expected {
		t.Fatalf("nonzero matcher field was removed: %q", output)
	}
}
