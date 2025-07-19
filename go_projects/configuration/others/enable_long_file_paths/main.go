// enable_long_paths.go
package main

import (
	"fmt"
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	fmt.Println("🔧 Enabling long file paths support...")

	if err := system_management_functions.Enable_long_file_paths(); err != nil {
		log.Fatalf("❌ Error: %v", err)
	}

	fmt.Println("🎉 Long file paths support is now enabled.")
}