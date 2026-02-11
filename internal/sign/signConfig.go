package sign

import (
	"crypto/elliptic"
	"io"
)

var (
	ecdsaStr   string = "ECDSA"
	rsaStr     string = "RSA"
	dsaStr     string = "DSA"
	ed25516Str string = "ED-25516"
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
		HashAlgo: sha256Str,
		Curve:    elliptic.P256(),
	}
}

var supportedSignAlogos []string = []string{ecdsaStr, rsaStr, dsaStr, ed25516Str}
var supportedHashAlogos []string = []string{sha256Str, sha512Str, argon2Str}
