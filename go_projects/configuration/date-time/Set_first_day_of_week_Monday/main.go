package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("📅 Setting Monday as the first day of the week...")

	err := system_management_functions.Set_first_day_of_week_Monday()
	if err != nil {
		log.Fatalf("❌ Failed to set first day of the week: %v", err)
	}

	log.Println("✅ First day of the week set to Monday successfully.")
}