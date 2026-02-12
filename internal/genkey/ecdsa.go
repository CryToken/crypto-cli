package genkey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var (
	errGenPrivateKey error = errors.New("ecdsa key generation failed")
	errWriteToFile   error = errors.New("write to file failed")
)

func generateECDSAkeyPair(out string) error {
	privateKeyFile := fmt.Sprintf("%s.pem", out)
	publicKeyFile := fmt.Sprintf("%s_public.pem", out)

	prvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return errGenPrivateKey
	}
	prvKeyBytes, _ := x509.MarshalECPrivateKey(prvKey)
	//fmt.Println("Priv key hex:", hex.EncodeToString(prvKeyBytes))

	block := pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: prvKeyBytes,
	}

	PrvPemBytes := pem.EncodeToMemory(&block)
	if err := os.WriteFile(privateKeyFile, PrvPemBytes, 0600); err != nil {
		return errWriteToFile
	}

	//Pubkey save part
	pubKeyBytes, _ := x509.MarshalPKIXPublicKey(&prvKey.PublicKey)
	fmt.Println("Pub key hex:", hex.EncodeToString(pubKeyBytes))
	pubBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}

	pubPemBytes := pem.EncodeToMemory(&pubBlock)
	if err := os.WriteFile(publicKeyFile, pubPemBytes, 0644); err != nil {
		return errWriteToFile
	}

	return nil
}
