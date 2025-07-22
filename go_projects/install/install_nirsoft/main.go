package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v4/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	// Constants
	downloadURL := "https://github.com/PeterCullenBurbery/Nirsoft/raw/main/nirsoft_package_enc_1.30.19.zip"
	baseDir := `C:\downloads\nirsoft`
	password := "nirsoft9876$"

	// ✅ Step 0: Ensure base directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create base directory: %v", err)
	}

	// Step 1: Exclude from Defender
	if err := system_management_functions.Exclude_from_Microsoft_Windows_Defender(baseDir); err != nil {
		log.Fatalf("❌ Failed to exclude from Defender: %v", err)
	}

	// Step 2: Generate safe timestamp
	timestamp, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to generate timestamp: %v", err)
	}
	safeTimestamp := date_time_functions.Safe_time_stamp(timestamp, 1)

	// Paths
	zipPath := filepath.Join(baseDir, "nirsoft_package_enc_1.30.19.zip")
	extractDir := filepath.Join(baseDir, safeTimestamp)

	// Step 3: Download the ZIP file
	log.Printf("⬇️ Downloading to: %s", zipPath)
	if err := system_management_functions.Download_file(zipPath, downloadURL); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	log.Println("✅ Download complete")

	// Step 4: Extract the password-protected ZIP
	log.Printf("📦 Extracting to: %s", extractDir)
	if err := system_management_functions.Extract_password_protected_zip(zipPath, extractDir, password); err != nil {
		log.Fatalf("❌ Extraction failed: %v", err)
	}
	log.Println("✅ Extraction complete")

	// Step 5: Add NirSoft\x64 to PATH
	x64Path := filepath.Join(extractDir, "NirSoft", "x64")
	log.Printf("➕ Adding to PATH: %s", x64Path)
	if err := system_management_functions.Add_to_path(x64Path); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}
	log.Println("✅ NirSoft\\x64 added to system PATH")
}