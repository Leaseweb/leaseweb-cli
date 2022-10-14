package cmd

import (
	"strings"

	LSW "github.com/LeaseWeb/leaseweb-go-sdk"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func init() {
	customerCmd.AddCommand(customerContactListCmd)
	customerCmd.AddCommand(customerGetCmd)
	rootCmd.AddCommand(customerCmd)
}

var customerCmd = &cobra.Command{
	Use:   "customer",
	Short: "Get information about your Account and manage your Contacts",
	Long:  "Get information about your Account and manage your Contacts",
}

var customerContactListCmd = &cobra.Command{
	Use:   "list-contacts",
	Short: "Retrieve the list of Contacts",
	Long:  "Retrieve the list of Contacts",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.CustomerAccountApi{}.ListContacts()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Email", "First name", "Last name", "Roles", "Primary roles")
			for i, contact := range result.Contacts {
				t.AddLine(i+1, contact.Id, contact.Email, contact.FirstName, contact.LastName, strings.Join(contact.Roles[:], ","), strings.Join(contact.PrimaryRoles[:], ","))
			}
			t.Print()
		}
	},
}

var customerGetCmd = &cobra.Command{
	Use:   "me",
	Short: "Retrieve details of logged in Account",
	Long:  "Retrieve details of logged in Account",
	Run: func(cmd *cobra.Command, args []string) {
		customerAccount, err := LSW.CustomerAccountApi{}.Get()
		if err == nil {
			t := tabby.New()
			t.AddLine("Name:", customerAccount.Name)
			t.AddLine("Vat number:", customerAccount.VatNumber)
			t.AddLine("City:", customerAccount.Address.City)
			t.AddLine("Country:", customerAccount.Address.Country)
			t.AddLine("PostalCode:", customerAccount.Address.PostalCode)
			t.AddLine("State:", customerAccount.Address.State)
			t.AddLine("StateCode:", customerAccount.Address.StateCode)
			t.AddLine("HouseNumber:", customerAccount.Address.HouseNumber)
			t.AddLine("Street:", customerAccount.Address.Street)
			t.Print()
		}
	},
}
