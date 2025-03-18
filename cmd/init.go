package cmd

import (
	"fmt"

	"github.com/go-sova/sova-cli/internal/project/api"
	"github.com/go-sova/sova-cli/internal/project/cli"
	"github.com/go-sova/sova-cli/pkg/questions"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project",
	Long: `Initialize a new project with the specified name.
This command will guide you through the project setup process.
If you don't provide a project name, you'll be prompted to enter one.
You can choose between different project types:
  - api: A Go API project with clean architecture
  - cli: A Go CLI project with clean architecture`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var projectType string
		var err error

		if len(args) > 0 {
			projectName = args[0]
		} else {
			projectName, err = questions.AskProjectName()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
		}

		projectType, err = questions.AskProjectType()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		switch projectType {
		case "api":
			apiCmd := api.InitCmd
			apiCmd.SetArgs([]string{projectName})
			err = apiCmd.Execute()
		case "cli":
			cliCmd := cli.InitCmd
			cliCmd.SetArgs([]string{projectName})
			err = cliCmd.Execute()
		default:
			err = fmt.Errorf("unsupported project type: %s", projectType)
		}

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
