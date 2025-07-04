try {
    Write-Host "🌐 Downloading configuration-003 repository from GitHub..."

    # Define GitHub ZIP URL
    $repoUrl = "https://github.com/PeterCullenBurbery/configuration-003/archive/refs/heads/main.zip"

    # Temp ZIP path and extraction directory
    $tempZip = Join-Path $env:TEMP "configuration-003.zip"
    $extractDir = Join-Path $env:TEMP ("configuration-003-" + [guid]::NewGuid().ToString())

    # Ensure TLS 1.2 (for PowerShell 5.1 and below)
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

    # Download the ZIP file
    $webClient = New-Object System.Net.WebClient
    $webClient.DownloadFile($repoUrl, $tempZip)

    Write-Host "📂 Extracting ZIP to: $extractDir"
    Expand-Archive -Path $tempZip -DestinationPath $extractDir -Force

    # Find the unzipped root directory (GitHub zips include a folder like repo-main)
    $unzippedRoot = Get-ChildItem -Path $extractDir | Where-Object { $_.PSIsContainer } | Select-Object -First 1
    if (-not $unzippedRoot) {
        throw "❌ Failed to detect extracted folder inside $extractDir"
    }

    $repoPath = $unzippedRoot.FullName

    # Compute the full path to call_installer.exe
    $exePath = Join-Path $repoPath "go_projects\call_installer\call_installer.exe"

    if (-not (Test-Path $exePath)) {
        throw "❌ call_installer.exe not found at expected location: $exePath"
    }

    Write-Host "🚀 Running call_installer.exe with repository path: $repoPath"
    & $exePath $repoPath

} catch {
    Write-Error "❌ Script failed: $_"
} finally {
    Write-Host "🧹 Temporary folder: $extractDir"
    # Optional cleanup:
    # Remove-Item -Recurse -Force $extractDir
}