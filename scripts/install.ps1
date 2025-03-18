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

# Download the CLI archive
Write-Host "Downloading $cliName from $downloadUrl..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $tarFile -Headers @{"User-Agent"="Mozilla/5.0"}

if (!(Test-Path -Path $tarFile)) {
    Write-Host "Error: Download failed." -ForegroundColor Red
    exit 1
}

# Extract the .tar.gz file using tar
Write-Host "Extracting files..."
tar -xzf $tarFile -C $installDir

# Locate the extracted binary - look for the Windows executable name
$extractedBinary = "${cliName}_windows_${arch}.exe"
$binaryPath = "$installDir\$extractedBinary"

if (!(Test-Path -Path $binaryPath)) {
    Write-Host "Error: Extracted binary not found." -ForegroundColor Red
    exit 1
}

# Rename the binary to the CLI name
Rename-Item -Path $binaryPath -NewName "$cliName.exe" -Force

# Add the install directory to PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($installDir -notin $path) {
    Write-Host "Adding $installDir to PATH..."
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installDir", [System.EnvironmentVariableTarget]::User)
}

# Clean up temporary files
Remove-Item -Path $tarFile -Force

Write-Host "`nInstallation completed successfully." -ForegroundColor Green
Write-Host "Restart your terminal or run 'refreshenv' if using Chocolatey."
Write-Host "Run '$cliName --help' to verify the installation."
