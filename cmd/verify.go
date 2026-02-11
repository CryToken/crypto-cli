/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/crytoken/crypto-cli/internal/verify"
	"github.com/spf13/cobra"
)

var verifyConfig *verify.VeryfiConfig

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		verify.Run(verifyConfig)

	},
}

func init() {
	verifyConfig = verify.InitConfig()
	rootCmd.AddCommand(verifyCmd)

	verifyCmd.Flags().StringVarP(&verifyConfig.Algorithm, "algorithm", "a", "", "Choose alogorithm to verify.")
	verifyCmd.Flags().StringVarP(&verifyConfig.Data, "file", "f", "", "choose file.")
	verifyCmd.Flags().StringVarP(&verifyConfig.PublicKey, "key", "k", "", "Choose public key file.")
	verifyCmd.Flags().StringVarP(&verifyConfig.Signature, "signature", "s", "", "Choose signature file.")
	verifyCmd.Flags().StringVar(&verifyConfig.HashAlgo, "hash", "SHA-256", "select data hash algorithm.")
}
