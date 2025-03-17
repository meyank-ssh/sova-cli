package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	generateOutput string
	generateForce  bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [component]",
	Short: "Generate project components",
	Long: `Generate various components for your project.
This command can be used to create new files, modules, or other
components based on predefined templates.

Example:
  sova generate controller User
  sova generate model User
  sova generate api User --output ./api`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		componentType := args[0]
		componentName := ""

		if len(args) > 1 {
			componentName = args[1]
		} else {
			PrintError("Component name is required")
			cmd.Help()
			os.Exit(1)
		}

		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			PrintError("Failed to get current directory: %v", err)
			os.Exit(1)
		}

		// Determine the output directory
		outputDir := cwd
		if generateOutput != "" {
			if filepath.IsAbs(generateOutput) {
				outputDir = generateOutput
			} else {
				outputDir = filepath.Join(cwd, generateOutput)
			}
		}

		PrintInfo("Generating %s: %s", componentType, componentName)
		PrintInfo("Output directory: %s", outputDir)

		// TODO: Implement component generation logic
		// This would call functions from the internal/templates package

		PrintSuccess("%s %s generated successfully!", componentType, componentName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Add flags specific to the generate command
	generateCmd.Flags().StringVarP(&generateOutput, "output", "o", "", "Output directory for generated files")
	generateCmd.Flags().BoolVarP(&generateForce, "force", "f", false, "Force generation even if files exist")

	// Bind flags to viper
	viper.BindPFlag("generate.output", generateCmd.Flags().Lookup("output"))
	viper.BindPFlag("generate.force", generateCmd.Flags().Lookup("force"))
}
