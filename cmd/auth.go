package cmd

import (
	"fmt"
	"os"
	"syscall"

	LSW "github.com/LeaseWeb/leaseweb-go-sdk"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the leaseweb account",
	Long:  "Log in to the leaseweb account",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Print("Enter api key: \n")
		apiKey, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(0)
		}
		WriteApiKey(string(apiKey))
		Login()
		_, err = LSW.CustomerAccountApi{}.Get()
		if err != nil {
			fmt.Println("Faild to login!")
			os.Exit(0)
		}
		fmt.Println("Logged in successfully!")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from the leaseweb account",
	Long:  "Log out from the leaseweb account",
	Run: func(cmd *cobra.Command, args []string) {
		Logout()
		fmt.Println("Logged out successfully!")
	},
}
