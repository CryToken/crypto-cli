package cmd

import (
	"github.com/crytoken/crypto-cli/internal/genkey"
	"github.com/spf13/cobra"
)

var genkeyCfg *genkey.GenkeyConfig

// genKeyCmd represents the genkey command
var genKeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "genkey key pairs",
	Long:  `You can decrypt encrypted files just use --file and --key.`,
	Run: func(cmd *cobra.Command, args []string) {

		genkey.Run(genkeyCfg, args)
	},
}

func init() {
	genkeyCfg = genkey.InitConfig()
	rootCmd.AddCommand(genKeyCmd)

	genKeyCmd.Flags().StringVarP(&genkeyCfg.Type, "type", "t", "", "app genkey -t [type]")
	genKeyCmd.Flags().StringVarP(&genkeyCfg.Output, "out", "o", "", "set out file name -o [filename]")
}
