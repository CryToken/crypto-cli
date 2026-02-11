package sign

import (
	"crypto/ed25519"
	"encoding/pem"
	"os"
)

func signEd25516(cfg *SignConfig) error {
	pemBytes, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(pemBytes)
	privateKey := ed25519.PrivateKey(block.Bytes)
	//Read data to sign
	dataToSign, err := os.ReadFile(cfg.Input)
	if err != nil {
		return err
	}

	//Signing
	signature := ed25519.Sign(privateKey, dataToSign)

	//Saving sig to file
	if err := os.WriteFile(cfg.Output, signature, 0600); err != nil {
		return err
	}

	return nil
}
