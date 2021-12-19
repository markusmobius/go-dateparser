package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"

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

	return ioutil.WriteFile(dstPath, bt, os.ModePerm)
}

func normalizeUnicode(str string) string {
	normalized, _, err := transform.String(unicodeTransformer, str)
	if err != nil {
		return str
	}

	return normalized
}

func cleanString(str string) string {
	str = normalizeUnicode(str)
	str = rxSanitizeAposthrope.ReplaceAllString(str, "'")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
	return str
}

func normalizeString(str string) string {
	str = normalizeUnicode(str)
	str = strings.ToLower(str)
	str = strings.Join(strings.Fields(str), " ")
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
	tracker := map[string]struct{}{}

	for _, item := range items {
		if useStringCleaner {
			item = cleanString(item)
		} else {
			item = normalizeString(item)
		}

		_, itemExist := tracker[item]
		if item != "" && !itemExist {
			tracker[item] = struct{}{}
			cleanedList = append(cleanedList, item)
		}
	}

	sort.Strings(cleanedList)
	return cleanedList
}

func cleanSimplifications(items ...SimplificationData) []SimplificationData {
	var cleanedSimplifications []SimplificationData
	tracker := map[string]struct{}{}

	for _, data := range items {
		key := data.Pattern + "==" + data.Replacement
		_, itemExist := tracker[key]

		if data.Pattern != "" && !itemExist {
			tracker[key] = struct{}{}
			cleanedSimplifications = append(cleanedSimplifications, data)
		}
	}

	return cleanedSimplifications
}
