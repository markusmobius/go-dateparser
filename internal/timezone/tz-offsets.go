package timezone

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type TimezoneOffsetData struct {
	Name   string
	Regex  *regexp.Regexp
	Offset int
}

func (tod TimezoneOffsetData) IsZero() bool {
	return tod.Name == "" && tod.Regex == nil && tod.Offset == 0
}

var timezoneOffsets, rxSearch, rxSearchIgnoreCase = func() ([]TimezoneOffsetData, *regexp.Regexp, *regexp.Regexp) {
	// Process timezone info list
	offsets := []TimezoneOffsetData{}
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
				offsets = append(offsets, TimezoneOffsetData{
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
					offsets = append(offsets, TimezoneOffsetData{
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
