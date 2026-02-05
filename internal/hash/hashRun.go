package hash

import (
	"fmt"
	"os"

	"github.com/crytoken/consl"
)

func Run(cfg *HashConfig) {
	err := parseHashCfg(cfg)
	if err != nil {
		fmt.Println("Ertor:", err)
		os.Exit(0)
	}

	switch cfg.Method {
	case "SHA1":
		err := sha1Router(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
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
	case "SHA3-256":
		err := sha3_256_Router(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "SHA3-512":
		err := sha3_512_Router(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	case "SHA4":
		err := sha4Router(cfg)
		if err != nil {
			consl.PrintRed("Error; ")
			fmt.Println(err)
		}
	default:
		//supportedAlgo := []string{"SHA1", "SHA256", "SHA512", "SHA3-256", "SHA3-512", "SHA4"}
		consl.PrintRed("No such algorithm as: ")
		fmt.Println(cfg.Method)
		consl.PrintGreen("Supported Algorims: ")
		fmt.Println(supportedAlgo)
	}
}
