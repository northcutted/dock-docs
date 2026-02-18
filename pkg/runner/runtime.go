package runner

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/northcutted/dock-docs/pkg/types"
)

// RuntimeRunner runs 'docker inspect' or 'podman inspect'
type RuntimeRunner struct {
	binary string
}

// Name returns the display name for this runner.
func (r *RuntimeRunner) Name() string {
	if r.binary != "" {
		return r.binary
	}
	return "runtime"
}

// IsAvailable checks whether a container runtime (docker or podman) is installed.
func (r *RuntimeRunner) IsAvailable() bool {
	// Check docker first
	if _, err := exec.LookPath("docker"); err == nil {
		r.binary = "docker"
		return true
	}
	// Fallback to podman
	if _, err := exec.LookPath("podman"); err == nil {
		r.binary = "podman"
		return true
	}
	return false
}

// Run executes 'docker inspect' or 'podman inspect' and parses the result.
// The provided context is used as the parent for the command timeout.
func (r *RuntimeRunner) Run(ctx context.Context, image string, verbose bool) (*types.ImageStats, error) {
	// Ensure binary is set if IsAvailable wasn't called (though it should be)
	if r.binary == "" {
		if !r.IsAvailable() {
			return nil, fmt.Errorf("no container runtime found (docker or podman)")
		}
	}

	runCtx, cancel := context.WithTimeout(ctx, TimeoutInspect)
	defer cancel()
	cmd := exec.CommandContext(runCtx, r.binary, "inspect", image)
	output, err := runCommand(cmd, verbose)
	if err != nil {
		return nil, err
	}

	return parseRuntimeInspect(output, image, r.binary)
}

// parseRuntimeInspect parses JSON output from 'docker inspect' or 'podman inspect'
// into ImageStats containing architecture, OS, size, and layer count.
func parseRuntimeInspect(output []byte, image string, binary string) (*types.ImageStats, error) {
	var inspect []struct {
		Architecture string `json:"Architecture"`
		Os           string `json:"Os"`
		Size         int64  `json:"Size"`
		RootFS       struct {
			Layers []string `json:"Layers"`
		} `json:"RootFS"`
	}

	if err := json.Unmarshal(output, &inspect); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s inspect output: %w", binary, err)
	}

	if len(inspect) == 0 {
		return nil, fmt.Errorf("no inspect data returned for image %s", image)
	}

	data := inspect[0]
	stats := &types.ImageStats{
		ImageTag:     image,
		Architecture: data.Architecture,
		OS:           data.Os,
		SizeBytes:    data.Size,
		TotalLayers:  len(data.RootFS.Layers),
	}

	return stats, nil
}
