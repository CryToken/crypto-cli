package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/crytoken/consl"
)

func sha512Router(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha512Text(cfg)
		if err != nil {
			return err
		}
	case "File":
		err := sha512File(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func sha512Text(cfg *HashConfig) error {
	data := []byte(cfg.Data)
	hasher := sha512.New()

	hasher.Write(data)
	result := hex.EncodeToString(hasher.Sum(nil))
	consl.PrintBlue("SHA-512:\n")
	fmt.Println(result)
	return nil
}

func sha512File(cfg *HashConfig) error {
	//Read file and close befor quit
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha512.New()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	consl.PrintBlue("SHA-512 File Hash:\n")
	fmt.Println(hashString)
	return nil
}
