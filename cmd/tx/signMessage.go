/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/spf13/cobra"
)

// SignCmd represents the signMessage command
var SignCmd = &cobra.Command{
	Use:   "sign-message",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		signedMessageHash, err := signMessage(
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
		)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err.Error(), gastParams.ColorReset)
			os.Exit(1)
		}
		fmt.Println("signed message:\n", signedMessageHash)
	},
}

func init() {
	// Flags and configuration settings.
	SignCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "message to sign")
	SignCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")

	SignCmd.MarkFlagRequired("data")
	SignCmd.MarkFlagRequired("private-key")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signMessageCmd.Flags().BoolP("toggle", "handleTraceTx", false, "Help message for toggle")
}
