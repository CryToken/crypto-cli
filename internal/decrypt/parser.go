package decrypt

import (
	"errors"
	"fmt"

	"github.com/crytoken/crypto-cli/internal/utils"
)

func parseCfg(cfg *Config) error {
	//Check that key is not empty
	if cfg.Key == "" {
		return errors.New("enter key by -k flag")
	}
	//Check the InputFile is choicen
	if cfg.InputFile == "" {
		return errors.New("enter filepath by -p flag")
	}

	if cfg.OutputFile == "" {
		res := fmt.Sprintf("%s_decr", cfg.InputFile)
		cfg.OutputFile = res
	}
	switch cfg.KeyMode {
	case "SHA256":
		cfg.KeyHash = utils.Sha256hash(cfg.Key)
	case "Sha1":
		cfg.KeyHash = utils.Sha256hash(cfg.Key)
	}
	return nil

}
