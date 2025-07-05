// To see VS Code settings, use:
// PowerShell: code $env:APPDATA\Code\User\settings.json

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Get VS Code settings.json path
	path, err := get_settings_path()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// Read existing settings.json if present
	var config map[string]interface{}
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			config = make(map[string]interface{}) // No file, start fresh
		} else {
			fmt.Fprintf(os.Stderr, "failed to read settings.json: %v\n", err)
			os.Exit(1)
		}
	} else {
		if err := json.Unmarshal(data, &config); err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse existing JSON: %v\n", err)
			os.Exit(1)
		}
	}

	// Backup original if it existed
	if data != nil {
		if err := make_backup(path, data); err != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to backup original settings.json: %v\n", err)
		}
	}

	// Get full Desktop path (e.g., C:\Users\peter\Desktop)
	desktop_path, err := get_resolved_desktop_path()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to determine desktop path: %v\n", err)
		os.Exit(1)
	}

	// Desired settings to insert/overwrite
	config["files.autoSave"] = "afterDelay"
	config["powershell.cwd"] = desktop_path
	config["terminal.integrated.cwd"] = desktop_path
	config["terminal.integrated.enableMultiLinePasteWarning"] = "never"
	config["terminal.integrated.persistentSessionScrollback"] = 10000000
	config["terminal.integrated.rightClickBehavior"] = "default"
	config["terminal.integrated.scrollback"] = 10000000
	config["workbench.startupEditor"] = "none"
	config["explorer.confirmDragAndDrop"] = false
	config["explorer.confirmDelete"] = false
	config["redhat.telemetry.enabled"] = true
	config["editor.renderWhitespace"] = "all"

	// YAML-specific editor settings
	config["[yaml]"] = map[string]interface{}{
		"editor.insertSpaces":      true,
		"editor.tabSize":           2,
		"editor.detectIndentation": false,
	}

	// Marshal JSON with indentation
	output, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal merged JSON: %v\n", err)
		os.Exit(1)
	}

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create directory %q: %v\n", dir, err)
		os.Exit(1)
	}

	// Atomic write: temp file -> rename
	tmp_path := path + ".tmp"
	if err := os.WriteFile(tmp_path, output, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write temp settings.json: %v\n", err)
		os.Exit(1)
	}
	if err := os.Rename(tmp_path, path); err != nil {
		fmt.Fprintf(os.Stderr, "failed to overwrite settings.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ VS Code settings.json updated successfully.")
}

func get_settings_path() (string, error) {
	appdata := os.Getenv("APPDATA")
	if appdata == "" {
		return "", errors.New("APPDATA environment variable not set")
	}
	return filepath.Join(appdata, "Code", "User", "settings.json"), nil
}

func make_backup(original_path string, data []byte) error {
	dir := filepath.Dir(original_path)
	base := filepath.Base(original_path)
	timestamp := time.Now().Format("20060102_150405")
	backup_name := fmt.Sprintf("%s.bak.%s", base, timestamp)
	backup_path := filepath.Join(dir, backup_name)
	return os.WriteFile(backup_path, data, 0o644)
}

func get_resolved_desktop_path() (string, error) {
	user_profile := os.Getenv("USERPROFILE")
	if user_profile == "" {
		return "", errors.New("USERPROFILE environment variable not set")
	}
	return filepath.Join(user_profile, "Desktop"), nil
}
