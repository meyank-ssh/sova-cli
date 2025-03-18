#!/bin/bash

VERSION="v0.1.1"
BINARY_NAME="sova"
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
GOOS=linux GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME} ./main.go
tar -czf $BUILD_DIR/${BINARY_NAME}_linux_amd64.tar.gz -C $BUILD_DIR ${BINARY_NAME}
rm $BUILD_DIR/${BINARY_NAME}

GOOS=darwin GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME} ./main.go
tar -czf $BUILD_DIR/${BINARY_NAME}_darwin_amd64.tar.gz -C $BUILD_DIR ${BINARY_NAME}
rm $BUILD_DIR/${BINARY_NAME}

GOOS=windows GOARCH=amd64 go build -ldflags "${BUILD_FLAGS_STR}" -o $BUILD_DIR/${BINARY_NAME}.exe ./main.go
tar -czf $BUILD_DIR/${BINARY_NAME}_windows_amd64.tar.gz -C $BUILD_DIR ${BINARY_NAME}.exe
rm $BUILD_DIR/${BINARY_NAME}.exe

# Create checksums
cd $BUILD_DIR
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    shasum -a 256 *.tar.gz > checksums.txt
else
    # Linux and others
    sha256sum *.tar.gz > checksums.txt
fi
cd ..

echo "Build complete! Archives are available in the $BUILD_DIR directory"