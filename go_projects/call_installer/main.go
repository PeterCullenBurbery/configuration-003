package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: call_installer.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]

	// Paths
	yamlPath := filepath.Join(baseDir, "what-to-install.yaml")
	installPackagesExe := filepath.Join(baseDir, "go_projects", "install_packages", "install_packages.exe")
	installCherryTreeExe := filepath.Join(baseDir, "go_projects", "install_cherry_tree", "install_cherry_tree.exe")

	// Check if what-to-install.yaml exists
	if _, err := os.Stat(yamlPath); err != nil {
		log.Fatalf("❌ Could not find what-to-install.yaml at: %s\n%v", yamlPath, err)
	}

	// Check if install_packages.exe exists
	if _, err := os.Stat(installPackagesExe); err != nil {
		log.Fatalf("❌ Could not find install_packages.exe at: %s\n%v", installPackagesExe, err)
	}

	// Run install_packages.exe
	log.Printf("📦 Running: %s %s\n", installPackagesExe, yamlPath)
	cmd1 := exec.Command(installPackagesExe, yamlPath)
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	if err := cmd1.Run(); err != nil {
		log.Fatalf("❌ install_packages.exe failed: %v", err)
	}
	log.Println("✅ Base package installation completed.")

	// Check if install_cherry_tree.exe exists
	if _, err := os.Stat(installCherryTreeExe); err != nil {
		log.Fatalf("❌ Could not find install_cherry_tree.exe at: %s\n%v", installCherryTreeExe, err)
	}

	// Run install_cherry_tree.exe
	log.Printf("🍒 Running: %s\n", installCherryTreeExe)
	cmd2 := exec.Command(installCherryTreeExe)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	if err := cmd2.Run(); err != nil {
		log.Fatalf("❌ install_cherry_tree.exe failed: %v", err)
	}
	log.Println("✅ Cherrytree installation completed.")
}