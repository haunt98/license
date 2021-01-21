package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	templatesPath = "templates"
)

// map template name with filename
var templates = map[string]templateInfo{
	"MIT": {
		filename: "mit",
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

func generateLicense(name string, values map[string]string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("empty license name")
	}

	templateInfo, ok := templates[name]
	if !ok {
		return "", fmt.Errorf("not support license %s", name)
	}

	// Read template
	path := filepath.Join(templatesPath, templateInfo.filename)
	templateRaw, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	// Replace template
	template := string(templateRaw)
	for _, arg := range templateInfo.args {
		value, ok := values[arg]
		if !ok {
			return "", fmt.Errorf("missing arg %s", arg)
		}

		template = strings.ReplaceAll(template, arg, value)
	}

	return template, nil
}
