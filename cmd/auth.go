package cmd

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	LSW "github.com/LeaseWeb/leaseweb-go-sdk"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var apiKeyPath string = "/.lsw"

func Login() {

	if !isFileExists(apiKeyPath) {
		writeFile(apiKeyPath, "")
	}

	LSW.InitLeasewebClient(readfile(apiKeyPath))
}

func Logout() {
	writeFile(apiKeyPath, "")
}

func getHomeDir() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Something went wrong! Home directory not found!")
		os.Exit(0)
	}
	return dirname
}

func writeFile(path, content string) {
	if os.WriteFile(getHomeDir()+path, []byte(content), 0644) != nil {
		fmt.Println("Something went wrong1")
		os.Exit(0)
	}
}

func readfile(path string) string {
	apiKey, err := os.ReadFile(getHomeDir() + path)
	if err != nil {
		fmt.Println("Something went wrong2")
		os.Exit(0)
	}
	return string(apiKey)
}

func isFileExists(path string) bool {
	_, err := os.Stat(getHomeDir() + path)

	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}

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
		writeFile(apiKeyPath, string(apiKey))
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
