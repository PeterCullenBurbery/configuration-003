package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/PeterCullenBurbery/go_functions_002/system_management_functions"
)

func main() {
	const scriptURL = "https://github.com/PeterCullenBurbery/configuration-003/raw/main/script.ps1"

	// Generate UUID-based temp directory
	uuidName := uuid.New().String()
	tempDir, err := os.MkdirTemp("", uuidName)
	if err != nil {
		fmt.Printf("❌ Failed to create temp directory: %v\n", err)
		return
	}
	defer fmt.Printf("🧹 Temp directory: %s\n", tempDir)

	// Download the script using your helper function
	scriptPath := filepath.Join(tempDir, "script.ps1")
	err = system_management_functions.Download_file(scriptPath, scriptURL)
	if err != nil {
		fmt.Printf("❌ Failed to download script: %v\n", err)
		return
	}
	fmt.Printf("✅ Script downloaded to: %s\n", scriptPath)

	// Detect PowerShell
	pwshPath := "pwsh.exe"
	if _, err := exec.LookPath(pwshPath); err != nil {
		fmt.Println("⚠️ pwsh.exe not found, falling back to powershell.exe")
		pwshPath = "powershell.exe"
	}

	// Execute the script
	cmd := exec.Command(pwshPath, "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", scriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("🚀 Running %s\n", pwshPath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Script execution failed: %v\n", err)
	}
}