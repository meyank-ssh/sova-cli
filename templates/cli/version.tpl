package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the version of the application
	Version = "0.1.0"
	// BuildDate is the date when the application was built
	BuildDate = "unknown"
	// GitCommit is the git commit hash
	GitCommit = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, build date, and git commit hash of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("{{.ProjectName}} version %s\n", Version)
		fmt.Printf("Built on %s\n", BuildDate)
		fmt.Printf("Git commit: %s\n", GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
} 