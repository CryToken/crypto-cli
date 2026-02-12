package genkey

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	errNotSupportedType error = errors.New("not supported key type")
	errNoOutFile        error = errors.New("there is no output file")
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
		if cfg.Type == "" {
			fmt.Println("You should set type of key pair with arg or -t flag")
			os.Exit(1)
		}
		if ok := isSupportedGenkeyType(cfg.Type); !ok {
			return errNotSupportedType
		}
		cfg.Type = strings.ToUpper(cfg.Type)
	}

	//Output file path checking
	if cfg.Output == "" {
		fmt.Println("You must provide provide output file name by -o flag.")
		return errNoOutFile
	}

	return nil
}
