package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

func main() {
	// --- Step 1: Install Python ---

	blob_url := "https://github.com/PeterCullenBurbery/python-projects-semipublic/blob/main/install_python/install_python_007/dist/part-007.exe"
	raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
	if err != nil {
		log.Fatalf("❌ Failed to convert blob URL: %v", err)
	}

	temp_dir := os.TempDir()
	python_installer_path := filepath.Join(temp_dir, "part_007.exe")

	log.Printf("⬇️  Downloading Python installer to: %s\n", python_installer_path)
	err = system_management_functions.Download_file(python_installer_path, raw_url)
	if err != nil {
		log.Fatalf("❌ Failed to download Python installer: %v", err)
	}

	log.Println("🚀 Running Python installer...")
	python_install_cmd := exec.Command(python_installer_path)
	python_install_cmd.Stdout = os.Stdout
	python_install_cmd.Stderr = os.Stderr
	if err := python_install_cmd.Run(); err != nil {
		log.Fatalf("❌ Python installer failed: %v", err)
	}
	log.Println("✅ Python installation completed.")

	// --- Step 2: Download python-packages.yaml from GitHub ---

	yaml_blob_url := "https://github.com/PeterCullenBurbery/configuration-003/blob/main/python-packages.yaml"
	yaml_raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(yaml_blob_url)
	if err != nil {
		log.Fatalf("❌ Failed to convert YAML URL: %v", err)
	}

	yaml_temp_path := filepath.Join(temp_dir, "python-packages.yaml")
	log.Printf("⬇️  Downloading YAML file to: %s\n", yaml_temp_path)
	err = system_management_functions.Download_file(yaml_temp_path, yaml_raw_url)
	if err != nil {
		log.Fatalf("❌ Failed to download YAML file: %v", err)
	}

	// --- Step 3: Parse YAML file to extract package names ---

	yaml_file, err := os.Open(yaml_temp_path)
	if err != nil {
		log.Fatalf("❌ Could not open YAML file: %v", err)
	}
	defer yaml_file.Close()

	var package_list []string
	scanner := bufio.NewScanner(yaml_file)
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
		log.Fatalf("❌ Failed to read YAML file: %v", err)
	}

	if len(package_list) == 0 {
		log.Println("ℹ️ No packages found in YAML file.")
		return
	}

	// --- Step 4: Locate pip executable ---

	pip_executable := "pip"
	_, err = exec.LookPath(pip_executable)
	if err != nil {
		fallback_pip := `C:\Program Files\Python313\Scripts\pip.exe`
		if _, stat_err := os.Stat(fallback_pip); stat_err == nil {
			log.Printf("⚠️ 'pip' not found in PATH, using fallback: %s", fallback_pip)
			pip_executable = fallback_pip
		} else {
			log.Fatalf("❌ 'pip' not found and fallback pip.exe missing.")
		}
	}

	// --- Step 5: Install packages ---

	log.Printf("🐍 Installing Python packages: %v\n", package_list)
	pip_args := append([]string{"install"}, package_list...)
	pip_cmd := exec.Command(pip_executable, pip_args...)
	pip_cmd.Stdout = os.Stdout
	pip_cmd.Stderr = os.Stderr

	if err := pip_cmd.Run(); err != nil {
		log.Fatalf("❌ pip install failed: %v", err)
	}
	log.Println("✅ All Python packages installed successfully.")
}
