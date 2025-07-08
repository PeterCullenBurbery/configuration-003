package main

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"

	"github.com/PeterCullenBurbery/go_functions_002/v3/date_time_functions"
	"github.com/PeterCullenBurbery/go_functions_002/v3/system_management_functions"
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
