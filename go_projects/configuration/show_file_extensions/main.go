package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("📂 Showing file extensions...")

	err := system_management_functions.Show_file_extensions(true)
	if err != nil {
		log.Fatalf("❌ Failed to show file extensions: %v", err)
	}

	log.Println("✅ File extensions are now visible.")
}