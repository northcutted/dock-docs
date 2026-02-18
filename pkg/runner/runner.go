package runner

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"time"

	"github.com/northcutted/dock-docs/pkg/installer"
)

// Command timeout defaults for external tool invocations.
// These prevent indefinite hangs if a tool stalls.
const (
	// TimeoutInspect is the timeout for quick container runtime commands
	// (docker inspect, docker manifest inspect).
	TimeoutInspect = 30 * time.Second

	// TimeoutScan is the timeout for longer-running analysis tools
	// (syft, grype, dive).
	TimeoutScan = 5 * time.Minute
)

// lookupTool resolves the path to an external tool binary. It checks
// the system PATH first and falls back to the dock-docs install
// directory (~/.dock-docs/bin/).
var lookupTool = func(name string) (string, error) {
	path, _, err := installer.FindTool(name)
	return path, err
}

// runCommand executes a command and handles verbose logging and error reporting
func runCommand(cmd *exec.Cmd, verbose bool) ([]byte, error) {
	if verbose {
		slog.Debug("running command", "cmd", cmd.String())
	}

	// Capture stdout and stderr separately if possible, or combined?
	// exec.Command.Output() captures stdout. Stderr is in ExitError.
	// But sometimes tools write useful info to stderr even on success.
	// Let's capture both if we can, but Output() is easiest for data.

	output, err := cmd.Output()
	if err != nil {
		var stderr []byte
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			stderr = exitErr.Stderr
		}
		return nil, fmt.Errorf("command failed: %w\nStderr: %s", err, string(stderr))
	}

	if verbose {
		// Print a snippet or full output? Full is safer for debug.
		slog.Debug("command output", "bytes", len(output), "output", string(output))
	}

	return output, nil
}

// EnsureImage checks if an image exists locally, and pulls it if not.
// The provided context is used as the parent for command timeouts.
func EnsureImage(ctx context.Context, image string, verbose bool) error {
	// Detect which container runtime is available
	binary := ""
	if _, err := exec.LookPath("docker"); err == nil {
		binary = "docker"
	} else if _, err := exec.LookPath("podman"); err == nil {
		binary = "podman"
	} else {
		return fmt.Errorf("no container runtime found (docker or podman)")
	}

	// Check if image exists
	inspectCtx, cancel := context.WithTimeout(ctx, TimeoutInspect)
	defer cancel()
	checkCmd := exec.CommandContext(inspectCtx, binary, "inspect", "--type=image", image)
	if err := checkCmd.Run(); err == nil {
		if verbose {
			slog.Debug("image found locally", "image", image)
		}
		return nil
	}

	// Image not found, pull it (pull can be slow, use scan timeout)
	slog.Info("pulling image", "image", image)
	pullCtx, pullCancel := context.WithTimeout(ctx, TimeoutScan)
	defer pullCancel()
	pullCmd := exec.CommandContext(pullCtx, binary, "pull", image)
	if _, err := runCommand(pullCmd, verbose); err != nil {
		return fmt.Errorf("failed to pull image %s: %w", image, err)
	}
	return nil
}
