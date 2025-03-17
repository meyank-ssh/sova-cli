package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	initTemplate string
	initForce    bool
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Long: `Initialize a new project with the specified name.
This will create a new directory with the project name and
set up the basic structure and files needed for your project.

Example:
  sova init my-awesome-project
  sova init my-awesome-project --template go-web
  sova init my-awesome-project --force`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := "my-project"
		if len(args) > 0 {
			projectName = args[0]
		}

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			PrintError("Failed to get current directory: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Current working directory: %s\n", cwd)

		// Create the project directory path
		projectDir := filepath.Join(cwd, projectName)
		fmt.Printf("Project directory will be: %s\n", projectDir)

		// Check if the directory already exists
		if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
			if !initForce {
				PrintError("Directory %s already exists. Use --force to overwrite.", projectName)
				os.Exit(1)
			}
			PrintWarning("Overwriting existing directory: %s", projectName)
			if err := os.RemoveAll(projectDir); err != nil {
				PrintError("Failed to remove existing directory: %v", err)
				os.Exit(1)
			}
		}

		PrintInfo("Initializing new project: %s", projectName)
		PrintInfo("Using template: %s", initTemplate)

		// Create project directory
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			PrintError("Failed to create project directory: %v", err)
			os.Exit(1)
		}
		fmt.Printf("Created project directory: %s\n", projectDir)

		// Create basic project structure
		dirs := []string{
			"cmd",
			"internal",
			"pkg",
			"api",
			"docs",
			"scripts",
			"test",
		}

		for _, dir := range dirs {
			dirPath := filepath.Join(projectDir, dir)
			err := os.MkdirAll(dirPath, 0755)
			if err != nil {
				PrintError("Failed to create directory %s: %v", dir, err)
				os.Exit(1)
			}
			fmt.Printf("Created directory: %s\n", dirPath)
		}

		// Create basic files
		files := map[string]string{
			"main.go": `package main

import "fmt"

func main() {
	fmt.Println("Hello from ` + projectName + `!")
}
`,
			"go.mod": `module ` + projectName + `

go 1.21
`,
			"README.md": `# ` + projectName + `

This project was generated using Sova CLI.

## Getting Started

1. Run the project:
   ` + "```" + `bash
   go run main.go
   ` + "```" + `

## Project Structure

- cmd/: Command line interfaces
- internal/: Private application code
- pkg/: Public libraries
- api/: API definitions
- docs/: Documentation
- scripts/: Build and maintenance scripts
- test/: Additional test files
`,
		}

		for filename, content := range files {
			filePath := filepath.Join(projectDir, filename)
			err := ioutil.WriteFile(filePath, []byte(content), 0644)
			if err != nil {
				PrintError("Failed to create file %s: %v", filename, err)
				os.Exit(1)
			}
			fmt.Printf("Created file: %s\n", filePath)
		}

		// Verify the project structure
		if err := filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(projectDir, path)
			if err != nil {
				return err
			}
			if rel == "." {
				return nil
			}
			if info.IsDir() {
				fmt.Printf("Directory: %s\n", rel)
			} else {
				fmt.Printf("File: %s (%d bytes)\n", rel, info.Size())
			}
			return nil
		}); err != nil {
			PrintError("Failed to verify project structure: %v", err)
		}

		PrintSuccess("Project %s initialized successfully!", projectName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Add flags specific to the init command
	initCmd.Flags().StringVarP(&initTemplate, "template", "t", "default", "Template to use for project initialization")
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "Force initialization even if directory exists")

	// Bind flags to viper
	viper.BindPFlag("init.template", initCmd.Flags().Lookup("template"))
	viper.BindPFlag("init.force", initCmd.Flags().Lookup("force"))
}
