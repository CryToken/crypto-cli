package utils

import (
	"fmt"
	"os"
	"strings"
)

func SetOutName(path string) string {
	var output string
	if strings.HasSuffix(path, "_enc") {
		output = strings.TrimSuffix(path, "_enc")
		_, err := os.Stat(output)
		if os.IsNotExist(err) {
			return output
		}
		for i := 1; i < 111; i++ {
			output = fmt.Sprintf("%s_%d", output, i)
			_, err := os.Stat(output)
			if os.IsNotExist(err) {
				return output
			}
		}
	}

	res := fmt.Sprintf("%s_decr", path)
	return res
}
