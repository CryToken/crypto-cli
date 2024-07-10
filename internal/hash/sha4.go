package hash

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/crytoken/consl"
	"github.com/crytoken/crypto-cli/pkg/sha4"
)

func sha4Router(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha4Text(cfg)
		if err != nil {
			return err
		}
	case "File":
		err := sha4File(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func sha4Text(cfg *HashConfig) error {
	data := []byte(cfg.Data)
	hasher := sha4.New()

	hasher.Write(data)
	result := hex.EncodeToString(hasher.Sum(nil))
	consl.PrintBlue("SHA-4:\n")
	fmt.Println(result)
	return nil
}

func sha4File(cfg *HashConfig) error {
	//Read file and close befor quit
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha4.New()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	consl.PrintBlue("SHA-4 File Hash:\n")
	fmt.Println(hashString)
	return nil
}
