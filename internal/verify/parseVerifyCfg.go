package verify

import (
	"errors"
	"fmt"
	"strings"

	"github.com/crytoken/crypto-cli/internal/tui"
)

var (
	errNotSupportedVerifyAlgo error = errors.New("not supported verify alogorithm")
	errNotSupportedHashAlgo   error = errors.New("not supported hash alogorithm")
	errEmptyKeyFile           error = errors.New("key file must be set by (-k flag)")
	errEmptySigFile           error = errors.New("signature file is reqired (-s flag)")
)

var (
	ecdsaStr   string = "ECDSA"
	rsaStr     string = "RSA"
	dsaStr     string = "DSA"
	ed25519Str string = "ED-25519"
)
var (
	sha256Str string = "SHA-256"
	sha512Str string = "SHA-512"
	argon2Str string = "argon2"
)

var supportedVerifyAlogos []string = []string{ecdsaStr, rsaStr, dsaStr, ed25519Str}
var supportedHashAlogos []string = []string{sha256Str, sha512Str, argon2Str}

func isSupportedVerifyAlgo(algo string) bool {
	algo = strings.ToUpper(algo)
	for _, supportedAlgo := range supportedVerifyAlogos {
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

func (cfg *VeryfiConfig) Parse() error {
	if cfg.Algorithm == "" {
		cfg.Algorithm = tui.ChoiceItem(supportedVerifyAlogos)
	}
	if ok := isSupportedVerifyAlgo(strings.ToUpper(cfg.Algorithm)); !ok {
		return errNotSupportedVerifyAlgo
	}

	if ok := isSupportedHashAlgo(strings.ToUpper(cfg.HashAlgo)); !ok {
		return errNotSupportedHashAlgo
	}

	if cfg.Data == "" {
		fmt.Println("Choose data file:")
		var dataFileChoose string
		if err := tui.SelectFile(&dataFileChoose); err != nil {
			return errEmptyKeyFile
		}
		cfg.Data = dataFileChoose
		fmt.Print(cfg.Data, "\n")
	}

	if cfg.PublicKey == "" {
		fmt.Println("Choose public key file:")
		var keyFile string
		if err := tui.SelectFile(&keyFile); err != nil {
			return errEmptyKeyFile
		}
		cfg.PublicKey = keyFile
	}

	if cfg.Signature == "" {
		fmt.Println("Choose signature file:")
		var sigFile string
		if err := tui.SelectFile(&sigFile); err != nil {
			return errEmptySigFile
		}
		cfg.Signature = sigFile
	}

	return nil

}
