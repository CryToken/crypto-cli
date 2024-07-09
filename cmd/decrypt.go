package cmd

import (
	"github.com/crytoken/crypto-cli/internal/decrypt"
	"github.com/spf13/cobra"
)

var decryptCfg *decrypt.Config

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt file",
	Long:  `You can decrypt encrypted files just use --file and --key.`,
	Run: func(cmd *cobra.Command, args []string) {
		decrypt.Run(decryptCfg)
	},
}

func init() {
	//Init Decrypt confing .
	decryptCfg = decrypt.InitCfg()
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	decryptCmd.Flags().StringVarP(&decryptCfg.Method, "algo", "a", "AES", "CHhooce algorithm to encrypt File")
	decryptCmd.Flags().StringVarP(&decryptCfg.MethodMode, "Mode", "M", "CFB", "Mode for Algorithm")
	decryptCmd.Flags().StringVarP(&decryptCfg.Key, "key", "k", "", "Key for cipher")
	decryptCmd.Flags().StringVarP(&decryptCfg.OutputFile, "out", "o", "", "path to output file")
	decryptCmd.Flags().StringVarP(&decryptCfg.InputFile, "file", "f", "", "File to encrypt")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
