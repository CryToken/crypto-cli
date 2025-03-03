package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func ChooseItem(items []string) string {
	prompt := promptui.Select{
		Label: "Select:",
		Items: items,
		Size:  7,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Println("Selection failed:", err)
		return ""
	}

	return result
}
