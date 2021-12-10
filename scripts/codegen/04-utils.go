package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"

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

func cleanString(str string) string {
	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}

func cleanList(initialItems []string, fnFilter func(string) bool, items ...string) []string {
	// Prepare tracker
	tracker := map[string]struct{}{}
	if initialItems != nil && len(initialItems) > 0 {
		for _, item := range initialItems {
			tracker[item] = struct{}{}
		}
	}

	// Create final list
	var finalList []string
	for _, item := range items {
		item = cleanString(item)

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

func mergeList(a, b []string) []string {
	return cleanList(nil, nil, append(a, b...)...)
}

func mergeStringMap(base, input map[string]string) map[string]string {
	mergedMap := map[string]string{}

	for key, value := range input {
		mergedMap[key] = value
	}

	for key, value := range base {
		if _, exist := mergedMap[key]; !exist {
			mergedMap[key] = value
		}
	}

	return mergedMap
}

func mergeStringListMap(base, input map[string][]string) map[string][]string {
	mergedMap := map[string][]string{}

	for key, value := range input {
		mergedMap[key] = value
	}

	for key, value := range base {
		if existingValue, exist := mergedMap[key]; !exist {
			mergedMap[key] = value
		} else {
			mergedMap[key] = mergeList(existingValue, value)
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
		SkipWords:             mergeList(base.SkipWords, input.SkipWords),
		PertainWords:          mergeList(base.PertainWords, input.PertainWords),
		NoWordSpacing:         base.NoWordSpacing || input.NoWordSpacing,
		SentenceSplitterGroup: sentenceSplitterGroup,
		January:               mergeList(base.January, input.January),
		February:              mergeList(base.February, input.February),
		March:                 mergeList(base.March, input.March),
		April:                 mergeList(base.April, input.April),
		May:                   mergeList(base.May, input.May),
		June:                  mergeList(base.June, input.June),
		July:                  mergeList(base.July, input.July),
		August:                mergeList(base.August, input.August),
		September:             mergeList(base.September, input.September),
		October:               mergeList(base.October, input.October),
		November:              mergeList(base.November, input.November),
		December:              mergeList(base.December, input.December),
		Monday:                mergeList(base.Monday, input.Monday),
		Tuesday:               mergeList(base.Tuesday, input.Tuesday),
		Wednesday:             mergeList(base.Wednesday, input.Wednesday),
		Thursday:              mergeList(base.Thursday, input.Thursday),
		Friday:                mergeList(base.Friday, input.Friday),
		Saturday:              mergeList(base.Saturday, input.Saturday),
		Sunday:                mergeList(base.Sunday, input.Sunday),
		AM:                    mergeList(base.AM, input.AM),
		PM:                    mergeList(base.PM, input.PM),
		Decade:                mergeList(base.Decade, input.Decade),
		Year:                  mergeList(base.Year, input.Year),
		Month:                 mergeList(base.Month, input.Month),
		Week:                  mergeList(base.Week, input.Week),
		Day:                   mergeList(base.Day, input.Day),
		Hour:                  mergeList(base.Hour, input.Hour),
		Minute:                mergeList(base.Minute, input.Minute),
		Second:                mergeList(base.Second, input.Second),
		In:                    mergeList(base.In, input.In),
		Ago:                   mergeList(base.Ago, input.Ago),

		RelativeType:      mergeStringListMap(base.RelativeType, input.RelativeType),
		RelativeTypeRegex: mergeStringListMap(base.RelativeTypeRegex, input.RelativeTypeRegex),
		Simplifications:   mergeStringMap(base.Simplifications, input.Simplifications),
	}
}
