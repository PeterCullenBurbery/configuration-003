package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	// Get USERPROFILE environment variable
	user_profile := os.Getenv("USERPROFILE")
	if user_profile == "" {
		fmt.Println("❌ USERPROFILE environment variable not found.")
		return
	}

	// Form PowerShell 5 profile path
	profile_path := filepath.Join(user_profile, "Documents", "WindowsPowerShell", "Microsoft.PowerShell_profile.ps1")
	profile_dir := filepath.Dir(profile_path)

	// Ensure the directory exists
	err := os.MkdirAll(profile_dir, 0755)
	if err != nil {
		fmt.Printf("❌ Failed to create profile directory: %v\n", err)
		return
	}

	// Convert blob URL to raw GitHub content URL
	blob_url := "https://github.com/PeterCullenBurbery/powershell-profile/blob/main/powershell-005-profile/Microsoft.PowerShell_profile.ps1"
	raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
	if err != nil {
		fmt.Printf("❌ Failed to convert blob to raw URL: %v\n", err)
		return
	}

	// Download the profile
	err = system_management_functions.Download_file(profile_path, raw_url)
	if err != nil {
		fmt.Printf("❌ Failed to download profile: %v\n", err)
		return
	}

	fmt.Printf("✅ PowerShell 5 profile downloaded to: %s\n", profile_path)
}