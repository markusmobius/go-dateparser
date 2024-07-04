package main

import (
	"encoding/json"
	"os"
	"sort"
	"strings"

	"github.com/elliotchance/pie/v2"
	"github.com/zyedidia/generic/mapset"
	"golang.org/x/text/transform"
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

	return os.WriteFile(dstPath, bt, os.ModePerm)
}

func normalizeUnicode(str string) string {
	normalized, _, err := transform.String(unicodeTransformer, str)
	if err != nil {
		return str
	}

	return normalized
}

func sanitizeSpaces(s string) string {
	s = rxNbsp.ReplaceAllString(s, " ")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func cleanString(str string) string {
	str = normalizeUnicode(str)
	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = sanitizeSpaces(str)
	return str
}

func normalizeString(str string) string {
	str = normalizeUnicode(str)
	str = strings.ToLower(str)
	str = sanitizeSpaces(str)
	return str
}

func normalizeCharset(str string) string {
	normalized, _, err := transform.String(nfkcTransformer, str)
	if err == nil {
		str = normalized
	}

	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = sanitizeSpaces(str)
	return str
}

func filterList(fnFilter func(string) bool, items ...string) []string {
	var filteredList []string
	for _, item := range items {
		if item != "" && (fnFilter == nil || fnFilter(item)) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}

func cleanList(useStringCleaner bool, items ...string) []string {
	var cleanedList []string
	for _, item := range items {
		if useStringCleaner {
			item = cleanString(item)
		} else {
			item = normalizeString(item)
		}
		cleanedList = append(cleanedList, item)
	}

	cleanedList = pie.Unique(cleanedList)
	sort.Strings(cleanedList)
	return cleanedList
}

func cleanSimplifications(items ...SimplificationData) []SimplificationData {
	var cleanedSimplifications []SimplificationData
	tracker := mapset.New[string]()

	for _, data := range items {
		key := data.Pattern + "==" + data.Replacement
		if data.Pattern != "" && !tracker.Has(key) {
			tracker.Put(key)
			cleanedSimplifications = append(cleanedSimplifications, data)
		}
	}

	return cleanedSimplifications
}

func cloneSlice[T any](source []T) []T {
	return append([]T{}, source...)
}

func cloneMap[K comparable, V any](source map[K]V) map[K]V {
	clone := make(map[K]V)
	for k, v := range source {
		clone[k] = v
	}
	return clone
}

func cloneMapSlice[K comparable, V any](source map[K][]V) map[K][]V {
	clone := make(map[K][]V)
	for k, v := range source {
		clone[k] = cloneSlice(v)
	}
	return clone
}

// reduceSlice removes item in current that already exist in input
func reduceSlice[T comparable](current, input []T) []T {
	tracker := mapset.New[T]()
	for _, i := range input {
		tracker.Put(i)
	}

	var result []T
	for _, c := range current {
		if !tracker.Has(c) {
			result = append(result, c)
		}
	}

	return result
}

// reduceMap removes item in current that already exist in input
func reduceMap[K comparable, V comparable](current, input map[K]V) map[K]V {
	result := make(map[K]V)
	for ck, cv := range current {
		if _, exist := input[ck]; !exist {
			result[ck] = cv
		}
	}
	return result
}

// reduceMapSlice removes item from slice in current that already exist in input
func reduceMapSlice[K comparable, V comparable](current, input map[K][]V) map[K][]V {
	result := make(map[K][]V)
	for ck, cv := range current {
		// If not exist in input, we can use as it is
		if iv, exist := input[ck]; !exist {
			result[ck] = append([]V{}, cv...)
		} else {
			// If already exist, we need to reduce the slice
			reduced := reduceSlice(cv, iv)
			if len(reduced) > 0 {
				result[ck] = reduced
			}
		}
	}
	return result
}

func reduceSimplifications(current, input []SimplificationData) []SimplificationData {
	tracker := mapset.New[string]()
	for _, data := range input {
		key := data.Pattern + "==" + data.Replacement
		tracker.Put(key)
	}

	var reducedSimplifications []SimplificationData
	for _, data := range current {
		key := data.Pattern + "==" + data.Replacement
		if !tracker.Has(key) {
			reducedSimplifications = append(reducedSimplifications, data)
		}
	}

	return reducedSimplifications
}
