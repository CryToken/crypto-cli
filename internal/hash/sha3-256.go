package hash

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/crytoken/consl"
	"golang.org/x/crypto/sha3"
)

func sha3_256_Router(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha3_256_Text(cfg)
		if err != nil {
			return err
		}
	case "File":
		err := sha3_256_File(cfg)
		if err != nil {
			return nil
		}
	}
	return nil
}

func sha3_256_Text(cfg *HashConfig) error {
	data := []byte(cfg.Data)
	hasher := sha3.New256()

	hasher.Write(data)
	result := hex.EncodeToString(hasher.Sum(nil))
	consl.PrintBlue("SHA3-256:\n")
	fmt.Println(result)
	return nil
}

func sha3_256_File(cfg *HashConfig) error {
	//Read file and close befor quit
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha3.New256()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	consl.PrintBlue("SHA3-256 File Hash:\n")
	fmt.Println(hashString)
	return nil
}
