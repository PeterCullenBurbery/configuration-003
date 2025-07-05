package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("🕒 Applying custom time format: HH.mm.ss...")

	err := system_management_functions.Set_time_pattern()
	if err != nil {
		log.Fatalf("❌ Failed to set time pattern: %v", err)
	}

	log.Println("✅ Time pattern configured successfully.")
}