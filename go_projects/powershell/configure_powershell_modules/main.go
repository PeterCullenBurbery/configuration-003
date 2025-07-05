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

func restartExplorer() {
	log.Println("🔄 Restarting Explorer...")
	if err := exec.Command("taskkill", "/f", "/im", "explorer.exe").Run(); err != nil {
		log.Fatalf("❌ Failed to stop Explorer: %v", err)
	}
	if err := exec.Command("explorer.exe").Start(); err != nil {
		log.Fatalf("❌ Failed to restart Explorer: %v", err)
	}
	log.Println("✅ Explorer restarted successfully.")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: configure_powershell_profiles.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]
	psPath := filepath.Join(baseDir, "go_projects", "powershell")

	runExecutable("powershell_modules.exe", filepath.Join(psPath, "powershell_modules", "powershell_modules.exe"), baseDir)
	runExecutable("powershell_005_profile.exe", filepath.Join(psPath, "powershell_005_profile", "powershell_005_profile.exe"), baseDir)
	runExecutable("powershell_007_profile.exe", filepath.Join(psPath, "powershell_007_profile", "powershell_007_profile.exe"), baseDir)

	restartExplorer()

	log.Println("🏁 PowerShell profile configuration completed.")
}