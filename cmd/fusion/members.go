package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/codegangsta/cli"
	"github.com/union-project/fusion"
)

var membersCommand = cli.Command{
	Name:  "members",
	Usage: "member operations",
	Subcommands: []cli.Command{
		listMembersCommand,
	},
}

var listMembersCommand = cli.Command{
	Name: "list",
	Aliases: []string{
		"ls",
	},
	Usage: "list members",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "order, o",
			Usage: "order by",
			Value: "first_name",
		},
	},
	Action: func(c *cli.Context) error {
		m, err := getManager(c)
		if err != nil {
			return err
		}

		members, err := m.Members(fusion.WithOrder(c.String("order")))
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
		fmt.Fprintln(w, "NAME\tTITLE\tSTATE\t")
		for _, member := range members {
			fmt.Fprintf(w, "%s %s\t%s\t%s\t\n", member.FirstName, member.LastName, member.Title, member.State)
		}
		w.Flush()

		return nil
	},
}
