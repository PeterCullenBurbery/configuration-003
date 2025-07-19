package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func run_executable(label, exe_path string, args ...string) {
	if _, err := os.Stat(exe_path); err != nil {
		log.Fatalf("❌ %s not found at: %s\n%v", label, exe_path, err)
	}
	log.Printf("🚀 Launching %s: %s %v\n", label, exe_path, args)
	cmd := exec.Command(exe_path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ %s failed: %v", label, err)
	}
	log.Printf("✅ %s completed.\n", label)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: configure_powershell_profiles.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	base_dir := os.Args[1]
	powershell_path := filepath.Join(base_dir, "go_projects", "powershell")

	// 🔧 Run PowerShell with exact command: Install-Module -Name PowershellFunctions
	log.Println("📦 Running: Install-Module -Name PowershellFunctions")
	cmd := exec.Command("pwsh", "-NoProfile", "-Command", "Install-Module -Name PowershellFunctions")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Install-Module failed: %v", err)
	}
	log.Println("✅ Install-Module completed.")

	run_executable("ip_address.exe", filepath.Join(powershell_path, "IP_address", "IP_address.exe"), base_dir)
	run_executable("powershell_005_profile.exe", filepath.Join(powershell_path, "powershell_005_profile", "powershell_005_profile.exe"), base_dir)
	run_executable("powershell_007_profile.exe", filepath.Join(powershell_path, "powershell_007_profile", "powershell_007_profile.exe"), base_dir)

	log.Println("🏁 PowerShell profile configuration completed.")
}