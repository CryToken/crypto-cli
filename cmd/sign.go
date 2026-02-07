/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/crytoken/crypto-cli/internal/sign"
	"github.com/spf13/cobra"
)

var signCfg *sign.SignConfig

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign your data",
	Long: `You can sign your data (-f flag) or Message (-m flag) with provided private key (-k flag).

	Provided private key should be in .pem format and contain private key for one of supported alogrithms

	Supported Alogorithms: ECDSA,
	`,
	Run: func(cmd *cobra.Command, args []string) {

		sign.Run(signCfg)
	},
}

func init() {
	signCfg = sign.InitSignConfig()

	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringVarP(&signCfg.Algorithm, "algorithm", "a", "ECDSA", "you can sign data with provided key(-k flag)")
	signCmd.Flags().StringVarP(&signCfg.KeyFile, "key", "k", "", "use keyFile for sign data")
	signCmd.Flags().StringVarP(&signCfg.Input, "input", "f", "", "choose file to sign")
	signCmd.Flags().StringVarP(&signCfg.Output, "out", "o", "signature", "provide filename to write signature")

}
