package main

import (
	"fmt"
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	minicondaDir := `C:\ProgramData\Miniconda3`
	scriptsDir := `C:\ProgramData\Miniconda3\Scripts`

	// Add Miniconda root
	fmt.Printf("➕ Adding %s to system PATH...\n", minicondaDir)
	if err := system_management_functions.Add_to_path(minicondaDir); err != nil {
		log.Fatalf("❌ Failed to add %s to PATH: %v", minicondaDir, err)
	} else {
		fmt.Printf("✅ Successfully added %s to PATH.\n", minicondaDir)
	}

	// Add Scripts subfolder
	fmt.Printf("➕ Adding %s to system PATH...\n", scriptsDir)
	if err := system_management_functions.Add_to_path(scriptsDir); err != nil {
		log.Fatalf("❌ Failed to add %s to PATH: %v", scriptsDir, err)
	} else {
		fmt.Printf("✅ Successfully added %s to PATH.\n", scriptsDir)
	}
}
