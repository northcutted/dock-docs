package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/northcutted/dock-docs/pkg/types"
)

// DiveRunner runs 'dive <image> --json output.json'
type DiveRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *DiveRunner) Name() string { return "dive" }

// IsAvailable checks whether the dive binary is installed.
func (r *DiveRunner) IsAvailable() bool {
	if path, err := lookupTool("dive"); err == nil {
		r.binary = path
		return true
	}
	return false
}

// detectPodmanSocket attempts to detect the Podman machine socket path.
// It returns a DOCKER_HOST value (e.g. "unix:///path/to/socket") if found,
// or an empty string if detection fails.
func detectPodmanSocket(ctx context.Context) string {
	socketCtx, cancel := context.WithTimeout(ctx, TimeoutInspect)
	defer cancel()
	pCmd := exec.CommandContext(socketCtx, "podman", "machine", "inspect")
	out, err := pCmd.Output()
	if err != nil {
		return ""
	}

	var machines []struct {
		ConnectionInfo struct {
			PodmanSocket *struct {
				Path string `json:"Path"`
			} `json:"PodmanSocket"`
		} `json:"ConnectionInfo"`
	}
	if json.Unmarshal(out, &machines) != nil || len(machines) == 0 {
		return ""
	}

	if machines[0].ConnectionInfo.PodmanSocket == nil || machines[0].ConnectionInfo.PodmanSocket.Path == "" {
		return ""
	}

	socketPath := machines[0].ConnectionInfo.PodmanSocket.Path
	if !strings.HasPrefix(socketPath, "unix://") {
		socketPath = "unix://" + socketPath
	}
	return socketPath
}

// Run executes dive against the given image and parses the efficiency results.
// The provided context is used as the parent for the command timeout.
func (r *DiveRunner) Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error) {
	// Create a temp file for output
	tmpFile, err := os.CreateTemp("", "dive-output-*.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		// Attempt to remove file, ignore error (best effort cleanup)
		_ = os.Remove(tmpFile.Name()) //nolint:gosec // path from os.CreateTemp, not user input
	}()
	// Close immediately, dive will write to it. Ignore error.
	_ = tmpFile.Close()

	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("dive not found")
		}
	}

	runCtx, cancel := context.WithTimeout(ctx, TimeoutScan)
	defer cancel()
	cmd := exec.CommandContext(runCtx, r.binary, image, "--json", tmpFile.Name()) //nolint:gosec // binary resolved from trusted lookup

	// Podman Support: If docker is missing but podman is present and DOCKER_HOST
	// is not set, try to detect the Podman machine socket automatically.
	if _, err := exec.LookPath("docker"); err != nil {
		if _, err := exec.LookPath("podman"); err == nil && os.Getenv("DOCKER_HOST") == "" {
			if socketPath := detectPodmanSocket(ctx); socketPath != "" {
				env := os.Environ()
				env = append(env, "DOCKER_HOST="+socketPath)
				cmd.Env = env
			}
		}
	}

	// Dive writes to file, but might output logs to stdout/stderr. capture or ignore?
	// cmd.CombinedOutput() might be useful for debugging if it fails.

	// Dive uses CombinedOutput because it writes analysis logs to stdout/stderr even with --json
	// But our runCommand assumes Output().
	// We can use runCommand here if we want consistent logging, but dive output (logs) is not the JSON.
	// The JSON is in the file.

	if verbose {
		slog.Debug("running command", "cmd", cmd.String()) //nolint:gosec // G706: cmd is constructed by us, not user input
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("dive failed: %w, output: %s", err, string(output))
	}

	if verbose {
		slog.Debug("command output", "bytes", len(output), "output", string(output)) //nolint:gosec // G706: output is from a tool we invoked
	}

	content, err := os.ReadFile(tmpFile.Name()) //nolint:gosec // path from os.CreateTemp, not user input
	if err != nil {
		return nil, fmt.Errorf("failed to read dive output: %w", err)
	}

	return parseDiveOutput(content)
}

// parseDiveOutput parses JSON output from dive's analysis file
// into ImageStats containing efficiency score and wasted bytes.
func parseDiveOutput(content []byte) (*types.ImageStats, error) {
	var diveOutput struct {
		Image struct {
			InefficientBytes uint64  `json:"inefficientBytes"`
			EfficiencyScore  float64 `json:"efficiencyScore"`
		} `json:"image"`
	}

	if err := json.Unmarshal(content, &diveOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dive output: %w", err)
	}

	stats := &types.ImageStats{
		Efficiency:  diveOutput.Image.EfficiencyScore * 100,
		WastedBytes: fmt.Sprintf("%.2f MB", float64(diveOutput.Image.InefficientBytes)/1024/1024),
	}

	return stats, nil
}
