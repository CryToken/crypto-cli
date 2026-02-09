/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/crytoken/crypto-cli/internal/tui"
	"github.com/spf13/cobra"
)

var isAdvanced bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if isAdvanced {
			fmt.Println("Advanced root")

			selectFunction(cmd, args)
		} else {
			selectFunction(cmd, args)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.crypto-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolVarP(&isAdvanced, "advance", "a", false, "set for advanced options")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func selectFunction(cmd *cobra.Command, args []string) {
	functions := []string{"Encrypt", "Decrypt", "Hash", "Sign", "Verify", "Checksum"}
	//f := utils.ChooseItem(functions)
	f := tui.ChoiceItem(functions)
	switch f {
	case "Encrypt":
		fmt.Println(args)
		EncryptCmd.Run(cmd, args)
	case "Decrypt":
		decryptCmd.Run(cmd, args)
	case "Hash":
		hashCmd.Run(cmd, args)
	case "Sign":
		signCmd.Run(cmd, args)
	case "Verify":
		verifyCmd.Run(cmd, args)
	}

}
