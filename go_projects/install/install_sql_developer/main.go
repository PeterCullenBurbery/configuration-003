package main

import (
	"log"
	"os"
	// "os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v4/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	downloadURL := "https://download.oracle.com/otn_software/java/sqldeveloper/sqldeveloper-24.3.1.347.1826-x64.zip"
	baseDir := `C:\downloads\sql-developer`
	fileName := "sqldeveloper-24.3.1.347.1826-x64.zip"
	downloadPath := filepath.Join(baseDir, fileName)

	// Ensure base download directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalf("❌ Failed to create base directory: %v", err)
	}

	// Download ZIP
	log.Println("🌐 Downloading SQL Developer...")
	if err := system_management_functions.Download_file(downloadPath, downloadURL); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	log.Println("✅ Download completed.")

	// Get safe timestamp
	rawTS, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to get timestamp: %v", err)
	}
	safeTS := date_time_functions.Safe_time_stamp(rawTS, 1)
	extractDir := filepath.Join(baseDir, safeTS)

	// Extract ZIP to timestamped folder
	log.Printf("📦 Extracting to: %s", extractDir)
	if err := system_management_functions.Extract_zip(downloadPath, extractDir); err != nil {
		log.Fatalf("❌ Extraction failed: %v", err)
	}
	log.Println("✅ Extraction complete.")

	// Path to sqldeveloper.exe
	exePath := filepath.Join(extractDir, "sqldeveloper", "sqldeveloper.exe")
	if _, err := os.Stat(exePath); err != nil {
		log.Fatalf("❌ Could not find sqldeveloper.exe: %v", err)
	}

	// Add to system PATH
	sqldeveloperBin := filepath.Dir(exePath)
	log.Printf("➕ Adding %s to system PATH...", sqldeveloperBin)
	if err := system_management_functions.Add_to_path(sqldeveloperBin); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}
	log.Println("✅ PATH updated.")

	// Create shortcut
	log.Println("🔗 Creating desktop shortcut...")
	err = system_management_functions.Create_desktop_shortcut(
		exePath,
		"SQL Developer.lnk",
		"Oracle SQL Developer",
		3,    // 3 = Maximized window
		true, // true = all users
	)
	if err != nil {
		log.Fatalf("❌ Failed to create shortcut: %v", err)
	}
	log.Println("✅ Shortcut created.")

	// // Download and execute part-004.exe from GitHub
	// log.Println("🌐 Downloading and executing part-004.exe from GitHub...")

	// blobURL := "https://github.com/PeterCullenBurbery/python-projects-semipublic/blob/main/gui_automation/sql_developer/sql_developer/sql_developer_004/dist/part-004.exe"
	// rawURL, err := system_management_functions.Convert_blob_to_raw_github_url(blobURL)
	// if err != nil {
	// 	log.Fatalf("❌ Failed to convert blob URL: %v", err)
	// }

	// tempDir := os.TempDir()
	// destinationPath := filepath.Join(tempDir, "part_004.exe")

	// log.Printf("⬇️  Downloading to: %s", destinationPath)
	// if err := system_management_functions.Download_file(destinationPath, rawURL); err != nil {
	// 	log.Fatalf("❌ Download failed: %v", err)
	// }
	// log.Println("✅ Download complete.")

	// log.Println("🚀 Running the downloaded program...")
	// cmd := exec.Command(destinationPath)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("❌ Execution failed: %v", err)
	// }
	// log.Println("✅ Execution complete.")
}