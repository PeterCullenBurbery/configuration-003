package main

import (
	"fmt"
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	ip, err := system_management_functions.Get_primary_ipv4_address()
	if err != nil {
		log.Fatalf("❌ Failed to get primary IPv4 address: %v", err)
	}
	fmt.Println("🌐 Primary IPv4 address:", ip)

	err = system_management_functions.Set_system_environment_variable("LOCAL_IPV4_ADDRESS", ip)
	if err != nil {
		log.Fatalf("❌ Failed to set LOCAL_IPV4_ADDRESS: %v", err)
	}

	fmt.Println("✅ LOCAL_IPV4_ADDRESS environment variable set successfully.")
}
