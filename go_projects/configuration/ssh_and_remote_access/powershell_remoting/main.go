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
	const sftp_subsystem = `Subsystem   sftp    C:/Windows/System32/OpenSSH/sftp-server.exe`

	// Step 4: Read sshd_config
	fmt.Println("📂 Checking sshd_config for Subsystem lines...")
	file, err := os.Open(config_path)
	if err != nil {
		fmt.Printf("❌ Failed to open sshd_config: %v\n", err)
		return
	}
	defer file.Close()

	var lines []string
	var insertIndex = -1
	foundPWSH := false
	foundSFTP := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		lower := strings.ToLower(trimmed)

		if strings.HasPrefix(lower, "subsystem") && strings.Contains(lower, "powershell") {
			foundPWSH = true
		}
		if strings.HasPrefix(lower, "subsystem") && strings.Contains(lower, "sftp") {
			// if incorrect, we’ll replace it later
			if !strings.Contains(trimmed, "C:/Windows/System32/OpenSSH/sftp-server.exe") {
				continue
			}
			foundSFTP = true
		}
		if insertIndex == -1 && (strings.HasPrefix(trimmed, "AllowGroups") || strings.HasPrefix(trimmed, "Match ")) {
			insertIndex = len(lines)
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Error reading sshd_config: %v\n", err)
		return
	}

	// Step 5: Modify the file
	updated := false
	var newLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		lower := strings.ToLower(trimmed)

		if strings.HasPrefix(lower, "subsystem") && strings.Contains(lower, "sftp") &&
			!strings.Contains(trimmed, "C:/Windows/System32/OpenSSH/sftp-server.exe") {
			fmt.Println("✏️ Fixing incorrect SFTP subsystem entry...")
			newLines = append(newLines, sftp_subsystem)
			updated = true
			continue
		}
		newLines = append(newLines, line)
	}

	if !foundPWSH {
		fmt.Println("➕ Adding PowerShell SSH subsystem...")
		if insertIndex == -1 {
			insertIndex = len(newLines)
		}
		newLines = append(newLines[:insertIndex], append([]string{powershell_subsystem}, newLines[insertIndex:]...)...)
		updated = true
	} else {
		fmt.Println("ℹ️ PowerShell SSH subsystem already present.")
	}

	if !foundSFTP {
		fmt.Println("➕ Adding SFTP subsystem...")
		if insertIndex == -1 {
			insertIndex = len(newLines)
		}
		newLines = append(newLines[:insertIndex], append([]string{sftp_subsystem}, newLines[insertIndex:]...)...)
		updated = true
	} else {
		fmt.Println("ℹ️ Correct SFTP subsystem already present.")
	}

	if updated {
		fmt.Println("💾 Writing updated sshd_config...")
		content := strings.Join(newLines, "\r\n") + "\r\n"
		if err := os.WriteFile(config_path, []byte(content), 0644); err != nil {
			fmt.Printf("❌ Failed to write sshd_config: %v\n", err)
			return
		}
		fmt.Println("✅ sshd_config updated.")
	} else {
		fmt.Println("✅ sshd_config is already up to date.")
	}

	// Step 6: Restart sshd
	fmt.Println("🔄 Restarting sshd service...")
	if err := run_powershell_command("Restart-Service sshd"); err != nil {
		fmt.Printf("❌ Failed to restart sshd service: %v\n", err)
		return
	}
	fmt.Println("✅ sshd service restarted successfully.")
}
