package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

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
	repo := "https://github.com/unicode-org/cldr-json.git"
	cloneDir := filepath.Join(RAW_DIR, "cldr-json")
	log.Info().Msgf("cloning CLDR %s from %s", CLDR_VERSION, repo)
	_, err := git.PlainClone(cloneDir, false, &git.CloneOptions{
		URL:           repo,
		Depth:         1,
		SingleBranch:  true,
		ReferenceName: plumbing.NewTagReferenceName(CLDR_VERSION),
	})
	if err != nil {
		return err
	}
	for _, module := range []string{"cldr-dates-full", "cldr-core", "cldr-units-full"} {
		if err := os.Rename(filepath.Join(cloneDir, "cldr-json", module), filepath.Join(RAW_DIR, module)); err != nil {
			return err
		}
	}
	return os.RemoveAll(cloneDir)
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
	bt, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	dstPath := filepath.Join(dstDir, "content_language.html")
	return os.WriteFile(dstPath, bt, os.ModePerm)
}
