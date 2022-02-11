package language

import (
	"strings"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/markusmobius/go-dateparser/internal/setting"
	"github.com/markusmobius/go-dateparser/internal/strutil"
)

func mapSkippedTokens(cfg *setting.Configuration, ld *data.LocaleData) strutil.Dict {
	skippedTokens := strutil.NewDict()
	if cfg != nil {
		skippedTokens.Add(cfg.SkipTokens...)
	}

	if ld != nil && ld.Name == "fi" {
		skippedTokens.Remove("t")
	}

	return skippedTokens
}

func join(tokens []string, separator string) string {
	if len(tokens) == 0 {
		return ""
	}

	joined := tokens[0]
	for i := 1; i < len(tokens); i++ {
		left, right := tokens[i-1], tokens[i]
		leftAlwaysKept := alwaysKeptTokens.Contain(left)
		rightAlwaysKept := alwaysKeptTokens.Contain(right)
		if !leftAlwaysKept && !rightAlwaysKept {
			joined += separator
		}

		joined += right
	}

	joined = strings.TrimSpace(joined)
	return joined
}

func translateWord(ld *data.LocaleData, word string) (string, bool) {
	if translation, exist := ld.RelativeType[word]; exist {
		return translation, true
	}

	if translation, exist := ld.Translations[word]; exist {
		return translation, true
	}

	return "", false
}

func isInDictionary(ld *data.LocaleData, word string) bool {
	_, exist := translateWord(ld, word)
	return exist
}
