package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func Run(cfg *Config) {
	err := parseCfg(cfg)
	if err != nil {
		fmt.Println("Errors:", err)
		return
	}
	encrypt(cfg)

}

func encrypt(cfg *Config) error {
	switch cfg.Method {
	case "AES":
		err := encryptAes(cfg)
		if err != nil {
			fmt.Println("Encryption went wrong", err)
		}
	}
	return nil
}

func encryptAes(cfg *Config) error {
	switch cfg.MethodMode {
	case "CFB":
		err := encryptAesCFB(cfg)
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
	return nil
}
