package main

import (
	"github.com/codegangsta/cli"
)

var syncCommand = cli.Command{
	Name:   "sync",
	Usage:  "start data sync",
	Action: syncAction,
}

func syncAction(c *cli.Context) error {
	m, err := getManager(c)
	if err != nil {
		return err
	}

	if err := m.Sync(); err != nil {
		return err
	}

	return nil
}
