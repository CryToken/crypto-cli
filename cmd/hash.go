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
	Short: "Hash data",
	Long: `You can hash file -f flag or text -t
	Choose Alogrithm with -a flag ("default:SHA-256")
	Also you may dont provide any flags to start interactive TUI menu.`,
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
	hashCmd.Flags().BoolVarP(&cfg.IsAdnvanced, "advance", "", false, "Advance mod for.")

}
