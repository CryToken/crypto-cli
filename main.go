package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/crytoken/crypto-cli/cmd"
)

func main() {

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		fmt.Println("\nInterrupted.")
		os.Exit(0)
	}()
	cmd.Execute()
}
