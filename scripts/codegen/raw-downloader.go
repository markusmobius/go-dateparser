package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func downloadRawData() error {
	// Clean current dst dir
	os.RemoveAll(RAW_DIR)

	// Fetch data from CLDR repository
	repos := []string{
		"https://github.com/unicode-cldr/cldr-dates-full.git",
		"https://github.com/unicode-cldr/cldr-core.git",
		"https://github.com/unicode-cldr/cldr-rbnf.git",
	}

	for _, repo := range repos {
		dirName := strings.TrimSuffix(path.Base(repo), ".git")
		dstDir := filepath.Join(RAW_DIR, dirName)

		log.Info().Msgf("cloning %s from %s", dirName, repo)
		_, err := git.PlainClone(dstDir, false, &git.CloneOptions{
			URL:           repo,
			Depth:         1,
			SingleBranch:  true,
			ReferenceName: plumbing.NewTagReferenceName(CLDR_VERSION)})
		if err != nil {
			return err
		}
	}

	return nil
}
