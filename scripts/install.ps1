# Set error handling
$ErrorActionPreference = "Stop"

# Define variables
$repoOwner = "go-sova"
$repoName = "sova-cli"
$cliName = "sova"
$arch = "amd64"
$installDir = "$env:LOCALAPPDATA\$cliName"

Write-Host "Installation Configuration:" -ForegroundColor Cyan
Write-Host "- Install Directory: $installDir"
Write-Host "- Architecture: $arch"
Write-Host "- CLI Name: $cliName"

# Ensure the install directory exists
if (!(Test-Path -Path $installDir)) {
    Write-Host "Creating installation directory: $installDir"
    New-Item -ItemType Directory -Path $installDir -Force | Out-Null
}

# Fetch the latest release tag from GitHub API
Write-Host "`nFetching latest release of $cliName..." -ForegroundColor Cyan
$latestRelease = (Invoke-RestMethod -Uri "https://api.github.com/repos/$repoOwner/$repoName/releases/latest" -Headers @{"User-Agent"="Mozilla/5.0"}).tag_name

if (!$latestRelease) {
    Write-Host "Error: Failed to retrieve latest release." -ForegroundColor Red
    exit 1
}

Write-Host "Latest release found: $latestRelease"

# Fix the file name to match your release files
$assetName = "${cliName}_windows_${arch}.tar.gz"
$downloadUrl = "https://github.com/$repoOwner/$repoName/releases/download/$latestRelease/$assetName"
$tarFile = "$env:TEMP\$assetName"

Write-Host "`nDownload Details:" -ForegroundColor Cyan
Write-Host "- Download URL: $downloadUrl"
Write-Host "- Local tar file path: $tarFile"

# Download the CLI archive
Write-Host "`nDownloading $cliName from $downloadUrl..."
try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $tarFile -Headers @{"User-Agent"="Mozilla/5.0"}
    Write-Host "Download completed successfully"
} catch {
    Write-Host "Error during download: $_" -ForegroundColor Red
    exit 1
}

if (!(Test-Path -Path $tarFile)) {
    Write-Host "Error: Download failed. File not found at: $tarFile" -ForegroundColor Red
    exit 1
}

Write-Host "Downloaded file size: $((Get-Item $tarFile).length) bytes"

# Extract the .tar.gz file using tar
Write-Host "`nExtracting files..." -ForegroundColor Cyan
Write-Host "Extracting from: $tarFile"
Write-Host "Extracting to: $installDir"

try {
    tar -xzf $tarFile -C $installDir
    Write-Host "Extraction completed"
} catch {
    Write-Host "Error during extraction: $_" -ForegroundColor Red
    exit 1
}

# First, clean up any ._ prefixed files
Write-Host "`nCleaning up macOS metadata files..." -ForegroundColor Cyan
Get-ChildItem -Path $installDir -Filter "._*" | ForEach-Object {
    Write-Host "Removing: $($_.FullName)"
    Remove-Item $_.FullName -Force
}

# List contents of install directory after cleanup
Write-Host "`nContents of install directory after cleanup ($installDir):" -ForegroundColor Cyan
Get-ChildItem -Path $installDir -Recurse | ForEach-Object {
    Write-Host "- $($_.FullName)"
}

# Now handle the correct executable
Write-Host "`nLocating and renaming executable..." -ForegroundColor Cyan
$targetExe = "sova_windows_amd64.exe"
$exeFile = Get-ChildItem -Path $installDir -Filter $targetExe | Select-Object -First 1

if ($exeFile) {
    Write-Host "Found executable: $($exeFile.FullName)"
    $targetPath = Join-Path $installDir "sova.exe"
    Write-Host "Renaming to: $targetPath"
    
    # Remove existing sova.exe if it exists
    if (Test-Path $targetPath) {
        Remove-Item $targetPath -Force
    }
    
    Move-Item -Path $exeFile.FullName -Destination $targetPath -Force
} else {
    Write-Host "Error: Could not find $targetExe in the extracted files" -ForegroundColor Red
    exit 1
}

# Add the install directory to PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($installDir -notin $path) {
    Write-Host "`nAdding installation directory to PATH: $installDir"
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installDir", [System.EnvironmentVariableTarget]::User)
    Write-Host "Successfully added to PATH"
}

# Clean up temporary files
Write-Host "`nCleaning up temporary files..."
Remove-Item -Path $tarFile -Force

Write-Host "`nInstallation completed successfully." -ForegroundColor Green
Write-Host "Restart your terminal or run 'refreshenv' if using Chocolatey."
Write-Host "Run 'sova --help' to verify the installation."
