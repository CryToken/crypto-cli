package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display current version of the app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("crypto-cli: v1.2.3")

	},
}

func init() {

	rootCmd.AddCommand(versionCmd)

}
