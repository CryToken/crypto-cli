package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/crytoken/consl"
)

func decryptAesCFB(cfg *Config) error {
	cipherFile, err := ioutil.ReadFile(cfg.InputFile)
	if err != nil {
		return errors.New("reading cipherFile err")
	}

	block, err := aes.NewCipher(cfg.KeyHash)
	if err != nil {
		return errors.New("err to create block")
	}
	iv := cipherFile[:aes.BlockSize]
	cipherFile = cipherFile[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherFile, cipherFile)

	err = ioutil.WriteFile(cfg.OutputFile, cipherFile, 0644)
	if err != nil {
		return errors.New("err to write to file")
	}
	return nil
}

func decryptAES_GCM(cfg *Config) error {
	//Open encrypted file
	encryptedFile, err := os.Open(cfg.InputFile)
	if err != nil {
		return errors.New("err to open file")
	}
	defer encryptedFile.Close()

	//Create output File
	outputFile, err := os.Create(cfg.OutputFile)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer outputFile.Close()

	block, err := aes.NewCipher(cfg.KeyHash)
	if err != nil {
		return errors.New("block init failed")
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return errors.New("init decrypt failed")
	}

	nonceSize := aesGCM.NonceSize()
	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(encryptedFile, nonce); err != nil {
		return err
	}

	buf := make([]byte, 1024+aesGCM.Overhead())
	for {
		n, err := encryptedFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		plaintext, err := aesGCM.Open(nil, nonce, buf[:n], nil)
		if err != nil {
			return err
		}
		if _, err := outputFile.Write(plaintext); err != nil {
			return err
		}
	}
	//Print result
	outDir := filepath.Dir(cfg.OutputFile)
	currentDir, _ := os.Getwd()
	if outDir == currentDir {
		cfg.OutputFile = filepath.Base(cfg.OutputFile)
	}

	consl.PrintGreen("Success:\n")
	consl.PrintCyan("    Encrypted File -> ")
	fmt.Println(cfg.OutputFile)

	return nil
}
