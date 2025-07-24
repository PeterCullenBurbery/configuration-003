package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var config_urls = map[string]string{
	"ApplicationConfig.json": "https://github.com/PeterCullenBurbery/ShareX/blob/main/defaults/ApplicationConfig.json",
	"HotkeysConfig.json":     "https://github.com/PeterCullenBurbery/ShareX/blob/main/defaults/HotkeysConfig.json",
	"UploadersConfig.json":   "https://github.com/PeterCullenBurbery/ShareX/blob/main/defaults/UploadersConfig.json",
}

type hotkey_info struct {
	Hotkey string `json:"Hotkey"`
	Win    bool   `json:"Win"`
}

type hotkey_entry struct {
	Hotkey_info   hotkey_info     `json:"HotkeyInfo"`
	Task_settings json.RawMessage `json:"TaskSettings"`
}

type hotkeys_config struct {
	Hotkeys             []hotkey_entry `json:"Hotkeys"`
	Application_version string         `json:"ApplicationVersion"`
}

func stop_sharex_process() {
	log.Println("🛑 Attempting to stop any running ShareX processes...")

	cmd := exec.Command("powershell", "-NoProfile", "-Command",
		`Get-Process -Name ShareX -ErrorAction SilentlyContinue | Stop-Process -Force`)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		log.Printf("⚠️ Failed to stop ShareX process (it may not be running): %v", err)
	} else {
		log.Println("✅ ShareX process stopped (if it was running).")
	}
}

