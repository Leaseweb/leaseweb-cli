package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of leaseweb cli",
	Long:  "Print the version number of leaseweb cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("leaseweb version 0.0.1")
	},
}
