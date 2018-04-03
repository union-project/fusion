package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
)

var billsCommand = cli.Command{
	Name:  "bills",
	Usage: "bill operations",
	Subcommands: []cli.Command{
		listBillsCommand,
	},
}

var listBillsCommand = cli.Command{
	Name: "list",
	Aliases: []string{
		"ls",
	},
	Usage: "list bills",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		m, err := getManager(c)
		if err != nil {
			return err
		}

		bills, err := m.Bills()
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
		fmt.Fprintln(w, "TITLE\tID\tTYPE\t")
		for _, bill := range bills {
			fmt.Fprintf(w, "%s %s\t%s\t%s\t\n", bill.Title, bill.BillID, bill.BillType)
		}
		w.Flush()

		return nil
	},
}
