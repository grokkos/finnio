package cmd

import (
	"fmt"
	"github.com/grokkos/finnio/internal/app"

	"github.com/spf13/cobra"
)

// sharesCmd represents the shares command
var sharesCmd = &cobra.Command{
	Use:   "shares",
	Short: "Shares",
	Long:  `Shares.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shares called")
	},
}

func init() {
	rootCmd.AddCommand(sharesCmd)
	app.App()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sharesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sharesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
