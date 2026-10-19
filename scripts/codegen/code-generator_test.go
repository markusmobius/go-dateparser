package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveGeneratedCodeKeepsTests(t *testing.T) {
	dir := t.TempDir()
	kept := map[string]bool{
		"00-locale-data.go":   false,
		"en.go":               false,
		"locale-data_test.go": true,
	}
	for name := range kept {
		if err := os.WriteFile(filepath.Join(dir, name), nil, 0644); err != nil {
			t.Fatal(err)
		}
	}

	if err := removeGeneratedCode(dir); err != nil {
		t.Fatalf("removeGeneratedCode(%q) = %v", dir, err)
	}

	for name, want := range kept {
		_, err := os.Stat(filepath.Join(dir, name))
		if got := err == nil; got != want {
			t.Errorf("after removeGeneratedCode, %s kept = %v; want %v", name, got, want)
		}
	}
}
