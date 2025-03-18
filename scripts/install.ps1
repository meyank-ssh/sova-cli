# Set error handling
$ErrorActionPreference = "Stop"

# Define variables
$repoOwner = "go-sova"
$repoName = "sova-cli"
$cliName = "sova"
$arch = "amd64"
$installDir = "$env:LOCALAPPDATA\$cliName"

Write-Host "Installing $cliName..." -ForegroundColor Cyan

# Ensure the install directory exists
if (!(Test-Path -Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir -Force | Out-Null
}

# Fetch the latest release tag from GitHub API
$latestRelease = (Invoke-RestMethod -Uri "https://api.github.com/repos/$repoOwner/$repoName/releases/latest" -Headers @{"User-Agent"="Mozilla/5.0"}).tag_name

if (!$latestRelease) {
    Write-Host "Error: Failed to retrieve latest release." -ForegroundColor Red
    exit 1
}

# Download the CLI archive
$assetName = "${cliName}_windows_${arch}.tar.gz"
$downloadUrl = "https://github.com/$repoOwner/$repoName/releases/download/$latestRelease/$assetName"
$tarFile = "$env:TEMP\$assetName"

try {
    Write-Host "Downloading latest version..." -ForegroundColor Cyan
    Invoke-WebRequest -Uri $downloadUrl -OutFile $tarFile -Headers @{"User-Agent"="Mozilla/5.0"}
} catch {
    Write-Host "Error: Download failed." -ForegroundColor Red
    exit 1
}

# Extract and set up the executable
try {
    Write-Host "Installing..." -ForegroundColor Cyan
    
    # Extract archive
    tar -xzf $tarFile -C $installDir
    
    # Clean up macOS metadata files
    Get-ChildItem -Path $installDir -Filter "._*" | Remove-Item -Force
    
    # Rename the executable
    $targetExe = "sova_windows_amd64.exe"
    $exeFile = Get-ChildItem -Path $installDir -Filter $targetExe | Select-Object -First 1
    
    if ($exeFile) {
        $targetPath = Join-Path $installDir "sova.exe"
        if (Test-Path $targetPath) {
            Remove-Item $targetPath -Force
        }
        Move-Item -Path $exeFile.FullName -Destination $targetPath -Force
    } else {
        throw "Installation failed: Required files not found."
    }
    
    # Add to PATH
    $path = [System.Environment]::GetEnvironmentVariable("Path", [System.EnvironmentVariableTarget]::User)
    if ($installDir -notin $path) {
        [System.Environment]::SetEnvironmentVariable("Path", "$path;$installDir", [System.EnvironmentVariableTarget]::User)
    }
    
    # Cleanup
    Remove-Item -Path $tarFile -Force
    
    Write-Host "`n$cliName installed successfully!" -ForegroundColor Green
    Write-Host "Please restart your terminal to use $cliName."
} catch {
    Write-Host "Error: Installation failed." -ForegroundColor Red
    exit 1
}
