package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type keybinding struct {
	Key     string `json:"key"`
	Command string `json:"command"`
	When    string `json:"when,omitempty"`
}

func main() {
	app_data := os.Getenv("APPDATA")
	if app_data == "" {
		fmt.Println("❌ APPDATA environment variable not set.")
		return
	}

	keybindings_path := filepath.Join(app_data, "Code", "User", "keybindings.json")

	var existing_bindings []keybinding

	if _, err := os.Stat(keybindings_path); err == nil {
		data, err := os.ReadFile(keybindings_path)
		if err == nil && len(data) > 0 {
			if err := json.Unmarshal(data, &existing_bindings); err != nil {
				fmt.Printf("❌ Failed to parse existing keybindings.json: %v\n", err)
				return
			}
		}
	} else if os.IsNotExist(err) {
		dir := filepath.Dir(keybindings_path)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Printf("❌ Failed to create directory %s: %v\n", dir, err)
			return
		}
		fmt.Println("ℹ️ keybindings.json does not exist, creating a new one.")
	} else {
		fmt.Printf("❌ Failed to stat keybindings.json: %v\n", err)
		return
	}

	new_bindings := []keybinding{
		{
			Key:     "ctrl+a",
			Command: "workbench.action.terminal.selectAll",
			When:    "terminalFocus",
		},
		{
			Key:     "ctrl+shift+a",
			Command: "workbench.action.terminal.copySelectionAsHtml",
			When:    "terminalFocus",
		},
	}

	for _, new_b := range new_bindings {
		found := false
		for _, existing_b := range existing_bindings {
			if existing_b.Key == new_b.Key && existing_b.Command == new_b.Command && existing_b.When == new_b.When {
				found = true
				break
			}
		}
		if !found {
			existing_bindings = append(existing_bindings, new_b)
		}
	}

	output, err := json.MarshalIndent(existing_bindings, "", "    ")
	if err != nil {
		fmt.Printf("❌ Failed to marshal keybindings: %v\n", err)
		return
	}

	if err := os.WriteFile(keybindings_path, output, 0o644); err != nil {
		fmt.Printf("❌ Failed to write keybindings.json: %v\n", err)
		return
	}

	fmt.Printf("✅ Successfully updated: %s\n", keybindings_path)
}
