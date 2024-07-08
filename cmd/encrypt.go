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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		encrypt.Run(config)
	},
}

func init() {
	// Initialize the config
	config = encrypt.InitCfg()

	// Add the encrypt command to the root command
	rootCmd.AddCommand(encryptCmd)

	// Add flags to the encrypt command
	encryptCmd.Flags().StringVarP(&config.Key, "key", "k", "", "Key for cipher")
	encryptCmd.Flags().StringVarP(&config.InputFile, "file", "f", "", "File to encrypt")
}
