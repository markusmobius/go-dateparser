package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"golang.org/x/text/unicode/norm"
	"gopkg.in/yaml.v2"
)

func parseJsonFile(dst interface{}, fPath string) error {
	f, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(dst)
}

func parseYamlFile(dst interface{}, fPath string) error {
	f, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return yaml.NewDecoder(f).Decode(dst)
}

func renderJSON(v interface{}, dstPath string) error {
	bt, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(dstPath, bt, os.ModePerm)
}

func normalizeString(str string) string {
	// Normalize unicode
	str = norm.NFKC.String(str)

	// Replace Python capture group
	str = rxPythonCaptureGroup.ReplaceAllString(str, "$$$1")
	return str
}

func cleanString(str string) string {
	str = normalizeString(str)
	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}

func cleanList(useStringCleaner bool, fnFilter func(string) bool, items ...string) []string {
	// Prepare tracker
	tracker := map[string]struct{}{}

	// Create final list
	var finalList []string
	for _, item := range items {
		if useStringCleaner {
			item = cleanString(item)
		} else {
			item = normalizeString(item)
		}

		// Make sure item not empty
		if item == "" {
			continue
		}

		// Make sure item tracked yet
		if _, exist := tracker[item]; exist {
			continue
		}

		// Make sure item passed the filter
		if fnFilter != nil && !fnFilter(item) {
			continue
		}

		// Save the item
		tracker[item] = struct{}{}
		finalList = append(finalList, item)
	}

	// Sort final list
	sort.Strings(finalList)

	if len(finalList) == 0 {
		return nil
	} else {
		return finalList
	}
}

func cleanEmptyStringListMap(m map[string][]string) map[string][]string {
	for key, list := range m {
		if len(list) == 0 {
			delete(m, key)
		}
	}

	if len(m) == 0 {
		return nil
	}

	return m
}

func mergeList(cleanString bool, a, b []string) []string {
	return cleanList(cleanString, nil, append(a, b...)...)
}

func mergeStringMap(base, input map[string]string) map[string]string {
	mergedMap := map[string]string{}

	for key, value := range input {
		key = normalizeString(key)
		mergedMap[key] = normalizeString(value)
	}

	for key, value := range base {
		key = normalizeString(key)
		if _, exist := mergedMap[key]; !exist {
			mergedMap[key] = normalizeString(value)
		}
	}

	return mergedMap
}

func mergeStringListMap(base, input map[string][]string) map[string][]string {
	mergedMap := map[string][]string{}

	for key, items := range input {
		key = normalizeString(key)
		mergedMap[key] = cleanList(false, nil, items...)
	}

	for key, items := range base {
		key = normalizeString(key)

		if existingValue, exist := mergedMap[key]; !exist {
			mergedMap[key] = cleanList(false, nil, items...)
		} else {
			mergedMap[key] = mergeList(false, existingValue, items)
		}
	}

	return mergedMap
}

func mergeLocaleData(base, input LocaleData) LocaleData {
	name := input.Name
	if name == "" {
		name = base.Name
	}

	dateOrder := input.DateOrder
	if dateOrder == "" {
		dateOrder = base.DateOrder
	}

	sentenceSplitterGroup := input.SentenceSplitterGroup
	if sentenceSplitterGroup == 0 {
		sentenceSplitterGroup = base.SentenceSplitterGroup
	}

	return LocaleData{
		Name:                  name,
		DateOrder:             dateOrder,
		SkipWords:             mergeList(true, base.SkipWords, input.SkipWords),
		PertainWords:          mergeList(true, base.PertainWords, input.PertainWords),
		NoWordSpacing:         base.NoWordSpacing || input.NoWordSpacing,
		SentenceSplitterGroup: sentenceSplitterGroup,
		January:               mergeList(true, base.January, input.January),
		February:              mergeList(true, base.February, input.February),
		March:                 mergeList(true, base.March, input.March),
		April:                 mergeList(true, base.April, input.April),
		May:                   mergeList(true, base.May, input.May),
		June:                  mergeList(true, base.June, input.June),
		July:                  mergeList(true, base.July, input.July),
		August:                mergeList(true, base.August, input.August),
		September:             mergeList(true, base.September, input.September),
		October:               mergeList(true, base.October, input.October),
		November:              mergeList(true, base.November, input.November),
		December:              mergeList(true, base.December, input.December),
		Monday:                mergeList(true, base.Monday, input.Monday),
		Tuesday:               mergeList(true, base.Tuesday, input.Tuesday),
		Wednesday:             mergeList(true, base.Wednesday, input.Wednesday),
		Thursday:              mergeList(true, base.Thursday, input.Thursday),
		Friday:                mergeList(true, base.Friday, input.Friday),
		Saturday:              mergeList(true, base.Saturday, input.Saturday),
		Sunday:                mergeList(true, base.Sunday, input.Sunday),
		AM:                    mergeList(true, base.AM, input.AM),
		PM:                    mergeList(true, base.PM, input.PM),
		Decade:                mergeList(true, base.Decade, input.Decade),
		Year:                  mergeList(true, base.Year, input.Year),
		Month:                 mergeList(true, base.Month, input.Month),
		Week:                  mergeList(true, base.Week, input.Week),
		Day:                   mergeList(true, base.Day, input.Day),
		Hour:                  mergeList(true, base.Hour, input.Hour),
		Minute:                mergeList(true, base.Minute, input.Minute),
		Second:                mergeList(true, base.Second, input.Second),
		In:                    mergeList(true, base.In, input.In),
		Ago:                   mergeList(true, base.Ago, input.Ago),

		RelativeType:      mergeStringListMap(base.RelativeType, input.RelativeType),
		RelativeTypeRegex: mergeStringListMap(base.RelativeTypeRegex, input.RelativeTypeRegex),
		Simplifications:   mergeStringMap(base.Simplifications, input.Simplifications),
	}
}
