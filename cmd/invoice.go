package cmd

import (
	"time"

	"github.com/cheynewallace/tabby"
	LSW "github.com/majidkarimizadeh/leaseweb-go-sdk"
	"github.com/spf13/cobra"
)

func init() {
	invoiceCmd.AddCommand(invoiceListCmd)
	invoiceCmd.AddCommand(invoiceGetCmd)
	rootCmd.AddCommand(invoiceCmd)
}

var invoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Get your invoice data",
	Long:  "Get your invoice data",
}

var invoiceListCmd = &cobra.Command{
	Use:   "list",
	Short: "An overview of all the invoices",
	Long:  "An overview of all the invoices",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := LSW.InvoiceApi{}.List()
		if err == nil {
			t := tabby.New()
			t.AddHeader("#", "Id", "Total", "Currency", "Status", "Due date")
			for i, invoice := range result.Invoices {
				theTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", invoice.DueDate)
				t.AddLine(i+1, invoice.Id, invoice.Total.String(), invoice.Currency, invoice.Status, theTime.Format("2006-01-02" /*" 15:04"*/))
			}
			t.Print()
		}
	},
}

var invoiceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "An overview of single invoice",
	Long:  "An overview of single invoice",
	Run: func(cmd *cobra.Command, args []string) {
		invoice, err := LSW.InvoiceApi{}.Get(args[0])
		if err == nil {
			dueDate, _ := time.Parse("2006-01-02T15:04:05Z07:00", invoice.DueDate)
			date, _ := time.Parse("2006-01-02T15:04:05Z07:00", invoice.Date)

			t := tabby.New()
			t.AddLine("Id:", invoice.Id)
			t.AddLine("Total:", invoice.Total.String())
			t.AddLine("Open amount:", invoice.OpenAmount.String())
			t.AddLine("Tax amount:", invoice.TaxAmount.String())
			t.AddLine("Currency:", invoice.Currency)
			t.AddLine("Status:", invoice.Status)
			t.AddLine("Due date:", dueDate.Format("2006-01-02"))
			t.AddLine("Date:", date.Format("2006-01-02"))
			t.AddLine("Is partial payment allowed:", invoice.IsPartialPaymentAllowed)
			// t.AddLine("Credits:", len(invoice.Credits))
			// t.AddLine("Lines:", len(invoice.Lines))
			t.Print()
		}
	},
}
