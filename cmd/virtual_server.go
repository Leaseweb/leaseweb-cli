package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
	"github.com/spf13/cobra"
)

func init() {
	virtualServerCmd.AddCommand(virtualServerlistCmd)
	virtualServerCmd.AddCommand(virtualServerGetCmd)
	virtualServerCmd.AddCommand(virtualServerPowerOnCmd)
	virtualServerCmd.AddCommand(virtualServerPowerOffCmd)
	rootCmd.AddCommand(virtualServerCmd)
}

var virtualServerCmd = &cobra.Command{
	Use:   "virtual-server",
	Short: "Control your virtual servers",
	Long:  "Control your virtual servers",
}

var virtualServerlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List your virtual servers",
	Long:  "List your virtual servers",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.VirtualServerApi{}.List()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Datacenter", "State", "Template", "Reference", "Firewall state")
			for i, server := range result.VirtualServers {
				t.AddLine(i+1, server.Id, server.DataCenter, server.State, server.Template, server.Reference, server.FirewallState)
			}
			t.Print()
		}
	},
}

var virtualServerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about a single virtual server",
	Long:  "Get information about a single virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := LSW.VirtualServerApi{}.Get(args[0])
		if err != nil {
			fmt.Println(err)
		}

		t := tabby.New()
		t.AddLine("Id:", server.Id)
		t.AddLine("Datacenter:", server.DataCenter)
		t.AddLine("State:", server.State)
		t.AddLine("Template:", server.Template)
		t.AddLine("Reference:", server.Reference)
		t.AddLine("Firewall state:", server.FirewallState)
		t.Print()
	},
}

var virtualServerPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power on the given virtual server",
	Long:  "Power on the given virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := LSW.VirtualServerApi{}.PowerOff(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Virtual server powered on!")
		}
	},
}

var virtualServerPowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power off the given virtual server",
	Long:  "Power off the given virtual server",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := LSW.VirtualServerApi{}.PowerOn(args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Virtual server powered off!")
		}
	},
}
