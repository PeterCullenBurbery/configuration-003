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
	// Step 1: Create VS Code desktop shortcut (for all users, maximized)
	vscode_path := `C:\Program Files\Microsoft VS Code\Code.exe`
	fmt.Println("📄 Creating all-users desktop shortcut to VS Code...")
	err := system_management_functions.Create_desktop_shortcut(
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