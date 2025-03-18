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

var rootCmd = &cobra.Command{
	Use:   "sova",
	Short: "Sova CLI - A tool for initializing projects",
	Long: `Sova CLI is a powerful tool for initializing projects 
with predefined templates and structures.

Available Commands:
  init        Initialize a new project with your desired settings
  version     Display version information
  help        Help about any command

Use 'sova init' to create a new project with your desired settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sova.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	rootCmd.Flags().BoolP("version", "V", false, "display version information")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".sova")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if verbose {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}

func PrintSuccess(format string, a ...interface{}) {
	color.Green(format, a...)
}

func PrintInfo(format string, a ...interface{}) {
	color.Blue(format, a...)
}

func PrintWarning(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

func PrintError(format string, a ...interface{}) {
	color.Red(format, a...)
}

func GetTemplatesDir() string {
	execPath, err := os.Executable()
	if err != nil {
		return findTemplatesInWorkingDir()
	}

	execDir := filepath.Dir(execPath)
	templateDirs := []string{
		filepath.Join(execDir, "templates"),
		filepath.Join(execDir, "..", "templates"),
		filepath.Join(execDir, "..", "..", "templates"),
	}

	for _, dir := range templateDirs {
		if isValidTemplateDir(dir) {
			return dir
		}
	}

	return findTemplatesInWorkingDir()
}

func findTemplatesInWorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		return "templates"
	}

	templateDirs := []string{
		filepath.Join(wd, "templates"),
		filepath.Join(wd, "..", "templates"),
		filepath.Join(wd, "..", "..", "templates"),
	}

	for _, dir := range templateDirs {
		if isValidTemplateDir(dir) {
			return dir
		}
	}

	return "templates"
}

func isValidTemplateDir(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		return false
	}

	requiredDirs := []string{"api", "cli"}
	for _, subdir := range requiredDirs {
		subdirPath := filepath.Join(dir, subdir)
		if _, err := os.Stat(subdirPath); err != nil {
			return false
		}
	}

	return true
}
