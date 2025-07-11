try {
    Write-Host "🌐 Downloading configuration-003 repository from GitHub..."

    # Define GitHub ZIP URL
    $repoUrl = "https://github.com/PeterCullenBurbery/configuration-003/archive/refs/heads/main.zip"

    # Temporary ZIP path and extraction directory
    $tempZip = Join-Path $env:TEMP "configuration-003.zip"
    $extractDir = Join-Path $env:TEMP ("configuration-003-" + [guid]::NewGuid().ToString())

    # Ensure TLS 1.2 for older PowerShell versions
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

    # Download the ZIP file
    $webClient = New-Object System.Net.WebClient
    $webClient.DownloadFile($repoUrl, $tempZip)

    Write-Host "📂 Extracting ZIP to: $extractDir"
    Expand-Archive -Path $tempZip -DestinationPath $extractDir -Force

    # Find the unzipped root directory
    $unzippedRoot = Get-ChildItem -Path $extractDir | Where-Object { $_.PSIsContainer } | Select-Object -First 1
    if (-not $unzippedRoot) {
        throw "❌ Could not locate extracted root folder inside: $extractDir"
    }

    $repoPath = $unzippedRoot.FullName
    $orchestrationPath = Join-Path $repoPath "go_projects\orchestration\orchestration.exe"
    $pinVsCodePath = Join-Path $repoPath "go_projects\configuration\apps\pin_vs_code_to_taskbar\pin_vs_code_to_taskbar.exe"

    if (-not (Test-Path $orchestrationPath)) {
        throw "❌ orchestration.exe not found at expected location: $orchestrationPath"
    }

    Write-Host "🚀 Running orchestration.exe with repository path:"
    Write-Host "    $repoPath"
    & $orchestrationPath $repoPath
} catch {
    Write-Error "❌ Script failed: $_"
} finally {
    Write-Host "🧹 Temporary folder: $extractDir"
    # Optional cleanup:
    # Remove-Item -Recurse -Force $extractDir
}