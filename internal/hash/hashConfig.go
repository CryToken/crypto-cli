package hash

type HashConfig struct {
	Method    string
	Mode      string
	Data      string
	InputFile string
}

func InitHashCfg() *HashConfig {
	return &HashConfig{}
}
