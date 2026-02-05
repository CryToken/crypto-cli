package tui

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/charmbracelet/huh"
)

func loadDirEntitis(path string) ([]string, error) {
	entilries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contained []string
	contained = append(contained, "..")

	for _, entri := range entilries {
		name := entri.Name()
		if entri.IsDir() {
			name = name + "/"
		}
		contained = append(contained, name)

	}

	sort.Strings(contained[1:])
	return contained, nil
}

func parentDir(path string) string {
	parent := filepath.Dir(filepath.Clean(path))
	if parent == "." {
		return "/"
	}
	return parent
}

func stringToOptions(list []string) []huh.Option[string] {

	var options []huh.Option[string]

	for _, entri := range list {
		option := huh.NewOption(entri, entri)
		options = append(options, option)
	}
	return options
}
