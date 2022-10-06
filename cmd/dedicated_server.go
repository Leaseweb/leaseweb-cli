package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
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
	Short: "Manage your dedicated servers",
	Long:  "Manage your dedicated servers",
}

var dedicatedServerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List your dedicated servers",
	Long:  "List your dedicated servers",
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
	Short: "Get information about a single server",
	Long:  "Get information about a single server",
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
	Short: "Power on the given server",
	Long:  "Power on the given server",
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
	Short: "Power off the given server",
	Long:  "Power off the given server",
	Run: func(cmd *cobra.Command, args []string) {
		err := LSW.DedicatedServerApi{}.PowerOffServer(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Server powered off!")
		}
	},
}
