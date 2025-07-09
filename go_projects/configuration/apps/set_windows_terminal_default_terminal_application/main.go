package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	// Registry path
	keyPath := `Console\%%Startup`

	// GUIDs for Windows Terminal
	const delegationConsoleGUID = "{2EACA947-7F5F-4CFA-BA87-8F7FBEEFBE69}"
	const delegationTerminalGUID = "{E12CFF52-A866-4C77-9A90-F570A7AA2C6B}"

	// Open or create the key
	key, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Fatalf("❌ Failed to open or create registry key: %v", err)
	}
	defer key.Close()

	// Set values
	if err := key.SetStringValue("DelegationConsole", delegationConsoleGUID); err != nil {
		log.Fatalf("❌ Failed to set DelegationConsole: %v", err)
	}

	if err := key.SetStringValue("DelegationTerminal", delegationTerminalGUID); err != nil {
		log.Fatalf("❌ Failed to set DelegationTerminal: %v", err)
	}

	fmt.Println("✅ Windows Terminal has been set as the default terminal application using GUIDs.")
}