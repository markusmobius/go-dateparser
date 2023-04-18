package timezone

import (
	"fmt"
	"sort"
	"strings"

	"github.com/markusmobius/go-dateparser/internal/regexp"
)

type OffsetData struct {
	Name   string
	Regex  *regexp.Regexp
	Offset int
}

func (tod OffsetData) IsZero() bool {
	return tod.Name == "" && tod.Regex == nil && tod.Offset == 0
}

var timezoneOffsets, rxSearch, rxSearchIgnoreCase = func() ([]OffsetData, *regexp.Regexp, *regexp.Regexp) {
	// Process timezone info list
	offsets := []OffsetData{}
	regexPatterns := map[string]struct{}{}
	searchPatterns := map[string]struct{}{}

	for _, tzInfo := range timezoneInfoList {
		for _, pattern := range tzInfo.RegexPatterns {
			// Main pattern
			for tzName, offset := range tzInfo.Timezones {
				searchPattern := regexp.QuoteMeta(tzName)
				regexPattern := fmt.Sprintf(pattern, searchPattern)
				if _, exist := regexPatterns[regexPattern]; exist {
					continue
				}

				searchPatterns[searchPattern] = struct{}{}
				offsets = append(offsets, OffsetData{
					Name:   tzName,
					Regex:  regexp.MustCompile(`(?i)` + regexPattern),
					Offset: offset,
				})
			}

			// Alternative patterns
			for replacer, replacement := range tzInfo.AlternativePatterns {
				for tzName, offset := range tzInfo.Timezones {
					searchPattern := regexp.QuoteMeta(tzName)
					searchPattern = replacer.ReplaceAllString(searchPattern, replacement)
					regexPattern := fmt.Sprintf(pattern, searchPattern)
					if _, exist := regexPatterns[regexPattern]; exist {
						continue
					}

					searchPatterns[searchPattern] = struct{}{}
					offsets = append(offsets, OffsetData{
						Name:   tzName,
						Regex:  regexp.MustCompile(`(?i)` + regexPattern),
						Offset: offset,
					})
				}
			}
		}
	}

	// Generate regex for searching
	var patterns []string
	for pattern := range searchPatterns {
		patterns = append(patterns, pattern)
	}

	sort.Strings(patterns)
	strSearchPatterns := strings.Join(patterns, "|")

	rxSearch := regexp.MustCompile(strSearchPatterns)
	rxSearchIgnoreCase := regexp.MustCompile("(?i)" + strSearchPatterns)
	return offsets, rxSearch, rxSearchIgnoreCase
}()
