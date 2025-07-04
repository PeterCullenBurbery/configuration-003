package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v2/system_management_functions"
)

func main() {
	log.Println("⏱️ Enabling seconds display in taskbar clock...")

	err := system_management_functions.Seconds_in_taskbar(true)
	if err != nil {
		log.Fatalf("❌ Failed to enable seconds in taskbar clock: %v", err)
	}

	log.Println("✅ Taskbar clock now shows seconds.")
}