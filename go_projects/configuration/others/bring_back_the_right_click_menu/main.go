// enable_right_click_menu.go

package main

import (
	"log"

	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	log.Println("🖱️ Enabling classic right-click menu (Windows 10 style)...")

	if err := system_management_functions.Bring_back_the_right_click_menu(); err != nil {
		log.Fatalf("❌ Failed to enable classic right-click menu: %v", err)
	}

	log.Println("🎉 Classic right-click menu enabled successfully.")
}
