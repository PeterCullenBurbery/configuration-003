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
		fmt.Println("❌ Usage: configuration.exe <path-to-configuration-003>")
		os.Exit(1)
	}

	baseDir := os.Args[1]
	configRoot := filepath.Join(baseDir, "go_projects", "configuration")

	steps := []struct {
		label   string
		exeName string
	}{
		{"🌙 dark_mode", "dark_mode.exe"},
		{"📍 set_start_menu_to_left", "set_start_menu_to_left.exe"},
		{"📄 show_file_extensions", "show_file_extensions.exe"},
		{"🫥 show_hidden_files", "show_hidden_files.exe"},
		{"🔍 hide_search_box", "hide_search_box.exe"},
		{"⏱️ seconds_in_taskbar", "seconds_in_taskbar.exe"},
	}

	for _, step := range steps {
		// Each executable is in a subfolder matching its name (minus .exe)
		exeDir := filepath.Join(configRoot, step.exeName[:len(step.exeName)-4])
		exePath := filepath.Join(exeDir, step.exeName)

		if _, err := os.Stat(exePath); err != nil {
			log.Fatalf("❌ Could not find %s at: %s\n%v", step.exeName, exePath, err)
		}

		log.Printf("%s Running: %s\n", step.label, exePath)
		cmd := exec.Command(exePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("❌ %s failed: %v", step.exeName, err)
		}
		log.Printf("✅ %s completed.\n", step.label)
	}
}
