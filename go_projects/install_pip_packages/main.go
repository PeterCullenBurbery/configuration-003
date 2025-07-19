package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: install_pip_packages.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	configuration_path := os.Args[1]
	yaml_file_path := filepath.Join(configuration_path, "python-packages.yaml")

	yaml_file, err := os.Open(yaml_file_path)
	if err != nil {
		log.Fatalf("❌ Could not open %s: %v", yaml_file_path, err)
	}
	defer yaml_file.Close()

	scanner := bufio.NewScanner(yaml_file)
	var package_list []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "-") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				package_list = append(package_list, fields[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("❌ Failed reading file: %v", err)
	}

	if len(package_list) == 0 {
		log.Println("ℹ️ No Python packages found in YAML.")
		return
	}

	// Try "pip" first
	pip_executable := "pip"

	// Check if "pip" exists in PATH
	_, err = exec.LookPath(pip_executable)
	if err != nil {
		// Fallback to known Miniconda path
		alt_pip := `C:\ProgramData\Miniconda3\Scripts\pip.exe`
		if _, stat_err := os.Stat(alt_pip); stat_err == nil {
			log.Printf("⚠️ 'pip' not in PATH. Falling back to: %s", alt_pip)
			pip_executable = alt_pip
		} else {
			log.Fatalf("❌ 'pip' not found and fallback pip.exe also missing.")
		}
	}

	pip_args := append([]string{"install"}, package_list...)
	pip_cmd := exec.Command(pip_executable, pip_args...)
	pip_cmd.Stdout = os.Stdout
	pip_cmd.Stderr = os.Stderr

	log.Printf("🐍 Installing Python packages: %v\n", package_list)
	if err := pip_cmd.Run(); err != nil {
		log.Fatalf("❌ pip install failed: %v", err)
	}

	log.Println("✅ Python packages installation completed.")
}