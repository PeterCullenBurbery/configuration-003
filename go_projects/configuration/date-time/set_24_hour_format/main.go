package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("🕓 Setting system to 24-hour time format...")

	err := system_management_functions.Set_24_hour_format()
	if err != nil {
		log.Fatalf("❌ Failed to set 24-hour time format: %v", err)
	}

	log.Println("✅ 24-hour time format applied successfully.")
}