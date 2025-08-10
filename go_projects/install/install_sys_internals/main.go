package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"

	"github.com/PeterCullenBurbery/go_functions_002/v6/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/v6/system_management_functions"
)

func main() {
	// Constants
	download_url := "https://download.sysinternals.com/files/SysinternalsSuite.zip"
	base_dir := `C:\downloads\sys-internals`
	zip_name := "SysinternalsSuite.zip"

	// ✅ Step 0: Ensure base directory exists
	if err := os.MkdirAll(base_dir, 0755); err != nil {
		log.Fatalf("❌ Failed to create base directory: %v", err)
	}

	// Step 1: Exclude from Defender
	if err := system_management_functions.Exclude_from_Microsoft_Windows_Defender(base_dir); err != nil {
		log.Fatalf("❌ Failed to exclude from Defender: %v", err)
	}

	// Step 2: Generate safe timestamp
	timestamp, err := date_time_functions.Date_time_stamp()
	if err != nil {
		log.Fatalf("❌ Failed to generate timestamp: %v", err)
	}
	safe_timestamp := date_time_functions.Safe_time_stamp(timestamp, 1)

	// Paths
	zip_path := filepath.Join(base_dir, zip_name)
	extract_dir := filepath.Join(base_dir, safe_timestamp)

	// Step 3: Download the ZIP file
	log.Printf("⬇️ Downloading to: %s", zip_path)
	if err := system_management_functions.Download_file(zip_path, download_url); err != nil {
		log.Fatalf("❌ Download failed: %v", err)
	}
	log.Println("✅ Download complete")

	// Step 4: Extract ZIP
	log.Printf("📦 Extracting to: %s", extract_dir)
	if err := system_management_functions.Extract_zip(zip_path, extract_dir); err != nil {
		log.Fatalf("❌ Extraction failed: %v", err)
	}
	log.Println("✅ Extraction complete")

	// Step 5: Add extract_dir to PATH
	log.Printf("➕ Adding to PATH: %s", extract_dir)
	if err := system_management_functions.Add_to_path(extract_dir); err != nil {
		log.Fatalf("❌ Failed to add to PATH: %v", err)
	}
	log.Println("✅ Sysinternals directory added to system PATH")

	// ✅ Step 6: Accept all Sysinternals EULAs
	if err := accept_all_sysinternals_eulas(); err != nil {
		log.Fatalf("❌ Failed to set EULAAccepted registry keys: %v", err)
	}
	log.Println("✅ EULA accepted for all Sysinternals tools")

	// ✅ Step 7: Copy and rename key utilities
	if err := copy_sysinternals_binaries(extract_dir); err != nil {
		log.Fatalf("❌ Failed to copy/rename executables: %v", err)
	}
	log.Println("✅ Copied and renamed key Sysinternals executables")
}

func accept_all_sysinternals_eulas() error {
	tool_names := []string{
		"Process Monitor", "Process Explorer", "PsExec", "PsFile", "PsGetSid", "PsInfo", "PsKill", "PsList",
		"PsLoggedon", "PsLogList", "PsPasswd", "PsPing", "PsService", "PsShutdown", "PsSuspend", "Autoruns",
		"AccessChk", "ADExplorer", "ADInsight", "Autologon", "BgInfo", "CacheSet", "ClockRes", "Contig",
		"CoreInfo", "CPUSTRES", "DbgView", "Desktops", "Disk2vhd", "DiskExt", "DiskMon", "DiskView", "Du",
		"EFSDump", "FindLinks", "Handle", "Hex2Dec", "Junction", "LDMDump", "ListDlls", "LiveKd", "LoadOrd",
		"LogonSessions", "MoveFile", "NotMyFault", "NTFSInfo", "PendMoves", "PipeList", "PortMon", "ProcDump",
		"RAMMap", "RegDelNull", "RegJump", "RU", "SDelete", "ShareEnum", "ShellRunAs", "SigCheck", "Streams",
		"Strings", "Sync", "Sysmon", "TcpView", "TestLimit", "VMMap", "VolumeID", "Whois", "WinObj", "ZoomIt",
	}

	for _, tool_name := range tool_names {
		key_path := `Software\Sysinternals\` + tool_name
		registry_key, _, err := registry.CreateKey(registry.CURRENT_USER, key_path, registry.SET_VALUE)
		if err != nil {
			return err
		}
		if err := registry_key.SetDWordValue("EulaAccepted", 1); err != nil {
			registry_key.Close()
			return err
		}
		registry_key.Close()
	}
	return nil
}

func copy_sysinternals_binaries(extract_dir string) error {
	copy_map := map[string][]string{
		"procexp64.exe": {
			"process_explorer.exe",
			"process_explorer64.exe",
		},
		"procmon64.exe": {
			"process_monitor.exe",
			"process_monitor64.exe",
		},
	}

	for source, destinations := range copy_map {
		source_path := filepath.Join(extract_dir, source)
		for _, dest := range destinations {
			dest_path := filepath.Join(extract_dir, dest)

			log.Printf("📁 Copying %s → %s", source, dest)
			if err := copy_file(source_path, dest_path); err != nil {
				return fmt.Errorf("failed to copy %s to %s: %w", source, dest, err)
			}
		}
	}
	return nil
}

func copy_file(src_path, dst_path string) error {
	src_file, err := os.Open(src_path)
	if err != nil {
		return err
	}
	defer src_file.Close()

	dst_file, err := os.Create(dst_path)
	if err != nil {
		return err
	}
	defer dst_file.Close()

	_, err = io.Copy(dst_file, src_file)
	return err
}
