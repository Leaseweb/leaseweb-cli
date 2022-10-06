package cmd

import (
	"fmt"
	"os"

	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
)

var apiKeyPath string = "/.lsw"

func Authenticate() {
	LSW.InitLeasewebClient(ReadApiKey())
}

func ReadApiKey() string {

	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(0)
	}

	apiKey, err := os.ReadFile(dirname + apiKeyPath)
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(0)
	}

	return string(apiKey)
}

func WriteApiKey(apiKey string) {

	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(0)
	}

	// TODO: encrypt the api key
	err = os.WriteFile(dirname+apiKeyPath, []byte(apiKey), 0644)
	if err != nil {
		fmt.Println("Something went wrong")
		os.Exit(0)
	}
}
