package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("📅 Setting long date pattern to yyyy-MM-dd-dddd...")

	err := system_management_functions.Set_long_date_pattern()
	if err != nil {
		log.Fatalf("❌ Failed to set long date pattern: %v", err)
	}

	log.Println("✅ Long date pattern applied successfully.")
}