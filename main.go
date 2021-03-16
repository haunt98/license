package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	appName = "license"

	outputFlag = "output"

	currentDir      = "."
	licenseFilename = "LICENSE"
)

var fmtErr = color.New(color.FgRed)

func main() {
	a := action{}

	app := &cli.App{
		Name:  appName,
		Usage: "generate LICENSE",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        outputFlag,
				Aliases:     []string{"o"},
				Usage:       "output directory",
				DefaultText: currentDir,
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
		output string
	}
}

func (a *action) Run(c *cli.Context) error {
	// Show help if there is nothing
	if c.NArg() == 0 && c.NumFlags() == 0 {
		return cli.ShowAppHelp(c)
	}

	a.getFlags(c)

	fmt.Printf("What LICENSE do you chose: ")
	licenseName := readStdin()

	license, err := generateLicense(licenseName)
	if err != nil {
		return fmt.Errorf("failed to generate license %s: %w", licenseName, err)
	}

	outputFile := filepath.Join(a.flags.output, licenseFilename)
	if err := ioutil.WriteFile(outputFile, []byte(license), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputFile, err)
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.output = c.String(outputFlag)
}
