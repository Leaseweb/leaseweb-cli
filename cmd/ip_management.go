package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
	"github.com/spf13/cobra"
)

func init() {
	ipManagementCmd.AddCommand(ipManagementlistCmd)
	ipManagementCmd.AddCommand(ipManagementGetCmd)
	rootCmd.AddCommand(ipManagementCmd)
}

var ipManagementCmd = &cobra.Command{
	Use:   "ip-management",
	Short: "Get information about your assigned IPs and manage them",
	Long:  "Get information about your assigned IPs and manage them",
}

var ipManagementlistCmd = &cobra.Command{
	Use:   "list",
	Short: "List your Ips",
	Long:  "List your Ips",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.IpManagementApi{}.List()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Ip", "Subnet Gateway", "Subnet Id", "Primary", "Version", "Null routed")
			for i, ip := range result.Ips {
				t.AddLine(i+1, ip.Ip, ip.Subnet.Gateway, ip.Subnet.Id, ip.Primary, ip.Version, ip.NullRouted)
			}
			t.Print()
		}
	},
}

var ipManagementGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about a single Ip",
	Long:  "Get information about a single Ip",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := LSW.IpManagementApi{}.Get(args[0])
		if err != nil {
			fmt.Println(err)
		}

		t := tabby.New()
		t.AddLine("Ip:", ip.Ip)
		t.AddLine("Floating Ip:", ip.FloatingIp)
		t.AddLine("Main Ip:", ip.MainIp)
		t.AddLine("Reverse lookup:", ip.ReverseLookup)
		t.AddLine("Type:", ip.Type)
		t.AddLine("Prefix length:", ip.PrefixLength.String())
		t.AddLine("Main Ip:", ip.MainIp)
		t.AddLine("Gateway:", ip.Gateway)
		t.AddLine("Network type:", ip.NetworkType)
		t.AddLine("Primary:", ip.Primary)
		t.AddLine("Version:", ip.Version)
		t.AddLine("Null routed:", ip.NullRouted)
		t.AddLine("Subnet id:", ip.Subnet.Id)
		t.AddLine("Subnet gateway:", ip.Subnet.Gateway)
		t.AddLine("Subnet network ip:", ip.Subnet.NetworkIp)
		t.AddLine("Subnet prefix length:", ip.Subnet.PrefixLength)
		t.Print()
	},
}
