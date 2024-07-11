package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/crytoken/consl"
)

func encryptAes(cfg *Config) error {
	switch cfg.MethodMode {
	case "CFB":
		err := encryptAesCFB(cfg)
		return err
	case "GCM":
		err := encryptAES_GCM(cfg)
		return err
	}
	return nil
}

func encryptAesCFB(cfg *Config) error {
	block, err := aes.NewCipher(cfg.KeyHash)
	if err != nil {
		return err
	}

	//Read file
	fileBytes, err := ioutil.ReadFile(cfg.InputFile)
	if err != nil {
		return err
	}

	//Cipher data
	cipherData := make([]byte, aes.BlockSize+len(fileBytes))
	iv := cipherData[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherData[aes.BlockSize:], fileBytes)

	err = ioutil.WriteFile(cfg.OutputFile, cipherData, 0644)
	if err != nil {
		return err
	}
	consl.PrintGreen("Success:\n")
	consl.PrintCyan("    Encrypted File -> ")
	fmt.Println(cfg.OutputFile)
	return nil
}

func encryptAES_GCM(cfg *Config) error {
	//Open file
	file, err := os.Open(cfg.InputFile)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	//Create file to save encrypted data
	outputFile, err := os.Create(cfg.OutputFile)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer outputFile.Close()

	//Start encrypted
	block, err := aes.NewCipher(cfg.KeyHash)
	if err != nil {
		return errors.New("block error")
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return errors.New("creating gcm cipher")
	}

	//Creating nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	if _, err := outputFile.Write(nonce); err != nil {
		return err
	}

	//Encrypting
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		ciphertext := aesGCM.Seal(nil, nonce, buf[:n], nil)
		if _, err := outputFile.Write(ciphertext); err != nil {
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
