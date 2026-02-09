package verify

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func verifyECDSA(cfg *VeryfiConfig) error {
	fmt.Println("Pubkey filepath:", cfg.PublicKey)
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

	fmt.Printf("Public key hex:%x\n", block.Bytes)
	pubKey, err := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), block.Bytes)
	if err != nil {
		return err
	}
	fmt.Printf("Pub key struct : %+v\n", pubKey)

	signature, err := os.ReadFile(cfg.Signature)
	if err != nil {
		return err
	}
	dataHash := sha256.Sum256([]byte(cfg.Data))
	if ok := ecdsa.VerifyASN1(pubKey, dataHash[:], signature); !ok {
		return errors.New("verify failed")
	}

	return nil
}
