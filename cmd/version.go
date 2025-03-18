package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/go-sova/sova-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sova CLI",
	Long:  `Display the version, build, and runtime information for Sova CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		info := version.GetInfo()

		if cmd.Flag("json").Value.String() == "true" {
			jsonOutput, _ := json.MarshalIndent(info, "", "  ")
			fmt.Println(string(jsonOutput))
			return
		}

		fmt.Printf("Sova CLI v%s\n", info.Version)
		if cmd.Flag("verbose").Value.String() == "true" {
			fmt.Printf("Build Date: %s\n", info.BuildDate)
			fmt.Printf("Git Commit: %s\n", info.GitCommit)
			fmt.Printf("Go Version: %s\n", info.GoVersion)
			fmt.Printf("Platform: %s\n", info.Platform)
		}
	},
}

func init() {
	versionCmd.Flags().BoolP("verbose", "v", false, "print detailed version information")
	versionCmd.Flags().Bool("json", false, "print version information as JSON")
	rootCmd.AddCommand(versionCmd)
}
