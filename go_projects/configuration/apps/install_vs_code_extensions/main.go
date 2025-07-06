package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Structure for YAML input
type VSCodeExtensions struct {
	VSCodeExtensions []string `yaml:"vs_code_extensions"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: install_vs_code_extensions.exe <path-to-yaml>")
		os.Exit(1)
	}
	yamlPath := os.Args[1]

	// Open the YAML file
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalf("❌ Failed to read YAML file: %v", err)
	}

	var extensions VSCodeExtensions
	if err := yaml.Unmarshal(yamlFile, &extensions); err != nil {
		log.Fatalf("❌ Failed to parse YAML file: %v", err)
	}

	codeCmd := `C:\Program Files\Microsoft VS Code\bin\code.cmd`

	for _, ext := range extensions.VSCodeExtensions {
		fmt.Printf("📦 Installing VS Code extension: %s\n", ext)
		cmd := exec.Command(codeCmd, "--install-extension", ext)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Failed to install extension %s: %v\n", ext, err)
		} else {
			fmt.Printf("✅ Successfully installed: %s\n", ext)
		}
	}
}