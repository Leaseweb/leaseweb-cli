package cmd

import (
	"fmt"

	LSW "github.com/LeaseWeb/leaseweb-go-sdk"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func init() {
	serviceCmd.AddCommand(servicelistCmd)
	serviceCmd.AddCommand(serviceGetCmd)
	rootCmd.AddCommand(serviceCmd)
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Get information about your Services and manage them",
	Long:  "Get information about your Services and manage them",
}

var servicelistCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve list of Services",
	Long:  "Retrieve list of Services",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.ServicesApi{}.List()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Product id", "Status", "Equipment id", "Contract id", "Reference")
			for i, service := range result.Services {
				t.AddLine(i+1, service.Id, service.ProductId, service.Status, service.EquipmentId, service.ContractId, service.Reference)
			}
			t.Print()
		}
	},
}

var serviceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details of a Service",
	Long:  "Retrieve details of a Service",
	Run: func(cmd *cobra.Command, args []string) {
		service, err := LSW.ServicesApi{}.Get(args[0])
		if err != nil {
			fmt.Println(err)
		}

		t := tabby.New()
		t.AddLine("Id:", service.Id)
		t.AddLine("Product id:", service.ProductId)
		t.AddLine("Equipment id:", service.EquipmentId)
		t.AddLine("Contract id:", service.ContractId)
		t.AddLine("Status:", service.Status)
		t.AddLine("Reference:", service.Reference)
		t.AddLine("Billing cycle:", service.BillingCycle)
		t.AddLine("Cancellable:", service.Cancellable)
		t.AddLine("Contract term:", service.ContractTerm)
		t.AddLine("Contract term end date:", service.ContractTermEndDate)
		t.AddLine("Currency:", service.Currency)
		t.AddLine("Delivery date:", service.DeliveryDate)
		t.AddLine("Delivery estimate:", service.DeliveryEstimate)
		t.AddLine("End date:", service.EndDate)
		t.AddLine("Order date:", service.OrderDate)
		t.AddLine("Price per frequency:", service.PricePerFrequency)
		t.AddLine("Start date:", service.StartDate)
		t.AddLine("Status:", service.Status)
		t.AddLine("Uncancellable:", service.Uncancellable)
		t.Print()
	},
}
