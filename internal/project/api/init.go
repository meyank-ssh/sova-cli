package api

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "api [project-name]",
	Short: "Initialize a new Go API project",
	Long: `Initialize a new Go API project with a clean architecture structure.
This command will create a new directory with the project name and set up all necessary files and directories.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		projectDir := filepath.Join(".", projectName)

		if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
			return fmt.Errorf("directory %s already exists", projectDir)
		}

		if err := os.MkdirAll(projectDir, 0755); err != nil {
			return fmt.Errorf("failed to create project directory: %v", err)
		}

		answers, err := questions.AskProjectQuestions("api")
		if err != nil {
			return fmt.Errorf("failed to get project configuration: %v", err)
		}

		answers.ProjectName = projectName

		generator := NewAPIProjectGenerator(projectName, projectDir, answers)

		files, dirs, err := generator.Generate()
		if err != nil {
			return fmt.Errorf("failed to generate project files: %v", err)
		}

		for _, dir := range dirs {
			dirPath := filepath.Join(projectDir, dir)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %v", dir, err)
			}
			fmt.Printf("Created directory: %s\n", dirPath)
		}

		if err := generator.WriteFiles(files); err != nil {
			return fmt.Errorf("failed to write files: %v", err)
		}

		fmt.Printf("\nProject %s created successfully!\n", projectName)
		fmt.Println("\nNext steps:")
		fmt.Printf("cd %s\n", projectName)
		fmt.Println("go mod tidy")
		fmt.Println("docker compose up -d")
		fmt.Println("go run cmd/main.go")
		fmt.Println("\nYour API will be available at http://localhost:8080")
		fmt.Println("Test the ping endpoint: curl http://localhost:8080/api/ping")

		return nil
	},
}
