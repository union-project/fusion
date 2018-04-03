package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/version"
)

func main() {
	app := cli.NewApp()
	app.Name = version.Name
	app.Usage = "data sync utility"
	app.Version = version.BuildVersion()
	app.Author = "@union-project"
	app.Email = ""
	app.Before = func(c *cli.Context) error {
		// enable debug
		if c.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Debug("debug enabled")
		}

		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "enable debug",
		},
		cli.StringFlag{
			Name:   "db-addr",
			Usage:  "Fusion DB Address",
			EnvVar: "FUSION_DB_ADDR",
		},
		cli.StringFlag{
			Name:   "propublica-api-key",
			Usage:  "ProPublica API Key",
			Value:  "",
			EnvVar: "FUSION_PROPUBLICA_API_KEY",
		},
	}
	app.Commands = []cli.Command{
		membersCommand,
		billsCommand,
		syncCommand,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
