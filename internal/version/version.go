// Package version holds the CLI version. Releases set it with
// -ldflags "-X github.com/yesly-dev/yesly-cli/internal/version.Version=<v>".
package version

// Version of this build.
var Version = "0.2.0"

// Commit and Date of this build, set by releases.
var (
	Commit = ""
	Date   = ""
)
