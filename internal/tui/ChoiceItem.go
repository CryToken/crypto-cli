package tui

import "github.com/charmbracelet/huh"

func ChoiceItem(list []string) string {
	options := stringToOptions(list)

	var choice string
	theme := huh.ThemeCatppuccin()
	huh.NewSelect[string]().Title("Select:").Options(options...).Value(&choice).Height(8).WithTheme(theme).Run()

	return choice
}
