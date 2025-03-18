# Set error handling
$ErrorActionPreference = "Stop"

# Define variables
$repoOwner = "go-sova"
$repoName = "sova-cli"
$cliName = "sova"
$arch = "amd64"
$installDir = "$env:LOCALAPPDATA\$cliName"

# Ensure the install directory exists
if (!(Test-Path -Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir -Force | Out-Null
}

# Fetch the latest release tag from GitHub API
Write-Host "Fetching latest release of $cliName..."
$latestRelease = (Invoke-RestMethod -Uri "https://api.github.com/repos/$repoOwner/$repoName/releases/latest").tag_name

if (!$latestRelease) {
    Write-Host "Error: Failed to retrieve latest release." -ForegroundColor Red
    exit 1
}

Write-Host "Latest release found: $latestRelease"

# Construct download URL
$assetName = "${cliName}_windows_${arch}.zip"
$downloadUrl = "https://github.com/$repoOwner/$repoName/releases/download/$latestRelease/$assetName"
$zipFile = "$env:TEMP\$assetName"

# Download the CLI archive
Write-Host "Downloading $cliName..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $zipFile

if (!(Test-Path -Path $zipFile)) {
    Write-Host "Error: Download failed." -ForegroundColor Red
    exit 1
}

# Extract the ZIP file
Write-Host "Extracting files..."
Expand-Archive -Path $zipFile -DestinationPath $installDir -Force

# Locate the extracted binary
$binaryPath = "$installDir\$cliName.exe"

if (!(Test-Path -Path $binaryPath)) {
    Write-Host "Error: Extracted binary not found." -ForegroundColor Red
    exit 1
}

# Add the install directory to PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($installDir -notin $path) {
    Write-Host "Adding $installDir to PATH..."
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installDir", [System.EnvironmentVariableTarget]::User)
}

# Clean up temporary files
Remove-Item -Path $zipFile -Force

Write-Host "`nInstallation completed successfully." -ForegroundColor Green
Write-Host "Restart your terminal or run 'refreshenv' if using Chocolatey."
Write-Host "Run '$cliName --help' to verify the installation."
