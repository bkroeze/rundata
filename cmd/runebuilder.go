package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bkroeze/go.utils"
	"github.com/bkroeze/runedata"

	"github.com/codegangsta/cli"
)

func MakeRuneFile(filename, outdir, table string) (string, error) {
	raw, err := ioutil.ReadFile(filename)

	if err != nil {
		return table, err
	}

	template := string(raw[:])

	formatted := template[:]

	if strings.Index(template, "<!-- runetable -->") > -1 {
		formatted = utils.InsertTextBetween("<!-- runetable -->", "<!-- /runetable -->", template, table)
	}

	return formatted, nil
}

func WriteOutfile(outdir, fname, text string) error {
	outfile := filepath.Join(outdir, filepath.Base(fname))
	fmt.Printf("    Writing: %s\n", outfile)
	err := ioutil.WriteFile(outfile, []byte(text), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func RuneBuildCmd(ctx *cli.Context) {
	filename := ctx.String("file")
	fmt.Printf("Loading runes from: %s", filename)
	runes, err := runedata.RunesFromFile(filename, true)

	if err != nil {
		panic(err)
	}

	files, err := filepath.Glob(filename)
	if err != nil {
		panic(err)
	}

	outdir := ctx.String("directory")
	if len(files) > 0 {
		err := os.MkdirAll(outdir, 0644) // o+rw,a+r
		if err != nil {
			panic(err)
		}
	}

	runetable := runedata.RunesToMDTable(runes)

	fmt.Printf("Building Rune Files in \"%s\" for\n", outdir)
	for i := 0; i < len(files); i++ {
		fmt.Printf("    %s\n", files[i])
		formatted, err := MakeRuneFile(files[i], outdir, runetable)
		if err != nil {
			panic(err)
		}
		WriteOutfile(outdir, files[i], formatted)
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
			Name:  "build",
			Usage: "Build files",
			Flags: []cli.Flag{
				runeflag,
				cli.StringFlag{
					Name:  "file, f",
					Value: "templates/*",
					Usage: "files to use as the template",
				},
				cli.StringFlag{
					Name:  "directory, d",
					Value: "out",
					Usage: "Output directory",
				},
			},
			Action: RuneBuildCmd,
		},
	}

	app.Run(os.Args)
}
