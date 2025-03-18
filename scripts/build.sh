#!/bin/bash

VERSION="v0.1.1"
BINARY_NAME="sova-cli"
BUILD_DIR="dist"
BUILD_DATE=$(date -u '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD)
MODULE="github.com/go-sova/sova-cli/internal/version"

# Create build directory
mkdir -p $BUILD_DIR

# Build flags
BUILD_FLAGS=(
    "-X '${MODULE}.Version=${VERSION}'"
    "-X '${MODULE}.BuildDate=${BUILD_DATE}'"
    "-X '${MODULE}.GitCommit=${GIT_COMMIT}'"
)
BUILD_FLAGS_STR=$(IFS=' '; echo "${BUILD_FLAGS[*]}")

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}-${VERSION}-linux-amd64 ./cmd/main.go
GOOS=linux GOARCH=arm64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}-${VERSION}-linux-arm64 ./cmd/main.go
GOOS=darwin GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}-${VERSION}-darwin-amd64 ./cmd/main.go
GOOS=darwin GOARCH=arm64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}-${VERSION}-darwin-arm64 ./cmd/main.go
GOOS=windows GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}-${VERSION}-windows-amd64.exe ./cmd/main.go

# Create checksums
cd $BUILD_DIR
sha256sum * > checksums.txt
cd ..

echo "Build complete! Binaries are available in the $BUILD_DIR directory" 