package main

import (
	"bufio"
	"os"
	"strings"
)

// Read input until newline,
// ignore empty line
func readStdin() string {
	bs := bufio.NewScanner(os.Stdin)
	for bs.Scan() {
		line := bs.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		return line
	}

	return ""
}
