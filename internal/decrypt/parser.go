package decrypt

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/crytoken/consl"
	"github.com/crytoken/crypto-cli/internal/utils"
	"github.com/crytoken/crypto-cli/pkg/sha4"
)

func parseCfg(cfg *Config, args []string) error {
	cfg.KeyMode = strings.ToUpper(cfg.KeyMode)
	cfg.MethodMode = strings.ToUpper(cfg.MethodMode)

	//Check the InputFile is choicen
	if cfg.InputFile == "" && len(args) < 1 {
		err := utils.SelectFile(&cfg.InputFile)
		if err != nil {
			return err
		}
		//return errors.New("choice file by -f flag or by argument")
	}
	if cfg.InputFile == "" && len(args) == 1 {
		cfg.InputFile = args[0]
	}

	//Check that key is not empty
	if cfg.Key == "" {
		err := utils.SetKeyValue(&cfg.Key)
		if err != nil {
			return err
		}
	}

	if cfg.OutputFile == "" {
		path, _ := os.Getwd()
		filename := filepath.Base(cfg.InputFile)
		res := fmt.Sprintf("%s/decr_%s", path, filename)
		cfg.OutputFile = res
	}
	switch cfg.KeyMode {
	case "SHA256":
		cfg.KeyHash = utils.Sha256hash(cfg.Key)
	case "SHA4":
		hasher := sha4.New()
		cfg.KeyHash = hasher.Hash([]byte(cfg.Key))
	default:
		consl.PrintRed("Chooce one of allowed keyHash mode\n")
		return errors.New("not supported keyHash mode")
	}
	return nil

}
