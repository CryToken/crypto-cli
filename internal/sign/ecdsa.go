package sign

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/pem"
	"errors"
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
	prvKey, err := ecdsa.ParseRawPrivateKey(cfg.Curve, pemBlock.Bytes)
	if err != nil {
		return errParsingPrivKey
	}

	//Data to sign
	dataToSign, err := os.ReadFile(cfg.Input)
	if err != nil {
		return err
	}
	signature, err := ecdsa.SignASN1(rand.Reader, prvKey, dataToSign)
	if err != nil {
		return err
	}

	if err := writeOutput(cfg.OutputWriter, signature); err != nil {
		return err
	}

	return nil
}

func writeOutput(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
