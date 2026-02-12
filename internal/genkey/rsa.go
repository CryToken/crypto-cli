package genkey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var (
	errPrivateValidationFailed error = errors.New("private key Validation failed")
)

func generateRSAkeyPair(out string) error {
	privateKeyFile := fmt.Sprintf("%s.pem", out)
	publicKeyFile := fmt.Sprintf("%s_public.pem", out)

	prvkey, err := rsa.GenerateKey(rand.Reader, 1024*4)
	if err != nil {
		return err
	}

	if err := prvkey.Validate(); err != nil {
		return errPrivateValidationFailed
	}

	prvKeyBytes := x509.MarshalPKCS1PrivateKey(prvkey)
	prvPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: prvKeyBytes,
	}
	encodedPrivateKey := pem.EncodeToMemory(prvPemBlock)
	if err := os.WriteFile(privateKeyFile, encodedPrivateKey, 0400); err != nil {
		return errWriteToFile
	}

	//Puvlic key part
	pubkey := prvkey.PublicKey
	pubkeyBytes := x509.MarshalPKCS1PublicKey(&pubkey)
	pubkeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubkeyBytes,
	}

	encodedPublicKey := pem.EncodeToMemory(pubkeyBlock)
	if err := os.WriteFile(publicKeyFile, encodedPublicKey, 0444); err != nil {
		return errWriteToFile
	}
	return nil
}
