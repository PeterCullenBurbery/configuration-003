package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func run_powershell_command(command string) error {
	const pwsh_path = `C:\Program Files\PowerShell\7\pwsh.exe`
	fmt.Printf("💻 Running PowerShell 7 command: %s\n", command)
	cmd := exec.Command(pwsh_path, "-Command", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Step 1: Enable PowerShell Remoting
	fmt.Println("🚀 Enabling PowerShell Remoting...")
	if err := run_powershell_command("Enable-PSRemoting -Force"); err != nil {
		fmt.Printf("❌ Failed to enable PowerShell remoting: %v\n", err)
		return
	}

	// Step 2: Set TrustedHosts
	fmt.Println("🔧 Setting WSMan TrustedHosts to '*'...")
	if err := run_powershell_command("Set-Item WSMan:\\localhost\\Client\\TrustedHosts -Value '*' -Force"); err != nil {
		fmt.Printf("❌ Failed to set TrustedHosts: %v\n", err)
		return
	}

	// Step 3: Define constants
	const config_path = `C:\ProgramData\ssh\sshd_config`
	const powershell_subsystem = `Subsystem   powershell   "C:\Program Files\PowerShell\7\pwsh.exe" -sshs -NoLogo`

	// Step 4: Check sshd_config for subsystem entry and insertion point
	fmt.Println("📂 Checking sshd_config for PowerShell SSH subsystem...")

	file, err := os.Open(config_path)
	if err != nil {
		fmt.Printf("❌ Failed to open sshd_config: %v\n", err)
		return
	}
	defer file.Close()

	var lines []string
	found := false
	insertIndex := -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		lines = append(lines, line)

		if strings.Contains(trimmed, "Subsystem") && strings.Contains(trimmed, "powershell") {
			found = true
		}

		// Find first suitable insertion point (before AllowGroups or Match block)
		if insertIndex == -1 && (strings.HasPrefix(trimmed, "AllowGroups") || strings.HasPrefix(trimmed, "Match ")) {
			insertIndex = len(lines) - 1
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Error reading sshd_config: %v\n", err)
		return
	}

	// Step 5: Insert if not found
	if !found {
		fmt.Println("➕ Adding PowerShell SSH subsystem to sshd_config...")

		if insertIndex == -1 {
			insertIndex = len(lines) // if no Match/AllowGroups found, append
		}

		lines = append(lines[:insertIndex], append([]string{powershell_subsystem}, lines[insertIndex:]...)...)
		content := strings.Join(lines, "\r\n") + "\r\n"

		err = os.WriteFile(config_path, []byte(content), 0644)
		if err != nil {
			fmt.Printf("❌ Failed to write sshd_config: %v\n", err)
			return
		}
		fmt.Println("✅ PowerShell SSH subsystem added.")
	} else {
		fmt.Println("ℹ️ PowerShell SSH subsystem already present. No changes needed.")
	}


	// Step 6: Restart sshd service
	fmt.Println("🔄 Restarting sshd service...")
	if err := run_powershell_command("Restart-Service sshd"); err != nil {
		fmt.Printf("❌ Failed to restart sshd service: %v\n", err)
		return
	}
	fmt.Println("✅ sshd service restarted successfully.")
}
