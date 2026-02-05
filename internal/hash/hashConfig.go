package hash

type HashConfig struct {
	Method      string
	Mode        string
	Data        string
	InputFile   string
	IsAdnvanced bool
}

func InitHashCfg() *HashConfig {
	return &HashConfig{}
}

var supportedAlgo []string = []string{"SHA1", "SHA256", "SHA512", "SHA3-256", "SHA3-512", "SHA4"}
