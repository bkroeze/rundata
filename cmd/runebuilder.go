package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func RuneTableCmd(ctx *cli.Context) {
	println("inserting runes into: ", c.Args().First())
}

func main() {
	app := cli.NewApp()
	app.Name = "runebuilder"
	app.Usage = "Inserts rune data into files"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "runes, r",
			Value: "runes.csv",
			Usage: "Rune CSV Filename",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "table",
			Usage:  "Insert a rune table",
			Action: RuneTableCmd,
		},
		{
			Name:  "detail",
			Usage: "Write rune details",
			Action: func(c *cli.Context) {
				println("writing details into: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
