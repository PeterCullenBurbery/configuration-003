package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func run_steps(base_dir string, category string, steps []struct {
	label    string
	exe_name string
}) {
	category_path := filepath.Join(base_dir, "go_projects", "configuration", category)

	for _, step := range steps {
		exe_dir := filepath.Join(category_path, step.exe_name[:len(step.exe_name)-4])
		exe_path := filepath.Join(exe_dir, step.exe_name)

		if _, err := os.Stat(exe_path); err != nil {
			log.Fatalf("❌ Could not find %s at: %s\n%v", step.exe_name, exe_path, err)
		}

		log.Printf("%s Running: %s\n", step.label, exe_path)
		cmd := exec.Command(exe_path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("❌ %s failed: %v", step.exe_name, err)
		}
		log.Printf("✅ %s completed.\n", step.label)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("❌ Usage: configuration.exe <path-to-configuration-003>")
		os.Exit(1)
	}
	base_dir := os.Args[1]

	explorer_steps := []struct {
		label    string
		exe_name string
	}{
		{"🌙 dark_mode", "dark_mode.exe"},
		{"📍 set_start_menu_to_left", "set_start_menu_to_left.exe"},
		{"📄 show_file_extensions", "show_file_extensions.exe"},
		{"🫥 show_hidden_files", "show_hidden_files.exe"},
		{"🔍 hide_search_box", "hide_search_box.exe"},
		{"⏱️ seconds_in_taskbar", "seconds_in_taskbar.exe"},
	}

	date_time_steps := []struct {
		label    string
		exe_name string
	}{
		{"🗓️ set_short_date_pattern", "set_short_date_pattern.exe"},
		{"📆 set_long_date_pattern", "set_long_date_pattern.exe"},
		{"⏰ set_time_pattern", "set_time_pattern.exe"},
		{"🕐 set_24_hour_format", "set_24_hour_format.exe"},
		{"📅 Set_first_day_of_week_Monday", "Set_first_day_of_week_Monday.exe"},
	}

	apps_steps := []struct {
		label    string
		exe_name string
	}{
		{"🛠️ configure_keyboard_shortcuts_for_vs_code", "configure_keyboard_shortcuts_for_vs_code.exe"},
		{"⚙️ configure_settings_for_vs_code", "configure_settings_for_vs_code.exe"},
		{"🪟 configure_settings_for_windows_terminal", "configure_settings_for_windows_terminal.exe"},
		{"📌 pin_vs_code_to_taskbar", "pin_vs_code_to_taskbar.exe"},
	}

	run_steps(base_dir, "explorer", explorer_steps)
	run_steps(base_dir, "date-time", date_time_steps)
	run_steps(base_dir, "apps", apps_steps)

	// // 🧩 Run install_vs_code_extensions.exe with path to vs-code-extensions.yaml
	// vs_code_yaml := filepath.Join(base_dir, "vs-code-extensions.yaml")
	// vs_code_ext_exe := filepath.Join(base_dir, "go_projects", "configuration", "apps", "install_vs_code_extensions", "install_vs_code_extensions.exe")

	// if _, err := os.Stat(vs_code_ext_exe); err != nil {
	// 	log.Fatalf("❌ Could not find install_vs_code_extensions.exe at: %s\n%v", vs_code_ext_exe, err)
	// }
	// if _, err := os.Stat(vs_code_yaml); err != nil {
	// 	log.Fatalf("❌ Could not find vs-code-extensions.yaml at: %s\n%v", vs_code_yaml, err)
	// }

	// log.Printf("🧩 Installing VS Code extensions using: %s %s\n", vs_code_ext_exe, vs_code_yaml)
	// cmd_vs_code := exec.Command(vs_code_ext_exe, vs_code_yaml)
	// cmd_vs_code.Stdout = os.Stdout
	// cmd_vs_code.Stderr = os.Stderr
	// if err := cmd_vs_code.Run(); err != nil {
	// 	log.Fatalf("❌ install_vs_code_extensions.exe failed: %v", err)
	// }
	// log.Println("✅ VS Code extensions installed.")

	// 🛡️ Run SSH setup
	ssh_exe := filepath.Join(base_dir, "go_projects", "configuration", "ssh", "ssh.exe")
	if _, err := os.Stat(ssh_exe); err != nil {
		log.Fatalf("❌ Could not find ssh.exe at: %s\n%v", ssh_exe, err)
	}
	log.Printf("🛡️ Running SSH setup: %s\n", ssh_exe)
	cmd_ssh := exec.Command(ssh_exe)
	cmd_ssh.Stdout = os.Stdout
	cmd_ssh.Stderr = os.Stderr
	if err := cmd_ssh.Run(); err != nil {
		log.Fatalf("❌ ssh.exe failed: %v", err)
	}
	log.Println("✅ SSH setup completed.")
}
