package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	appName = "license"

	nameFlag = "name"
)

var fmtErr = color.New(color.FgRed)

func main() {
	a := action{}

	app := &cli.App{
		Name:  appName,
		Usage: "generate LICENSE",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    nameFlag,
				Aliases: []string{"n"},
				Usage:   "which `LICENSE`",
			},
		},
		Action: a.Run,
	}

	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr.Printf("[%s error]: ", appName)
		fmt.Printf("%s\n", err.Error())
	}
}

type action struct {
	flags struct {
		name string
	}
}

func (a *action) Run(c *cli.Context) error {
	// Show help if there is nothing
	if c.NArg() == 0 && c.NumFlags() == 0 {
		return cli.ShowAppHelp(c)
	}

	a.getFlags(c)

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.name = c.String(nameFlag)
}
