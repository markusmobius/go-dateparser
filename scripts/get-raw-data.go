//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const cldrVersion = "31.0.1"

func main() {
	// Find target dir in arguments
	rawDir := "./raw_data"
	if len(os.Args) > 1 {
		rawDir = os.Args[1]
	}

	// Clean current raw data
	os.RemoveAll(rawDir)

	// Fetch data from CLDR repository
	repos := []string{
		"https://github.com/unicode-cldr/cldr-dates-full.git",
		"https://github.com/unicode-cldr/cldr-core.git",
		"https://github.com/unicode-cldr/cldr-rbnf.git",
	}

	for _, repo := range repos {
		dirName := strings.TrimSuffix(path.Base(repo), ".git")
		dstDir := filepath.Join(rawDir, dirName)

		log.Printf("cloning %s from %s", dirName, repo)
		git.PlainClone(dstDir, false, &git.CloneOptions{
			URL:           repo,
			Depth:         1,
			SingleBranch:  true,
			Progress:      os.Stdout,
			ReferenceName: plumbing.NewTagReferenceName(cldrVersion),
		})

		fmt.Println()
	}
}
