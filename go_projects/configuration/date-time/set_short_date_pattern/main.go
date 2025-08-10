package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("📅 Setting short date pattern to yyyy-MM-dd-dddd...")

	err := system_management_functions.Set_short_date_pattern()
	if err != nil {
		log.Fatalf("❌ Failed to set short date pattern: %v", err)
	}

	log.Println("✅ Short date pattern applied successfully.")
}
