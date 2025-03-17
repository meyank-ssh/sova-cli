package tests

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLICommands(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sova-cli-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name          string
		args          []string
		expectedOut   string
		expectedError bool
	}{
		{
			name:          "Version command",
			args:          []string{"version"},
			expectedOut:   "Sova CLI version",
			expectedError: false,
		},
		{
			name:          "Help command",
			args:          []string{"help"},
			expectedOut:   "Available Commands:",
			expectedError: false,
		},
		{
			name:          "Init command with project name",
			args:          []string{"init", "test-project"},
			expectedOut:   "Project initialized successfully",
			expectedError: false,
		},
		{
			name:          "Invalid command",
			args:          []string{"invalid-command"},
			expectedOut:   "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "../main.go"}, tc.args...)...)
			cmd.Dir = tempDir

			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			output := stdout.String() + stderr.String()

			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !strings.Contains(output, tc.expectedOut) {
				t.Errorf("Expected output containing %q, got %q", tc.expectedOut, output)
			}
		})
	}
}

func TestCLIFlags(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "sova-cli-flags-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name          string
		args          []string
		flags         []string
		expectedOut   string
		expectedError bool
	}{
		{
			name:          "Init with template flag",
			args:          []string{"init", "test-project"},
			flags:         []string{"--template", "default"},
			expectedOut:   "Project initialized successfully",
			expectedError: false,
		},
		{
			name:          "Init with invalid template",
			args:          []string{"init", "test-project"},
			flags:         []string{"--template", "nonexistent"},
			expectedOut:   "",
			expectedError: true,
		},
		{
			name:          "Version with json flag",
			args:          []string{"version"},
			flags:         []string{"--json"},
			expectedOut:   "{",
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmdArgs := append([]string{"run", "../main.go"}, tc.args...)
			cmdArgs = append(cmdArgs, tc.flags...)
			cmd := exec.Command("go", cmdArgs...)
			cmd.Dir = tempDir

			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			output := stdout.String() + stderr.String()

			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !strings.Contains(output, tc.expectedOut) {
				t.Errorf("Expected output containing %q, got %q", tc.expectedOut, output)
			}
		})
	}
}
