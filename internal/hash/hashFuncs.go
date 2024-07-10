package hash

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func Run(cfg *HashConfig) {
	err := parseHashCfg(cfg)
	if err != nil {
		fmt.Println("Ertor:", err)
		os.Exit(1)
	}

	switch cfg.Method {
	case "SHA256":
		err := sha256router(cfg)
		if err != nil {
			fmt.Println("Hash err:", err)
		}
	}
}

// SHA256 Router and hadlers
func sha256router(cfg *HashConfig) error {
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
	fmt.Printf("SHA256 Hash:\n%s\n", result)
	return nil
}
