package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the version of the application (injected at build time)
	Version = "dev"
	// Commit is the git commit hash (injected at build time)
	Commit = "none"
	// Date is the build date (injected at build time)
	Date = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  "Print the version, commit hash, and build date of dock-docs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(stdout, "dock-docs %s\n", Version)
		fmt.Fprintf(stdout, "  commit: %s\n", Commit)
		fmt.Fprintf(stdout, "  built:  %s\n", Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
