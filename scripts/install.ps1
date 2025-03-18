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

# List contents of install directory
Write-Host "`nContents of install directory ($installDir):" -ForegroundColor Cyan
Get-ChildItem -Path $installDir -Recurse | ForEach-Object {
    Write-Host "- $($_.FullName)"
}

# Locate the extracted binary - look for the Windows executable name
$extractedBinary = "${cliName}_windows_${arch}.exe"
$binaryPath = "$installDir\$extractedBinary"

Write-Host "`nLooking for binary:" -ForegroundColor Cyan
Write-Host "- Expected binary name: $extractedBinary"
Write-Host "- Expected full path: $binaryPath"

if (!(Test-Path -Path $binaryPath)) {
    Write-Host "`nError: Extracted binary not found at: $binaryPath" -ForegroundColor Red
    Write-Host "Searching for any .exe files in install directory..."
    $exeFiles = Get-ChildItem -Path $installDir -Filter "*.exe" -Recurse
    if ($exeFiles) {
        Write-Host "Found these .exe files:"
        $exeFiles | ForEach-Object {
            Write-Host "- $($_.FullName)"
        }
    } else {
        Write-Host "No .exe files found in install directory"
    }
    exit 1
}

# Rename the binary to the CLI name
Write-Host "`nRenaming binary:" -ForegroundColor Cyan
Write-Host "- From: $binaryPath"
Write-Host "- To: $installDir\$cliName.exe"
Rename-Item -Path $binaryPath -NewName "$cliName.exe" -Force

# Add the install directory to PATH if not already present
$path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
if ($installDir -notin $path) {
    Write-Host "`nAdding $installDir to PATH..."
    [System.Environment]::SetEnvironmentVariable("Path", "$path;$installDir", [System.EnvironmentVariableTarget]::User)
}

# Clean up temporary files
Write-Host "`nCleaning up temporary files..."
Remove-Item -Path $tarFile -Force

Write-Host "`nInstallation completed successfully." -ForegroundColor Green
Write-Host "Restart your terminal or run 'refreshenv' if using Chocolatey."
Write-Host "Run '$cliName --help' to verify the installation."
