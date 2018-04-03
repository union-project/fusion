package main

import (
	"github.com/codegangsta/cli"
	"github.com/union-project/fusion/manager"
)

func getManager(c *cli.Context) (*manager.Manager, error) {
	cfg := &manager.Config{
		DBAddr:           c.GlobalString("db-addr"),
		ProPublicaAPIKey: c.GlobalString("propublica-api-key"),
	}

	return manager.NewManager(cfg)
}
