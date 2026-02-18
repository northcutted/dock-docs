package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/northcutted/dock-docs/pkg/installer"
)

var (
	setupDir   string
	setupForce bool
	setupCheck bool
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Install external tool dependencies (syft, grype, dive)",
	Long: `Downloads and installs the external tools required for image analysis.

Tools installed:
  - syft   (SBOM generation)       from anchore/syft
  - grype  (vulnerability scanning) from anchore/grype
  - dive   (layer efficiency)       from wagoodman/dive

Binaries are downloaded from upstream GitHub Releases and placed in
~/.dock-docs/bin/ by default. System-installed tools (found in PATH)
are always preferred; the install directory serves as a fallback.`,
	Example: `  # Install all missing tools
  dock-docs setup

  # Reinstall even if already present
  dock-docs setup --force

  # Install to a custom directory
  dock-docs setup --dir /usr/local/bin

  # Check status only (no downloads)
  dock-docs setup --check`,
	RunE: runSetup,
}

func init() {
	setupCmd.Flags().StringVar(&setupDir, "dir", "", "Install directory (default: ~/.dock-docs/bin)")
	setupCmd.Flags().BoolVar(&setupForce, "force", false, "Reinstall tools even if already present")
	setupCmd.Flags().BoolVar(&setupCheck, "check", false, "Show tool status without installing")

	rootCmd.AddCommand(setupCmd)
}

func runSetup(cmd *cobra.Command, args []string) error {
	// Resolve install directory
	installDir := setupDir
	if installDir == "" {
		dir, err := installer.DefaultInstallDir()
		if err != nil {
			return fmt.Errorf("failed to determine install directory: %w", err)
		}
		installDir = dir
	}

	// --check: show status and exit
	if setupCheck {
		return printToolStatus(installDir)
	}

	// Create directory if needed
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", installDir, err)
	}

	slog.Info("install directory", "path", installDir)

	// Install missing (or all if --force)
	if err := installer.InstallAll(installDir, setupForce); err != nil {
		return err
	}

	// Print final status
	fmt.Fprintln(stdout)
	return printToolStatus(installDir)
}

func printToolStatus(installDir string) error {
	statuses := installer.Status(installDir)

	fmt.Fprintln(stdout, "Tool Status:")
	for _, s := range statuses {
		if s.Installed {
			fmt.Fprintf(stdout, "  [OK]      %-6s  (%s: %s)\n", s.Name, s.Source, s.Path)
		} else {
			fmt.Fprintf(stdout, "  [MISSING] %-6s\n", s.Name)
		}
	}
	return nil
}
