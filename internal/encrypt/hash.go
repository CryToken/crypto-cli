package encrypt

import "crypto/sha256"

func sha256hash(input string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	result := hasher.Sum(nil)
	return result
}
