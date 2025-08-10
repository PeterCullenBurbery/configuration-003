package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("👀 Enabling visibility of hidden files...")

	err := system_management_functions.Show_hidden_files(true)
	if err != nil {
		log.Fatalf("❌ Failed to show hidden files: %v", err)
	}

	log.Println("✅ Hidden files are now visible.")
}
