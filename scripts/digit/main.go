package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const (
	RAW_PATH     = "./raw-data/digit/scripts.txt"
	GO_CODE_PATH = "./internal/digit/digit.go"
	UNIDATA_URL  = "https://www.unicode.org/Public/UNIDATA/Scripts.txt"
)

var (
	log = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05",
	}).With().Timestamp().Logger()

	rxParser = regexp.MustCompile(`^([^\.\s]+)\.{2}([^\.\s]+)[^\]]+\]\s([^\.]+)\.{2}([^\.]+)$`)
)

type DigitRange struct {
	Start string
	End   string
	Name  string
}

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

	// Parse file
	digitRanges, err := parseUnicodeScripts()
	if err != nil {
		return err
	}

	// Generate code
	return generateCode(digitRanges)
}

func downloadRawData() error {
	// Clean current dst dir
	rawDir := filepath.Dir(RAW_PATH)
	os.RemoveAll(rawDir)
	os.MkdirAll(rawDir, os.ModePerm)

	// Download file
	resp, err := http.Get(UNIDATA_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Save to storage
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(RAW_PATH, bt, os.ModePerm)
}

func parseUnicodeScripts() ([]DigitRange, error) {
	// Open file
	f, err := os.Open(RAW_PATH)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read each line
	var digitRanges []DigitRange

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Normalize line
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// We only care about digit
		if !strings.Contains(line, "# Nd") {
			continue
		}

		// Parse using regex
		parts := rxParser.FindStringSubmatch(line)
		if len(parts) != 5 {
			continue
		}

		digitRanges = append(digitRanges, DigitRange{
			Start: `\U` + padZero(parts[1]),
			End:   `\U` + padZero(parts[2]),
			Name:  fmt.Sprintf("%s - %s", parts[3], parts[4]),
		})
	}

	return digitRanges, nil
}

func generateCode(data []DigitRange) error {
	// Clean dst dir
	codeDir := filepath.Dir(GO_CODE_PATH)
	os.RemoveAll(codeDir)
	os.MkdirAll(codeDir, os.ModePerm)

	// Generate code
	b := bytes.NewBuffer(nil)
	err := tpl.Execute(b, data)
	if err != nil {
		return err
	}

	// Format code
	code := b.Bytes()
	code, err = format.Source(b.Bytes())
	if err != nil {
		return err
	}

	// Save to file
	return ioutil.WriteFile(GO_CODE_PATH, code, os.ModePerm)
}

func padZero(str string) string {
	return strings.Repeat("0", 8-len(str)) + str
}
