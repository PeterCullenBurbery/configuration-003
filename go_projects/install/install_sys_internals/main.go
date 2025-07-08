package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v3/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	// Constants
	downloadURL := "https://download.sysinternals.com/files/SysinternalsSuite.zip"
	baseDir := `C:\downloads\sys-internals`
	zipName := "SysinternalsSuite.zip"

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
	zipPath := filepath.Join(baseDir, zipName)
	extractDir := filepath.Join(baseDir, safeTimestamp)

	// Step 3: Download the ZIP file
	log.Printf("⬇️ Downloading to: %s", zipPath)
	if err := system_management_functions.Download_file(zipPath, downloadURL); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	log.Println("✅ Download complete")

	// Step 4: Extract ZIP (not password protected)
	log.Printf("📦 Extracting to: %s", extractDir)
	if err := system_management_functions.Extract_zip(zipPath, extractDir); err != nil {
		log.Fatalf("❌ Extraction failed: %v", err)
	}
	log.Println("✅ Extraction complete")

	// Step 5: Add extractDir to PATH
	log.Printf("➕ Adding to PATH: %s", extractDir)
	if err := system_management_functions.Add_to_path(extractDir); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}
	log.Println("✅ Sysinternals directory added to system PATH")
}
