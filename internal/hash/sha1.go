package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/crytoken/consl"
)

func sha1Router(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha1Text(cfg)
		if err != nil {
			return err
		}
	case "File":
		err := sha1File(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func sha1Text(cfg *HashConfig) error {
	data := []byte(cfg.Data)
	hasher := sha1.New()

	hasher.Write(data)
	result := hex.EncodeToString(hasher.Sum(nil))
	consl.PrintBlue("SHA-1:\n")
	fmt.Println(result)
	return nil
}

func sha1File(cfg *HashConfig) error {
	//Read file and close befor quit
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha1.New()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	consl.PrintBlue("SHA-1 File Hash:\n")
	fmt.Println(hashString)
	return nil
}
