package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	pwd, _ := os.Getwd()
	pid := os.Getpid()
	fmt.Println(pwd, pid)

	pwd = parentDir(pwd)
	fmt.Println(pwd)

	path := "/home/fedor/golang/crypto-cli/ftoenc_enc"
	out := setOutName(path)
	fmt.Println(out)

}
func parentDir(path string) string {
	parent := filepath.Dir(filepath.Clean(path))
	if parent == "." {
		return "/"
	}
	return parent
}

func setOutName(path string) string {
	var output string
	if strings.HasSuffix(path, "_enc") {
		output = strings.TrimSuffix(path, "_enc")
		_, err := os.Stat(output)
		if os.IsNotExist(err) {
			return output
		}
		for i := 1; i < 111; i++ {
			output = fmt.Sprintf("%s(%d)", output, i)
			_, err := os.Stat(output)
			if os.IsNotExist(err) {
				return output
			}
		}
	}

	return path
}
