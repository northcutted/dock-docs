package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/northcutted/dock-docs/pkg/installer"
)

// checkToolStatus returns a string indicating the status of required tools.
func checkToolStatus() string {
	tools := []string{"syft", "grype", "dive"}
	var status strings.Builder
	status.WriteString("\nPrerequisites:\n")

	// Check for Docker or Podman
	if _, err := exec.LookPath("docker"); err == nil {
		status.WriteString("  [OK] docker\n")
	} else if _, err := exec.LookPath("podman"); err == nil {
		status.WriteString("  [OK] podman\n")
	} else {
		status.WriteString("  [MISSING] docker or podman (required for dynamic analysis)\n")
	}

	for _, tool := range tools {
		if path, source, err := installer.FindTool(tool); err == nil {
			fmt.Fprintf(&status, "  [OK] %s (%s: %s)\n", tool, source, path)
		} else {
			fmt.Fprintf(&status, "  [MISSING] %s (run 'dock-docs setup' to install)\n", tool)
		}
	}
	return status.String()
}
