package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("⬅️ Setting Start menu to left...")

	err := system_management_functions.Set_start_menu_to_left()
	if err != nil {
		log.Fatalf("❌ Failed to set Start menu alignment: %v", err)
	}

	log.Println("✅ Start menu is now aligned to the left.")
}
