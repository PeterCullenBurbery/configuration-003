package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v5/system_management_functions"
)

func main() {
	// Google Drive file ID and download URL
	driveFileID := "1sKv7YP2hTq322iFaiZ0lpZoOSXzDHR91"
	downloadURL := "https://drive.google.com/uc?export=download&id=" + driveFileID

	// Output path
	downloadDir := `C:\downloads\oracle-database`
	fileName := "test_image.jpg"
	downloadPath := filepath.Join(downloadDir, fileName)

	// Ensure target folder exists
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create folder %s: %v", downloadDir, err)
	}

	// Download the file
	log.Printf("⬇️ Downloading test image to: %s", downloadPath)
	if err := system_management_functions.Download_file(downloadPath, downloadURL); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	log.Println("✅ Download complete")
}