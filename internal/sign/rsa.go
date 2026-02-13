package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func signRSA(cfg *SignConfig) error {
	prvKeyPEMbytes, err := os.ReadFile(cfg.KeyFile)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(prvKeyPEMbytes)
	if block == nil {
		return errInvalidPem
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return errParsingPrivKey
	}

	//Read and Hash data file
	dataToSign, err := os.ReadFile(cfg.Input)
	if err != nil {
		return errReadingFile
	}

	dataHash := sha256.Sum256(dataToSign)

	//Signing data digest

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dataHash[:])
	if err != nil {
		return err
	}

	if err := os.WriteFile(cfg.Output, signature, 0660); err != nil {
		return err
	}

	return nil
}
