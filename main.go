package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/make-go-great/color-go"
	"github.com/make-go-great/ioe-go"
	"github.com/urfave/cli/v2"
)

const (
	name  = "license"
	usage = "generate LICENSE quickly"

	commandGenerateName  = "generate"
	commandGenerateUsage = "generate LICENSE"

	flagOutputName  = "output"
	flagOutputUsage = "output directory"

	currentDir      = "."
	licenseFilename = "LICENSE"
)

var (
	// command aliases
	generateAliases = []string{"gen"}

	// flag aliases
	outputAliases = []string{"o"}
)

func main() {
	a := action{}

	app := &cli.App{
		Name:  name,
		Usage: usage,
		Commands: []*cli.Command{
			{
				Name:    commandGenerateName,
				Aliases: generateAliases,
				Usage:   commandGenerateUsage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        flagOutputName,
						Aliases:     outputAliases,
						Usage:       flagOutputUsage,
						DefaultText: currentDir,
					},
				},
				Action: a.RunGenerate,
			},
		},
		Action: a.RunHelp,
	}

	if err := app.Run(os.Args); err != nil {
		color.PrintAppError(name, err.Error())
	}
}

type action struct {
	flags struct {
		output string
	}
}

func (a *action) RunHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) RunGenerate(c *cli.Context) error {
	a.getFlags(c)

	fmt.Printf("What LICENSE do you chose: ")
	licenseName := ioe.ReadInput()

	license, err := generateLicense(licenseName)
	if err != nil {
		return fmt.Errorf("failed to generate license %s: %w", licenseName, err)
	}

	outputFile := filepath.Join(a.flags.output, licenseFilename)
	if err := os.WriteFile(outputFile, []byte(license), os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file %s: %w", outputFile, err)
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.output = c.String(flagOutputName)
}
