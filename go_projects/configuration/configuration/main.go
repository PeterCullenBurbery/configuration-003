package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: configuration.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]

	// Path to the dark_mode.exe binary
	darkModeExe := filepath.Join(baseDir, "go_projects", "configuration", "dark_mode", "dark_mode.exe")

	if _, err := os.Stat(darkModeExe); err != nil {
		log.Fatalf("❌ Could not find dark_mode.exe at: %s\n%v", darkModeExe, err)
	}

	log.Printf("🌙 Running: %s\n", darkModeExe)
	cmd := exec.Command(darkModeExe)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ dark_mode.exe failed: %v", err)
	}
	log.Println("✅ dark_mode.exe completed.")
}
