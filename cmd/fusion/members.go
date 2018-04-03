package main

import (
	"fmt"

	"github.com/codegangsta/cli"
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
	Action: func(c *cli.Context) error {
		m, err := getManager(c)
		if err != nil {
			return err
		}

		members, err := m.Members()
		if err != nil {
			return err
		}

		fmt.Println(members)

		return nil
	},
}
