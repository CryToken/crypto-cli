package decrypt

import (
	"fmt"
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
		err := decryptAesRouter(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func decryptAesRouter(cfg *Config) error {
	switch cfg.MethodMode {
	case "CFB":
		err := decryptAesCFB(cfg)
		if err != nil {
			return err
		}
	case "GCM":
		err := decryptAES_GCM(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}
