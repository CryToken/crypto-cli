package hash

import "errors"

func parseHashCfg(cfg *HashConfig) error {
	if cfg.Data == "" && cfg.InputFile == "" {
		return errors.New("choose one mode -t or -f")
	}
	if cfg.Data != "" && cfg.InputFile != "" {
		return errors.New("only one can be choicen -t or -f")
	}
	if cfg.Data != "" {
		cfg.Mode = "Text"
	}
	if cfg.InputFile != "" {
		cfg.Mode = "File"
	}
	return nil
}
