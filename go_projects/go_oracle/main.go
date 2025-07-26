package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	// Define download targets
	// Step 1: Define Oracle Instant Client download targets
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
	// Step 2: Set working directories
	zip_directory := `C:\downloads\oracle_instant_client\zips`
	extract_directory := `C:\downloads\oracle_instant_client\instantclient_23_8`
	final_path := filepath.Join(extract_directory, "instantclient_23_8")

	// Ensure zip and extract directories exist
	// Step 3: Ensure directories exist
	if err := os.MkdirAll(zip_directory, 0755); err != nil {
		log.Fatalf("❌ Could not create zip_directory: %v", err)
	}
	if err := os.MkdirAll(extract_directory, 0755); err != nil {
		log.Fatalf("❌ Could not create extract_directory: %v", err)
	}

	// Download and extract each file
	// Step 4: Download and extract Oracle files
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
	// Step 5: Add Oracle Instant Client to system PATH
	fmt.Println("➕ Adding Oracle Instant Client to system PATH...")
	if err := system_management_functions.Add_to_path(final_path); err != nil {
		log.Fatalf("❌ Failed to add Oracle path to PATH: %v", err)
	}
	fmt.Println("✅ Oracle Instant Client path added.")

	// Step 6: Install MSYS2
	fmt.Println("📦 Installing MSYS2 via Chocolatey...")
	if err := system_management_functions.Choco_install("msys2"); err != nil {
		log.Fatalf("❌ Failed to install MSYS2: %v", err)
	}
	fmt.Println("✅ MSYS2 installed.")

	// Step 7: Run pacman -Syu
	fmt.Println("🔄 Updating MSYS2 with pacman -Syu...")
	err := exec.Command(`C:\tools\msys64\msys2_shell.cmd`, "-defterm", "-no-start", "-mingw64", "-c", "pacman -Syu --noconfirm").Run()
	if err != nil {
		log.Fatalf("❌ Failed to run pacman -Syu: %v", err)
	}
	fmt.Println("✅ MSYS2 packages updated.")

	// Step 8: Install mingw-w64-x86_64-gcc
	fmt.Println("🛠 Installing mingw-w64-x86_64-gcc...")
	err = exec.Command(`C:\tools\msys64\msys2_shell.cmd`, "-defterm", "-no-start", "-mingw64", "-c", "pacman -S --noconfirm mingw-w64-x86_64-gcc").Run()
	if err != nil {
		log.Fatalf("❌ Failed to install mingw-w64-x86_64-gcc: %v", err)
	}
	fmt.Println("✅ GCC installed.")

	// Step 9: Add gcc path to system PATH
	gcc_bin_path := `C:\tools\msys64\mingw64\bin`
	fmt.Println("➕ Adding GCC to system PATH...")
	if err := system_management_functions.Add_to_path(gcc_bin_path); err != nil {
		log.Fatalf("❌ Failed to add GCC path to PATH: %v", err)
	}
	fmt.Println("✅ GCC path added to system PATH.")

	// Step 10: Done
	fmt.Println("🎉 Setup complete: Oracle Instant Client + MSYS2 + GCC.")
}