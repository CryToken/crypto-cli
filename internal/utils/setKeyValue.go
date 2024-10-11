package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crytoken/consl"
)

func SetKeyValue(str *string) error {
	var enteredData string
	scanner := bufio.NewScanner(os.Stdin) // Create a new scanner for input

	fmt.Println("Enter passkey :")
	for {
		// Scan input from user
		if scanner.Scan() {
			enteredData = scanner.Text()
		} else {
			return scanner.Err() // Return error if scanning fails
		}

		// Validate input length
		if len(enteredData) > 0 {
			break // Exit the loop if input is valid
		}

		consl.PrintYellow("Key must be > 1 \n")
	}

	*str = enteredData // Set the passed string to the entered value
	return nil
}
