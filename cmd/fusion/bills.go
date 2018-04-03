package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/union-project/fusion"
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
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "limit, l",
			Usage: "limit number of results",
			Value: 0,
		},
		cli.StringFlag{
			Name:  "order, o",
			Usage: "order by",
			Value: "updated_at DESC",
		},
	},
	Action: func(c *cli.Context) error {
		m, err := getManager(c)
		if err != nil {
			return err
		}

		opts := []fusion.QueryOpt{}
		limit := c.Int("limit")
		if limit > 0 {
			opts = append(opts, fusion.WithLimit(limit))
		}
		opts = append(opts, fusion.WithOrder(c.String("order")))

		bills, err := m.Bills(opts...)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
		fmt.Fprintln(w, "TITLE\tID\tTYPE\tINTRODUCED\tUPDATED\t")
		for _, bill := range bills {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", truncate(bill.Title, 80), bill.BillID, bill.BillType, bill.IntroducedAt, bill.UpdatedAt)
		}
		w.Flush()

		return nil
	},
}

func truncate(str string, num int) string {
	b := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		b = str[0:num] + "..."
	}
	return b
}
