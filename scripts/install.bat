@echo off
setlocal

:: Define download URL
set URL=https://raw.githubusercontent.com/meyanksingh/sova-cli/meyank/InstallFix/scripts/install.ps1
set PS_SCRIPT=install.ps1

:: Check if curl exists, else use PowerShell to download
where curl >nul 2>nul
if %errorlevel% neq 0 (
    echo curl not found, using PowerShell to download...
    powershell -Command "Invoke-WebRequest -Uri '%URL%' -OutFile '%PS_SCRIPT%'"
) else (
    echo Downloading install.ps1 using curl...
    curl -fsSL %URL% -o %PS_SCRIPT%
)

:: Run the PowerShell script with Execution Policy Bypass
powershell -NoProfile -ExecutionPolicy Bypass -File %PS_SCRIPT%

del %PS_SCRIPT%

endlocal