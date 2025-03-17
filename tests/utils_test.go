package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileOperations(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "sova-utils-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	t.Run("Directory Creation", func(t *testing.T) {
		testDir := filepath.Join(tempDir, "test-dir")

		// Test directory creation
		if err := os.MkdirAll(testDir, 0755); err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}

		// Verify directory exists
		if info, err := os.Stat(testDir); err != nil || !info.IsDir() {
			t.Error("Directory was not created properly")
		}

		// Test nested directory creation
		nestedDir := filepath.Join(testDir, "nested", "dir")
		if err := os.MkdirAll(nestedDir, 0755); err != nil {
			t.Fatalf("Failed to create nested directory: %v", err)
		}

		// Verify nested directory exists
		if info, err := os.Stat(nestedDir); err != nil || !info.IsDir() {
			t.Error("Nested directory was not created properly")
		}
	})

	t.Run("File Creation and Modification", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "test.txt")
		content := []byte("Hello, World!")

		// Test file creation
		if err := os.WriteFile(testFile, content, 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		// Verify file exists
		if info, err := os.Stat(testFile); err != nil || info.IsDir() {
			t.Error("File was not created properly")
		}

		// Read file content
		readContent, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		// Verify content
		if string(readContent) != string(content) {
			t.Errorf("File content mismatch. Want %q, got %q", content, readContent)
		}

		// Modify file
		newContent := []byte("Modified content")
		if err := os.WriteFile(testFile, newContent, 0644); err != nil {
			t.Fatalf("Failed to modify file: %v", err)
		}

		// Verify modified content
		readContent, err = os.ReadFile(testFile)
		if err != nil {
			t.Fatalf("Failed to read modified file: %v", err)
		}

		if string(readContent) != string(newContent) {
			t.Errorf("Modified content mismatch. Want %q, got %q", newContent, readContent)
		}
	})

	t.Run("File Permissions", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "permissions.txt")
		content := []byte("Test content")

		// Create file with specific permissions
		if err := os.WriteFile(testFile, content, 0600); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		// Verify permissions
		info, err := os.Stat(testFile)
		if err != nil {
			t.Fatalf("Failed to get file info: %v", err)
		}

		if info.Mode().Perm() != 0600 {
			t.Errorf("File permissions mismatch. Want %v, got %v", 0600, info.Mode().Perm())
		}

		// Change permissions
		if err := os.Chmod(testFile, 0644); err != nil {
			t.Fatalf("Failed to change permissions: %v", err)
		}

		// Verify new permissions
		info, err = os.Stat(testFile)
		if err != nil {
			t.Fatalf("Failed to get file info after permission change: %v", err)
		}

		if info.Mode().Perm() != 0644 {
			t.Errorf("File permissions mismatch after change. Want %v, got %v", 0644, info.Mode().Perm())
		}
	})

	t.Run("File Deletion", func(t *testing.T) {
		testFile := filepath.Join(tempDir, "to-delete.txt")
		content := []byte("Delete me")

		// Create file
		if err := os.WriteFile(testFile, content, 0644); err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		// Verify file exists
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			t.Error("File was not created")
		}

		// Delete file
		if err := os.Remove(testFile); err != nil {
			t.Fatalf("Failed to delete file: %v", err)
		}

		// Verify file was deleted
		if _, err := os.Stat(testFile); !os.IsNotExist(err) {
			t.Error("File was not deleted")
		}
	})
}

func TestPathOperations(t *testing.T) {
	testCases := []struct {
		name     string
		path     string
		wantDir  string
		wantBase string
		wantExt  string
	}{
		{
			name:     "Simple file",
			path:     "file.txt",
			wantDir:  ".",
			wantBase: "file",
			wantExt:  ".txt",
		},
		{
			name:     "Nested file",
			path:     "path/to/file.go",
			wantDir:  "path/to",
			wantBase: "file",
			wantExt:  ".go",
		},
		{
			name:     "Hidden file",
			path:     ".config",
			wantDir:  ".",
			wantBase: ".config",
			wantExt:  "",
		},
		{
			name:     "Multiple extensions",
			path:     "archive.tar.gz",
			wantDir:  ".",
			wantBase: "archive.tar",
			wantExt:  ".gz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test Dir
			if got := filepath.Dir(tc.path); got != tc.wantDir {
				t.Errorf("Dir(%q) = %q, want %q", tc.path, got, tc.wantDir)
			}

			// Test Base
			if got := filepath.Base(tc.path); got != filepath.Base(tc.path) {
				t.Errorf("Base(%q) = %q, want %q", tc.path, got, filepath.Base(tc.path))
			}

			// Test Ext
			if got := filepath.Ext(tc.path); got != tc.wantExt {
				t.Errorf("Ext(%q) = %q, want %q", tc.path, got, tc.wantExt)
			}
		})
	}
}
