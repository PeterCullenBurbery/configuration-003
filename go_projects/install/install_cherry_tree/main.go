package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/system_management_functions"
)

func main() {
	downloadURL := "https://www.giuspen.net/software/cherrytree_1.5.0.0_win64_setup.exe"
	downloadDir := `C:\downloads\cherry-tree`
	logDir := `C:\logs\cherry-tree`
	fileName := "cherrytree_1.5.0.0_win64_setup.exe"
	installPath := filepath.Join(downloadDir, fileName)

	// Create download and log directories
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create download directory: %v", err)
	}
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create log directory: %v", err)
	}

	// Get safe timestamp for log file
	timestamp, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to generate timestamp: %v", err)
	}
	safeTimestamp := date_time_functions.Safe_time_stamp(timestamp, 1)
	logFilePath := filepath.Join(logDir, safeTimestamp+".log")

	// Download the installer using your shared function
	if err := system_management_functions.Download_file(installPath, downloadURL); err != nil {
		log.Fatalf("❌ Failed to download installer: %v", err)
	}
	log.Printf("✅ Downloaded to: %s", installPath)

	// Prepare and run the PowerShell installer command
	log.Println("🚀 Running installer...")

	powershellArgs := []string{
		"-NoProfile", "-ExecutionPolicy", "Bypass", "-Command",
		fmt.Sprintf(`Start-Process -FilePath "%s" -ArgumentList '/VERYSILENT','/SUPPRESSMSGBOXES','/NORESTART','/SP-','/LOG=%s' -Wait`, installPath, logFilePath),
	}

	cmd := exec.Command("powershell", powershellArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Installation failed: %v", err)
	}

	log.Println("✅ Cherrytree installation completed.")
}