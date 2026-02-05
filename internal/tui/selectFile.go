package tui

import (
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func SelectFile(res *string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dirEntries, err := loadDirEntitis(pwd)
	if err != nil {
		return err
	}
	options := stringToOptions(dirEntries)

	var choice string
	tuiTheme := huh.ThemeCatppuccin()
	huh.NewSelect[string]().Title("Select file").Options(options...).Value(&choice).Height(8).WithTheme(tuiTheme).Run()

	for {

		if choice == ".." {
			pwd = parentDir(pwd)
			newDirEntries, err := loadDirEntitis(pwd)
			if err != nil {
				return err
			}
			options = stringToOptions(newDirEntries)
			huh.NewSelect[string]().Title("Select file").Options(options...).Value(&choice).Height(18).Run()
			continue

		}

		if strings.HasSuffix(choice, "/") {
			choice = strings.TrimRight(choice, "/")
			pwd = pwd + "/" + choice
			dirEntries, err = loadDirEntitis(pwd)
			if err != nil {
				return err
			}
			options = stringToOptions(dirEntries)
			huh.NewSelect[string]().Title("Select file").Options(options...).Value(&choice).Height(8).Run()

		} else {
			*res = pwd + "/" + choice

			break
		}
	}
	return nil

}
