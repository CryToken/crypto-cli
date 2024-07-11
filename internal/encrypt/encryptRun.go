package encrypt

import (
	"fmt"
)

func Run(cfg *Config, args []string) {
	err := parseCfg(cfg, args)
	if err != nil {
		fmt.Println("Errors:", err)
		return
	}
	encryptRouter(cfg)

}

func encryptRouter(cfg *Config) error {
	switch cfg.Method {
	case "AES":
		err := encryptAes(cfg)
		if err != nil {
			fmt.Println("Encryption went wrong", err)
		}
	}
	return nil
}
