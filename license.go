package main

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/haunt98/ioe-go"
)

const (
	templatesPath = "templates"
)

//go:embed templates/*
var embedFS embed.FS

// map template name with filename
// always use upercase for license name
var templates = map[string]templateInfo{
	"MIT": {
		filename: "mit.txt",
		args: []string{
			"[year]",
			"[fullname]",
		},
	},
}

type templateInfo struct {
	filename string
	args     []string
}

func generateLicense(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty license name")
	}
	name = strings.ToUpper(name)

	templateInfo, ok := templates[name]
	if !ok {
		return "", fmt.Errorf("not support license %s", name)
	}

	// Get correct path of license
	path := filepath.Join(templatesPath, templateInfo.filename)

	// Read template
	templateRaw, err := embedFS.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	// Replace template
	template := string(templateRaw)
	for _, arg := range templateInfo.args {
		fmt.Printf("What is your %s: ", arg)
		value := ioe.ReadInput()

		template = strings.ReplaceAll(template, arg, value)
	}

	return template, nil
}
