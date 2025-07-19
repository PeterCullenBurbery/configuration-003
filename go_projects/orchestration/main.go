package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func run_executable(label, exe_path string, args ...string) {
	if _, err := os.Stat(exe_path); err != nil {
		log.Fatalf("❌ %s not found at: %s\n%v", label, exe_path, err)
	}
	log.Printf("🚀 Launching %s: %s %v\n", label, exe_path, args)
	cmd := exec.Command(exe_path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ %s failed: %v", label, err)
	}
	log.Printf("✅ %s completed.\n", label)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: orchestration.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	base_dir := os.Args[1]

	install_things_path := filepath.Join(base_dir, "go_projects", "install", "install_things", "install_things.exe")
	configuration_path := filepath.Join(base_dir, "go_projects", "configuration", "configuration", "configuration.exe")
	configure_powershell_path := filepath.Join(base_dir, "go_projects", "powershell", "configure_powershell_modules", "configure_powershell_modules.exe")
	install_pip_packages_path := filepath.Join(base_dir, "go_projects", "install_pip_packages", "install_pip_packages.exe")

	run_executable("install_things.exe", install_things_path, base_dir)
	run_executable("configuration.exe", configuration_path, base_dir)
	run_executable("configure_powershell_modules.exe", configure_powershell_path, base_dir)
	run_executable("install_pip_packages.exe", install_pip_packages_path, base_dir)

	log.Println("🏁 orchestration.exe finished successfully.")
}