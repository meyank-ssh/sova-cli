package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sova",
	Short: "Sova CLI - A tool for generating project boilerplate",
	Long: `Sova CLI is a powerful tool for initializing and generating 
project boilerplate code. It helps you quickly set up new projects
with predefined templates and structures.

Use 'sova init' to create a new project or 'sova generate' to 
generate specific components for your existing project.`,
	// This will make the root command show help when executed without subcommands
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sova.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "V", false, "display version information")

	// Bind flags to viper
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".sova" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sova")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}

// PrintSuccess prints a success message in green
func PrintSuccess(format string, a ...interface{}) {
	color.Green(format, a...)
}

// PrintInfo prints an info message in blue
func PrintInfo(format string, a ...interface{}) {
	color.Blue(format, a...)
}

// PrintWarning prints a warning message in yellow
func PrintWarning(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

// PrintError prints an error message in red
func PrintError(format string, a ...interface{}) {
	color.Red(format, a...)
}

// GetTemplatesDir returns the path to the templates directory
func GetTemplatesDir() string {
	// In a real-world scenario, this would be more sophisticated
	// For now, we'll assume the templates are in the same directory as the executable
	execPath, err := os.Executable()
	if err != nil {
		return "templates"
	}

	execDir := filepath.Dir(execPath)
	return filepath.Join(execDir, "templates")
}
