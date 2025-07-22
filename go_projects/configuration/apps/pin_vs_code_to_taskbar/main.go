package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/PeterCullenBurbery/go_functions_002/v4/system_management_functions"
)

// Create_desktop_shortcut creates a .lnk shortcut on the desktop.
// It accepts the target path, shortcut name (optional), description (optional),
// window style (3 = maximized), and allUsers flag.
func Create_desktop_shortcut(target_path, shortcut_name, description string, window_style int, all_users bool) error {
	// Ensure target exists
	if _, err := os.Stat(target_path); os.IsNotExist(err) {
		return fmt.Errorf("❌ Target path does not exist: %s", target_path)
	}

	// Determine desktop path
	var desktopPath string
	if all_users {
		public := os.Getenv("PUBLIC")
		desktopPath = filepath.Join(public, "Desktop")

		// Ensure the Public Desktop folder exists
		if _, err := os.Stat(desktopPath); os.IsNotExist(err) {
			fmt.Printf("📁 Public Desktop missing — creating: %s\n", desktopPath)
			if err := os.MkdirAll(desktopPath, 0755); err != nil {
				return fmt.Errorf("❌ Failed to create Public Desktop folder: %w", err)
			}
		}
	} else {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("❌ Could not determine current user: %w", err)
		}
		desktopPath = filepath.Join(usr.HomeDir, "Desktop")
	}

	// Determine shortcut name if empty
	if shortcut_name == "" {
		base := filepath.Base(target_path)
		shortcut_name = strings.TrimSuffix(base, filepath.Ext(base)) + ".lnk"
	}

	// Final shortcut path
	shortcutPath := filepath.Join(desktopPath, shortcut_name)

	// Log input info
	fmt.Printf("📄 Creating shortcut:\n")
	fmt.Printf("   📁 Target Path: %s\n", target_path)
	fmt.Printf("   📝 Shortcut Name: %s\n", shortcut_name)
	fmt.Printf("   📂 Shortcut Path: %s\n", shortcutPath)

	// Initialize COM
	if err := ole.CoInitialize(0); err != nil {
		return fmt.Errorf("❌ Failed to initialize COM: %w", err)
	}
	defer ole.CoUninitialize()

	// Create Shell COM object
	shell, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return fmt.Errorf("❌ Failed to create WScript.Shell COM object: %w", err)
	}
	defer shell.Release()

	dispatch, err := shell.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return fmt.Errorf("❌ Failed to get IDispatch: %w", err)
	}
	defer dispatch.Release()

	// Create the shortcut
	shortcutRaw, err := oleutil.CallMethod(dispatch, "CreateShortcut", shortcutPath)
	if err != nil {
		return fmt.Errorf("❌ Failed to create shortcut: %w", err)
	}
	shortcut := shortcutRaw.ToIDispatch()
	defer shortcut.Release()

	// Set properties
	if _, err := oleutil.PutProperty(shortcut, "TargetPath", target_path); err != nil {
		return fmt.Errorf("❌ Failed to set TargetPath: %w", err)
	}
	_, _ = oleutil.PutProperty(shortcut, "WorkingDirectory", filepath.Dir(target_path))
	_, _ = oleutil.PutProperty(shortcut, "WindowStyle", window_style)
	_, _ = oleutil.PutProperty(shortcut, "Description", description)
	_, _ = oleutil.PutProperty(shortcut, "IconLocation", fmt.Sprintf("%s, 0", target_path))

	// Save the shortcut
	if _, err := oleutil.CallMethod(shortcut, "Save"); err != nil {
		return fmt.Errorf("❌ Failed to save shortcut: %w", err)
	}

	fmt.Printf("✅ Shortcut created at: %s\n", shortcutPath)
	return nil
}

func main() {
	// Step 1: Create VS Code desktop shortcut (for all users, maximized)
	vscode_path := `C:\Program Files\Microsoft VS Code\Code.exe`
	fmt.Println("📄 Creating all-users desktop shortcut to VS Code...")
	err := Create_desktop_shortcut(
		vscode_path,
		"VSCode.lnk",
		"Visual Studio Code (Maximized)",
		3,    // 3 = Maximized
		true, // true = all users
	)
	if err != nil {
		log.Fatalf("❌ Failed to create shortcut: %v", err)
	}

	// Step 2: Convert GitHub blob to raw URL
	blob_url := "https://github.com/PeterCullenBurbery/python-projects-semipublic/blob/main/pin_vs_code_to_taskbar/pin_vs_code_to_taskbar_001/dist/part-001.exe"
	fmt.Println("🔗 Converting GitHub blob to raw URL...")
	raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(blob_url)
	if err != nil {
		log.Fatalf("❌ Failed to convert blob URL: %v", err)
	}

	// Step 3: Create temporary directory
	temp_dir, err := os.MkdirTemp("", "pin_vs_code_to_taskbar_*")
	if err != nil {
		log.Fatalf("❌ Failed to create temp directory: %v", err)
	}

	// Step 4: Download EXE into temp directory
	exe_path := filepath.Join(temp_dir, "pin_vs_code_to_taskbar.exe")
	fmt.Printf("📥 Downloading EXE to %s...\n", exe_path)
	if err := system_management_functions.Download_file(exe_path, raw_url); err != nil {
		log.Fatalf("❌ Failed to download EXE: %v", err)
	}

	// Step 5: Execute the EXE
	fmt.Println("🚀 Executing pin_vs_code_to_taskbar.exe...")
	cmd := exec.Command(exe_path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Execution failed: %v", err)
	}

	fmt.Println("✅ Finished successfully.")
}