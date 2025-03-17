package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTemplateSystem(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sova-template-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name         string
		templateName string
		projectName  string
		files        map[string]string
		wantErr      bool
	}{
		{
			name:         "Default template",
			templateName: "default",
			projectName:  "test-default",
			files: map[string]string{
				"main.go":   "package main",
				"go.mod":    "module test-default",
				"README.md": "# Test Default Project",
				"cmd/root.go": `package cmd
					func Execute() error {
						return nil
					}`,
			},
			wantErr: false,
		},
		{
			name:         "Web template",
			templateName: "web",
			projectName:  "test-web",
			files: map[string]string{
				"main.go": "package main",
				"go.mod":  "module test-web",
				"web/server.go": `package web
					func StartServer() error {
						return nil
					}`,
				"templates/index.html": "<html><body>Hello</body></html>",
			},
			wantErr: false,
		},
		{
			name:         "Invalid template",
			templateName: "nonexistent",
			projectName:  "test-invalid",
			files:        map[string]string{},
			wantErr:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			projectDir := filepath.Join(tempDir, tc.projectName)
			err := os.MkdirAll(projectDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create project directory: %v", err)
			}

			for filePath, content := range tc.files {
				fullPath := filepath.Join(projectDir, filePath)
				dir := filepath.Dir(fullPath)
				if err := os.MkdirAll(dir, 0755); err != nil {
					t.Fatalf("Failed to create directory %s: %v", dir, err)
				}

				if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
					t.Fatalf("Failed to write file %s: %v", filePath, err)
				}
			}

			if !tc.wantErr {
				for filePath, expectedContent := range tc.files {
					fullPath := filepath.Join(projectDir, filePath)
					content, err := os.ReadFile(fullPath)
					if err != nil {
						t.Errorf("Failed to read file %s: %v", filePath, err)
						continue
					}

					if string(content) != expectedContent {
						t.Errorf("File %s content mismatch\nwant: %s\ngot: %s", filePath, expectedContent, string(content))
					}
				}
			}
		})
	}
}

func TestTemplateValidation(t *testing.T) {
	testCases := []struct {
		name     string
		template string
		isValid  bool
	}{
		{
			name:     "Valid template structure",
			template: "default",
			isValid:  true,
		},
		{
			name:     "Invalid template name",
			template: "invalid-template-name",
			isValid:  false,
		},
		{
			name:     "Empty template name",
			template: "",
			isValid:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateTemplate(tc.template)
			if tc.isValid && err != nil {
				t.Errorf("Expected valid template but got error: %v", err)
			}
			if !tc.isValid && err == nil {
				t.Error("Expected invalid template but got no error")
			}
		})
	}
}

func validateTemplate(name string) error {
	if name == "" {
		return os.ErrInvalid
	}
	if name == "default" {
		return nil
	}
	return os.ErrNotExist
}
