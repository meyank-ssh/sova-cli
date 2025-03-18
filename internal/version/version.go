package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of Sova CLI
	Version = "dev"

	// BuildDate is the date when the binary was built
	BuildDate = "unknown"

	// GitCommit is the git commit hash
	GitCommit = "unknown"

	// GoVersion is the version of Go used to build the binary
	GoVersion = runtime.Version()

	// Platform is the operating system and architecture
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

// Info returns version information
type Info struct {
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
	GitCommit string `json:"git_commit"`
	GoVersion string `json:"go_version"`
	Platform  string `json:"platform"`
}

// GetInfo returns the version information
func GetInfo() Info {
	return Info{
		Version:   Version,
		BuildDate: BuildDate,
		GitCommit: GitCommit,
		GoVersion: GoVersion,
		Platform:  Platform,
	}
}
