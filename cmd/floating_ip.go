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
	Short: "Get information about your Floating Ips and manage them",
	Long:  "Get information about your Floating Ips and manage them",
}

var floatingIplistCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve the list of Floating Ips",
	Long:  "Retrieve the list of Floating Ips",
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
	Short: "Retrieve details of Floating Ip",
	Long:  "Retrieve details of Floating Ip",
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
