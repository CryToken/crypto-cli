package verify

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func verifyECDSA(cfg *VeryfiConfig) error {

	pubkeyPEM, err := os.ReadFile(cfg.PublicKey)
	if err != nil {
		return err
	}

	block, rest := pem.Decode(pubkeyPEM)
	if block == nil {
		fmt.Println("block nil")
		fmt.Printf("Rest part: %x\n", rest)
		os.Exit(1)
	}

	//fmt.Printf("Public key hex:%x\n", block.Bytes)
	pubKeyX, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pubKey, ok := pubKeyX.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("parsing public key error")
	}
	//fmt.Printf("Pub key struct : %+v\n", pubKey)

	//Readimg signature
	signature, err := os.ReadFile(cfg.Signature)
	if err != nil {
		return err
	}
	//Reading data
	dataToVerify, err := os.ReadFile(cfg.Data)
	if err != nil {
		return err
	}

	dataHash := sha256.Sum256(dataToVerify)
	if ok := ecdsa.VerifyASN1(pubKey, dataHash[:], signature); !ok {
		return errors.New("verify failed")
	}

	return nil
}
