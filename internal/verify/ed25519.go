package verify

import (
	"crypto/ed25519"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func verifyEd25519(cfg *VeryfiConfig) error {
	pemBytes, err := os.ReadFile(cfg.PublicKey)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return errors.New("nil block.invalid pubkey pem format")
	}

	publicKey := ed25519.PublicKey(block.Bytes)

	//Read data to verify
	dataToVerify, err := os.ReadFile(cfg.Data)
	if err != nil {
		return err
	}

	//Read signature
	signature, err := os.ReadFile(cfg.Signature)
	if err != nil {
		return err
	}

	//Verifing
	if ok := ed25519.Verify(publicKey, dataToVerify, signature); !ok {
		return errors.New("not valid signature")
	}

	fmt.Println("Signature is valid.")
	return nil

}
