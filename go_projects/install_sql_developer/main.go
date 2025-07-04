package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/system_management_functions"
)

func main() {
	const downloadURL = "https://download.oracle.com/otn_software/java/sqldeveloper/sqldeveloper-24.3.1.347.1826-x64.zip"
	const zipName = "sqldeveloper-24.3.1.347.1826-x64.zip"

	// Get timestamp before download
	t1, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to get download timestamp: %v", err)
	}
	safeTimestamp1 := date_time_functions.Safe_time_stamp(t1, 1)
	baseDir := filepath.Join(`C:\downloads\sql-developer`, safeTimestamp1)

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create download directory: %v", err)
	}

	zipPath := filepath.Join(baseDir, zipName)

	// Download the file
	log.Printf("📥 Downloading SQL Developer to: %s", zipPath)
	if err := system_management_functions.Download_file(zipPath, downloadURL); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}

	// Get timestamp after download
	t2, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to get extraction timestamp: %v", err)
	}
	safeTimestamp2 := date_time_functions.Safe_time_stamp(t2, 1)
	extractDir := filepath.Join(baseDir, safeTimestamp2)

	// Extract ZIP
	log.Printf("📂 Extracting to: %s", extractDir)
	if err := system_management_functions.Extract_zip(zipPath, extractDir); err != nil {
		log.Fatalf("❌ Extraction failed: %v", err)
	}
	log.Println("✅ Extraction complete.")

	// Path to sqldeveloper.exe
	finalExe := filepath.Join(extractDir, "sqldeveloper", "sqldeveloper.exe")
	log.Printf("🚀 SQL Developer located at: %s", finalExe)

	// Create desktop shortcut (maximized, all users)
	if err := system_management_functions.Create_desktop_shortcut(
		finalExe,
		"SQL Developer.lnk",
		"Launch Oracle SQL Developer",
		3,    // WindowStyle 3 = Maximized
		true, // All users
	); err != nil {
		log.Fatalf("❌ Failed to create desktop shortcut: %v", err)
	}
	log.Println("📎 Desktop shortcut created successfully (maximized).")
}