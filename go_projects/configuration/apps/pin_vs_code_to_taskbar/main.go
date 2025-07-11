package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
)

func main() {
	// Step 1: Convert GitHub blob URL to raw
	raw_url, err := system_management_functions.Convert_blob_to_raw_github_url(
		"https://github.com/PeterCullenBurbery/pin-to-taskbar/blob/main/pin_to_taskbar.exe",
	)
	if err != nil {
		log.Fatalf("❌ Failed to convert URL: %v", err)
	}

	// Step 2: Define and create target directory
	base_dir := `C:\downloads\pin-to-taskbar`
	if err := os.MkdirAll(base_dir, 0755); err != nil {
		log.Fatalf("❌ Failed to create directory %s: %v", base_dir, err)
	}

	// Step 3: Define path to download into
	pinner_path := filepath.Join(base_dir, "pin_to_taskbar.exe")

	// Step 4: Download the pin_to_taskbar.exe
	fmt.Println("📥 Downloading pin_to_taskbar.exe...")
	if err := system_management_functions.Download_file(pinner_path, raw_url); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	fmt.Println("✅ Downloaded to:", pinner_path)

	// Step 5: Add the directory to PATH
	fmt.Println("➕ Adding to system PATH...")
	if err := system_management_functions.Add_to_path(base_dir); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}
	fmt.Println("✅ Added to system PATH.")

	// Step 6: Create VS Code desktop shortcut (for all users, maximized)
	vscode_path := `C:\Program Files\Microsoft VS Code\Code.exe`
	fmt.Println("📄 Creating all-users desktop shortcut to VS Code...")
	err = system_management_functions.Create_desktop_shortcut(
		vscode_path,
		"VSCode.lnk",
		"Visual Studio Code (Maximized)",
		3,    // 3 = Maximized
		true, // true = all users
	)
	if err != nil {
		log.Fatalf("❌ Failed to create shortcut: %v", err)
	}

	// Step 7: Locate the shortcut in public desktop and pin it
	public_desktop := filepath.Join(os.Getenv("PUBLIC"), "Desktop")
	shortcut := filepath.Join(public_desktop, "VSCode.lnk")
	fmt.Println("📌 Pinned all-users shortcut:", shortcut)

	cmd := exec.Command(pinner_path, shortcut)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Failed to pin shortcut to taskbar: %v", err)
	}
	fmt.Println("✅ VS Code shortcut (all users) pinned to taskbar.")

	// Step 8: Restart File Explorer to reflect the taskbar pin visually
	fmt.Println("🔁 Restarting File Explorer to reflect changes...")
	if err := system_management_functions.Restart_file_explorer(); err != nil {
		log.Fatalf("❌ Failed to restart Explorer: %v", err)
	}
	fmt.Println("✅ File Explorer restarted.")
}