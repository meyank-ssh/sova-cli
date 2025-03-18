package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CreateDirIfNotExists(dirname string) error {
	if !DirExists(dirname) {
		return os.MkdirAll(dirname, os.ModePerm)
	}
	return nil
}

func WriteFile(filename string, data []byte) error {
	dir := filepath.Dir(filename)
	if err := CreateDirIfNotExists(dir); err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ReadFile(filename string) ([]byte, error) {
	if !FileExists(filename) {
		return nil, fmt.Errorf("file not found: %s", filename)
	}
	return os.ReadFile(filename)
}

func CopyFile(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstDir := filepath.Dir(dst)
	if err := CreateDirIfNotExists(dstDir); err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", src)
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func GetFileExtension(filename string) string {
	return strings.TrimPrefix(filepath.Ext(filename), ".")
}

func GetFileNameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
}

func IsTextFile(filename string) bool {
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

func GetCurrentYear() string {
	return time.Now().Format("2006")
}
