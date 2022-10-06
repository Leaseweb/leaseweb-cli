package cmd

import (
	"fmt"
	"os"
	"syscall"

	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func init() {
	rootCmd.AddCommand(authCmd)
}

var authCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate your API call",
	Long:  "Authenticate your API call",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Enter api key: \n")
		apiKey, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(0)
		}
		WriteApiKey(string(apiKey))
		Authenticate()
		_, err = LSW.CustomerAccountApi{}.Get()
		if err != nil {
			fmt.Println("Faild to authenticate api key!")
			os.Exit(0)
		}
		fmt.Println("Authenticated")
	},
}
