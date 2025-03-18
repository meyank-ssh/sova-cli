#!/bin/bash

set -e  # Stop script on error

OS_TYPE=$(uname -s 2>/dev/null || echo "Windows")  
ARCH="amd64"

case "$OS_TYPE" in
    Linux*) OS="linux";;
    Darwin*) OS="darwin";;
    CYGWIN*|MINGW*|MSYS*) OS="windows";;
    Windows) OS="windows";;
    *) echo "Unsupported OS: $OS_TYPE"; exit 1;;
esac

if [ "$OS" == "windows" ]; then
    echo "Detected Windows: Please use the Windows installer (install.bat or install.ps1)"
    exit 1
fi

REPO_OWNER="go-sova"
REPO_NAME="sova-cli"
CLI_NAME="sova"
INSTALL_DIR="/usr/local/bin"

echo "Detected OS: $OS"
echo "Fetching latest release of $CLI_NAME..."

LATEST_RELEASE=$(curl -fsSL "https://api.github.com/repos/$REPO_OWNER/$REPO_NAME/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
    echo "Error: Failed to get the latest release version."
    exit 1
fi

echo "Latest release found: $LATEST_RELEASE"

ASSET_NAME="${CLI_NAME}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO_OWNER/$REPO_NAME/releases/download/$LATEST_RELEASE/$ASSET_NAME"

echo "Downloading $CLI_NAME from $DOWNLOAD_URL..."
curl -fsSL -o "$ASSET_NAME" "$DOWNLOAD_URL"

if [ ! -f "$ASSET_NAME" ]; then
    echo "Error: Download failed."
    exit 1
fi

echo "Extracting files..."
tar -xzf "$ASSET_NAME"

EXTRACTED_BINARY="${CLI_NAME}_${OS}_${ARCH}" 

if [ ! -f "$EXTRACTED_BINARY" ]; then
    echo "Error: Extracted binary not found."
    rm -f "$ASSET_NAME"
    exit 1
fi

mv "$EXTRACTED_BINARY" "$CLI_NAME"

echo "Installing $CLI_NAME to $INSTALL_DIR..."
chmod +x "$CLI_NAME"
sudo mv "$CLI_NAME" "$INSTALL_DIR/$CLI_NAME"

rm -f "$ASSET_NAME"

echo "Installation completed successfully."
echo "Run '$CLI_NAME --help' to verify the installation."
