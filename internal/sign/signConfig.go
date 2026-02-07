package sign

import (
	"crypto/elliptic"
	"io"
)

var (
	ecdsaStr string = "ECDSA"
	rsaStr   string = "RSA"
	dsaStr   string = "DSA"
)
var (
	sha256Str string = "SHA-256"
	sha512Str string = "SHA-512"
	argon2Str string = "argon2"
)

type SignConfig struct {
	Algorithm    string
	KeyFile      string
	IsAdnvanced  bool
	Input        string
	HashAlgo     string
	Curve        elliptic.Curve
	Output       string
	OutputWriter io.Writer
}

func InitSignConfig() *SignConfig {
	return &SignConfig{
		HashAlgo: "SHA-256",
		Curve:    elliptic.P256(),
	}
}

var supportedSignAlogos []string = []string{ecdsaStr, rsaStr, dsaStr}
var supportedHashAlogos []string = []string{sha256Str, sha512Str, argon2Str}
