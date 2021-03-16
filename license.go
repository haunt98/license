package main

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
)

const (
	templatesPath = "templates"
)

//go:embed templates/*
var embedFS embed.FS

// map template name with filename
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
		value := readStdin()

		template = strings.ReplaceAll(template, arg, value)
	}

	return template, nil
}
