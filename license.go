package main

import (
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/make-go-great/ioe-go"
)

const (
	templatesPath = "templates"
)

var ErrInvalidLicense = errors.New("invalid license")

//go:embed templates/*
var embedFS embed.FS

// map template name with filename
// always use upercase for license name
var templates = map[string]templateInfo{
	"MIT": {
		templateFilename: "mit.txt",
		outputFilename:   "LICENSE",
		args: []string{
			"[year]",
			"[fullname]",
		},
	},
	"GNU GPLv3": {
		templateFilename: "gnu_gplv3.txt",
		outputFilename:   "COPYING",
	},
}

type templateInfo struct {
	templateFilename string
	outputFilename   string
	args             []string
}

func generateLicense(name string) (string, string, error) {
	if name == "" {
		return "", "", fmt.Errorf("empty license name: %w", ErrInvalidLicense)
	}

	isSupportTemplate := false
	var templateInfo templateInfo
	for templateName := range templates {
		if strings.EqualFold(templateName, name) {
			isSupportTemplate = true
			templateInfo = templates[templateName]
		}
	}

	if !isSupportTemplate {
		return "", "", fmt.Errorf("not support license %s: %w", name, ErrInvalidLicense)
	}

	// Get correct path of license
	path := filepath.Join(templatesPath, templateInfo.templateFilename)

	// Read template
	templateRaw, err := embedFS.ReadFile(path)
	if err != nil {
		return "", "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	// Replace template info args
	licenseData := string(templateRaw)
	for _, arg := range templateInfo.args {
		fmt.Printf("What is your %s: ", arg)
		value := ioe.ReadInput()

		licenseData = strings.ReplaceAll(licenseData, arg, value)
	}

	return licenseData, templateInfo.outputFilename, nil
}
