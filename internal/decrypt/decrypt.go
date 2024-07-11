package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io/ioutil"
)

func Run(cfg *Config, args []string) {
	err := parseCfg(cfg, args)
	if err != nil {
		fmt.Println("Decrypting error:", err)
	}

	err = decryptFile(cfg)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func decryptFile(cfg *Config) error {
	switch cfg.Method {
	case "AES":
		err := decryptAes(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func decryptAes(cfg *Config) error {
	switch cfg.MethodMode {
	case "CFB":
		err := decryptAesCFB(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

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
