package hash

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func Run(cfg *HashConfig) {
	err := parseHashCfg(cfg)
	if err != nil {
		fmt.Println("Somthing went wrong:", err)
		os.Exit(1)
	}

	switch cfg.Method {
	case "SHA256":
		err := sha256hash(cfg)
		if err != nil {
			fmt.Println("Hash err:", err)
		}
	}
}

func sha256hash(cfg *HashConfig) error {
	switch cfg.Mode {
	case "Text":
		err := sha256Text(cfg)
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
	fmt.Println("SHA256 Hash:", result)
	return nil
}
