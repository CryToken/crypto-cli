package encrypt

type Config struct {
	Method     string
	MethodMode string
	InputFile  string
	OutputFile string
	KeyMode    string
	Key        string
	KeyHash    []byte
}

func InitCfg() *Config {
	return &Config{
		Method:     "AES",
		MethodMode: "CFB",
		KeyMode:    "SHA256",
		OutputFile: "",
	}
}
