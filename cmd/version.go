package cmd

import (
	"fmt"

	"github.com/bm611/go-ph/internal/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long: `Display version information for go-ph including version number,
build information, and other metadata.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := config.GetBuildInfo()

		fmt.Printf("%s version %s\n", buildInfo.AppName, buildInfo.Version)
		fmt.Printf("Description: %s\n", config.Description)
		fmt.Printf("Author: %s\n", config.Author)
		fmt.Printf("Repository: %s\n", config.Repository)

		if buildInfo.BuildTime != "unknown" {
			fmt.Printf("Build Time: %s\n", buildInfo.BuildTime)
		}

		if buildInfo.GitCommit != "unknown" {
			fmt.Printf("Git Commit: %s\n", buildInfo.GitCommit)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
