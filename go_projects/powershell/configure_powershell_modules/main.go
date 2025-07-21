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

	// Install PowershellFunctions005 with Windows PowerShell
	log.Println("📦 Installing PowershellFunctions005 with Windows PowerShell (powershell.exe)")
	exec.Command("powershell.exe", "-NoProfile", "-Command", "Install-PackageProvider -Name NuGet -MinimumVersion 2.8.5.201 -Force").Run()
	exec.Command("powershell.exe", "-NoProfile", "-Command", "Set-PSRepository -Name PSGallery -InstallationPolicy Trusted").Run()

	cmd := exec.Command("powershell.exe", "-NoProfile", "-Command", "Install-Module -Name PowershellFunctions005 -Force")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Printf("⚠️ Install-Module PowershellFunctions005 with powershell.exe failed: %v", err)
	} else {
		log.Println("✅ Installed PowershellFunctions005 with powershell.exe")
	}

	// Install PowershellFunctions007 with PowerShell 7+
	pwsh_path := "pwsh"
	if _, err := exec.LookPath(pwsh_path); err != nil {
		alt := `C:\Program Files\PowerShell\7\pwsh.exe`
		if _, err := os.Stat(alt); err == nil {
			pwsh_path = alt
			log.Printf("ℹ️ Using fallback pwsh path: %s", pwsh_path)
		} else {
			log.Fatalln("❌ pwsh not found — cannot install PowershellFunctions007")
		}
	}

	log.Println("📦 Installing PowershellFunctions007 with pwsh")
	exec.Command(pwsh_path, "-NoProfile", "-Command", "Set-PSRepository -Name PSGallery -InstallationPolicy Trusted").Run()

	install_cmd := exec.Command(pwsh_path, "-NoProfile", "-Command", "Install-Module -Name PowershellFunctions007 -Force -AllowClobber")
	install_cmd.Stdout = os.Stdout
	install_cmd.Stderr = os.Stderr
	if err := install_cmd.Run(); err != nil {
		log.Fatalf("❌ Install-Module PowershellFunctions007 with pwsh failed: %v", err)
	} else {
		log.Println("✅ Installed PowershellFunctions007 with pwsh")
	}

	// Execute profile config tools
	run_executable("ip_address.exe", filepath.Join(powershell_path, "IP_address", "ip_address.exe"), base_dir)
	run_executable("powershell_005_profile.exe", filepath.Join(powershell_path, "powershell_005_profile", "powershell_005_profile.exe"), base_dir)
	run_executable("powershell_007_profile.exe", filepath.Join(powershell_path, "powershell_007_profile", "powershell_007_profile.exe"), base_dir)

	log.Println("🏁 PowerShell profile configuration completed.")
}