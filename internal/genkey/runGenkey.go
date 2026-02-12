package genkey

import (
	"fmt"
	"os"

	"github.com/crytoken/consl"
)

func Run(cfg *GenkeyConfig, args []string) {
	if err := parseGenkeyConfig(cfg, args); err != nil {
		fmt.Println("Parsing failed:", err)
		os.Exit(1)
	}

	switch cfg.Type {
	case ecdsaStr:
		if err := generateECDSAkeyPair(cfg.Output); err != nil {
			consl.PrintRed("ECDSA key pair generation failed: ")
			fmt.Println(err)
			os.Exit(1)
		}
	case ed25519Str:
		if err := generateED25519keyPair(cfg.Output); err != nil {
			consl.PrintRed("Ed-25519 key pair generation failed: ")
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
