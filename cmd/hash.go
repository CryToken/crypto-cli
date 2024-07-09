/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/crytoken/crypto-cli/internal/hash"
	"github.com/spf13/cobra"
)

var cfg *hash.HashConfig

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		hash.Run(cfg)
	},
}

func init() {
	cfg = hash.InitHashCfg()
	rootCmd.AddCommand(hashCmd)

	hashCmd.Flags().StringVarP(&cfg.Method, "algo", "a", "SHA256", "Choice hash Algorithm")
	hashCmd.Flags().StringVarP(&cfg.Data, "text", "t", "", "Input text for hashing")
	hashCmd.Flags().StringVarP(&cfg.InputFile, "file", "f", "", "Choose file for hash")

}
