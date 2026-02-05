package hash

import (
	"errors"
	"strings"

	"github.com/crytoken/crypto-cli/internal/tui"
)

func parseHashCfg(cfg *HashConfig) error {

	if cfg.Data == "" && cfg.InputFile == "" {
		if err := tui.SelectFile(&cfg.InputFile); err != nil {
			return err
		}

	}
	if cfg.Data != "" && cfg.InputFile != "" {
		return errors.New("only one can be choicen -t or -f")
	}

	if cfg.IsAdnvanced {

		choice := tui.ChoiceItem(supportedAlgo)
		if choice != "" {
			cfg.Method = choice
		}
	}
	if cfg.Data != "" {
		cfg.Mode = "Text"
	}
	if cfg.InputFile != "" {
		cfg.Mode = "File"
	}
	cfg.Method = strings.ToUpper(cfg.Method)
	return nil
}
