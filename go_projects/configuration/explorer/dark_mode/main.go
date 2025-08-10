package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("🌙 Setting Windows to dark mode...")

	// Pass true to restart Explorer and apply changes immediately
	err := system_management_functions.Set_dark_mode(true)
	if err != nil {
		log.Fatalf("❌ Failed to set dark mode: %v", err)
	}

	log.Println("✅ Dark mode has been enabled successfully.")
}
