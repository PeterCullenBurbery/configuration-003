package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	// Define download targets
	files_to_download := []struct {
		file_name string
		url       string
	}{
		{
			file_name: "instantclient-basic-windows.x64-23.8.0.25.04.zip",
			url:       "https://download.oracle.com/otn_software/nt/instantclient/2380000/instantclient-basic-windows.x64-23.8.0.25.04.zip",
		},
		{
			file_name: "instantclient-sdk-windows.x64-23.8.0.25.04.zip",
			url:       "https://download.oracle.com/otn_software/nt/instantclient/2380000/instantclient-sdk-windows.x64-23.8.0.25.04.zip",
		},
	}

	// Set working directories
	zip_directory := `C:\downloads\oracle_instant_client\zips`
	extract_directory := `C:\downloads\oracle_instant_client\instantclient_23_8`
	final_path := filepath.Join(extract_directory, "instantclient_23_8")

	// Ensure zip and extract directories exist
	if err := os.MkdirAll(zip_directory, 0755); err != nil {
		log.Fatalf("❌ Could not create zip_directory: %v", err)
	}
	if err := os.MkdirAll(extract_directory, 0755); err != nil {
		log.Fatalf("❌ Could not create extract_directory: %v", err)
	}

	// Download and extract each file
	for _, file := range files_to_download {
		zip_path := filepath.Join(zip_directory, file.file_name)
		fmt.Printf("📥 Downloading %s...\n", file.url)

		if err := system_management_functions.Download_file(zip_path, file.url); err != nil {
			log.Fatalf("❌ Failed to download %s: %v", file.file_name, err)
		}
		fmt.Printf("✅ Downloaded: %s\n", zip_path)

		fmt.Printf("📦 Extracting %s...\n", zip_path)
		if err := system_management_functions.Extract_zip(zip_path, extract_directory); err != nil {
			log.Fatalf("❌ Failed to extract %s: %v", file.file_name, err)
		}
		fmt.Printf("✅ Extracted to: %s\n", extract_directory)
	}

	// Add to system PATH
	fmt.Println("➕ Adding Instant Client to system PATH...")
	if err := system_management_functions.Add_to_path(final_path); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}

	fmt.Println("🎉 Oracle Instant Client Basic + SDK setup complete.")
}