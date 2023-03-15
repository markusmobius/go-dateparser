package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func downloadRawData() error {
	os.RemoveAll(RAW_DIR)

	err := downloadCldrData()
	if err != nil {
		return err
	}

	err = downloadW3ContentLanguage()
	return err
}

func downloadCldrData() error {
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

func downloadW3ContentLanguage() error {
	// Prepare dst dir
	dstDir := filepath.Join(RAW_DIR, "w3techs")
	err := os.MkdirAll(dstDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Download page
	url := "https://w3techs.com/technologies/overview/content_language"
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Save to file
	bt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dstPath := filepath.Join(dstDir, "content_language.html")
	return os.WriteFile(dstPath, bt, os.ModePerm)
}
