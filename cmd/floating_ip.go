package cmd

import (
	"fmt"

	"github.com/cheynewallace/tabby"
	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
	"github.com/spf13/cobra"
)

func init() {
	floatingIpCmd.AddCommand(floatingIplistCmd)
	floatingIpCmd.AddCommand(floatingIpGetCmd)
	rootCmd.AddCommand(floatingIpCmd)
}

var floatingIpCmd = &cobra.Command{
	Use:   "floating-ip",
	Short: "Manage your Floating IPs",
	Long:  "Manage your Floating IPs",
}

var floatingIplistCmd = &cobra.Command{
	Use:   "list",
	Short: "A Floating IP range is bound to a particular site or metro.",
	Long:  "A Floating IP range is bound to a particular site or metro.",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.FloatingIpApi{}.ListRanges()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Range", "Location", "Type")
			for i, ip := range result.Ranges {
				t.AddLine(i+1, ip.Id, ip.Range, ip.Location, ip.Type)
			}
			t.Print()
		}
	},
}

var floatingIpGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about a single server",
	Long:  "Get information about a single server",
	Run: func(cmd *cobra.Command, args []string) {
		ip, err := LSW.FloatingIpApi{}.GetRange(args[0])
		if err != nil {
			fmt.Println(err)
		}

		t := tabby.New()
		t.AddLine("Id:", ip.Id)
		t.AddLine("Range:", ip.Range)
		t.AddLine("Location:", ip.Location)
		t.AddLine("Type:", ip.Type)
		t.AddLine("Customer Id:", ip.CustomerId)
		t.AddLine("SalesOrgId:", ip.SalesOrgId)
		t.Print()
	},
}
