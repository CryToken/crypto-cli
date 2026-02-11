package sign

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/crytoken/crypto-cli/internal/tui"
)

var (
	errNotSupportedSignAlgo error = errors.New("not supported sign alogorithm")
	errNotSupportedHashAlgo error = errors.New("not supported hash alogorithm")
	errEmptyKeyFile         error = errors.New("key file must be set by (-k flag)")
)

func (cfg *SignConfig) Parse() error {
	if cfg.Algorithm == "" {
		cfg.Algorithm = tui.ChoiceItem(supportedSignAlogos)
	}
	//Check sing algoruhms
	cfg.Algorithm = strings.ToUpper(cfg.Algorithm)
	if ok := isSupportedAlgo(cfg.Algorithm); !ok {
		return errNotSupportedSignAlgo
	}
	//Check hash func
	cfg.HashAlgo = strings.ToUpper(cfg.HashAlgo)
	if ok := isSupportedHashAlgo(cfg.HashAlgo); !ok {
		return errNotSupportedHashAlgo
	}

	if cfg.KeyFile == "" {
		var keyChoose string
		fmt.Println("You need to choose private key file:")
		if err := tui.SelectFile(&keyChoose); err != nil {
			return err
		}
		cfg.KeyFile = keyChoose
		if cfg.KeyFile == "" {
			return errEmptyKeyFile
		}

	}

	//Input part
	if cfg.Input == "" {
		var choose string
		fmt.Println("You need to choose file to sign.")
		if err := tui.SelectFile(&choose); err != nil {
			return err
		}
		cfg.Input = choose
	}

	//Output part
	if cfg.Output == "" {
		cfg.OutputWriter = os.Stdout

	} else {
		f, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_RDWR, 0600)
		if err != nil {
			fmt.Println(cfg.Output, err)
			return errOpenFile
		}
		cfg.OutputWriter = f
	}
	return nil

}

func isSupportedAlgo(algo string) bool {
	algo = strings.ToUpper(algo)
	for _, supportedAlgo := range supportedSignAlogos {
		if algo == supportedAlgo {
			return true
		}
	}
	return false
}

func isSupportedHashAlgo(algo string) bool {
	algo = strings.ToUpper(algo)
	for _, supportedAlgo := range supportedHashAlogos {
		if algo == supportedAlgo {
			return true
		}
	}
	return false
}
