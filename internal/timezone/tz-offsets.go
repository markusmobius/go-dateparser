package timezone

import (
	"fmt"
	"regexp"
	"time"
)

type TimezoneOffsetData struct {
	Name   string
	Regex  *regexp.Regexp
	Offset time.Duration
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
				searchPattern := tzName
				regexPattern := fmt.Sprintf(pattern, searchPattern)
				if _, exist := regexPatterns[regexPattern]; exist {
					continue
				}

				searchPatterns[searchPattern] = struct{}{}
				offsets = append(offsets, TimezoneOffsetData{
					Name:   tzName,
					Regex:  regexp.MustCompile(`(?i)` + regexPattern),
					Offset: time.Duration(offset) * time.Second,
				})
			}

			// Alternative patterns
			for replacer, replacement := range tzInfo.AlternativePatterns {
				for tzName, offset := range tzInfo.Timezones {
					searchPattern := replacer.ReplaceAllString(tzName, replacement)
					regexPattern := fmt.Sprintf(pattern, searchPattern)
					if _, exist := regexPatterns[regexPattern]; exist {
						continue
					}

					searchPatterns[searchPattern] = struct{}{}
					offsets = append(offsets, TimezoneOffsetData{
						Name:   tzName,
						Regex:  regexp.MustCompile(`(?i)` + regexPattern),
						Offset: time.Duration(offset) * time.Second,
					})
				}
			}
		}
	}

	// Generate regex for searching
	strSearchPatterns := ""
	for pattern := range searchPatterns {
		strSearchPatterns += pattern + "|"
	}
	strSearchPatterns = strSearchPatterns[:len(strSearchPatterns)-1]

	rxSearch := regexp.MustCompile(strSearchPatterns)
	rxSearchIgnoreCase := regexp.MustCompile("(?i)" + strSearchPatterns)
	return offsets, rxSearch, rxSearchIgnoreCase
}()
