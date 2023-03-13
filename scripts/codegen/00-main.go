package main

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/elliotchance/pie/v2"
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

	err := rootCmd.Execute()
	if err != nil {
		log.Panic().Err(err)
	}
}

func rootCmdHandler(cmd *cobra.Command, args []string) error {
	// Parse flags
	skipRawDownload, _ := cmd.Flags().GetBool("skip-raw")

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

	// Generate map between a language and its descendant (if any)
	languageMap := createLanguageMap(languageOrder)

	// Generate map of locale order
	localeOrder := createLocaleOrder(languageLocalesMap, languageOrder)

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

	// Generate code
	os.RemoveAll(GO_CODE_DIR)
	os.MkdirAll(GO_CODE_DIR, os.ModePerm)

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
			return err
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
			// If there are non empty translations, remove the empty ones
			nonEmptyTranslations := pie.Filter(translations, func(t string) bool {
				return t != ""
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
