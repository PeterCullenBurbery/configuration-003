package main

import (
	"fmt"
	"os"

	"github.com/PeterCullenBurbery/go_functions_002/v5/system_management_functions"
)

func main() {
	fmt.Println("🧼 Starting PATH cleanup...")

	err := system_management_functions.Clean_path()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to clean PATH: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ PATH cleanup completed successfully.")
}