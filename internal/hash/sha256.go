package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/crytoken/consl"
)

// SHA256 Router and hadlers
func sha256router(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha256Text(cfg)
		if err != nil {
			return err
		}
	case "File":
		err := sha256File(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func sha256Text(cfg *HashConfig) error {
	data := []byte(cfg.Data)
	hasher := sha256.New()
	hasher.Write(data)
	result := fmt.Sprintf("%x", hasher.Sum(nil))
	consl.PrintBlue("SHA-256:\n")
	fmt.Printf("%s\n", result)
	return nil
}

func sha256File(cfg *HashConfig) error {
	//Read file and close befor quit
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return err
	}

	defer file.Close()

	hash := sha256.New()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	consl.PrintBlue("SHA-256 File Hash:\n")
	fmt.Println(hashString)
	return nil
}
