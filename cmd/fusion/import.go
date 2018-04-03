package main

import (
	"fmt"
	"runtime"

	"github.com/codegangsta/cli"
)

var importCommand = cli.Command{
	Name:   "import",
	Usage:  "import data archive",
	Action: importAction,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "workers, w",
			Usage: "set number of workers",
			Value: runtime.NumCPU(),
		},
	},
}

func importAction(c *cli.Context) error {
	m, err := getManager(c)
	if err != nil {
		return err
	}

	archivePath := c.Args().First()
	if archivePath == "" {
		return fmt.Errorf("you must specify an archive path")
	}

	if err := m.ImportBillArchive(archivePath, c.Int("workers")); err != nil {
		return err
	}

	return nil
}
