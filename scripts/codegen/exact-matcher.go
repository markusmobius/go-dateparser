package main

import (
	"crypto/sha256"
	"fmt"
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"regexp/syntax"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func exactMatcherName(pattern string) string {
	if pattern == "" {
		return "nil"
	}
	digest := sha256.Sum256([]byte(pattern))
	return fmt.Sprintf("matchExact%x", digest)
}

func generateExactMatchers(locales map[string]LocaleData, executable string) ([]byte, error) {
	patterns := map[string]struct{}{}
	for _, locale := range locales {
		if pattern := locale.ExactCombinedRegexPattern; pattern != "" {
			patterns[pattern] = struct{}{}
		}
	}
	var ordered []string
	for pattern := range patterns {
		ordered = append(ordered, pattern)
	}
	slices.Sort(ordered)

	var rules strings.Builder
	rules.WriteString(exactMatcherTemplate)
	for _, pattern := range ordered {
		expression, err := exactMatcherRule(pattern)
		if err != nil {
			return nil, fmt.Errorf("exact matcher %q: %w", pattern, err)
		}
		fmt.Fprintf(&rules, `
func %s(input string) bool {
	cursor, marker, limit := 0, 0, len(input)
	_ = marker
	/*!use:re2c:exact
		%s { return cursor == limit }
		* { return false }
		$ { return false }
	*/
}
`, exactMatcherName(pattern), expression)
	}

	directory, err := os.MkdirTemp("", "dateparser-re2go-")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(directory)
	inputPath := filepath.Join(directory, "exact.re")
	outputPath := filepath.Join(directory, "exact.go")
	if err := os.WriteFile(inputPath, []byte(rules.String()), 0644); err != nil {
		return nil, err
	}
	command := exec.Command(executable, "-W", "-F", "--utf8", "--input-encoding", "utf8", "--no-generation-date", "-i", inputPath, "-o", outputPath)
	if output, err := command.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("re2go: %w: %s", err, strings.TrimSpace(string(output)))
	}
	generated, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}
	return format.Source(generated)
}

func exactMatcherRule(pattern string) (string, error) {
	parsed, err := syntax.Parse("(?i)"+pattern, syntax.Perl)
	if err != nil {
		return "", err
	}
	if parsed.Op != syntax.OpConcat || len(parsed.Sub) < 3 ||
		parsed.Sub[0].Op != syntax.OpBeginText || parsed.Sub[len(parsed.Sub)-1].Op != syntax.OpEndText {
		return "", fmt.Errorf("pattern must be anchored to the whole input")
	}
	return re2goExpression(&syntax.Regexp{Op: syntax.OpConcat, Sub: parsed.Sub[1 : len(parsed.Sub)-1]})
}

func re2goExpression(parsed *syntax.Regexp) (string, error) {
	switch parsed.Op {
	case syntax.OpEmptyMatch:
		return `""`, nil
	case syntax.OpLiteral:
		var characters []string
		for _, value := range parsed.Rune {
			values := []rune{value}
			if parsed.Flags&syntax.FoldCase != 0 {
				for folded := unicode.SimpleFold(value); folded != value; folded = unicode.SimpleFold(folded) {
					values = append(values, folded)
				}
			}
			slices.Sort(values)
			var character strings.Builder
			character.WriteByte('[')
			for _, member := range values {
				character.WriteString(re2goRune(member))
			}
			character.WriteByte(']')
			characters = append(characters, character.String())
		}
		return strings.Join(characters, " "), nil
	case syntax.OpCharClass:
		var character strings.Builder
		character.WriteByte('[')
		for index := 0; index < len(parsed.Rune); index += 2 {
			character.WriteString(re2goRune(parsed.Rune[index]))
			if parsed.Rune[index] != parsed.Rune[index+1] {
				character.WriteByte('-')
				character.WriteString(re2goRune(parsed.Rune[index+1]))
			}
		}
		character.WriteByte(']')
		return character.String(), nil
	case syntax.OpCapture:
		child, err := re2goExpression(parsed.Sub[0])
		return "(" + child + ")", err
	case syntax.OpConcat, syntax.OpAlternate:
		var children []string
		for _, child := range parsed.Sub {
			expression, err := re2goExpression(child)
			if err != nil {
				return "", err
			}
			children = append(children, expression)
		}
		separator := " "
		if parsed.Op == syntax.OpAlternate {
			separator = " | "
		}
		return "(" + strings.Join(children, separator) + ")", nil
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest, syntax.OpRepeat:
		child, err := re2goExpression(parsed.Sub[0])
		if err != nil {
			return "", err
		}
		operator := "*"
		switch parsed.Op {
		case syntax.OpPlus:
			operator = "+"
		case syntax.OpQuest:
			operator = "?"
		case syntax.OpRepeat:
			maximum := strconv.Itoa(parsed.Max)
			if parsed.Max < 0 {
				maximum = ""
			}
			operator = fmt.Sprintf("{%d,%s}", parsed.Min, maximum)
		}
		return "(" + child + ")" + operator, nil
	default:
		return "", fmt.Errorf("unsupported regex operation %s", parsed.Op)
	}
}

func re2goRune(value rune) string {
	if value <= 0xff {
		return fmt.Sprintf(`\x%02X`, value)
	}
	if value <= 0xffff {
		return fmt.Sprintf(`\u%04X`, value)
	}
	return fmt.Sprintf(`\U%08X`, value)
}

const exactMatcherTemplate = `package data

func exactPeek(input string, cursor int) byte {
	if cursor >= len(input) {
		return 0
	}
	return input[cursor]
}

/*!rules:re2c:exact
	re2c:eof = 0;
	re2c:yyfill:enable = 0;
	re2c:define:YYCTYPE = byte;
	re2c:define:YYPEEK = "exactPeek(input, cursor)";
	re2c:define:YYSKIP = "cursor++";
	re2c:define:YYBACKUP = "marker = cursor";
	re2c:define:YYRESTORE = "cursor = marker";
	re2c:define:YYLESSTHAN = "limit <= cursor";
*/
`
