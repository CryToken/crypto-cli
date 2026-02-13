package genkey

import (
	"errors"
	"fmt"
	"strings"

	"github.com/crytoken/crypto-cli/internal/tui"
)

var (
	errNotSupportedType error = errors.New("not supported key type")
)

var (
	ecdsaStr   string = "ECDSA"
	rsaStr     string = "RSA"
	dsaStr     string = "DSA"
	ed25519Str string = "ED-25519"
)

var supportedTypes []string = []string{ecdsaStr, rsaStr, dsaStr, ed25519Str}

func isSupportedGenkeyType(t string) bool {
	t = strings.ToUpper(t)
	for _, supportedAlgo := range supportedTypes {
		if t == supportedAlgo {
			return true
		}
	}
	return false
}

func parseGenkeyConfig(cfg *GenkeyConfig, args []string) error {
	//Type checking
	if len(args) == 1 {
		if ok := isSupportedGenkeyType(args[0]); !ok {
			return errNotSupportedType
		}
		cfg.Type = strings.ToUpper(args[0])
	} else {
		fmt.Println("Select key type:")
		if cfg.Type == "" {
			cfg.Type = tui.ChoiceItem(supportedTypes)
		}
		if ok := isSupportedGenkeyType(cfg.Type); !ok {
			return errNotSupportedType
		}
		cfg.Type = strings.ToUpper(cfg.Type)
	}

	//Output file path checking
	if cfg.Output == "" {
		var userInput string
		printEnterOutput(cfg.Type)
		fmt.Scanln(&userInput)

		cfg.Output = userInput
	}

	if cfg.Output == "" {
		cfg.Output = strings.ToLower(cfg.Type)
	}

	return nil
}

func printEnterOutput(t string) {
	fmt.Printf("Enter out file name (default: %s):", strings.ToLower(t))
}
