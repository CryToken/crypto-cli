package verify

type VeryfiConfig struct {
	Signature string
	PublicKey string
	Data      string
	Algorithm string
	HashAlgo  string
}

func InitConfig() *VeryfiConfig {
	return &VeryfiConfig{
		Algorithm: "ECDSA",
	}
}
