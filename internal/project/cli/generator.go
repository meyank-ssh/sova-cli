package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-sova/sova-cli/pkg/questions"
)

type CLIProjectGenerator struct {
	ProjectName string
	ProjectDir  string
	Answers     *questions.ProjectAnswers
}

func NewCLIProjectGenerator(projectName, projectDir string, answers *questions.ProjectAnswers) *CLIProjectGenerator {
	return &CLIProjectGenerator{
		ProjectName: projectName,
		ProjectDir:  projectDir,
		Answers:     answers,
	}
}

func (g *CLIProjectGenerator) Generate() (map[string]string, []string, error) {
	dirs := []string{
		"cmd",
		"internal",
		"test",
	}

	files := make(map[string]string)

	files["cmd/root.go"] = g.generateRootCmdFile()

	files["cmd/command1.go"] = g.generateCommand1File()
	files["cmd/command2.go"] = g.generateCommand2File()

	files["internal/utils.go"] = g.generateUtilsFile()
	files["internal/config.go"] = g.generateConfigFile()

	files["main.go"] = g.generateMainFile()

	files["README.md"] = g.generateReadmeFile()

	files["go.mod"] = g.generateGoModFile()

	return files, dirs, nil
}

func (g *CLIProjectGenerator) generateRootCmdFile() string {
	return fmt.Sprintf(`package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "%s",
	Short: "A brief description of your application",
	Long: %s,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.%s.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".%s" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".%s")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
`, g.ProjectName, "`A longer description that spans multiple lines and likely contains\nexamples and usage of using your application. For example:\n\nCobra is a CLI library for Go that empowers applications.\nThis application is a tool to generate the needed files\nto quickly create a Cobra application.`", g.ProjectName, g.ProjectName, g.ProjectName)
}

func (g *CLIProjectGenerator) generateCommand1File() string {
	return fmt.Sprintf(`package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// command1Cmd represents the command1 command
var command1Cmd = &cobra.Command{
	Use:   "command1",
	Short: "A brief description of your command",
	Long: %s,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command1 called")
	},
}

func init() {
	rootCmd.AddCommand(command1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
`, "`A longer description that spans multiple lines and likely contains examples\nand usage of using your command. For example:\n\nCobra is a CLI library for Go that empowers applications.\nThis application is a tool to generate the needed files\nto quickly create a Cobra application.`")
}

func (g *CLIProjectGenerator) generateCommand2File() string {
	return fmt.Sprintf(`package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// command2Cmd represents the command2 command
var command2Cmd = &cobra.Command{
	Use:   "command2",
	Short: "A brief description of your command",
	Long: %s,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command2 called")
	},
}

func init() {
	rootCmd.AddCommand(command2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
`, "`A longer description that spans multiple lines and likely contains examples\nand usage of using your command. For example:\n\nCobra is a CLI library for Go that empowers applications.\nThis application is a tool to generate the needed files\nto quickly create a Cobra application.`")
}

func (g *CLIProjectGenerator) generateUtilsFile() string {
	return `package internal

import (
	"fmt"
	"strings"
)

// Constants for terminal colors
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
)

// PrintInfo prints an info message to the console
func PrintInfo(format string, a ...interface{}) {
	fmt.Printf(ColorBlue+"INFO: "+format+ColorReset+"\n", a...)
}

// PrintSuccess prints a success message to the console
func PrintSuccess(format string, a ...interface{}) {
	fmt.Printf(ColorGreen+"SUCCESS: "+format+ColorReset+"\n", a...)
}

// PrintWarning prints a warning message to the console
func PrintWarning(format string, a ...interface{}) {
	fmt.Printf(ColorYellow+"WARNING: "+format+ColorReset+"\n", a...)
}

// PrintError prints an error message to the console
func PrintError(format string, a ...interface{}) {
	fmt.Printf(ColorRed+"ERROR: "+format+ColorReset+"\n", a...)
}

// StringInSlice checks if a string is in a slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
`
}

func (g *CLIProjectGenerator) generateConfigFile() string {
	return fmt.Sprintf(`package internal

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	// Add your configuration fields here
	LogLevel string
	APIKey   string
}

// LoadConfig loads the configuration from disk
func LoadConfig(configPath string) (*Config, error) {
	// Set default configuration file path if not provided
	if configPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configPath = filepath.Join(home, ".%s.yaml")
	}

	// Initialize config
	viper.SetConfigFile(configPath)

	// Set defaults
	viper.SetDefault("LogLevel", "info")
	viper.SetDefault("APIKey", "")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if the config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// Create a new config struct
	config := &Config{}

	// Unmarshal the config
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
`, g.ProjectName)
}

func (g *CLIProjectGenerator) generateMainFile() string {
	return fmt.Sprintf(`package main

import "%s/cmd"

func main() {
	cmd.Execute()
}
`, g.ProjectName)
}

func (g *CLIProjectGenerator) generateReadmeFile() string {
	installCmd := fmt.Sprintf("go get -u github.com/yourusername/%s", g.ProjectName)

	return fmt.Sprintf(`# %s

A CLI application built with Go and Cobra.

## Installation

`+"```bash"+`
%s
`+"```"+`

## Usage

`+"```bash"+`
%s [command]
`+"```"+`

### Available Commands:

* command1: Description of command1
* command2: Description of command2
* help: Help about any command

### Flags:

* --config string: config file (default is $HOME/.%s.yaml)
* -h, --help: help for %s
* -t, --toggle: Help message for toggle

Use "%s [command] --help" for more information about a command.

## Development

1. Clone the repository
2. Install dependencies with `+"`go mod tidy`"+`
3. Run with `+"`go run main.go`"+`

## Building

Build a binary with:

`+"```bash"+`
go build -o %s
`+"```"+`

## License

This project is licensed under the MIT License - see the LICENSE file for details.
`, g.ProjectName, installCmd, g.ProjectName, g.ProjectName, g.ProjectName, g.ProjectName, g.ProjectName)
}

func (g *CLIProjectGenerator) generateGoModFile() string {
	return fmt.Sprintf(`module %s

go 1.21

require (
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.1
)
`, g.ProjectName)
}

func (g *CLIProjectGenerator) WriteFiles(files map[string]string) error {
	for filename, content := range files {
		filePath := filepath.Join(g.ProjectDir, filename)

		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filename, err)
		}

		fmt.Printf("Created file: %s\n", filePath)
	}

	return nil
}
