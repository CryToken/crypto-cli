package cmd

import (
	"github.com/crytoken/crypto-cli/internal/encrypt"
	"github.com/spf13/cobra"
)

// config variable should be declared but initialized in the init function
var config *encrypt.Config

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt file",
	Long: `You can encrypt file with some algorithms:
	To choice algo use flag -a or --algo.
	To set pass use flag -k or --key,
	To choice a file to encrypt use flag -f or --file`,
	Run: func(cmd *cobra.Command, args []string) {
		encrypt.Run(config, args)
	},
}

func init() {
	// Initialize the config
	config = encrypt.InitCfg()

	// Add the encrypt command to the root command
	rootCmd.AddCommand(encryptCmd)

	// Add flags to the encrypt command
	encryptCmd.Flags().StringVarP(&config.Method, "algo", "a", "AES", "CHhooce algorithm to encrypt File")
	encryptCmd.Flags().StringVarP(&config.MethodMode, "Mode", "M", "CFB", "Mode for Algorithm")
	encryptCmd.Flags().StringVarP(&config.Key, "key", "k", "", "Key for cipher")
	encryptCmd.Flags().StringVarP(&config.KeyMode, "keyHash", "K", "SHA256", "Key mode ,defauld its hash of key")
	encryptCmd.Flags().StringVarP(&config.OutputFile, "out", "o", "", "name for output file")
	encryptCmd.Flags().StringVarP(&config.InputFile, "file", "f", "", "File to encrypt")
}
