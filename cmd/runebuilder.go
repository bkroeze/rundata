package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bkroeze/go.utils"
	"github.com/bkroeze/runedata"

	"github.com/codegangsta/cli"
)

func MakeRuneTable(filename, outdir, table string) error {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	template := string(raw[:])

	fmt.Printf("index: %i", strings.Index(template, "<!-- runetable -->"))
	formatted := utils.InsertTextBetween("<!-- runetable -->", "<!-- /runetable -->", template, table)

	println(formatted)
	return nil
}

func RuneTableCmd(ctx *cli.Context) {
	filename := ctx.String("runes")
	println("Loading runes from: %s", filename)
	runes, err := runedata.RunesFromFile(filename, true)

	if err != nil {
		panic(err)
	}

	filename = ctx.String("file")
	files, err := filepath.Glob(filename)
	if err != nil {
		panic(err)
	}
	outdir := ctx.String("directory")

	runetable := runedata.RunesToMDTable(runes)

	fmt.Printf("Building Rune Tables in \"%s\" for\n", outdir)
	for i := 0; i < len(files); i++ {
		fmt.Printf("    %s\n", files[i])
		MakeRuneTable(files[i], outdir, runetable)
	}

}

func main() {
	app := cli.NewApp()
	app.Name = "runebuilder"
	app.Usage = "Inserts rune data into files"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		println("Please run with the \"-h\" flag for help")
	}

	runeflag := cli.StringFlag{
		Name:  "runes, r",
		Value: "runes.csv",
		Usage: "Rune CSV Filename",
	}

	app.Commands = []cli.Command{
		{
			Name:  "table",
			Usage: "Insert a rune table",
			Flags: []cli.Flag{
				runeflag,
				cli.StringFlag{
					Name:  "file, f",
					Value: "templates/*",
					Usage: "file to use as the template",
				},
				cli.StringFlag{
					Name:  "directory, d",
					Value: "out",
					Usage: "Output directory",
				},
			},
			Action: RuneTableCmd,
		}, {
			Name:  "detail",
			Usage: "Write rune details",
			Action: func(c *cli.Context) {
				println("writing details into: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
