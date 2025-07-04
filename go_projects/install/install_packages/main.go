package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/PeterCullenBurbery/go_functions_002/system_management_functions"
	"github.com/PeterCullenBurbery/go_functions_002/yaml_functions"
)

type WhatToInstall struct {
	Winget []string `yaml:"winget"`
	Choco  []string `yaml:"choco"`
}

func load_yaml(path string) (*WhatToInstall, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to read YAML file: %w", err)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("❌ Failed to unmarshal YAML: %w", err)
	}

	install_map := yaml_functions.GetCaseInsensitiveMap(data, "what to install")
	if install_map == nil {
		return nil, fmt.Errorf("❌ 'what to install' section not found")
	}

	winget := yaml_functions.GetCaseInsensitiveList(install_map, "winget")
	choco := yaml_functions.GetCaseInsensitiveList(install_map, "choco")

	return &WhatToInstall{
		Winget: winget,
		Choco:  choco,
	}, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: install_packages.exe <path-to-what-to-install.yaml>")
		os.Exit(1)
	}

	yamlPath := os.Args[1]
	absYamlPath, err := filepath.Abs(yamlPath)
	if err != nil {
		log.Fatalf("❌ Failed to resolve YAML path: %v", err)
	}

	log.Printf("📄 Loading install list from %s...", absYamlPath)
	installList, err := load_yaml(absYamlPath)
	if err != nil {
		log.Fatalf("❌ Error loading YAML: %v", err)
	}

	// Install winget packages
	for _, id := range installList.Winget {
		err := system_management_functions.Winget_install(id, id)
		if err != nil {
			log.Printf("❌ Winget install failed for %s: %v", id, err)
		}
	}

	// Install choco packages
	for _, id := range installList.Choco {
		err := system_management_functions.Choco_install(id)
		if err != nil {
			log.Printf("❌ Choco install failed for %s: %v", id, err)
		}
	}
}
