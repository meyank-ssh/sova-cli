package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

func TestTemplateSystem(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "sova-template-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testCases := []struct {
		name         string
		templateName string
		content      string
		data         map[string]interface{}
		expected     string
		wantErr      bool
	}{
		{
			name:         "Valid Go template",
			templateName: "main.go.tpl",
			content: `package main

import "fmt"

func main() {
	fmt.Println("Hello from {{.ProjectName}}!")
}`,
			data: map[string]interface{}{
				"ProjectName": "test-project",
			},
			expected: `package main

import "fmt"

func main() {
	fmt.Println("Hello from test-project!")
}`,
			wantErr: false,
		},
		{
			name:         "Valid README template",
			templateName: "README.md.tpl",
			content: `# {{.ProjectName}}

This project was generated using Sova CLI.

## Getting Started

1. Run the project:
` + "```" + `bash
go run main.go
` + "```" + ``,
			data: map[string]interface{}{
				"ProjectName": "test-project",
			},
			expected: `# test-project

This project was generated using Sova CLI.

## Getting Started

1. Run the project:
` + "```" + `bash
go run main.go
` + "```" + ``,
			wantErr: false,
		},
		{
			name:         "Invalid template syntax",
			templateName: "invalid.tpl",
			content:      "Hello {{.ProjectName", // Missing closing brace
			data: map[string]interface{}{
				"ProjectName": "test-project",
			},
			expected: "",
			wantErr:  true,
		},
		{
			name:         "Missing required variable",
			templateName: "missing.tpl",
			content:      "Hello {{.NonexistentVar}}!",
			data: map[string]interface{}{
				"ProjectName": "test-project",
			},
			expected: "",
			wantErr:  true,
		},
		{
			name:         "Complex template with nested data",
			templateName: "complex.tpl",
			content: `# {{.ProjectName}}

{{if .HasTests}}
## Testing
To run tests:
` + "```" + `bash
go test ./...
` + "```" + `
{{end}}

{{range .Dependencies}}
- {{.Name}} v{{.Version}}
{{end}}`,
			data: map[string]interface{}{
				"ProjectName": "test-project",
				"HasTests":    true,
				"Dependencies": []struct {
					Name    string
					Version string
				}{
					{"cobra", "1.0.0"},
					{"viper", "1.7.0"},
				},
			},
			expected: `# test-project

## Testing
To run tests:
` + "```" + `bash
go test ./...
` + "```" + `

- cobra v1.0.0
- viper v1.7.0
`,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create template file
			templatePath := filepath.Join(tempDir, tc.templateName)
			err := os.WriteFile(templatePath, []byte(tc.content), 0644)
			if err != nil {
				t.Fatalf("Failed to create template file: %v", err)
			}

			// Verify template file exists
			if _, err := os.Stat(templatePath); os.IsNotExist(err) {
				t.Errorf("Template file %s was not created", tc.templateName)
			}

			// Parse and execute template
			tmpl, err := template.New(tc.templateName).Parse(tc.content)
			if tc.wantErr {
				if err == nil {
					// Try executing the template to catch execution errors
					var buf bytes.Buffer
					err = tmpl.Execute(&buf, tc.data)
				}
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("Failed to parse template: %v", err)
			}

			// Execute template
			var buf bytes.Buffer
			err = tmpl.Execute(&buf, tc.data)
			if err != nil {
				t.Fatalf("Failed to execute template: %v", err)
			}

			// Compare output
			if buf.String() != tc.expected {
				t.Errorf("Template output mismatch.\nWant:\n%s\nGot:\n%s", tc.expected, buf.String())
			}
		})
	}
}
