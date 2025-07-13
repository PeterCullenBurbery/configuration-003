package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// List of directories to create
	dirs := []string{
		`C:\GitHub-repositories`,
		`C:\go-projects`,
		`C:\python-projects`,
	}

	for _, dir := range dirs {
		absPath := filepath.Clean(dir)
		err := os.MkdirAll(absPath, 0755)
		if err != nil {
			fmt.Printf("❌ Failed to create %s: %v\n", absPath, err)
		} else {
			fmt.Printf("📁 Created: %s\n", absPath)
		}
	}
}