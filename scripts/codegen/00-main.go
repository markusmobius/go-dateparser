package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
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
	createLanguageMap(languageOrder)

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

	// Render CLDR locale data to JSON
	os.RemoveAll(CLDR_DIR)
	os.MkdirAll(CLDR_DIR, os.ModePerm)

	for language, localeData := range cldrLocaleData {
		dstPath := filepath.Join(CLDR_DIR, language+".json")
		err = renderJSON(&localeData, dstPath)
		if err != nil {
			return err
		}
	}

	// Render supplemental locale data to JSON
	os.RemoveAll("data/supplementary/json")
	os.MkdirAll("data/supplementary/json", os.ModePerm)

	for language, localeData := range supplementalLocaleData {
		dstPath := filepath.Join("data/supplementary/json", language+".json")
		err = renderJSON(&localeData, dstPath)
		if err != nil {
			return err
		}
	}

	return nil
}
