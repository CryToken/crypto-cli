package verify

import (
	"fmt"
	"os"

	"github.com/crytoken/consl"
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
		consl.PrintGreen("Signature: is valid.\n")
		os.Exit(0)
	case ed25519Str:
		if err := verifyEd25519(cfg); err != nil {
			consl.PrintRed(err, "\n")
			os.Exit(1)
		}
	case rsaStr:
		if err := verifyRSA(cfg); err != nil {
			consl.PrintRed(err, "\n")
			os.Exit(1)
		}
	}
}
