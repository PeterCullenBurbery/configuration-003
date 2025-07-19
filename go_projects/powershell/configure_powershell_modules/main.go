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

	// Install for PowerShell 5
	log.Println("📦 Installing PowershellFunctions with Windows PowerShell (powershell.exe)")

	exec.Command("powershell.exe", "-NoProfile", "-Command", "Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force").Run()
	exec.Command("powershell.exe", "-NoProfile", "-Command", "Set-PSRepository -Name PSGallery -InstallationPolicy Trusted").Run()

	cmd1 := exec.Command("powershell.exe", "-NoProfile", "-Command", "Install-Module -Name PowershellFunctions")
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	if err := cmd1.Run(); err != nil {
		log.Printf("⚠️ Install-Module with powershell.exe failed: %v", err)
	} else {
		log.Println("✅ Installed with powershell.exe")
	}

	// Install for PowerShell 7
	log.Println("📦 Installing PowershellFunctions with PowerShell 7 (pwsh)")
	pwsh_path := "pwsh"
	if _, err := exec.LookPath(pwsh_path); err != nil {
		alt_path := `C:\Program Files\PowerShell\7\pwsh.exe`
		if _, err := os.Stat(alt_path); err == nil {
			pwsh_path = alt_path
			log.Printf("ℹ️ Using fallback path for pwsh: %s\n", pwsh_path)
		} else {
			log.Printf("⚠️ pwsh not found at default locations: %v", err)
			pwsh_path = "" // prevent usage
		}
	}

	if pwsh_path != "" {
		exec.Command(pwsh_path, "-NoProfile", "-Command", "Set-PSRepository -Name PSGallery -InstallationPolicy Trusted").Run()

		cmd2 := exec.Command(pwsh_path, "-NoProfile", "-Command", "Install-Module -Name PowershellFunctions")
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		if err := cmd2.Run(); err != nil {
			log.Printf("⚠️ Install-Module with pwsh failed: %v", err)
		} else {
			log.Println("✅ Installed with pwsh")
		}
	} else {
		log.Println("⚠️ Skipped PowerShell 7 installation: pwsh not found")
	}

	run_executable("ip_address.exe", filepath.Join(powershell_path, "IP_address", "IP_address.exe"), base_dir)
	run_executable("powershell_005_profile.exe", filepath.Join(powershell_path, "powershell_005_profile", "powershell_005_profile.exe"), base_dir)
	run_executable("powershell_007_profile.exe", filepath.Join(powershell_path, "powershell_007_profile", "powershell_007_profile.exe"), base_dir)

	log.Println("🏁 PowerShell profile configuration completed.")
}