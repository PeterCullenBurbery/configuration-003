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
	// installCherryTreeExe := filepath.Join(baseDir, "go_projects", "install_cherry_tree", "install_cherry_tree.exe")
	// installMinicondaExe := filepath.Join(baseDir, "go_projects", "install_miniconda", "install_miniconda.exe")
	// installSQLDeveloperExe := filepath.Join(baseDir, "go_projects", "install_sql_developer", "install_sql_developer.exe")
	installNirsoftExe := filepath.Join(baseDir, "go_projects", "install_nirsoft", "install_nirsoft.exe")

	// === Check config file exists ===
	if _, err := os.Stat(yamlPath); err != nil {
		log.Fatalf("❌ Could not find what-to-install.yaml at: %s\n%v", yamlPath, err)
	}

	// === Run install_packages.exe ===
	if _, err := os.Stat(installPackagesExe); err != nil {
		log.Fatalf("❌ Could not find install_packages.exe at: %s\n%v", installPackagesExe, err)
	}
	log.Printf("📦 Running: %s %s\n", installPackagesExe, yamlPath)
	cmd1 := exec.Command(installPackagesExe, yamlPath)
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	if err := cmd1.Run(); err != nil {
		log.Fatalf("❌ install_packages.exe failed: %v", err)
	}
	log.Println("✅ Base package installation completed.")

	// // === Run install_cherry_tree.exe ===
	// if _, err := os.Stat(installCherryTreeExe); err != nil {
	// 	log.Fatalf("❌ Could not find install_cherry_tree.exe at: %s\n%v", installCherryTreeExe, err)
	// }
	// log.Printf("🍒 Running: %s\n", installCherryTreeExe)
	// cmd2 := exec.Command(installCherryTreeExe)
	// cmd2.Stdout = os.Stdout
	// cmd2.Stderr = os.Stderr
	// if err := cmd2.Run(); err != nil {
	// 	log.Fatalf("❌ install_cherry_tree.exe failed: %v", err)
	// }
	// log.Println("✅ Cherrytree installation completed.")

	// // === Run install_miniconda.exe ===
	// if _, err := os.Stat(installMinicondaExe); err != nil {
	// 	log.Fatalf("❌ Could not find install_miniconda.exe at: %s\n%v", installMinicondaExe, err)
	// }
	// log.Printf("🐍 Running: %s\n", installMinicondaExe)
	// cmd3 := exec.Command(installMinicondaExe)
	// cmd3.Stdout = os.Stdout
	// cmd3.Stderr = os.Stderr
	// if err := cmd3.Run(); err != nil {
	// 	log.Fatalf("❌ install_miniconda.exe failed: %v", err)
	// }
	// log.Println("✅ Miniconda installation completed.")

	// // === Run install_sql_developer.exe ===
	// if _, err := os.Stat(installSQLDeveloperExe); err != nil {
	// 	log.Fatalf("❌ Could not find install_sql_developer.exe at: %s\n%v", installSQLDeveloperExe, err)
	// }
	// log.Printf("🧠 Running: %s\n", installSQLDeveloperExe)
	// cmd4 := exec.Command(installSQLDeveloperExe)
	// cmd4.Stdout = os.Stdout
	// cmd4.Stderr = os.Stderr
	// if err := cmd4.Run(); err != nil {
	// 	log.Fatalf("❌ install_sql_developer.exe failed: %v", err)
	// }
	// log.Println("✅ SQL Developer installation completed.")

	// === Run install_nirsoft.exe ===
	if _, err := os.Stat(installNirsoftExe); err != nil {
		log.Fatalf("❌ Could not find install_nirsoft.exe at: %s\n%v", installNirsoftExe, err)
	}
	log.Printf("🧰 Running: %s\n", installNirsoftExe)
	cmd5 := exec.Command(installNirsoftExe)
	cmd5.Stdout = os.Stdout
	cmd5.Stderr = os.Stderr
	if err := cmd5.Run(); err != nil {
		log.Fatalf("❌ install_nirsoft.exe failed: %v", err)
	}
	log.Println("✅ NirSoft installation completed.")
}
