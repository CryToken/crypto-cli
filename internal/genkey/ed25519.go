package genkey

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"os"
)

func generateED25519keyPair(out string) error {
	outPrivateKey := fmt.Sprintf("%s.pem", out)
	outPublicKey := fmt.Sprintf("%s_public.pem", out)

	pub, prv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	prvRawBytes := []byte(prv)
	prvPemBlock := &pem.Block{
		Type:  "Ed-25519 PRIVATE KEY",
		Bytes: prvRawBytes,
	}

	encodedPrvKey := pem.EncodeToMemory(prvPemBlock)
	if err := os.WriteFile(outPrivateKey, encodedPrvKey, 0400); err != nil {
		return errWriteToFile
	}

	//Public key pair
	pubRawBytes := []byte(pub)
	pubPemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubRawBytes,
	}

	encodedPublicKey := pem.EncodeToMemory(pubPemBlock)
	if err := os.WriteFile(outPublicKey, encodedPublicKey, 0444); err != nil {
		return errWriteToFile
	}

	return nil
}
