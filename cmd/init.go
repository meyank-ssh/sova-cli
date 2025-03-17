package cmd

import (
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

		// Create the project directory path
		projectDir := filepath.Join(cwd, projectName)

		// Check if the directory already exists
		if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
			if !initForce {
				PrintError("Directory %s already exists. Use --force to overwrite.", projectName)
				os.Exit(1)
			}
			PrintWarning("Overwriting existing directory: %s", projectName)
		}

		PrintInfo("Initializing new project: %s", projectName)
		PrintInfo("Using template: %s", initTemplate)

		// TODO: Implement project initialization logic
		// This would call functions from the internal/project package

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
