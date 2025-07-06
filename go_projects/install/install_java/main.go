package main

import (
	"fmt"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	fmt.Println("🛠️ Starting Java installation...")

	if err := system_management_functions.Install_Java(); err != nil {
		fmt.Printf("❌ Java installation failed: %v\n", err)
		return
	}

	fmt.Println("✅ Java installation completed successfully.")
}
