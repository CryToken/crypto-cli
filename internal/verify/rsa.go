package verify

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/crytoken/consl"
)

func verifyRSA(cfg *VeryfiConfig) error {
	pubKeyPemBytes, err := os.ReadFile(cfg.PublicKey)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(pubKeyPemBytes)
	if block == nil {
		return errInvalidPem
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}

	//Reading and Hash data
	dataToVerify, err := os.ReadFile(cfg.Data)
	if err != nil {
		return err
	}
	dataHash := sha256.Sum256(dataToVerify)

	//Reading signature
	signature, err := os.ReadFile(cfg.Signature)
	if err != nil {
		return err
	}

	//Verifing
	if err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, dataHash[:], signature); err != nil {
		return err
	}

	consl.PrintGreen("Signature is Valid.\n")
	return nil

}
