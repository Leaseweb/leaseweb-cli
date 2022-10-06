package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "leaseweb",
	Short: "leaseweb-cli is a tool for managing resources",
	Long:  "leaseweb-cli is a tool for managing resources",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
