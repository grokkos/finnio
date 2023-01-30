package cmd

import (
	"fmt"
	"github.com/grokkos/finnio/internal/app"

	"github.com/spf13/cobra"
)

// sharesCmd represents the shares command
var sharesCmd = &cobra.Command{
	Use:   "shares",
	Short: "Use to retrieve portfolio data",
	Long:  `Current price, Previous closing Profit and Loss calculation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shares called")
	},
}

func init() {
	rootCmd.AddCommand(sharesCmd)
	app.App()
	//TODO work on flags
}
