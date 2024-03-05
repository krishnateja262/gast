/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/spf13/cobra"
)

// estimateGasCmd represents the TryEstimateGas command
var estimateGasCmd = &cobra.Command{
	Use:   "estimate-gas",
	Short: "Provides an estimate of the gas required to execute a given transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		estimatedTxGas := TryEstimateGas(gastParams.TxRpcUrlValue, gastParams.FromValue, gastParams.ToValue, gastParams.TxDataValue, gastParams.WeiValue)
		fmt.Printf("Estimated gas: %s%d%s\n", gastParams.ColorGreen, estimatedTxGas, gastParams.ColorReset)
	},
}

func init() {
	// Flags and configuration settings.
	estimateGasCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")
	estimateGasCmd.Flags().StringVarP(&gastParams.FromValue, "from", "f", "", "sender")
	estimateGasCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "recipient")
	estimateGasCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "data")
	estimateGasCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "wei")

	// Mark flags required
	estimateGasCmd.MarkFlagsRequiredTogether("url", "from", "to", "data", "wei")
}
