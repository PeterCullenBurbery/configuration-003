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
		fmt.Println("❌ Usage: install_things.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]

	// Paths
	yamlPath := filepath.Join(baseDir, "what-to-install.yaml")
	projectsDir := filepath.Join(baseDir, "go_projects", "install")

	installSteps := []struct {
		label   string
		exeName string
		args    []string
	}{
		{"☕ install_java", "install_java.exe", nil},
		{"📦 install_packages", "install_packages.exe", []string{yamlPath}},
		{"🍒 install_cherry_tree", "install_cherry_tree.exe", nil},
		// {"🐍 install_miniconda", "install_miniconda.exe", nil},
		// {"🐍 install_Python", "install_Python.exe", nil},
		// {"🧠 install_sql_developer", "install_sql_developer.exe", nil},
		{"🧰 install_nirsoft", "install_nirsoft.exe", nil},
		// {"🔧 install_sys_internals", "install_sys_internals.exe", nil},
		{"📸 install_ShareX", "install_ShareX.exe", nil},
	}

	for _, step := range installSteps {
		exePath := filepath.Join(projectsDir, step.exeName[:len(step.exeName)-4], step.exeName)

		if _, err := os.Stat(exePath); err != nil {
			log.Fatalf("❌ Could not find %s at: %s\n%v", step.exeName, exePath, err)
		}

		log.Printf("%s Running: %s %v\n", step.label, exePath, step.args)
		cmd := exec.Command(exePath, step.args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("❌ %s failed: %v", step.exeName, err)
		}
		log.Printf("✅ %s completed.\n", step.label)
	}
}
