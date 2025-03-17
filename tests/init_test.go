package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestProjectInitialization(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sova-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name        string
		projectName string
		template    string
		wantDirs    []string
		wantFiles   []string
		wantErr     bool
	}{
		{
			name:        "Basic project creation",
			projectName: "test-project",
			template:    "default",
			wantDirs: []string{
				"cmd",
				"internal",
				"pkg",
				"api",
				"docs",
				"scripts",
				"test",
			},
			wantFiles: []string{
				"main.go",
				"go.mod",
				"README.md",
			},
			wantErr: false,
		},
		{
			name:        "Web project creation",
			projectName: "web-project",
			template:    "go-web",
			wantDirs: []string{
				"cmd",
				"internal",
				"pkg",
				"api",
				"docs",
				"scripts",
				"test",
				"web",
				"templates",
				"static",
			},
			wantFiles: []string{
				"main.go",
				"go.mod",
				"README.md",
				"web/handlers.go",
				"web/middleware.go",
				"web/routes.go",
			},
			wantErr: false,
		},
		{
			name:        "CLI project creation",
			projectName: "cli-project",
			template:    "cli",
			wantDirs: []string{
				"cmd",
				"internal",
				"pkg",
				"docs",
			},
			wantFiles: []string{
				"main.go",
				"go.mod",
				"README.md",
				"cmd/root.go",
				"cmd/version.go",
			},
			wantErr: false,
		},
		{
			name:        "Invalid template",
			projectName: "invalid-project",
			template:    "nonexistent",
			wantDirs:    []string{},
			wantFiles:   []string{},
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			projectDir := filepath.Join(tempDir, tc.projectName)
			cmd := exec.Command("go", "run", "../main.go", "init", tc.projectName, "--template", tc.template)
			cmd.Dir = tempDir
			output, err := cmd.CombinedOutput()

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none. Output: %s", string(output))
				}
				return
			} else if err != nil {
				t.Fatalf("Failed to run command: %v\nOutput: %s", err, string(output))
			}

			for _, dir := range tc.wantDirs {
				dirPath := filepath.Join(projectDir, dir)
				if _, err := os.Stat(dirPath); os.IsNotExist(err) {
					t.Errorf("Expected directory %s does not exist", dir)
				}
			}

			for _, file := range tc.wantFiles {
				filePath := filepath.Join(projectDir, file)
				info, err := os.Stat(filePath)
				if os.IsNotExist(err) {
					t.Errorf("Expected file %s does not exist", file)
					continue
				}
				if info.Size() == 0 {
					t.Errorf("File %s exists but is empty", file)
				}
			}

			if !tc.wantErr {
				buildCmd := exec.Command("go", "build", "./...")
				buildCmd.Dir = projectDir
				if output, err := buildCmd.CombinedOutput(); err != nil {
					t.Errorf("Project failed to build: %v\nOutput: %s", err, string(output))
				}
			}
		})
	}
}
