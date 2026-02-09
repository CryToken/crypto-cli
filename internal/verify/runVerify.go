package verify

import (
	"fmt"
	"os"
)

func Run(cfg *VeryfiConfig) {
	if err := cfg.Parse(); err != nil {
		fmt.Println("Parsing config err:", err)
		os.Exit(1)
	}

	switch cfg.Algorithm {
	case ecdsaStr:
		if err := verifyECDSA(cfg); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Signature: is Valid.")
		os.Exit(0)
	}
}
