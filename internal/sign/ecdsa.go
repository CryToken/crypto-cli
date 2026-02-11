package sign

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	errOpenFile       error = errors.New("open file failed")
	errParsingPrivKey error = errors.New("error parsing private key")
)

func signECDSA(cfg *SignConfig) error {
	f, err := os.Open(cfg.KeyFile)
	if err != nil {
		return errOpenFile
	}
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	buf := make([]byte, fileInfo.Size())
	_, err = io.ReadFull(f, buf)
	if err != nil && err != io.EOF {
		return err
	}
	pemBlock, _ := pem.Decode(buf)
	if pemBlock == nil {
		fmt.Println("Invalid PEM key format ")
		os.Exit(1)
	}

	prvKey, err := x509.ParseECPrivateKey(pemBlock.Bytes)
	if err != nil {
		return errParsingPrivKey
	}

	//Data to sign
	dataToSign, err := os.ReadFile(cfg.Input)
	if err != nil {
		return err
	}
	dataHash := sha256.Sum256(dataToSign)
	signature, err := ecdsa.SignASN1(rand.Reader, prvKey, dataHash[:])
	if err != nil {
		return err
	}

	if err := writeOutput(cfg.OutputWriter, signature); err != nil {
		return err
	}

	fmt.Println("File signed succesfuly.\nSignature save in:", cfg.Output)
	return nil
}

func writeOutput(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
