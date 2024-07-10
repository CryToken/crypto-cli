package hash

import (
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
	case "SHA512":
		err := sha512Router(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
