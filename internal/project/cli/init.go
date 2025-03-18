package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "cli [project-name]",
	Short: "Initialize a new Go CLI project",
	Long: `Initialize a new Go CLI project with a clean architecture structure.
This command will create a new directory with the project name and set up all necessary files and directories.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectDir := filepath.Join(".", projectName)

		if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
			fmt.Printf("Error: directory %s already exists\n", projectDir)
			return
		}

		if err := os.MkdirAll(projectDir, 0755); err != nil {
			fmt.Printf("Error: failed to create project directory: %v\n", err)
			return
		}

		answers, err := questions.AskProjectQuestions("cli")
		if err != nil {
			fmt.Printf("Error: failed to get project configuration: %v\n", err)
			return
		}

		answers.ProjectName = projectName

		generator := NewCLIProjectGenerator(projectName, projectDir, answers)

		files, dirs, err := generator.Generate()
		if err != nil {
			fmt.Printf("Error: failed to generate project files: %v\n", err)
			return
		}

		for _, dir := range dirs {
			dirPath := filepath.Join(projectDir, dir)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				fmt.Printf("Error: failed to create directory %s: %v\n", dir, err)
				return
			}
			fmt.Printf("Created directory: %s\n", dirPath)
		}

		if err := generator.WriteFiles(files); err != nil {
			fmt.Printf("Error: failed to write files: %v\n", err)
			return
		}

		fmt.Printf("\nProject %s created successfully!\n", projectName)
		fmt.Println("\nNext steps:")
		fmt.Printf("1. cd %s\n", projectName)
		fmt.Println("2. go mod tidy")
		fmt.Println("3. go run main.go")
		fmt.Println("\nTry your CLI commands:")
		fmt.Printf("   ./%s command1\n", projectName)
		fmt.Printf("   ./%s command2\n", projectName)
	},
}
