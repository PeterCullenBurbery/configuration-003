package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/date_time_functions"
)

func main() {
	downloadURL := "https://www.giuspen.net/software/cherrytree_1.5.0.0_win64_setup.exe"
	downloadDir := `C:\downloads\cherry-tree`
	logDir := `C:\logs\cherry-tree`
	fileName := "cherrytree_1.5.0.0_win64_setup.exe"

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

	// Download the installer
	downloadPath := filepath.Join(downloadDir, fileName)
	if err := downloadFile(downloadPath, downloadURL); err != nil {
		log.Fatalf("❌ Failed to download installer: %v", err)
	}
	log.Printf("✅ Downloaded to: %s", downloadPath)

	// Prepare PowerShell command
	psCommand := fmt.Sprintf(`Start-Process -FilePath "%s" -ArgumentList "/VERYSILENT", "/SUPPRESSMSGBOXES", "/NORESTART", "/SP-", "/LOG=%s" -Wait`, downloadPath, logFilePath)

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", psCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("🚀 Running installer...")
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Installation failed: %v", err)
	}
	log.Println("✅ Cherrytree installation completed.")
}

func downloadFile(filepath string, url string) error {
	// Send HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP GET failed: %w", err)
	}
	defer resp.Body.Close()

	// Create the destination file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Copy data
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy data: %w", err)
	}

	return nil
}
