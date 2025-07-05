package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	module_dir := filepath.Join("C:", "powershell-modules", "MyModule")

	// Ensure target directory exists
	if err := os.MkdirAll(module_dir, 0755); err != nil {
		fmt.Printf("❌ Failed to create module directory: %v\n", err)
		return
	}

	files := map[string]string{
		"MyModule.psd1": "https://github.com/PeterCullenBurbery/configuration-003/blob/main/powershell-modules/MyModule/MyModule.psd1",
		"MyModule.psm1": "https://github.com/PeterCullenBurbery/configuration-003/blob/main/powershell-modules/MyModule/MyModule.psm1",
	}

	for filename, blob_url := range files {
		raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
		if err != nil {
			fmt.Printf("❌ %v\n", err)
			continue
		}

		dest_path := filepath.Join(module_dir, filename)
		if err := system_management_functions.Download_file(dest_path, raw_url); err != nil {
			fmt.Printf("❌ Failed to download %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Downloaded %s to %s\n", filename, dest_path)
	}

	// Add to PSModulePath using the .psm1 file path
	psm1_path := filepath.Join(module_dir, "MyModule.psm1")
	if err := system_management_functions.Add_to_ps_module_path(psm1_path); err != nil {
		fmt.Printf("❌ Failed to add to PSModulePath: %v\n", err)
	} else {
		fmt.Printf("✅ Successfully added to PSModulePath: %s\n", psm1_path)
	}
}