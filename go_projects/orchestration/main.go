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
		fmt.Println("❌ Usage: orchestration.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]
	installThingsExe := filepath.Join(baseDir, "go_projects", "install", "install_things", "install_things.exe")

	if _, err := os.Stat(installThingsExe); err != nil {
		log.Fatalf("❌ install_things.exe not found at: %s\n%v", installThingsExe, err)
	}

	log.Printf("🚀 Launching installer: %s %s\n", installThingsExe, baseDir)
	cmd := exec.Command(installThingsExe, baseDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ install_things.exe failed: %v", err)
	}

	log.Println("✅ orchestration.exe finished successfully.")
}