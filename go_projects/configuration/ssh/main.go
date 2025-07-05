package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	log.Println("🔐 Enabling SSH service and firewall rule...")

	if err := system_management_functions.Enable_SSH(); err != nil {
		log.Fatalf("❌ SSH service setup failed: %v", err)
	}

	if err := system_management_functions.Enable_SSH_through_firewall(); err != nil {
		log.Fatalf("❌ SSH firewall setup failed: %v", err)
	}

	log.Println("✅ SSH service and firewall rule configured successfully.")
}
