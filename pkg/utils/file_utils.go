package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// CreateDirIfNotExists creates a directory if it doesn't exist
func CreateDirIfNotExists(dirname string) error {
	if !DirExists(dirname) {
		return os.MkdirAll(dirname, os.ModePerm)
	}
	return nil
}

// WriteFile writes data to a file, creating the file if it doesn't exist
func WriteFile(filename string, data []byte) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExists(dir); err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// ReadFile reads the contents of a file
func ReadFile(filename string) ([]byte, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf("file not found: %s", filename)
	}
	return os.ReadFile(filename)
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	// Get source file info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Check if source is a regular file
	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination directory if it doesn't exist
	dstDir := filepath.Dir(dst)
	if err := CreateDirIfNotExists(dstDir); err != nil {
		return err
	}

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy file contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Set file permissions
	return os.Chmod(dst, srcInfo.Mode())
}

// CopyDir recursively copies a directory from src to dst
func CopyDir(src, dst string) error {
	// Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Check if source is a directory
	if !srcInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	// Create destination directory
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	// Read source directory entries
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Copy each entry
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectory
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy file
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetFileExtension returns the extension of a file
func GetFileExtension(filename string) string {
	return strings.TrimPrefix(filepath.Ext(filename), ".")
}

// GetFileNameWithoutExtension returns the filename without extension
func GetFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
}

// IsTextFile checks if a file is a text file
func IsTextFile(filename string) bool {
	// Common text file extensions
	textExtensions := map[string]bool{
		"txt":  true,
		"md":   true,
		"go":   true,
		"js":   true,
		"ts":   true,
		"html": true,
		"css":  true,
		"json": true,
		"yaml": true,
		"yml":  true,
		"xml":  true,
		"sh":   true,
		"bat":  true,
		"py":   true,
		"rb":   true,
		"java": true,
		"c":    true,
		"cpp":  true,
		"h":    true,
		"hpp":  true,
	}

	ext := GetFileExtension(filename)
	return textExtensions[ext]
}

// GetCurrentYear returns the current year as a string
func GetCurrentYear() string {
	return time.Now().Format("2006")
}
