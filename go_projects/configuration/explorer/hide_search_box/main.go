package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("🔍 Hiding the search box on the taskbar...")

	err := system_management_functions.Hide_search_box(true)
	if err != nil {
		log.Fatalf("❌ Failed to hide search box: %v", err)
	}

	log.Println("✅ Search box is now hidden.")
}
