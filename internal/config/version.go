package config

// Version information
const (
	// Version is the current version of go-ph
	Version = "1.0.0"

	// AppName is the application name
	AppName = "go-ph"

	// Description is the application description
	Description = "CLI tool for fetching Product Hunt launches"

	// Author is the application author
	Author = "go-ph contributors"

	// Repository is the source code repository
	Repository = "https://github.com/bm611/go-ph"
)

// BuildInfo contains build-time information
type BuildInfo struct {
	Version   string
	AppName   string
	BuildTime string
	GitCommit string
}

// GetBuildInfo returns the current build information
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   Version,
		AppName:   AppName,
		BuildTime: "unknown", // This can be set during build time
		GitCommit: "unknown", // This can be set during build time
	}
}
