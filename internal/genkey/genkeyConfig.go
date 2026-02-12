package genkey

type GenkeyConfig struct {
	Type   string
	Output string
}

func InitConfig() *GenkeyConfig {
	return &GenkeyConfig{}
}
