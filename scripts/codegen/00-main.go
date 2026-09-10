package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/markusmobius/go-dateparser/internal/data"
	"github.com/spf13/cobra"
	"github.com/zyedidia/generic/mapset"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go run scripts/codegen/*.go",
		Short: "Generate code for locales data",
		RunE:  rootCmdHandler,
	}

	rootCmd.Flags().Bool("skip-raw", false, "skip downloading raw data")
	rootCmd.Flags().Bool("keep-language-order", false, "retain the existing language detection order")
	rootCmd.Flags().String("re2go", "re2go", "path to the re2go lexer generator")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func rootCmdHandler(cmd *cobra.Command, args []string) error {
	// Parse flags
	skipRawDownload, _ := cmd.Flags().GetBool("skip-raw")
	re2go, _ := cmd.Flags().GetString("re2go")
	re2go, err := exec.LookPath(re2go)
	if err != nil {
		return fmt.Errorf("locale generation requires re2go; add it to PATH or set --re2go: %w", err)
	}

	// Download raw data if required
	if !skipRawDownload {
		err := downloadRawData()
		if err != nil {
			return err
		}
	}

	// Generate map between language and its locales
	languageLocalesMap, err := createLanguageLocales()
	if err != nil {
		return err
	}

	// Generate order of languages based on its popularity
	languageOrder, err := createLanguageOrder(languageLocalesMap)
	if err != nil {
		return err
	}
	keepLanguageOrder, _ := cmd.Flags().GetBool("keep-language-order")
	if keepLanguageOrder {
		sort.SliceStable(languageOrder, func(first, second int) bool {
			firstOrder, firstKnown := data.LanguageOrder[languageOrder[first]]
			secondOrder, secondKnown := data.LanguageOrder[languageOrder[second]]
			if firstKnown != secondKnown {
				return firstKnown
			}
			return firstKnown && firstOrder < secondOrder
		})
	}

	// Generate map between a language and its descendant (if any)
	languageMap := createLanguageMap(languageOrder)

	// Parse CLDR data
	cldrLocaleData, err := parseAllCldrData(languageLocalesMap)
	if err != nil {
		return err
	}

	// Parse supplementary data
	supplementalLocaleData, err := parseAllSupplementaryData(languageLocalesMap)
	if err != nil {
		return err
	}

	// Generate map of locale order
	localeOrder := createLocaleOrder(languageLocalesMap, languageOrder)

	// Merge locale data
	finalLocaleData := map[string]LocaleData{}
	for language, locales := range languageLocalesMap {
		// Fetch suplemental data
		supplementalData, supplementalExist := supplementalLocaleData[language]

		// Process language data
		languageData, languageExist := cldrLocaleData[language]
		if !languageExist && !supplementalExist {
			continue
		}

		// Finalize and save language data
		languageData = languageData.Merge(supplementalData)
		languageData, err := finalizeData(languageData)
		if err != nil {
			return err
		} else {
			finalLocaleData[language] = languageData
		}

		// Process sub locales data
		for _, locale := range locales {
			// Create sub locale data
			localeData, localeExist := cldrLocaleData[locale]
			if !localeExist && !supplementalExist {
				continue
			}
			if !localeExist {
				localeData.Name = locale
			}
			if override, exists := supplementalLocaleData[locale]; exists {
				localeData = override.Merge(localeData)
			}

			localeData = localeData.Merge(supplementalData)
			localeData = localeData.Merge(languageData)

			// Finalize and save sub locale data
			localeData, err = finalizeData(localeData)
			if err != nil {
				return err
			}

			localeData = localeData.Reduce(languageData)
			localeData.Parent = language
			finalLocaleData[locale] = localeData
		}
	}

	exactMatchers, err := generateExactMatchers(finalLocaleData, re2go)
	if err != nil {
		return err
	}

	// Generate code
	os.RemoveAll(GO_CODE_DIR)
	os.MkdirAll(GO_CODE_DIR, os.ModePerm)
	if err := os.WriteFile(filepath.Join(GO_CODE_DIR, "05-exact-matchers.go"), exactMatchers, 0644); err != nil {
		return err
	}

	path := filepath.Join(GO_CODE_DIR, "00-locale-data.go")
	err = generateCode("locale-map", &finalLocaleData, path)
	if err != nil {
		return err
	}

	path = filepath.Join(GO_CODE_DIR, "01-language-order.go")
	err = generateCode("lang-order", &languageOrder, path)
	if err != nil {
		return err
	}

	path = filepath.Join(GO_CODE_DIR, "02-language-map.go")
	err = generateCode("lang-map", &languageMap, path)
	if err != nil {
		return err
	}

	path = filepath.Join(GO_CODE_DIR, "03-language-locales-map.go")
	err = generateCode("lang-loc-map", &languageLocalesMap, path)
	if err != nil {
		return err
	}

	path = filepath.Join(GO_CODE_DIR, "04-locale-order.go")
	err = generateCode("locale-order", &localeOrder, path)
	if err != nil {
		return err
	}

	for language, locales := range languageLocalesMap {
		var listLocaleData []LocaleData

		locales = append([]string{language}, locales...)
		for _, locale := range locales {
			if data, exist := finalLocaleData[locale]; exist {
				listLocaleData = append(listLocaleData, data)
			}
		}

		fName := strings.ToLower(language)
		fName = strings.ReplaceAll(fName, " ", "")
		path = filepath.Join(GO_CODE_DIR, fName+".go")
		err = generateLocaleDataCode(path, listLocaleData)
		if err != nil {
			return fmt.Errorf("gen locale %q failed: %w", language, err)
		}
	}

	// Render JSON (not used by code, just for debugging)
	os.RemoveAll(JSON_DIR)
	jsonDirs := map[string]map[string]LocaleData{
		filepath.Join(JSON_DIR, "cldr"):          cldrLocaleData,
		filepath.Join(JSON_DIR, "supplementary"): supplementalLocaleData,
		filepath.Join(JSON_DIR, "final"):         finalLocaleData,
	}

	for dir, locales := range jsonDirs {
		os.MkdirAll(dir, os.ModePerm)

		for language, data := range locales {
			dstPath := filepath.Join(dir, language+".json")
			err = renderJSON(&data, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func finalizeData(ld LocaleData) (LocaleData, error) {
	// Validate regex patterns
	if err := ld.Validate(); err != nil {
		return ld, err
	}

	// Clean up translations
	skipWords := mapset.New[string]()
	for _, w := range ld.SkipWords {
		skipWords.Put(w)
	}

	for word, translations := range ld.Translations {
		// If this word is supposed to be skipped, clear its translation
		if skipWords.Has(word) {
			translations = []string{""}
		} else {
			// Remove empty translations
			nonEmptyTranslations := slices.Clone(translations)
			nonEmptyTranslations = slices.DeleteFunc(nonEmptyTranslations, func(t string) bool {
				return t == ""
			})

			if len(nonEmptyTranslations) > 0 {
				translations = nonEmptyTranslations
			}
		}

		// Sort by the shortest
		sort.Slice(translations, func(i, j int) bool {
			transI, transJ := translations[i], translations[j]
			lenI := utf8.RuneCountInString(transI)
			lenJ := utf8.RuneCountInString(transJ)
			if lenI != lenJ {
				return lenI < lenJ
			}
			return transI < transJ
		})

		ld.Translations[word] = translations
	}

	// Save words that is known or always kept no matter what the language is
	for _, token := range importantTokens {
		ld.AddCharset(token)
		ld.Translations[token] = []string{token}
	}

	// Generate combined patterns
	ld.CombineRegexPatterns()
	ld.GenerateAbbreviations()
	ld.GenerateKnownWords()
	return ld, nil
}
