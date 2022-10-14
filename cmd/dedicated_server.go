package cmd

import (
	"fmt"

	LSW "github.com/LeaseWeb/leaseweb-go-sdk"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func init() {
	dedicatedServerCmd.AddCommand(dedicatedServerlistCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerGetCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOnCmd)
	dedicatedServerCmd.AddCommand(dedicatedServerPowerOffCmd)
	rootCmd.AddCommand(dedicatedServerCmd)
}

var dedicatedServerCmd = &cobra.Command{
	Use:   "dedicated-server",
	Short: "Get information about your Dedicated Servers and manage them",
	Long:  "Get information about your Dedicated Servers and manage them",
}

var dedicatedServerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve the list of Dedicated Servers",
	Long:  "Retrieve the list of Dedicated Servers",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.DedicatedServerApi{}.List()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Asset id", "Rack", "Site", "Suite", "Unit", "Rack type")
			for i, server := range result.Servers {
				t.AddLine(i+1, server.Id, server.AssetId, server.Location.Rack, server.Location.Site, server.Location.Suite, server.Location.Unit, server.Rack.Type)
			}
			t.Print()
		}
	},
}

var dedicatedServerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details of Dedicated Server",
	Long:  "Retrieve details of Dedicated Server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := LSW.DedicatedServerApi{}.Get(args[0])
		if err != nil {
			fmt.Println(err)
		}

		t := tabby.New()
		t.AddLine("Id:", server.Id)
		t.AddLine("Asset Id:", server.AssetId)
		t.AddLine("Location (Rack):", server.Location.Rack)
		t.AddLine("Location (Site):", server.Location.Site)
		t.AddLine("Location (Suite):", server.Location.Suite)
		t.AddLine("Location (Unit):", server.Location.Unit)
		t.AddLine("Rack Type:", server.Rack.Type)
		t.Print()
	},
}

var dedicatedServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-on a Dedicated Server",
	Long:  "Power-on a Dedicated Server",
	Run: func(cmd *cobra.Command, args []string) {
		err := LSW.DedicatedServerApi{}.PowerOnServer(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Server powered on!")
		}
	},
}

var dedicatedServerPowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power-off a Dedicated Server",
	Long:  "Power-off a Dedicated Server",
	Run: func(cmd *cobra.Command, args []string) {
		err := LSW.DedicatedServerApi{}.PowerOffServer(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Server powered off!")
		}
	},
}
