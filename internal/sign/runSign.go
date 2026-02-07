package sign

import (
	"fmt"
	"os"
)

func Run(cfg *SignConfig) {
	if err := cfg.Parse(); err != nil {
		fmt.Println("Parsing config err:", err)
		os.Exit(1)
	}

	//fmt.Printf("You run sign cmd\n%+v\n", cfg)

	switch cfg.Algorithm {
	case ecdsaStr:
		if err := signECDSA(cfg); err != nil {
			fmt.Printf("ecdsa signig failed:%s\n", err)
			os.Exit(1)
		}
	case rsaStr:
		if err := signRSA(cfg); err != nil {
			fmt.Printf("rsa signig failed:%s\n", err)
			os.Exit(1)
		}
	case dsaStr:
		if err := signDSA(cfg); err != nil {
			fmt.Printf("dsa signig failed:%s\n", err)
			os.Exit(1)
		}
	}
}
