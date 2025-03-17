package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	Version   = "0.1.0"
	BuildDate = "unknown"
	GitCommit = "unknown"
	GoVersion = runtime.Version()
	Platform  = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sova CLI",
	Long:  `Display the version, build, and runtime information for Sova CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sova CLI v%s\n", Version)

		if cmd.Flag("verbose").Value.String() == "true" {
			fmt.Printf("Build Date: %s\n", BuildDate)
			fmt.Printf("Git Commit: %s\n", GitCommit)
			fmt.Printf("Go Version: %s\n", GoVersion)
			fmt.Printf("Platform: %s\n", Platform)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
