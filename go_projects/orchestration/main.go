package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func runExecutable(label, exePath string, args ...string) {
	if _, err := os.Stat(exePath); err != nil {
		log.Fatalf("❌ %s not found at: %s\n%v", label, exePath, err)
	}
	log.Printf("🚀 Launching %s: %s %v\n", label, exePath, args)
	cmd := exec.Command(exePath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ %s failed: %v", label, err)
	}
	log.Printf("✅ %s completed.\n", label)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: orchestration.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]

	installThings := filepath.Join(baseDir, "go_projects", "install", "install_things", "install_things.exe")
	configuration := filepath.Join(baseDir, "go_projects", "configuration", "configuration", "configuration.exe")
	configurePowershell := filepath.Join(baseDir, "go_projects", "powershell", "configure_powershell_modules", "configure_powershell_modules.exe")

	runExecutable("install_things.exe", installThings, baseDir)
	runExecutable("configuration.exe", configuration, baseDir)
	runExecutable("configure_powershell_modules.exe", configurePowershell, baseDir)

	log.Println("🏁 orchestration.exe finished successfully.")
}