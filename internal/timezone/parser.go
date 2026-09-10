package timezone

import (
	"strings"
	"unicode/utf8"
)

var timezoneWords = func() map[string]struct{} {
	words := map[string]struct{}{}
	for _, info := range timezoneInfoList {
		for name := range info.Timezones {
			if !strings.ContainsAny(name, "+-") {
				words[strings.ToUpper(name)] = struct{}{}
			}
		}
	}
	return words
}()

func IsTimezoneToken(token string) bool {
	return rxTimezoneToken.MatchString(strings.TrimSpace(token))
}

// PopTzOffset extracts timezone data from str then return the
// str with the timezone code removed.
func PopTzOffset(str string) (string, OffsetData) {
	if couldContainTimezone(str) && rxSearchIgnoreCase.MatchString(str) {
		for _, offsetData := range timezoneOffsets {
			idxs := offsetData.Regex.FindStringIndex(str)
			if len(idxs) == 0 {
				continue
			}

			str = str[:idxs[0]+1] + str[idxs[1]:]
			return str, offsetData
		}
	}

	return str, OffsetData{}
}

func couldContainTimezone(input string) bool {
	var word [16]byte
	for index := 0; index < len(input); index++ {
		character := input[index]
		if character >= utf8.RuneSelf || character == '+' || character == '-' {
			return true
		}
		if !isASCIITimezoneLetter(character) || (index > 0 && isASCIITimezoneLetter(input[index-1])) {
			continue
		}
		end := index
		for end < len(input) {
			character = input[end]
			if !isASCIITimezoneLetter(character) && (character < '0' || character > '9') && character != '_' {
				break
			}
			if end-index == len(word) {
				return true
			}
			if character >= 'a' && character <= 'z' {
				character -= 'a' - 'A'
			}
			word[end-index] = character
			end++
		}
		if _, exists := timezoneWords[string(word[:end-index])]; exists {
			return true
		}
	}
	return false
}

func isASCIITimezoneLetter(character byte) bool {
	character |= 'a' - 'A'
	return character >= 'a' && character <= 'z'
}

// WordIsTz check if the specified word is a timezone code.
func WordIsTz(word string) bool {
	return rxSearch.MatchString(word)
}
