package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLICommands(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "sova-cmd-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Build the CLI binary for testing
	binaryPath := filepath.Join(tempDir, "sova")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, "../main.go")
	if output, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build CLI: %v\nOutput: %s", err, string(output))
	}

	testCases := []struct {
		name        string
		args        []string
		expectedOut string
		expectedErr string
		wantErr     bool
		checkOutput bool
	}{
		{
			name:        "Version command",
			args:        []string{"version"},
			expectedOut: "Sova CLI v",
			checkOutput: true,
			wantErr:     false,
		},
		{
			name:        "Help command",
			args:        []string{"help"},
			expectedOut: "Sova CLI - A tool for generating project boilerplate",
			checkOutput: true,
			wantErr:     false,
		},
		{
			name:        "Invalid command",
			args:        []string{"invalid-command"},
			expectedErr: "unknown command",
			wantErr:     true,
		},
		{
			name:        "Init without project name",
			args:        []string{"init"},
			expectedOut: "my-project", // Should use default project name
			checkOutput: true,
			wantErr:     false,
		},
		{
			name:        "Init with invalid template",
			args:        []string{"init", "test-project", "--template", "invalid-template"},
			expectedErr: "unknown template",
			wantErr:     true,
		},
		{
			name:        "Generate without component",
			args:        []string{"generate"},
			expectedErr: "requires at least 1 arg",
			wantErr:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new command
			cmd := exec.Command(binaryPath, tc.args...)
			cmd.Dir = tempDir

			// Run the command and capture output
			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			// Check error condition
			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if !strings.Contains(outputStr, tc.expectedErr) {
					t.Errorf("Expected error containing %q, got %q", tc.expectedErr, outputStr)
				}
				return
			}

			// Check success condition
			if err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, outputStr)
				return
			}

			// Check output if required
			if tc.checkOutput && !strings.Contains(outputStr, tc.expectedOut) {
				t.Errorf("Expected output containing %q, got %q", tc.expectedOut, outputStr)
			}

			// Additional checks for specific commands
			switch tc.args[0] {
			case "init":
				// Check if project directory was created
				projectName := "my-project"
				if len(tc.args) > 1 {
					projectName = tc.args[1]
				}
				projectDir := filepath.Join(tempDir, projectName)
				if _, err := os.Stat(projectDir); os.IsNotExist(err) {
					t.Errorf("Project directory was not created: %s", projectDir)
				}
			case "generate":
				if len(tc.args) > 2 {
					// Check if generated file exists
					componentPath := filepath.Join(tempDir, strings.ToLower(tc.args[2])+".go")
					if _, err := os.Stat(componentPath); os.IsNotExist(err) {
						t.Errorf("Generated component file was not created: %s", componentPath)
					}
				}
			}
		})
	}
}

func TestCLIFlags(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "sova-flag-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Build the CLI binary for testing
	binaryPath := filepath.Join(tempDir, "sova")
	buildCmd := exec.Command("go", "build", "-o", binaryPath, "../main.go")
	if output, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build CLI: %v\nOutput: %s", err, string(output))
	}

	testCases := []struct {
		name        string
		args        []string
		expectedOut string
		wantErr     bool
	}{
		{
			name:        "Version with verbose flag",
			args:        []string{"version", "--verbose"},
			expectedOut: "Go Version:",
			wantErr:     false,
		},
		{
			name:        "Init with force flag",
			args:        []string{"init", "test-project", "--force"},
			expectedOut: "Overwriting existing directory",
			wantErr:     false,
		},
		{
			name:        "Init with template flag",
			args:        []string{"init", "test-project", "--template", "go-web"},
			expectedOut: "Using template: go-web",
			wantErr:     false,
		},
		{
			name:        "Global verbose flag",
			args:        []string{"--verbose", "init", "test-project"},
			expectedOut: "debug",
			wantErr:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new command
			cmd := exec.Command(binaryPath, tc.args...)
			cmd.Dir = tempDir

			// Run the command and capture output
			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, outputStr)
				return
			}

			if !strings.Contains(outputStr, tc.expectedOut) {
				t.Errorf("Expected output containing %q, got %q", tc.expectedOut, outputStr)
			}
		})
	}
}
