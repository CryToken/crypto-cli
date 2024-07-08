package utils

import "crypto/sha256"

func Sha256hash(input string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	result := hasher.Sum(nil)
	return result
}
