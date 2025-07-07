package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	downloadURL := "https://repo.anaconda.com/miniconda/Miniconda3-latest-Windows-x86_64.exe"
	downloadDir := `C:\downloads\miniconda`
	installerName := "Miniconda3-latest-Windows-x86_64.exe"
	installDir := `C:\ProgramData\Miniconda3`

	// Ensure the download directory exists
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create download directory: %v", err)
	}

	downloadPath := filepath.Join(downloadDir, installerName)

	// Use your shared Download_file function
	if err := system_management_functions.Download_file(downloadPath, downloadURL); err != nil {
		log.Fatalf("❌ Failed to download Miniconda installer: %v", err)
	}
	log.Printf("✅ Downloaded Miniconda installer to: %s", downloadPath)

	// Construct PowerShell install command
	arguments := fmt.Sprintf(`"/S", "/InstallationType=AllUsers", "/RegisterPython=1", "/D=%s"`, installDir)
	psCommand := fmt.Sprintf(`Start-Process -FilePath "%s" -ArgumentList %s -Wait -NoNewWindow`, downloadPath, arguments)

	// Run the PowerShell command
	log.Println("🚀 Running Miniconda installer...")
	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", psCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Installation failed: %v", err)
	}

	log.Println("✅ Miniconda installation completed.")

	// Add Miniconda and Scripts folder to PATH
	if err := system_management_functions.Add_to_path(filepath.Join(installDir)); err != nil {
		log.Printf("⚠️ Failed to add %s to PATH: %v", installDir, err)
	} else {
		log.Printf("✅ Added %s to system PATH.", installDir)
	}

	scriptsPath := filepath.Join(installDir, "Scripts")
	if err := system_management_functions.Add_to_path(scriptsPath); err != nil {
		log.Printf("⚠️ Failed to add %s to PATH: %v", scriptsPath, err)
	} else {
		log.Printf("✅ Added %s to system PATH.", scriptsPath)
	}
}
