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
		fmt.Println("❌ Usage: call_installer.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]

	// Resolve paths
	yamlPath := filepath.Join(baseDir, "what-to-install.yaml")
	exePath := filepath.Join(baseDir, "go_projects", "install_packages", "install_packages.exe")

	// Confirm files exist
	if _, err := os.Stat(yamlPath); err != nil {
		log.Fatalf("❌ Could not find what-to-install.yaml at: %s\n%v", yamlPath, err)
	}

	if _, err := os.Stat(exePath); err != nil {
		log.Fatalf("❌ Could not find install_packages.exe at: %s\n%v", exePath, err)
	}

	// Log
	log.Printf("📦 Calling: %s %s\n", exePath, yamlPath)

	// Execute
	cmd := exec.Command(exePath, yamlPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ install_packages.exe failed: %v", err)
	}

	log.Println("✅ Installation completed successfully.")
}