func get_documents_paths() ([]string, error) {
	cmd := exec.Command("powershell", "-NoProfile", "-Command",
		`Get-LocalUser | Where-Object { $_.Enabled -eq $true } | ForEach-Object { $d="C:\Users\$($_.Name)\Documents"; if (Test-Path $d) { $d } }`)
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Documents paths: %w", err)
	}

	var paths []string
	for _, line := range strings.Split(string(output), "\n") {
		path := filepath.Clean(strings.TrimSpace(line))
		if path == "" || !strings.HasPrefix(path, `C:\Users\`) {
			continue
		}
		paths = append(paths, path)
	}
	return paths, nil
}

func download_default_configs(target_dir string) error {
	err := os.MkdirAll(target_dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create ShareX config folder: %w", err)
	}

	for filename, blob_url := range config_urls {
		raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
		if err != nil {
			log.Printf("❌ Skipping %s: invalid blob URL (%s)", filename, blob_url)
			continue
		}
		dest := filepath.Join(target_dir, filename)
		log.Printf("⬇️ Downloading %s to %s", filename, dest)
		if err := system_management_functions.Download_file(dest, raw_url); err != nil {
			log.Printf("❌ Failed to download %s: %v", filename, err)
		} else {
			log.Printf("✅ Downloaded %s", filename)
		}
	}
	return nil
}

func patch_application_config(config_path string) error {
	original_data, err := os.ReadFile(config_path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	updated_data := original_data

	updates := []struct {
		json_path string
		value     interface{}
	}{
		{"DefaultTaskSettings.AfterUploadJob", "None"},
		{"DefaultTaskSettings.GeneralSettings.PlaySoundAfterCapture", false},
		{"DefaultTaskSettings.GeneralSettings.PlaySoundAfterUpload", false},
		{"DefaultTaskSettings.GeneralSettings.PlaySoundAfterAction", false},
		{"DefaultTaskSettings.GeneralSettings.ShowToastNotificationAfterTaskCompleted", false},
		{"DefaultTaskSettings.CaptureSettings.ShowCursor", false},
		{"DefaultTaskSettings.AdvancedSettings.RegionCaptureDisableAnnotation", true},
		{"Language", "en-US"},
		{"ShowUploadWarning", false},
	}

	for _, update := range updates {
		updated_data, err = sjson.SetBytes(updated_data, update.json_path, update.value)
		if err != nil {
			return fmt.Errorf("failed to update %s: %w", update.json_path, err)
		}
	}

	backup_path := config_path + ".bak"
	if err := os.WriteFile(backup_path, original_data, 0644); err != nil {
		return fmt.Errorf("failed to write backup: %w", err)
	}

	if err := os.WriteFile(config_path, updated_data, 0644); err != nil {
		return fmt.Errorf("failed to write updated config: %w", err)
	}

	log.Printf("✅ Patched %s and saved backup.", config_path)
	return nil
}

func patch_hotkeys_config(config_path string) error {
	original_data, err := os.ReadFile(config_path)
	if err != nil {
		return fmt.Errorf("failed to read hotkeys config: %w", err)
	}

	updated_data := original_data
	hotkeys := gjson.GetBytes(original_data, "Hotkeys")
	if !hotkeys.Exists() || !hotkeys.IsArray() {
		return fmt.Errorf("hotkeys array not found or invalid")
	}

	for i, item := range hotkeys.Array() {
		hotkey_val := item.Get("HotkeyInfo.Hotkey").String()
		switch hotkey_val {
		case "PrintScreen, Control":
			path := fmt.Sprintf("Hotkeys.%d.HotkeyInfo.Hotkey", i)
			updated_data, err = sjson.SetBytes(updated_data, path, "R, Shift, Alt")
		case "PrintScreen":
			path := fmt.Sprintf("Hotkeys.%d.HotkeyInfo.Hotkey", i)
			updated_data, err = sjson.SetBytes(updated_data, path, "F, Shift, Alt")
		}
		if err != nil {
			return fmt.Errorf("failed to update hotkey at index %d: %w", i, err)
		}
	}

	// Check for existing Shift+Alt+D
	var parsed hotkeys_config
	if err := json.Unmarshal(updated_data, &parsed); err != nil {
		return fmt.Errorf("failed to re-parse hotkey config: %w", err)
	}

	already_exists := false
	for _, entry := range parsed.Hotkeys {
		if entry.Hotkey_info.Hotkey == "D, Shift, Alt" && !entry.Hotkey_info.Win {
			already_exists = true
			log.Println("⚠️ Hotkey 'D, Shift, Alt' already exists. Skipping addition.")
			break
		}
	}

	if !already_exists {
		blob_url := "https://github.com/PeterCullenBurbery/ShareX/blob/main/new_folder/shift-alt-d.json"
		raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
		if err != nil {
			return fmt.Errorf("failed to convert blob URL: %w", err)
		}
		temp_path := filepath.Join(os.TempDir(), "shift-alt-d.json")
		if err := system_management_functions.Download_file(temp_path, raw_url); err != nil {
			return fmt.Errorf("failed to download new hotkey: %w", err)
		}
		new_data, err := os.ReadFile(temp_path)
		if err != nil {
			return fmt.Errorf("failed to read downloaded hotkey: %w", err)
		}
		var new_entry hotkey_entry
		if err := json.Unmarshal(new_data, &new_entry); err != nil {
			return fmt.Errorf("failed to parse downloaded hotkey JSON: %w", err)
		}
		parsed.Hotkeys = append(parsed.Hotkeys, new_entry)

		updated_data, err = json.MarshalIndent(parsed, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal patched hotkeys JSON: %w", err)
		}
	}

	backup_path := config_path + ".bak"
	if err := os.WriteFile(backup_path, original_data, 0644); err != nil {
		return fmt.Errorf("failed to write hotkeys backup: %w", err)
	}

	if err := os.WriteFile(config_path, updated_data, 0644); err != nil {
		return fmt.Errorf("failed to write updated hotkeys config: %w", err)
	}

	log.Printf("✅ Patched %s and saved backup.", config_path)
	return nil
}

func main() {
	const package_name = "sharex"

	log.Printf("📦 Installing %s using Chocolatey...", package_name)
	err := system_management_functions.Choco_install(package_name)
	if err != nil {
		log.Fatalf("❌ Failed to install %s: %v", package_name, err)
	}
	log.Printf("✅ %s installation completed successfully.", package_name)

	stop_sharex_process()

	doc_paths, err := get_documents_paths()
	if err != nil {
		log.Fatalf("❌ Failed to get Documents folders: %v", err)
	}

	for _, doc := range doc_paths {
		sharex_folder := filepath.Join(doc, "ShareX")
		log.Printf("📁 Processing: %s", sharex_folder)

		if err := download_default_configs(sharex_folder); err != nil {
			log.Printf("⚠️ Failed to set up defaults in %s: %v", sharex_folder, err)
			continue
		}

		app_config := filepath.Join(sharex_folder, "ApplicationConfig.json")
		if err := patch_application_config(app_config); err != nil {
			log.Printf("⚠️ Failed to patch ApplicationConfig: %v", err)
		}

		hotkey_config := filepath.Join(sharex_folder, "HotkeysConfig.json")
		if err := patch_hotkeys_config(hotkey_config); err != nil {
			log.Printf("⚠️ Failed to patch HotkeysConfig: %v", err)
		}
	}
}
